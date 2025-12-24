package htmlheadingelement

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

var htmlheadingelementinterface js.Value

// HtmlHeadingElement struct
type HtmlHeadingElement struct {
	htmlelement.HtmlElement
}

type HtmlHeadingElementFrom interface {
	HtmlHeadingElement_() HtmlHeadingElement
}

func (h HtmlHeadingElement) HtmlHeadingElement_() HtmlHeadingElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlheadingelementinterface = js.Global().Get("HTMLHeadingElement"); htmlheadingelementinterface.Error() != nil {
			htmlheadingelementinterface = js.Undefined()
		}
		js.Register(htmlheadingelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlheadingelementinterface
}

func NewH1(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h1"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH2(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h2"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH3(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h3"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH4(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h4"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH5(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h5"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH6(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h6"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHeadingElement, error) {
	var h HtmlHeadingElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlHeadingElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHeadingElement, error) {
	var h HtmlHeadingElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlHeadingElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
