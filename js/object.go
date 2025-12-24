package js

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNotImpl = errors.New("js not implemented")

	//ErrUnableGetFunctName ErrUnableGetConstructName error
	//ErrUnableGetConstruct ErrUnableGetConstruct error
	ErrUnableGetConstruct = errors.New("Unable to get the constructor")
	//ErrNotImplementedFunc ErrNotImplementedFunc error
/*	ErrNotImplementedFunc = errors.New("Function.prototype.apply was called on undefined, which is a undefined and not a function")
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
	ErrNotABaseObject = errors.New("Not a base object")*/
)

// Base Object where all herited from
type ( /*
		Object interface {
			//BaseObject_() IObject
			Empty() bool
			GetValueByKey(string) Value
			GetValueByIndex(int) Value
			SetValueByKey(string, interface{})
			Call(string, ...interface{}) Value
			Discover() (interface{}, error)
			ConstructName() (string, error)
			SetObjectValue(Value) Object
			String() (string, error)
			ToString() (string, error)
			GetObjectValue() Value
			Length() int
			Bind(Object) (interface{}, error)
			Implement(string) (bool, error)
			Class() (string, error)
			SetFunc(string, func(Value, []Value) any)
			SetAttribute(string, interface{}) error
			Export(string)
			GetAttributeString(string) (string, error)
			GetAttributeGlobal(string) (interface{}, error)
			SetAttributeString(string, string) error
			GetAttributeBool(string) (bool, error)
			SetAttributeBool(string, bool) error
			GetAttributeInt(string) (int, error)
			GetAttributeInt64(string) (int64, error)
			SetAttributeInt(string, int) error
			GetAttributeDouble(string) (float64, error)
			SetAttributeDouble(string, float64) error
			CallInt64(string) (int64, error)
			CallInt(string) (int, error)
			CallBool(string) (bool, error)
			Debug(msg string) error
			//GoValue_(object Value) interface{}
			Class_() string
			ToString_() string
			ConstructName_() string
			GetAttributeString_(attribute string) string

			Keys() (Array, error)
		}*/

	Object struct {
		value Value
	}
)

