package html

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

var (
	ErrNotAnEvent = errors.New("Object is not an Event")
	//ErrNotImplemented ErrNotImplemented error
	ErrEventNotImplemented = errors.New("Browser not implemented Event")
	//ErrNotAnEventTarget ErrNotAnEventTarget error
	ErrNotAnEventListener = errors.New("Object is not an EventListener")
)

type (
	Event struct {
		js.Object
	}
)

var eventtargetinterface js.Value
var eventinterface js.Value

func init() {
	js.RegisterInterface(EventInterface)
}

// GetInterface get the JS interface of event
func EventInterface() js.Value {
	sync.OnceFunc(func() {
		if eventinterface = js.Global().Get("Event"); eventinterface.IsUndefined() {
			eventinterface = js.Undefined()
		}

		js.Register(eventinterface, func(v js.Value) (interface{}, error) {
			return ToEvent(v)
		})

	})

	return eventinterface
}

func ToEvent(obj js.Value) (Event, error) {
	var e Event
	var err error
	if eventi := EventInterface(); !eventi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventi) {
				e.Object.SetValue(obj)

			} else {
				err = ErrNotAnEvent
			}
		}
	} else {
		err = errors.ErrNotImplemented
	}

	return e, err
}
func ToEventTarget(obj js.Value) (EventTarget, error) {
	var e eventTarget
	var err error
	if eti := EventListenerInterface(); !eti.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eti) {
				e.SetValue(obj)
			} else {
				err = ErrNotAnEventListener
			}
		}
	}

	return e, err
}
func (e Event) Target() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	obj = e.Get("target")
	if obj.IsUndefined() || obj.IsNull() {
		err = js.ErrUndefinedValue
	} else {
		bobj, err = js.Discover(obj)
	}

	return bobj, err
}

func (e Event) CurrentTarget() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	obj = e.Get("currentTarget")
	if obj.IsUndefined() || obj.IsNull() {
		err = js.ErrUndefinedValue
	} else {
		bobj, err = js.Discover(obj)
	}

	return bobj, err
}

func (e Event) PreventDefault() {
	e.Call("preventDefault")
}

func (e Event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}

func (e Event) StopPropagation() {
	e.Call("stopPropagation")
}

func (e Event) Bubbles() bool {
	v, err := e.Get("bubbles").Bool()
	if err != nil {
		return false
	}
	return v
}

func (e Event) Cancelable() bool {
	v, err := e.Get("cancelable").Bool()
	if err != nil {
		return false
	}
	return v
}

func (e Event) Composed() bool {
	v, err := e.Get("composed").Bool()
	if err != nil {
		return false
	}
	return v
}

func (e Event) EventPhase() int {
	v, err := e.Get("eventPhase").Int()
	if err != nil {
		return 0
	}
	return v
}

func (e Event) Type() string {
	v, err := e.Get("type").String()
	if err != nil {
		return ""
	}
	return v
}

func (e Event) IsTrusted() bool {
	v, err := e.Get("isTrusted").Bool()
	if err != nil {
		return false
	}
	return v
}
