package htmliframeelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`i= document.createElement("iframe")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLIFrameElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("i"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLIFrameElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AllowPaymentRequest", "resultattempt": false},
	{"method": "ContentDocument", "type": "error", "resultattempt": ErrNoContentDocument},
	{"method": "Height", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "Name", "resultattempt": ""},
	{"method": "Width", "resultattempt": ""},
	{"method": "Srcdoc", "resultattempt": ""},
	{"method": "SetAllowPaymentRequest", "args": []interface{}{true}, "gettermethod": "AllowPaymentRequest", "resultattempt": true},
	{"method": "SetHeight", "args": []interface{}{"value"}, "gettermethod": "Height", "resultattempt": "value"},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "SetName", "args": []interface{}{"value"}, "gettermethod": "Name", "resultattempt": "value"},
	{"method": "SetWidth", "args": []interface{}{"value"}, "gettermethod": "Width", "resultattempt": "value"},
	{"method": "SetSrcdoc", "args": []interface{}{"value"}, "gettermethod": "Srcdoc", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("i"); test.AssertErr(t, obj.Error()) {

		if iframe, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, iframe, result)
			}

		}

	}
}
