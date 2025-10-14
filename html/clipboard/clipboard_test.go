package clipboard

import (
	"fmt"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/blob"
	"github.com/volts-dev/vertex/html/clipboarditem"
)

var clipitem clipboarditem.ClipboardItem

var methodsAttempt []map[string]interface{}

func TestMain(m *testing.M) {
	var err error
	var b blob.Blob
	reflect.SetSyscall()

	if b, err = blob.New("{\"hello\"}"); err != nil {
		fmt.Printf("error %s\n", err.Error())
	}
	blobitem := map[string]blob.Blob{"application/json": b}

	if clipitem, err = clipboarditem.New(blobitem); err != nil {
		fmt.Printf("error %s\n", err.Error())
	}

	methodsAttempt = []map[string]interface{}{
		{"method": "Read", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "ReadText", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "Write", "args": []interface{}{[]clipboarditem.ClipboardItem{clipitem}}, "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "WriteText", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	}

	js.Eval(`clip=window.navigator.clipboard`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("clip"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Clipboard", nav.ConstructName_())

		}
	}

}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("clip"); test.AssertErr(t, obj.Error()) {

		if clip, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, clip, result)
			}

		}

	}
}
