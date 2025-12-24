package abortcontroller

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object AbortController]", a.ToString_())

	}
}
func TestNewFromJSObject(t *testing.T) {

	js.Eval("abortctrl=new AbortController()")

	if obj := js.Global().Get("abortctrl"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object AbortController]", d.ToString_())

		}
	}

}

func TestAbort(t *testing.T) {

	var isAborted bool = false
	if a, err := New(); test.AssertErr(t, err) {

		if as, err := a.Signal(); test.AssertErr(t, err) {

			as.OnAbort(func(e event.Event) error {
				isAborted = true
				return nil
			})
			a.Abort()

			test.AssertExpect(t, true, isAborted)

		}
	}
}
