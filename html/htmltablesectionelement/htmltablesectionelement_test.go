package htmltablesectionelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`tbody=document.createElement("tbody")
	tfoot=document.createElement("tfoot")
	thead=document.createElement("thead")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if tr, err := NewTBody(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}
		if tr, err := NewTFoot(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}
		if tr, err := NewTHead(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("tbody"); test.AssertErr(t, obj.Error()) {

		if tbody, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableSectionElement", tbody.ConstructName_())
		}

	}

	if obj := js.Global().Get("thead"); test.AssertErr(t, obj.Error()) {

		if thead, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableSectionElement", thead.ConstructName_())
		}

	}

	if obj := js.Global().Get("tfoot"); test.AssertErr(t, obj.Error()) {

		if tfoot, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableSectionElement", tfoot.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Rows", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("tbody"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertRow(t *testing.T) {

	js.Eval(`tbodyi=document.createElement("tbody")
	`)

	if obj := js.Global().Get("tbodyi"); test.AssertErr(t, obj.Error()) {

		if tbody, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if r, err := tbody.InsertRow(); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTableRowElement", r.ConstructName_())
				if cr, err := tbody.Rows(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, cr.Length())
				}

			}
		}

	}
}

func TestDeleteRow(t *testing.T) {

	js.Eval(`tbodyr=document.createElement("tbody")
	`)

	if obj := js.Global().Get("tbodyr"); test.AssertErr(t, obj.Error()) {

		if tbody, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			tbody.InsertRow()
			tbody.InsertRow()
			if rows, err := tbody.Rows(); test.AssertErr(t, err) {

				test.AssertExpect(t, 2, rows.Length())

				test.AssertErr(t, tbody.DeleteRow(0))
				if rows, err := tbody.Rows(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, rows.Length())
				}

			}

		}

	}
}
