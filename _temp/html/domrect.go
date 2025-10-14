package html

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/js/reflect"
	"github.com/volts-dev/vertex/js"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var (
	ErrNotAnDOMRect = errors.New("The given value must be an DOMRect")
)

var domrectinterface js.Value

// GetJSInterface get the JS interface
func GetDOMRectInterface() js.Value {

	singleton.Do(func() {
		if domrectinterface = js.Global().Get("DOMRect"); domrectinterface.IsUndefined() {
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
	DOMRectReadOnly
}

type DOMRectFrom interface {
	DOMRect_() DOMRect
}

func (d DOMRect) DOMRect_() DOMRect {
	return d
}

func NewDOMRect() (DOMRect, error) {

	var d DOMRect
	var obj js.Object
	var err error
	if di := GetInterface(); !di.IsUndefined() {

		if obj, err = js.ToObject(di); err == nil {
			//d.SetValue(obj)
			d.Object = obj
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func ToDOMRect(obj js.Value) (DOMRect, error) {
	var d DOMRect
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				//d.SetValue(obj)
				d.SetValue(obj)
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
	d.SetAttributeDouble("height", value)
	return nil
}
func (d DOMRect) SetLeft(value float64) error {
	d.SetAttributeDouble("left", value)
	return nil
}
func (d DOMRect) SetRight(value float64) error {

	d.SetAttributeDouble("right", value)
	return nil
}

func (d DOMRect) SetTop(value float64) error {
	d.SetAttributeDouble("top", value)
	return nil
}

func (d DOMRect) SetWidth(value float64) error {
	d.SetAttributeDouble("width", value)
	return nil
}

func (d DOMRect) SetX(value float64) error {
	d.SetAttributeDouble("x", value)
	return nil
}

func (d DOMRect) SetY(value float64) error {
	d.SetAttributeDouble("y", value)
	return nil
}

func (d DOMRect) RectReadOnly() (DOMRectReadOnly, error) {
	var ro DOMRectReadOnly
	var err error
	if di := GetDOMRectReadOnlyInterface(); !di.IsUndefined() {
		if objro, err := reflect.Call(di, "fromRect", d.Object); err == nil {
			ro, err = ToDOMRectReadOnly(objro)
		}

	} else {
		err = js.ErrNotImplemented
	}
	return ro, err
}
