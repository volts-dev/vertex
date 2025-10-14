package clipboarditem

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`blob=new Blob()
	clipitem= new ClipboardItem({
        ['file']: blob
      })`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("clipitem"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "ClipboardItem", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Types", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "GetType", "args": []interface{}{"file"}, "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("clipitem"); test.AssertErr(t, obj.Error()) {

		if clip, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, clip, result)
			}

		}

	}
}
