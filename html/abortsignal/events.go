package abortsignal

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (a AbortSignal) OnAbort(handler func(e event.Event)) (js.Func, error) {

	return a.AddEventListener("abort", handler)
}
