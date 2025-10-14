package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBRequest

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/domexception"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/eventtarget"
)

var classIDBRequest string = "IDBRequest"

// IDBRequest struct
type IDBRequest struct {
	eventtarget.EventTarget
}

type IDBRequestFrom interface {
	IDBRequest_() IDBRequest
}

func (i IDBRequest) IDBRequest_() IDBRequest {
	return i
}

var singletonIDBRequest sync.Once

var idbrequestinterface js.Value

func IDBRequestGetInterface() js.Value {

	singletonIDBRequest.Do(func() {

		if idbrequestinterface = js.Global().Get(classIDBRequest); idbrequestinterface.Error() != nil {
			idbrequestinterface = js.Undefined()
		}

		js.Register(idbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBRequestNewFromJSObject(v)
		})
		IDBTransactionGetInterface()
		IDBDatabaseGetInterface()
		IDBCursorGetInterface()
	})
	return idbrequestinterface
}

func IDBRequestNewFromJSObject(obj js.Value) (IDBRequest, error) {
	var i IDBRequest
	var err error
	if ai := IDBRequestGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBRequest) OnError(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("error", handler)
}

func (i IDBRequest) OnSuccess(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("success", handler)
}

func (i IDBRequest) Error() (domexception.DomException, error) {
	var err error
	var obj js.Value
	var e domexception.DomException
	if obj = i.GetValueByKey("error"); obj.Error() == nil {
		e, err = domexception.NewFromJSObject(obj)
	}
	return e, err
}

func (i IDBRequest) ReadyState() (string, error) {
	return i.GetAttributeString("readyState")
}

func (i IDBRequest) Result() (interface{}, error) {

	var err error
	var obj js.Value
	var ret interface{}

	if obj = i.GetValueByKey("result"); obj.Error() == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			ret, err = js.Discover(obj)
		}
	}

	return ret, err
}

func (i IDBRequest) Source() (interface{}, error) {

	var err error
	var obj js.Value
	var ret interface{}

	if obj = i.GetValueByKey("source"); obj.Error() == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			ret, err = js.Discover(obj)
		}

	}

	return ret, err

}

func (i IDBRequest) Transaction() (IDBTransaction, error) {

	var err error
	var obj js.Value
	var it IDBTransaction
	var ret interface{}

	if obj = i.GetValueByKey("transaction"); obj.Error() == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if ret, err = js.Discover(obj); err == nil {

				if tfrom, ok := ret.(IDBTransactionFrom); ok {

					it = tfrom.IDBTransaction_()

				} else {
					err = ErrNotAnIDBTransaction
				}

			}
		}
	}
	return it, err
}
