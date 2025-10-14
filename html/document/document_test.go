package document

import (
	"errors"
	"strings"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/element"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestDomain(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if str, err := d.Domain(); test.AssertErr(t, err) {
			test.AssertExpect(t, "localhost", str)
		}

		if err := d.SetDomain("testing.com"); err == nil {
			t.Error("Must return error")
		} else {

			test.AssertExpect(t, "Failed to set the 'domain' property on 'Document': 'testing.com' is not a suffix of 'localhost'.", err.Error())

		}

	}

}

func TestTitle(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		if str, err := d.Title(); test.AssertErr(t, err) {
			test.AssertExpect(t, "Go wasm", str)

		}

		test.AssertErr(t, d.SetTitle("Hello"))

		if str, err := d.Title(); test.AssertErr(t, err) {
			test.AssertExpect(t, "Hello", str)

		}

	}

}

func TestBody(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		if b, err := d.Body(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLBodyElement]", b.ToString_())
		}

	}

}

func TestActiveElement(t *testing.T) {
	if d, err := New(); test.AssertErr(t, err) {
		if el, err := d.ActiveElement(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLBodyElement]", el.ToString_())
		}
	}
}

func TestCharacterSet(t *testing.T) {
	if d, err := New(); test.AssertErr(t, err) {
		if s, err := d.CharacterSet(); test.AssertErr(t, err) {
			test.AssertExpect(t, "UTF-8", s)
		}

	}
}

func TestChildElementCount(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		if c, err := d.ChildElementCount(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, c)
		}

	}
}

func TestChildren(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		if collections, err := d.Children(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLHtmlElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}
}

func TestCompatMode(t *testing.T) {

	//var expectElementString []string = []string{"", ""}
	if d, err := New(); test.AssertErr(t, err) {
		if compatmode, err := d.CompatMode(); test.AssertErr(t, err) {
			test.AssertExpect(t, "CSS1Compat", compatmode)
		}

	}
}

func TestContentType(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if str, err := d.ContentType(); test.AssertErr(t, err) {
			if str != "text/html" {
				t.Errorf("Content Type must be text/html, have %s", str)
			}
		}

	}

}
func TestDocumentElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.DocumentElement(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}

func TestDocumentURI(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if uri, err := d.DocumentURI(); test.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(uri, expect) {
				t.Errorf("Must contain %s have %s", expect, uri)
			}

		}

	}
}

func TestEmbeds(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("emb=document.createElement(\"embed\");document.body.appendChild(emb)")

		if collections, err := d.Embeds(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLEmbedElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("emb.parentNode.removeChild(emb)")
	}
}

func TestFirstElementChild(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.FirstElementChild(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}
func TestForms(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("form=document.createElement(\"form\");document.body.appendChild(form)")

		if collections, err := d.Forms(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLFormElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}

		js.Eval("form.parentNode.removeChild(form)")

	}
}

func TestFullscreenElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.FullscreenElement(); test.AssertErr(t, err) {
			test.AssertExpect(t, "", elem.ToString_())
		}

	}

}

func TestHead(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		if h, err := d.Head(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLHeadElement]", h.ToString_())
		}

	}

}

func TestHidden(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if h, err := d.Hidden(); test.AssertErr(t, err) {
			test.AssertExpect(t, false, h)
		}

	}

}

func TestImages(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("img=document.createElement(\"img\");document.body.appendChild(img)")

		if collections, err := d.Images(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLImageElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("img.parentNode.removeChild(img)")
	}
}

func TestLastElementChild(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.LastElementChild(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}

func TestLinks(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("a=document.createElement(\"a\");a.href=\"testing://localhost\";document.body.appendChild(a)")

		if collections, err := d.Links(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "testing://localhost", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("a.parentNode.removeChild(a)")
	}
}

func TestPictureInPictureElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.PictureInPictureElement(); test.AssertErr(t, err) {
			if !elem.Empty() {
				t.Error("PictureInPicture must be empty")
			}

		}

	}
}

func TestPictureInPictureEnabled(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if ok, err := d.PictureInPictureEnabled(); test.AssertErr(t, err) {
			test.AssertExpect(t, true, ok)
		}

	}
}

func TestPlugins(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("emb=document.createElement(\"embed\");document.body.appendChild(emb)")

		if collections, err := d.Plugins(); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLEmbedElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("emb.parentNode.removeChild(emb)")
	}

}

func TestScripts(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if collections, err := d.Scripts(); test.AssertErr(t, err) {
			test.AssertExpect(t, 2, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLScriptElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}

}
func TestScrollingElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.ScrollingElement(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}
func TestVisibilityState(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if state, err := d.VisibilityState(); test.AssertErr(t, err) {
			test.AssertExpect(t, "visible", state)
		}

	}

}

func TestLastModified(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if state, err := d.LastModified(); test.AssertErr(t, err) {
			if len(state) == 0 {
				t.Errorf("Must have value")
			}
		}

	}

}

func TestReadyState(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if state, err := d.ReadyState(); test.AssertErr(t, err) {
			test.AssertExpect(t, "complete", state)
		}

	}

}

