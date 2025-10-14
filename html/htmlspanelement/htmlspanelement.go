package htmlspanelement

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

var htmlspanelementinterface js.Value

// HtmlSpanElement struct
type HtmlSpanElement struct {
	htmlelement.HtmlElement
}

type HtmlSpanElementFrom interface {
	HtmlSpanElement_() HtmlSpanElement
}

func (h HtmlSpanElement) HtmlSpanElement_() HtmlSpanElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlspanelementinterface = js.Global().Get("HTMLSpanElement"); htmlspanelementinterface.Error() != nil {
			htmlspanelementinterface = js.Undefined()
		}
		js.Register(htmlspanelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlspanelementinterface
}

func New(d document.Document) (HtmlSpanElement, error) {
	var err error

	var h HtmlSpanElement
	var e element.Element

	if e, err = d.CreateElement("span"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSpanElement, error) {
	var h HtmlSpanElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLSpanElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSpanElement, error) {
	var h HtmlSpanElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLSpanElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
