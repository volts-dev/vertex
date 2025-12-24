package serviceworker

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkerinterface js.Value

// GetInterface get the JS interface of serviceworker
func GetInterface() js.Value {

	singleton.Do(func() {

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
	eventtarget.EventTarget
}

type ServiceWorkerFrom interface {
	ServiceWorker_() ServiceWorker
}

func (s ServiceWorker) ServiceWorker_() ServiceWorker {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorker, error) {
	var s ServiceWorker

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.SetObjectValue(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorker) ScriptURL() (string, error) {

	return s.GetAttributeString("scriptURL")
}

func (s ServiceWorker) State() (string, error) {

	return s.GetAttributeString("state")
}
