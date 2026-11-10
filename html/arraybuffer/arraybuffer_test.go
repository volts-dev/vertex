package arraybuffer

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

	if a, err := New(8); test.AssertErr(t, err) {

		test.AssertExpect(t, int64(8), a.ByteLength())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("ab = new ArrayBuffer()")

	if obj := js.Global().Get("ab"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ArrayBuffer]", d.ToString())

		}
	}

}

func TestSlice(t *testing.T) {

	if a, err := New(32); test.AssertErr(t, err) {

		if b, err := a.Slice(10); test.AssertErr(t, err) {

			test.AssertExpect(t, int64(22), b.ByteLength())

		}

		if b, err := a.Slice(10, 16); test.AssertErr(t, err) {

			test.AssertExpect(t, int64(6), b.ByteLength())

		}

	}
}

func TestIsView(t *testing.T) {

	js.Eval("customuint16=new Uint16Array()")
	if obj := js.Global().Get("customuint16"); test.AssertErr(t, obj.Error()) {
		if a, err := js.ToObject(obj); test.AssertErr(t, err) {
			if ok := IsView(a); ok {

				test.AssertExpect(t, true, ok)

			}
		}

	}
	js.Eval("customuint16=\"string\"")
	if obj := js.Global().Get("customuint16"); test.AssertErr(t, obj.Error()) {
		if a, err := js.ToObject(obj); test.AssertErr(t, err) {
			if ok := IsView(a); ok {
				test.AssertExpect(t, false, ok)
			}
		}

	}
}
