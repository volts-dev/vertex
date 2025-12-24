package js

import (
	"errors"
	"sync"
)

var (
	ErrNotAnArray = errors.New("The given value must be an Array")
	//ErrNotImplemented = errors.New("Browser not implemented Array")
)

func init() {
	//js.RegisterInterface(GetArrayInterface)
}

var singletonArray sync.Once

var arrayinterface Value

// GetArrayInterface get the JS interface Array
func GetArrayInterface() Value {
	singletonArray.Do(func() {
		var err error
		if arrayinterface = Global().Get("Array"); err != nil {
			arrayinterface = Undefined()
		}
		Register(arrayinterface, func(v Value) (interface{}, error) {
			return NewArrayFromJSObject(v)
		})
	})

	return arrayinterface
}

// Array struct
type Array struct {
	Object
}

type ArrayFrom interface {
	Array_() Array
}

func (a Array) Array_() Array {
	return a
}

func NewEmpty(size int) (Array, error) {

	var a Array
	var obj Value
	var err error
	if ai := GetArrayInterface(); !ai.IsUndefined() {

		if obj = ai.New(ValueOf(size)); obj.Error() == nil {
			a.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func From(iterable interface{}, f ...func(interface{}) interface{}) (Array, error) {
	var a Array
	var err error
	var obj Value
	var opts []interface{}
	var jsfunc Func

	if ai := GetArrayInterface(); !ai.IsUndefined() {
		opts = append(opts, ValueOf(iterable))
		if f != nil && len(f) == 1 {
			jsfunc = FuncOf(func(this Value, args []Value) interface{} {
				b := f[0](GoValue_(args[0]))
				return ValueOf(b)
			})
			opts = append(opts, jsfunc)

		}

		if obj = ai.Call("from", opts...); obj.Error() == nil {
			a.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func Of(values ...interface{}) (Array, error) {

	var a Array
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, ValueOf(value))
	}
	if ai := GetArrayInterface(); !ai.IsUndefined() {
		a.SetObjectValue(ai.Call("of", arrayJS...))
		return a, nil
	}
	return a, ErrNotImplemented

}

func NewArray(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}
	var obj Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, ValueOf(value))
	}
	if ai := GetArrayInterface(); !ai.IsUndefined() {

		if obj = ai.New(arrayJS...); obj.Error() == nil {
			a.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err

}

func NewArrayFromJSObject(obj Value) (Array, error) {
	var a Array
	var err error
	if ai := GetArrayInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnArray
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Array) Length() (int, error) {

	return a.GetAttributeInt("length")

}

func (a Array) Concat(a2 Array) (Array, error) {

	var err error
	var obj Value
	var newArr Array

	if obj = a.Call("concat", a2.GetObjectValue()); obj.Error() == nil {

		newArr, err = NewArrayFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) CopyWithin(cible int, opts ...int) (Array, error) {

	var err error
	var obj Value
	var newArr Array
	var arrayJS []interface{}

	arrayJS = append(arrayJS, ValueOf(cible))

	for _, opt := range opts {
		arrayJS = append(arrayJS, ValueOf(opt))
	}

	if obj = a.Call("copyWithin", arrayJS...); obj.Error() == nil {

		newArr, err = NewArrayFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Entries() (Iterator, error) {
	var err error
	var obj Value
	var iter Iterator

	if obj = a.Call("entries"); obj.Error() == nil {
		iter, err = NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (a Array) Every(f func(interface{}) bool) (bool, error) {
	var err error
	var obj Value
	var result bool

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("every", jsfunc); obj.Error() == nil {
		if obj.Type() == TypeBoolean {
			return obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

// Fill (value, begin, end)
func (a Array) Fill(i interface{}, opts ...int) error {
	var arrayJS []interface{}
	arrayJS = append(arrayJS, ValueOf(i))

	for _, opt := range opts {
		arrayJS = append(arrayJS, ValueOf(opt))
	}

	return a.Call("fill", arrayJS...).Error()
}

func (a Array) Filter(f func(interface{}) bool) (Array, error) {
	var err error
	var obj Value
	var newArr Array

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("filter", jsfunc); obj.Error() == nil {
		newArr, err = NewArrayFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Find(f func(interface{}) bool) (interface{}, error) {
	var err error
	var obj Value
	var i interface{}

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("find", jsfunc); obj.Error() == nil {
		if obj.Type() != TypeUndefined {
			i = GoValue_(obj)
		}

	}
	jsfunc.Release()
	return i, err
}

func (a Array) FindIndex(f func(interface{}) bool) (int, error) {
	var err error
	var obj Value
	var index int = -1

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("findIndex", jsfunc); obj.Error() == nil {
		if obj.Type() == TypeNumber {
			return obj.Int()
		}
	}
	jsfunc.Release()
	return index, err
}

func (a Array) Flat(opts ...int) (Array, error) {
	var err error
	var arrayJS []interface{}
	var obj Value
	var newArr Array

	if len(opts) < 2 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, ValueOf(opt))
		}
	}

	if obj = a.Call("flat", arrayJS...); obj.Error() == nil {
		newArr, err = NewArrayFromJSObject(obj)
	}
	return newArr, err
}

func (a Array) FlatMap(f func(interface{}, int) interface{}) (Array, error) {
	var err error
	var obj Value
	var newArr Array

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]), ValueToInt(args[1]))
		return b
	})

	if obj = a.Call("flatMap", jsfunc); obj.Error() == nil {
		newArr, err = NewArrayFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) ForEach(f func(interface{})) error {
	var err error

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		f(GoValue_(args[0]))
		return nil
	})

	err = a.Call("forEach", jsfunc).Error()
	jsfunc.Release()
	return err
}

func (a Array) Includes(i interface{}) (bool, error) {
	var err error
	var obj Value
	var result bool

	if obj = a.Call("includes", ValueOf(i)); obj.Error() == nil {
		if obj.Type() == TypeBoolean {
			return obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err
}

func (a Array) IndexOf(i interface{}) (int, error) {
	var err error
	var obj Value
	var index int = -1

	if obj = a.Call("indexOf", ValueOf(i)); obj.Error() == nil {
		if obj.Type() == TypeNumber {
			return obj.Int()
		}
	}

	return index, err
}

func IsArray(bobj Object) (bool, error) {

	var err error
	var result bool
	var obj Value

	if ai := GetArrayInterface(); !ai.IsUndefined() {

		if obj = ai.Call("isArray", bobj.GetObjectValue()); obj.Error() == nil {
			if obj.Type() == TypeBoolean {
				return obj.Bool()
			} else {
				err = ErrObjectNotBool
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return result, err
}

func (a Array) Join(separator string) (string, error) {
	var err error
	var result string
	var obj Value

	if obj = a.Call("join", ValueOf(separator)); obj.Error() == nil {
		if obj.Type() == TypeString {
			return obj.String()
		} else {
			err = ErrObjectNotString
		}
	}
	return result, err
}

func (a Array) Keys() (Iterator, error) {
	var err error
	var obj Value
	var iter Iterator

	if obj = a.Call("keys"); obj.Error() == nil {
		iter, err = NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (a Array) LastIndexOf(i interface{}) (int, error) {
	var err error
	var obj Value
	var index int = -1

	if obj = a.Call("lastIndexOf", ValueOf(i)); obj.Error() == nil {
		if obj.Type() == TypeNumber {
			return obj.Int()
		}
	}

	return index, err
}

func (a Array) Map(f func(interface{}) interface{}) (Array, error) {
	var err error
	var obj Value
	var newArr Array

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("map", jsfunc); obj.Error() == nil {
		newArr, err = NewArrayFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Pop() error {
	return a.Call("pop").Error()
}

func (a Array) Push(i interface{}) (int, error) {
	var err error
	var obj Value
	var index int = -1

	if obj = a.Call("push", ValueOf(i)); obj.Error() == nil {
		if obj.Type() == TypeNumber {
			return obj.Int()
		}
	}

	return index, err

}

func (a Array) Reduce(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj Value
	var newValue interface{}
	var argCall []interface{}

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, ValueOf(args[i]))
		}

		b := f(GoValue_(args[0]), GoValue_(args[1]), arrayJS...)
		return ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, ValueOf(initialValue[0]))
	}
	if obj = a.Call("reduce", argCall...); obj.Error() == nil {
		newValue = GoValue_(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) ReduceRight(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj Value
	var newValue interface{}
	var argCall []interface{}

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, ValueOf(args[i]))
		}

		b := f(GoValue_(args[0]), GoValue_(args[1]), arrayJS...)
		return ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, ValueOf(initialValue[0]))
	}
	if obj = a.Call("reduceRight", argCall...); obj.Error() == nil {
		newValue = GoValue_(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) Reverse() error {
	return a.Call("reverse").Error()
}

func (a Array) Shift() (interface{}, error) {
	var err error
	var obj Value
	var i interface{}
	if obj = a.Call("shift"); obj.Error() == nil {
		i = GoValue_(obj)
	}
	return i, err
}

func (a Array) Slice(opts ...int) (Array, error) {

	var err error
	var obj Value
	var newArr Array
	var arrayJS []interface{}
	for _, opt := range opts {
		arrayJS = append(arrayJS, ValueOf(opt))
	}

	if obj = a.Call("slice", arrayJS...); obj.Error() == nil {

		newArr, err = NewArrayFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Some(f func(interface{}) bool) (bool, error) {
	var err error
	var obj Value
	var result bool

	jsfunc := FuncOf(func(this Value, args []Value) interface{} {
		b := f(GoValue_(args[0]))
		return ValueOf(b)
	})

	if obj = a.Call("some", jsfunc); obj.Error() == nil {
		if obj.Type() == TypeBoolean {
			return obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

func (a Array) Sort() error {
	return a.Call("sort").Error()
}

func (a Array) Splice(begin, suppress int, values ...interface{}) error {
	var arrayJS []interface{}
	arrayJS = append(arrayJS, ValueOf(begin), ValueOf(suppress))

	for _, value := range values {
		arrayJS = append(arrayJS, ValueOf(value))
	}

	return a.Call("splice", arrayJS...).Error()
}

func (a Array) ToLocaleString() (string, error) {

	return a.GetAttributeString("toLocaleString")
}

func (a Array) Unshift(values ...interface{}) (int, error) {

	var err error
	var arrayJS []interface{}
	var obj Value
	var index int = -1

	for _, value := range values {
		arrayJS = append(arrayJS, ValueOf(value))
	}
	if obj = a.Call("unshift", arrayJS...); obj.Error() == nil {
		if obj.Type() == TypeNumber {
			return obj.Int()
		}
	}
	return index, err
}

func (a Array) Values() (Iterator, error) {
	var err error
	var obj Value
	var iter Iterator

	if obj = a.Call("values"); obj.Error() == nil {
		iter, err = NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (a Array) SetValue(index int, i interface{}) error {

	var obj interface{}
	if objGo, ok := i.(ObjectFrom); ok {

		obj = objGo.GetObjectValue()
	} else {
		obj = i
	}

	a.GetObjectValue().SetIndex(index, obj)
	return nil
}

func (a Array) GetValue(index int) (interface{}, error) {

	obj := a.GetObjectValue().Index(index)
	return GoValue_(obj), nil
}

func NewArray_(values ...interface{}) Array {
	a, err := NewArray(values...)

	if err != nil {
		a.Debug(err.Error())
	}
	return a
}

func Of_(values ...interface{}) Array {
	a, err := Of(values...)

	if err != nil {
		a.Debug(err.Error())
	}
	return a
}

func NewEmpty_(size int) Array {
	a, err := NewEmpty(size)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}
func ArrayFrom_(iterable interface{}, f ...func(interface{}) interface{}) Array {

	a, err := From(iterable, f...)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}

func NewArrayFromJSObject_(obj Value) Array {
	a, err := NewArrayFromJSObject(obj)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}
func (a Array) Concat_(a2 Array) Array {

	arr, err := a.Concat(a2)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) CopyWithin_(cible int, opts ...int) Array {
	arr, err := a.CopyWithin(cible, opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr

}

func (a Array) Filter_(f func(interface{}) bool) Array {
	arr, err := a.Filter(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Flat_(opts ...int) Array {
	arr, err := a.Flat(opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) FlatMap_(f func(interface{}, int) interface{}) Array {
	arr, err := a.FlatMap(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Map_(f func(interface{}) interface{}) Array {
	arr, err := a.Map(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Slice_(opts ...int) Array {
	arr, err := a.Slice(opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}
