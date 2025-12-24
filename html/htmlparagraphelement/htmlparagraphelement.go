package htmlparagraphelement

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

var htmlparagraphelementinterface js.Value

// HtmlParagraphElement struct
type HtmlParagraphElement struct {
	htmlelement.HtmlElement
}

type HtmlParagraphElementFrom interface {
	HtmlParagraphElement_() HtmlParagraphElement
}

func (h HtmlParagraphElement) HtmlParagraphElement_() HtmlParagraphElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlparagraphelementinterface = js.Global().Get("HTMLParagraphElement"); htmlparagraphelementinterface.Error() != nil {
			htmlparagraphelementinterface = js.Undefined()
		}
		js.Register(htmlparagraphelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlparagraphelementinterface
}

func New(d document.Document) (HtmlParagraphElement, error) {
	var err error

	var h HtmlParagraphElement
	var e element.Element

	if e, err = d.CreateElement("p"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLParagraphElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLParagraphElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
