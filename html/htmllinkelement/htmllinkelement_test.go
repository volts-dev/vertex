package htmllinkelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`l=document.createElement("link")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLLinkElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLLinkElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "As", "resultattempt": ""},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Media", "resultattempt": ""},
	{"method": "Href", "resultattempt": ""},
	{"method": "Hreflang", "resultattempt": ""},
	{"method": "ReferrerPolicy", "resultattempt": ""},
	{"method": "Rel", "resultattempt": ""},
	{"method": "RelList", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "Sizes", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "Type", "resultattempt": ""},
	{"method": "SetAs", "args": []interface{}{"font"}, "gettermethod": "As", "resultattempt": "font"},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetHref", "args": []interface{}{"myFont.woff2"}, "gettermethod": "Href", "type": "contains", "resultattempt": "/myFont.woff2"},
	{"method": "SetHreflang", "args": []interface{}{"lang"}, "gettermethod": "Hreflang", "resultattempt": "lang"},
	{"method": "SetReferrerPolicy", "args": []interface{}{"no-referrer"}, "gettermethod": "ReferrerPolicy", "resultattempt": "no-referrer"},
	{"method": "SetRel", "args": []interface{}{"stylesheet"}, "gettermethod": "Rel", "resultattempt": "stylesheet"},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if anchor, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, anchor, result)
			}

		}

	}
}
