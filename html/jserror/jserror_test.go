package jserror

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

func TestNew(t *testing.T) {

	if e, err := New("throw error"); test.AssertErr(t, err) {

		test.AssertExpect(t, "Error: throw error", e.ToString_())

	}

	var message string = "message error"

	if e, err := New(message); test.AssertErr(t, err) {
		test.AssertExpect(t, "Error: message error", e.ToString_())

		if str, err := e.Message(); test.AssertErr(t, err) {
			test.AssertExpect(t, message, str)
		}
		message = "message error2"
		e.SetMessage(message)

		test.AssertExpect(t, "Error: message error2", e.ToString_())

	}

	var customname string = "custom name"
	if e, err := New(message); test.AssertErr(t, err) {
		e.SetName(customname)
		test.AssertExpect(t, "custom name: message error2", e.ToString_())

		if str, err := e.Name(); test.AssertErr(t, err) {
			test.AssertExpect(t, customname, str)
		}
	}

}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("err=new Error()")

	if obj := js.Global().Get("err"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Error", d.ToString_())

		}
	}

}
