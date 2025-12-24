package worker

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/messageevent"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var workerinterface js.Value

// GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {

		if workerinterface = js.Global().Get("Worker"); workerinterface.Error() != nil {
			workerinterface = js.Undefined()
		}
		js.Register(workerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		messageevent.GetInterface()

	})

	return workerinterface
}

type Worker struct {
	eventtarget.EventTarget
}

type WorkerFrom interface {
	Worker_() Worker
}

func (w Worker) Worker_() Worker {
	return w
}

func NewFromJSObject(obj js.Value) (Worker, error) {
	var w Worker

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.SetObjectValue(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func New(url string, opts ...map[string]interface{}) (Worker, error) {

	var arrayJS []interface{}
	var w Worker
	var err error
	var obj js.Value

	arrayJS = append(arrayJS, js.ValueOf(url))

	if len(opts) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(opts[0]))
	}

	if workeri := GetInterface(); !workeri.IsUndefined() {

		if obj = workeri.New(arrayJS...); obj.Error() == nil {
			w.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return w, err
}

func (w Worker) PostMessage(message string, transfer ...js.Array) error {

	var arrayJS []interface{}

	var err error

	arrayJS = append(arrayJS, js.ValueOf(message))

	if len(transfer) > 0 {
		arrayJS = append(arrayJS, transfer[0].GetObjectValue())
	}

	err = w.Call("postMessage", arrayJS...).Error()

	return err

}

func (w Worker) Terminate() error {

	var err error
	err = w.Call("terminate").Error()

	return err

}
