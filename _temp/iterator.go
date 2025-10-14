package js

import (
	"errors"
)

// Iterator iterator
type Iterator struct {
	Object
}

var (
	//ErrNotImplemented ErrNotImplemented error
	EOI           = errors.New("End of iterator")
	NotAnIterator = errors.New("Object Not an iterator")
)

type IteratorFrom interface {
	Iterator_() Iterator
}

func (i Iterator) Iterator_() Iterator {
	return i
}

func ToIterator(obj Value) (Iterator, error) {
	var i Iterator
	var err error
	var objf Value

	symbol := Global().Get("Symbol")
	it := symbol.Get("iterator")

	if objf = obj.Invoke(it); objf.Error() == nil {
		if objf.Type() == TypeFunction {
			i.SetValue(obj)
		} else {
			err = NotAnIterator
		}
	}

	return i, err
}

func pairValues(obj Value) (interface{}, interface{}) {

	var value interface{}
	var index interface{}

	if obj.Type() == TypeObject {
		if obj.Length() == 2 {

			index = GoValueExt(obj.Index(0))

			value = GoValueExt(obj.Index(1))

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
	var obj, doneobj, valueobj Value
	var index interface{}
	var value interface{}

	if obj = i.Call("next"); obj.Error() == nil {

		if doneobj = obj.Get("done"); doneobj.Error() == nil {
			if doneobj.Type() == TypeBoolean {
				done = ValueToBool(doneobj)
			} else {
				err = ErrObjectNotBool
			}
		}
		if done {
			err = EOI

		} else {

			if valueobj = obj.Get("value"); valueobj.Error() == nil {
				if valueobj.Type() == TypeObject {
					index, value = pairValues(valueobj)
				} else {
					value, err = GoValue(valueobj)
				}

			}
		}

	}
	return index, value, err
}
