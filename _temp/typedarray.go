package js

import (
	"errors"
	"sync"
)

var (
	//ErrNotAUint8Array ErrNotAUint8Array error
	ErrNotAUint8Array                  = errors.New("Object is not a Uint8Array")
	ErrNotImplementedUint8Array        = errors.New("Browser not implemented Uint8Array")
	ErrNotAUint8ClampedArray           = errors.New("Object is not a Uint8ClampedArray")
	ErrNotImplementedUint8ClampedArray = errors.New("Browser not implemented Uint8ClampedArray")
	ErrNotAInt8Array                   = errors.New("Object is not a Int8Array")
	ErrNotImplementedInt8Array         = errors.New("Browser not implemented Int8Array")
	ErrNotAUint16Array                 = errors.New("Object is not a Uint16Array")
	ErrNotImplementedUint16Array       = errors.New("Browser not implemented Uint16Array")
	ErrNotAInt16Array                  = errors.New("Object is not a Int16Array")
	ErrNotImplementedInt16Array        = errors.New("Browser not implemented Int16Array")
	ErrNotAInt32Array                  = errors.New("Object is not a Int32Array")
	ErrNotImplementedInt32Array        = errors.New("Browser not implemented Int32Array")
	ErrNotAUint32Array                 = errors.New("Object is not a Uint32Array")
	ErrNotImplementedUint32Array       = errors.New("Browser not implemented Uint32Array")
	ErrNotAFloat32Array                = errors.New("Object is not a Float32Array")
	ErrNotImplementedFloat32Array      = errors.New("Browser not implemented Float32Array")
	ErrNotAFloat64Array                = errors.New("Object is not a Float64Array")
	ErrNotImplementedFloat64Array      = errors.New("Browser not implemented Float64Array")
)

func init() {
	RegisterInterface(GetTypedArrayInterface[Uint8Array])
	RegisterInterface(GetTypedArrayInterface[Uint8ClampedArray])
	RegisterInterface(GetTypedArrayInterface[Uint16Array])
	RegisterInterface(GetTypedArrayInterface[Uint32Array])
	RegisterInterface(GetTypedArrayInterface[Int8Array])
	RegisterInterface(GetTypedArrayInterface[Int16Array])
	RegisterInterface(GetTypedArrayInterface[Int32Array])
	RegisterInterface(GetTypedArrayInterface[Float32Array])
	RegisterInterface(GetTypedArrayInterface[Float64Array])

}

var singletonuint8array sync.Once
var singletonuint16array sync.Once
var singletonuint32array sync.Once
var singletonint8array sync.Once
var singletonint16array sync.Once
var singletonint32array sync.Once
var singletonFloat32array sync.Once
var singletonarray sync.Once

var uint8arrayinterface Value
var uint16arrayinterface Value
var uint32arrayinterface Value
var int8arrayinterface Value
var int16arrayinterface Value
var int32arrayinterface Value
var Float32arrayinterface Value
var Float64arrayinterface Value

type (
	// TypedArray struct
	TypedArray struct {
		Array
	}

	TypedArrayFrom interface {
		TypedArray_() TypedArray
	}

	// Uint8Array struct
	Uint8Array struct {
		TypedArray
	}
	// Uint8ClampedArray struct
	Uint8ClampedArray struct {
		TypedArray
	}

	Uint8ClampedArrayFrom interface {
		Uint8ClampedArray_() Uint8ClampedArray
	}

	Uint8ArrayFrom interface {
		Uint8Array_() Uint8Array
	}

	// Uint16Array struct
	Uint16Array struct {
		TypedArray
	}

	Uint16ArrayFrom interface {
		Uint16Array_() Uint16Array
	}

	// Uint32Array struct
	Uint32Array struct {
		TypedArray
	}

	Uint32ArrayFrom interface {
		Uint32Array_() Uint32Array
	}

	// Uint8Array struct
	Int8Array struct {
		TypedArray
	}

	Int8ArrayFrom interface {
		Uint8Array_() Uint8Array
	}

	// Int16Array struct
	Int16Array struct {
		TypedArray
	}

	Int16ArrayFrom interface {
		Int16Array_() Int16Array
	}

	// Int32Array struct
	Int32Array struct {
		TypedArray
	}

	Int32ArrayFrom interface {
		Int32Array_() Int32Array
	}

	// Float32Array struct
	Float32Array struct {
		TypedArray
	}

	Float32ArrayFrom interface {
		Float32Array_() Float32Array
	}

	// Float64Array struct
	Float64Array struct {
		TypedArray
	}

	Float64ArrayFrom interface {
		Float64Array_() Float64Array
	}
)

