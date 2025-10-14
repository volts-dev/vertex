package iterator

import (
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/helper"
)

// Iterator iterator
type Iterator struct {
	js.Object
}

type IteratorFrom interface {
	Iterator_() Iterator
}

func (i Iterator) Iterator_() Iterator {
	return i
}

func NewFromJSObject(obj js.Value) (Iterator, error) {
	var i Iterator
	var err error
	var objf js.Value

	symbol := js.Global().Get("Symbol")
	it := symbol.Get("iterator")

	if objf = obj.Invoke(it); objf.Error() == nil {
		if objf.Type() == js.TypeFunction {
			i.SetObjectValue(obj)
		} else {
			err = NotAnIterator
		}
	}

	return i, err
}

func pairValues(obj js.Value) (interface{}, interface{}) {

	var value interface{}
	var index interface{}

	if obj.Type() == js.TypeObject {
		if obj.Length() == 2 {

			index = helper.GoValue_(obj.Index(0))

			value = helper.GoValue_(obj.Index(1))

		}

	}
	return index, value
}

/* Parse using

for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

}
*/

func (i Iterator) Next() (interface{}, interface{}, error) {

	var err error
	var done bool = true
	var obj, doneobj, valueobj js.Value
	var index interface{}
	var value interface{}

	if obj = i.Call("next"); obj.Error() == nil {

		if doneobj = obj.Get("done"); doneobj.Error() == nil {
			if doneobj.Type() == js.TypeBoolean {
				done = helper.ValueToBool(doneobj)
			} else {
				err = js.ErrObjectNotBool
			}
		}
		if done {
			err = EOI

		} else {

			if valueobj = obj.Get("value"); valueobj.Error() == nil {
				if valueobj.Type() == js.TypeObject {
					index, value = pairValues(valueobj)
				} else {
					value, err = helper.GoValue(valueobj)
				}

			}
		}

	}
	return index, value, err
}
