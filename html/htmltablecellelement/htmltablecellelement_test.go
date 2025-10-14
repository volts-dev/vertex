package htmltablecellelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`td= document.createElement("td")
	th= document.createElement("th")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if td, err := NewTd(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableCellElement", td.ConstructName_())
		}

		if th, err := NewTh(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableCellElement", th.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("td"); test.AssertErr(t, obj.Error()) {

		if td, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableCellElement", td.ConstructName_())
		}

	}

	if obj := js.Global().Get("th"); test.AssertErr(t, obj.Error()) {

		if th, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableCellElement", th.ConstructName_())
		}

	}
}
