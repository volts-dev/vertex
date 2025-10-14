package htmlselectelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/htmloptionelement"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`f=document.createElement("form")
	s= document.createElement("select")
	f.appendChild(s)
	o= document.createElement("option")
	o.value="t3st"
	s.appendChild(o)
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if selectobj, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLSelectElement", selectobj.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if selectobj, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLSelectElement", selectobj.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Autofocus", "resultattempt": false},
	{"method": "SetAutofocus", "args": []interface{}{true}, "gettermethod": "Autofocus", "resultattempt": true},
	{"method": "Disabled", "resultattempt": false},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "Length", "resultattempt": 1},
	{"method": "Name", "resultattempt": ""},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "Options", "type": "constructnamechecking", "resultattempt": "HTMLOptionsCollection"},
	{"method": "Multiple", "resultattempt": false},
	{"method": "SetMultiple", "args": []interface{}{true}, "gettermethod": "Multiple", "resultattempt": true},
	{"method": "Required", "resultattempt": false},
	{"method": "SetRequired", "args": []interface{}{true}, "gettermethod": "Required", "resultattempt": true},
	{"method": "SelectedIndex", "resultattempt": 0},
	{"method": "SetSelectedIndex", "args": []interface{}{0}, "gettermethod": "SelectedIndex", "resultattempt": 0},
	{"method": "SelectedOptions", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Size", "resultattempt": 0},
	{"method": "SetSize", "args": []interface{}{3}, "gettermethod": "Size", "resultattempt": 3},
	{"method": "Type", "resultattempt": "select-multiple"},
	{"method": "Value", "resultattempt": "t3st"},
	{"method": "SetValue", "args": []interface{}{"t3st"}, "gettermethod": "Value", "resultattempt": "t3st"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "WillValidate", "resultattempt": false},
	{"method": "ReportValidity", "resultattempt": true},
	{"method": "SetCustomValidity", "args": []interface{}{"hello"}, "type": "error", "resultattempt": nil},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("s"); test.AssertErr(t, obj.Error()) {

		if selectobj, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, selectobj, result)
			}

		}

	}
}

func TestAdd(t *testing.T) {
	js.Eval(`
	sadd= document.createElement("select")
	`)

	if obj := js.Global().Get("sadd"); test.AssertErr(t, obj.Error()) {

		if sadd, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if l, err := sadd.Length(); test.AssertErr(t, err) {
				test.AssertExpect(t, 0, l)

				if option, err := htmloptionelement.Option("test"); test.AssertErr(t, err) {

					test.AssertErr(t, sadd.Add(option))

					if l2, err := sadd.Length(); test.AssertErr(t, err) {
						test.AssertExpect(t, 1, l2)

					}

				}
			}

		}

	}

}

func TestItem(t *testing.T) {
	js.Eval(`
	sitem= document.createElement("select")
	oitem= document.createElement("option")
	oitem.name="t3stitem"
	sitem.appendChild(oitem)

	`)

	if obj := js.Global().Get("sitem"); test.AssertErr(t, obj.Error()) {

		if sitem, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if item, err := sitem.Item(0); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLOptionElement", item.ConstructName_())
			}

		}

	}

}

func TestNamedItem(t *testing.T) {
	js.Eval(`
	sitemn= document.createElement("select")
	oitemn= document.createElement("option")
	oitemn.id="t3stitem"
	sitemn.appendChild(oitemn)

	`)

	if obj := js.Global().Get("sitemn"); test.AssertErr(t, obj.Error()) {

		if sitem, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if item, err := sitem.NamedItem("t3stitem"); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLOptionElement", item.ConstructName_())
			}

		}

	}

}

func TestRemove(t *testing.T) {
	js.Eval(`
	sitemr= document.createElement("select")
	oitemr= document.createElement("option")
	sitemr.appendChild(oitemr)
	`)

	if obj := js.Global().Get("sitemr"); test.AssertErr(t, obj.Error()) {

		if sitemr, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if l, err := sitemr.Length(); test.AssertErr(t, err) {
				test.AssertExpect(t, 1, l)

				test.AssertErr(t, sitemr.Remove(0))
				if l2, err := sitemr.Length(); test.AssertErr(t, err) {
					test.AssertExpect(t, 0, l2)

				}

			}

		}

	}

}
