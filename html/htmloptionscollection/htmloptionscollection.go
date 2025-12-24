package htmloptionscollection

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/htmlcollection"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmloptionscollectioninterface js.Value

// HTMLOptionsCollection struct
type HtmlOptionsCollection struct {
	htmlcollection.HtmlCollection
}

type HtmlOptionsCollectionFrom interface {
	HtmlOptionsCollection_() HtmlOptionsCollection
}

func (h HtmlOptionsCollection) HTMLOptionsCollection_() HtmlOptionsCollection {
	return h
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if htmloptionscollectioninterface = js.Global().Get("HTMLOptionsCollection"); htmloptionscollectioninterface.Error() != nil {
			htmloptionscollectioninterface = js.Undefined()
		}
		js.Register(htmloptionscollectioninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmloptionscollectioninterface
}

func NewFromJSObject(obj js.Value) (HtmlOptionsCollection, error) {
	var h HtmlOptionsCollection
	var err error
	if fli := GetInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fli) {
				h.SetObjectValue(obj)
			} else {
				err = ErrNotAnHTMLOptionsCollection
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlOptionsCollection) Length() (int, error) {

	return h.GetAttributeInt("length")

}
