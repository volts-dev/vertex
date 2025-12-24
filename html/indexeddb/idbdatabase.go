package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBDatabase

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/domstringlist"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/eventtarget"
)

// IDBDatabase struct
type IDBDatabase struct {
	eventtarget.EventTarget
}

type IDBDatabaseFrom interface {
	IDBDatabase_() IDBDatabase
}

func (i IDBDatabase) IDBDatabase_() IDBDatabase {
	return i
}

var singletonIDBDatabase sync.Once

var idbdatabaseinterface js.Value

func IDBDatabaseGetInterface() js.Value {

	singletonIDBDatabase.Do(func() {

		if idbdatabaseinterface = js.Global().Get("IDBDatabase"); idbdatabaseinterface.Error() != nil {
			idbdatabaseinterface = js.Undefined()
		}

		js.Register(idbdatabaseinterface, func(v js.Value) (interface{}, error) {

			return IDBDatabaseNewFromJSObject(v)
		})
	})

	return idbdatabaseinterface
}

func IDBDatabaseNewFromJSObject(obj js.Value) (IDBDatabase, error) {
	var i IDBDatabase
	var err error
	if ai := IDBDatabaseGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBDatabase
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBDatabase) Close() error {
	var err error
	err = i.Call("close").Error()
	return err
}

func (i IDBDatabase) DeleteObjectStore(name string) error {
	var err error
	err = i.Call("deleteObjectStore", js.ValueOf(name)).Error()
	return err
}

func (i IDBDatabase) CreateObjectStore(name string, options ...map[string]interface{}) (IDBObjectStore, error) {
	var err error
	var obj js.Value
	var arrayJS []interface{}
	var s IDBObjectStore
	arrayJS = append(arrayJS, js.ValueOf(name))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}
	if obj = i.Call("createObjectStore", arrayJS...); obj.Error() == nil {
		s, err = IDBObjectStoreNewFromJSObject(obj)
	}

	return s, err
}

func (i IDBDatabase) Transaction(store interface{}, mode ...string) (IDBTransaction, error) {
	var err error
	var obj js.Value
	var arrayJS []interface{}
	var t IDBTransaction

	//array of string ['my-store-name']
	if arr, ok := store.(js.Array); ok {
		arrayJS = append(arrayJS, arr.GetObjectValue())
		//store name
	} else if storename, ok := store.(string); ok {
		arrayJS = append(arrayJS, js.ValueOf(storename))
	} else {
		err = ErrBadStoreType
	}

	if len(mode) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(mode[0]))
	}

	if obj = i.Call("transaction", arrayJS...); obj.Error() == nil {
		t, err = IDBTransactionNewFromJSObject(obj)
	}
	return t, err
}

func (i IDBDatabase) getAttributeInt(attribute string) (int64, error) {

	var err error
	var obj js.Value
	var ret int64

	if obj = i.GetValueByKey(attribute); obj.Error() == nil {

		if obj.Type() == js.TypeNumber {
			v, _ := obj.Float()
			ret = int64(v)
		} else {
			err = js.ErrObjectNotNumber
		}
	}
	return ret, err
}

func (i IDBDatabase) Name() (string, error) {
	return i.GetAttributeString("name")
}

func (i IDBDatabase) Version() (int64, error) {
	return i.getAttributeInt("version")
}

func (i IDBDatabase) ObjectStoreNames() (domstringlist.DOMStringList, error) {

	var err error
	var obj js.Value
	var d domstringlist.DOMStringList

	if obj = i.GetValueByKey("objectStoreNames"); obj.Error() == nil {
		d, err = domstringlist.NewFromJSObject(obj)
	}
	return d, err
}

func (i IDBDatabase) OnAbort(handler func(e event.Event) error) (js.Func, error) {

	return i.AddEventListener("onabort", handler)
}

func (i IDBDatabase) OnError(handler func(e event.Event) error) (js.Func, error) {

	return i.AddEventListener("onerror", handler)
}

func (i IDBDatabase) OnVersionChange(handler func(e event.Event) error) (js.Func, error) {

	return i.AddEventListener("onversionchange", handler)
}
