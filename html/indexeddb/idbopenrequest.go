package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBOpenDBRequest

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

// IDBOpenRequest struct
type IDBOpenDBRequest struct {
	IDBRequest
}

type IDBOpenDBRequestFrom interface {
	IDBOpenDBRequest_() IDBOpenDBRequest
}

func (i IDBOpenDBRequest) IDBOpenDBRequest_() IDBOpenDBRequest {
	return i
}

var singletonIDBOpenRequest sync.Once

var idbopendbrequestinterface js.Value

func IDBOpenDBRequestGetInterface() js.Value {

	singletonIDBOpenRequest.Do(func() {

		if idbopendbrequestinterface = js.Global().Get("IDBOpenDBRequest"); idbopendbrequestinterface.Error() != nil {
			idbopendbrequestinterface = js.Undefined()
		}

		js.Register(idbopendbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBOpenDBRequestNewFromJSObject(v)
		})
		IDBRequestGetInterface()
	})
	return idbopendbrequestinterface
}

func IDBOpenDBRequestNewFromJSObject(obj js.Value) (IDBOpenDBRequest, error) {
	var i IDBOpenDBRequest
	var err error
	if ai := IDBOpenDBRequestGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBOpenDBRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBOpenDBRequest) OnBlocked(handler func(e event.Event) error) (js.Func, error) {

	return i.AddEventListener("blocked", handler)
}

func (i IDBOpenDBRequest) OnUpgradeNeeded(handler func(e event.Event) error) (js.Func, error) {

	return i.AddEventListener("upgradeneeded", handler)
}
