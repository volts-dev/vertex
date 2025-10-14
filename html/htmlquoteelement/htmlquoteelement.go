package htmlquoteelement

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

var htmlquoteelementinterface js.Value

// HtmlQuoteElement struct
type HtmlQuoteElement struct {
	htmlelement.HtmlElement
}

type HtmlQuoteElementFrom interface {
	HtmlQuoteElement_() HtmlQuoteElement
}

func (h HtmlQuoteElement) HtmlQuoteElement_() HtmlQuoteElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlquoteelementinterface = js.Global().Get("HTMLQuoteElement"); htmlquoteelementinterface.Error() != nil {
			htmlquoteelementinterface = js.Undefined()
		}
		js.Register(htmlquoteelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlquoteelementinterface
}

func New(d document.Document) (HtmlQuoteElement, error) {
	var err error

	var h HtmlQuoteElement
	var e element.Element

	if e, err = d.CreateElement("q"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewBlockQuote(d document.Document) (HtmlQuoteElement, error) {
	var err error

	var h HtmlQuoteElement
	var e element.Element

	if e, err = d.CreateElement("blockquote"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlQuoteElement, error) {
	var h HtmlQuoteElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLQuoteElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlQuoteElement, error) {
	var h HtmlQuoteElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLQuoteElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlQuoteElement) Cite() (string, error) {

	return h.GetAttributeString("cite")
}

func (h HtmlQuoteElement) SetCite(value string) error {
	return h.SetAttributeString("cite", value)
}
