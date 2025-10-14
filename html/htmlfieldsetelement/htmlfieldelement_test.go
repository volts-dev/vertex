package htmlfieldsetelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`f= document.createElement("fieldset")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLFieldSetElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("f"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLFieldSetElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Disabled", "resultattempt": false},
	{"method": "Elements", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Form", "type": "error", "resultattempt": ErrNoForm},
	{"method": "Name", "resultattempt": ""},
	{"method": "Type", "resultattempt": "fieldset"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
	{"method": "WillValidate", "resultattempt": false},
	{"method": "ReportValidity", "resultattempt": true},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "SetCustomValidity", "args": []interface{}{"hello"}, "type": "error", "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("f"); test.AssertErr(t, obj.Error()) {

		if button, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, button, result)
			}

		}

	}
}
