package websocket

// https://developer.mozilla.org/fr/docs/Web/API/WebSocket

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/blob"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/messageevent"
	"github.com/volts-dev/vertex/js/arraybuffer"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var wsinterface js.Value

// Websocket struct
type WebSocket struct {
	eventtarget.EventTarget
}

type WebSocketFrom interface {
	WebSocket_() WebSocket
}

func (w WebSocket) WebSocket_() WebSocket {
	return w
}

const (
	BlobType        = "blob"
	ArrayBufferType = "arraybuffer"
)

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if wsinterface = js.Global().Get("WebSocket"); wsinterface.Error() != nil {
			wsinterface = js.Undefined()
		}

		messageevent.GetInterface()
		js.Register(wsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return wsinterface
}

func NewFromJSObject(obj js.Value) (WebSocket, error) {
	var w WebSocket
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				w.SetObjectValue(obj)

			} else {
				err = ErrNotAWebSocket
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return w, err
}

// New Get a new channel broadcast
func New(url string) (WebSocket, error) {
	var ws WebSocket
	var err error
	var obj js.Value

	if wsi := GetInterface(); !wsi.IsUndefined() {
		if obj = wsi.New(js.ValueOf(url)); obj.Error() == nil {
			ws.SetObjectValue(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return ws, err
}

func (w WebSocket) setHandler(jshandlername string, handler func(e event.Event)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})

	w.GetObjectValue().Set(jshandlername, jsfunc)
}

// SetOnOpen Set onOpen Handler
func (w WebSocket) SetOnOpen(handler func(e event.Event)) {

	w.setHandler("onopen", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (w WebSocket) SetOnClose(handler func(e event.Event)) {
	w.setHandler("onclose", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (w WebSocket) SetOnError(handler func(e event.Event)) {
	w.setHandler("onerror", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (w WebSocket) SetOnMessage(handler func(e messageevent.MessageEvent)) {
	w.setHandler("onmessage", func(e event.Event) {

		if obj, err := js.Discover(e.GetObjectValue()); err == nil {

			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

// OnOpen Set onOpen Handler
func (w WebSocket) OnOpen(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("open", handler)
}

// OnClose Set onClose Handler
func (w WebSocket) OnClose(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("close", handler)
}

// OnError Set onError Handler
func (w WebSocket) OnError(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("error", handler)
}

func (w WebSocket) BinaryType() (string, error) {

	var err error
	var obj js.Value
	if obj = w.GetValueByKey("binaryType"); obj.Error() == nil {

		return obj.String()
	}
	return "", err

}

func (w WebSocket) SetBinaryType(binaryType string) error {

	switch binaryType {
	case BlobType:
	case ArrayBufferType:
	default:
		return ErrSetBadBinaryType
	}

	w.GetObjectValue().Set("binaryType", js.ValueOf(binaryType))

	return nil

}

// OnError Set onError Handler
func (w WebSocket) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return w.AddEventListener("message", func(e event.Event) {

		if obj, err := js.Discover(e.GetObjectValue()); err == nil {
			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

func (w WebSocket) Send(data interface{}) error {
	var object js.Value

	var err error
	switch d := data.(type) {
	case string:
		object = js.ValueOf(d)
	case arraybuffer.ArrayBuffer:
		object = d.GetObjectValue()
	case blob.Blob:
		object = d.GetObjectValue()
	default:
		err = ErrSendUnknownType
	}

	err = w.Call("send", object).Error()

	return err
}

func (w WebSocket) Close() error {

	var err error
	err = w.Call("close").Error()
	return err
}

func (w WebSocket) Protocol() (string, error) {

	var err error
	var obj js.Value
	if obj = w.GetValueByKey("protocol"); obj.Error() == nil {

		return obj.String()
	}
	return "", err

}

func (w WebSocket) BufferedAmount() (int, error) {
	var err error
	var obj js.Value
	if obj = w.GetValueByKey("bufferedAmount"); obj.Error() == nil {

		return obj.Int()
	}
	return 0, err
}

func (w WebSocket) ReadyState() (int, error) {
	var err error
	var obj js.Value
	if obj = w.GetValueByKey("readyState"); obj.Error() == nil {

		return obj.Int()
	}
	return 0, err
}

func (w WebSocket) Url() (string, error) {

	var err error
	var obj js.Value
	if obj = w.GetValueByKey("url"); obj.Error() == nil {

		return obj.String()
	}
	return "", err

}