func (t TypedArray) TypedArray_() TypedArray {
	return t
}

func (t TypedArray) Bytes() ([]byte, error) {
	var err error
	var buffer []byte
	var l int
	if l, err = t.Length(); err == nil {
		buffer = make([]byte, l)
		if _, err = CopyBytesToGo(buffer, t.Value()); err == nil {
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

	return CopyBytesToGo(buffer, t.Value())

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

	return CopyBytesToJS(t.Value(), buffer)
}

func (t TypedArray) Buffer() (ArrayBuffer, error) {

	var err error
	var obj Value

	if obj = t.Get("buffer"); obj.Error() == nil {

		return ToArrayBuffer(obj)

	}

	return ArrayBuffer{}, err
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
	var obj Value
	var newArr interface{}

	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, ValueOf(opt))
		}
	}

	if obj = t.Call("subarray", arrayJS...); obj.Error() == nil {
		newArr, err = Discover(obj)

	}

	return newArr, err
}

func newTypedArrayFrom(interfacejs Value, f func(Value) (interface{}, error), iterable interface{}) (interface{}, error) {
	var obj Value
	var err error
	var newArr interface{}
	if obj = interfacejs.Call("from", ValueOf(iterable)); obj.Error() == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}

func newTypedArrayOf(interfacejs Value, f func(Value) (interface{}, error), values ...interface{}) (interface{}, error) {
	var newArr interface{}
	var arrayJS []interface{}
	var obj Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, ValueOf(value))
	}
	if obj = interfacejs.Call("of", arrayJS...); obj.Error() == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}

func GetTypedArrayInterface[T Uint8Array | Uint8ClampedArray | Uint16Array | Uint32Array | Int8Array | Int16Array | Int32Array | Float32Array | Float64Array]() (v Value) {
	var itf Value

	singletonarray.Do(func() {
		var typeName string
		var vv any = v

		switch vv.(type) {
		case Uint8Array:
			typeName = "Uint8Array"
		case Uint8ClampedArray:
			typeName = "Uint8ClampedArray"
		case Uint16Array:
			typeName = "Uint16Array"
		case Uint32Array:
			typeName = "Uint32Array"
		case Int8Array:
			typeName = "Int8Array"
		case Int16Array:
			typeName = "Int16Array"
		case Int32Array:
			typeName = "Uint32Array"
		case Float32Array:
			typeName = "Float32Array"
		case Float64Array:
			typeName = "Float64Array"
			//itf = Float64arrayinterface
		}

		if itf = Global().Get(typeName); itf.Error() != nil {
			itf = Undefined()
		}

		Register(itf, func(v Value) (interface{}, error) {
			return ToTypedArray[T](v)
		})
	})

	return itf
}
func NewTypedArray[T Uint8Array | Uint8ClampedArray | Uint16Array | Uint32Array | Int8Array | Int16Array | Int32Array | Float32Array | Float64Array](value interface{}) (v T, e error) {
	var obj Value
	var err error
	if ai := GetTypedArrayInterface[T](); !ai.IsUndefined() {
		if obj = ai.New(ValueOf(value)); obj.Error() == nil {
			//a.BaseObject = a.SetObject(objnew)
			var i any = v
			if vv, ok := i.(Object); ok {
				vv.SetValue(obj)
			}
		}

	} else {
		err = ErrNotImplementedFloat32Array
	}

	return v, err
}

func NewTypedArrayOf[T Uint8Array | Uint8ClampedArray | Uint16Array | Uint32Array | Int8Array | Int16Array | Int32Array | Float32Array | Float64Array](values ...interface{}) (v T, e error) {
	arr, err := newTypedArrayOf(GetTypedArrayInterface[T](), func(v Value) (interface{}, error) {
		return ToTypedArray[T](v)
	}, values...)
	return arr.(T), err
}

func ToTypedArray[T Uint8Array | Uint8ClampedArray | Uint16Array | Uint32Array | Int8Array | Int16Array | Int32Array | Float32Array | Float64Array](obj Value) (v T, e error) {
	var err error

	if ui := GetTypedArrayInterface[T](); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				//u.BaseObject = u.SetObject(obj)

				var i any = v
				if vv, ok := i.(Object); ok {
					vv.SetValue(obj)
				}
			} else {
				err = ErrNotAFloat64Array
			}
		}
	} else {
		err = ErrNotImplementedFloat64Array
	}

	return v, err
}
