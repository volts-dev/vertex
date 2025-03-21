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
	Type  int
	Value struct {
	}

	// This is a syscall/js placeholder.
	Func struct {
		Value // the JavaScript function that invokes the Go function
	}
)

func Global() Value {
	panic(errNotImpl)
}

func (t Type) String() string {
	panic(errNotImpl)
}
func (t Type) isObject() bool {
	return t == TypeObject || t == TypeFunction
}

// Type alias to syscall/js
func (v Value) Type() Type {
	panic(errNotImpl)
}

func (v Value) Get(p string) Value {
	panic(errNotImpl)
}

func (v Value) Set(p string, x interface{}) {
	panic(errNotImpl)
}

func (v Value) Index(i int) Value {
	panic(errNotImpl)
}

func (v Value) SetIndex(i int, x interface{}) {
	panic(errNotImpl)
}

func (v Value) Length() int {
	panic(errNotImpl)
}

func (v Value) Call(m string, args ...interface{}) Value {
	panic(errNotImpl)
}

func (v Value) Invoke(args ...interface{}) Value {
	panic(errNotImpl)
}

func (v Value) New(args ...interface{}) Value {
	panic(errNotImpl)
}

func (v Value) Float() float64 {
	panic(errNotImpl)
}

func (v Value) Int() int {
	panic(errNotImpl)
}

func (v Value) Bool() bool {
	panic(errNotImpl)
}

func (v Value) Truthy() bool {
	panic(errNotImpl)
}

func (v Value) String() string {
	return "undefined"
}

func (v Value) InstanceOf(t Value) bool {
	panic(errNotImpl)
}

func (v Value) IsUndefined() bool {
	panic(errNotImpl)
}

func (v Value) IsNull() bool {
	panic(errNotImpl)
}

// Release alias to syscall/js
func (c Func) Release() {
	panic(errNotImpl)
}

// Error alias to syscall/js
func (e Error) Error() string {
	panic(errNotImpl)
}
