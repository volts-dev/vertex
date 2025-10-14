package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBFactory

import (
	"sync"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"

	"github.com/volts-dev/vertex/html/promise"
)

var singletonIDBFactory sync.Once

var idbfactoryinterface js.Value

// GetInterface get the JS interface
func GetIDBFactoryInterface() js.Value {

	singletonIDBFactory.Do(func() {

		if idbfactoryinterface = js.Global().Get("IDBFactory"); idbfactoryinterface.Error() != nil {
			idbfactoryinterface = js.Undefined()
		}
	})
	return idbfactoryinterface
}

// IDBFactory struct
type IDBFactory struct {
	js.Object
}

type IDBFactoryFrom interface {
	IDBFactory_() IDBFactory
}

func (i IDBFactory) IDBFactory_() IDBFactory {
	return i
}

func IDBFactoryNewFromJSObject(obj js.Value) (IDBFactory, error) {
	var i IDBFactory
	var err error
	if ai := GetIDBFactoryInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBFactory
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (f IDBFactory) genericRequest(method string, dbname string, option ...string) (IDBOpenDBRequest, error) {
	var err error
	var i IDBOpenDBRequest
	var idbobj js.Value

	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(dbname))

	if len(option) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(option[0]))
	}

	if idbobj = f.Call(method, arrayJS...); idbobj.Error() == nil {
		i, err = IDBOpenDBRequestNewFromJSObject(idbobj)

	}

	return i, err

}

func (f IDBFactory) Cmp(a, b interface{}) (int, error) {
	var arrayJS []interface{}
	var err error
	var obj js.Value
	var result int
	arrayJS = append(arrayJS, js.ValueOf(a), js.ValueOf(b))
	if obj = f.Call("cmp", arrayJS...); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		} else {
			err = object.ErrObjectNotNumber
		}
	}
	return result, err

}

func (f IDBFactory) Open(dbname string, option ...string) (IDBOpenDBRequest, error) {

	return f.genericRequest("open", dbname, option...)
}

func (f IDBFactory) DeleteDatabase(dbname string, option ...string) (IDBOpenDBRequest, error) {
	return f.genericRequest("deleteDatabase", dbname, option...)

}

func (f IDBFactory) Databases() (promise.Promise, error) {
	//not support in firefox
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = f.Call("databases"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
