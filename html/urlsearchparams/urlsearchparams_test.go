package urlsearchparams

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestNew(t *testing.T) {

	if h, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "URLSearchParams", h.ConstructName_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`u= new URLSearchParams()
	`)

	if obj := js.Global().Get("u"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "URLSearchParams", h.ConstructName_())
		}

	}
}

func TestAppend(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))
		if v, err := u.Get("TestKey"); test.AssertErr(t, err) {

			test.AssertExpect(t, "1234", v)
		}
	}

}

func TestDelete(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))

		test.AssertErr(t, u.Delete("TestKey"))

		_, err := u.Get("TestKey")

		test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
	}

}

func TestHas(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))

		if b, err := u.Has("TestKey"); test.AssertErr(t, err) {
			test.AssertExpect(t, true, b)
		}
		test.AssertErr(t, u.Delete("TestKey"))

		if b, err := u.Has("TestKey"); test.AssertErr(t, err) {
			test.AssertExpect(t, false, b)
		}

	}

}

func TestKeys(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))

		if it, err := u.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object URLSearchParams Iterator]", it.ToString_())
		}
	}
}

func TestValues(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))

		if it, err := u.Keys(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object URLSearchParams Iterator]", it.ToString_())
		}
	}
}

func TestSet(t *testing.T) {

	if u, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, u.Append("TestKey", "1234"))

		test.AssertErr(t, u.Set("TestKey", "4567"))
		if v, err := u.Get("TestKey"); test.AssertErr(t, err) {

			test.AssertExpect(t, "4567", v)
		}

	}
}

func TestSort(t *testing.T) {

	if u, err := New("c=4&a=2&b=3&a=1"); test.AssertErr(t, err) {

		if err := u.Sort(); test.AssertErr(t, err) {

			test.AssertExpect(t, "a=2&a=1&b=3&c=4", u.ToString_())
		}

	}
}

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
