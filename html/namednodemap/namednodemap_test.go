package namednodemap

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/attr"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`b=document.createElement("button")
	b.setAttribute("hello","world")
	b.setAttributeNS("name","high","low")
	listattr=b.attributes
	`)

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("listattr"); test.AssertErr(t, obj.Error()) {
		if namednodemap, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object NamedNodeMap]", namednodemap.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "Attr"},
	{"method": "GetNamedItem", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "Attr"},
	{"method": "GetNamedItemNS", "args": []interface{}{"name", "high"}, "type": "constructnamechecking", "resultattempt": "Attr"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("listattr"); test.AssertErr(t, obj.Error()) {

		if namednodemap, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, namednodemap, result)
			}

		}

	}
}

func TestSetNamedItem(t *testing.T) {
	js.Eval(`b1=document.createElement("button")
	listattr1=b1.attributes
	attr1=document.createAttribute("hello");

	`)
	if obj := js.Global().Get("listattr1"); test.AssertErr(t, obj.Error()) {

		if namednodemap1, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if objattr := js.Global().Get("attr1"); test.AssertErr(t, objattr.Error()) {

				if attr, err := attr.NewFromJSObject(objattr); test.AssertErr(t, err) {

					test.AssertErr(t, namednodemap1.SetNamedItem(attr))
					if item, err := namednodemap1.Item(0); test.AssertErr(t, err) {

						test.AssertExpect(t, "[object Attr]", item.ToString_())
					}

				}

			}

		}
	}

}

func TestRemoveNamedItem(t *testing.T) {
	js.Eval(`br=document.createElement("button")
	br.setAttribute("hello","world")
	listattrr=br.attributes
	`)
	if obj := js.Global().Get("listattrr"); test.AssertErr(t, obj.Error()) {

		if namednodemap, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertErr(t, namednodemap.RemoveNamedItem("hello"))
			_, err := namednodemap.Item(0)
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))

		}
	}

}

func TestSetNamedItemNS(t *testing.T) {
	js.Eval(`bns1=document.createElement("button")
	listattrns1=bns1.attributes
	attrns1=document.createAttributeNS("namespace","hello")

	`)
	if obj := js.Global().Get("listattrns1"); test.AssertErr(t, obj.Error()) {

		if namednodemap, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if objattr := js.Global().Get("attrns1"); test.AssertErr(t, objattr.Error()) {

				if attr, err := attr.NewFromJSObject(objattr); test.AssertErr(t, err) {

					test.AssertErr(t, namednodemap.SetNamedItemNS(attr))
					if item, err := namednodemap.Item(0); test.AssertErr(t, err) {

						test.AssertExpect(t, "[object Attr]", item.ToString_())
					}

				}

			}

		}
	}

}

func TestRemoveNamedItemNS(t *testing.T) {
	js.Eval(`brns=document.createElement("button")
	brns.setAttributeNS("namespace","hello","world")
	listattrrns=brns.attributes
	`)
	if obj := js.Global().Get("listattrrns"); test.AssertErr(t, obj.Error()) {

		if namednodemap, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertErr(t, namednodemap.RemoveNamedItemNS("namespace", "hello"))
			_, err := namednodemap.Item(0)
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))

		}
	}

}
