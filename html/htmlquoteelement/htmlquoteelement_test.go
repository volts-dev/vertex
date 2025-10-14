package htmlquoteelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`
	q= document.createElement("blockquote")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLQuoteElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("q"); test.AssertErr(t, obj.Error()) {

		if meter, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLQuoteElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Cite", "resultattempt": ""},
	{"method": "SetCite", "args": []interface{}{"test"}, "gettermethod": "Cite", "type": "contains", "resultattempt": "/test"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("q"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
