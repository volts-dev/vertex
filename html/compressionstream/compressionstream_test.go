package compressionstream

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`enc=new CompressionStream("gzip")`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("enc"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "CompressionStream", nav.ConstructName_())

		}
	}

	encstream, err := New("gzip")
	test.AssertExpect(t, nil, err)
	test.AssertExpect(t, "CompressionStream", encstream.ConstructName_())
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Readable", "type": "constructnamechecking", "resultattempt": "ReadableStream"},
	{"method": "Writable", "type": "constructnamechecking", "resultattempt": "WritableStream"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("enc"); test.AssertErr(t, obj.Error()) {

		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, nav, result)
			}

		}

	}
}
