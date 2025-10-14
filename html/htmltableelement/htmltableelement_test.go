package htmltableelement

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`t=document.createElement("table")
	c=document.createElement("caption")
	c.testContent="title"
	t.appendChild(c)
	thead=document.createElement("thead")
	t.appendChild(thead)
	r=document.createElement("row")
	t.insertRow(r)
	tfoot=document.createElement("tfoot")
	t.appendChild(tfoot)

	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if table, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLTableElement", table.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("t"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTableElement", table.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Rows", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Caption", "type": "constructnamechecking", "resultattempt": "HTMLTableCaptionElement"},
	{"method": "TBodies", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "TFoot", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "THead", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "CreateCaption", "type": "constructnamechecking", "resultattempt": "HTMLTableCaptionElement"},
	{"method": "CreateTHead", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "CreateTFoot", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("t"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertRow(t *testing.T) {

	js.Eval(`tir=document.createElement("table")
	`)

	if obj := js.Global().Get("tir"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if r, err := table.InsertRow(); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTableRowElement", r.ConstructName_())
				if cr, err := table.Rows(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, cr.Length())
				}

			}
		}

	}
}

func TestDeleteRow(t *testing.T) {

	js.Eval(`tdr=document.createElement("table")
	`)

	if obj := js.Global().Get("tdr"); test.AssertErr(t, obj.Error()) {

		if table, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			table.InsertRow()
			table.InsertRow()
			if cr, err := table.Rows(); test.AssertErr(t, err) {

				test.AssertExpect(t, 2, cr.Length())

				test.AssertErr(t, table.DeleteRow(0))
				if cr2, err := table.Rows(); test.AssertErr(t, err) {

					test.AssertExpect(t, 1, cr2.Length())
				}

			}

		}

	}
}

func TestDeleteCaption(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if table, err := New(doc); test.AssertErr(t, err) {

			table.CreateCaption()
			_, err := table.Caption()
			test.AssertExpect(t, false, errors.Is(err, js.ErrUndefinedValue))
			test.AssertErr(t, table.DeleteCaption())
			_, err = table.Caption()
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
		}

	}

}

func TestDeleteTFoot(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if table, err := New(doc); test.AssertErr(t, err) {

			table.CreateTFoot()
			_, err := table.TFoot()
			test.AssertExpect(t, false, errors.Is(err, js.ErrUndefinedValue))
			test.AssertErr(t, table.DeleteTFoot())
			_, err = table.TFoot()
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
		}

	}

}

func TestDeleteTHead(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if table, err := New(doc); test.AssertErr(t, err) {

			table.CreateTHead()
			_, err := table.THead()
			test.AssertExpect(t, false, errors.Is(err, js.ErrUndefinedValue))
			test.AssertErr(t, table.DeleteTHead())
			_, err = table.THead()
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
		}

	}

}
