package eventtarget

// https://developer.mozilla.org/fr/docs/Web/API/EventTarget/EventTarget
import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/core/console"
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
		eventtargetinterface = js.Global().Get("EventTarget")
		if eventtargetinterface.Error() != nil {
			eventtargetinterface = js.Undefined()
		} else {
			js.Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
				return NewFromJSObject(v)
			})
		}
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

// New 创建一个新的 EventTarget 实例
func New() (EventTarget, error) {
	eti := GetInterface()
	if eti.IsUndefined() {
		return EventTarget{}, ErrNotImplemented
	}

	obj := eti.New()
	if err := obj.Error(); err != nil {
		return EventTarget{}, err
	}

	e := EventTarget{}
	e.SetObjectValue(obj)
	return e, nil
}

// NewFromJSObject 从 JavaScript Object 创建 EventTarget
func NewFromJSObject(obj js.Value) (EventTarget, error) {
	if obj == nil {
		return EventTarget{}, js.ErrUndefinedValue
	}

	if obj.IsUndefined() || obj.IsNull() {
		return EventTarget{}, js.ErrUndefinedValue
	}

	eti := GetInterface()
	if eti.IsUndefined() {
		return EventTarget{}, ErrNotImplemented
	}

	if !obj.InstanceOf(eti) {
		return EventTarget{}, ErrNotAnEventTarget
	}

	e := EventTarget{}
	e.SetObjectValue(obj)
	return e, nil
}

// AddEventListener 添加事件监听器
// 注意：返回的 js.Func 需要保持引用以防止 GC，removeEventListener 时需要同一引用
func (e EventTarget) AddEventListener(name string, handler func(e event.Event) error) (js.Func, error) {
	if handler == nil {
		return nil, ErrInvalidHandler
	}

	// 创建包装器，捕获事件并调用处理函数
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			console.Warn("Event handler called with no arguments")
			return nil
		}

		// 创建事件对象
		if evt, err := event.NewFromJSObject(args[0]); err != nil {
			console.Error("Failed to create event object:", err)
			return nil
		} else {
			// 调用用户处理函数
			if err := handler(evt); err != nil {
				console.Error("Event handler error:", err)
			}
		}
		return nil
	})

	// 添加事件监听器
	callErr := e.Call("addEventListener", js.ValueOf(name), cb)
	if callErr.Error() != nil {
		cb.Release()
		return nil, callErr.Error()
	}

	return cb, nil
}

// RemoveEventListener 移除事件监听器
// 注意：handler 参数无法用于精确匹配移除，建议使用 RemoveEventListenerByFunc 并传入相同的 js.Func
func (e EventTarget) RemoveEventListener(name string, handler func(e event.Event) error) error {
	if handler == nil {
		return ErrInvalidHandler
	}

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})
	defer cb.Release()

	return e.Call("removeEventListener", js.ValueOf(name), cb).Error()
}

func (e EventTarget) RemoveEventListenerWithFunc(name string, handler js.Func) error {
	if handler == nil {
		return ErrInvalidHandler
	}

	defer handler.Release()
	return e.Call("removeEventListener", js.ValueOf(name), handler).Error()
}

// DispatchEvent 分派一个事件到 EventTarget
func (e EventTarget) DispatchEvent(event event.Event) error {
	var err error
	err = e.Call("dispatchEvent", event.GetObjectValue()).Error()
	return err
}
