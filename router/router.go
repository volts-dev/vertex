//go:build js && wasm

package router

import (
	"regexp"
	"sync"

	"github.com/volts-dev/vertex/element"
)

type (
	regexpRoute struct {
		regexp           *regexp.Regexp
		componentCreator func() element.IElement
	}

	Router struct {
		sync.RWMutex
		routes           map[string]func() IElement
		routesWithRegexp []regexpRoute
	}
)

func New(name string) *Router {
	return &Router{
		routes: make(map[string]func() IElement),
	}
}

func (r *Router) Route(path string, componentCreator func() Composer) {
	r.Lock()
	defer r.Unlock()

	r.routes[path] = componentCreator
}

func (r *Router) RouteWithRegexp(pattern string, newComponent func() IElement) {
	r.Lock()
	defer r.Unlock()

	r.routesWithRegexp = append(r.routesWithRegexp, regexpRoute{
		regexp:           regexp.MustCompile(pattern),
		componentCreator: newComponent,
	})
}

func (r *Router) Macth(path string) (Composer, bool) {
	r.RLock()
	defer r.RUnlock()

	if creator, routed := r.routes[path]; routed {
		return creator, true
	}

	for _, rwr := range r.routesWithRegexp {
		if rwr.regexp.MatchString(path) {
			return rwr.componentCreator(), true
		}
	}

	return nil, false
}
