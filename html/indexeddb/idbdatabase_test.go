package indexeddb

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

// cant direct call but discover do it for us
func TestIDBDatabaseNewFromJSObject(t *testing.T) {

	var io chan bool = make(chan bool)

	if obj := js.Global().Get("iddb"); test.AssertErr(t, obj.Error()) {

		if factory, err := IDBFactoryNewFromJSObject(obj); test.AssertErr(t, err) {

			if openrequest, err := factory.Open("db", "1"); test.AssertErr(t, err) {

				openrequest.OnSuccess(func(e event.Event) {

					if result, err := openrequest.Result(); err == nil {

						if db, ok := result.(IDBDatabaseFrom); test.AssertExpect(t, true, ok) {

							test.AssertExpect(t, "[object IDBDatabase]", db.IDBDatabase_().ToString_())
							io <- true

						}
					}

				})

			}

			factory.DeleteDatabase("test")
		}

	}

	select {
	case <-io:
	case <-time.After(time.Duration(500) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

var methodsIDBDatabaseAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": "databasemethods"},
	{"method": "Version", "resultattempt": int64(1)},
	{"method": "CreateObjectStore", "args": []interface{}{"utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}}, "type": "constructnamechecking", "resultattempt": "IDBObjectStore"},
	{"method": "ObjectStoreNames", "type": "constructnamechecking", "resultattempt": "DOMStringList"},
	{"method": "DeleteObjectStore", "args": []interface{}{"utilisateur"}, "type": "error", "resultattempt": nil},
	{"method": "Close", "type": "error", "resultattempt": nil},
}

func TestIDBDatabaseMethods(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	dbdatabase=iddb.open("databasemethods")
	`)

	var io chan bool = make(chan bool)

	if obj := js.Global().Get("dbdatabase"); test.AssertErr(t, obj.Error()) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); test.AssertExpect(t, true, ok) {

						for _, result := range methodsIDBDatabaseAttempt {
							test.InvokeCheck(t, db.IDBDatabase_(), result)
						}

						io <- true

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
	iddb.deleteDatabase("dbdatabase")
	`)
}
