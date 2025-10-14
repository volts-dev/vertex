package customevent

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

	if d, err := New("hello", "world"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object CustomEvent]", d.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("customevent=new CustomEvent(\"hello\")")

	if obj := js.Global().Get("customevent"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object CustomEvent]", d.ToString_())

		}
	}

}

func TestDetail(t *testing.T) {

	if d, err := New("hello", "world"); test.AssertErr(t, err) {

		if detail, err := d.Detail(); test.AssertErr(t, err) {
			test.AssertExpect(t, "world", detail)
		}

	}
}
