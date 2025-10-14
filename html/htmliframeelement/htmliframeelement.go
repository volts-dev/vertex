package htmliframeelement

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

var htmliframelementinterface js.Value

// HtmlHeadingElement struct
type HtmlIFrameElement struct {
	htmlelement.HtmlElement
}

type HtmlIFrameElementFrom interface {
	HtmlIFrameElement_() HtmlIFrameElement
}

func (h HtmlIFrameElement) HtmlIFrameElement_() HtmlIFrameElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmliframelementinterface = js.Global().Get("HTMLIFrameElement"); htmliframelementinterface.Error() != nil {
			htmliframelementinterface = js.Undefined()
		}

		js.Register(htmliframelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmliframelementinterface
}

func New(d document.Document) (HtmlIFrameElement, error) {
	var err error

	var h HtmlIFrameElement
	var e element.Element

	if e, err = d.CreateElement("iframe"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlIFrameElement, error) {
	var h HtmlIFrameElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlIFrameElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlIFrameElement, error) {
	var h HtmlIFrameElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlIFrameElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlIFrameElement) AllowPaymentRequest() (bool, error) {
	return h.GetAttributeBool("allowPaymentRequest")
}

func (h HtmlIFrameElement) SetAllowPaymentRequest(value bool) error {
	return h.SetAttributeBool("allowPaymentRequest", value)
}

func (h HtmlIFrameElement) ContentDocument() (document.Document, error) {
	var err error
	var obj js.Value
	var doc document.Document

	if obj = h.GetValueByKey("contentDocument"); obj.Error() == nil {
		if !obj.IsNull() {
			doc, err = document.NewFromJSObject(obj)
		} else {
			err = ErrNoContentDocument
		}

	}

	return doc, err
}

func (h HtmlIFrameElement) Height() (string, error) {
	return h.GetAttributeString("height")
}

func (h HtmlIFrameElement) SetHeight(value string) error {
	return h.SetAttributeString("height", value)
}

func (h HtmlIFrameElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlIFrameElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlIFrameElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlIFrameElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlIFrameElement) Srcdoc() (string, error) {
	return h.GetAttributeString("srcdoc")
}

func (h HtmlIFrameElement) SetSrcdoc(value string) error {
	return h.SetAttributeString("srcdoc", value)
}

func (h HtmlIFrameElement) Width() (string, error) {
	return h.GetAttributeString("width")
}

func (h HtmlIFrameElement) SetWidth(value string) error {
	return h.SetAttributeString("width", value)
}
