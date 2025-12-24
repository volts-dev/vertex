package typedarray

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonFloat64array sync.Once

var Float64arrayinterface js.Value

// Float64Array struct
type Float64Array struct {
	TypedArray
}

type Float64ArrayFrom interface {
	Float64Array_() Float64Array
}

func (u Float64Array) Float64Array_() Float64Array {
	return u
}

// GetFloat64ArrayInterface get the JS interface of Float64Array
func GetFloat64ArrayInterface() js.Value {

	singletonFloat64array.Do(func() {
		if Float64arrayinterface = js.Global().Get("Float64Array"); Float64arrayinterface.Error() != nil {
			Float64arrayinterface = js.Undefined()
		}
		js.Register(Float64arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFloat64FromJSObject(v)
		})
	})

	return Float64arrayinterface
}

func NewFloat64Array(value interface{}) (Float64Array, error) {

	var a Float64Array
	var objnew js.Value
	var err error
	if ai := GetFloat64ArrayInterface(); !ai.IsUndefined() {
		if objnew = ai.New(js.ValueOf(value)); objnew.Error() == nil {
			a.SetObjectValue(objnew)
		}

	} else {
		err = ErrNotImplementedFloat64Array
	}

	return a, err
}

func NewFloat64ArrayFrom(iterable interface{}) (Float64Array, error) {

	arr, err := newTypedArrayFrom(GetFloat64ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat64FromJSObject(v)
	}, iterable)
	return arr.(Float64Array), err
}

func NewFloat64ArrayOf(values ...interface{}) (Float64Array, error) {

	arr, err := newTypedArrayOf(GetFloat64ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat64FromJSObject(v)
	}, values...)
	return arr.(Float64Array), err
}

func NewFloat64FromJSObject(obj js.Value) (Float64Array, error) {
	var u Float64Array
	var err error
	if ui := GetFloat64ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAFloat64Array
			}
		}
	} else {
		err = ErrNotImplementedFloat64Array
	}

	return u, err
}
