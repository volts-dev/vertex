//go:build js && wasm

package js

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"syscall/js"
	"unsafe"
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

var (
	console js.Value

	global    = js.Global()
	null      = js.Null()
	undefined = js.Undefined()
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

func init() {
	console = js.Global().Get("console")
}

func RecoverHandler(r any) {
	console.Call("log", fmt.Sprintf("wasm: panic : %v", r))
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		console.Call("log", fmt.Sprintf("  %s %d\n", file, line))
	}
}

// Undefined alias to syscall/js
func Undefined() Value {
	return &value{v: undefined}
}

// Null alias to syscall/js
func Null() Value {
	return &value{v: null}
}

func Global() Value {
	return &value{v: global}
}

// ValueOf alias to syscall/js
func ValueOf(x any) Value {
	if objGo, ok := x.(ObjectFrom); ok {
		return objGo.GetObjectValue()
	}

	switch x := x.(type) {
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr, unsafe.Pointer,
		float32, float64, string, []any, map[string]any:
		return &value{v: js.ValueOf(x)}
	default:
		v := reflect.ValueOf(x)
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface:
			if v.IsNil() {
				return &value{v: null}
			}
			return ValueOf(v.Elem())

		case reflect.Struct:
			t := v.Type()
			s := global.Get("Object").New()
			n := v.NumField()
			for i := 0; i < n; i++ {
				if f := v.Field(i); f.CanInterface() {
					k := nameOf(t.Field(i))
					s.Set(k, ValueOf(f))
				}
			}
			return &value{v: s}
		default:
			panic("ValueOf: invalid value")
		}
	}
}

// FuncOf alias to syscall/js
func FuncOf(fn func(Value, []Value) any) Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) any {
		args2 := make([]Value, len(args))
		for i := range args {
			args2[i] = &value{v: args[i]}
		}

		result := fn(&value{v: this}, args2)
		return unwrap(result)
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

func (v *value) Equal(other Value) bool {
	return v.v.Equal(other.(*value).v)
}

// Type alias to syscall/js
func (v *value) Type() Type {
	return Type(v.v.Type())
}

func (v *value) Get(p string) (result Value) {
	// 如果已经有错误，返回包含错误的新值
	if v.err != nil {
		return &value{err: v.err}
	}

	// 检查值的有效性
	if !v.v.Truthy() {
		return &value{v: v.v, err: fmt.Errorf("wasm: get property '%s' on %s", p, v.v.Type().String())}
	}

	// 捕获 panic 并返回错误
	defer func() {
		if r := recover(); r != nil {
			result = &value{v: v.v, err: fmt.Errorf("wasm: panic getting property '%s': %v", p, r)}
			RecoverHandler(r)
		}
	}()

	return &value{v: v.v.Get(p)}
}

func (v *value) Set(p string, x any) (result Value) {
	if v.err != nil {
		return v
	}

	if !v.v.Truthy() {
		v.err = fmt.Errorf("wasm: set property '%s' on %s", p, v.v.Type().String())
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			result = &value{v: v.v, err: fmt.Errorf("wasm: panic setting property '%s': %v", p, r)}
			RecoverHandler(r)
		}
	}()

	//v.v.Set(p, x)
	v.v.Set(p, unwrap(x))
	return v
}

func (v *value) Call(m string, args ...any) (result Value) {
	if v.err != nil {
		return v
	}

	if !v.v.Truthy() {
		v.err = fmt.Errorf("wasm: call method '%s' on %s", m, v.v.Type().String())
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			result = &value{v: v.v, err: fmt.Errorf("wasm: panic calling method '%s': %v", m, r)}
			RecoverHandler(r)
		}
	}()

	processedArgs, err := v.processArgs(args)
	if err != nil {
		v.err = err
		return v
	}

	return &value{
		v: v.v.Call(m, processedArgs...),
	}
}

func (v *value) Invoke(args ...any) Value {
	vv := v.v.Invoke(fixArgsToGojs(args)...)
	return &value{v: vv}
}

func (v *value) Index(i int) Value {
	return &value{v: v.v.Index(i)}
}

func (v *value) SetIndex(i int, x any) {
	v.v.SetIndex(i, unwrap(x))
}

func (v *value) Length() int {
	return v.v.Length()
}

func (v *value) New(args ...any) (result Value) {
	if v.err != nil {
		return v
	}

	if v.v.Type() != js.TypeFunction {
		v.err = fmt.Errorf("wasm: new called on non-function type: %s", v.v.Type().String())
		return v
	}

	defer func() {
		if r := recover(); r != nil {
			result = &value{v: v.v, err: fmt.Errorf("wasm: panic in new: %v", r)}
			RecoverHandler(r)
		}
	}()

	processedArgs, err := v.processArgs(args)
	if err != nil {
		v.err = err
		return v
	}

	res := v.v.New(processedArgs...)
	return &value{v: res}
}

func (v *value) Float() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}

	if v.v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("value is not a number (got %s)", v.v.Type())
	}

	return v.v.Float(), nil
}

func (v *value) Int() (int, error) {
	if v.err != nil {
		return 0, v.err
	}

	if v.v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("value is not a number (got %s)", v.v.Type())
	}

	return v.v.Int(), nil
}

func (v *value) Bool() (bool, error) {
	if v.err != nil {
		return false, v.err
	}

	if v.v.Type() != js.TypeBoolean {
		return false, fmt.Errorf("value is not a boolean (got %s)", v.v.Type())
	}

	return v.v.Bool(), nil
}

func (v *value) String() (string, error) {
	if v.err != nil {
		return "", v.err
	}

	if v.v.Type() != js.TypeString {
		return "", fmt.Errorf("value is not a string (got %s)", v.v.Type())
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

func (v *value) IsNaN() bool {
	return v.v.IsNaN()
}

func (v *value) Error() error {
	return v.err
}

// processArgs 展开参数列表中的所有 *SafeValue。
// 如果任何参数本身包含错误，则返回该错误。
func (v *value) processArgs(args []any) ([]any, error) {
	if v.err != nil {
		return nil, v.err
	}

	if len(args) == 0 {
		return nil, nil
	}

	processed := make([]any, len(args))
	for i, arg := range args {
		processed[i] = unwrap(arg)
	}

	return processed, nil
}

func fixArgsToGojs(args []any) []any {
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
func unwrap(val any) any {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case *value:
		return v.v
	case function:
		return v.f
	case error:
		return js.ValueOf(v.Error())
	/*case map[string]any:
		m := make(map[string]any, len(v))
		for key, value := range v {
			m[key] = value
		}
		return js.ValueOf(m)

	case []any:
	s := make([]any, len(v))
	for i, value := range v {
		s[i] = value)
	}
	return js.ValueOf(s)
	*/
	case []string:
		s := make([]any, len(v))
		for i, value := range v {
			s[i] = value
		}
		return js.ValueOf(s)
	default:
		return js.ValueOf(val)
	}
}

// nameOf returns the JS tag name, otherwise the field name.
func nameOf(sf reflect.StructField) string {
	name := sf.Tag.Get("js")
	if name == "" {
		name = sf.Tag.Get("json")
	}
	if name == "" {
		return sf.Name
	}
	return name
}
