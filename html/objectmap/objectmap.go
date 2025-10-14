package objectmap

import (
	"sync"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/helper"
	"github.com/volts-dev/vertex/js/object"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js/iterator"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var mapinterface js.Value

// GetInterface get the JS interface of object channel
func GetInterface() js.Value {

	singleton.Do(func() {

		if mapinterface = js.Global().Get("Map"); mapinterface.Error() != nil {
			mapinterface = js.Undefined()
		}
		js.Register(mapinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return mapinterface
}

// ObjectMap
type ObjectMap struct {
	js.Object
}

type ObjectMapFrom interface {
	ObjectMap_() ObjectMap
}

func (o ObjectMap) ObjectMap_() ObjectMap {
	return o
}

func NewFromJSObject(obj js.Value) (ObjectMap, error) {
	var o ObjectMap
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				o.SetObjectValue(obj)

			} else {
				err = ErrNotAMap
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func NewFromBaseObject(b js.Object) (ObjectMap, error) {

	return New(b)
}

func New(values ...interface{}) (ObjectMap, error) {

	var o ObjectMap
	var err error
	var obj js.Value
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	if omi := GetInterface(); !omi.IsUndefined() {

		if obj = omi.New(arrayJS...); obj.Error() == nil {
			o.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o ObjectMap) Clear() error {
	var err error
	err = o.Call("clear").Error()
	return err
}

func (o ObjectMap) Delete(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = o.Call("delete", js.ValueOf(key)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = o.Call("entries"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) ForEach(f func(value, index interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		f(helper.GoValue_(args[0]), helper.GoValue_(args[1]))
		return nil
	})

	err = o.Call("forEach", jsfunc).Error()
	jsfunc.Release()
	return err
}

func (o ObjectMap) Get(key interface{}) (interface{}, error) {
	var err error
	var obj js.Value
	var result interface{}
	if obj = o.Call("get", js.ValueOf(key)); obj.Error() == nil {
		result, err = helper.GoValue(obj)
	}
	return result, err
}

func (o ObjectMap) Has(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = o.Call("has", js.ValueOf(key)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = o.Call("keys"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Set(key interface{}, value interface{}) error {
	var err error
	err = o.Call("set", js.ValueOf(key), js.ValueOf(value)).Error()
	return err
}

func (o ObjectMap) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = o.Call("values"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Size() (int, error) {
	return o.GetAttributeInt("size")
}
