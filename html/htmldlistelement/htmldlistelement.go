package htmldlistelement

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

var htmldlistinterface js.Value

// HtmlDListElement struct
type HtmlDListElement struct {
	htmlelement.HtmlElement
}

type HtmlDListElementFrom interface {
	HtmlDListElement_() HtmlDListElement
}

func (h HtmlDListElement) HtmlDivElement_() HtmlDListElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmldlistinterface = js.Global().Get("HTMLDListElement"); htmldlistinterface.Error() != nil {
			htmldlistinterface = js.Undefined()
		}
		js.Register(htmldlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldlistinterface
}

func New(d document.Document) (HtmlDListElement, error) {
	var err error

	var h HtmlDListElement
	var e element.Element

	if e, err = d.CreateElement("dl"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDListElement, error) {
	var h HtmlDListElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlDListElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDListElement, error) {
	var h HtmlDListElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlDListElement
			}
		}
	} else {
		err = ErrNotImplemented

	}
	return h, err
}