func init() {
	//js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var objectinterface Value

// GetInterface get the JS interface
func GetObjectInterface() Value {
	singleton.Do(func() {
		if objectinterface = Global().Get("Object"); objectinterface.Error() != nil {
			objectinterface = Undefined()
		}

		Register(objectinterface, func(v Value) (interface{}, error) {
			return ToObject(v)
		})
	})

	return objectinterface
}

// NewFromJSObject Build a Object from a Js  Value Object given
func ToObject(obj Value) (Object, error) {
	var o Object
	if obj.IsUndefined() || obj.IsNull() {
		return o, ErrUndefinedValue
	}

	//o.SetObjectValue(obj)
	return Object{
		value: obj,
	}, nil
}

// Base Object_ Return the current BaseObject
func (b *Object) BaseObject_() Object {
	return *b
}

func (b *Object) Equal(other any) bool {
	if v, ok := other.(ObjectFrom); ok {
		fmt.Println("Equal", b)
		return b.value.Equal(v.GetObjectValue())
	}

	return false
}

// Empty check if the struct is an empty Struct or have a JS  Value attached
func (b Object) Empty() bool {
	return b.value == nil
}

// Get Get  Value of Object and handle err
func (b Object) GetValueByKey(name string) Value {
	return b.value.Get(name)
}

// Get Get  Value of Object and handle err
func (b Object) GetValueByIndex(index int) Value {
	return b.value.Index(index)
}

// Set Set  Value of Object and handle err
func (b Object) SetValueByKey(name string, value interface{}) {
	b.value.Set(name, value)
}

// Call
func (b Object) Call(name string, args ...interface{}) Value {
	return b.value.Call(name, args...)
}

// Discover Use Discover of this struct
func (b Object) Discover() (interface{}, error) {
	return nil, nil
}

// ConstructName Get the construct name
func (b Object) ConstructName() (string, error) {
	var construct Value
	var constructname string
	if construct = b.value.Get("constructor"); construct.IsUndefined() {
		return GetFuncName(construct)
	}
	return constructname, ErrUnableGetConstruct
}

// SetObject Set the JS value Object to this struct
func (b *Object) SetObjectValue(value Value) *Object {
	if b != nil {
		b.value = value
		return b
	}

	*b = Object{value: value}
	return b
}

// String Get the current string representation of the js  Value attached to this struct
func (b Object) String() (string, error) {
	return b.value.String()
}

// ToString Get the current toString representation of the js  Value attached to this struct
func (b Object) ToString() (string, error) {
	var value Value
	if b.value.Type() == TypeObject {
		if value = b.Call("toString"); !value.IsUndefined() {
			return value.String()
		}
	}

	return "", ErrNotAnObject
}

// Value Equivalent to String()
func (b Object) GetObjectValue() Value {
	if b.value != nil {
		return b.value
	}

	return Undefined()
}

// Length Length of the JS.Value attached of this strict
func (b Object) Length() int {
	return b.value.Length()
}

// Bind Bind
func (b Object) Bind(to Object) (interface{}, error) {
	var err error
	var bindObj Value
	var gobj interface{}

	if bindObj = b.Call("bind", to.GetObjectValue()); !bindObj.IsUndefined() {
		//gobj, err = Discover(bindObj)

	}
	return gobj, err
}

// Implement Check if the stuct implement a given name method
func (b Object) Implement(method string) (bool, error) {

	var obj Value

	var err error

	if obj = b.value.Get(method); obj.IsUndefined() {
		if obj.Type() == TypeFunction {
			return true, nil
		}
	}

	return false, err
}

func (b Object) Class() (string, error) {
	var err error
	var objconstructor, objname Value
	var classname string

	if objconstructor = b.value.Get("constructor"); objconstructor.IsUndefined() {
		if objname = objconstructor.Get("name"); objname.IsUndefined() {
			return objname.String()
		}
	}

	return classname, err
}
func (o Object) Keys() (Array, error) {
	var err error
	var obj Value
	var newArr Array

	if ai := GetObjectInterface(); !ai.IsUndefined() {
		if obj = ai.Call("keys", o.value); obj.Error() == nil {
			newArr, err = NewArrayFromJSObject(obj)

		}

	}

	return newArr, err
}

func (b Object) SetFunc(attribute string, f func(this Value, args []Value) any) {
	b.value.Set(attribute, FuncOf(f))
}

func (b Object) SetAttribute(attribute string, i interface{}) error {
	b.value.Set(attribute, ValueOf(i))
	return b.value.Error()
}

func (b Object) Export(name string) {
	Global().Set(name, b.value)
}

func (b Object) GetAttributeString(attribute string) (string, error) {
	var err error
	var obj Value
	var ret = ""

	obj = b.value.Get(attribute)
	fmt.Println("GetAttributeString", obj)
	if obj.IsUndefined() || obj.IsNull() {
		err = ErrUndefinedValue
	} else {
		if obj.Type() == TypeString {
			return obj.String()
		} else {
			err = ErrObjectNotString
		}
	}

	return ret, err
}

func (b Object) GetAttributeGlobal(attribute string) (interface{}, error) {
	panic(ErrNotImpl)

	/*
	   var err error

	   	var obj  Value
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

func (b Object) SetAttributeString(attribute string, value string) error {
	b.value.Set(attribute, ValueOf(value))
	//return b.Set(attribute, ValueOf(value))
	return b.value.Error()
}

func (b Object) GetAttributeBool(attribute string) (bool, error) {
	var err error
	var obj Value
	var ret bool

	obj = b.value.Get(attribute)
	if obj.Type() == TypeBoolean {
		return obj.Bool()
	} else {
		err = ErrObjectNotBool
	}

	return ret, err
}

func (b Object) SetAttributeBool(attribute string, value bool) error {
	b.value.Set(attribute, ValueOf(value))
	return b.value.Error()
}

func (b Object) GetAttributeInt(attribute string) (int, error) {

	var err error
	var obj Value
	var result int

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == TypeNumber {
			return obj.Int()
		} else {
			err = ErrObjectNotNumber
		}
	}

	return result, err
}

func (b Object) GetAttributeInt64(attribute string) (int64, error) {
	var err error
	var obj Value
	var ret int64

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == TypeNumber {
			retf, err := obj.Float()
			return int64(retf), err
		} else {
			err = ErrObjectNotNumber
		}
	}

	return ret, err
}

func (b Object) SetAttributeInt(attribute string, value int) error {
	b.value.Set(attribute, ValueOf(value))
	return b.value.Error()
}

func (b Object) GetAttributeDouble(attribute string) (float64, error) {
	var err error
	var obj Value
	var result float64

	obj = b.value.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == TypeNumber {
			return obj.Float()
		} else {
			err = ErrObjectNotDouble
		}
	}

	return result, err
}

func (b Object) SetAttributeDouble(attribute string, value float64) error {
	b.value.Set(attribute, ValueOf(value))
	return b.value.Error()
}

// CallInt64 Call method given and return a 64bit int
func (b Object) CallInt64(method string) (int64, error) {
	obj := b.value.Call(method)
	if obj.Type() == TypeNumber {
		fv, err := obj.Float()
		return int64(fv), err
	}

	return 0, ErrObjectNotNumber
}

// CallInt Call method given and return int
func (b Object) CallInt(method string) (int, error) {
	obj := b.value.Call(method)
	if obj.Type() == TypeNumber {
		return obj.Int()
	}

	return 0, ErrObjectNotNumber
}

// CallInt64 Call method given and return a bool int
func (b Object) CallBool(method string) (bool, error) {
	obj := b.value.Call(method)
	if obj.Type() == TypeBoolean {
		return obj.Bool()
	}

	return false, ErrObjectNotBool
}

func (b Object) Debug(msg string) error {
	return Debug(msg)
}

func (b Object) Class_() string {

	var c string
	var err error

	if c, err = b.Class(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b Object) ToString_() string {

	var c string
	var err error

	if c, err = b.ToString(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b Object) ConstructName_() string {

	var c string
	var err error

	if c, err = b.ConstructName(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b Object) GetAttributeString_(attribute string) string {
	var c string
	var err error

	if c, err = b.GetAttributeString(attribute); err != nil {
		b.Debug(err.Error())
	}

	return c
}
