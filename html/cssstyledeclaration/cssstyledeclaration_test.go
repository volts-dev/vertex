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

			test.AssertExpect(t, "[object CSSStyleDeclaration]", o.ToString())

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
				test.AssertExpect(t, "[object CSSStyleRule]", p.ToString())

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

			test.AssertExpect(t, "red", o.GetPropertyValue("color"))

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

				test.AssertExpect(t, "blue", o.GetPropertyValue("color"))

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

			test.AssertExpect(t, "color", o.Item(0))

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

			test.AssertExpect(t, "important", o.GetPropertyPriority("color"))

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

			test.AssertExpect(t, "red", o.RemoveProperty("color"))
			test.AssertExpect(t, "", o.GetPropertyValue("color"))

		}
	}

}
