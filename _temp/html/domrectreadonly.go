package html

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var (
	ErrNotAnDOMRectReadOnly = errors.New("The given value must be an DOMRectReadOnly")
)

var domrectreadonlyinterface js.Value

// GetInterface get the JS interface
func GetDOMRectReadOnlyInterface() js.Value {

	singleton.Do(func() {

		if domrectreadonlyinterface = js.Global().Get("DOMRectReadOnly"); domrectreadonlyinterface.IsUndefined() {
			domrectreadonlyinterface = js.Undefined()
		}
		js.Register(domrectreadonlyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domrectreadonlyinterface
}

// DOMRectReadOnly struct
type DOMRectReadOnly struct {
	js.Object
}

type DOMRectReadOnlyFrom interface {
	DOMRectReadOnly_() DOMRectReadOnly
}

func (d DOMRectReadOnly) DOMRectReadOnly_() DOMRectReadOnly {
	return d
}

func NewDOMRectReadOnly() (DOMRectReadOnly, error) {
	var d DOMRectReadOnly
	var obj js.Object
	var err error
	if di := GetDOMRectReadOnlyInterface(); !di.IsUndefined() {

		if obj, err = js.ToObject(di); err == nil {
			//d.SetValue(obj)
			d.Object = obj
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func ToDOMRectReadOnly(obj js.Value) (DOMRectReadOnly, error) {
	var d DOMRectReadOnly
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetValue(obj)

			} else {
				err = ErrNotAnDOMRectReadOnly
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DOMRectReadOnly) Bottom() (float64, error) {
	return d.GetAttributeDouble("bottom")
}

func (d DOMRectReadOnly) Height() (float64, error) {
	return d.GetAttributeDouble("height")
}

func (d DOMRectReadOnly) Left() (float64, error) {
	return d.GetAttributeDouble("left")
}
func (d DOMRectReadOnly) Right() (float64, error) {
	return d.GetAttributeDouble("right")
}
func (d DOMRectReadOnly) Top() (float64, error) {
	return d.GetAttributeDouble("top")
}
func (d DOMRectReadOnly) Width() (float64, error) {
	return d.GetAttributeDouble("width")
}

func (d DOMRectReadOnly) X() (float64, error) {
	return d.GetAttributeDouble("x")
}

func (d DOMRectReadOnly) Y() (float64, error) {
	return d.GetAttributeDouble("y")
}

// Impossible cyclic import use RectReadOnly
func (d DOMRectReadOnly) FromRect() {
	//TODO IMPLEMENT
}

func (d DOMRectReadOnly) ToJSON() (string, error) {

	return d.GetAttributeString("toJSON")
}
