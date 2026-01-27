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
		Equal(other Value) bool
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
		IsNaN() bool
		Error() error
	}

	Func interface {
		Value
		Release()
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

// GetFuncName 获取函数名称，从 JS 对象的 name 属性中读取
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

// Eval 执行 JavaScript 代码并返回结果
func Eval(str string) (Value, error) {
	if str == "" {
		return Undefined(), fmt.Errorf("eval: empty string")
	}

	f := (&value{v: global}).Call("eval", str)
	if err := f.Error(); err != nil {
		return nil, fmt.Errorf("eval failed: %w", err)
	}

	return f, nil
}

// Self 获取全局 self 对象（在 Web Workers 中使用）
func Self() (Value, error) {
	self := (&value{v: global}).Get("self")
	if err := self.Error(); err != nil {
		return nil, fmt.Errorf("failed to get self: %w", err)
	}

	if self.IsUndefined() {
		return nil, fmt.Errorf("self is undefined")
	}

	return self, nil
}

// Get is a shorthand for Global().Get().
func Get(name string) Value {
	return (&value{v: global}).Get(name)
}

// Set is a shorthand for Global().Set().
func Set(name string, v interface{}) {
	(&value{v: global}).Set(name, v)
}

// Call is a shorthand for Global().Call().
func Call(name string, args ...interface{}) Value {
	return (&value{v: global}).Call(name, args...)
}

func Reflect() Value {
	return (&value{v: global}).Get("Reflect")
}
