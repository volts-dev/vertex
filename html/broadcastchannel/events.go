package broadcastchannel

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/messageevent"
)

func (c BroadcastChannel) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return c.AddEventListener("message", func(e event.Event) error {

		if globalObj, err := js.Discover(e.GetObjectValue()); err == nil {

			if m, ok := globalObj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
		return nil
	})
}

func (c BroadcastChannel) OnMessageError(handler func(e event.Event) error) (js.Func, error) {

	return c.AddEventListener("messageerror", handler)
}
