package cssrule

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	prule=document.styleSheets[0].rules[0].style.parentRule
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("prule"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object CSSStyleRule]", o.ToString_())

		}
	}

}

func TestCssText(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	prule=document.styleSheets[0].rules[0].style.parentRule
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("prule"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := o.CssText(); test.AssertErr(t, err) {

				test.AssertExpect(t, "h1 { color: red; font-size: 50px; }", v)
			}

		}
	}

}

func TestParentRule(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "@supports (display: flex) { @media screen and (min-width: 900px) { article { display: flex; } } }"
	document.head.appendChild(s)
	rule=document.styleSheets[0].rules[0].cssRules[0]
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("rule"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if p, err := o.ParentRule(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object CSSSupportsRule]", p.ToString_())

			}

		}
	}

}

func TestParentStyleSheet(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "@supports (display: flex) { @media screen and (min-width: 900px) { article { display: flex; } } }"
	document.head.appendChild(s)
	rule=document.styleSheets[0].rules[0].cssRules[0]
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("rule"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if p, err := o.ParentStyleSheet(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object CSSStyleSheet]", p.ToString_())

			}

		}
	}

}
