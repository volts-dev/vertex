package abortcontroller

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/abortsignal"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var abortcontrollerinterface js.Value

// GetInterface get the JS interface abortcontroller
func GetInterface() js.Value {

	singleton.Do(func() {

		if abortcontrollerinterface = js.Global().Get("AbortController"); abortcontrollerinterface.Error() != nil {
			abortcontrollerinterface = js.Undefined()
		}
		js.Register(abortcontrollerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return abortcontrollerinterface
}

// AbortController struct
type AbortController struct {
	js.Object
}

type AbortControllerFrom interface {
	AbortController_() AbortController
}

func (a AbortController) AbortController_() AbortController {
	return a
}

func NewFromJSObject(obj js.Value) (AbortController, error) {
	var a AbortController
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnAbortController
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func New() (AbortController, error) {

	var a AbortController
	var obj js.Value
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj = ai.New(); obj.Error() == nil {
			a.SetObjectValue(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func (a AbortController) Signal() (abortsignal.AbortSignal, error) {
	var err error
	var obj js.Value
	var as abortsignal.AbortSignal
	if obj = a.GetValueByKey("signal"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			as, err = abortsignal.NewFromJSObject(obj)
		}
	}
	return as, err
}

func (a AbortController) Abort() error {
	var err error
	err = a.Call("abort").Error()
	return err
}
