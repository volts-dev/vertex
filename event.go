//go:build js && wasm

package vertex

import (
	"errors"
	"sync"

	"github.com/volts-dev/volts/vertex/core/js"
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
		Object
	}

	EventListener struct {
		Event
	}
)

var eventtargetinterface js.Value
var eventinterface js.Value

func init() {
	RegisterInterface(EventInterface)
	RegisterInterface(EventListenerInterface)
}

// GetInterface get the JS interface of event
func EventInterface() js.Value {
	sync.OnceFunc(func() {
		if eventinterface = js.Global().Get("Event"); eventinterface.IsUndefined() {
			eventinterface = js.Undefined()
		}

		Register(eventinterface, func(v js.Value) (interface{}, error) {
			return ToEvent(v)
		})

	})

	return eventinterface
}

// GetJSInterface get the JS interface
func EventListenerInterface() js.Value {
	sync.OnceFunc(func() {
		if eventtargetinterface = js.Global().Get("EventTarget"); eventtargetinterface.IsUndefined() {
			eventtargetinterface = js.Undefined()
		}

		Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
			return ToEvent(v)
		})
	})

	return eventtargetinterface
}
func ToEvent(obj js.Value) (Event, error) {
	var e Event
	var err error
	if eventi := EventInterface(); !eventi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventi) {
				e.Object = e.SetObject(obj)

			} else {
				err = ErrNotAnEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return e, err
}
func ToEventListener(obj js.Value) (EventListener, error) {
	var e EventListener
	var err error
	if eti := EventListenerInterface(); !eti.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(eti) {
				e.Object = e.SetObject(obj)

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
		err = ErrUndefinedValue
	} else {
		bobj, err = Discover(obj)
	}

	return bobj, err
}
func (e Event) CurrentTarget() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	obj = e.Get("currentTarget")
	if obj.IsUndefined() || obj.IsNull() {
		err = ErrUndefinedValue
	} else {
		bobj, err = Discover(obj)
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

func (e EventListener) AddEventListener(name string, handler func(e Event)) (js.Func, error) {
	var err error
	var cb js.Func
	if handler != nil {
		cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if e, err := ToEvent(args[0]); err == nil {
				handler(e)
			}
			return nil
		})

		e.Call("addEventListener", js.ValueOf(name), cb)
	}

	return cb, err
}

func (e EventListener) RemoveEventListener(f js.Func, typeevent string) error {
	var err error
	e.Call("removeEventListener", typeevent, f)
	f.Release()
	return err
}

func (e EventListener) DispatchEvent(event Event) error {
	var err error
	e.Call("dispatchEvent", event.JSObject())
	return err
}
