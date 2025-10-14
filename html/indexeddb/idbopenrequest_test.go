package indexeddb

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"
)

func TestIDBOpenDBRequestNewFromJSObject(t *testing.T) {
	js.Eval(`iddb=window.indexedDB
	opendbrequest=iddb.open("openrequest")
	`)

	if obj := js.Global().Get("opendbrequest"); test.AssertErr(t, obj.Error()) {

		if opendbrequest, err := IDBOpenDBRequestNewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "IDBOpenDBRequest", opendbrequest.ConstructName_())
		}

	}
	js.Eval(`
	opendbrequest=iddb.deleteDatabase("openrequest")
	`)

}
