package stream

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

	if s, err := NewReadableStream(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object ReadableStream]", s.ToString_())

	}

	if s, err := NewWritableStream(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object WritableStream]", s.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("r=new ReadableStream();w=new WritableStream();")

	if obj := js.Global().Get("r"); test.AssertErr(t, obj.Error()) {
		if d, err := NewReadableStreamFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ReadableStream]", d.ToString_())

		}
	}
	if obj := js.Global().Get("w"); test.AssertErr(t, obj.Error()) {
		if d, err := NewWriteableStreamFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object WritableStream]", d.ToString_())

		}
	}

}

func TestLocked(t *testing.T) {
	if s, err := NewReadableStream(); test.AssertErr(t, err) {

		if locked, err := s.Locked(); test.AssertErr(t, err) {

			test.AssertExpect(t, false, locked)
		}

	}

	if s, err := NewWritableStream(); test.AssertErr(t, err) {

		if locked, err := s.Locked(); test.AssertErr(t, err) {

			test.AssertExpect(t, false, locked)
		}

	}
}

func TestCancelReadable(t *testing.T) {
	if s, err := NewReadableStream(); test.AssertErr(t, err) {

		if pcancel, err := s.Cancel(); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Promise]", pcancel.ToString_())
		}

	}
}

func TestAbortWritable(t *testing.T) {
	if w, err := NewWritableStream(); test.AssertErr(t, err) {

		if pabort, err := w.Abort("i want"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Promise]", pabort.ToString_())
		}

	}
}

func TestCloseWritable(t *testing.T) {
	if w, err := NewWritableStream(); test.AssertErr(t, err) {

		if pabort, err := w.Close(); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Promise]", pabort.ToString_())
		}

	}
}

func TestGetReader(t *testing.T) {
	if s, err := NewReadableStream(); test.AssertErr(t, err) {

		if reader, err := s.GetReader(); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ReadableStreamDefaultReader]", reader.ToString_())
		}
		//doc say that the stream is locked when get reader so check it
		if locked, err := s.Locked(); test.AssertErr(t, err) {

			test.AssertExpect(t, true, locked)
		}
	}
}

func TestTeeReadable(t *testing.T) {
	if s, err := NewReadableStream(); test.AssertErr(t, err) {

		if a, err := s.Tee(); test.AssertErr(t, err) {

			if test.AssertExpect(t, 2, len(a)) {

				for i := 0; i < 2; i++ {
					test.AssertExpect(t, "[object ReadableStream]", a[i].ToString_())
				}

			}

		}
	}
}
