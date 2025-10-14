package html

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	//ErrNotImplemented      = errors.New("Browser not implemented HTMLCollection")
	ErrNotAnHTMLCollection = errors.New("Object is not a HTMLCollection")
)

func init() {
	js.RegisterInterface(HtmlCollectionInterface)
}

var singleton sync.Once

var htmlcollectioninterface js.Value

// HTMLCollection struct
type HTMLCollection struct {
	js.Object
}

type HtmlCollectionFrom interface {
	HtmlCollection_() HTMLCollection
}

func (h HTMLCollection) HtmlCollection_() HTMLCollection {
	return h
}

// GetInterface get the JS interface of formdata
func HtmlCollectionInterface() js.Value {
	singleton.Do(func() {
		if htmlcollectioninterface = js.Global().Get("HTMLCollection"); htmlcollectioninterface.IsUndefined() {
			htmlcollectioninterface = js.Undefined()
		}

		js.Register(htmlcollectioninterface, func(v js.Value) (interface{}, error) {
			return ToHTMLCollection(v)
		})
	})

	return htmlcollectioninterface
}

func ToHTMLCollection(obj js.Value) (HTMLCollection, error) {
	var h HTMLCollection
	var err error
	if fli := HtmlCollectionInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(fli) {
				h.SetValue(obj)
			} else {
				err = ErrNotAnHTMLCollection
			}
		}
	} else {
		err = js.ErrNotImplemented
	}

	return h, err
}

func (h HTMLCollection) Item(index int) (interface{}, error) {

	var i interface{}
	var err error
	obj := h.Value().Index(index)
	if !obj.IsUndefined() && !obj.IsNull() {
		i, err = js.Discover(obj)
	}

	return i, err
}
