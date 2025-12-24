package htmltablecellelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltablecellelementinterface js.Value

// HtmlTableCellElement struct
type HtmlTableCellElement struct {
	htmlelement.HtmlElement
}

type HtmlTableCellElementFrom interface {
	HtmlTableCellElement_() HtmlTableCellElement
}

func (h HtmlTableCellElement) HtmlTableCellElement_() HtmlTableCellElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmltablecellelementinterface = js.Global().Get("HTMLTableCellElement"); htmltablecellelementinterface.Error() != nil {
			htmltablecellelementinterface = js.Undefined()
		}
		js.Register(htmltablecellelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecellelementinterface
}

func NewTd(d document.Document) (HtmlTableCellElement, error) {
	var err error

	var h HtmlTableCellElement
	var e element.Element

	if e, err = d.CreateElement("td"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewTh(d document.Document) (HtmlTableCellElement, error) {
	var err error

	var h HtmlTableCellElement
	var e element.Element

	if e, err = d.CreateElement("th"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableCellElement, error) {
	var h HtmlTableCellElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHTMLTableCellElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableCellElement, error) {
	var h HtmlTableCellElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHTMLTableCellElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
