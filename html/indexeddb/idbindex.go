package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBIndex

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
)

func init() {

	initinterface.RegisterInterface(GetIDBIndexInterface)
	initinterface.RegisterInterface(GetIDBFactoryInterface)
	initinterface.RegisterInterface(GetIDBKeyRangeInterface)
	initinterface.RegisterInterface(IDBOpenDBRequestGetInterface)
	initinterface.RegisterInterface(IDBDatabaseGetInterface)
	initinterface.RegisterInterface(IDBCursorGetInterface)
	initinterface.RegisterInterface(IDBDatabaseGetInterface)
	initinterface.RegisterInterface(IDBCursorWithValueGetInterface)
	initinterface.RegisterInterface(IDBObjectStoreGetInterface)
	initinterface.RegisterInterface(IDBRequestGetInterface)
	initinterface.RegisterInterface(IDBTransactionGetInterface)

}

var singletonIDBIndex sync.Once

var idbindexinterface js.Value

// GetIDBIndexInterface get the JS interface
func GetIDBIndexInterface() js.Value {

	singletonIDBIndex.Do(func() {

		if idbindexinterface = js.Global().Get("IDBIndex"); idbindexinterface.Error() != nil {
			idbindexinterface = js.Undefined()
		}
		js.Register(idbindexinterface, func(v js.Value) (interface{}, error) {
			return IDBDIndexNewFromJSObject(v)
		})
	})
	return idbindexinterface
}

// IDBIndex struct
type IDBIndex struct {
	js.Object
}

type IDBIndexFrom interface {
	IDBIndex_() IDBIndex
}

func (i IDBIndex) IDBIndex_() IDBIndex {
	return i
}

func IDBDIndexNewFromJSObject(obj js.Value) (IDBIndex, error) {
	var i IDBIndex
	var err error
	if ai := GetIDBIndexInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBIndex
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBIndex) KeyPath() (string, error) {
	return i.GetAttributeString("keyPath")
}

func (i IDBIndex) Name() (string, error) {
	return i.GetAttributeString("name")
}

func (i IDBIndex) MultiEntry() (bool, error) {
	return i.GetAttributeBool("multiEntry")
}

func (i IDBIndex) ObjectStore() (IDBObjectStore, error) {

	var err error
	var obj js.Value
	var store IDBObjectStore

	if obj = i.GetValueByKey("objectStore"); obj.Error() == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = object.ErrNotAnObject

		} else {

			store, err = IDBObjectStoreNewFromJSObject(obj)
		}
	}

	return store, err
}

func (i IDBIndex) Unique() (bool, error) {
	return i.GetAttributeBool("unique")
}

func (i IDBIndex) callMethodKey(method string, key ...interface{}) (IDBRequest, error) {
	var obj js.Value
	var o IDBRequest
	var err error
	var arrayJS []interface{}

	if len(key) > 0 {

		var objkey interface{}

		if o, ok := key[0].(IDBKeyRange); ok {
			objkey = o.GetObjectValue()
		} else {
			objkey = js.ValueOf(key[0])
		}

		arrayJS = append(arrayJS, objkey)
	}

	if obj = i.Call(method, arrayJS...); obj.Error() == nil {
		o, err = IDBRequestNewFromJSObject(obj)
	}

	return o, err
}

func (i IDBIndex) Count(key ...interface{}) (IDBRequest, error) {
	return i.callMethodKey("count", key...)
}

func (i IDBIndex) Get(key interface{}) (IDBRequest, error) {
	return i.callMethodKey("get", key)
}

func (i IDBIndex) GetKey(key interface{}) (IDBRequest, error) {

	return i.callMethodKey("getKey", key)
}

func (i IDBIndex) getAll(method string, option ...interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objquery js.Value
	var arrayJS []interface{}

	if len(option) > 1 {
		if rangequery, ok := option[0].(IDBKeyRange); ok {
			objquery = rangequery.GetObjectValue()
		} else {
			objquery = js.ValueOf(option[0])
		}
		arrayJS = append(arrayJS, objquery)
	}
	if len(option) > 2 {
		if count, ok := option[0].(int); ok {
			arrayJS = append(arrayJS, js.ValueOf(count))
		}

	}

	if obj = i.Call(method, arrayJS...); obj.Error() == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBIndex) GetAll(option ...interface{}) (IDBRequest, error) {
	return i.getAll("getAll", option...)
}

func (i IDBIndex) GetAllKeys(option ...interface{}) (IDBRequest, error) {
	return i.getAll("getAllKeys", option...)
}

func (i IDBIndex) openCursorWithMethod(method string, options ...interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objquery js.Value
	var arrayJS []interface{}

	if len(options) > 1 {
		if rangequery, ok := options[0].(IDBKeyRange); ok {
			objquery = rangequery.GetObjectValue()
			arrayJS = append(arrayJS, objquery)
		}

	}

	if len(options) > 2 {
		if direction, ok := options[1].(string); ok {
			objquery = js.ValueOf(direction)
			arrayJS = append(arrayJS, objquery)
		}

	}

	if obj = i.Call("openCursor", arrayJS...); obj.Error() == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBIndex) OpenCursor(options ...interface{}) (IDBRequest, error) {

	return i.openCursorWithMethod("openCursor", options...)
}

func (i IDBIndex) OpenKeyCursor(options ...interface{}) (IDBRequest, error) {

	return i.openCursorWithMethod("openKeyCursor", options...)
}
