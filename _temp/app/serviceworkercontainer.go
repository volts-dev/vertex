package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetServiceWorkerContainerInterface)
}

var ErrNotAServiceWorkerContainer = errors.New("Object is not a ServiceWorkerContainer object")
var ErrControllerNotDefined = errors.New("Controller not defined")
var singletonServiceWorkerContainer sync.Once

var serviceworkercontainerinterface js.Value

// GetInterface get the JS interface of serviceworkercontainer
func GetServiceWorkerContainerInterface() js.Value {
	singletonServiceWorkerContainer.Do(func() {
		if serviceworkercontainerinterface = js.Global().Get("ServiceWorkerContainer"); serviceworkercontainerinterface.Error() != nil {
			serviceworkercontainerinterface = js.Undefined()
		}
		js.Register(serviceworkercontainerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		GetPromiseInterface()
		GetServiceWorkerInterface()

	})

	return serviceworkercontainerinterface
}

type ServiceWorkerContainer struct {
	html.EventTarget
}

type ServiceWorkerContainerFrom interface {
	ServiceWorkerContainer_() ServiceWorkerContainer
}

func (s ServiceWorkerContainer) ServiceWorkerContainer_() ServiceWorkerContainer {
	return s
}

func ToServiceWorkerContainer(obj js.Value) (ServiceWorkerContainer, error) {
	var s ServiceWorkerContainer

	if si := GetServiceWorkerContainerInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			//s.BaseObject = s.SetObject(obj)
			s.SetValue(obj)
			return s, nil

		}
	}

	return s, js.ErrNotImplemented
}

func (s ServiceWorkerContainer) Controller() (ServiceWorker, error) {

	var err error
	var obj interface{}
	var sw ServiceWorker
	var ok bool

	if obj, err = s.GetAttributeGlobal("controller"); err == nil {

		if obj == nil {
			return sw, ErrControllerNotDefined
		} else {
			if sw, ok = obj.(ServiceWorker); !ok {

				err = ErrNotAServiceWorker

			}
		}

	}

	return sw, err

}

func (s ServiceWorkerContainer) Ready() (Promise, error) {
	var err error
	var obj interface{}
	var p Promise
	var ok bool

	if obj, err = s.GetAttributeGlobal("ready"); err == nil {

		if p, ok = obj.(Promise); !ok {
			err = ErrNotAPromise
		}

	}

	return p, err

}

func (s ServiceWorkerContainer) GetRegistration(clientURL string) (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = s.Call("getRegistration", js.ValueOf(clientURL)); obj.Error() == nil {

		p, err = ToPromise(obj)

	}
	return p, err

}

func (s ServiceWorkerContainer) Register(url string, options ...map[string]interface{}) (Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p Promise

	arrayJS = append(arrayJS, js.ValueOf(url))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = s.Call("register", arrayJS...); obj.Error() == nil {

		p, err = ToPromise(obj)

	}

	return p, err
}

func (s ServiceWorkerContainer) OnControllerChange(handler func(e html.Event)) (js.Func, error) {

	return s.AddEventListener("controllerchange", handler)
}

func (s ServiceWorkerContainer) OnMessage(handler func(e html.Event)) (js.Func, error) {

	return s.AddEventListener("message", handler)
}

func (s ServiceWorkerContainer) OnError(handler func(e html.Event)) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