func TestReferrer(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if state, err := d.Referrer(); test.AssertErr(t, err) {
			test.AssertExpect(t, "", state)
		}

	}

}
func TestURL(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if state, err := d.URL(); test.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(state, expect) {
				t.Errorf("Must contain %s have %s", expect, state)
			}
		}

	}

}

func TestCookie(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if str, err := d.Cookie(); test.AssertErr(t, err) {

			test.AssertExpect(t, "", str)

		}
		test.AssertErr(t, d.SetCookie("hello world"))
		if str, err := d.Cookie(); test.AssertErr(t, err) {

			test.AssertExpect(t, "hello world", str)
		}

	}

}

func TestAppend(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if err := d.Append("hello"); err == nil {

			t.Error("Must return an error")

		}
	}

}

func TestCreateAttribute(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if attr, err := d.CreateAttribute("myattr"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Attr]", attr.ToString_())
			if n, err := attr.Name(); test.AssertErr(t, err) {
				test.AssertExpect(t, "myattr", n)
			}

			if n, err := attr.LocalName(); test.AssertErr(t, err) {
				test.AssertExpect(t, "myattr", n)
			}

			if n := attr.GetObjectValue(); test.AssertErr(t, n.Error()) {
				test.AssertExpect(t, "", n)
			}

		}
	}

}

func TestCreateComment(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if comment, err := d.CreateComment("com"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Comment]", comment.ToString_())

		}
	}
}

func TestCreateDocumentFragment(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if fragment, err := d.CreateDocumentFragment(); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DocumentFragment]", fragment.ToString_())

		}
	}

}

func TestCreateHTMLElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.CreateHTMLElement("test"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}

		if elem, err := d.CreateHTMLElement("input"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateHTMLElement("button"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateElement(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.CreateElement("test"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}
		if elem, err := d.CreateElement("input"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateElement("button"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateElementNS(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "test"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}
		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "input"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "button"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateEvent(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if event, err := d.CreateEvent("KeyboardEvent"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object KeyboardEvent]", event.ToString_())

		}

		if _, err := d.CreateEvent("testEvent"); err == nil {

			t.Error("Must return an err")

		}
	}

}

func TestCreateTextNode(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		if textnode, err := d.CreateTextNode("test"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Text]", textnode.ToString_())

		}

	}

}

func TestGetElementsByClassName(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.className=\"test\"")

		if collections, err := d.GetElementsByClassName("test"); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("div.parentNode.removeChild(div)")

		if collections, err := d.GetElementsByClassName("test"); test.AssertErr(t, err) {
			test.AssertExpect(t, 0, collections.Length())

		}
	}

}

func TestGetElementsByTagName(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.className=\"test\"")

		if collection, err := d.GetElementsByTagName("div"); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collection.Length())

			if item, err := collection.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("div.parentNode.removeChild(div)")

		if collection, err := d.GetElementsByTagName("div"); test.AssertErr(t, err) {
			test.AssertExpect(t, 0, collection.Length())

		}
	}

}

func TestGetElementsByTagNameNS(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div)")

		if collection, err := d.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div"); test.AssertErr(t, err) {
			test.AssertExpect(t, 1, collection.Length())

			if item, err := collection.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", item.(js.ObjectFrom).BaseObject_().ToString_())
			}
		}
		js.Eval("div.parentNode.removeChild(div)")

		if collection, err := d.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div"); test.AssertErr(t, err) {
			test.AssertExpect(t, 0, collection.Length())

		}

	}

}

func TestGetElementById(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if item, err := d.GetElementById("testid"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())

		}

		js.Eval("div.parentNode.removeChild(div)")
		if _, err := d.GetElementById("testid"); !errors.Is(err, ErrElementNotFound) {
			t.Errorf("Must return err %s", ErrElementNotFound.Error())
		}

	}

}

func TestQuerySelector(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if item, err := d.QuerySelector("#testid"); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())

		}

		js.Eval("div.parentNode.removeChild(div)")
		if _, err := d.QuerySelector("#testid"); !errors.Is(err, ErrElementNotFound) {
			t.Errorf("Must return err %s", ErrElementNotFound.Error())
		}

	}

}

func TestQuerySelectorAll(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if nodeslist, err := d.QuerySelectorAll("div"); test.AssertErr(t, err) {

			test.AssertExpect(t, 1, nodeslist.Length())
			if item, err := nodeslist.Item(0); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())
			}
		}

		js.Eval("div.parentNode.removeChild(div)")
		if nodeslist, err := d.QuerySelectorAll("divv"); test.AssertErr(t, err) {
			test.AssertExpect(t, 0, nodeslist.Length())
		}

	}

}

func TestImportNode(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {
		js.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if div, err := d.GetElementById("testid"); test.AssertErr(t, err) {

			if clone, err := d.ImportNode(div.Node, true); test.AssertErr(t, err) {
				div.SetID("te")
				if clondelem, err := element.NewFromJSObject(clone.(js.ObjectFrom).Value()); test.AssertErr(t, err) {

					test.AssertExpect(t, "testid", clondelem.ID_())

				}

			}

		}

	}

}
