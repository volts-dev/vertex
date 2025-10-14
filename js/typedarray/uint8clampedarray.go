package typedarray

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonuint8campledarray sync.Once

var uint8clampedarrayinterface js.Value

// Uint8ClampedArray struct
type Uint8ClampedArray struct {
	TypedArray
}

type Uint8ClampedArrayFrom interface {
	Uint8ClampedArray_() Uint8ClampedArray
}

func (u Uint8ClampedArray) Uint8ClampedArray_() Uint8ClampedArray {
	return u
}

// GetUint8ClampedArrayInterface get the JS interface of GetUint8ClampedArrayInterface
func GetUint8ClampedArrayInterface() js.Value {

	singletonuint8campledarray.Do(func() {

		if uint8clampedarrayinterface = js.Global().Get("Uint8ClampedArray"); uint8clampedarrayinterface.Error() != nil {
			uint8clampedarrayinterface = js.Undefined()
		}
		js.Register(uint8clampedarrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint8ClampedFromJSObject(v)
		})
	})

	return uint8clampedarrayinterface
}

func NewUint8ClampedArray(value interface{}) (Uint8ClampedArray, error) {
	var a Uint8ClampedArray
	var objnew js.Value
	var err error
	if ai := GetUint8ClampedArrayInterface(); !ai.IsUndefined() {
		if objnew = ai.New(js.ValueOf(value)); objnew.Error() == nil {
			a.SetObjectValue(objnew)
		}
	} else {
		err = ErrNotImplementedUint8ClampedArray
	}
	return a, err
}

func NewUint8ClampedArrayFrom(iterable interface{}) (Uint8ClampedArray, error) {

	arr, err := newTypedArrayFrom(GetUint8ClampedArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8ClampedFromJSObject(v)
	}, iterable)
	return arr.(Uint8ClampedArray), err
}

func NewUint8ClampedArrayOf(values ...interface{}) (Uint8ClampedArray, error) {

	arr, err := newTypedArrayOf(GetUint8ClampedArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8ClampedFromJSObject(v)
	}, values...)
	return arr.(Uint8ClampedArray), err
}

func NewUint8ClampedFromJSObject(obj js.Value) (Uint8ClampedArray, error) {
	var u Uint8ClampedArray
	var err error
	if ui := GetUint8ClampedArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAUint8ClampedArray
			}
		}
	} else {
		err = ErrNotImplementedUint8ClampedArray
	}

	return u, err
}
