package htmlprogresselement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/nodelist"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlprogresselementinterface js.Value

// HtmlProgressElement struct
type HtmlProgressElement struct {
	htmlelement.HtmlElement
}

type HtmlProgressElementFrom interface {
	HtmlProgressElement_() HtmlProgressElement
}

func (h HtmlProgressElement) HtmlProgressElement_() HtmlProgressElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlprogresselementinterface = js.Global().Get("HTMLProgressElement"); htmlprogresselementinterface.Error() != nil {
			htmlprogresselementinterface = js.Undefined()
		}
		js.Register(htmlprogresselementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlprogresselementinterface
}

func New(d document.Document) (HtmlProgressElement, error) {
	var err error

	var h HtmlProgressElement
	var e element.Element

	if e, err = d.CreateElement("progress"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlProgressElement, error) {
	var h HtmlProgressElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlProgressElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlProgressElement, error) {
	var h HtmlProgressElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlProgressElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlProgressElement) Max() (float64, error) {
	return h.GetAttributeDouble("max")
}

func (h HtmlProgressElement) SetMax(value float64) error {
	return h.SetAttributeDouble("max", value)
}

func (h HtmlProgressElement) Position() (float64, error) {
	return h.GetAttributeDouble("position")
}

func (h HtmlProgressElement) Value() (float64, error) {
	return h.GetAttributeDouble("value")
}

func (h HtmlProgressElement) SetValue(value float64) error {
	return h.SetAttributeDouble("value", value)
}

func (h HtmlProgressElement) Labels() (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj = h.GetValueByKey("labels"); obj.Error() == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}

	return nlist, err
}
