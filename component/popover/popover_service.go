//go:build js && wasm

package popover

import (
	"github.com/volts-dev/vertex/js"
)

type PopoverService struct{}

type ServiceOptions struct {
	Animation        bool
	Arrow            bool
	CloseOnClickAway interface{} // bool or func(target js.Value) bool
	CloseOnEscape    bool
	FixedPosition    bool
	Position         string
	Class            string
	OnClose          func()
	OnPositioned     func(popover js.Value, solution string)
}

func NewService() *PopoverService {
	return &PopoverService{}
}

// Add creates and mounts a popover programmatically.
func (s *PopoverService) Add(target js.Value, options ServiceOptions) func() {
	doc := js.Global().Get("document")
	popoverEl := doc.Call("createElement", "web-popover")

	// Set attributes
	popoverEl.Set("target", js.ValueOf(target)) // Vertex should handle passing the element object

	if options.Position != "" {
		popoverEl.Call("setAttribute", "position", options.Position)
	}

	if options.Class != "" {
		popoverEl.Call("setAttribute", "class", options.Class)
	}

	if !options.Arrow {
		popoverEl.Call("setAttribute", "arrow", "false")
	}

	if !options.Animation {
		popoverEl.Call("setAttribute", "animation", "false")
	}

	// For complex props, we might need to set them directly if the component supports it.
	// Since web-popover is registered, we can't easily access the Go instance before it's mounted,
	// but Vertex's reactive system might handle attribute changes.

	// In the JS version, they mount it to an overlay. Here we'll append to document.body.
	doc.Get("body").Call("appendChild", popoverEl)

	remove := func() {
		if popoverEl.Truthy() && popoverEl.Get("parentNode").Truthy() {
			popoverEl.Call("remove")
		}
		if options.OnClose != nil {
			options.OnClose()
		}
	}

	return remove
}

// Global service instance
var DefaultService = NewService()

func Add(target js.Value, options ServiceOptions) func() {
	return DefaultService.Add(target, options)
}
