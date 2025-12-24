package htmlselectelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/htmlformelement"
	"github.com/volts-dev/vertex/html/htmloptionelement"
	"github.com/volts-dev/vertex/html/htmloptionscollection"
	"github.com/volts-dev/vertex/html/validitystate"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlselectelementinterface js.Value

// HtmlSelectElement struct
type HtmlSelectElement struct {
	htmlelement.HtmlElement
}

type HtmlSelectElementFrom interface {
	HtmlSelectElement_() HtmlSelectElement
}

func (h HtmlSelectElement) HtmlSelectElement_() HtmlSelectElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlselectelementinterface = js.Global().Get("HTMLSelectElement"); htmlselectelementinterface.Error() != nil {
			htmlselectelementinterface = js.Undefined()
		}
		js.Register(htmlselectelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlselectelementinterface
}

func New(d document.Document) (HtmlSelectElement, error) {
	var err error

	var h HtmlSelectElement
	var e element.Element

	if e, err = d.CreateElement("select"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSelectElement, error) {
	var h HtmlSelectElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLSelectElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSelectElement, error) {
	var h HtmlSelectElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLSelectElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlSelectElement) Autofocus() (bool, error) {
	return h.GetAttributeBool("autofocus")
}

func (h HtmlSelectElement) SetAutofocus(value bool) error {
	return h.SetAttributeBool("autofocus", value)
}

func (h HtmlSelectElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlSelectElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlSelectElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var f htmlformelement.HtmlFormElement
	if obj = h.GetValueByKey("form"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			f, err = htmlformelement.NewFromJSObject(obj)
		}
	}
	return f, err
}

func (h HtmlSelectElement) Length() (int, error) {
	return h.GetAttributeInt("length")
}

func (h HtmlSelectElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlSelectElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlSelectElement) Options() (htmloptionscollection.HtmlOptionsCollection, error) {

	var err error
	var obj js.Value
	var optioncollection htmloptionscollection.HtmlOptionsCollection

	if obj = h.GetValueByKey("options"); obj.Error() == nil {

		optioncollection, err = htmloptionscollection.NewFromJSObject(obj)
	}

	return optioncollection, err
}

func (h HtmlSelectElement) Multiple() (bool, error) {
	return h.GetAttributeBool("multiple")
}

func (h HtmlSelectElement) SetMultiple(value bool) error {
	return h.SetAttributeBool("multiple", value)
}

func (h HtmlSelectElement) Required() (bool, error) {
	return h.GetAttributeBool("required")
}

func (h HtmlSelectElement) SetRequired(value bool) error {
	return h.SetAttributeBool("required", value)
}

func (h HtmlSelectElement) SelectedIndex() (int, error) {
	return h.GetAttributeInt("selectedIndex")
}

func (h HtmlSelectElement) SetSelectedIndex(value int) error {
	return h.SetAttributeInt("selectedIndex", value)
}

func (h HtmlSelectElement) SelectedOptions() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("selectedOptions"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlSelectElement) Size() (int, error) {
	return h.GetAttributeInt("size")
}

func (h HtmlSelectElement) SetSize(value int) error {
	return h.SetAttributeInt("size", value)
}

func (h HtmlSelectElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlSelectElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj = h.GetValueByKey("validity"); obj.Error() == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlSelectElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlSelectElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}

func (h HtmlSelectElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlSelectElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlSelectElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlSelectElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlSelectElement) SetCustomValidity(message string) error {

	err := h.Call("setCustomValidity", js.ValueOf(message)).Error()
	return err
}

func (h HtmlSelectElement) Add(elem htmloptionelement.HtmlOptionElement, before ...interface{}) error {
	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, elem.GetObjectValue())

	for _, value := range before {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	err = h.Call("add", arrayJS...).Error()
	return err
}

func (h HtmlSelectElement) Item(index int) (htmloptionelement.HtmlOptionElement, error) {

	var optelem htmloptionelement.HtmlOptionElement
	var jsobj js.Value
	var err error

	if jsobj = h.Call("item", js.ValueOf(index)); jsobj.Error() == nil {
		optelem, err = htmloptionelement.NewFromJSObject(jsobj)
	}
	return optelem, err
}

func (h HtmlSelectElement) NamedItem(str string) (htmloptionelement.HtmlOptionElement, error) {

	var optelem htmloptionelement.HtmlOptionElement
	var jsobj js.Value
	var err error

	if jsobj = h.Call("namedItem", js.ValueOf(str)); jsobj.Error() == nil {
		optelem, err = htmloptionelement.NewFromJSObject(jsobj)
	}
	return optelem, err
}

func (h HtmlSelectElement) Remove(index int) error {
	err := h.Call("remove", js.ValueOf(index)).Error()
	return err
}
