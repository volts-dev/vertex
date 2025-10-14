package element

import (
	"testing"

	"github.com/volts-dev/vertex/html/attr"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`
	elementspan= document.createElement("span")
	document.body.appendChild(elementspan)
	element= document.createElement("title")
	element.setAttribute("hello","world")
	listattr=element.attributes
	attr=listattr.item(0)
	div=document.createElement("div")
	div.id="pouet"
	element.appendChild(div)
	collection=element.children
	document.body.appendChild(element)
	element2= document.createElement("br")
	document.body.appendChild(element2)
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("element"); test.AssertErr(t, obj.Error()) {

		if e, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTitleElement", e.ConstructName_())
		}

	}

}

func TestItemFromHTMLCollection(t *testing.T) {

	if obj := js.Global().Get("collection"); test.AssertErr(t, obj.Error()) {

		if c, err := htmlcollection.NewFromJSObject(obj); test.AssertErr(t, err) {

			if e, err := ItemFromHTMLCollection(c, 0); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLDivElement", e.ConstructName_())

			}
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Attributes", "type": "constructnamechecking", "resultattempt": "NamedNodeMap"},
	{"method": "ChildElementCount", "resultattempt": 1},
	{"method": "Children", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "ClassList", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "ClassName", "resultattempt": ""},
	{"method": "SetClassName", "args": []interface{}{"n2"}, "gettermethod": "ClassName", "resultattempt": "n2"},
	{"method": "ClientHeight", "resultattempt": 0},
	{"method": "ClientLeft", "resultattempt": 0},
	{"method": "ClientTop", "resultattempt": 0},
	{"method": "ClientWidth", "resultattempt": 0},
	{"method": "ID", "resultattempt": ""},

	{"method": "QuerySelector", "args": []interface{}{"#pouet"}, "type": "constructnamechecking", "resultattempt": "HTMLDivElement"},
	{"method": "QuerySelectorAll", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "NodeList"},

	{"method": "SetID", "args": []interface{}{"test"}, "gettermethod": "ID", "resultattempt": "test"},
	{"method": "InnerHTML", "resultattempt": "<div id=\"pouet\"></div>"},
	{"method": "SetInnerHTML", "args": []interface{}{"test"}, "gettermethod": "InnerHTML", "resultattempt": "test"},
	{"method": "LocalName", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "NamespaceURI", "resultattempt": "http://www.w3.org/1999/xhtml"},
	{"method": "NextElementSibling", "type": "constructnamechecking", "resultattempt": "HTMLBRElement"},
	{"method": "PreviousElementSibling", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "TagName", "resultattempt": "TITLE"},
	{"method": "Prefix", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "ScrollHeight", "resultattempt": 0},
	{"method": "SetScrollHeight", "args": []interface{}{0}, "gettermethod": "ScrollHeight", "resultattempt": 0},
	{"method": "ScrollLeft", "resultattempt": 0},
	{"method": "SetScrollLeft", "args": []interface{}{0}, "gettermethod": "ScrollLeft", "resultattempt": 0},
	{"method": "ScrollTop", "resultattempt": 0},
	{"method": "SetScrollTop", "args": []interface{}{0}, "gettermethod": "ScrollTop", "resultattempt": 0},
	{"method": "ScrollWidth", "resultattempt": 0},
	{"method": "SetScrollWidth", "args": []interface{}{0}, "gettermethod": "ScrollWidth", "resultattempt": 0},
	{"method": "Closest", "args": []interface{}{"body"}, "type": "constructnamechecking", "resultattempt": "HTMLBodyElement"},
	{"method": "GetAttribute", "args": []interface{}{"hello"}, "resultattempt": "world"},
	{"method": "GetAttributeNS", "args": []interface{}{"name", "hello"}, "type": "error", "resultattempt": ErrAttributeEmpty},
	{"method": "GetAttributeNames", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "GetBoundingClientRect", "type": "constructnamechecking", "resultattempt": "DOMRect"},
	{"method": "GetClientRects", "type": "constructnamechecking", "resultattempt": "DOMRectList"},
	{"method": "GetElementsByClassName", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "GetElementsByTagName", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "GetElementsByTagNameNS", "args": []interface{}{"namespace", "div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "HasAttribute", "args": []interface{}{"hello"}, "resultattempt": true},
	{"method": "HasPointerCapture", "args": []interface{}{0}, "resultattempt": false},
	{"method": "Matches", "args": []interface{}{"#test"}, "resultattempt": true},

	{"method": "OuterHTML", "resultattempt": "<title hello=\"world\" class=\"n2\" id=\"test\">test</title>"},
	{"method": "SetOuterHTML", "args": []interface{}{"<title helloZ=\"world\" class=\"n2\" id=\"test\">test</title>"}, "gettermethod": "OuterHTML", "resultattempt": "<title hello=\"world\" class=\"n2\" id=\"test\">test</title>"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("element"); test.AssertErr(t, obj.Error()) {

		if elem, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, elem, result)
			}

		}

	}
}
func TestOwnerElementForAttr(t *testing.T) {

	if objattr := js.Global().Get("attr"); test.AssertErr(t, objattr.Error()) {

		if attr, err := attr.NewFromJSObject(objattr); test.AssertErr(t, err) {

			if elem, err := OwnerElementForAttr(attr); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTitleElement", elem.ConstructName_())

			}
		}
	}

}

func TestAfter(t *testing.T) {

	js.Eval(`
	elementspanaf= document.createElement("span")
	elementdivaf= document.createElement("div")
	elementspanaf.appendChild(elementdivaf)
	elementbraf= document.createElement("br")

	`)

	if objspan := js.Global().Get("elementspanaf"); test.AssertErr(t, objspan.Error()) {

		if span, err := NewFromJSObject(objspan); test.AssertErr(t, err) {

			if objdiv := js.Global().Get("elementdivaf"); test.AssertErr(t, objdiv.Error()) {

				if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {
					if objbr := js.Global().Get("elementbraf"); test.AssertErr(t, objbr.Error()) {

						if br, err := NewFromJSObject(objbr); test.AssertErr(t, err) {

							test.AssertErr(t, div.After(br))

							if val, err := span.OuterHTML(); test.AssertErr(t, err) {
								test.AssertExpect(t, "<span><div></div><br></span>", val)

							}

						}
					}

				}
			}

		}
	}

}

func TestAppend(t *testing.T) {

	js.Eval(`
	elementspanap= document.createElement("span")
	elementdivap= document.createElement("div")

	`)

	if objspan := js.Global().Get("elementspanap"); test.AssertErr(t, objspan.Error()) {

		if span, err := NewFromJSObject(objspan); test.AssertErr(t, err) {

			if objdiv := js.Global().Get("elementdivap"); test.AssertErr(t, objdiv.Error()) {

				if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {

					test.AssertErr(t, span.AppendChild(div.Node))

					if er, err := span.FirstChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "HTMLDivElement", er.ConstructName_())

					}

				}
			}
		}
	}

}

func TestPrepend(t *testing.T) {

	js.Eval(`
	elementspanap= document.createElement("span")
	elementdivap= document.createElement("div")

	`)

	if objspan := js.Global().Get("elementspanap"); test.AssertErr(t, objspan.Error()) {

		if span, err := NewFromJSObject(objspan); test.AssertErr(t, err) {

			if objdiv := js.Global().Get("elementdivap"); test.AssertErr(t, objdiv.Error()) {

				if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {

					test.AssertErr(t, span.Prepend(div))

					if er, err := span.FirstChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "HTMLDivElement", er.ConstructName_())

					}

				}
			}
		}
	}

}

func TestBefore(t *testing.T) {

	js.Eval(`
	elementspanbf= document.createElement("span")
	elementdivbf= document.createElement("div")
	elementspanbf.appendChild(elementdivbf)
	elementbrbf= document.createElement("br")
	`)

	if objspan := js.Global().Get("elementspanbf"); test.AssertErr(t, objspan.Error()) {

		if span, err := NewFromJSObject(objspan); test.AssertErr(t, err) {

			if objdiv := js.Global().Get("elementdivbf"); test.AssertErr(t, objdiv.Error()) {

				if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {
					if objbr := js.Global().Get("elementbrbf"); test.AssertErr(t, objbr.Error()) {

						if br, err := NewFromJSObject(objbr); test.AssertErr(t, err) {

							test.AssertErr(t, div.Before(br))

							if val, err := span.OuterHTML(); test.AssertErr(t, err) {
								test.AssertExpect(t, "<span><br><div></div></span>", val)

							}

						}
					}

				}
			}

		}
	}

}

func TestInsertAdjacentElement(t *testing.T) {

	js.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp := js.Global().Get("elementp"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			if objbr := js.Global().Get("elementbri"); test.AssertErr(t, objbr.Error()) {

				if br, err := NewFromJSObject(objbr); test.AssertErr(t, err) {

					if elem, err := p.InsertAdjacentElement("afterbegin", br); test.AssertErr(t, err) {

						test.AssertExpect(t, "[object HTMLBRElement]", elem.ToString_())
						if val, err := p.OuterHTML(); test.AssertErr(t, err) {
							test.AssertExpect(t, "<p><br>hello</p>", val)

						}

					}

				}
			}

		}
	}

}

func TestInsertAdjacentHTML(t *testing.T) {

	js.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp := js.Global().Get("elementp"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.InsertAdjacentHTML("afterbegin", "<div>test</div>"))

			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p><div>test</div>hello</p>", val)

			}

		}
	}

}

func TestInsertAdjacentText(t *testing.T) {

	js.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp := js.Global().Get("elementp"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.InsertAdjacentText("beforeend", "this is text <br>"))

			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p>hellothis is text &lt;br&gt;</p>", val)

			}

		}
	}

}

func TestRemove(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")
	div= document.createElement("div")
	p.appendChild(div)

	`)

	if objdiv := js.Global().Get("div"); test.AssertErr(t, objdiv.Error()) {

		if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {

			if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

				if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

					test.AssertErr(t, div.Remove())
					if val, err := p.OuterHTML(); test.AssertErr(t, err) {
						test.AssertExpect(t, "<p></p>", val)

					}

				}
			}

		}
	}

}

