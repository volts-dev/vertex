package domstringlist

//

import (
	"sync"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"

	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var domstringlistinterface js.Value

// DOMRectLists struct
type DOMStringList struct {
	js.Object
}

type DOMStringListFrom interface {
	DOMStringList_() DOMStringList
}

func (d DOMStringList) DOMStringList_() DOMStringList {
	return d
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if domstringlistinterface = js.Global().Get("DOMStringList"); domstringlistinterface.Error() != nil {
			domstringlistinterface = js.Undefined()
		}
		js.Register(domstringlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domstringlistinterface
}

func NewFromJSObject(obj js.Value) (DOMStringList, error) {
	var d DOMStringList
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				d.SetObjectValue(obj)

			} else {
				err = ErrNotAnDOMStringList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMStringList) Item(index int) js.Value {
	var obj js.Value
	obj = d.GetObjectValue().Index(index)
	return obj
}

func (d DOMStringList) Contains(search string) (bool, error) {

	var err error
	var obj js.Value
	var result bool
	if obj = d.Call("contains", js.ValueOf(search)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err
}
