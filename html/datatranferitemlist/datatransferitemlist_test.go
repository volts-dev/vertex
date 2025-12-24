package datatranferitemlist

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/file"
	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`dt=new DataTransfer()
	itemlist=dt.items`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj = js.Global().Get("itemlist"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DataTransferItemList]", d.ToString_())

		}
	}

}

func TestAdd(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj = js.Global().Get("itemlist"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			if f, err := file.New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

				test.AssertErr(t, d.Add(f))
				if l, err := d.Length(); test.AssertErr(t, err) {
					test.AssertExpect(t, 1, l)
				}

			}

		}
	}

}

func TestClear(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj = js.Global().Get("itemlist"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {

			if f, err := file.New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

				test.AssertErr(t, d.Add(f))
				test.AssertErr(t, d.Clear())
				if l, err := d.Length(); test.AssertErr(t, err) {
					test.AssertExpect(t, 0, l)
				}

			}

		}
	}

}

func TestRemove(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj = js.Global().Get("itemlist"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {
			d.Clear()

			if f, err := file.New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

				test.AssertErr(t, d.Add(f))

				test.AssertErr(t, d.Remove(0))
				if l, err := d.Length(); test.AssertErr(t, err) {
					test.AssertExpect(t, 0, l)
				}

			}

		}
	}

}

func TestDataTransferItem(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj = js.Global().Get("itemlist"); test.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); test.AssertErr(t, err) {
			d.Clear()

			if f, err := file.New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {
				test.AssertErr(t, d.Add(f))

				if item, err := d.DataTransferItem(0); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object DataTransferItem]", item.ToString_())
				}

				_, err := d.DataTransferItem(1)

				test.AssertExpect(t, err, js.ErrUndefinedValue)
			}

		}
	}

}
