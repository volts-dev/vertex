package filelist

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	file2 = new File(['(⌐□_□)'], 'chucknorris2.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	dt.items.add(file2)
	inputfiles=dt.files`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj = js.Global().Get("inputfiles"); test.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object FileList]", f.ToString_())

		}
	}
}

func TestLength(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj = js.Global().Get("inputfiles"); test.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); test.AssertErr(t, err) {
			if l, err := f.Length(); test.AssertErr(t, err) {
				test.AssertExpect(t, 2, l)
			}

		}
	}
}

func TestItem(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj = js.Global().Get("inputfiles"); test.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); test.AssertErr(t, err) {
			if file, err := f.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object File]", file.ToString_())
			}

		}
	}
}
