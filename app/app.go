package app

import (
	"net/url"
	"sync"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/context"
	"github.com/volts-dev/vertex/core/router"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/storage"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

var (
	defaultApplication *application

// defaultRouter = router.New("")
// window = newBrowserWindow()
// singleton sync.Once
)

type (
	application struct {
		ctx            *context.Context
		localStorage   storage.Storage
		sessionStorage storage.Storage
		browser        Browser
		router         *router.Router
		internalURLs   []string
		resolveURL     func(string) string
		//originPage     *requestPage
		//lastVisitedURL *url.URL

		//nodes   nodeManager
		//updates updateManager
		//body    HTMLBody

		dispatches chan func()
		defers     chan func()
		goroutines sync.WaitGroup

		//asynchronousActionHandlers map[string]ActionHandler
		//actions                    actionManager
		//states                     stateManager
	}
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
	js.Init()
}

func Default() *application {
	if defaultApplication == nil {
		defaultApplication = New()
	}

	return defaultApplication

}

func New() *application {
	localStorage, err := storage.New()
	if err != nil {
		console.Error(err)
		panic(err)
	}

	sessionStorage, err := storage.New()
	if err != nil {
		console.Error(err)
		panic(err)
	}

	return &application{
		localStorage:   localStorage,
		sessionStorage: sessionStorage,
	}
}

func (self *application) Start() {
	console.Info("start")
	if IsServer {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			js.RecoverHandler(err)
		}
		//displayLoadError(err)

	}()

	env, err := Getenv("VERTEX_STATIC_RESOURCES_URL")
	if err != nil && env != "" {
		console.Error(err)
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
	self.Navigate(window.Default().URL(), false)
	// 保持 Go 进程不退出
	select {}
}

func (self *application) Navigate(destination *url.URL, updateHistory bool) {
}

func (self *application) Use(mid Middleware) {

}
