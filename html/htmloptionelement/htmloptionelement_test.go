package htmloptionelement

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
	f=document.createElement("form")
	s= document.createElement("select")
	o= document.createElement("option")
	s.appendChild(o)
	f.appendChild(s)
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLOptionElement", b.ConstructName_())
		}

	}
}

func TestOption(t *testing.T) {

	if o, err := Option("test"); test.AssertErr(t, err) {
		test.AssertExpect(t, "HTMLOptionElement", o.ConstructName_())
	}

}
func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("o"); test.AssertErr(t, obj.Error()) {

		if meter, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLOptionElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "DefaultSelected", "resultattempt": false},
	{"method": "SetDefaultSelected", "args": []interface{}{true}, "gettermethod": "DefaultSelected", "resultattempt": true},
	{"method": "Disabled", "resultattempt": false},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "Index", "resultattempt": 0},
	{"method": "Label", "resultattempt": ""},
	{"method": "SetLabel", "args": []interface{}{"test"}, "gettermethod": "Label", "resultattempt": "test"},
	{"method": "Selected", "resultattempt": true},
	{"method": "SetSelected", "args": []interface{}{false}, "gettermethod": "Selected", "resultattempt": false},
	{"method": "Text", "resultattempt": ""},
	{"method": "SetText", "args": []interface{}{"test"}, "gettermethod": "Text", "resultattempt": "test"},
	{"method": "Value", "resultattempt": "test"},
	{"method": "SetValue", "args": []interface{}{"t3st"}, "gettermethod": "Value", "resultattempt": "t3st"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("o"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
