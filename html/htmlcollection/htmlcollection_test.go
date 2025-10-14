package htmlcollection

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`ch= document.children
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("ch"); test.AssertErr(t, obj.Error()) {

		if c, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLCollection", c.ConstructName_())
		}

	}
}

func TestItem(t *testing.T) {

	if obj := js.Global().Get("ch"); test.AssertErr(t, obj.Error()) {

		if c, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if n, err := c.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "HTMLHtmlElement", n.(js.Object).ConstructName_())

			}
		}

	}
}
