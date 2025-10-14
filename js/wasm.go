//go:build js && wasm

package js

import (
	"errors"
	"fmt"
	"syscall/js"
)

const (
	TypeUndefined = Type(js.TypeUndefined)
	TypeNull      = Type(js.TypeNull)
	TypeBoolean   = Type(js.TypeBoolean)
	TypeNumber    = Type(js.TypeNumber)
	TypeString    = Type(js.TypeString)
	TypeSymbol    = Type(js.TypeSymbol)
	TypeObject    = Type(js.TypeObject)
	TypeFunction  = Type(js.TypeFunction)
)

// Type alias to syscall/js
type (
	// Value alias to syscall/js
	// value js.Value
	value struct {
		v   js.Value
		err error
	}

	// Func alias to syscall/js
	function struct {
		Value
		f js.Func // proxy for this func from syscall/js
	}
)

// Undefined alias to syscall/js
func Undefined() Value {
	return &value{v: js.Undefined()}
}

// Null alias to syscall/js
func Null() Value {
	return &value{v: js.Null()}
}

func Global() Value {
	return &value{v: js.Global()}
}

// ValueOf alias to syscall/js
func ValueOf(x interface{}) Value {
	if objGo, ok := x.(ObjectFrom); ok {
		return objGo.Value()
	}

	return &value{v: js.ValueOf(x)}
}

// FuncOf alias to syscall/js
func FuncOf(fn func(Value, []Value) interface{}) Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		args2 := make([]Value, len(args))
		for i := range args {
			args2[i] = &value{v: args[i]}
		}
		return fn(&value{v: this}, args2)
	})

	return function{
		Value: &value{v: f.Value},
		f:     f,
	}
}

// CopyBytesToGo alias to syscall/js
func CopyBytesToGo(dst []byte, src Value) (int, error) {
	v := src.(*value)
	return js.CopyBytesToGo(dst, v.v), nil
}

// CopyBytesToJS alias to syscall/js
func CopyBytesToJS(dst Value, src []byte) (int, error) {
	v := dst.(*value)
	return js.CopyBytesToJS(v.v, src), nil
}

// Type alias to syscall/js
func (v *value) Type() Type {
	return Type(v.v.Type())
}

func (v *value) Get(p string) Value {
	if v.err != nil {
		return v
	}

	if !v.v.Truthy() {
		v.err = fmt.Errorf("wasm: get property '%s' on %s", p, v.v.Type().String())
		return v
	}

	// 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			v.err = fmt.Errorf("wasm: panic getting property '%s': %v", p, r)
		}
	}()

	res := v.v.Get(p)
	return &value{v: res}
}

func (v *value) Set(p string, x interface{}) Value {
	if v.err != nil {
		return v
	}
	if !v.v.Truthy() {
		v.err = fmt.Errorf("wasm: set property '%s' on %s", p, v.v.Type().String())
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			v.err = fmt.Errorf("wasm: panic setting property '%s': %v", p, r)
		}
	}()

	v.v.Set(p, unwrap(x))
	return v
}

func (v *value) Call(m string, args ...interface{}) Value {
	if v.err != nil {
		return v
	}
	if !v.v.Truthy() {
		v.err = fmt.Errorf("wasm: call method '%s' on %s", m, v.v.Type().String())
		return v
	}

	processedArgs, err := v.processArgs(args)
	if err != nil {
		v.err = err
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			v.err = fmt.Errorf("wasm: panic calling method '%s': %v", m, r)
		}
	}()

	res := v.v.Call(m, processedArgs...)
	return &value{v: res}
}

func (v *value) Invoke(args ...interface{}) Value {
	vv := v.v.Invoke(fixArgsToGojs(args)...)
	return &value{v: vv}
}

func (v *value) Index(i int) Value {
	return &value{v: v.v.Index(i)}
}

func (v *value) SetIndex(i int, x interface{}) {
	v.v.SetIndex(i, x)
}

func (v *value) Length() int {
	return v.v.Length()
}

