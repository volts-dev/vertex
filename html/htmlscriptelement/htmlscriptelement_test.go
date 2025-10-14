package htmlscriptelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`s= document.createElement("script")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if a, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLScriptElement", a.ConstructName_())
		}

	}

}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if a, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLScriptElement", a.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Type", "resultattempt": ""},
	{"method": "SetType", "args": []interface{}{"module"}, "gettermethod": "Type", "resultattempt": "module"},
	{"method": "Src", "resultattempt": ""},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "Async", "resultattempt": true},
	{"method": "SetAsync", "args": []interface{}{false}, "gettermethod": "Async", "resultattempt": false},
	{"method": "Defer", "resultattempt": false},
	{"method": "SetDefer", "args": []interface{}{true}, "gettermethod": "Defer", "resultattempt": true},
	{"method": "Text", "resultattempt": ""},
	{"method": "SetText", "args": []interface{}{"test"}, "gettermethod": "Text", "resultattempt": "test"},
	{"method": "NoModule", "resultattempt": false},
	{"method": "SetNoModule", "args": []interface{}{true}, "gettermethod": "NoModule", "resultattempt": true},
	{"method": "ReferrerPolicy", "resultattempt": ""},
	{"method": "SetReferrerPolicy", "args": []interface{}{"no-referrer"}, "gettermethod": "ReferrerPolicy", "resultattempt": "no-referrer"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if area, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, area, result)
			}

		}

	}
}
