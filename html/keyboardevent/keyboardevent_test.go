package keyboardevent

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

	if k, err := New("keypress", map[string]interface{}{"code": "KeyA"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object KeyboardEvent]", k.ToString())
		test.AssertExpect(t, "KeyA", k.Code())
	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\")")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object KeyboardEvent]", k.ToString())

		}
	}

}

func TestAltKey(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{altKey:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v := k.AltKey(); v != true {
				test.AssertExpect(t, true, v)
			}

		}
	}

}

func TestCtrlKey(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{ctrlKey:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, true, k.CtrlKey())

		}
	}

}

func TestKey(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{key:\"Enter\"})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Enter", k.Key())

		}
	}

}

func TestCode(t *testing.T) {

	//independent position (FR position) == (press q obtain KeyA, press a obtain KeyQ )

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{code:\"KeyA\"})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "KeyA", k.Code())

		}
	}

}

func TestIsComposing(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{isComposing:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, true, k.IsComposing())

		}
	}

}

func TestLocation(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{location:1})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, int64(1), k.Location())

		}
	}

}

func TestMetaKey(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{metaKey:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, true, k.MetaKey())

		}
	}

}

func TestRepeat(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{repeat:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, true, k.Repeat())

		}
	}

}

func TestShiftKey(t *testing.T) {

	js.Eval("keypress=new KeyboardEvent(\"keyup\",{shiftKey:true})")

	if obj := js.Global().Get("keypress"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, true, k.ShiftKey())

		}
	}

}
