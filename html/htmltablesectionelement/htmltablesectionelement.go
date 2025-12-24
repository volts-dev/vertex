package htmltablesectionelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/htmltablerowelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltablesectionelementinterface js.Value

// HtmlTableRowElement struct
type HtmlTableSectionElement struct {
	htmlelement.HtmlElement
}

type HtmlTableSectionElementFrom interface {
	HtmlTableSectionElement_() HtmlTableSectionElement
}

func (h HtmlTableSectionElement) HtmlTableSectionElement_() HtmlTableSectionElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltablesectionelementinterface = js.Global().Get("HTMLTableSectionElement"); htmltablesectionelementinterface.Error() != nil {
			htmltablesectionelementinterface = js.Undefined()
		}
		js.Register(htmltablesectionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmltablesectionelementinterface
}

func NewTBody(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("tbody"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewTHead(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("thead"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewTFoot(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("tfoot"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableSectionElement, error) {
	var h HtmlTableSectionElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableSectionElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableSectionElement, error) {
	var h HtmlTableSectionElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableSectionElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableSectionElement) Rows() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("rows"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableSectionElement) InsertRow(index ...int) (htmltablerowelement.HtmlTableRowElement, error) {
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

func (h HtmlTableSectionElement) DeleteRow(index int) error {

	var err error
	err = h.Call("deleteRow", js.ValueOf(index)).Error()
	return err
}
