package js

import (
	gojs "syscall/js"
)

const (
	TypeUndefined = Type(gojs.TypeUndefined)
	TypeNull      = Type(gojs.TypeNull)
	TypeBoolean   = Type(gojs.TypeBoolean)
	TypeNumber    = Type(gojs.TypeNumber)
	TypeString    = Type(gojs.TypeString)
	TypeSymbol    = Type(gojs.TypeSymbol)
	TypeObject    = Type(gojs.TypeObject)
	TypeFunction  = Type(gojs.TypeFunction)
)

// Type alias to syscall/js
type (
	Type gojs.Type
	// Value alias to syscall/js
	Value gojs.Value
	// Func alias to syscall/js
	Func struct {
		Value
		f gojs.Func // proxy for this func from syscall/js
	}

	// Error alias to syscall/js
	Error struct {
		// Value is the underlying JavaScript error value.
		Value
	}
)

// Undefined alias to syscall/js
func Undefined() Value {
	return Value(gojs.Undefined())
}

// Null alias to syscall/js
func Null() Value {
	return Value(gojs.Null())
}

// ValueOf alias to syscall/js
func ValueOf(x interface{}) Value {
	return Value(gojs.ValueOf(x))
}

// FuncOf alias to syscall/js
func FuncOf(fn func(Value, []Value) interface{}) Func {
	fn2 := func(this gojs.Value, args []gojs.Value) interface{} {
		args2 := make([]Value, len(args))
		for i := range args {
			args2[i] = Value(args[i])
		}
		return fn(Value(this), args2)
	}

	f := gojs.FuncOf(fn2)

	ret := Func{
		Value: Value(f.Value),
		f:     f,
	}

	return ret
}

// CopyBytesToGo alias to syscall/js
func CopyBytesToGo(dst []byte, src Value) int {
	return gojs.CopyBytesToGo(dst, gojs.Value(src))
}

// CopyBytesToJS alias to syscall/js
func CopyBytesToJS(dst Value, src []byte) int {
	return gojs.CopyBytesToJS(gojs.Value(dst), src)
}
func Global() Value {
	return Value(gojs.Global())
}

func (t Type) _String() string {
	return gojs.Type(t).String()
}
func (t Type) isObject() bool {
	return t == TypeObject || t == TypeFunction
}

// Type alias to syscall/js
func (v Value) Type() Type {
	return Type(gojs.Value(v).Type())
}

func (v Value) Get(p string) Value {
	return Value(gojs.Value(v).Get(p))
}

func (v Value) Set(p string, x interface{}) {
	gojs.Value(v).Set(p, x)
}

func (v Value) Index(i int) Value {
	return Value(gojs.Value(v).Index(i))
}

func (v Value) SetIndex(i int, x interface{}) {
	gojs.Value(v).SetIndex(i, x)
}

func (v Value) Length() int {
	return gojs.Value(v).Length()
}

func (v Value) Call(m string, args ...interface{}) Value {
	return Value(gojs.Value(v).Call(m, fixArgsToGojs(args)...))
}

func (v Value) Invoke(args ...interface{}) Value {
	return Value(gojs.Value(v).Invoke(fixArgsToGojs(args)...))
}

func (v Value) New(args ...interface{}) Value {
	return Value(gojs.Value(v).New(fixArgsToGojs(args)...))
}

func (v Value) Float() float64 {
	return gojs.Value(v).Float()
}

func (v Value) Int() int {
	return gojs.Value(v).Int()
}

func (v Value) Bool() bool {
	return gojs.Value(v).Bool()
}

func (v Value) Truthy() bool {
	return gojs.Value(v).Truthy()
}

func (v Value) String() string {
	return gojs.Value(v).String()
}

func (v Value) InstanceOf(t Value) bool {
	return gojs.Value(v).InstanceOf(gojs.Value(t))
}

func (v Value) IsUndefined() bool {
	return gojs.Value(v).IsUndefined()
}

func (v Value) IsNull() bool {
	return gojs.Value(v).IsNull()
}

func fixArgsToGojs(args []interface{}) []interface{} {
	for i := 0; i < len(args); i++ {
		v := args[i]
		if val, ok := v.(Value); ok {
			args[i] = gojs.Value(val) // convert to gojs.Value
		}
		if f, ok := v.(Func); ok {
			args[i] = f.f
		}
		// if ta, ok := v.(TypedArray); ok {
		// 	args[i] = gojs.TypedArray{Value: gojs.Value(ta.Value)}
		// }
	}
	return args
}

// Release alias to syscall/js
func (c Func) Release() {
	// return (*(*sjs.Func)(unsafe.Pointer(&c))).Release()
	c.f.Release()
}

// Error alias to syscall/js
func (e Error) Error() string {
	return gojs.Error{Value: gojs.Value(e.Value)}.Error()
}
