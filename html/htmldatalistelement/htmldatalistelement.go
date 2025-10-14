package htmldatalistelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmldatalistelementinterface js.Value

// HtmlDataElement struct
type HtmlDataListElement struct {
	htmlelement.HtmlElement
}

type HtmlDataListElementFrom interface {
	HtmlDataListElement_() HtmlDataListElement
}

func (h HtmlDataListElement) HtmlDataListElement_() HtmlDataListElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmldatalistelementinterface = js.Global().Get("HTMLDataListElement"); htmldatalistelementinterface.Error() != nil {
			htmldatalistelementinterface = js.Undefined()
		}
		js.Register(htmldatalistelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldatalistelementinterface
}

func New(d document.Document) (HtmlDataListElement, error) {
	var err error

	var h HtmlDataListElement
	var e element.Element

	if e, err = d.CreateElement("datalist"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDataListElement, error) {
	var h HtmlDataListElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlDataListElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDataListElement, error) {
	var h HtmlDataListElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlDataListElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlDataListElement) Options() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("options"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
