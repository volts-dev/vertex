package array

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/helper"
	"github.com/volts-dev/vertex/js/iterator"
)

func init() {
	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var arrayinterface js.Value

// GetInterface get the JS interface Array
func GetInterface() js.Value {
	singleton.Do(func() {
		var err error
		if arrayinterface = js.Global().Get("Array"); err != nil {
			arrayinterface = js.Undefined()
		}
		js.Register(arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return arrayinterface
}

// Array struct
type Array struct {
	js.Object
}

type ArrayFrom interface {
	Array_() Array
}

func (a Array) Array_() Array {
	return a
}

func NewEmpty(size int) (Array, error) {

	var a Array
	var obj js.Value
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.New(js.ValueOf(size)); obj.Error() == nil {
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
	var obj js.Value
	var opts []interface{}
	var jsfunc js.Func

	if ai := GetInterface(); !ai.IsUndefined() {
		opts = append(opts, js.ValueOf(iterable))
		if f != nil && len(f) == 1 {
			jsfunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				b := f[0](helper.GoValue_(args[0]))
				return js.ValueOf(b)
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
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if ai := GetInterface(); !ai.IsUndefined() {
		a.SetObjectValue(ai.Call("of", arrayJS...))
		return a, nil
	}
	return a, ErrNotImplemented

}

func New(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}
	var obj js.Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.New(arrayJS...); obj.Error() == nil {
			a.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err

}

func NewFromJSObject(obj js.Value) (Array, error) {
	var a Array
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
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
	var obj js.Value
	var newArr Array

	if obj = a.Call("concat", a2.GetObjectValue()); obj.Error() == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) CopyWithin(cible int, opts ...int) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(cible))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj = a.Call("copyWithin", arrayJS...); obj.Error() == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = a.Call("entries"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (a Array) Every(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("every", jsfunc); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

// Fill (value, begin, end)
func (a Array) Fill(i interface{}, opts ...int) error {
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(i))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	return a.Call("fill", arrayJS...).Error()
}

func (a Array) Filter(f func(interface{}) bool) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("filter", jsfunc); obj.Error() == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Find(f func(interface{}) bool) (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("find", jsfunc); obj.Error() == nil {
		if obj.Type() != js.TypeUndefined {
			i = helper.GoValue_(obj)
		}

	}
	jsfunc.Release()
	return i, err
}

func (a Array) FindIndex(f func(interface{}) bool) (int, error) {
	var err error
	var obj js.Value
	var index int = -1

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("findIndex", jsfunc); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}
	jsfunc.Release()
	return index, err
}

func (a Array) Flat(opts ...int) (Array, error) {
	var err error
	var arrayJS []interface{}
	var obj js.Value
	var newArr Array

	if len(opts) < 2 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if obj = a.Call("flat", arrayJS...); obj.Error() == nil {
		newArr, err = NewFromJSObject(obj)
	}
	return newArr, err
}

func (a Array) FlatMap(f func(interface{}, int) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]), helper.ValueToInt(args[1]))
		return b
	})

	if obj = a.Call("flatMap", jsfunc); obj.Error() == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) ForEach(f func(interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(helper.GoValue_(args[0]))
		return nil
	})

	err = a.Call("forEach", jsfunc).Error()
	jsfunc.Release()
	return err
}

func (a Array) Includes(i interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = a.Call("includes", js.ValueOf(i)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (a Array) IndexOf(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1

	if obj = a.Call("indexOf", js.ValueOf(i)); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return index, err
}

func IsArray(bobj js.Object) (bool, error) {

	var err error
	var result bool
	var obj js.Value

	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.Call("isArray", bobj.GetObjectValue()); obj.Error() == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				err = js.ErrObjectNotBool
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
	var obj js.Value

	if obj = a.Call("join", js.ValueOf(separator)); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrObjectNotString
		}
	}
	return result, err
}

func (a Array) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = a.Call("keys"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (a Array) LastIndexOf(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1

	if obj = a.Call("lastIndexOf", js.ValueOf(i)); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return index, err
}

func (a Array) Map(f func(interface{}) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("map", jsfunc); obj.Error() == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Pop() error {
	return a.Call("pop").Error()
}

func (a Array) Push(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1

	if obj = a.Call("push", js.ValueOf(i)); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return index, err

}

func (a Array) Reduce(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj js.Value
	var newValue interface{}
	var argCall []interface{}

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, js.ValueOf(args[i]))
		}

		b := f(helper.GoValue_(args[0]), helper.GoValue_(args[1]), arrayJS...)
		return js.ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj = a.Call("reduce", argCall...); obj.Error() == nil {
		newValue = helper.GoValue_(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) ReduceRight(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj js.Value
	var newValue interface{}
	var argCall []interface{}

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, js.ValueOf(args[i]))
		}

		b := f(helper.GoValue_(args[0]), helper.GoValue_(args[1]), arrayJS...)
		return js.ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj = a.Call("reduceRight", argCall...); obj.Error() == nil {
		newValue = helper.GoValue_(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) Reverse() error {
	return a.Call("reverse").Error()
}

func (a Array) Shift() (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}
	if obj = a.Call("shift"); obj.Error() == nil {
		i = helper.GoValue_(obj)
	}
	return i, err
}

func (a Array) Slice(opts ...int) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{}
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj = a.Call("slice", arrayJS...); obj.Error() == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Some(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(helper.GoValue_(args[0]))
		return js.ValueOf(b)
	})

	if obj = a.Call("some", jsfunc); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
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
	arrayJS = append(arrayJS, js.ValueOf(begin), js.ValueOf(suppress))

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	return a.Call("splice", arrayJS...).Error()
}

func (a Array) ToLocaleString() (string, error) {

	return a.GetAttributeString("toLocaleString")
}

func (a Array) Unshift(values ...interface{}) (int, error) {

	var err error
	var arrayJS []interface{}
	var obj js.Value
	var index int = -1

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if obj = a.Call("unshift", arrayJS...); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}
	return index, err
}

func (a Array) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = a.Call("values"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (a Array) SetValue(index int, i interface{}) error {

	var obj interface{}
	if objGo, ok := i.(js.ObjectFrom); ok {

		obj = objGo.Value()
	} else {
		obj = i
	}

	a.GetObjectValue().SetIndex(index, obj)
	return nil
}

func (a Array) GetValue(index int) (interface{}, error) {

	obj := a.GetObjectValue().Index(index)
	return helper.GoValue_(obj), nil
}
