package htmltableelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/htmltablecaptionelement"
	"github.com/volts-dev/vertex/html/htmltablerowelement"
	"github.com/volts-dev/vertex/html/htmltablesectionelement"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltableelementinterface js.Value

// HtmlTableElement struct
type HtmlTableElement struct {
	htmlelement.HtmlElement
}

type HtmlTableElementFrom interface {
	HtmlTableElement_() HtmlTableElement
}

func (h HtmlTableElement) HtmlTableElement_() HtmlTableElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltableelementinterface = js.Global().Get("HTMLTableElement"); htmltableelementinterface.Error() != nil {
			htmltableelementinterface = js.Undefined()
		}
		js.Register(htmltableelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmltableelementinterface
}

func New(d document.Document) (HtmlTableElement, error) {
	var err error

	var h HtmlTableElement
	var e element.Element

	if e, err = d.CreateElement("table"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableElement, error) {
	var h HtmlTableElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableColElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableElement, error) {
	var h HtmlTableElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableColElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableElement) Caption() (htmltablecaptionelement.HtmlTableCaptionElement, error) {

	var err error
	var obj js.Value
	var caption htmltablecaptionelement.HtmlTableCaptionElement

	if obj = h.GetValueByKey("caption"); obj.Error() == nil {

		caption, err = htmltablecaptionelement.NewFromJSObject(obj)
	}

	return caption, err

}

func (h HtmlTableElement) SetCaption(caption htmltablecaptionelement.HtmlTableCaptionElement) error {
	h.SetValueByKey("caption", caption.GetObjectValue())
	return h.GetObjectValue().Error()
}

func (h HtmlTableElement) getCollectionMethod(method string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey(method); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableElement) getElemMethod(method string) (htmltablesectionelement.HtmlTableSectionElement, error) {

	var err error
	var obj js.Value
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj = h.GetValueByKey(method); obj.Error() == nil {

		elem, err = htmltablesectionelement.NewFromJSObject(obj)
	}

	return elem, err
}

func (h HtmlTableElement) Rows() (htmlcollection.HtmlCollection, error) {
	return h.getCollectionMethod("rows")

}

func (h HtmlTableElement) TBodies() (htmlcollection.HtmlCollection, error) {
	return h.getCollectionMethod("tBodies")
}

func (h HtmlTableElement) TFoot() (htmltablesectionelement.HtmlTableSectionElement, error) {
	return h.getElemMethod("tFoot")
}

func (h HtmlTableElement) THead() (htmltablesectionelement.HtmlTableSectionElement, error) {
	return h.getElemMethod("tHead")
}

func (h HtmlTableElement) CreateCaption() (htmltablecaptionelement.HtmlTableCaptionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablecaptionelement.HtmlTableCaptionElement

	if obj = h.Call("createCaption"); obj.Error() == nil {
		elem, err = htmltablecaptionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTHead() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj = h.Call("createTHead"); obj.Error() == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTFoot() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj = h.Call("createTFoot"); obj.Error() == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteCaption() error {

	err := h.Call("deleteCaption").Error()
	return err
}

func (h HtmlTableElement) DeleteTHead() error {

	err := h.Call("deleteTHead").Error()
	return err
}

func (h HtmlTableElement) DeleteTFoot() error {

	err := h.Call("deleteTFoot").Error()
	return err
}

func (h HtmlTableElement) InsertRow(index ...int) (htmltablerowelement.HtmlTableRowElement, error) {
	var obj js.Value
	var err error
	var elem htmltablerowelement.HtmlTableRowElement

	var arrayJS []interface{}

	if len(index) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(index[0]))
	}

	if obj = h.Call("insertRow", arrayJS...); obj.Error() == nil {
		elem, err = htmltablerowelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteRow(index int) error {

	var err error
	err = h.Call("deleteRow", js.ValueOf(index)).Error()
	return err
}
