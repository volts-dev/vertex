package event

// partial implemented
// https://developer.mozilla.org/fr/docs/Web/API/Event

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var eventinterface js.Value

// Event Event struct
type Event struct {
	js.Object
}

type EventFrom interface {
	Event_() Event
}

func (e Event) Event_() Event {
	return e
}

// GetInterface get the JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		if eventinterface = js.Global().Get("Event"); eventinterface.Error() != nil {
			eventinterface = js.Undefined()
		}
		js.Register(eventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return eventinterface
}

// New Create a event
func New(typeevent string, init ...map[string]interface{}) (Event, error) {
	var e Event
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if ei := GetInterface(); !ei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj = ei.New(arrayJS...); obj.Error() == nil {
			e.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (Event, error) {
	var e Event
	var err error
	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventi) {
				e.SetObjectValue(obj)

			} else {
				err = ErrNotAnEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func (e Event) Target() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj = e.GetValueByKey("target"); obj.Error() == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			bobj, err = js.Discover(obj)
		}

	}
	return bobj, err
}
func (e Event) CurrentTarget() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj = e.GetValueByKey("currentTarget"); obj.Error() == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			bobj, err = js.Discover(obj)
		}

	}
	return bobj, err
}

func (e Event) PreventDefault() error {
	var err error
	err = e.Call("preventDefault").Error()

	return err
}

func (e Event) StopImmediatePropagation() error {
	var err error
	err = e.Call("stopImmediatePropagation").Error()

	return err
}

func (e Event) StopPropagation() error {
	var err error
	err = e.Call("stopPropagation").Error()

	return err
}

func (e Event) Bubbles() (bool, error) {
	return e.GetAttributeBool("bubbles")
}

func (e Event) Cancelable() (bool, error) {
	return e.GetAttributeBool("cancelable")
}

func (e Event) Composed() (bool, error) {
	return e.GetAttributeBool("composed")
}

func (e Event) EventPhase() (int, error) {
	return e.GetAttributeInt("eventPhase")
}

func (e Event) Type() (string, error) {
	return e.GetAttributeString("type")
}

func (e Event) IsTrusted() (bool, error) {
	return e.GetAttributeBool("isTrusted")
}
