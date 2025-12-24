package htmltablecolelement

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

var htmltablecolelementinterface js.Value

// HtmlTableColElement struct
type HtmlTableColElement struct {
	htmlelement.HtmlElement
}

type HtmlTableColElementFrom interface {
	HtmlTableColElement_() HtmlTableColElement
}

func (h HtmlTableColElement) HtmlTableColElement_() HtmlTableColElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltablecolelementinterface = js.Global().Get("HTMLTableColElement"); htmltablecolelementinterface.Error() != nil {
			htmltablecolelementinterface = js.Undefined()
		}
		js.Register(htmltablecolelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecolelementinterface
}

func New(d document.Document) (HtmlTableColElement, error) {
	var err error

	var h HtmlTableColElement
	var e element.Element

	if e, err = d.CreateElement("col"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableColElement, error) {
	var h HtmlTableColElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableColElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableColElement, error) {
	var h HtmlTableColElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableColElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableColElement) Span() (int, error) {
	return h.GetAttributeInt("span")
}

func (h HtmlTableColElement) SetSpan(value int) error {
	return h.SetAttributeInt("span", value)
}
