package object

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/array"
)

var (
	ErrNotImpl = errors.New("js not implemented")

	//ErrUnableGetFunctName ErrUnableGetConstructName error
	//ErrUnableGetConstruct ErrUnableGetConstruct error
	ErrUnableGetConstruct = errors.New("Unable to get the constructor")
	//ErrNotImplementedFunc ErrNotImplementedFunc error
	ErrNotImplementedFunc = errors.New("Function.prototype.apply was called on undefined, which is a undefined and not a function")
	//ErrUndefinedValue ErrUndefinedValue error
	//ErrNotAnObject ErrNotAnObject error
	ErrNotAnObject = errors.New("The given value must be an object")
	//ErrObjectNotNumber ErrObjectNotNumber error
	ErrObjectNotNumber = errors.New("The given object is not a number")
	//ErrObjectNotDouble ErrObjectNotDouble error
	ErrObjectNotDouble = errors.New("The given object is not a double")
	//ErrObjectNotString ErrObjectNotString error
	ErrObjectNotString = errors.New("The given object is not a string")
	//ErrObjectNotBool ErrObjectNotBool error
	ErrObjectNotBool = errors.New("The given object is not boolean")
	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnMEv = errors.New("The given value must be an Message Event")
	//ErrNotImplemented ErrNotImplemented error
	//ErrNotImplemented ErrNotImplemented error
	ErrNotABaseObject = errors.New("Not a base object")
)

// Base Object where all herited from
type (
	Object interface {
		js.Object
		Keys() (array.Array, error)
	}

	object struct {
		value js.Value
	}
)

func init() {
	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var objectinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {
	singleton.Do(func() {
		if objectinterface = js.Global().Get("Object"); objectinterface.Error() != nil {
			objectinterface = js.Undefined()
		}
		js.Register(objectinterface, func(v js.Value) (interface{}, error) {
			return ToObject(v)
		})
	})

	return objectinterface
}

// NewFromJSObject Build a Object from a Js js.Value Object given
func ToObject(obj js.Value) (Object, error) {
	var o Object
	if obj.IsUndefined() {
		return o, js.ErrUndefinedValue
	}

	//o.SetObjectValue(obj)
	return &object{
		value: obj,
	}, nil
}

// Base Object_ Return the current BaseObject
func (b *object) BaseObject_() js.Object {
	return b
}

// Empty check if the struct is an empty Struct or have a JS js.Value attached
func (b object) Empty() bool {
	return b.value == nil
}

// Get Get js.Value of Object and handle err
func (b object) GetValueByKey(name string) js.Value {
	return b.value.Get(name)
}

// Get Get js.Value of Object and handle err
func (b object) GetValueByIndex(index int) js.Value {
	return b.value.Index(index)
}

// Set Set js.Value of Object and handle err
func (b object) SetValueByKey(name string, value interface{}) {
	b.value.Set(name, value)
}

// Call
func (b object) Call(name string, args ...interface{}) js.Value {
	return b.value.Call(name, args...)
}

// Discover Use Discover of this struct
func (b object) Discover() (interface{}, error) {
	return nil, nil
}

// ConstructName Get the construct name
func (b object) ConstructName() (string, error) {
	var construct js.Value
	var constructname string
	if construct = b.value.Get("constructor"); construct.IsUndefined() {
		return js.GetFuncName(construct)
	}
	return constructname, ErrUnableGetConstruct
}

// SetObject Set the JS value Object to this struct
func (b *object) SetObjectValue(value js.Value) js.Object {
	b.value = value
	return b
}

// String Get the current string representation of the js js.Value attached to this struct
func (b object) String() (string, error) {
	return b.value.String()
}

// ToString Get the current toString representation of the js js.Value attached to this struct
func (b object) ToString() (string, error) {
	var value js.Value
	if b.value.Type() == js.TypeObject {
		if value = b.Call("toString"); !value.IsUndefined() {
			return value.String()
		}
	}

	return "", js.ErrNotAnObject
}

// js.Value Equivalent to String()
func (b object) GetObjectValue() js.Value {
	if b.value != nil {
		return b.value
	}
	return js.Undefined()

}

// Length Length of the JS.Value attached of this strict
func (b object) Length() int {
	return b.value.Length()
}

// Bind Bind
func (b object) Bind(to js.Object) (interface{}, error) {
	var err error
	var bindObj js.Value
	var gobj interface{}

	if bindObj = b.Call("bind", to.GetObjectValue()); !bindObj.IsUndefined() {
		//gobj, err = Discover(bindObj)

	}
	return gobj, err
}

// Implement Check if the stuct implement a given name method
func (b object) Implement(method string) (bool, error) {

	var obj js.Value

	var err error

	if obj = b.value.Get(method); obj.IsUndefined() {
		if obj.Type() == js.TypeFunction {
			return true, nil
		}
	}

	return false, err
}

func (b object) Class() (string, error) {
	var err error
	var objconstructor, objname js.Value
	var classname string

	if objconstructor = b.value.Get("constructor"); objconstructor.IsUndefined() {
		if objname = objconstructor.Get("name"); objname.IsUndefined() {
			return objname.String()
		}
	}

	return classname, err
}
func (o object) Keys() (array.Array, error) {
	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj = ai.Call("keys", o.value); obj.Error() == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func (b object) SetFunc(attribute string, f func(this js.Value, args []js.Value) any) {
	b.value.Set(attribute, js.FuncOf(f))
}

func (b object) SetAttribute(attribute string, i interface{}) error {
	b.value.Set(attribute, js.ValueOf(i))
	return b.value.Error()
}

func (b object) Export(name string) {
	js.Global().Set(name, b.value)
}

func (b object) GetAttributeString(attribute string) (string, error) {
	var err error
	var obj js.Value
	var ret = ""

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = js.ErrUndefinedValue
	} else {
		if obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.Type() == js.TypeString {
				return obj.String()
			} else {
				err = js.ErrObjectNotString
			}
		}

	}

	return ret, err
}

