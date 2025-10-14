package indexeddb

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

var methodsIDBIndexeAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "KeyPath", "resultattempt": "hello"},
	{"method": "Name", "resultattempt": "hello"},
	{"method": "MultiEntry", "resultattempt": false},
	{"method": "ObjectStore", "type": "constructnamechecking", "resultattempt": "IDBObjectStore"},
	{"method": "Unique", "resultattempt": false},
	{"method": "Count", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Get", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetKey", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAll", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAllKeys", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenKeyCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
}

func TestIDBIndexMethods(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	objectstore=iddb.open("idbindex")
	`)

	var io chan bool = make(chan bool)

	if obj := js.Global().Get("objectstore"); test.AssertErr(t, obj.Error()) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); test.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("user", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							if index, err := store.CreateIndex("hello", "hello"); test.AssertErr(t, err) {

								for _, result := range methodsIDBIndexeAttempt {
									test.InvokeCheck(t, index, result)
								}

								io <- true

							}
						}

					}
				}

			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	js.Eval(`
	iddb.deleteDatabase("idbindex")
	`)
}
