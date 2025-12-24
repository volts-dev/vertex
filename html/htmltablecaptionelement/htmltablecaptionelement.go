package htmltablecaptionelement

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

var htmltablecaptionelementinterface js.Value

// HtmlTableCaptionElement struct
type HtmlTableCaptionElement struct {
	htmlelement.HtmlElement
}

type HtmlTableCaptionElementFrom interface {
	HtmlTableCaptionElement_() HtmlTableCaptionElement
}

func (h HtmlTableCaptionElement) HtmlTableCaptionElement_() HtmlTableCaptionElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltablecaptionelementinterface = js.Global().Get("HTMLTableCaptionElement"); htmltablecaptionelementinterface.Error() != nil {
			htmltablecaptionelementinterface = js.Undefined()
		}

		js.Register(htmltablecaptionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecaptionelementinterface
}

func New(d document.Document) (HtmlTableCaptionElement, error) {
	var err error

	var h HtmlTableCaptionElement
	var e element.Element

	if e, err = d.CreateElement("caption"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableCaptionElement, error) {
	var h HtmlTableCaptionElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableCaptionElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableCaptionElement, error) {
	var h HtmlTableCaptionElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableCaptionElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
