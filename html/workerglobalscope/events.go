package workerglobalscope

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (w WorkerGlobalScope) OnError(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("error", handler)
}

func (w WorkerGlobalScope) OnLanguageChange(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("languagechange", handler)
}

func (w WorkerGlobalScope) OnOffline(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("offline", handler)
}

func (w WorkerGlobalScope) OnOnline(handler func(e event.Event) error) (js.Func, error) {

	return w.AddEventListener("online", handler)
}
