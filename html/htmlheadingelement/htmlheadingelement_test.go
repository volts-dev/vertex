package htmlheadingelement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`h= document.createElement("h1")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if h, err := NewH1(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH2(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH3(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH4(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH5(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}
		if h, err := NewH6(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("h"); test.AssertErr(t, obj.Error()) {

		if h, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

	}
}
