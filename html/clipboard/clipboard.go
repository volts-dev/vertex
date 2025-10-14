package clipboard

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/clipboarditem"
	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var clipboardinterface js.Value

// Clipboard struct
type Clipboard struct {
	eventtarget.EventTarget
}

type ClipboardFrom interface {
	Clipboard_() Clipboard
}

func (c Clipboard) Clipboard_() Clipboard {
	return c
}

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if clipboardinterface = js.Global().Get("Clipboard"); clipboardinterface.Error() != nil {
			clipboardinterface = js.Undefined()
		}

		js.Register(clipboardinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})

		clipboarditem.GetInterface()
		promise.GetInterface()

	})

	return clipboardinterface
}

func NewFromJSObject(obj js.Value) (Clipboard, error) {
	var c Clipboard

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			c.SetObjectValue(obj)
			return c, nil

		}
	}

	return c, ErrNotImplemented
}

func (c Clipboard) Read() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = c.Call("read"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) ReadText() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = c.Call("readText"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) Write(data []clipboarditem.ClipboardItem) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise
	var arrayJS []interface{}

	for _, value := range data {
		arrayJS = append(arrayJS, value.GetObjectValue())
	}

	if obj = c.Call("write", arrayJS); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) WriteText(data string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = c.Call("writeText", js.ValueOf(data)); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
