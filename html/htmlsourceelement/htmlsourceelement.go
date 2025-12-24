package htmlsourceelement

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

var htmlsourceelementinterface js.Value

// HtmlSourceElement struct
type HtmlSourceElement struct {
	htmlelement.HtmlElement
}

type HtmlSourceElementFrom interface {
	HtmlSourceElement_() HtmlSourceElement
}

func (h HtmlSourceElement) HtmlSourceElement_() HtmlSourceElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlsourceelementinterface = js.Global().Get("HTMLSourceElement"); htmlsourceelementinterface.Error() != nil {
			htmlsourceelementinterface = js.Undefined()
		}
		js.Register(htmlsourceelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlsourceelementinterface
}

func New(d document.Document) (HtmlSourceElement, error) {
	var err error

	var h HtmlSourceElement
	var e element.Element

	if e, err = d.CreateElement("source"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSourceElement, error) {
	var h HtmlSourceElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLSourceElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSourceElement, error) {
	var h HtmlSourceElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLSourceElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlSourceElement) Media() (string, error) {
	return h.GetAttributeString("media")
}

func (h HtmlSourceElement) SetMedia(value string) error {
	return h.SetAttributeString("media", value)
}

func (h HtmlSourceElement) Sizes() (string, error) {
	return h.GetAttributeString("sizes")
}

func (h HtmlSourceElement) SetSizes(value string) error {
	return h.SetAttributeString("sizes", value)
}

func (h HtmlSourceElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlSourceElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlSourceElement) SrcSet() (string, error) {
	return h.GetAttributeString("srcset")
}

func (h HtmlSourceElement) SetSrcSet(value string) error {
	return h.SetAttributeString("srcset", value)
}

func (h HtmlSourceElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlSourceElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}
