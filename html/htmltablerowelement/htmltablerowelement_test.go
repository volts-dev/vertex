package htmltablerowelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`tr=document.createElement("tr")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if tr, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableRowElement", tr.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("tr"); test.AssertErr(t, obj.Error()) {

		if tr, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableRowElement", tr.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "RowIndex", "resultattempt": -1},
	{"method": "SectionRowIndex", "resultattempt": -1},
	{"method": "Cells", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("tr"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertCell(t *testing.T) {

	js.Eval(`tri=document.createElement("tr")
	`)

	if obj := js.Global().Get("tri"); test.AssertErr(t, obj.Error()) {

		if tr, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if cell, err := tr.InsertCell(); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTableCellElement", cell.ConstructName_())
				if cells, err := tr.Cells(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, cells.Length())
				}

			}
		}

	}
}

func TestDeleteCell(t *testing.T) {

	js.Eval(`trd=document.createElement("tr")
	`)

	if obj := js.Global().Get("trd"); test.AssertErr(t, obj.Error()) {

		if tr, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			tr.InsertCell()
			tr.InsertCell()
			if cells, err := tr.Cells(); test.AssertErr(t, err) {

				test.AssertExpect(t, 2, cells.Length())

				test.AssertErr(t, tr.DeleteCell(0))
				if cells2, err := tr.Cells(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, cells2.Length())
				}

			}
		}

	}
}
