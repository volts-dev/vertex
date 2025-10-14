package nodelist

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`
	i1=document.createElement("input")
	i1.name="toto"
	i2=document.createElement("input")
	i2.name="toto"
	document.body.appendChild(i1)
	document.body.appendChild(i2)
	list=document.getElementsByName("toto")
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {
		if rect, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object NodeList]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "HTMLInputElement"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {

		if image, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, image, result)
			}

		}

	}
}
