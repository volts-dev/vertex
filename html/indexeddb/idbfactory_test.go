package indexeddb

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"
)

func TestIDBFactoryNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("iddb"); test.AssertErr(t, obj.Error()) {

		if factory, err := IDBFactoryNewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "IDBFactory", factory.ConstructName_())
		}

	}

}

var methodsFactoryAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Open", "args": []interface{}{"dbtest"}, "type": "constructnamechecking", "resultattempt": "IDBOpenDBRequest"},
	{"method": "DeleteDatabase", "args": []interface{}{"dbtest"}, "type": "constructnamechecking", "resultattempt": "IDBOpenDBRequest"},
	{"method": "Databases", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Cmp", "args": []interface{}{2, 2}, "resultattempt": 0},
}

func TestFactoryMethods(t *testing.T) {

	if obj := js.Global().Get("iddb"); test.AssertErr(t, obj.Error()) {

		if factory, err := IDBFactoryNewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsFactoryAttempt {
				test.InvokeCheck(t, factory, result)
			}

		}

	}
}
