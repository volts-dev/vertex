package htmlpreelement

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

var htmlpreelementinterface js.Value

// HtmlPreElement struct
type HtmlPreElement struct {
	htmlelement.HtmlElement
}

type HtmlPreElementFrom interface {
	HtmlPreElement_() HtmlPreElement
}

func (h HtmlPreElement) HtmlPreElement_() HtmlPreElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlpreelementinterface = js.Global().Get("HTMLPreElement"); htmlpreelementinterface.Error() != nil {
			htmlpreelementinterface = js.Undefined()
		}

		js.Register(htmlpreelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlpreelementinterface
}

func New(d document.Document) (HtmlPreElement, error) {
	var err error

	var h HtmlPreElement
	var e element.Element

	if e, err = d.CreateElement("pre"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlPreElement, error) {
	var h HtmlPreElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLPreElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlPreElement, error) {
	var h HtmlPreElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLPreElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
