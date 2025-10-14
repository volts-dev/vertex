package htmlmetaelement

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

var htmlmetaelementinterface js.Value

// HtmlMetaElement struct
type HtmlMetaElement struct {
	htmlelement.HtmlElement
}

type HtmlMetaElementFrom interface {
	HtmlMetaElement_() HtmlMetaElement
}

func (h HtmlMetaElement) HtmlMetaElement_() HtmlMetaElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlmetaelementinterface = js.Global().Get("HTMLMetaElement"); htmlmetaelementinterface.Error() != nil {
			htmlmetaelementinterface = js.Undefined()
		}
		js.Register(htmlmetaelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlmetaelementinterface
}

func New(d document.Document) (HtmlMetaElement, error) {
	var err error

	var h HtmlMetaElement
	var e element.Element

	if e, err = d.CreateElement("meta"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMetaElement, error) {
	var h HtmlMetaElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLMetaElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMetaElement, error) {
	var h HtmlMetaElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLMetaElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlMetaElement) Content() (string, error) {
	return h.GetAttributeString("content")
}

func (h HtmlMetaElement) SetContent(value string) error {
	return h.SetAttributeString("content", value)
}

func (h HtmlMetaElement) HttpEquiv() (string, error) {
	return h.GetAttributeString("httpEquiv")
}

func (h HtmlMetaElement) SetHttpEquiv(value string) error {
	return h.SetAttributeString("httpEquiv", value)
}

func (h HtmlMetaElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlMetaElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}
