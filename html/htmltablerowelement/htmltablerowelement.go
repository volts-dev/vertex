package htmltablerowelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/htmltablecellelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltablerowelementinterface js.Value

// HtmlTableRowElement struct
type HtmlTableRowElement struct {
	htmlelement.HtmlElement
}

type HtmlTableRowElementFrom interface {
	HtmlTableRowElement_() HtmlTableRowElement
}

func (h HtmlTableRowElement) HtmlTableRowElement_() HtmlTableRowElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltablerowelementinterface = js.Global().Get("HTMLTableRowElement"); htmltablerowelementinterface.Error() != nil {
			htmltablerowelementinterface = js.Undefined()
		}
		js.Register(htmltablerowelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablerowelementinterface
}

func New(d document.Document) (HtmlTableRowElement, error) {
	var err error

	var h HtmlTableRowElement
	var e element.Element

	if e, err = d.CreateElement("tr"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableRowElement, error) {
	var h HtmlTableRowElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableRowElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableRowElement, error) {
	var h HtmlTableRowElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableRowElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableRowElement) Cells() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = h.GetValueByKey("cells"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableRowElement) RowIndex() (int, error) {
	return h.GetAttributeInt("rowIndex")
}

func (h HtmlTableRowElement) SectionRowIndex() (int, error) {
	return h.GetAttributeInt("sectionRowIndex")
}

func (h HtmlTableRowElement) InsertCell(index ...int) (htmltablecellelement.HtmlTableCellElement, error) {
	var obj js.Value
	var err error
	var elem htmltablecellelement.HtmlTableCellElement
	var arrayJS []interface{}

	if len(index) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(index[0]))
	}

	if obj = h.Call("insertCell", arrayJS...); obj.Error() == nil {
		elem, err = htmltablecellelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableRowElement) DeleteCell(index int) error {

	var err error
	err = h.Call("deleteCell", js.ValueOf(index)).Error()

	return err
}
