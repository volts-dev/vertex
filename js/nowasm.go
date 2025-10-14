//go:build !js || !wasm

package js

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
		err error
	}

	// This is a syscall/js placeholder.
	function struct {
		Value // the JavaScript function that invokes the Go function
	}
)

func Undefined() Value {
	return &value{
		err: errNotImpl,
	}
}

func Null() Value {
	panic(errNotImpl)
}

func Global() Value {
	panic(errNotImpl)
}

// ValueOf alias to syscall/js
func ValueOf(x interface{}) Value {
	panic(errNotImpl)
}

// FuncOf alias to syscall/js
func FuncOf(fn func(Value, []Value) interface{}) Func {
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

// Type alias to syscall/js
func (v *value) Type() Type {
	panic(errNotImpl)
}

func (v *value) Get(p string) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Set(p string, x interface{}) Value {
	v.err = errNotImpl
	return v
}
func (v *value) Call(m string, args ...interface{}) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Invoke(args ...interface{}) Value {
	v.err = errNotImpl
	return v
}

func (v *value) Index(i int) Value {
	v.err = errNotImpl
	return v
}

func (v *value) SetIndex(i int, x interface{}) {
	panic(errNotImpl)
}

func (v *value) Length() int {
	panic(errNotImpl)
}

func (v *value) New(args ...interface{}) Value {
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

func (v *value) Error() error {
	return errNotImpl
}

// Release alias to syscall/js
func (c function) Release() {
	panic(errNotImpl)
}
