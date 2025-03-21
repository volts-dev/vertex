package vertex

import (
	"errors"

	"github.com/volts-dev/vertex/core/js"
)

var (
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
	ErrNotImplemented = errors.New("Browser not implemented Object")
	//ErrNotImplemented ErrNotImplemented error
	ErrNotABaseObject = errors.New("Not a base object")
	//ErrUnableGetFunctName ErrUnableGetConstructName error
	ErrUnableGetFunctName = errors.New("Unable to get the func name")
	//ErrUnableGetConstruct ErrUnableGetConstruct error
	ErrUnableGetConstruct = errors.New("Unable to get the constructor")
	//ErrNotImplementedFunc ErrNotImplementedFunc error
	ErrNotImplementedFunc = errors.New("Function.prototype.apply was called on undefined, which is a undefined and not a function")
	//ErrUndefinedValue ErrUndefinedValue error
	ErrUndefinedValue = errors.New("Undefined value")
)

// Base Object where all herited from
type Object struct {
	object *js.Value
}

// ObjectFrom Interface to check if Object is a BaseObject
type ObjectFrom interface {
	JSObject() js.Value
	BaseObject_() Object
}

// NewFromJSObject Build a Object from a Js Value Object given
func ToObject(obj js.Value) (Object, error) {
	var o Object
	if obj.IsUndefined() {
		return o, ErrUndefinedValue
	}
	o.object = &obj
	return o, nil

}
func GetFuncName(inter js.Value) (string, error) {
	var obj js.Value
	var err error
	var name string

	obj = inter.Get("name")
	if !obj.IsUndefined() {
		name = obj.String()
	} else {
		err = ErrUnableGetFunctName
	}

	return name, err
}

func GoValue(object js.Value) (interface{}, error) {
	var err error
	switch object.Type() {
	case js.TypeNumber:
		/*
			if v, err := IsInteger(object); err == nil && v {
				return int64(object.Float()), nil
			}*/
		return object.Float(), nil
	case js.TypeString:
		return object.String(), nil
	case js.TypeBoolean:
		return object.Bool(), nil
	case js.TypeNull:
		return nil, nil
	}

	obj, err := Discover(object)

	return obj, err
}

// Base Object_ Return the current BaseObject
func (b Object) BaseObject_() Object {
	return b
}

// Empty check if the struct is an empty Struct or have a JS Value attached
func (b Object) Empty() bool {
	return b.object == nil
}

// Get Get Value of object and handle err
func (b Object) Get(name string) js.Value {
	return b.object.Get(name)
}

// Get Get Value of object and handle err
func (b Object) GetIndex(index int) js.Value {
	return b.object.Index(index)
}

// Set Set Value of object and handle err
func (b Object) Set(name string, value interface{}) {
	b.object.Set(name, value)
}

// Call
func (b Object) Call(name string, args ...interface{}) js.Value {
	return b.object.Call(name, args...)
}

// Discover Use Discover of this struct
func (b Object) Discover() (interface{}, error) {
	return nil, nil
}

// ConstructName Get the construct name
func (b Object) ConstructName() (string, error) {
	var construct js.Value
	var constructname string
	if construct = b.Get("constructor"); construct.IsUndefined() {
		return GetFuncName(construct)
	}
	return constructname, ErrUnableGetConstruct
}

// SetObject Set the JS value Object to this struct
func (b Object) SetObject(object js.Value) Object {
	b.object = &object
	return b
}

// JSObject Give the JS Value Object attach to this struct
func (b Object) JSObject() js.Value {
	if b.object != nil {
		return *b.object
	} else {
		return js.Undefined()
	}

}

// String Get the current string representation of the js Value attached to this struct
func (b Object) String() string {
	return b.object.String()
}

// ToString Get the current toString representation of the js Value attached to this struct
func (b Object) ToString() (string, error) {
	var value js.Value
	if b.JSObject().Type() == js.TypeObject {
		if value = b.Call("toString"); !value.IsUndefined() {
			return value.String(), nil
		}
	}

	return "", ErrNotAnObject
}

// Value Equivalent to String()
func (b Object) Value() string {
	return b.object.String()
}

// Length Length of the JS.Value attached of this strict
func (b Object) Length() int {
	return b.object.Length()
}