func TestRemoveAttribute(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")
	p.setAttribute("hello","world")
	p.setAttribute("hello1","world1")
	`)

	if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.RemoveAttribute("hello"))
			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p hello1=\"world1\"></p>", val)

			}

		}
	}

}

func TestRemoveAttributeNS(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")
	p.setAttribute("hello","world")
	p.setAttribute("hello1","world1")
	p.setAttributeNS("space","hello","world")
	`)

	if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.RemoveAttributeNS("space", "hello"))
			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p hello=\"world\" hello1=\"world1\"></p>", val)

			}

		}
	}

}

func TestReplaceChildren(t *testing.T) {
	js.Eval(`
	div= document.createElement("div")
	pTemp = document.createElement("p")
	pTemp.innerText = "remove me"
	div.append(pTemp)
	span = document.createElement("span")
	span.innerText = "done"
	`)

	if objdiv := js.Global().Get("div"); test.AssertErr(t, objdiv.Error()) {
		if div, err := NewFromJSObject(objdiv); test.AssertErr(t, err) {
			if objspan := js.Global().Get("span"); test.AssertErr(t, objspan.Error()) {
				if span, err := NewFromJSObject(objspan); test.AssertErr(t, err) {
					test.AssertErr(t, div.ReplaceChildren("well ", span.Node))
					if val, err := div.OuterHTML(); test.AssertErr(t, err) {
						test.AssertExpect(t, "<div>well <span>done</span></div>", val)
					}
				}
			}
		}
	}
}

