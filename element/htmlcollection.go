package element

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/core/js"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	//ErrNotImplemented      = errors.New("Browser not implemented HTMLCollection")
	ErrNotAnHTMLCollection = errors.New("Object is not a HTMLCollection")
)

func init() {
	RegisterInterface(HtmlCollectionInterface)
}

var singleton sync.Once

var htmlcollectioninterface js.Value

// HTMLCollection struct
type HtmlCollection struct {
	Object
}

type HtmlCollectionFrom interface {
	HtmlCollection_() HtmlCollection
}

func (h HtmlCollection) HtmlCollection_() HtmlCollection {
	return h
}

// GetInterface get the JS interface of formdata
func HtmlCollectionInterface() js.Value {
	singleton.Do(func() {
		if htmlcollectioninterface = js.Global().Get("HTMLCollection"); htmlcollectioninterface.IsUndefined() {
			htmlcollectioninterface = js.Undefined()
		}

		Register(htmlcollectioninterface, func(v js.Value) (interface{}, error) {
			return ToHtmlCollection(v)
		})
	})

	return htmlcollectioninterface
}

func ToHtmlCollection(obj js.Value) (HtmlCollection, error) {
	var h HtmlCollection
	var err error
	if fli := HtmlCollectionInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {
			if obj.InstanceOf(fli) {
				h.SetObject(obj)
			} else {
				err = ErrNotAnHTMLCollection
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlCollection) Item(index int) (interface{}, error) {

	var i interface{}
	var err error
	obj := h.JSObject().Index(index)
	if !obj.IsUndefined() && !obj.IsNull() {
		i, err = Discover(obj)
	}

	return i, err
}
