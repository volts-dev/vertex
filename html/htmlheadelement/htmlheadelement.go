package htmlheadelement

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

var htmlheadelementinterface js.Value

// HtmlHeadElement struct
type HtmlHeadElement struct {
	htmlelement.HtmlElement
}

type HtmlHeadElementFrom interface {
	HtmlHeadElement_() HtmlHeadElement
}

func (h HtmlHeadElement) HtmlHeadElement_() HtmlHeadElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlheadelementinterface = js.Global().Get("HTMLHeadElement"); htmlheadelementinterface.Error() != nil {
			htmlheadelementinterface = js.Undefined()
		}
		js.Register(htmlheadelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlheadelementinterface
}

func New(d document.Document) (HtmlHeadElement, error) {
	var err error

	var h HtmlHeadElement
	var e element.Element

	if e, err = d.CreateElement("head"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHeadElement, error) {
	var h HtmlHeadElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlHeadElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHeadElement, error) {
	var h HtmlHeadElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlHeadElement
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return h, err
}
