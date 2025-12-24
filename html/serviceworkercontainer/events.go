package serviceworkercontainer

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (s ServiceWorkerContainer) OnControllerChange(handler func(e event.Event) error) (js.Func, error) {

	return s.AddEventListener("controllerchange", handler)
}

func (s ServiceWorkerContainer) OnMessage(handler func(e event.Event) error) (js.Func, error) {

	return s.AddEventListener("message", handler)
}

func (s ServiceWorkerContainer) OnError(handler func(e event.Event) error) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
