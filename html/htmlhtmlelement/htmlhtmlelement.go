package htmlhtmlelement

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

var htmlhtmlelementinterface js.Value

// HtmlHtmlElement struct
type HtmlHtmlElement struct {
	htmlelement.HtmlElement
}

type HtmlHtmlElementFrom interface {
	HtmlHtmlElement_() HtmlHtmlElement
}

func (h HtmlHtmlElement) HtmlHtmlElement_() HtmlHtmlElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlhtmlelementinterface = js.Global().Get("HTMLHtmlElement"); htmlhtmlelementinterface.Error() != nil {
			htmlhtmlelementinterface = js.Undefined()
		}
		js.Register(htmlhtmlelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlhtmlelementinterface
}

func New(d document.Document) (HtmlHtmlElement, error) {
	var err error

	var h HtmlHtmlElement
	var e element.Element

	if e, err = d.CreateElement("html"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHtmlElement, error) {
	var h HtmlHtmlElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlHtmlElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHtmlElement, error) {
	var h HtmlHtmlElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlHtmlElement
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return h, err
}
