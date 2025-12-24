package window

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (w Window) OnHashChange(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("hashchange", handler)
}

func (w Window) OnPopState(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("popstate", handler)
}
