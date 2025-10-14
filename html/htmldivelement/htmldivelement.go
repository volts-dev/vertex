package htmldivelement

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

var htmldivelementinterface js.Value

// HtmlDetailsElement struct
type HtmlDivElement struct {
	htmlelement.HtmlElement
}

type HtmlDivElementFrom interface {
	HtmlDivElement_() HtmlDivElement
}

func (h HtmlDivElement) HtmlDivElement_() HtmlDivElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmldivelementinterface = js.Global().Get("HTMLDivElement"); htmldivelementinterface.Error() != nil {
			htmldivelementinterface = js.Undefined()
		}
		js.Register(htmldivelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldivelementinterface
}

func New(d document.Document) (HtmlDivElement, error) {
	var err error

	var h HtmlDivElement
	var e element.Element

	if e, err = d.CreateElement("div"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDivElement, error) {
	var h HtmlDivElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlDivElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDivElement, error) {
	var h HtmlDivElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlDivElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
