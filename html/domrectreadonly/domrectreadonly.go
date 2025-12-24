package domrectreadonly

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var domrectreadonlyinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if domrectreadonlyinterface = js.Global().Get("DOMRectReadOnly"); domrectreadonlyinterface.Error() != nil {
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

func New() (DOMRectReadOnly, error) {

	var d DOMRectReadOnly
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

func NewFromJSObject(obj js.Value) (DOMRectReadOnly, error) {
	var d DOMRectReadOnly
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetObjectValue(obj)

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
