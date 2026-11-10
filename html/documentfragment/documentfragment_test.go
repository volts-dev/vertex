package documentfragment

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`df=new DocumentFragment()`)

	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if o, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DocumentFragment]", o.ToString())

		}
	}

}

func TestNew(t *testing.T) {

	js.Eval(`df=new DocumentFragment()`)
	if d, err := New(); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object DocumentFragment]", d.ToString())

	}
}

func TestChildElementCount(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	df.append(div)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, d.ChildElementCount())

		}
	}

}

func TestChildren(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.Children(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLCollection]", c.ToString())
			}

		}
	}

}

func TestFirstElementChild(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.FirstElementChild(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())
			}

		}
	}

}

func TestLastElementChild(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.LastElementChild(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLSpanElement]", c.ToString())
			}

		}
	}

}

func TestQueryAppend(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	b=document.createElement("b")
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if e, err := element.NewFromJSObject(divobj); test.AssertErr(t, err) {
					test.AssertErr(t, d.Append(e))

					test.AssertExpect(t, 1, d.ChildElementCount())

					span := js.Global().Get("span")
					b := js.Global().Get("b")
					es, _ := element.NewFromJSObject(span)
					eb, _ := element.NewFromJSObject(b)

					test.AssertErr(t, d.Append(es, eb))

					test.AssertExpect(t, 3, d.ChildElementCount())

					if c, err := d.FirstElementChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())
					}
					if c, err := d.LastElementChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "[object HTMLElement]", c.ToString())
					}

				}

			}

		}
	}

}

func TestQueryPrepend(t *testing.T) {

	js.Eval(`
	df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	b=document.createElement("b")
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if e, err := element.NewFromJSObject(divobj); test.AssertErr(t, err) {
					test.AssertErr(t, d.Prepend(e))

					test.AssertExpect(t, 1, d.ChildElementCount())

					span := js.Global().Get("span")
					b := js.Global().Get("b")
					es, _ := element.NewFromJSObject(span)
					eb, _ := element.NewFromJSObject(b)

					test.AssertErr(t, d.Prepend(es, eb))

					test.AssertExpect(t, 3, d.ChildElementCount())

					if c, err := d.FirstElementChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "[object HTMLSpanElement]", c.ToString())
					}
					if c, err := d.LastElementChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())
					}
				}

			}

		}
	}

}

func TestQuerySelector(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.QuerySelector("div"); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())
			}

			_, err := d.QuerySelector("button")
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))

		}
	}

}

func TestQuerySelectorAll(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.QuerySelectorAll("div"); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object NodeList]", c.ToString())
			}
		}
	}

}

func TestReplaceChild(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	b=document.createElement("b")
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			b := js.Global().Get("b")

			eb, _ := element.NewFromJSObject(b)

			if c, err := d.FirstElementChild(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())

				if old, err := d.ReplaceChild(eb.Node, c.Node); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object HTMLDivElement]", old.ToString())
					if c2, err := d.FirstElementChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "[object HTMLElement]", c2.ToString())

					}
				}

			}

		}
	}

}
func TestGetElementById(t *testing.T) {

	js.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	div.id="test"
	df.append(div)
	span.id="test2"
	df.append(span)
	
	`)
	if obj := js.Global().Get("df"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			if c, err := d.GetElementById("test"); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", c.ToString())
			}
			if c, err := d.GetElementById("test2"); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLSpanElement]", c.ToString())
			}
			_, err := d.GetElementById("unknown")
			test.AssertExpect(t, true, errors.Is(err, js.ErrUndefinedValue))
		}
	}

}
