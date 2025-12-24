package serviceworkerregistration

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (s ServiceWorkerRegistration) OnUpdateFound(handler func(e event.Event) error) (js.Func, error) {

	return s.AddEventListener("updatefound", handler)
}
