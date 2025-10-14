package domexception

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("err= new DOMException()")

	if obj := js.Global().Get("err"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "DOMException", d.ConstructName_())

		}
	}

}

func TestException(t *testing.T) {

	if e, err := New(); test.AssertErr(t, err) {
		test.AssertExpect(t, "Error", e.ToString_())
	}

	var message string = "message error"

	if e, err := New(message); test.AssertErr(t, err) {
		test.AssertExpect(t, "Error: message error", e.ToString_())

		if str, err := e.Message(); test.AssertErr(t, err) {
			test.AssertExpect(t, message, str)
		}

	}

	var customname string = "custom name"
	if e, err := New(message, customname); test.AssertErr(t, err) {
		test.AssertExpect(t, "custom name: message error", e.ToString_())

		if str, err := e.Name(); test.AssertErr(t, err) {
			test.AssertExpect(t, customname, str)
		}

	}

}
