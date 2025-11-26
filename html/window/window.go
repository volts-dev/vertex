package window

import (
	"net/url"
	"sync"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/history"
	"github.com/volts-dev/vertex/html/indexeddb"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/location"
	"github.com/volts-dev/vertex/html/navigator"
	"github.com/volts-dev/vertex/html/storage"
	"github.com/volts-dev/vertex/js"
)

func init() {
	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once
var windowinterface js.Value
var defaultWindow *Window

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {
	singleton.Do(func() {
		if windowinterface = js.Global().Get("Window"); windowinterface.Error() != nil {
			windowinterface = js.Undefined()
		}
		js.Register(windowinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		navigator.GetInterface()
		history.GetInterface()
		location.GetInterface()
		storage.GetInterface()
	})

	return windowinterface
}

type Window struct {
	eventtarget.EventTarget
}

type WindowFrom interface {
	Window_() Window
}

func Default() *Window {
	if defaultWindow != nil {
		return defaultWindow
	}

	w, err := js.Self()
	if err != nil {
		return nil
	}
	win, ok := w.(Window)
	if !ok {
		console.Error("cant allocate self")
		return nil
	}
	defaultWindow = &win
	return defaultWindow
}

func (w Window) Window_() Window {
	return w
}

func NewFromJSObject(obj js.Value) (Window, error) {
	var w Window

	if wi := GetInterface(); !wi.IsUndefined() {

		if obj.InstanceOf(wi) {
			w.SetObjectValue(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func New() (Window, error) {
	var err error
	var w Window
	var windowObj js.Value
	if windowObj = js.Global().Get("window"); err == nil {

		w, err = NewFromJSObject(windowObj)

	}
	return w, err
}

func (w Window) URL() *url.URL {
	rawurl, _ := w.
		GetValueByKey("location").
		Get("href").
		String()

	u, _ := url.Parse(rawurl)
	return u
}

func (w Window) Blur() error {
	return w.GetValueByKey("blur").Error()
}

func (w Window) Document() (document.Document, error) {
	var err error
	var obj js.Value
	var d document.Document

	if obj = w.GetValueByKey("document"); obj.Error() == nil {
		d, err = document.NewFromJSObject(obj)
	}

	return d, err
}

func (w Window) History() (history.History, error) {
	var err error
	var obj js.Value
	var h history.History

	if obj = w.GetValueByKey("history"); obj.Error() == nil {
		h, err = history.NewFromJSObject(obj)
	}

	return h, err
}

func (w Window) Location() (location.Location, error) {
	var err error
	var obj js.Value
	var l location.Location

	if obj = w.GetValueByKey("location"); obj.Error() == nil {
		l, err = location.NewFromJSObject(obj)
	}

	return l, err
}

func (w Window) LocalStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj = w.GetValueByKey("localStorage"); obj.Error() == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) SessionStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj = w.GetValueByKey("sessionStorage"); obj.Error() == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) IndexdedDB() (indexeddb.IDBFactory, error) {
	var err error
	var obj js.Value
	var i indexeddb.IDBFactory

	if obj = w.GetValueByKey("indexedDB"); obj.Error() == nil {
		i, err = indexeddb.IDBFactoryNewFromJSObject(obj)
	}

	return i, err
}

func (w Window) Navigator() (navigator.Navigator, error) {
	var err error
	var obj js.Value
	var n navigator.Navigator

	if obj = w.GetValueByKey("navigator"); obj.Error() == nil {
		n, err = navigator.NewFromJSObject(obj)
	}

	return n, err
}
