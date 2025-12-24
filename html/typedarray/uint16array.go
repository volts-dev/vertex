package typedarray

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonuint16array sync.Once

var uint16arrayinterface js.Value

// Uint16Array struct
type Uint16Array struct {
	TypedArray
}

type Uint16ArrayFrom interface {
	Uint16Array_() Uint16Array
}

func (u Uint16Array) Uint16Array_() Uint16Array {
	return u
}

// GetInterface get the JS interface
func GetUint16ArrayInterface() js.Value {

	singletonuint16array.Do(func() {

		if uint16arrayinterface = js.Global().Get("Uint16Array"); uint16arrayinterface.Error() != nil {
			uint16arrayinterface = js.Undefined()
		}
		js.Register(uint16arrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint16FromJSObject(v)
		})
	})

	return uint16arrayinterface
}

func NewUint16Array(value interface{}) (Uint16Array, error) {
	var a Uint16Array
	var objnew js.Value
	var err error
	if ai := GetUint16ArrayInterface(); !ai.IsUndefined() {
		if objnew = ai.New(js.ValueOf(value)); objnew.Error() == nil {
			a.SetObjectValue(objnew)
		}
	} else {
		err = ErrNotImplementedUint16Array
	}
	return a, err
}

func NewUint16ArrayFrom(iterable interface{}) (Uint16Array, error) {

	arr, err := newTypedArrayFrom(GetUint16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint16FromJSObject(v)
	}, iterable)
	return arr.(Uint16Array), err

}

func NewUint16ArrayOf(values ...interface{}) (Uint16Array, error) {

	arr, err := newTypedArrayOf(GetUint16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint16FromJSObject(v)
	}, values...)
	return arr.(Uint16Array), err
}

func NewUint16FromJSObject(obj js.Value) (Uint16Array, error) {
	var u Uint16Array
	var err error
	if ui := GetUint16ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAUint16Array
			}
		}
	} else {
		err = ErrNotImplementedUint16Array
	}

	return u, err
}
