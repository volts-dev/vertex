package htmllabelelement

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

var htmllabelelementinterface js.Value

// HtmlLabelElement struct
type HtmlLabelElement struct {
	htmlelement.HtmlElement
}

type HtmlLabelElementFrom interface {
	HtmlLabelElement_() HtmlLabelElement
}

func (h HtmlLabelElement) HtmlLabelElement_() HtmlLabelElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmllabelelementinterface = js.Global().Get("HTMLLabelElement"); htmllabelelementinterface.Error() != nil {
			htmllabelelementinterface = js.Undefined()
		}
		js.Register(htmllabelelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllabelelementinterface
}

func New(d document.Document) (HtmlLabelElement, error) {
	var err error

	var h HtmlLabelElement
	var e element.Element

	if e, err = d.CreateElement("label"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLabelElement, error) {
	var h HtmlLabelElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLLabelElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLabelElement, error) {
	var h HtmlLabelElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLLabelElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLabelElement) Control() (htmlelement.HtmlElement, error) {
	var err error
	var obj js.Value
	var htmlelem htmlelement.HtmlElement

	if obj = h.GetValueByKey("control"); obj.Error() == nil {

		htmlelem, err = htmlelement.NewFromJSObject(obj)
	}

	return htmlelem, err
}

func (h HtmlLabelElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var formelem htmlformelement.HtmlFormElement

	if obj = h.GetValueByKey("form"); obj.Error() == nil {

		formelem, err = htmlformelement.NewFromJSObject(obj)
	}

	return formelem, err
}

func (h HtmlLabelElement) HtmlFor() (string, error) {
	return h.GetAttributeString("htmlFor")
}

func (h HtmlLabelElement) SetHtmlFor(value string) error {
	return h.SetAttributeString("htmlFor", value)
}
