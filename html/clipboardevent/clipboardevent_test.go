package clipboardevent

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/datatransfer"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	event=new ClipboardEvent(dt)`)
	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("event"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ClipboardEvent]", d.ToString_())

		}
	}

}

func TestNew(t *testing.T) {

	if dt, err := datatransfer.New(); test.AssertErr(t, err) {
		if d, err := New(dt); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ClipboardEvent]", d.ToString_())

		}
	}

}

/*
func TestClipboardData(t *testing.T) {

	if obj := js.Global().Get("event"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if data, err := d.ClipboardData(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object ClipboardEvent]", data.ToString_())
			}

		}
	}
}
*/
