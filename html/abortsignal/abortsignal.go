package abortsignal

import (
	"sync"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var abortsignalinterface js.Value

// GetInterface get the JS interface abortsignal
func GetInterface() js.Value {

	singleton.Do(func() {

		if abortsignalinterface = js.Global().Get("AbortSignal"); abortsignalinterface.Error() != nil {
			abortsignalinterface = js.Undefined()
		}
		js.Register(abortsignalinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return abortsignalinterface
}

// AbortSignal struct
type AbortSignal struct {
	eventtarget.EventTarget
}

type AbortSignalFrom interface {
	AbortSignal_() AbortSignal
}

func NewFromJSObject(obj js.Value) (AbortSignal, error) {
	var a AbortSignal
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnAbortSignal
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a AbortSignal) Aborted() (bool, error) {

	return a.GetAttributeBool("aborted")

}

func (a AbortSignal) Abort() (AbortSignal, error) {
	var err error
	var obj js.Value
	var as AbortSignal
	if obj = a.Call("abort"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			as, err = NewFromJSObject(obj)
		}
	}
	return as, err
}
