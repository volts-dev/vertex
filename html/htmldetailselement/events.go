package htmldetailselement

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (h HtmlDetailsElement) OnToggle(handler func(e event.Event) error) (js.Func, error) {

	return h.AddEventListener("toggle", handler)
}
