package blob

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/js/array"
	"github.com/volts-dev/vertex/js/arraybuffer"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/js/typedarray"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(); test.AssertErr(t, err) {

		if s, err := a.Size(); test.AssertErr(t, err) {

			test.AssertExpect(t, int64(0), s)

		}
	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("blob1=new Blob()")

	if obj := js.Global().Get("blob1"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Blob]", d.ToString_())

		}
	}

}

func TestNewWithArrayBuffer(t *testing.T) {

	if a, err := arraybuffer.New(8); test.AssertErr(t, err) {
		if viewuint8, err := typedarray.NewInt8Array(a); test.AssertErr(t, err) {
			viewuint8.Fill(7)

			if ab, err := NewWithArrayBuffer(a); test.AssertErr(t, err) {

				if s, err := ab.Size(); test.AssertErr(t, err) {

					test.AssertExpect(t, int64(8), s)

				}

			}

		}
	}

}

func TestNewWith2ArrayBuffer(t *testing.T) {

	if a, err := arraybuffer.New(8); test.AssertErr(t, err) {
		if viewuint8, err := typedarray.NewInt8Array(a); test.AssertErr(t, err) {
			viewuint8.Fill(7)

			astring := array.From_("Hello World")
			if struint8, err := typedarray.NewUint8ArrayFrom(astring); test.AssertErr(t, err) {

				if appendblob, err := New(viewuint8, struint8); test.AssertErr(t, err) {

					if s, err := appendblob.Size(); test.AssertErr(t, err) {

						test.AssertExpect(t, int64(19), s)

					}

				}

			}

		}
	}

}

func TestIsClosed(t *testing.T) {

	if a, err := New(); test.AssertErr(t, err) {

		_, err := a.IsClosed()
		test.AssertExpect(t, js.ErrNotImplementedFunc, err)
	}
}

func TestClosed(t *testing.T) {

	if a, err := New(); test.AssertErr(t, err) {

		err := a.Close()
		test.AssertExpect(t, js.ErrNotImplementedFunc, err)
	}
}

func TestSlice(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); test.AssertErr(t, err) {

		if b, err := struint8.Buffer(); test.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); test.AssertErr(t, err) {

				if blob2, err := ab.Slice(0, 6); test.AssertErr(t, err) {
					if s, err := blob2.Size(); test.AssertErr(t, err) {

						test.AssertExpect(t, int64(6), s)

					}
				}

			}
		}

	}

}

func TestStream(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); test.AssertErr(t, err) {

		if b, err := struint8.Buffer(); test.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); test.AssertErr(t, err) {

				if stream, err := ab.Stream(); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object ReadableStream]", stream.ToString_())
				}

			}
		}

	}

}

func TestArrayBuffer(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); test.AssertErr(t, err) {

		if b, err := struint8.Buffer(); test.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); test.AssertErr(t, err) {

				if blobbuffer, err := ab.ArrayBuffer(); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object Promise]", blobbuffer.ToString_())
				}

			}
		}

	}

}
func TestText(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); test.AssertErr(t, err) {

		if b, err := struint8.Buffer(); test.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); test.AssertErr(t, err) {

				if blobtext, err := ab.Text(); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object Promise]", blobtext.ToString_())
				}

			}
		}

	}

}
