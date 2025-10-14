package indexeddb

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
)

// IDBCursor struct
type IDBCursor struct {
	eventtarget.EventTarget
}

type IDBCursorFrom interface {
	IDBCursor_() IDBCursor
}

func (i IDBCursor) IDBCursor_() IDBCursor {
	return i
}

var singletonIDBCursor sync.Once

var idbcursorinterface js.Value

func IDBCursorGetInterface() js.Value {

	singletonIDBCursor.Do(func() {

		if idbcursorinterface = js.Global().Get("IDBCursor"); idbcursorinterface.Error() != nil {
			idbcursorinterface = js.Undefined()
		}

		js.Register(idbcursorinterface, func(v js.Value) (interface{}, error) {
			return IDBCursorNewFromJSObject(v)
		})
		IDBCursorWithValueGetInterface()
	})

	return idbcursorinterface
}

func IDBCursorNewFromJSObject(obj js.Value) (IDBCursor, error) {
	var i IDBCursor
	var err error
	if ai := IDBDatabaseGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBCursor
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBCursor) Direction() (string, error) {
	return i.GetAttributeString("direction")
}

func (i IDBCursor) Key() (interface{}, error) {
	return i.GetAttributeGlobal("key")
}

func (i IDBCursor) PrimaryKey() (interface{}, error) {
	return i.GetAttributeGlobal("primaryKey")
}

func (i IDBCursor) Advance(count int) error {
	var err error
	err = i.Call("advance", js.ValueOf(count)).Error()
	return err
}

func (i IDBCursor) Request() (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest

	if obj = i.GetValueByKey("request"); obj.Error() == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}

func (i IDBCursor) Source() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj = i.GetValueByKey("source"); obj.Error() == nil {

		bobj, err = js.Discover(obj)
	}
	return bobj, err
}

func (i IDBCursor) Continue(option ...interface{}) error {
	//var err error
	var arrayJS []interface{}

	if len(option) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(option[0]))
	}

	i.Call("continue", arrayJS...)

	return nil
}

func (i IDBCursor) Delete() (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest

	if obj = i.Call("delete"); obj.Error() == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}

func (i IDBCursor) Update(value interface{}) (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(value))
	if obj = i.Call("update", arrayJS...); obj.Error() == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}
