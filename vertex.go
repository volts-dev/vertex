package vertex

import (
	"context"
	"runtime"

	"github.com/volts-dev/vertex/core/js"
	"github.com/volts-dev/vertex/router"
)

const (
	// IsClient reports whether the code is running as a client in the
	// WebAssembly binary (app.wasm).
	IsClient = runtime.GOARCH == "wasm" && runtime.GOOS == "js"

	// IsServer reports whether the code is running on a server for
	// pre-rendering purposes.
	IsServer = runtime.GOARCH != "wasm" || runtime.GOOS != "js"
)

var (
	defaultRouter = router.New("")
	//window = newBrowserWindow()
)

func Init() {
	for _, f := range interfaces {
		f()
	}

	interfaces = make([]func() js.Value, 0)
}

// Route associates a given path with a function that generates a new Composer
// component. When a user navigates to the specified path, the function
// newComponent is invoked to create and mount the associated component.
//
// Example:
//
//	Route("/home", func() Composer {
//	    return NewHomeComponent()
//	})
func Route(path string, newComponent func() IComponent) {
	defaultRouter.Route(path, newComponent)
}

func Start() {
	if IsServer {
		return
	}

	defer func() {
		err := recover()
		displayLoadError(err)
		panic(err)
	}()

	resolveURL := clientResourceResolver(Getenv("GOAPP_STATIC_RESOURCES_URL"))
	originPage := makeRequestPage(Window().URL(), resolveURL)

	engine := newEngine(context.Background(),
		&routes,
		resolveURL,
		&originPage,
		actionHandlers,
	)

	engine.Navigate(Window().URL(), false)
	engine.Start(120)
}
