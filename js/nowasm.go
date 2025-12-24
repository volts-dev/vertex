//go:build !js || !wasm

package js

import (
	"fmt"
	"runtime"
)

// Constants that enumerates the JavaScript types.
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

// Type represents the JavaScript type of a Value.
type (
	value struct {
		v   any
		err error
	}

	// This is a syscall/js placeholder.
	function struct {
		Value // the JavaScript function that invokes the Go function
	}
)

var (
	global    = value{v: 2, err: errNotImpl}
	null      = value{v: 1, err: errNotImpl}
	undefined = value{v: 0, err: errNotImpl}
)

func RecoverHandler(r any) {
	fmt.Printf("wasm: panic : %v", r)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Printf("  %s %d\n", file, line)
	}
}

func Undefined() Value {
	return &undefined
}

func Null() Value {
	return &null
}

func Global() Value {
	return &global
}

// ValueOf alias to syscall/js
func ValueOf(x any) Value {
	panic(errNotImpl)
}

// FuncOf alias to syscall/js
func FuncOf(fn func(Value, []Value) any) Func {
	return function{
		Value: &value{err: errNotImpl},
	}
}

func CopyBytesToGo(dst []byte, src Value) (int, error) {
	panic(errNotImpl)
}

func CopyBytesToJS(dst Value, src []byte) (int, error) {
	panic(errNotImpl)
}

func (v *value) Equal(other Value) bool {
	panic(errNotImpl)
}

// Type alias to syscall/js
func (v *value) Type() Type {
	panic(errNotImpl)
}

func (v *value) Get(p string) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Set(p string, x any) Value {
	v.err = errNotImpl
	return v
}
func (v *value) Call(m string, args ...any) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Invoke(args ...any) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Index(i int) Value {
	v.err = errNotImpl
	return v
}

func (v *value) SetIndex(i int, x any) {
	panic(errNotImpl)
}

func (v *value) Length() int {
	panic(errNotImpl)
}

func (v *value) New(args ...any) Value {
	panic(errNotImpl)
}

func (v *value) Float() (float64, error) {
	panic(errNotImpl)
}

func (v *value) Int() (int, error) {
	panic(errNotImpl)
}

func (v *value) Bool() (bool, error) {
	panic(errNotImpl)
}

func (v *value) Truthy() bool {
	panic(errNotImpl)
}

func (v *value) String() (string, error) {
	panic(errNotImpl)
}

func (v *value) Bytes() ([]byte, error) {
	panic(errNotImpl)
}
func (v *value) InstanceOf(t Value) bool {
	panic(errNotImpl)
}

func (v *value) IsUndefined() bool {
	panic(errNotImpl)
}

func (v *value) IsNull() bool {
	panic(errNotImpl)
}

func (v *value) IsNaN() bool {
	panic(errNotImpl)
}

func (v *value) Error() error {
	return errNotImpl
}

// Release alias to syscall/js
func (c function) Release() {
	panic(errNotImpl)
}

func unwrap(val any) any {
	panic(errNotImpl)
}
