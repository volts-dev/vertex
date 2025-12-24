package htmllegendelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/htmlformelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmllegendelementinterface js.Value

// HtmlLegendElement struct
type HtmlLegendElement struct {
	htmlelement.HtmlElement
}

type HtmlLegendElementFrom interface {
	HtmlLegendElement_() HtmlLegendElement
}

func (h HtmlLegendElement) HtmlLegendElement_() HtmlLegendElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmllegendelementinterface = js.Global().Get("HTMLLegendElement"); htmllegendelementinterface.Error() != nil {
			htmllegendelementinterface = js.Undefined()
		}
		js.Register(htmllegendelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllegendelementinterface
}

func New(d document.Document) (HtmlLegendElement, error) {
	var err error

	var h HtmlLegendElement
	var e element.Element

	if e, err = d.CreateElement("legend"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLegendElement, error) {
	var h HtmlLegendElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLLegendElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLegendElement, error) {
	var h HtmlLegendElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLLegendElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLegendElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var formelem htmlformelement.HtmlFormElement

	if obj = h.GetValueByKey("form"); obj.Error() == nil {

		formelem, err = htmlformelement.NewFromJSObject(obj)
	}

	return formelem, err
}

func (h HtmlLegendElement) AccessKey() (string, error) {
	return h.GetAttributeString("accessKey")
}

func (h HtmlLegendElement) SetAccessKey(value string) error {
	return h.SetAttributeString("accessKey", value)
}
