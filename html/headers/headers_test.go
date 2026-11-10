package headers

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestNew(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "Headers", h.ConstructName_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`h= new Headers()
	`)

	if obj := js.Global().Get("h"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Headers", h.ConstructName_())
		}

	}
}

func TestAppend(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertExpect(t, "1234", h.Get("X-custom"))
	}

}

func TestDelete(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertErr(t, h.Delete("X-custom"))

		h.Get("X-custom")

	}

}

func TestHas(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertExpect(t, true, h.Has("X-custom"))
		test.AssertErr(t, h.Delete("X-custom"))

		test.AssertExpect(t, false, h.Has("X-custom"))

	}

}

func TestKeys(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Headers Iterator]", it.ToString())
		}
	}
}

func TestValues(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Headers Iterator]", it.ToString())
		}
	}
}

func TestSet(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertErr(t, h.Set("X-custom", "4567"))

		test.AssertExpect(t, "4567", h.Get("X-custom"))

	}
}

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
