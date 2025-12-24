package htmldetailselement

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

var htmldetailselementinterface js.Value

// HtmlDetailsElement struct
type HtmlDetailsElement struct {
	htmlelement.HtmlElement
}

type HtmlDetailsElementFrom interface {
	HtmlDetailsElement_() HtmlDetailsElement
}

func (h HtmlDetailsElement) HtmlDetailsElement_() HtmlDetailsElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmldetailselementinterface = js.Global().Get("HTMLDetailsElement"); htmldetailselementinterface.Error() != nil {
			htmldetailselementinterface = js.Undefined()
		}

		js.Register(htmldetailselementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldetailselementinterface
}

func New(d document.Document) (HtmlDetailsElement, error) {
	var err error

	var h HtmlDetailsElement
	var e element.Element

	if e, err = d.CreateElement("details"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDetailsElement, error) {
	var h HtmlDetailsElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlDetailsElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDetailsElement, error) {
	var h HtmlDetailsElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlDetailsElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlDetailsElement) Open() (bool, error) {
	return h.GetAttributeBool("open")
}

func (h HtmlDetailsElement) SetOpen(value bool) error {
	return h.SetAttributeBool("open", value)
}
