package domtokenlist

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`l=document.createElement("link")
	l.rel="a b c"
	list=l.relList
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "DOMTokenList", list.ConstructName_())

		}
	}

}

func TestItem(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if i, err := list.Item(0); test.AssertErr(t, err) {

				test.AssertExpect(t, "a", i)
			}

		}
	}

}

func TestContains(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if b, err := list.Contains("b"); test.AssertErr(t, err) {

				test.AssertExpect(t, true, b)
			}

			if b, err := list.Contains("d"); test.AssertErr(t, err) {

				test.AssertExpect(t, false, b)
			}

		}
	}

}

func TestAdd(t *testing.T) {

	js.Eval(`ladd=document.createElement("link")
	ladd.rel="a b c"
	listadd=ladd.relList
	`)

	if obj := js.Global().Get("listadd"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertErr(t, list.Add("d", "e"))
			test.AssertExpect(t, "a b c d e", list.ToString_())

		}
	}

}

func TestRemove(t *testing.T) {

	js.Eval(`lr=document.createElement("link")
	lr.rel="a b c"
	listr=lr.relList
	`)

	if obj := js.Global().Get("listr"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertErr(t, list.Remove("b", "c"))
			test.AssertExpect(t, "a", list.ToString_())

		}
	}

}

func TestReplace(t *testing.T) {

	js.Eval(`lre=document.createElement("link")
	lre.rel="a b c"
	listre=lre.relList
	`)

	if obj := js.Global().Get("listre"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertErr(t, list.Replace("b", "k"))
			test.AssertExpect(t, "a k c", list.ToString_())

		}
	}

}

func TestToggle(t *testing.T) {

	js.Eval(`lt=document.createElement("link")
	lt.rel="a b c"
	listt=lt.relList
	`)

	if obj := js.Global().Get("listt"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if b, err := list.Toggle("b"); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)
				test.AssertExpect(t, "a c", list.ToString_())

			}

		}
	}

}

func TestSupports(t *testing.T) {

	js.Eval(`ls=document.createElement("link")
	ls.rel="a b c"
	lists=ls.relList
	`)

	if obj := js.Global().Get("lists"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if b, err := list.Supports(""); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)

			}

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Entries", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
	{"method": "Keys", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
	{"method": "Values", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {

		if maphtml, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, maphtml, result)
			}

		}

	}
}

func TestForEach(t *testing.T) {

	if obj := js.Global().Get("list"); test.AssertErr(t, obj.Error()) {
		if list, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			var i int
			list.ForEach(func(s string) {

				i++

			})
			test.AssertExpect(t, 3, i)
		}
	}

}
