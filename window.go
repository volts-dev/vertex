//go:build js && wasm

package vertex

import (
	"errors"
	"net/url"
	"sync"

	"github.com/volts-dev/volts/vertex/core/js"
)

type (
	IWindow interface {
		URL() *url.URL
		Size() (width int, height int)
		CursorPosition() (x, y int)
		setCursorPosition(x, y int)
		GetElementByID(id string) js.Value // TODO: make it IComponent
	}

	window struct {
		EventListener
		body    IComponent
		cursorX int
		cursorY int
	}
)

var windowInterface js.Value
var defaultWindow *window

func init() {
	RegisterInterface(WindowInterface)
}

// GetInterface get the JS interface of formdata
func WindowInterface() js.Value {
	sync.OnceFunc(func() {
		var err error
		if windowInterface, err = baseobject.Get(js.Global(), "Window"); err != nil {
			windowInterface = js.Undefined()
		}

		Register(windowInterface, func(v js.Value) (interface{}, error) {
			return ToWindow(v)
		})
		navigator.GetInterface()
		history.GetInterface()
		location.GetInterface()
		storage.GetInterface()
	})

	return windowInterface
}

func ToWindow(obj js.Value) (window, error) {
	var w window
	if wi := WindowInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.SetObject(obj)
			return w, nil
		}
	}

	return w, ErrNotImplemented
}
func Window() *window {
	return defaultWindow
}

func bindingWindow() *window {
	return &window{Value: js.Global()}
}

func (w *window) URL() *url.URL {
	rawurl := w.
		Get("location").
		Get("href").
		String()

	u, _ := url.Parse(rawurl)
	return u
}

func (w *window) Size() (width int, height int) {
	getSize := func(axis string) int {
		size := w.Get("inner" + axis)
		if !size.Truthy() {
			size = w.
				Get("document").
				Get("documentElement").
				Get("client" + axis)
		}
		if !size.Truthy() {
			size = w.
				Get("document").
				Get("body").
				Get("client" + axis)
		}

		if size.Type() != js.TypeNumber {
			return 0
		}
		return size.Int()
	}

	return getSize("Width"), getSize("Height")
}

func (w *window) CursorPosition() (x, y int) {
	return w.cursorX, w.cursorY
}

func (w *window) setCursorPosition(x, y int) {
	w.cursorX = x
	w.cursorY = y
}

func (w *window) GetElementByID(id string) js.Value {
	return w.Get("document").Call("getElementById", id)
}

func (w *window) ScrollToID(id string) {
	if elem := w.GetElementByID(id); elem.Truthy() {
		elem.Call("scrollIntoView")
	}
}

func (w *window) setBody(body IComponent) {
	w.body = body
}

func (w *window) createElement(tag, xmlns string) (js.Value, error) {
	var element js.Value
	if xmlns == "" {
		element = w.Get("document").Call("createElement", tag)
	} else {
		element = w.Get("document").Call("createElementNS", xmlns, tag)
	}

	if !element.Truthy() {
		return nil, errors.New("creating javascript element failed").
			WithTag("tag", tag).
			WithTag("xmlns", xmlns)
	}
	return element, nil
}

func (w *window) createTextNode(v string) js.Value {
	return w.Get("document").Call("createTextNode", v)
}

func (w *window) addHistory(u *url.URL) {
	u.Scheme = w.URL().Scheme
	u.Host = w.URL().Host
	w.Get("history").Call("pushState", nil, "", u.String())
}

func (w *window) replaceHistory(u *url.URL) {
	u.Scheme = w.URL().Scheme
	u.Host = w.URL().Host
	w.Get("history").Call("replaceState", nil, "", u.String())
}
