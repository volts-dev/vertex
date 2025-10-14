package indexeddb

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

func TestIDBRequestNewFromJSObject(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	dbrequest=iddb.open("openrequest2")
	`)

	if obj := js.Global().Get("dbrequest"); test.AssertErr(t, obj.Error()) {

		if openrequest, err := IDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "IDBOpenDBRequest", openrequest.ConstructName_())
		}

	}
	js.Eval(`
	dbrequest=iddb.deleteDatabase("openrequest2")
	`)

}

var methodsIDBRequestAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Error", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "ReadyState", "resultattempt": "done"},
	{"method": "Source", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "Transaction", "type": "constructnamechecking", "resultattempt": "IDBTransaction"},
}

func TestIDBRequestMethods(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	dbrequest=iddb.open("openrequestmethods")
	`)

	var io chan bool = make(chan bool)

	if obj := js.Global().Get("dbrequest"); test.AssertErr(t, obj.Error()) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				for _, result := range methodsIDBRequestAttempt {
					test.InvokeCheck(t, openrequest, result)
				}

				io <- true
			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	js.Eval(`
	dbrequest=iddb.deleteDatabase("openrequestmethods")
	`)
}
