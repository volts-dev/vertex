package serviceworkerregistration

import (
	"fmt"
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/navigationpreloadmanager"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/pushmanager"
	"github.com/volts-dev/vertex/html/serviceworker"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkerregistrationinterface js.Value

// GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {

		if serviceworkerregistrationinterface = js.Global().Get("ServiceWorkerRegistration"); serviceworkerregistrationinterface.Error() != nil {
			serviceworkerregistrationinterface = js.Undefined()
		}
		js.Register(serviceworkerregistrationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()
		serviceworker.GetInterface()
		navigationpreloadmanager.GetInterface()
		pushmanager.GetInterface()
	})

	return serviceworkerregistrationinterface
}

type ServiceWorkerRegistration struct {
	eventtarget.EventTarget
}

type ServiceWorkerRegistrationFrom interface {
	ServiceWorkerRegistration_() ServiceWorkerRegistration
}

func (s ServiceWorkerRegistration) ServiceWorkerRegistration_() ServiceWorkerRegistration {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorkerRegistration, error) {
	var s ServiceWorkerRegistration

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.SetObjectValue(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorkerRegistration) getserviceworkerAttribute(attribute string) (serviceworker.ServiceWorker, error) {
	var err error
	var obj interface{}
	var sw serviceworker.ServiceWorker
	var ok bool

	if obj, err = s.GetAttributeGlobal(attribute); err == nil {

		if obj == nil {
			return sw, js.ErrUndefinedValue
		} else {
			if sw, ok = obj.(serviceworker.ServiceWorker); !ok {
				fmt.Printf("-->%v\n", obj.(js.Object).ConstructName_())
				err = serviceworker.ErrNotAServiceWorker

			}
		}

	}

	return sw, err
}

func (s ServiceWorkerRegistration) Active() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("active")
}

func (s ServiceWorkerRegistration) Index() (int, error) {

	return s.GetAttributeInt("index")
}

func (s ServiceWorkerRegistration) Installing() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("installing")
}

func (s ServiceWorkerRegistration) Scope() (string, error) {

	return s.GetAttributeString("scope")
}

func (s ServiceWorkerRegistration) Waiting() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("waiting")
}

func (s ServiceWorkerRegistration) NavigationPreload() (navigationpreloadmanager.NavigationPreloadManager, error) {

	var err error
	var obj interface{}
	var n navigationpreloadmanager.NavigationPreloadManager
	var ok bool

	if obj, err = s.GetAttributeGlobal("navigationPreload"); err == nil {
		if n, ok = obj.(navigationpreloadmanager.NavigationPreloadManager); !ok {
			err = navigationpreloadmanager.ErrNotANavigationPreloadManager
		}
	}

	return n, err
}

func (s ServiceWorkerRegistration) PushManager() (pushmanager.PushManager, error) {

	var err error
	var obj interface{}
	var p pushmanager.PushManager
	var ok bool

	if obj, err = s.GetAttributeGlobal("pushManager"); err == nil {
		if p, ok = obj.(pushmanager.PushManager); !ok {
			err = pushmanager.ErrNotAPushManager
		}
	}

	return p, err
}

func (s ServiceWorkerRegistration) GetNotifications(title string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = s.Call("getNotifications", arrayJS...); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err

}

func (s ServiceWorkerRegistration) ShowNotification(title string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = s.Call("shownotification", arrayJS...); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err

}

func (s ServiceWorkerRegistration) Unregister() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = s.Call("getRegistration"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerRegistration) Update() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = s.Call("update"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerRegistration) UpdateViaCache() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = s.Call("updateViaCache"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}
