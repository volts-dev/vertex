package htmlmeterelement

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

var htmlmeterelementinterface js.Value

// HtmlMeterElement struct
type HtmlMeterElement struct {
	htmlelement.HtmlElement
}

type HtmlMeterElementFrom interface {
	HtmlMeterElement_() HtmlMeterElement
}

func (h HtmlMeterElement) HtmlMeterElement_() HtmlMeterElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlmeterelementinterface = js.Global().Get("HTMLMeterElement"); htmlmeterelementinterface.Error() != nil {
			htmlmeterelementinterface = js.Undefined()
		}
		js.Register(htmlmeterelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlmeterelementinterface
}

func New(d document.Document) (HtmlMeterElement, error) {
	var err error

	var h HtmlMeterElement
	var e element.Element

	if e, err = d.CreateElement("meter"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMeterElement, error) {
	var h HtmlMeterElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLMeterElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMeterElement, error) {
	var h HtmlMeterElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLMeterElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlMeterElement) High() (float64, error) {
	return h.GetAttributeDouble("high")
}

func (h HtmlMeterElement) SetHigh(value float64) error {
	return h.SetAttributeDouble("high", value)
}

func (h HtmlMeterElement) Low() (float64, error) {
	return h.GetAttributeDouble("low")
}

func (h HtmlMeterElement) SetLow(value float64) error {
	return h.SetAttributeDouble("low", value)
}

func (h HtmlMeterElement) Max() (float64, error) {
	return h.GetAttributeDouble("max")
}

func (h HtmlMeterElement) SetMax(value float64) error {
	return h.SetAttributeDouble("max", value)
}

func (h HtmlMeterElement) Min() (float64, error) {
	return h.GetAttributeDouble("min")
}

func (h HtmlMeterElement) SetMin(value float64) error {
	return h.SetAttributeDouble("min", value)
}

func (h HtmlMeterElement) Optimum() (float64, error) {
	return h.GetAttributeDouble("optimum")
}

func (h HtmlMeterElement) SetOptimum(value float64) error {
	return h.SetAttributeDouble("optimum", value)
}

func (h HtmlMeterElement) Value() (float64, error) {
	return h.GetAttributeDouble("value")
}

func (h HtmlMeterElement) SetValue(value float64) error {
	return h.SetAttributeDouble("value", value)
}

func (h HtmlMeterElement) Labels() (nodelist.NodeList, error) {
	var obj js.Value
	var err error
	var arr nodelist.NodeList
	if obj = h.GetValueByKey("labels"); obj.Error() == nil {

		arr, err = nodelist.NewFromJSObject(obj)
	}
	return arr, err
}
