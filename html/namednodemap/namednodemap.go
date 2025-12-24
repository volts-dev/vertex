package namednodemap

// https://developer.mozilla.org/fr/docs/Web/API/NamedNodeMap

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/attr"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var namednodemapinterface js.Value

// NamedNodeMap struct
type NamedNodeMap struct {
	js.Object
}

type NamedNodeMapFrom interface {
	NamedNodeMap_() NamedNodeMap
}

func (n NamedNodeMap) NamedNodeMap_() NamedNodeMap {
	return n
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if namednodemapinterface = js.Global().Get("NamedNodeMap"); namednodemapinterface.Error() != nil {
			namednodemapinterface = js.Undefined()
		}
		js.Register(namednodemapinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return namednodemapinterface
}

func NewFromJSObject(obj js.Value) (NamedNodeMap, error) {
	var n NamedNodeMap
	var err error
	if nli := GetInterface(); !nli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(nli) {
				n.SetObjectValue(obj)

			} else {
				err = ErrNotANamedNodeMap
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return n, err
}

func (n NamedNodeMap) Item(index int) (attr.Attr, error) {

	return attr.NewFromJSObject(n.GetObjectValue().Index(index))
}

func (n NamedNodeMap) GetNamedItem(name string) (attr.Attr, error) {
	var elemObject js.Value
	var newAttr attr.Attr
	var err error

	if elemObject = n.Call("getNamedItem", js.ValueOf(name)); elemObject.Error() == nil {
		if elemObject.IsUndefined() {
			err = ErrNotNameAttr
		} else {
			newAttr, err = attr.NewFromJSObject(elemObject)
		}
	}

	return newAttr, err
}

func (n NamedNodeMap) SetNamedItem(a attr.Attr) error {
	var err error
	err = n.Call("setNamedItem", a.GetObjectValue()).Error()
	return err
}

func (n NamedNodeMap) RemoveNamedItem(name string) error {
	var err error
	err = n.Call("removeNamedItem", js.ValueOf(name)).Error()
	return err
}

func (n NamedNodeMap) GetNamedItemNS(namespace string, name string) (attr.Attr, error) {
	var err error
	var elemObject js.Value
	var newAttr attr.Attr

	if elemObject = n.Call("getNamedItemNS", js.ValueOf(namespace), js.ValueOf(name)); elemObject.Error() == nil {

		if elemObject.IsUndefined() {
			err = ErrNotNameAttr

		} else {

			newAttr, err = attr.NewFromJSObject(elemObject)

		}
	}

	return newAttr, err

}

func (n NamedNodeMap) SetNamedItemNS(a attr.Attr) error {
	var err error
	err = n.Call("setNamedItemNS", a.GetObjectValue()).Error()
	return err
}

func (n NamedNodeMap) RemoveNamedItemNS(namespace string, name string) error {
	var err error
	err = n.Call("removeNamedItemNS", js.ValueOf(namespace), js.ValueOf(name)).Error()
	return err
}