func (v *value) New(args ...interface{}) Value {
	if v.err != nil {
		return v
	}
	if v.v.Type() != js.TypeFunction {
		v.err = fmt.Errorf("wasm: new called on non-function type: %s", v.v.Type().String())
		return v
	}

	processedArgs, err := v.processArgs(args)
	if err != nil {
		v.err = err
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			v.err = fmt.Errorf("wasm: panic in new: %v", r)
		}
	}()

	res := v.v.New(processedArgs...)
	return &value{v: res}
}

func (v *value) Float() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}
	if v.v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("wasm: value is not a number (got %s)", v.v.Type())
	}
	return v.v.Float(), nil
}

func (v *value) Int() (int, error) {
	if v.err != nil {
		return 0, v.err
	}

	if v.v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("wasm: value is not a number (got %s)", v.v.Type())
	}

	return v.v.Int(), nil
}

func (v *value) Bool() (bool, error) {
	if v.err != nil {
		return false, v.err
	}
	if v.v.Type() != js.TypeBoolean {
		return false, fmt.Errorf("wasm: value is not a boolean (got %s)", v.v.Type())
	}
	return v.v.Bool(), nil
}

func (v *value) String() (string, error) {
	if v.err != nil {
		return "", v.err
	}
	if v.v.Type() != js.TypeString {
		return "", fmt.Errorf("wasm: value is not a string (got %s)", v.v.Type())
	}
	return v.v.String(), nil
}

// Bytes 将 JS 中的 Uint8Array 复制到 Go 的 []byte 中。
func (v *value) Bytes() ([]byte, error) {
	if v.err != nil {
		return nil, v.err
	}

	if !v.v.Truthy() || v.v.Type() != js.TypeObject {
		return nil, fmt.Errorf("wasm: value is not a Uint8Array (got %s)", v.v.Type())
	}

	lengthVal := v.Get("length")
	if lengthVal.Error() != nil {
		return nil, fmt.Errorf("wasm: could not get length of array: %w", lengthVal.Error())
	}

	length, err := lengthVal.Int()
	if err != nil {
		return nil, fmt.Errorf("wasm: array length is not a number: %w", err)
	}

	dst := make([]byte, length)
	copied := js.CopyBytesToGo(dst, v.v)
	if copied != length {
		return nil, errors.New("wasm: failed to copy all bytes from Uint8Array")
	}

	return dst, nil
}

func (v *value) Truthy() bool {
	return v.v.Truthy()
}

func (v *value) InstanceOf(t Value) bool {
	if vv, ok := t.(*value); ok {
		return v.v.InstanceOf(vv.v)
	}
	return false
}

func (v *value) IsUndefined() bool {
	return v.v.IsUndefined()
}

func (v *value) IsNull() bool {
	return v.v.IsNull()
}

func (v *value) Error() error {
	return v.err
}

// processArgs 展开参数列表中的所有 *SafeValue。
// 如果任何参数本身包含错误，则返回该错误。
func (v *value) processArgs(args []interface{}) ([]interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}
	processed := make([]interface{}, len(args))
	for i, arg := range args {
		if s, ok := arg.(*value); ok {
			if s.err != nil {
				return nil, fmt.Errorf("argument %d has an error: %w", i, s.err)
			}
			processed[i] = s.v
		} else {
			processed[i] = arg
		}
	}
	return processed, nil
}

func fixArgsToGojs(args []interface{}) []interface{} {
	for i := 0; i < len(args); i++ {
		v := args[i]
		if val, ok := v.(value); ok {
			args[i] = val.v // convert to js.Value
		}
		if f, ok := v.(function); ok {
			args[i] = f.f
		}
		// if ta, ok := v.(TypedArray); ok {
		// 	args[i] = js.TypedArray{Value: js.Value(ta.Value)}
		// }
	}
	return args
}

// Release alias to syscall/js
func (c function) Release() {
	// return (*(*sjs.Func)(unsafe.Pointer(&c))).Release()
	c.f.Release()
}

// unwrap an argument if it's a *SafeValue
func unwrap(arg interface{}) interface{} {
	if v, ok := arg.(*value); ok {
		return v.v
	}
	return arg
}
