package app

import (
	"sync"

	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/initinterface"
)

var (
	//defaultRouter = router.New("")
	//window = newBrowserWindow()
	singleton sync.Once
)

// Route associates a given path with a function that generates a new Composer
// component. When a user navigates to the specified path, the function
// newComponent is invoked to create and mount the associated component.
//
// Example:
//
//	Route("/home", func() Composer {
//	    return NewHomeComponent()
//	})
func Route(path string, rootElement func() htmlelement.HtmlElement) {
	//defaultRouter.Route(path, rootElement)
}

func RouteWithRegexp(pattern string, rootElement func() htmlelement.HtmlElement) {
	//defaultRouter.RouteWithRegexp(pattern, rootElement)
}
func Init() {
	initinterface.Init()
}

func Start() {
	if IsServer {
		return
	}

	defer func() {
		err := recover()
		//displayLoadError(err)
		panic(err)
	}()

	env, err := Getenv("VERTEX_STATIC_RESOURCES_URL")
	if err != nil && env != "" {
		return
	}

	/*
		resolveURL := clientResourceResolver(env)
		originPage := newPage(DefaultWindow().URL(), resolveURL)

		engine := newEngine(
			context.Background(),
			defaultRouter,
			resolveURL,
			&originPage,
			//actionHandlers,
		)

		engine.Navigate(window.Default().URL(), false)
		engine.Start(120)
	*/
}
