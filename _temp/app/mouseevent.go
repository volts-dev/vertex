package app

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetMouseEventInterface)
}

var ErrNotAMouseEvent = errors.New("Object is not a MouseEvent")

var mouseeventinterface js.Value

// MouseEvent MouseEvent struct
type MouseEvent struct {
	//Must be herited from mouseevent
	html.Event
}

type MouseEventFrom interface {
	MouseEvent_() MouseEvent
}

func (m MouseEvent) MouseEvent_() MouseEvent {
	return m
}

// GetInterface get MouseEvent interface
func GetMouseEventInterface() js.Value {

	singleton.Do(func() {

		if mouseeventinterface = js.Global().Get("MouseEvent"); mouseeventinterface.Error() != nil {
			mouseeventinterface = js.Undefined()
		}
		js.Register(mouseeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		html.EventListenerInterface()
	})

	return mouseeventinterface
}

func NewMouseEvent(typeevent string, init ...map[string]interface{}) (MouseEvent, error) {

	var m MouseEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if mei := GetMouseEventInterface(); !mei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj = mei.New(arrayJS...); obj.Error() == nil {
			//m.BaseObject = m.SetObject(obj)
			m.SetValue(obj)
		}

	} else {
		err = js.ErrNotImplemented
	}
	return m, err
}

func ToMouseEvent(obj js.Value) (MouseEvent, error) {
	var m MouseEvent
	var err error
	if mei := GetMouseEventInterface(); !mei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(mei) {
				//m.BaseObject = m.SetObject(obj)
				m.SetValue(obj)
			} else {
				err = ErrNotAMouseEvent
			}
		}
	} else {
		err = js.ErrNotImplemented
	}
	return m, err
}

func (m MouseEvent) AltKey() (bool, error) {

	return m.GetAttributeBool("altKey")
}

func (m MouseEvent) Button() (int, error) {

	return m.GetAttributeInt("button")
}

func (m MouseEvent) Buttons() (int, error) {

	return m.GetAttributeInt("buttons")
}

func (m MouseEvent) ClientX() (float64, error) {

	return m.GetAttributeDouble("clientX")
}

func (m MouseEvent) ClientY() (float64, error) {

	return m.GetAttributeDouble("clientY")
}

func (m MouseEvent) CtrlKey() (bool, error) {

	return m.GetAttributeBool("ctrlKey")
}

func (m MouseEvent) MetaKey() (bool, error) {

	return m.GetAttributeBool("metaKey")
}

func (m MouseEvent) MovementX() (int, error) {

	return m.GetAttributeInt("movementX")
}

func (m MouseEvent) MovementY() (int, error) {

	return m.GetAttributeInt("movementY")
}

func (m MouseEvent) OffsetX() (float64, error) {

	return m.GetAttributeDouble("offsetX")
}

func (m MouseEvent) OffsetY() (float64, error) {

	return m.GetAttributeDouble("offsetY")
}

func (m MouseEvent) PageX() (int, error) {

	return m.GetAttributeInt("pageX")
}

func (m MouseEvent) PageY() (int, error) {

	return m.GetAttributeInt("pageY")
}

func (m MouseEvent) Region() (string, error) {

	return m.GetAttributeString("region")
}

func (m MouseEvent) RelatedTarget() (html.EventTarget, error) {
	var err error
	var obj interface{}

	var e html.EventTarget

	if obj, err = m.GetAttributeGlobal("relatedTarget"); err == nil {

		if obj != nil {
			if efrom, ok := obj.(html.EventTargetFrom); ok {
				e = efrom.EventTarget_()
			}
		} else {
			err = js.ErrUndefinedValue
		}

	}
	return e, err
}

func (m MouseEvent) ScreenX() (float64, error) {

	return m.GetAttributeDouble("screenX")
}

func (m MouseEvent) ScreenY() (float64, error) {

	return m.GetAttributeDouble("screenY")
}

func (m MouseEvent) ShiftKey() (bool, error) {

	return m.GetAttributeBool("shiftKey")
}

func (m MouseEvent) X() (float64, error) {

	return m.GetAttributeDouble("x")
}

func (m MouseEvent) Y() (float64, error) {

	return m.GetAttributeDouble("y")
}

func (m MouseEvent) GetModifierState(args string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = m.Call("getModifierState", js.ValueOf(args)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}