func (b object) GetAttributeGlobal(attribute string) (interface{}, error) {
	panic(ErrNotImpl)

	/*
	   var err error

	   	var obj js.Value
	   	var objGlobal interface{}

	   	obj = b.value.Get(attribute)
	   	if obj.IsUndefined() {
	   		err = ErrUndefinedValue
	   	} else {
	   		objGlobal, err = GoValue(obj)
	   	}

	   	return objGlobal, err
	*/
}

func (b object) SetAttributeString(attribute string, value string) error {
	b.value.Set(attribute, js.ValueOf(value))
	//return b.Set(attribute, ValueOf(value))
	return b.value.Error()
}

func (b object) GetAttributeBool(attribute string) (bool, error) {
	var err error
	var obj js.Value
	var ret bool

	obj = b.value.Get(attribute)
	if obj.Type() == js.TypeBoolean {
		return obj.Bool()
	} else {
		err = js.ErrObjectNotBool
	}

	return ret, err
}

func (b object) SetAttributeBool(attribute string, value bool) error {
	b.value.Set(attribute, js.ValueOf(value))
	return b.value.Error()
}

func (b object) GetAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = js.ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			return obj.Int()
		} else {
			err = js.ErrObjectNotNumber
		}
	}

	return result, err
}

func (b object) GetAttributeInt64(attribute string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = js.ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			retf, err := obj.Float()
			return int64(retf), err
		} else {
			err = js.ErrObjectNotNumber
		}
	}

	return ret, err
}

func (b object) SetAttributeInt(attribute string, value int) error {
	b.value.Set(attribute, js.ValueOf(value))
	return b.value.Error()
}

func (b object) GetAttributeDouble(attribute string) (float64, error) {
	var err error
	var obj js.Value
	var result float64

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = js.ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			return obj.Float()
		} else {
			err = js.ErrObjectNotDouble
		}
	}

	return result, err
}

func (b object) SetAttributeDouble(attribute string, value float64) error {
	b.value.Set(attribute, js.ValueOf(value))
	return b.value.Error()
}

// CallInt64 Call method given and return a 64bit int
func (b object) CallInt64(method string) (int64, error) {
	obj := b.value.Call(method)
	if obj.Type() == js.TypeNumber {
		fv, err := obj.Float()
		return int64(fv), err
	}

	return 0, js.ErrObjectNotNumber
}

// CallInt Call method given and return int
func (b object) CallInt(method string) (int, error) {
	obj := b.value.Call(method)
	if obj.Type() == js.TypeNumber {
		return obj.Int()
	}

	return 0, js.ErrObjectNotNumber
}

// CallInt64 Call method given and return a bool int
func (b object) CallBool(method string) (bool, error) {
	obj := b.value.Call(method)
	if obj.Type() == js.TypeBoolean {
		return obj.Bool()
	}

	return false, js.ErrObjectNotBool
}

func (b object) Debug(msg string) error {
	return js.Debug(msg)
}
