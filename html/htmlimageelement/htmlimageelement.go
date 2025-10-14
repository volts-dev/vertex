package htmlimageelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlimageelementinterface js.Value

// HtmlImageElement struct
type HtmlImageElement struct {
	htmlelement.HtmlElement
}

type HtmlImageElementFrom interface {
	HtmlImageElement_() HtmlImageElement
}

func (h HtmlImageElement) HtmlImageElement_() HtmlImageElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlimageelementinterface = js.Global().Get("HTMLImageElement"); htmlimageelementinterface.Error() != nil {
			htmlimageelementinterface = js.Undefined()
		}
		js.Register(htmlimageelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlimageelementinterface
}

func New(d document.Document) (HtmlImageElement, error) {
	var err error

	var h HtmlImageElement
	var e element.Element

	if e, err = d.CreateElement("img"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlImageElement, error) {
	var h HtmlImageElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmImageElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlImageElement, error) {
	var h HtmlImageElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmImageElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlImageElement) Alt() (string, error) {
	return h.GetAttributeString("alt")
}

func (h HtmlImageElement) SetAlt(value string) error {
	return h.SetAttributeString("alt", value)
}

func (h HtmlImageElement) Complete() (bool, error) {
	return h.GetAttributeBool("complete")
}

func (h HtmlImageElement) CrossOrigin() (string, error) {
	return h.GetAttributeString("crossOrigin")
}

func (h HtmlImageElement) SetCrossOrigin(value string) error {
	return h.SetAttributeString("crossOrigin", value)
}

func (h HtmlImageElement) CurrentSrc() (string, error) {
	return h.GetAttributeString("currentSrc")
}

func (h HtmlImageElement) Decoding() (string, error) {
	return h.GetAttributeString("decoding")
}

func (h HtmlImageElement) SetDecoding(value string) error {
	return h.SetAttributeString("decoding", value)
}
func (h HtmlImageElement) Height() (int, error) {
	return h.GetAttributeInt("height")
}

func (h HtmlImageElement) SetHeight(value int) error {
	return h.SetAttributeInt("height", value)
}

func (h HtmlImageElement) IsMap() (bool, error) {
	return h.GetAttributeBool("isMap")
}

func (h HtmlImageElement) SetIsMap(value bool) error {
	return h.SetAttributeBool("isMap", value)
}

func (h HtmlImageElement) Loading() (string, error) {
	return h.GetAttributeString("loading")
}

func (h HtmlImageElement) SetLoading(value string) error {
	return h.SetAttributeString("loading", value)
}

func (h HtmlImageElement) NaturalHeight() (int, error) {
	return h.GetAttributeInt("naturalHeight")
}

func (h HtmlImageElement) NaturalWidth() (int, error) {
	return h.GetAttributeInt("naturalWidth")
}

func (h HtmlImageElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlImageElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlImageElement) Width() (int, error) {
	return h.GetAttributeInt("width")
}

func (h HtmlImageElement) SetWidth(value int) error {
	return h.SetAttributeInt("width", value)
}

func (h HtmlImageElement) X() (int, error) {
	return h.GetAttributeInt("x")
}

func (h HtmlImageElement) Y() (int, error) {
	return h.GetAttributeInt("y")
}

func (h HtmlImageElement) Decode() (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj = h.Call("decode"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
