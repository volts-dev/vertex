package htmldataelement

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

var htmldataelementinterface js.Value

// HtmlDataElement struct
type HtmlDataElement struct {
	htmlelement.HtmlElement
}

type HtmlDataElementFrom interface {
	HtmlDataElement_() HtmlDataElement
}

func (h HtmlDataElement) HtmlDataElement_() HtmlDataElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmldataelementinterface = js.Global().Get("HTMLDataElement"); htmldataelementinterface.Error() != nil {
			htmldataelementinterface = js.Undefined()
		}
		js.Register(htmldataelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldataelementinterface
}

func New(d document.Document) (HtmlDataElement, error) {
	var err error

	var h HtmlDataElement
	var e element.Element

	if e, err = d.CreateElement("data"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDataElement, error) {
	var h HtmlDataElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {

		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlDataElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDataElement, error) {
	var h HtmlDataElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlDataElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlDataElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlDataElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}
