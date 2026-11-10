package typedarray

import (
	"errors"

	"github.com/volts-dev/vertex/html/arraybuffer"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetFloat32ArrayInterface)
	js.RegisterInterface(GetFloat64ArrayInterface)
	js.RegisterInterface(GetInt8ArrayInterface)
	js.RegisterInterface(GetInt16ArrayInterface)
	js.RegisterInterface(GetInt32ArrayInterface)
	js.RegisterInterface(GetUint8ArrayInterface)
	js.RegisterInterface(GetUint8ClampedArrayInterface)
	js.RegisterInterface(GetUint16ArrayInterface)
	js.RegisterInterface(GetUint32ArrayInterface)

}

// TypedArray struct
type TypedArray struct {
	js.Array
}

type TypedArrayFrom interface {
	TypedArray_() TypedArray
}

func (t TypedArray) TypedArray_() TypedArray {
	return t
}

func (t TypedArray) Bytes() []byte {
	var err error
	var buffer []byte
	buffer = make([]byte, t.Length())
	if _, err = js.CopyBytesToGo(buffer, t.GetObjectValue()); err == nil {
		return buffer
	}

	return buffer
}

func (t TypedArray) CopyBytes(buffer []byte) (int, error) {
	var l int
	l = t.Length()
	if len(buffer) < l {
		return 0, errors.New("Increase your buffer size")
	}

	return js.CopyBytesToGo(buffer, t.GetObjectValue())
}

func (t TypedArray) CopyFromBytes(buffer []byte) (int, error) {

	if len(buffer) < t.Length() {
		return 0, errors.New("Increase your buffer size")
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

func (t TypedArray) ByteLength() int64 {

	return t.GetAttributeInt64("byteLength")
}

func (t TypedArray) ByteOffset() int64 {

	return t.GetAttributeInt64("byteOffset")
}

func (t TypedArray) BYTES_PER_ELEMENT() int {

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
