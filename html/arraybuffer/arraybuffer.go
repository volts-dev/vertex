package arraybuffer

//partial implemented (herited from function)
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var arraybufferinterface js.Value

// GetInterface get the JS interface ArrayBuffer
func GetInterface() js.Value {
	singleton.Do(func() {
		if arraybufferinterface = js.Global().Get("ArrayBuffer"); arraybufferinterface.Error() != nil {
			arraybufferinterface = js.Undefined()
		}
		js.Register(arraybufferinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return arraybufferinterface
}

// ArrayBuffer struct
type ArrayBuffer struct {
	js.Object
}

type ArrayBufferFrom interface {
	ArrayBuffer_() ArrayBuffer
}

func (a ArrayBuffer) ArrayBuffer_() ArrayBuffer {
	return a
}

func New(size int) (ArrayBuffer, error) {

	var a ArrayBuffer
	var obj js.Value
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.New(js.ValueOf(size)); obj.Error() == nil {
			a.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func NewFromJSObject(obj js.Value) (ArrayBuffer, error) {
	var a ArrayBuffer
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnArrayBuffer
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a ArrayBuffer) ByteLength() (int64, error) {

	return a.GetAttributeInt64("byteLength")
}

func (a ArrayBuffer) Slice(begin int, end ...int) (ArrayBuffer, error) {

	var optjs []interface{}
	var err error
	var obj js.Value
	var ret ArrayBuffer

	optjs = append(optjs, js.ValueOf(begin))
	if len(end) > 0 {
		optjs = append(optjs, js.ValueOf(end[0]))
	}

	if obj = a.Call("slice", optjs...); obj.Error() == nil {

		ret, err = NewFromJSObject(obj)
	}
	return ret, err
}

func IsView(i interface{}) (bool, error) {
	var ret bool
	var err error
	var obj js.Value

	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.Call("isView", js.ValueOf(i)); obj.Error() == nil {

			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				err = js.ErrObjectNotBool
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return ret, err
}
