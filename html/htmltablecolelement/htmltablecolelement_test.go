package htmltablecolelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`col= document.createElement("col")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if col, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableColElement", col.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("col"); test.AssertErr(t, obj.Error()) {

		if col, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableColElement", col.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Span", "resultattempt": 1},
	{"method": "SetSpan", "args": []interface{}{10}, "gettermethod": "Span", "resultattempt": 10},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("col"); test.AssertErr(t, obj.Error()) {

		if selectobj, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, selectobj, result)
			}

		}

	}
}
