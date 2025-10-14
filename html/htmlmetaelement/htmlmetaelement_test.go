package htmlmetaelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`m= document.createElement("meta")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if meta, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLMetaElement", meta.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("m"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLMetaElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Content", "resultattempt": ""},
	{"method": "SetContent", "args": []interface{}{"test"}, "gettermethod": "Content", "resultattempt": "test"},
	{"method": "HttpEquiv", "resultattempt": ""},
	{"method": "SetHttpEquiv", "args": []interface{}{"test"}, "gettermethod": "HttpEquiv", "resultattempt": "test"},
	{"method": "Name", "resultattempt": ""},
	{"method": "SetName", "args": []interface{}{"test"}, "gettermethod": "Name", "resultattempt": "test"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("m"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
