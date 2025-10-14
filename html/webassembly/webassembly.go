package webassembly

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/js/arraybuffer"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var webassemblyinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if webassemblyinterface = js.Global().Get("WebAssembly"); webassemblyinterface.Error() != nil {
			webassemblyinterface = js.Undefined()
		}

		js.Register(webassemblyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return webassemblyinterface
}

// WebAssembly struct
type WebAssembly struct {
	js.Object
}

type WebAssemblyFrom interface {
	WebAssembly_() WebAssembly
}

func (w WebAssembly) WebAssembly_() WebAssembly {
	return w
}

func New() (WebAssembly, error) {

	var w WebAssembly

	if wi := GetInterface(); !wi.IsUndefined() {

		w.SetObjectValue(wi)
		return w, nil
	}
	return w, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (WebAssembly, error) {
	var w WebAssembly
	var err error
	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(wi) {
				w.SetObjectValue(obj)
			} else {
				err = ErrNotAWebAssembly
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return w, err
}

func (w WebAssembly) InstantiateStreaming(source interface{}, imports js.Value) (promise.Promise, error) {

	var obj js.Value
	var err error
	var p promise.Promise

	if s, ok := source.(js.ObjectFrom); ok {
		if obj = w.Call("instantiateStreaming", s.Value(), imports); obj.Error() == nil {
			p, err = promise.NewFromJSObject(obj)

		}
	}

	return p, err
}

func (w WebAssembly) Instantiate(source arraybuffer.ArrayBuffer, imports js.Value) (promise.Promise, error) {
	var obj js.Value
	var err error
	var p promise.Promise

	if obj = w.Call("instantiate", source.GetObjectValue(), imports); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
