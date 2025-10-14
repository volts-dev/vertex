package domrectlist

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`d=document.createElement("div")
	document.body.appendChild(d)
	rectlist=d.getClientRects()
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("rectlist"); test.AssertErr(t, obj.Error()) {
		if rect, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DOMRectList]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "DOMRect"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("rectlist"); test.AssertErr(t, obj.Error()) {

		if image, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, image, result)
			}

		}

	}
}
