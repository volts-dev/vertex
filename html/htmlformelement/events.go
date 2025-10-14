package htmlformelement

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (h HtmlFormElement) OnFormData(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("formdata", handler)
}

func (h HtmlFormElement) OnReset(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("reset", handler)
}

func (h HtmlFormElement) OnSubmit(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("submit", handler)
}