func TestSetAttribute(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")

	`)

	if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.SetAttribute("space", "hello"))
			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p space=\"hello\"></p>", val)

			}

		}
	}

}

func TestSetAttributeNS(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")

	`)

	if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			test.AssertErr(t, p.SetAttributeNS("space", "hello", "world"))
			if val, err := p.OuterHTML(); test.AssertErr(t, err) {
				test.AssertExpect(t, "<p hello=\"world\"></p>", val)

			}

		}
	}

}

func TestToggleAttribute(t *testing.T) {

	js.Eval(`
	p= document.createElement("p")
	`)

	if objp := js.Global().Get("p"); test.AssertErr(t, objp.Error()) {

		if p, err := NewFromJSObject(objp); test.AssertErr(t, err) {

			if b, err := p.ToggleAttribute("disabled", true); test.AssertErr(t, err) {

				test.AssertExpect(t, true, b)
				if val, err := p.OuterHTML(); test.AssertErr(t, err) {
					test.AssertExpect(t, "<p disabled=\"\"></p>", val)

				}
			}

		}
	}

}

var methodsAttemptAccessibility []map[string]interface{} = []map[string]interface{}{

	{"method": "AriaAtomic", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaAtomic", "args": []interface{}{"true"}, "gettermethod": "AriaAtomic", "resultattempt": "true"},

	{"method": "AriaAutoComplete", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaAutoComplete", "args": []interface{}{"inline"}, "gettermethod": "AriaAutoComplete", "resultattempt": "inline"},

	{"method": "AriaBusy", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaBusy", "args": []interface{}{"true"}, "gettermethod": "AriaBusy", "resultattempt": "true"},

	{"method": "AriaChecked", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaChecked", "args": []interface{}{"true"}, "gettermethod": "AriaChecked", "resultattempt": "true"},

	{"method": "AriaColCount", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaColCount", "args": []interface{}{"2"}, "gettermethod": "AriaColCount", "resultattempt": "2"},

	{"method": "AriaColIndex", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaColIndex", "args": []interface{}{"1"}, "gettermethod": "AriaColIndex", "resultattempt": "1"},

	{"method": "AriaColIndexText", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaColIndexText", "args": []interface{}{"11"}, "gettermethod": "AriaColIndexText", "resultattempt": "11"},

	{"method": "AriaColSpan", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaColSpan", "args": []interface{}{"1"}, "gettermethod": "AriaColSpan", "resultattempt": "1"},

	{"method": "AriaCurrent", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaCurrent", "args": []interface{}{"page"}, "gettermethod": "AriaCurrent", "resultattempt": "page"},

	{"method": "AriaDescription", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaDescription", "args": []interface{}{"test"}, "gettermethod": "AriaDescription", "resultattempt": "test"},

	{"method": "AriaDisabled", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaDisabled", "args": []interface{}{"true"}, "gettermethod": "AriaDisabled", "resultattempt": "true"},

	{"method": "AriaExpanded", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaExpanded", "args": []interface{}{"true"}, "gettermethod": "AriaExpanded", "resultattempt": "true"},

	{"method": "AriaHasPopup", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaHasPopup", "args": []interface{}{"true"}, "gettermethod": "AriaHasPopup", "resultattempt": "true"},

	{"method": "AriaHidden", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaHidden", "args": []interface{}{"true"}, "gettermethod": "AriaHidden", "resultattempt": "true"},

	{"method": "AriaKeyShortcuts", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaKeyShortcuts", "args": []interface{}{"true"}, "gettermethod": "AriaKeyShortcuts", "resultattempt": "true"},

	{"method": "AriaLabel", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaLabel", "args": []interface{}{"true"}, "gettermethod": "AriaLabel", "resultattempt": "true"},

	{"method": "AriaLevel", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaLevel", "args": []interface{}{"1"}, "gettermethod": "AriaLevel", "resultattempt": "1"},

	{"method": "AriaLive", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaLive", "args": []interface{}{"assertive"}, "gettermethod": "AriaLive", "resultattempt": "assertive"},

	{"method": "AriaModal", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaModal", "args": []interface{}{"true"}, "gettermethod": "AriaModal", "resultattempt": "true"},

	{"method": "AriaMultiline", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaMultiline", "args": []interface{}{"true"}, "gettermethod": "AriaMultiline", "resultattempt": "true"},

	{"method": "AriaMultiSelectable", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaMultiSelectable", "args": []interface{}{"true"}, "gettermethod": "AriaMultiSelectable", "resultattempt": "true"},

	{"method": "AriaOrientation", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaOrientation", "args": []interface{}{"horizontal"}, "gettermethod": "AriaOrientation", "resultattempt": "horizontal"},

	{"method": "AriaPlaceholder", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaPlaceholder", "args": []interface{}{"true"}, "gettermethod": "AriaPlaceholder", "resultattempt": "true"},

	{"method": "AriaPosInSet", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaPosInSet", "args": []interface{}{"1"}, "gettermethod": "AriaPosInSet", "resultattempt": "1"},

	{"method": "AriaPressed", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaPressed", "args": []interface{}{"true"}, "gettermethod": "AriaPressed", "resultattempt": "true"},

	{"method": "AriaReadOnly", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaReadOnly", "args": []interface{}{"true"}, "gettermethod": "AriaReadOnly", "resultattempt": "true"},

	{"method": "AriaRelevant", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRelevant", "args": []interface{}{"text"}, "gettermethod": "AriaRelevant", "resultattempt": "text"},

	{"method": "AriaRequired", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRequired", "args": []interface{}{"true"}, "gettermethod": "AriaRequired", "resultattempt": "true"},

	{"method": "AriaRoleDescription", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRoleDescription", "args": []interface{}{"test"}, "gettermethod": "AriaRoleDescription", "resultattempt": "test"},

	{"method": "AriaRowCount", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRowCount", "args": []interface{}{"1"}, "gettermethod": "AriaRowCount", "resultattempt": "1"},

	{"method": "AriaRowIndex", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRowIndex", "args": []interface{}{"1"}, "gettermethod": "AriaRowIndex", "resultattempt": "1"},

	{"method": "AriaRowIndexText", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRowIndexText", "args": []interface{}{"1"}, "gettermethod": "AriaRowIndexText", "resultattempt": "1"},

	{"method": "AriaRowSpan", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaRowSpan", "args": []interface{}{"1"}, "gettermethod": "AriaRowSpan", "resultattempt": "1"},

	{"method": "AriaSelected", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaSelected", "args": []interface{}{"true"}, "gettermethod": "AriaSelected", "resultattempt": "true"},

	{"method": "AriaSetSize", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaSetSize", "args": []interface{}{"true"}, "gettermethod": "AriaSetSize", "resultattempt": "true"},

	{"method": "AriaSort", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaSort", "args": []interface{}{"ascending"}, "gettermethod": "AriaSort", "resultattempt": "ascending"},

	{"method": "AriaValueMax", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaValueMax", "args": []interface{}{"9"}, "gettermethod": "AriaValueMax", "resultattempt": "9"},

	{"method": "AriaValueMin", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaValueMin", "args": []interface{}{"2"}, "gettermethod": "AriaValueMin", "resultattempt": "2"},

	{"method": "AriaValueNow", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaValueNow", "args": []interface{}{"2"}, "gettermethod": "AriaValueNow", "resultattempt": "2"},

	{"method": "AriaValueText", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "SetAriaValueText", "args": []interface{}{"2"}, "gettermethod": "AriaValueText", "resultattempt": "2"},
}

func TestMethodsAccessibility(t *testing.T) {

	js.Eval(`
	
	elementaccessibility= document.createElement("title")
	`)

	if obj := js.Global().Get("elementaccessibility"); test.AssertErr(t, obj.Error()) {

		if elem, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttemptAccessibility {
				test.InvokeCheck(t, elem, result)
			}

		}

	}
}
