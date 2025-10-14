package htmltimeelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltimeelementinterface js.Value

// HtmlTimeElement struct
type HtmlTimeElement struct {
	htmlelement.HtmlElement
}

type HtmlTimeElementFrom interface {
	HtmlTimeElement_() HtmlTimeElement
}

func (h HtmlTimeElement) HtmlTimeElement_() HtmlTimeElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltimeelementinterface = js.Global().Get("HTMLTimeElement"); htmltimeelementinterface.Error() != nil {
			htmltimeelementinterface = js.Undefined()
		}
		js.Register(htmltimeelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltimeelementinterface
}

func New(d document.Document) (HtmlTimeElement, error) {
	var err error

	var h HtmlTimeElement
	var e element.Element

	if e, err = d.CreateElement("time"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTimeElement, error) {
	var h HtmlTimeElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTimeElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTimeElement, error) {
	var h HtmlTimeElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTimeElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTimeElement) DateTime() (string, error) {
	return h.GetAttributeString("dateTime")
}

func (h HtmlTimeElement) SetDateTime(value string) error {
	return h.SetAttributeString("dateTime", value)
}
