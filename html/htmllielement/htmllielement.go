package htmllielement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmllielementinterface js.Value

// HtmlLIElement struct
type HtmlLIElement struct {
	htmlelement.HtmlElement
}

type HtmlLIElementFrom interface {
	HtmlLIElement_() HtmlLIElement
}

func (h HtmlLIElement) HtmlLIElement_() HtmlLIElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmllielementinterface = js.Global().Get("HTMLLIElement"); htmllielementinterface.Error() != nil {
			htmllielementinterface = js.Undefined()
		}
		js.Register(htmllielementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllielementinterface
}

func New(d document.Document) (HtmlLIElement, error) {
	var err error

	var h HtmlLIElement
	var e element.Element

	if e, err = d.CreateElement("li"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLIElement, error) {
	var h HtmlLIElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLLIElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLIElement, error) {
	var h HtmlLIElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLLIElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLIElement) Value() (int, error) {
	return h.GetAttributeInt("value")
}

func (h HtmlLIElement) SetValue(value int) error {
	return h.SetAttributeInt("value", value)
}
