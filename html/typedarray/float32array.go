package typedarray

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonFloat32array sync.Once

var Float32arrayinterface js.Value

// Float32Array struct
type Float32Array struct {
	TypedArray
}

type Float32ArrayFrom interface {
	Float32Array_() Float32Array
}

func (u Float32Array) Float32Array_() Float32Array {
	return u
}

// GetFloat32ArrayInterface get the JS interface of Float32Array
func GetFloat32ArrayInterface() js.Value {

	singletonFloat32array.Do(func() {

		if Float32arrayinterface = js.Global().Get("Float32Array"); Float32arrayinterface.Error() != nil {
			Float32arrayinterface = js.Undefined()
		}
		js.Register(Float32arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFloat32FromJSObject(v)
		})
	})

	return Float32arrayinterface
}

func NewFloat32Array(value interface{}) (Float32Array, error) {

	var a Float32Array
	var objnew js.Value
	var err error
	if ai := GetFloat32ArrayInterface(); !ai.IsUndefined() {
		if objnew = ai.New(js.ValueOf(value)); objnew.Error() == nil {
			a.SetObjectValue(objnew)
		}

	} else {
		err = ErrNotImplementedFloat32Array
	}

	return a, err
}

func NewFloat32ArrayFrom(iterable interface{}) (Float32Array, error) {

	arr, err := newTypedArrayFrom(GetFloat32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat32FromJSObject(v)
	}, iterable)
	return arr.(Float32Array), err
}

func NewFloat32ArrayOf(values ...interface{}) (Float32Array, error) {

	arr, err := newTypedArrayOf(GetFloat32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat32FromJSObject(v)
	}, values...)
	return arr.(Float32Array), err
}

func NewFloat32FromJSObject(obj js.Value) (Float32Array, error) {
	var u Float32Array
	var err error
	if ui := GetFloat32ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAFloat32Array
			}
		}
	} else {
		err = ErrNotImplementedFloat32Array
	}

	return u, err
}
