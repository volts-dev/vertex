package htmlscriptelement

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

var htmlscriptelementinterface js.Value

// HtmlScriptElement struct
type HtmlScriptElement struct {
	htmlelement.HtmlElement
}

type HtmlScriptElementFrom interface {
	HtmlScriptElement_() HtmlScriptElement
}

func (h HtmlScriptElement) HtmlScriptElement_() HtmlScriptElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlscriptelementinterface = js.Global().Get("HTMLScriptElement"); htmlscriptelementinterface.Error() != nil {
			htmlscriptelementinterface = js.Undefined()
		}

		js.Register(htmlscriptelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlscriptelementinterface
}

func New(d document.Document) (HtmlScriptElement, error) {
	var err error

	var h HtmlScriptElement
	var e element.Element

	if e, err = d.CreateElement("script"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlScriptElement, error) {
	var h HtmlScriptElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLScriptElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlScriptElement, error) {
	var h HtmlScriptElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLScriptElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlScriptElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlScriptElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlScriptElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlScriptElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlScriptElement) Async() (bool, error) {
	return h.GetAttributeBool("async")
}

func (h HtmlScriptElement) SetAsync(value bool) error {
	return h.SetAttributeBool("async", value)
}

func (h HtmlScriptElement) Defer() (bool, error) {
	return h.GetAttributeBool("defer")
}

func (h HtmlScriptElement) SetDefer(value bool) error {
	return h.SetAttributeBool("defer", value)
}

func (h HtmlScriptElement) Text() (string, error) {

	return h.GetAttributeString("text")
}

func (h HtmlScriptElement) SetText(value string) error {
	return h.SetAttributeString("text", value)
}

func (h HtmlScriptElement) NoModule() (bool, error) {
	return h.GetAttributeBool("noModule")
}

func (h HtmlScriptElement) SetNoModule(value bool) error {
	return h.SetAttributeBool("noModule", value)
}

func (h HtmlScriptElement) ReferrerPolicy() (string, error) {
	return h.GetAttributeString("referrerPolicy")
}

func (h HtmlScriptElement) SetReferrerPolicy(value string) error {
	return h.SetAttributeString("referrerPolicy", value)
}
