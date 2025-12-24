package htmlselectelement

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (h HtmlSelectElement) OnInput(handler func(e event.Event) error) (js.Func, error) {
	return h.AddEventListener("input", handler)
}