// Bind Bind
func (b Object) Bind(to Object) (interface{}, error) {
	var err error
	var bindObj js.Value
	var gobj interface{}

	if bindObj = b.Call("bind", to.JSObject()); !bindObj.IsUndefined() {
		//gobj, err = Discover(bindObj)

	}
	return gobj, err
}

// Implement Check if the stuct implement a given name method
func (b Object) Implement(method string) (bool, error) {

	var obj js.Value

	var err error

	if obj = b.Get(method); obj.IsUndefined() {
		if obj.Type() == js.TypeFunction {
			return true, nil
		}
	}

	return false, err
}

func (b Object) Class() (string, error) {
	var err error
	var objconstructor, objname js.Value
	var classname string

	if objconstructor = b.Get("constructor"); objconstructor.IsUndefined() {
		if objname = objconstructor.Get("name"); objname.IsUndefined() {
			classname = objname.String()
		}
	}

	return classname, err
}

func (b Object) SetFunc(attribute string, f func(this js.Value, args []js.Value) any) {
	b.Set(attribute, js.FuncOf(f))
}

func (b Object) SetAttribute(attribute string, i interface{}) {
	b.Set(attribute, js.ValueOf(i))
}

func (b Object) Export(name string) {
	js.Global().Set(name, *b.object)
}

func (b Object) GetAttributeString(attribute string) (string, error) {
	var err error
	var obj js.Value
	var ret = ""

	obj = b.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {
		if obj.IsNull() {
			err = ErrUndefinedValue
		} else {
			if obj.Type() == js.TypeString {
				ret = obj.String()
			} else {
				err = ErrObjectNotString
			}
		}

	}

	return ret, err
}

func (b Object) GetAttributeGlobal(attribute string) (interface{}, error) {
	var err error
	var obj js.Value
	var objGlobal interface{}

	obj = b.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {
		objGlobal, err = GoValue(obj)
	}

	return objGlobal, err

}

func (b Object) SetAttributeString(attribute string, value string) {
	b.Set(attribute, js.ValueOf(value))
	//return b.Set(attribute, js.ValueOf(value))
}

func (b Object) GetAttributeBool(attribute string) (bool, error) {
	var err error
	var obj js.Value
	var ret bool

	obj = b.Get(attribute)
	if obj.Type() == js.TypeBoolean {
		ret = obj.Bool()
	} else {
		err = ErrObjectNotBool
	}

	return ret, err
}

func (b Object) SetAttributeBool(attribute string, value bool) {
	b.Set(attribute, js.ValueOf(value))
}

func (b Object) GetAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	obj = b.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = ErrObjectNotNumber
		}
	}

	return result, err
}

func (b Object) GetAttributeInt64(attribute string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64

	obj = b.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			ret = int64(obj.Float())
		} else {
			err = ErrObjectNotNumber
		}
	}

	return ret, err
}

func (b Object) SetAttributeInt(attribute string, value int) {
	b.Set(attribute, js.ValueOf(value))
}

func (b Object) GetAttributeDouble(attribute string) (float64, error) {
	var err error
	var obj js.Value
	var result float64

	obj = b.Get(attribute)
	if obj.IsUndefined() {
		err = ErrUndefinedValue
	} else {

		if obj.Type() == js.TypeNumber {
			result = obj.Float()
		} else {
			err = ErrObjectNotDouble
		}
	}

	return result, err
}

func (b Object) SetAttributeDouble(attribute string, value float64) {

	b.Set(attribute, js.ValueOf(value))
}

// CallInt64 Call method given and return a 64bit int
func (b Object) CallInt64(method string) (int64, error) {

	var err error
	var obj js.Value
	var ret int64

	obj = b.Call(method)
	if obj.Type() == js.TypeNumber {
		ret = int64(obj.Float())
	} else {
		err = ErrObjectNotNumber
	}

	return ret, err
}

// CallInt Call method given and return int
func (b Object) CallInt(method string) (int, error) {

	var err error
	var obj js.Value
	var ret int

	obj = b.Call(method)
	if obj.Type() == js.TypeNumber {
		ret = obj.Int()
	} else {
		err = ErrObjectNotNumber
	}
	return ret, err
}

// CallInt64 Call method given and return a bool int
func (b Object) CallBool(method string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	obj = b.Call(method)
	if obj.Type() == js.TypeBoolean {
		result = obj.Bool()
	} else {
		err = ErrObjectNotBool
	}

	return result, err
}
