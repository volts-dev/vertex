package htmldataelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`d= document.createElement("data")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLDataElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("d"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLDataElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Value", "resultattempt": ""},
	{"method": "SetValue", "args": []interface{}{"value"}, "gettermethod": "Value", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("d"); test.AssertErr(t, obj.Error()) {

		if button, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, button, result)
			}

		}

	}
}
