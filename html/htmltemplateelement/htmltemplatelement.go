package htmltemplateelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/documentfragment"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once
var htmltemplateelementinterface js.Value

// HTMLTemplateElement struct
type HTMLTemplateElement struct {
	htmlelement.HtmlElement
}

type HTMLTemplateElementFrom interface {
	HTMLTemplateElement_() HTMLTemplateElement
}

func (h HTMLTemplateElement) HTMLTemplateElement_() HTMLTemplateElement {
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

func New(d document.Document) (HTMLTemplateElement, error) {
	var err error
	var h HTMLTemplateElement
	var e element.Element

	if e, err = d.CreateElement("template"); err == nil {
		h, err = NewFromElement(e)
	}

	return h, err
}

func NewFromElement(elem element.Element) (HTMLTemplateElement, error) {
	var h HTMLTemplateElement
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

func NewFromJSObject(obj js.Value) (HTMLTemplateElement, error) {
	var h HTMLTemplateElement
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

func (h HTMLTemplateElement) Content() (documentfragment.DocumentFragment, error) {
	var err error
	var fragment documentfragment.DocumentFragment

	if obj := h.GetValueByKey("content"); obj.Error() == nil {
		fragment, err = documentfragment.NewFromJSObject(obj)
	}

	return fragment, err
}
