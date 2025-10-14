package htmlbaseelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`b= document.createElement("base")
	b.href="https://myuser:mypass@www.test.com:444?q=123#tag"
	b.target="thistarget"
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLBaseElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("b"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLBaseElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Href", "resultattempt": "https://myuser:mypass@www.test.com:444/?q=123#tag"},
	{"method": "Target", "resultattempt": "thistarget"},
	{"method": "SetHref", "args": []interface{}{"http://pp:ss@www.noone.com:444?q=456#nosecure"}, "gettermethod": "Href", "resultattempt": "http://pp:ss@www.noone.com:444/?q=456#nosecure"},
	{"method": "SetTarget", "args": []interface{}{"yes"}, "gettermethod": "Target", "resultattempt": "yes"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("b"); test.AssertErr(t, obj.Error()) {

		if base, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, base, result)
			}

		}

	}
}
