package attr

import (
	"sync"

	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var attrinterface js.Value

// GetInterface get the JS interface Attr
func GetInterface() js.Value {

	singleton.Do(func() {

		if attrinterface = js.Global().Get("Attr"); attrinterface.Error() != nil {
			attrinterface = js.Undefined()
		}
		js.Register(attrinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return attrinterface
}

type Attr struct {
	node.Node
}

type AttrFrom interface {
	Attr_() Attr
}

func (a Attr) Attr_() Attr {
	return a
}

func NewFromJSObject(obj js.Value) (Attr, error) {
	var a Attr
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnAttr
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Attr) Name() (string, error) {
	return a.GetAttributeString("name")
}

func (a Attr) NamespaceURI() (string, error) {
	return a.GetAttributeString("namespaceURI")
}

func (a Attr) LocalName() (string, error) {
	return a.GetAttributeString("localName")
}

func (a Attr) Prefix() (string, error) {
	return a.GetAttributeString("prefix")
}

func (a Attr) Value() (string, error) {
	return a.GetAttributeString("value")
}

//use element.OwnerElementForAttr
//func (a Attr) OwnerElementObjet()
