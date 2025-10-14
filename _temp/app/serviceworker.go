package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetServiceWorkerInterface)
}

var ErrNotAServiceWorker = errors.New("Object is not a ServiceWorker object")

var singletonServiceWorker sync.Once

var serviceworkerinterface js.Value

// GetInterface get the JS interface of serviceworker
func GetServiceWorkerInterface() js.Value {

	singletonServiceWorker.Do(func() {
		if serviceworkerinterface = js.Global().Get("ServiceWorker"); serviceworkerinterface.Error() != nil {
			serviceworkerinterface = js.Undefined()
		}
		js.Register(serviceworkerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		//promise.GetInterface()

	})

	return serviceworkerinterface
}

type ServiceWorker struct {
	html.EventTarget
}

type ServiceWorkerFrom interface {
	ServiceWorker_() ServiceWorker
}

func (s ServiceWorker) ServiceWorker_() ServiceWorker {
	return s
}

func ToServiceWorker(obj js.Value) (ServiceWorker, error) {
	var s ServiceWorker

	if si := GetServiceWorkerInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			//s.BaseObject = s.SetObject(obj)
			s.SetValue(obj)
			return s, nil

		}
	}

	return s, js.ErrNotImplemented
}

func (s ServiceWorker) ScriptURL() (string, error) {

	return s.GetAttributeString("scriptURL")
}

func (s ServiceWorker) State() (string, error) {

	return s.GetAttributeString("state")
}
func (s ServiceWorker) OnStateChange(handler func(e html.Event)) (js.Func, error) {

	return s.AddEventListener("statechange", handler)
}

func (s ServiceWorker) OnError(handler func(e html.Event)) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
