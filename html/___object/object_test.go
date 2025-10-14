package object

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`
	objnew= new Object()
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if o, err := New(); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object Object]", o.ToString_())
	}

}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`
	objnew= new Object()
	`)

	if obj := js.Global().Get("objnew"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Object]", o.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Keys", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "Values", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "Map", "type": "constructnamechecking", "resultattempt": "Map"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("Array"); test.AssertErr(t, obj.Error()) {

		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, o, result)
			}

		}

	}
}
