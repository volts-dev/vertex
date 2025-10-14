package htmltitleelement

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

var htmltitleelementinterface js.Value

// HtmlTemplatelement struct
type HtmlTitleElement struct {
	htmlelement.HtmlElement
}

type HtmlTitleElementFrom interface {
	HtmlTitleElement_() HtmlTitleElement
}

func (h HtmlTitleElement) HtmlTitleElement_() HtmlTitleElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltitleelementinterface = js.Global().Get("HTMLTitleElement"); htmltitleelementinterface.Error() != nil {
			htmltitleelementinterface = js.Undefined()
		}
		js.Register(htmltitleelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltitleelementinterface
}

func New(d document.Document) (HtmlTitleElement, error) {
	var err error

	var h HtmlTitleElement
	var e element.Element

	if e, err = d.CreateElement("title"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTitleElement, error) {
	var h HtmlTitleElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTitleElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTitleElement, error) {
	var h HtmlTitleElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTitleElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTitleElement) Text() (string, error) {
	return h.GetAttributeString("text")
}

func (h HtmlTitleElement) SetText(value string) error {
	return h.SetAttributeString("text", value)
}
