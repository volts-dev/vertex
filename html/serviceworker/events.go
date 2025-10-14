package serviceworker

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (s ServiceWorker) OnStateChange(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("statechange", handler)
}

func (s ServiceWorker) OnError(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
