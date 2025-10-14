package serviceworkercontainer

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/serviceworker"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkercontainerinterface js.Value

// GetInterface get the JS interface of serviceworkercontainer
func GetInterface() js.Value {

	singleton.Do(func() {

		if serviceworkercontainerinterface = js.Global().Get("ServiceWorkerContainer"); serviceworkercontainerinterface.Error() != nil {
			serviceworkercontainerinterface = js.Undefined()
		}
		js.Register(serviceworkercontainerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()
		serviceworker.GetInterface()

	})

	return serviceworkercontainerinterface
}

type ServiceWorkerContainer struct {
	eventtarget.EventTarget
}

type ServiceWorkerContainerFrom interface {
	ServiceWorkerContainer_() ServiceWorkerContainer
}

func (s ServiceWorkerContainer) ServiceWorkerContainer_() ServiceWorkerContainer {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorkerContainer, error) {
	var s ServiceWorkerContainer

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.SetObjectValue(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorkerContainer) Controller() (serviceworker.ServiceWorker, error) {

	var err error
	var obj interface{}
	var sw serviceworker.ServiceWorker
	var ok bool

	if obj, err = s.GetAttributeGlobal("controller"); err == nil {

		if obj == nil {
			return sw, ErrControllerNotDefined
		} else {
			if sw, ok = obj.(serviceworker.ServiceWorker); !ok {

				err = serviceworker.ErrNotAServiceWorker

			}
		}

	}

	return sw, err

}

func (s ServiceWorkerContainer) Ready() (promise.Promise, error) {
	var err error
	var obj interface{}
	var p promise.Promise
	var ok bool

	if obj, err = s.GetAttributeGlobal("ready"); err == nil {

		if p, ok = obj.(promise.Promise); !ok {
			err = promise.ErrNotAPromise
		}

	}

	return p, err

}

func (s ServiceWorkerContainer) GetRegistration(clientURL string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = s.Call("getRegistration", js.ValueOf(clientURL)); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerContainer) Register(url string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(url))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = s.Call("register", arrayJS...); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err
}
