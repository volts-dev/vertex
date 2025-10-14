package datatransferitem

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

// doc say that can accept add(string) but dont work return (Failed to execute 'add' on 'DataTransferItemList': parameter 1 is not of type 'File'.)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	item1=dt.items[0]`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj = js.Global().Get("item1"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DataTransferItem]", d.ToString_())

		}
	}

}

func TestKind(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj = js.Global().Get("item1"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			if kind, err := d.Kind(); test.AssertErr(t, err) {
				test.AssertExpect(t, "file", kind)
			}

		}
	}

}

func TestType(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj = js.Global().Get("item1"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			if typed, err := d.Type(); test.AssertErr(t, err) {
				test.AssertExpect(t, "image/png", typed)
			}

		}
	}

}
func TestGetAsFile(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj = js.Global().Get("item1"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			if file, err := d.GetAsFile(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object File]", file.ToString_())
			}

		}
	}

}
