package js

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type (
	Value interface {
		Type() Type
		Get(p string) Value
		Set(p string, x interface{}) Value
		Index(i int) Value
		SetIndex(i int, x interface{})
		Length() int
		Call(m string, args ...interface{}) Value
		Invoke(args ...interface{}) Value
		New(args ...interface{}) Value
		Float() (float64, error)
		Int() (int, error)
		Bool() (bool, error)
		String() (string, error)
		Bytes() ([]byte, error)
		Truthy() bool
		InstanceOf(t Value) bool
		IsUndefined() bool
		IsNull() bool
		Error() error
	}

	Func interface {
		Value
		Release()
	}

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
	}

	// ObjectFrom Interface to check if Object is a BaseObject
	ObjectFrom interface {
		Value() Value
		BaseObject_() Object
	}

	__Error interface {
		Error() string
	}

	Type int
)

var (
	errNotImpl            = errors.New("js not implemented")
	ErrNotImplemented     = errors.New("Browser not implemented this feature")
	ErrNotImplementedFunc = errors.New("Function.prototype.apply was called on undefined, which is a undefined and not a function")
	ErrUndefinedValue     = errors.New("Undefined value")
	ErrUnableGetFunctName = errors.New("Unable to get the func name")

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

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "undefined"
	case TypeNull:
		return "null"
	case TypeBoolean:
		return "boolean"
	case TypeNumber:
		return "number"
	case TypeString:
		return "string"
	case TypeSymbol:
		return "symbol"
	case TypeObject:
		return "object"
	case TypeFunction:
		return "function"
	default:
		panic("bad type")
	}
}
func (t Type) isObject() bool {
	return t == TypeObject || t == TypeFunction
}

// WasmExecJsPath find wasm_exec.js in the local Go distribution and return it's path.
// Return error if not found.
func WasmExecJsPath() (string, error) {
	b, err := exec.Command("go", "env", "GOROOT").CombinedOutput()
	if err != nil {
		return "", err
	}
	bstr := strings.TrimSpace(string(b))
	if bstr == "" {
		return "", fmt.Errorf("failed to find wasm_exec.js, empty path from `go env GOROOT`")
	}

	p := filepath.Join(bstr, "misc/wasm/wasm_exec.js")
	_, err = os.Stat(p)
	if err != nil {
		return "", err
	}

	return p, nil
}

// MustWasmExecJsPath find wasm_exec.js in the local Go distribution and return it's path.
// Panic if not found.
func MustWasmExecJsPath() string {
	s, err := WasmExecJsPath()
	if err != nil {
		panic(err)
	}
	return s
}
func GetFuncName(inter Value) (string, error) {
	var obj Value
	var err error
	var name string

	obj = inter.Get("name")
	if !obj.IsUndefined() {
		return obj.String()
	} else {
		err = ErrUnableGetFunctName
	}

	return name, err
}

func Eval(str string) (Value, error) {
	f := Global().Call("eval", str)
	return f, f.Error()
}

func Self() (interface{}, error) {

	var err error
	var self Value

	if self = Global().Get("self"); self.Error() == nil {
		return Discover(self)
	}

	return nil, err
}
