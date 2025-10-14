package htmlmapelement

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

var htmlmapelementinterface js.Value

// HtmlMapElement struct
type HtmlMapElement struct {
	htmlelement.HtmlElement
}

type HtmlMapElementFrom interface {
	HtmlMapElement_() HtmlMapElement
}

func (h HtmlMapElement) HtmlMapElement_() HtmlMapElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlmapelementinterface = js.Global().Get("HTMLMapElement"); htmlmapelementinterface.Error() != nil {
			htmlmapelementinterface = js.Undefined()
		}
		js.Register(htmlmapelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlmapelementinterface
}

func New(d document.Document) (HtmlMapElement, error) {
	var err error

	var h HtmlMapElement
	var e element.Element

	if e, err = d.CreateElement("map"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMapElement, error) {
	var h HtmlMapElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLMapElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMapElement, error) {
	var h HtmlMapElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLMapElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlMapElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlMapElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlMapElement) Areas() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("areas"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
