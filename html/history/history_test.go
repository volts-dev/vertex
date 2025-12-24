package history

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/objectmap"
	"github.com/volts-dev/vertex/js/reflect"
)

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`h= window.history
	`)

	if obj := js.Global().Get("h"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object History]", h.ToString_())
		}

	}
}

func TestState(t *testing.T) {

	js.Eval(`h= window.history
	`)

	if obj := js.Global().Get("h"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			_, err := h.State()

			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
			o, _ := objectmap.New(js.NewArray_(js.NewArray_("title", "teststate")))
			h.PushState(o, "hello")
			state, err := h.State()
			test.AssertExpect(t, "[object Map]", state.(objectmap.ObjectMap).ObjectMap_().ToString_())
		}

	}

}

func TestLength(t *testing.T) {

	js.Eval(`h= window.history
	`)

	if obj := js.Global().Get("h"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if l, err := h.Length(); test.AssertErr(t, err) {

				test.AssertExpect(t, true, l >= 1)
				o, _ := objectmap.New(js.NewArray_(js.NewArray_("title", "testLength1")))
				h.PushState(o, "hello")
				if l2, err := h.Length(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, l2-l)
				}
			}
		}

	}

}

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
