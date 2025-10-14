package htmlembedelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`e= document.createElement("embed")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLEmbedElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("e"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLEmbedElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Height", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "Type", "resultattempt": ""},
	{"method": "Width", "resultattempt": ""},
	{"method": "SetHeight", "args": []interface{}{"value"}, "gettermethod": "Height", "resultattempt": "value"},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "SetType", "args": []interface{}{"value"}, "gettermethod": "Type", "resultattempt": "value"},
	{"method": "SetWidth", "args": []interface{}{"value"}, "gettermethod": "Width", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("e"); test.AssertErr(t, obj.Error()) {

		if button, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, button, result)
			}

		}

	}
}
