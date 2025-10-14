package htmltemplateelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/documentfragment"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltemplateelementinterface js.Value

// HtmlTemplateElement struct
type HtmlTemplateElement struct {
	htmlelement.HtmlElement
}

type HtmlTemplateElementFrom interface {
	HtmlTemplateElement_() HtmlTemplateElement
}

func (h HtmlTemplateElement) HtmlTemplateElement_() HtmlTemplateElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltemplateelementinterface = js.Global().Get("HTMLTemplateElement"); htmltemplateelementinterface.Error() != nil {
			htmltemplateelementinterface = js.Undefined()
		}
		js.Register(htmltemplateelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltemplateelementinterface
}

func New(d document.Document) (HtmlTemplateElement, error) {
	var err error

	var h HtmlTemplateElement
	var e element.Element

	if e, err = d.CreateElement("template"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTemplateElement, error) {
	var h HtmlTemplateElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTemplateElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTemplateElement, error) {
	var h HtmlTemplateElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTemplateElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTemplateElement) Content() (documentfragment.DocumentFragment, error) {
	var err error
	var obj js.Value
	var fragment documentfragment.DocumentFragment

	if obj = h.GetValueByKey("content"); obj.Error() == nil {

		fragment, err = documentfragment.NewFromJSObject(obj)
	}

	return fragment, err
}
