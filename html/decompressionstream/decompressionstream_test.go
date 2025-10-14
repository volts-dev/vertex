package decompressionstream

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`dec=new DecompressionStream("gzip")`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("dec"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "DecompressionStream", nav.ConstructName_())

		}
	}

	decstream, err := New("gzip")
	test.AssertExpect(t, nil, err)
	test.AssertExpect(t, "DecompressionStream", decstream.ConstructName_())
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Readable", "type": "constructnamechecking", "resultattempt": "ReadableStream"},
	{"method": "Writable", "type": "constructnamechecking", "resultattempt": "WritableStream"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("dec"); test.AssertErr(t, obj.Error()) {

		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, nav, result)
			}

		}

	}
}
