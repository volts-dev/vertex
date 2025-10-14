package htmlbrelement

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

var htmlbrelementinterface js.Value

// HtmlBRElement struct
type HtmlBRElement struct {
	htmlelement.HtmlElement
}

type HtmlBRElementFrom interface {
	HtmlBRElement_() HtmlBRElement
}

func (h HtmlBRElement) HtmlBRElement_() HtmlBRElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlbrelementinterface = js.Global().Get("HTMLBRElement"); htmlbrelementinterface.Error() != nil {
			htmlbrelementinterface = js.Undefined()
		}
		js.Register(htmlbrelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlbrelementinterface
}

func New(d document.Document) (HtmlBRElement, error) {
	var err error

	var h HtmlBRElement
	var e element.Element

	if e, err = d.CreateElement("br"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlBRElement, error) {
	var h HtmlBRElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlBrElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlBRElement, error) {
	var h HtmlBRElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {

		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlBrElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
