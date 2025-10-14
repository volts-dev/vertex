package js

import "errors"

//partial implemented (herited from function)
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer

func init() {

	RegisterInterface(GetArrayBufferInterface)
}

var ErrNotAnArrayBuffer = errors.New("The given value must be an arrayBuffer")

var arraybufferinterface Value

// GetArrayBufferInterface get the JS interface ArrayBuffer
func GetArrayBufferInterface() Value {

	singleton.Do(func() {
		if arraybufferinterface = Global().Get("ArrayBuffer"); arraybufferinterface.Error() != nil {
			arraybufferinterface = Undefined()
		}
		Register(arraybufferinterface, func(v Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return arraybufferinterface
}

// ArrayBuffer struct
type ArrayBuffer struct {
	Object
}

type ArrayBufferFrom interface {
	ArrayBuffer_() ArrayBuffer
}

func (a ArrayBuffer) ArrayBuffer_() ArrayBuffer {
	return a
}

func NewArrayBuffer(size int) (ArrayBuffer, error) {

	var a ArrayBuffer
	var obj Value
	var err error
	if ai := GetArrayBufferInterface(); !ai.IsUndefined() {

		if obj = ai.New(ValueOf(size)); obj.Error() == nil {
			//a.BaseObject = a.SetObject(obj)
			a.SetValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func ToArrayBuffer(obj Value) (ArrayBuffer, error) {
	var a ArrayBuffer
	var err error
	if ai := GetArrayBufferInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				//a.BaseObject = a.SetObject(obj)
				a.SetValue(obj)
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
	var obj Value
	var ret ArrayBuffer

	optjs = append(optjs, ValueOf(begin))
	if len(end) > 0 {
		optjs = append(optjs, ValueOf(end[0]))
	}

	if obj = a.Call("slice", optjs...); obj.Error() == nil {
		ret, err = ToArrayBuffer(obj)
	}
	return ret, err
}

func IsView(i interface{}) (bool, error) {
	var ret bool
	var err error
	var obj Value

	if ai := GetArrayBufferInterface(); !ai.IsUndefined() {

		if obj = ai.Call("isView", ValueOf(i)); obj.Error() == nil {

			if obj.Type() == TypeBoolean {
				return obj.Bool()
			} else {
				err = ErrObjectNotBool
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return ret, err
}
