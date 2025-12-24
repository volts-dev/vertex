package eventtarget

// https://developer.mozilla.org/fr/docs/Web/API/EventTarget/EventTarget
import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var eventtargetinterface js.Value

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if eventtargetinterface = js.Global().Get("EventTarget"); eventtargetinterface.Error() != nil {
			eventtargetinterface = js.Undefined()
		}

		js.Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return eventtargetinterface
}

type EventTarget struct {
	event.Event
}

type EventTargetFrom interface {
	EventTarget_() EventTarget
}

func (e EventTarget) EventTarget_() EventTarget {
	return e
}

func New() (EventTarget, error) {

	var e EventTarget
	var obj js.Value
	var err error
	if eti := GetInterface(); !eti.IsUndefined() {

		if obj = eti.New(); obj.Error() == nil {
			e.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (EventTarget, error) {
	var e EventTarget
	var err error
	if eti := GetInterface(); !eti.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eti) {
				e.SetObjectValue(obj)

			} else {
				err = ErrNotAnEventTarget
			}
		}
	}

	return e, err
}

func (e EventTarget) AddEventListener(name string, handler func(e event.Event) error) (js.Func, error) {
	var err error
	var cb js.Func
	if handler != nil {
		cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if e, err := event.NewFromJSObject(args[0]); err == nil {
				handler(e)
			}
			return nil
		})
		defer cb.Release()

		err = e.Call("addEventListener", js.ValueOf(name), cb).Error()
	}

	return cb, err
}

func (e EventTarget) RemoveEventListener(name string, handler func(e event.Event) error) error {
	var err error
	var cb js.Func
	if handler != nil {
		cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if e, err := event.NewFromJSObject(args[0]); err == nil {
				handler(e)
			}
			return nil
		})
		defer cb.Release()
		err = e.Call("removeEventListener", js.ValueOf(name), cb).Error()
	}

	return err
}

func (e EventTarget) DispatchEvent(event event.Event) error {
	var err error
	err = e.Call("dispatchEvent", event.GetObjectValue()).Error()
	return err
}
