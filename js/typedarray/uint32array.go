package typedarray

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonuint32array sync.Once

var uint32arrayinterface js.Value

// Uint32Array struct
type Uint32Array struct {
	TypedArray
}

type Uint32ArrayFrom interface {
	Uint32Array_() Uint32Array
}

func (u Uint32Array) Uint32Array_() Uint32Array {
	return u
}

// GetInterface get the JS interface
func GetUint32ArrayInterface() js.Value {

	singletonuint32array.Do(func() {

		if uint32arrayinterface = js.Global().Get("Uint32Array"); uint32arrayinterface.Error() != nil {
			uint32arrayinterface = js.Undefined()
		}
		js.Register(uint32arrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint32FromJSObject(v)
		})
	})

	return uint32arrayinterface
}

func NewUint32Array(value interface{}) (Uint32Array, error) {

	var a Uint32Array
	var objnew js.Value
	var err error
	if ai := GetUint32ArrayInterface(); !ai.IsUndefined() {
		if objnew = ai.New(js.ValueOf(value)); objnew.Error() == nil {
			a.SetObjectValue(objnew)
		}

	} else {
		err = ErrNotImplementedUint32Array
	}

	return a, err
}

func NewUint32ArrayFrom(iterable interface{}) (Uint32Array, error) {

	arr, err := newTypedArrayFrom(GetUint32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint32FromJSObject(v)
	}, iterable)
	return arr.(Uint32Array), err
}

func NewUint32ArrayOf(values ...interface{}) (Uint32Array, error) {

	arr, err := newTypedArrayOf(GetUint32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint32FromJSObject(v)
	}, values...)
	return arr.(Uint32Array), err
}

func NewUint32FromJSObject(obj js.Value) (Uint32Array, error) {
	var u Uint32Array
	var err error
	if ui := GetUint32ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAUint32Array
			}
		}
	} else {
		err = ErrNotImplementedUint32Array
	}

	return u, err
}
