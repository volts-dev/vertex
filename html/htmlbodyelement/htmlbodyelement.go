package htmlbodyelement

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

var htmlbodyelementinterface js.Value

// HtmlBodyElement struct
type HtmlBodyElement struct {
	htmlelement.HtmlElement
}

type HtmlBodyElementFrom interface {
	HtmlBodyElement_() HtmlBodyElement
}

func (h HtmlBodyElement) HtmlBodyElement_() HtmlBodyElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlbodyelementinterface = js.Global().Get("HTMLBodyElement"); htmlbodyelementinterface.Error() != nil {
			htmlbodyelementinterface = js.Undefined()
		}
		js.Register(htmlbodyelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlbodyelementinterface
}

func New(d document.Document) (HtmlBodyElement, error) {
	var err error

	var h HtmlBodyElement
	var e element.Element

	if e, err = d.CreateElement("body"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlBodyElement, error) {
	var h HtmlBodyElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlBodyElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlBodyElement, error) {
	var h HtmlBodyElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlBodyElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
