package htmlelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLElement

import (
	"sync"

	"github.com/volts-dev/vertex/html/cssstyledeclaration"
	"github.com/volts-dev/vertex/html/dragevent"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/js"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlelementinterface js.Value

// HtmlInputElement struct
type HtmlElement struct {
	element.Element
}

type HtmlElementFrom interface {
	HtmlElement_() HtmlElement
}

func (h HtmlElement) HtmlElement_() HtmlElement {
	return h
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlelementinterface = js.Global().Get("HTMLElement"); htmlelementinterface.Error() != nil {
			htmlelementinterface = js.Undefined()
		}
		js.Register(htmlelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		dragevent.GetInterface()
	})

	return htmlelementinterface
}

func NewFromJSObject(obj js.Value) (HtmlElement, error) {
	var h HtmlElement
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlElement
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromElement(elem element.Element) (HtmlElement, error) {
	var h HtmlElement
	var err error

	if ai := GetInterface(); !ai.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(ai) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlElement) AccessKey() (string, error) {
	return h.GetAttributeString("accessKey")
}

func (h HtmlElement) SetAccessKey(value string) error {
	return h.SetAttributeString("accessKey", value)
}

func (h HtmlElement) AccessKeyLabel() (string, error) {
	return h.GetAttributeString("accessKeyLabel")
}

func (h HtmlElement) SetAccessKeyLabel(value string) error {
	return h.SetAttributeString("accessKeyLabel", value)
}

func (h HtmlElement) ClassName() (string, error) {
	return h.GetAttributeString("className")
}

func (h HtmlElement) SetClassName(classname string) error {
	return h.SetAttributeString("className", classname)
}

func (h HtmlElement) ContentEditable() (string, error) {
	return h.GetAttributeString("contentEditable")
}

func (h HtmlElement) SetContentEditable(value string) error {
	return h.SetAttributeString("contentEditable", value)
}

func (h HtmlElement) IsContentEditable() (bool, error) {
	return h.GetAttributeBool("isContentEditable")
}

func (h HtmlElement) Dataset(name string) (interface{}, error) {
	var err error
	var obj, objv js.Value
	var ret interface{}

	if obj = h.GetValueByKey("dataset"); obj.Error() == nil {

		if objv = obj.Get(name); objv.Error() == nil {
			if !objv.IsUndefined() {
				ret, err = js.GoValue(objv)
			} else {
				err = ErrDatasetNotFound
			}

		}

	}

	return ret, err

}

func (h HtmlElement) SetDataset(name string, value interface{}) error {

	var err error
	var obj js.Value

	if obj = h.GetValueByKey("dataset"); obj.Error() == nil {
		err = obj.Set(name, js.ValueOf(value)).Error()
	}
	return err
}

func (h HtmlElement) Dir() (string, error) {
	return h.GetAttributeString("dir")
}

func (h HtmlElement) Hidden() (bool, error) {
	return h.GetAttributeBool("hidden")
}

func (h HtmlElement) SetHidden(value bool) error {
	return h.SetAttributeBool("hidden", value)
}

func (h HtmlElement) SetDir(value string) error {
	return h.SetAttributeString("dir", value)
}

func (h HtmlElement) Lang() (string, error) {
	return h.GetAttributeString("lang")
}

func (h HtmlElement) InnerText() (string, error) {

	return h.GetAttributeString("innerText")
}

func (h HtmlElement) SetInnerText(value string) error {

	return h.SetAttributeString("innerText", value)
}

func (h HtmlElement) SetLang(value string) error {
	return h.SetAttributeString("lang", value)
}

func (h HtmlElement) OffsetHeight() (int, error) {
	return h.GetAttributeInt("offsetHeight")
}

func (h HtmlElement) OffsetLeft() (int, error) {
	return h.GetAttributeInt("offsetLeft")
}

func (h HtmlElement) OffsetParent() (js.Object, error) {
	var err error
	var obj js.Value
	var ret js.Object

	if obj = h.GetValueByKey("offsetParent"); obj.Error() == nil {
		if !obj.IsUndefined() {
			if !obj.IsNull() {
				ret, err = js.ToObject(obj)
			} else {
				err = ErrParentNotFound
			}

		} else {
			err = js.ErrNotAnObject
		}

	}
	return ret, err
}

func (h HtmlElement) OffsetTop() (int, error) {
	return h.GetAttributeInt("offsetTop")
}

func (h HtmlElement) OffsetWidth() (int, error) {
	return h.GetAttributeInt("offsetWidth")
}

func (h HtmlElement) Title() (string, error) {
	return h.GetAttributeString("title")
}

func (h HtmlElement) SetTitle(value string) error {
	return h.SetAttributeString("title", value)
}

func (h HtmlElement) Blur() error {
	err := h.Call("blur").Error()
	return err
}

func (h HtmlElement) Click() error {
	err := h.Call("click").Error()
	return err
}

func (h HtmlElement) Focus() error {
	err := h.Call("focus").Error()
	return err
}

func (h HtmlElement) Style() (cssstyledeclaration.CSSStyleDeclaration, error) {
	var err error
	var obj js.Value
	var ret cssstyledeclaration.CSSStyleDeclaration

	if obj = h.GetValueByKey("style"); obj.Error() == nil {

		if !obj.IsUndefined() {
			ret, err = cssstyledeclaration.NewFromJSObject(obj)
		} else {
			err = js.ErrNotAnObject
		}

	}
	return ret, err
}
