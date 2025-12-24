package htmltextareaelement

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (h HtmlTextAreaElement) OnInput(handler func(e event.Event) error) (js.Func, error) {
	return h.AddEventListener("input", handler)
}
