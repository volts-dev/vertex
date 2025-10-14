package eventsource

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/messageevent"
)

var singleton sync.Once

var sseinterface js.Value

// GetJSInterface Get the Event Source Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		if sseinterface = js.Global().Get("EventSource"); sseinterface.Error() != nil {
			sseinterface = js.Undefined()
		}

		messageevent.GetInterface()
		js.Register(sseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return sseinterface
}

// EventSource struct
type EventSource struct {
	eventtarget.EventTarget
}

type EventSourceFrom interface {
	EventSource_() EventSource
}

func (e EventSource) EventSource_() EventSource {
	return e
}

func NewFromJSObject(obj js.Value) (EventSource, error) {
	var e EventSource
	var err error
	if eventsourcei := GetInterface(); !eventsourcei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventsourcei) {

				e.SetObjectValue(obj)
			} else {
				err = ErrNotAnEventSource
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func New(url string, opts ...interface{}) (EventSource, error) {

	var arrayJS []interface{}
	var e EventSource
	var err error
	var obj js.Value

	arrayJS = append(arrayJS, url)
	for _, value := range opts {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	if eventsourcei := GetInterface(); !eventsourcei.IsUndefined() {

		if obj = eventsourcei.New(arrayJS...); obj.Error() == nil {
			e.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func (e EventSource) ReadyState() (int, error) {
	var err error
	var obj js.Value
	if obj = e.GetValueByKey("readyState"); obj.Error() == nil {

		return obj.Int()
	}
	return 0, err
}

func (e EventSource) Url() (string, error) {

	var err error
	var obj js.Value
	if obj = e.GetValueByKey("url"); obj.Error() == nil {

		return obj.String()
	}
	return "", err

}

func (e EventSource) Close() error {
	var err error
	err = e.Call("close").Error()
	return err
}

func (e EventSource) WithCredentials() (bool, error) {
	return e.GetAttributeBool("withCredentials")
}

func (sse EventSource) setHandler(jshandlername string, handler func(e event.Event)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})

	sse.GetObjectValue().Set(jshandlername, jsfunc)
}

// SetOnOpen Set onOpen Handler
func (sse EventSource) SetOnOpen(handler func(e event.Event)) {

	sse.setHandler("onopen", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnClose(handler func(e event.Event)) {
	sse.setHandler("onclose", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnError(handler func(e event.Event)) {
	sse.setHandler("onerror", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnMessage(handler func(e messageevent.MessageEvent)) {
	sse.setHandler("onmessage", func(e event.Event) {

		if obj, err := js.Discover(e.GetObjectValue()); err == nil {

			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

// OnOpen Set onOpen Handler
func (e EventSource) OnOpen(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("open", handler)
}

// OnClose Set onClose Handler
func (e EventSource) OnClose(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("close", handler)
}

// OnError Set onError Handler
func (e EventSource) OnError(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("error", handler)
}
