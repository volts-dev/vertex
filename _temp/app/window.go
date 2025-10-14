package app

import (
	"errors"
	"net/url"
	"sync"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

type (
	__IWindow interface {
		html.EventTarget
		URL() *url.URL
		Size() (width int, height int)
		//CursorPosition() (x, y int)
		//setCursorPosition(x, y int)
		//GetElementByID(id string) js.Value // TODO: make it IComponent
		AddHistory(url *url.URL)
		ScrollToID(id string)
		Document() *Document
	}

	Window struct {
		html.EventTarget
		document *Document
		body     html.IHTMLElement
		cursorX  int
		cursorY  int
	}
)

var windowInterface js.Value
var defaultWindow *Window
var singleton sync.Once

func init() {
	js.RegisterInterface(WindowInterface)
}

// GetInterface get the JS interface of formdata
func WindowInterface() js.Value {
	sync.OnceFunc(func() {
		var err error
		if windowInterface = js.Global().Get("Window"); err != nil {
			windowInterface = js.Undefined()
		}

		js.Register(windowInterface, func(v js.Value) (interface{}, error) {
			return ToWindow(v)
		})

		GetNavigatorInterface()
		GetHistoryInterface()
		GetLocationInterface()
		GetStorageInterface()
	})

	return windowInterface
}

func ToWindow(obj js.Value) (Window, error) {
	var w Window
	if wi := WindowInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.SetValue(obj)
			return w, nil
		}
	}

	return w, js.ErrNotImplemented
}
func DefaultWindow() *Window {
	return defaultWindow
}

func bindingWindow() *Window {
	w, err := ToWindow(js.Global())
	if err != nil {
		console.Error(err)
		return nil
	}
	defaultWindow = &w
	return defaultWindow
}

func (w *Window) Document() *Document {
	if w.document == nil {
		document, err := ToDocument(w.
			Get("document"))
		if err != nil {
			console.Error(err)
			return nil
		}
		w.document = &document
	}
	return w.document
}
func (w Window) URL() *url.URL {
	rawurl, _ := w.
		Get("location").
		Get("href").
		String()

	u, _ := url.Parse(rawurl)
	return u
}

func (w Window) Size() (width int, height int) {
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

		v, err := size.Int()
		if err != nil {
			return 0
		}
		if size.Type() != js.TypeNumber {
			return 0
		}
		return v
	}

	return getSize("Width"), getSize("Height")
}

func (w *Window) CursorPosition() (x, y int) {
	return w.cursorX, w.cursorY
}

func (w *Window) setCursorPosition(x, y int) {
	w.cursorX = x
	w.cursorY = y
}

func (w *Window) GetElementByID(id string) js.Value {
	return w.Get("document").Call("getElementById", id)
}

func (w *Window) ScrollToID(id string) {
	if elem := w.GetElementByID(id); elem.Truthy() {
		elem.Call("scrollIntoView")
	}
}

func (w *Window) setBody(body html.IHTMLElement) {
	w.body = body
}

func (w *Window) CreateElement(tag, xmlns string) (html.IHTMLElement, error) {
	var jv js.Value
	if xmlns == "" {
		jv = w.Get("document").Call("createElement", tag)
	} else {
		jv = w.Get("document").Call("createElementNS", xmlns, tag)
	}

	if !jv.Truthy() {
		return nil, errors.New("creating javascript element failed")
	}

	return html.ToElement(jv)
}

func (w *Window) createTextNode(v string) js.Value {
	return w.document.Call("createTextNode", v)
}

func (w *Window) AddHistory(u *url.URL) {
	u.Scheme = w.URL().Scheme
	u.Host = w.URL().Host
	w.Get("history").Call("pushState", nil, "", u.String())
}

func (w *Window) replaceHistory(u *url.URL) {
	u.Scheme = w.URL().Scheme
	u.Host = w.URL().Host
	w.Get("history").Call("replaceState", nil, "", u.String())
}
