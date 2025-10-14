package html

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

type (
	EventTarget interface {
		js.Object
		AddEventListener(name string, handler func(e Event)) (js.Func, error)
		RemoveEventListener(f js.Func, typeevent string) error
		DispatchEvent(event Event) error
	}

	eventTarget struct {
		js.Object
	}

	EventTargetFrom interface {
		EventTarget_() EventTarget
	}
)

func init() {
	js.RegisterInterface(EventListenerInterface)
}

// GetJSInterface get the JS interface
func EventListenerInterface() js.Value {
	sync.OnceFunc(func() {
		if eventtargetinterface = js.Global().Get("EventTarget"); eventtargetinterface.IsUndefined() {
			eventtargetinterface = js.Undefined()
		}

		js.Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
			return ToEvent(v)
		})
	})

	return eventtargetinterface
}

func (e eventTarget) AddEventListener(name string, handler func(e Event)) (js.Func, error) {
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

func (e eventTarget) RemoveEventListener(f js.Func, typeevent string) error {
	var err error
	e.Call("removeEventListener", typeevent, f)
	f.Release()
	return err
}

func (e eventTarget) DispatchEvent(event Event) error {
	var err error
	e.Call("dispatchEvent", event.Value)
	return err
}
