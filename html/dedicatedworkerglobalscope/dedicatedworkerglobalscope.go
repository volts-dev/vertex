package dedicatedworkerglobalscope

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/messageevent"
	"github.com/volts-dev/vertex/html/workerglobalscope"
	"github.com/volts-dev/vertex/js/array"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var dedicatedworkerglobalscopeinterface js.Value

// GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {

		if dedicatedworkerglobalscopeinterface = js.Global().Get("DedicatedWorkerGlobalScope"); dedicatedworkerglobalscopeinterface.Error() != nil {
			dedicatedworkerglobalscopeinterface = js.Undefined()
		}
		js.Register(dedicatedworkerglobalscopeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		messageevent.GetInterface()

	})

	return dedicatedworkerglobalscopeinterface
}

type DedicatedWorkerGlobalScope struct {
	workerglobalscope.WorkerGlobalScope
}

type DedicatedWorkerGlobalScopeFrom interface {
	DedicatedWorkerGlobalScope_() DedicatedWorkerGlobalScope
}

func (d DedicatedWorkerGlobalScope) DedicatedWorkerGlobalScope_() DedicatedWorkerGlobalScope {
	return d
}

func NewFromJSObject(obj js.Value) (DedicatedWorkerGlobalScope, error) {
	var d DedicatedWorkerGlobalScope

	if di := GetInterface(); !di.IsUndefined() {
		if obj.InstanceOf(di) {
			d.SetObjectValue(obj)
			return d, nil

		}
	}

	return d, ErrNotImplemented
}

func (d DedicatedWorkerGlobalScope) PostMessage(message string, transfer ...array.Array) error {

	var arrayJS []interface{}

	var err error

	arrayJS = append(arrayJS, js.ValueOf(message))

	if len(transfer) > 0 {
		arrayJS = append(arrayJS, transfer[0].GetObjectValue())
	}

	err = d.Call("postMessage", arrayJS...).Error()

	return err

}

func (d DedicatedWorkerGlobalScope) Name() (string, error) {

	return d.GetAttributeString("name")
}

func (d DedicatedWorkerGlobalScope) Close() error {

	var err error
	err = d.Call("close").Error()

	return err

}
