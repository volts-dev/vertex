package htmlhrelement

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

var htmlhrelementinterface js.Value

// HtmlHrElement struct
type HtmlHrElement struct {
	htmlelement.HtmlElement
}

type HtmlHrElementFrom interface {
	HtmlHrElement_() HtmlHrElement
}

func (h HtmlHrElement) HtmlHrElement_() HtmlHrElement {
	return h
}

func GetInterface() js.Value {
	singleton.Do(func() {
		if htmlhrelementinterface = js.Global().Get("HTMLHRElement"); htmlhrelementinterface.Error() != nil {
			htmlhrelementinterface = js.Undefined()
		}
		js.Register(htmlhrelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlhrelementinterface
}

func New(d document.Document) (HtmlHrElement, error) {
	var err error

	var h HtmlHrElement
	var e element.Element

	if e, err = d.CreateElement("hr"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHrElement, error) {
	var h HtmlHrElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlHrElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHrElement, error) {
	var h HtmlHrElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlHrElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
