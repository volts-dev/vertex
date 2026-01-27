package router

import (
	"context"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/window"
)

type Router struct {
}

func New() *Router {
	router := &Router{}

	window.Default().AddEventListener("click", router.onClick)
	window.Default().AddEventListener("popstate", router.onPopState)

	return router
}

func (self *Router) Route(url string, com component.Component) {

}

func (self *Router) Handle(ctx context.Context) {

}

func (self *Router) onClick(e event.Event) error {

	return nil
}

func (self *Router) onPopState(e event.Event) error {

	return nil
}
