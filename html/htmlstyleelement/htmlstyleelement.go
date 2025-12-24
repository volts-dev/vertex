package htmlstyleelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/stylesheet"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlstylelementinterface js.Value

// HtmlStyleElement struct
type HtmlStyleElement struct {
	htmlelement.HtmlElement
}

type HtmlStyleElementFrom interface {
	HtmlStyleElement_() HtmlStyleElement
}

func (h HtmlStyleElement) HtmlStyleElement_() HtmlStyleElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlstylelementinterface = js.Global().Get("HTMLStyleElement"); htmlstylelementinterface.Error() != nil {
			htmlstylelementinterface = js.Undefined()
		}
		js.Register(htmlstylelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlstylelementinterface
}

func New(d document.Document) (HtmlStyleElement, error) {
	var err error

	var h HtmlStyleElement
	var e element.Element

	if e, err = d.CreateElement("style"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlStyleElement, error) {
	var h HtmlStyleElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLStyleElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlStyleElement, error) {
	var h HtmlStyleElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLStyleElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlStyleElement) Media() (string, error) {
	return h.GetAttributeString("media")
}

func (h HtmlStyleElement) SetMedia(value string) error {
	return h.SetAttributeString("media", value)
}

func (h HtmlStyleElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlStyleElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlStyleElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlStyleElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlStyleElement) Sheet() (stylesheet.StyleSheet, error) {
	var err error
	var obj js.Value
	var s stylesheet.StyleSheet
	if obj = h.GetValueByKey("sheet"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			s, err = stylesheet.NewFromJSObject(obj)
		}
	}
	return s, err
}
