package object

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/objectmap"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var objectinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if objectinterface = js.Global().Get("Object"); objectinterface.Error() != nil {
			objectinterface = js.Undefined()
		}
		js.Register(objectinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return objectinterface
}

// Object struct
type Object struct {
	js.Value
}

type ObjectFrom interface {
	Object_() Object
}

func (o Object) Object_() Object {
	return o
}

func New() (Object, error) {
	var o Object
	var err error
	var obj js.Value
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.New(); obj.Error() == nil {
			o.Value = obj
		}

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func NewFromJSObject(obj js.Value) (Object, error) {
	var o Object
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				o.Value = obj

			} else {
				err = ErrNotAnObject
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o Object) Keys() (js.Array, error) {

	var err error
	var obj js.Value
	var newArr js.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj = ai.Call("keys", o.Value); obj.Error() == nil {
			newArr, err = js.NewArrayFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Values() (js.Array, error) {

	var err error
	var obj js.Value
	var newArr js.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj = ai.Call("values", o.Value); obj.Error() == nil {
			newArr, err = js.NewArrayFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Map() (objectmap.ObjectMap, error) {
	var err error
	var obj js.Value
	var newMap objectmap.ObjectMap

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj = ai.Call("entries", o.Value); obj.Error() == nil {
			newMap, err = objectmap.New(obj)

		}

	}
	return newMap, err
}
