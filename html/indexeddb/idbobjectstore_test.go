package indexeddb

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

var methodsIDBObjectStoreAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "CreateIndex", "args": []interface{}{"email", "emailkey", map[string]interface{}{"unique": true}}, "type": "constructnamechecking", "resultattempt": "IDBIndex"},
	{"method": "Add", "args": []interface{}{map[string]interface{}{"email": "yesmymaail", "data": "name"}}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Get", "args": []interface{}{"yesmymaail"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAll", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAllKeys", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetKey", "args": []interface{}{"email"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Index", "args": []interface{}{"email"}, "type": "constructnamechecking", "resultattempt": "IDBIndex"},
	{"method": "Put", "args": []interface{}{map[string]interface{}{"email": "yesm2", "data": "name"}}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenKeyCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},

	{"method": "Delete", "args": []interface{}{"yesmymaail"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Count", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Clear", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "DeleteIndex", "args": []interface{}{"email"}, "type": "error", "resultattempt": nil},
}

func TestIDBObjectStoreMethods(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	objectstore=iddb.open("objectstore")
	`)

	var io chan bool = make(chan bool)

	if obj := js.Global().Get("objectstore"); test.AssertErr(t, obj.Error()) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) error {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); test.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							test.AssertExpect(t, "[object IDBObjectStore]", store.ToString_())

							for _, result := range methodsIDBObjectStoreAttempt {
								test.InvokeCheck(t, store, result)
							}

							io <- true

						}

					}
				}
				return nil

			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	js.Eval(`
	iddb.deleteDatabase("objectstore")
	`)
}
