package typedarray

import (
	"errors"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/array"
	"github.com/volts-dev/vertex/js/arraybuffer"
)

func init() {

	initinterface.RegisterInterface(GetFloat32ArrayInterface)
	initinterface.RegisterInterface(GetFloat64ArrayInterface)
	initinterface.RegisterInterface(GetInt8ArrayInterface)
	initinterface.RegisterInterface(GetInt16ArrayInterface)
	initinterface.RegisterInterface(GetInt32ArrayInterface)
	initinterface.RegisterInterface(GetUint8ArrayInterface)
	initinterface.RegisterInterface(GetUint8ClampedArrayInterface)
	initinterface.RegisterInterface(GetUint16ArrayInterface)
	initinterface.RegisterInterface(GetUint32ArrayInterface)

}

// TypedArray struct
type TypedArray struct {
	array.Array
}

type TypedArrayFrom interface {
	TypedArray_() TypedArray
}

func (t TypedArray) TypedArray_() TypedArray {
	return t
}

func (t TypedArray) Bytes() ([]byte, error) {
	var err error
	var buffer []byte
	var l int
	if l, err = t.Length(); err == nil {
		buffer = make([]byte, l)
		if _, err = js.CopyBytesToGo(buffer, t.GetObjectValue()); err == nil {
			return buffer, nil
		}
	}

	return buffer, err
}

func (t TypedArray) CopyBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return js.CopyBytesToGo(buffer, t.GetObjectValue())

}

func (t TypedArray) CopyFromBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return js.CopyBytesToJS(t.GetObjectValue(), buffer)
}

func (t TypedArray) Buffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var obj js.Value

	if obj = t.GetValueByKey("buffer"); obj.Error() == nil {
		return arraybuffer.NewFromJSObject(obj)

	}

	return arraybuffer.ArrayBuffer{}, err
}

func (t TypedArray) ByteLength() (int64, error) {

	return t.GetAttributeInt64("byteLength")
}

func (t TypedArray) ByteOffset() (int64, error) {

	return t.GetAttributeInt64("byteOffset")
}

func (t TypedArray) BYTES_PER_ELEMENT() (int, error) {

	return t.GetAttributeInt("BYTES_PER_ELEMENT")
}

func (t TypedArray) Subarray(opts ...int) (interface{}, error) {

	var err error
	var arrayJS []interface{}
	var obj js.Value
	var newArr interface{}

	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if obj = t.Call("subarray", arrayJS...); obj.Error() == nil {
		newArr, err = js.Discover(obj)

	}

	return newArr, err
}

func newTypedArrayFrom(interfacejs js.Value, f func(js.Value) (interface{}, error), iterable interface{}) (interface{}, error) {
	var obj js.Value
	var err error
	var newArr interface{}
	if obj = interfacejs.Call("from", js.ValueOf(iterable)); obj.Error() == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}

func newTypedArrayOf(interfacejs js.Value, f func(js.Value) (interface{}, error), values ...interface{}) (interface{}, error) {

	var newArr interface{}
	var arrayJS []interface{}
	var obj js.Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if obj = interfacejs.Call("of", arrayJS...); obj.Error() == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}
