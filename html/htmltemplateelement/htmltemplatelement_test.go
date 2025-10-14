package htmltemplateelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`t= document.createElement("template")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if template, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTemplateElement", template.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("t"); test.AssertErr(t, obj.Error()) {

		if template, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTemplateElement", template.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Content", "type": "constructnamechecking", "resultattempt": "DocumentFragment"},
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
