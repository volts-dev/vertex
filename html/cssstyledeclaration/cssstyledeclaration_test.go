package cssstyledeclaration

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
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object CSSStyleDeclaration]", o.ToString_())

		}
	}

}

func TestParentRule(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if p, err := o.ParentRule(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object CSSStyleRule]", p.ToString_())

			}

		}
	}

}

func TestGetPropertyValue(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := o.GetPropertyValue("color"); test.AssertErr(t, err) {

				test.AssertExpect(t, "red", v)
			}

		}
	}

}

func TestSetProperty(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if err := o.SetProperty("color", "blue"); test.AssertErr(t, err) {

				if v, err := o.GetPropertyValue("color"); test.AssertErr(t, err) {

					test.AssertExpect(t, "blue", v)
				}

			}

		}
	}

}

func TestItem(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := o.Item(0); test.AssertErr(t, err) {

				test.AssertExpect(t, "color", v)

			}

		}
	}

}

func TestGetPropertyPriority(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red!important; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := o.GetPropertyPriority("color"); test.AssertErr(t, err) {

				test.AssertExpect(t, "important", v)

			}

		}
	}

}
func TestRemoveProperty(t *testing.T) {

	js.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj := js.Global().Get("style"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := o.RemoveProperty("color"); test.AssertErr(t, err) {
				test.AssertExpect(t, "red", v)
				if v2, err := o.GetPropertyValue("color"); test.AssertErr(t, err) {

					test.AssertExpect(t, "", v2)
				}

			}

		}
	}

}
