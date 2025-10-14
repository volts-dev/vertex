package htmlsourceelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`s=document.createElement("source")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if source, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLSourceElement", source.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if source, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLSourceElement", source.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Media", "resultattempt": ""},
	{"method": "Sizes", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "SrcSet", "resultattempt": ""},
	{"method": "Type", "resultattempt": ""},

	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetSizes", "args": []interface{}{"no"}, "gettermethod": "Sizes", "resultattempt": "no"},
	{"method": "SetSrcSet", "args": []interface{}{"small.jpg 1x, large.jpg 2x"}, "gettermethod": "SrcSet", "resultattempt": "small.jpg 1x, large.jpg 2x"},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if anchor, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, anchor, result)
			}

		}

	}
}
