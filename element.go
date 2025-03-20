//go:build js && wasm

package vertex

import (
	"errors"
	"sync"

	"github.com/volts-dev/volts/vertex/core/js"
)

var (
	ErrNotAnElement    = errors.New("Object is not an Element")
	ErrElementNoChilds = errors.New("Element has no childs")
	ErrAttributeEmpty  = errors.New("Attribute is empty")
	ErrInsertAdjacent  = errors.New("Insert Adjacent failed")
	//ErrElementNotFound  = errors.New("Element not Found")
	//ErrElementsNotFound = errors.New("Elements not Found")
	ErrSendUnknownType = errors.New("Unknown type send data provide to send method")
)

type (
	IElement interface {
		// JSValue returns the javascript value linked to the element.
		Value() js.Value

		// Reports whether the element is mounted.
		Mounted() bool

		parent() IElement
		setParent(IElement) IElement
	}

	// Element stucture for basic strucutre nodes (html elements, etc)
	Element struct {
		Node
		Parent   IElement // if element is root, parent is nil
		tag      string
		xmlns    string
		children []IElement

		Attrs    func() Map
		RAttrs   Map                // last rendered Attrs
		Handlers map[string]Handler // events handlers: onClick, onHover

		HTML  func() string
		RHTML string

		Childes, OldChildes []interface{}

		RefName string

		Component *Component // can be nil

	}
)

var (
	ErrElementNotImplemented = errors.New("Browser not implemented Element")
)

var elementinterface js.Value
var customElements = make(map[string]IElement)

func init() {
	RegisterInterface(ElementInterface)
}

func DefineElement(tag string, element IElement) {
	customElements[tag] = element

}

func CreateElement(tag string) (IElement, error) {
	if element, ok := customElements[tag]; ok {
		return element, nil
	}

	return nil, ErrElementNotImplemented
}

// GetJSInterface get the JS interface
func ElementInterface() js.Value {
	sync.OnceFunc(func() {
		if elementinterface = js.Global().Get("Element"); elementinterface.IsUndefined() {
			elementinterface = js.Undefined()
		}

		Register(elementinterface, func(v js.Value) (interface{}, error) {
			return ToElement(v)
		})

	})

	return elementinterface
}
func ToElement(obj js.Value) (Element, error) {
	var e Element
	var err error
	if ei := ElementInterface(); !ei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(ei) {
				e.SetObject(obj)

			} else {
				err = ErrNotAnElement
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func (e Element) getAttributeElement(attribute string) (Element, error) {
	var nodeObject js.Value
	var newElement Element
	var err error

	if nodeObject, err = e.Get(attribute); err == nil {

		if nodeObject.IsUndefined() {
			err = ErrElementNoChilds

		} else {

			newElement, err = NewFromJSObject(nodeObject)

		}

	}

	return newElement, err
}

func (e Element) Attributes() (namednodemap.NamedNodeMap, error) {

	var err error
	var obj js.Value
	var namednmap namednodemap.NamedNodeMap

	if obj, err = e.Get("attributes"); err == nil {
		namednmap, err = namednodemap.NewFromJSObject(obj)
	}
	return namednmap, err
}

func (e Element) ChildElementCount() (int, error) {
	return e.GetAttributeInt("childElementCount")
}

func (e Element) Children() (HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection HtmlCollection

	if obj, err = e.Get("children"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (e Element) ClassList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj = e.Get("classList"); !obj.IsUndefined() {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (e Element) ClassName() (string, error) {
	return e.GetAttributeString("className")
}

func (e Element) SetClassName(value string) {
	e.SetAttributeString("className", value)
}

func (e Element) ClientHeight() (int, error) {
	return e.GetAttributeInt("clientHeight")
}

func (e Element) ClientLeft() (int, error) {
	return e.GetAttributeInt("clientLeft")
}

func (e Element) ClientTop() (int, error) {
	return e.GetAttributeInt("clientTop")
}

func (e Element) ClientWidth() (int, error) {
	return e.GetAttributeInt("clientWidth")
}

func (e Element) ComputedRole() (string, error) {
	return e.GetAttributeString("computedRole")
}

func (e Element) ID() (string, error) {
	return e.GetAttributeString("id")
}

func (e Element) SetID(value string) {
	e.SetAttributeString("id", value)
}

func (e Element) InnerHTML() (string, error) {
	return e.GetAttributeString("innerHTML")
}

func (e Element) SetInnerHTML(value string) {
	e.SetAttributeString("innerHTML", value)
}

func (e Element) LocalName() (string, error) {

	return e.GetAttributeString("localname")
}

func (e Element) NamespaceURI() (string, error) {

	return e.GetAttributeString("namespaceURI")
}

func (e Element) NextElementSibling() (Element, error) {
	return e.getAttributeElement("nextElementSibling")
}

func (e Element) OuterHTML() (string, error) {
	return e.GetAttributeString("outerHTML")
}

func (e Element) SetOuterHTML(value string) {
	e.SetAttributeString("outerHTML", value)
}

func (e Element) Prefix() (string, error) {

	return e.GetAttributeString("prefix")
}

func (e Element) PreviousElementSibling() (Element, error) {
	return e.getAttributeElement("previousElementSibling")
}

func (e Element) ScrollHeight() (int, error) {

	return e.GetAttributeInt("scrollHeight")
}

func (e Element) SetScrollHeight(value int) {

	e.SetAttributeInt("scrollHeight", value)

}

func (e Element) ScrollLeft() (int, error) {

	return e.GetAttributeInt("scrollLeft")
}

func (e Element) SetScrollLeft(value int) {

	e.SetAttributeInt("scrollLeft", value)

}

func (e Element) ScrollTop() (int, error) {

	return e.GetAttributeInt("scrollTop")
}

func (e Element) SetScrollTop(value int) {

	e.SetAttributeInt("scrollTop", value)

}

func (e Element) ScrollWidth() (int, error) {

	return e.GetAttributeInt("scrollWidth")
}

func (e Element) SetScrollWidth(value int) {

	e.SetAttributeInt("scrollWidth", value)

}

func (e Element) TagName() (string, error) {

	return e.GetAttributeString("tagName")
}
