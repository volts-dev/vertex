package arraybuffer

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(8); test.AssertErr(t, err) {

		if l, err := a.ByteLength(); test.AssertErr(t, err) {

			test.AssertExpect(t, int64(8), l)

		}
	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("ab = new ArrayBuffer()")

	if obj := js.Global().Get("ab"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ArrayBuffer]", d.ToString_())

		}
	}

}

func TestSlice(t *testing.T) {

	if a, err := New(32); test.AssertErr(t, err) {

		if b, err := a.Slice(10); test.AssertErr(t, err) {

			if l, err := b.ByteLength(); test.AssertErr(t, err) {

				test.AssertExpect(t, int64(22), l)

			}

		}

		if b, err := a.Slice(10, 16); test.AssertErr(t, err) {

			if l, err := b.ByteLength(); test.AssertErr(t, err) {

				test.AssertExpect(t, int64(6), l)

			}

		}

	}
}

func TestIsView(t *testing.T) {

	js.Eval("customuint16=new Uint16Array()")
	if obj := js.Global().Get("customuint16"); test.AssertErr(t, obj.Error()) {
		if a, err := object.ToObject(obj); test.AssertErr(t, err) {
			if ok, err := IsView(a); test.AssertErr(t, err) {

				test.AssertExpect(t, true, ok)

			}
		}

	}
	js.Eval("customuint16=\"string\"")
	if obj := js.Global().Get("customuint16"); test.AssertErr(t, obj.Error()) {
		if a, err := object.ToObject(obj); test.AssertErr(t, err) {
			if ok, err := IsView(a); test.AssertErr(t, err) {
				test.AssertExpect(t, false, ok)
			}
		}

	}
}
