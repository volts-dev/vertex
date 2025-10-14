package htmlstyleelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`style=document.createElement("style")
	style.type="text/css"
	style.textContent="p { color: #26b72b; }"
	document.head.appendChild(style)
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if source, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLStyleElement", source.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {

		if source, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLStyleElement", source.ConstructName_())
		}
	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Sheet", "type": "constructnamechecking", "resultattempt": "CSSStyleSheet"},
	{"method": "Media", "resultattempt": ""},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Type", "resultattempt": "text/css"},
	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {

		if anchor, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, anchor, result)
			}

		}

	}
}
