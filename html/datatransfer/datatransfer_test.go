package datatransfer

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)`)
	m.Run()

}

func TestNew(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object DataTransfer]", d.ToString_())

	}
}

func TestFiles(t *testing.T) {

	if obj := js.Global().Get("dt"); test.AssertErr(t, obj.Error()) {

		if dt, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if files, err := dt.Files(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object FileList]", files.ToString_())
			}
		}
	}
}

func TestItems(t *testing.T) {

	if obj := js.Global().Get("dt"); test.AssertErr(t, obj.Error()) {

		if dt, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if items, err := dt.Items(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object DataTransferItemList]", items.ToString_())
			}
		}
	}
}

func TestTypes(t *testing.T) {

	if obj := js.Global().Get("dt"); test.AssertErr(t, obj.Error()) {

		if dt, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if types, err := dt.Types(); test.AssertErr(t, err) {
				test.AssertExpect(t, "Files", types.ToString_())
			}
		}
	}
}
