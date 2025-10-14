package router

import (
	"regexp"
	"sync"

	"github.com/volts-dev/vertex/html/element"
)

type (
	regexpRoute struct {
		regexp           *regexp.Regexp
		componentCreator func() element.IHTMLElement
	}

	Router struct {
		sync.RWMutex
		routes           map[string]func() element.IHTMLElement
		routesWithRegexp []regexpRoute
	}
)

func New(name string) *Router {
	return &Router{
		routes: make(map[string]func() element.IHTMLElement),
	}
}

func (r *Router) Route(path string, componentCreator func() element.IHTMLElement) {
	r.Lock()
	defer r.Unlock()

	r.routes[path] = componentCreator
}

func (r *Router) RouteWithRegexp(pattern string, newComponent func() element.IHTMLElement) {
	r.Lock()
	defer r.Unlock()

	r.routesWithRegexp = append(r.routesWithRegexp, regexpRoute{
		regexp:           regexp.MustCompile(pattern),
		componentCreator: newComponent,
	})
}

func (r *Router) Macth(path string) (func() element.IHTMLElement, bool) {
	r.RLock()
	defer r.RUnlock()

	if creator, routed := r.routes[path]; routed {
		return creator, true
	}

	for _, rwr := range r.routesWithRegexp {
		if rwr.regexp.MatchString(path) {
			return rwr.componentCreator, true
		}
	}

	return nil, false
}
