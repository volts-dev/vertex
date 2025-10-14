package domrect

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/domrectreadonly"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var domrectinterface js.Value

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if domrectinterface = js.Global().Get("DOMRect"); domrectinterface.Error() != nil {
			domrectinterface = js.Undefined()
		}
		js.Register(domrectinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domrectinterface
}

// DOMRectReadOnly struct
type DOMRect struct {
	domrectreadonly.DOMRectReadOnly
}

type DOMRectFrom interface {
	DOMRect_() DOMRect
}

func (d DOMRect) DOMRect_() DOMRect {
	return d
}

func New() (DOMRect, error) {

	var d DOMRect
	var obj js.Value
	var err error
	if di := GetInterface(); !di.IsUndefined() {

		if obj = di.New(); obj.Error() == nil {
			d.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func NewFromJSObject(obj js.Value) (DOMRect, error) {
	var d DOMRect
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetObjectValue(obj)

			} else {
				err = ErrNotAnDOMRect
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DOMRect) SetHeight(value float64) error {

	return d.SetAttributeDouble("height", value)
}
func (d DOMRect) SetLeft(value float64) error {

	return d.SetAttributeDouble("left", value)
}
func (d DOMRect) SetRight(value float64) error {

	return d.SetAttributeDouble("right", value)
}

func (d DOMRect) SetTop(value float64) error {

	return d.SetAttributeDouble("top", value)
}

func (d DOMRect) SetWidth(value float64) error {

	return d.SetAttributeDouble("width", value)
}

func (d DOMRect) SetX(value float64) error {

	return d.SetAttributeDouble("x", value)
}

func (d DOMRect) SetY(value float64) error {

	return d.SetAttributeDouble("y", value)
}

func (d DOMRect) RectReadOnly() (domrectreadonly.DOMRectReadOnly, error) {
	var ro domrectreadonly.DOMRectReadOnly
	var err error
	if di := domrectreadonly.GetInterface(); !di.IsUndefined() {

		if objro := di.Call("fromRect", d.GetObjectValue()); objro.Error() == nil {
			ro, err = domrectreadonly.NewFromJSObject(objro)
		}

	} else {
		err = domrectreadonly.ErrNotImplemented
	}
	return ro, err
}
