package workerglobalscope

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var workerglobalscopeinterface js.Value

// GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {

		if workerglobalscopeinterface = js.Global().Get("WorkerGlobalScope"); workerglobalscopeinterface.Error() != nil {
			workerglobalscopeinterface = js.Undefined()
		}
		js.Register(workerglobalscopeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return workerglobalscopeinterface
}

type WorkerGlobalScope struct {
	eventtarget.EventTarget
}

type WorkerGlobalScopeFrom interface {
	WorkerGlobalScope_() WorkerGlobalScope
}

func (w WorkerGlobalScope) WorkerGlobalScope_() WorkerGlobalScope {
	return w
}

func NewFromJSObject(obj js.Value) (WorkerGlobalScope, error) {
	var w WorkerGlobalScope

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.SetObjectValue(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func (w WorkerGlobalScope) ImportsScripts(values ...string) error {

	var err error
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	err = w.Call("importScripts", arrayJS...).Error()

	return err
}
