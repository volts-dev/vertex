package htmlfieldsetelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/validitystate"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlfieldsetelementinterface js.Value

// HtmlFieldSetElement struct
type HtmlFieldSetElement struct {
	htmlelement.HtmlElement
}

type HtmlFieldSetElementFrom interface {
	HtmlFieldSetElement_() HtmlFieldSetElement
}

func (h HtmlFieldSetElement) HtmlFieldSetElement_() HtmlFieldSetElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlfieldsetelementinterface = js.Global().Get("HTMLFieldSetElement"); htmlfieldsetelementinterface.Error() != nil {
			htmlfieldsetelementinterface = js.Undefined()
		}
		js.Register(htmlfieldsetelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlfieldsetelementinterface
}

func New(d document.Document) (HtmlFieldSetElement, error) {
	var err error

	var h HtmlFieldSetElement
	var e element.Element

	if e, err = d.CreateElement("fieldset"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlFieldSetElement, error) {
	var h HtmlFieldSetElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlFieldSetElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlFieldSetElement, error) {
	var h HtmlFieldSetElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlFieldSetElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlFieldSetElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlFieldSetElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlFieldSetElement) Elements() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("elements"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlFieldSetElement) Form() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("form"); obj.Error() == nil {
		if !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrNoForm
		}

	}

	return collection, err
}

func (h HtmlFieldSetElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlFieldSetElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlFieldSetElement) Type() (string, error) {

	return h.GetAttributeString("type")
}

func (h HtmlFieldSetElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlFieldSetElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj = h.GetValueByKey("validity"); obj.Error() == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlFieldSetElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlFieldSetElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlFieldSetElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlFieldSetElement) SetCustomValidity(message string) error {

	err := h.Call("setCustomValidity", js.ValueOf(message)).Error()
	return err
}
