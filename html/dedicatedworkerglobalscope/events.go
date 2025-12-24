package dedicatedworkerglobalscope

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/messageevent"
)

func (d DedicatedWorkerGlobalScope) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {
	return d.AddEventListener("message", func(e event.Event) error {
		if globalObj, err := js.Discover(e.GetObjectValue()); err == nil {

			if m, ok := globalObj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
		return nil
	})
}

func (d DedicatedWorkerGlobalScope) OnMessageError(handler func(e event.Event) error) (js.Func, error) {
	return d.AddEventListener("onmessageerror", handler)
}
