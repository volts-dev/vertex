package domrectlist

//

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/domrect"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var domrectlistinterface js.Value

// DOMRectLists struct
type DOMRectList struct {
	js.Object
}

type DOMRectListFrom interface {
	DOMRectList_() DOMRectList
}

func (d DOMRectList) DOMRectList_() DOMRectList {
	return d
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if domrectlistinterface = js.Global().Get("DOMRectList"); domrectlistinterface.Error() != nil {
			domrectlistinterface = js.Undefined()
		}
		js.Register(domrectlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domrectlistinterface
}

func NewFromJSObject(obj js.Value) (DOMRectList, error) {
	var d DOMRectList
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				d.SetObjectValue(obj)

			} else {
				err = ErrNotAnDOMRectList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMRectList) Item(index int) (domrect.DOMRect, error) {

	return domrect.NewFromJSObject(d.GetObjectValue().Index(index))
}
