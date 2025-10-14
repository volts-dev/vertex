package indexeddb

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

// IDBCursorWithValue struct
type IDBCursorWithValue struct {
	IDBCursor
}

type IDBCursorWithValueFrom interface {
	IDBCursorWithValue_() IDBCursorWithValue
}

func (i IDBCursorWithValue) IDBCursorWithValue_() IDBCursorWithValue {
	return i
}

var singletonIDBCursorWithValue sync.Once

var idbcursorinterfacewithvalue js.Value

func IDBCursorWithValueGetInterface() js.Value {

	singletonIDBCursorWithValue.Do(func() {

		if idbcursorinterfacewithvalue = js.Global().Get("IDBCursorWithValue"); idbcursorinterfacewithvalue.Error() != nil {
			idbcursorinterfacewithvalue = js.Undefined()
		}

		js.Register(idbcursorinterfacewithvalue, func(v js.Value) (interface{}, error) {
			return IDBCursorWithValueNewFromJSObject(v)
		})
	})

	return idbcursorinterfacewithvalue
}

func IDBCursorWithValueNewFromJSObject(obj js.Value) (IDBCursorWithValue, error) {
	var i IDBCursorWithValue
	var err error
	if ai := IDBCursorWithValueGetInterface(); !ai.IsUndefined() {
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
