package htmllabelelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`l= document.createElement("label")
	i= document.createElement("input")
	l.appendChild(i)
	l.appendChild(document.createElement("input"))
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLLabelElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLLabelElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Control", "type": "constructnamechecking", "resultattempt": "HTMLInputElement"},
	{"method": "Form", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "HtmlFor", "resultattempt": ""},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if button, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, button, result)
			}

		}

	}
}
