package htmltimeelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`t= document.createElement("time")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if ti, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTimeElement", ti.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("t"); test.AssertErr(t, obj.Error()) {

		if ti, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTimeElement", ti.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "DateTime", "resultattempt": ""},
	{"method": "SetDateTime", "args": []interface{}{"13h"}, "gettermethod": "DateTime", "resultattempt": "13h"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("t"); test.AssertErr(t, obj.Error()) {

		if selectobj, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, selectobj, result)
			}

		}

	}
}
