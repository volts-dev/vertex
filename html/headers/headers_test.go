package headers

import (
	"errors"
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
		if v, err := h.Get("X-custom"); test.AssertErr(t, err) {

			test.AssertExpect(t, "1234", v)
		}
	}

}

func TestDelete(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertErr(t, h.Delete("X-custom"))

		_, err := h.Get("X-custom")

		test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
	}

}

func TestHas(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		if b, err := h.Has("X-custom"); test.AssertErr(t, err) {
			test.AssertExpect(t, true, b)
		}
		test.AssertErr(t, h.Delete("X-custom"))

		if b, err := h.Has("X-custom"); test.AssertErr(t, err) {
			test.AssertExpect(t, false, b)
		}

	}

}

func TestKeys(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Headers Iterator]", it.ToString_())
		}
	}
}

func TestValues(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Headers Iterator]", it.ToString_())
		}
	}
}

func TestSet(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, h.Append("X-custom", "1234"))

		test.AssertErr(t, h.Set("X-custom", "4567"))
		if v, err := h.Get("X-custom"); test.AssertErr(t, err) {

			test.AssertExpect(t, "4567", v)
		}

	}
}

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
