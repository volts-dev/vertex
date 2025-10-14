package node

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`node= document.createElement("title")
	node.text="hello"
	div=document.createElement("div")
	node.appendChild(div)
	span=document.createElement("span")
	span.appendChild(node)
	br=document.createElement("br")
	span.appendChild(br)
	document.body.appendChild(span)


	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if ti, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLTitleElement", ti.ConstructName_())
		}

	}

}

func TestNodeValue(t *testing.T) {
	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if c, err := node.FirstChild(); test.AssertErr(t, err) {
				if v, err := c.NodValue(); test.AssertErr(t, err) {
					test.AssertExpect(t, "hello", v)
				}

			}
		}

	}

}

func TestSetNodeValue(t *testing.T) {

	js.Eval(`nodeset= document.createElement("title")
	nodeset.text="hello"
	`)
	if obj := js.Global().Get("nodeset"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if c, err := node.FirstChild(); test.AssertErr(t, err) {

				test.AssertErr(t, c.SetNodeValue("world"))

				if v, err := c.NodValue(); test.AssertErr(t, err) {
					test.AssertExpect(t, "world", v)
				}

			}
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "BaseURI", "type": "contains", "resultattempt": "http://localhost"},
	{"method": "FirstChild", "type": "constructnamechecking", "resultattempt": "Text"},
	{"method": "IsConnected", "resultattempt": true},
	{"method": "LastChild", "type": "constructnamechecking", "resultattempt": "HTMLDivElement"},
	{"method": "NextSibling", "type": "constructnamechecking", "resultattempt": "HTMLBRElement"},
	{"method": "NodeName", "resultattempt": "TITLE"},
	{"method": "NodeType", "resultattempt": 1},
	{"method": "OwnerDocument", "type": "constructnamechecking", "resultattempt": "HTMLDocument"},
	{"method": "ParentNode", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "ParentElement", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "TextContent", "resultattempt": "hello"},
	{"method": "SetTextContent", "args": []interface{}{"mytitle"}, "gettermethod": "TextContent", "resultattempt": "mytitle"},
	{"method": "GetRootNode", "type": "constructnamechecking", "resultattempt": "HTMLDocument"},
	{"method": "IsDefaultNamespace", "args": []interface{}{"none"}, "resultattempt": false},
	{"method": "LookupPrefix", "args": []interface{}{"none"}, "resultattempt": ""},
	{"method": "LookupNamespaceURI", "args": []interface{}{"none"}, "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, node, result)
			}

		}

	}
}

func TestPreviousSibling(t *testing.T) {

	if obj := js.Global().Get("br"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if c, err := node.PreviousSibling(); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTitleElement", c.ConstructName_())

			}
		}

	}

}

func TestAppendChild(t *testing.T) {

	js.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
`)
	if obj := js.Global().Get("appendnode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if div, err := NewFromJSObject(divobj); test.AssertErr(t, err) {

					test.AssertErr(t, node.AppendChild(div))

					if c, err := node.FirstChild(); test.AssertErr(t, err) {
						test.AssertExpect(t, "HTMLDivElement", c.ConstructName_())

					}

				}
			}

		}

	}

}

func TestCloneNode(t *testing.T) {

	js.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"
`)
	if obj := js.Global().Get("clonenode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); test.AssertErr(t, err) {

				test.AssertExpect(t, "HTMLTitleElement", clone.ConstructName_())
				test.AssertErr(t, clone.SetTextContent("world"))

				v1, _ := clone.TextContent()
				v2, _ := node.TextContent()
				test.AssertExpect(t, true, v1 != v2)

			}

		}

	}

}

func TestCompareDocumentPosition(t *testing.T) {

	js.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj := js.Global().Get("appendnode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if div, err := NewFromJSObject(divobj); test.AssertErr(t, err) {

					if n, err := node.CompareDocumentPosition(div); test.AssertErr(t, err) {
						test.AssertExpect(t, 20, n)

					}

				}
			}

		}

	}

}

func TestContains(t *testing.T) {

	js.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj := js.Global().Get("appendnode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if div, err := NewFromJSObject(divobj); test.AssertErr(t, err) {

					if b, err := node.Contains(div); test.AssertErr(t, err) {
						test.AssertExpect(t, true, b)

					}

					node.RemoveChild(div)

					if b, err := node.Contains(div); test.AssertErr(t, err) {
						test.AssertExpect(t, false, b)

					}
				}
			}

		}

	}

}

func TestHasChildNodes(t *testing.T) {

	js.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj := js.Global().Get("appendnode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if div, err := NewFromJSObject(divobj); test.AssertErr(t, err) {

					if b, err := node.HasChildNodes(); test.AssertErr(t, err) {
						test.AssertExpect(t, true, b)

					}
					node.RemoveChild(div)

					if b, err := node.HasChildNodes(); test.AssertErr(t, err) {
						test.AssertExpect(t, false, b)

					}
				}
			}

		}

	}

}

func TestInsertBefore(t *testing.T) {

	js.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	span=document.createElement("span")
	appendnode.appendChild(div)

`)
	if obj := js.Global().Get("appendnode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if divobj := js.Global().Get("div"); test.AssertErr(t, divobj.Error()) {

				if div, err := NewFromJSObject(divobj); test.AssertErr(t, err) {

					if spanobj := js.Global().Get("span"); test.AssertErr(t, spanobj.Error()) {

						if span, err := NewFromJSObject(spanobj); test.AssertErr(t, err) {

							if n, err := node.InsertBefore(span, div); test.AssertErr(t, err) {

								test.AssertExpect(t, "HTMLSpanElement", n.ConstructName_())

								if next, err := node.FirstChild(); test.AssertErr(t, err) {

									test.AssertExpect(t, "HTMLSpanElement", next.ConstructName_())

								}

							}

						}
					}

				}
			}

		}

	}

}

func TestIsEqualNode(t *testing.T) {

	js.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"

`)
	if obj := js.Global().Get("clonenode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); test.AssertErr(t, err) {

				if b, err := clone.IsEqualNode(node); test.AssertErr(t, err) {

					test.AssertExpect(t, true, b)

				}
				clone.SetTextContent("world")
				if b, err := clone.IsEqualNode(node); test.AssertErr(t, err) {

					test.AssertExpect(t, false, b)

				}

			}

		}

	}

}

func TestIsSameNode(t *testing.T) {

	js.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"

`)
	if obj := js.Global().Get("clonenode"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); test.AssertErr(t, err) {

				if b, err := clone.IsSameNode(node); test.AssertErr(t, err) {

					test.AssertExpect(t, false, b)

				}
				if b, err := node.IsSameNode(node); test.AssertErr(t, err) {

					test.AssertExpect(t, true, b)

				}

			}

		}

	}

}

func TestNormalize(t *testing.T) {
	js.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
`)
	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if text2obj := js.Global().Get("text1"); test.AssertErr(t, text2obj.Error()) {

				if text1, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

					if text2obj := js.Global().Get("text2"); test.AssertErr(t, text2obj.Error()) {

						if text2, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)

							test.AssertErr(t, node.Normalize())

							if textn, err := node.FirstChild(); test.AssertErr(t, err) {

								test.AssertExpect(t, "-Partie 1 Partie 2 -", textn.TextContent_())

							}

						}
					}

				}
			}

		}
	}

}

func TestRemoveChild(t *testing.T) {
	js.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
`)
	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if text2obj := js.Global().Get("text1"); test.AssertErr(t, text2obj.Error()) {

				if text1, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

					if text2obj := js.Global().Get("text2"); test.AssertErr(t, text2obj.Error()) {

						if text2, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)

							if removetext, err := node.RemoveChild(text1); test.AssertErr(t, err) {

								test.AssertExpect(t, "Text", removetext.ConstructName_())

								if n, err := node.FirstChild(); test.AssertErr(t, err) {

									test.AssertExpect(t, "Partie 2 -", n.TextContent_())

								}

							}

						}
					}

				}
			}

		}
	}

}

func TestReplaceChild(t *testing.T) {
	js.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
	text3=  document.createTextNode("Partie 3 -")
`)
	if obj := js.Global().Get("node"); test.AssertErr(t, obj.Error()) {

		if node, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if text2obj := js.Global().Get("text1"); test.AssertErr(t, text2obj.Error()) {

				if text1, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

					if text2obj := js.Global().Get("text2"); test.AssertErr(t, text2obj.Error()) {

						if text2, err := NewFromJSObject(text2obj); test.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)
							if text3obj := js.Global().Get("text3"); test.AssertErr(t, text3obj.Error()) {

								if text3, err := NewFromJSObject(text3obj); test.AssertErr(t, err) {
									if replace, err := node.ReplaceChild(text3, text1); test.AssertErr(t, err) {

										test.AssertExpect(t, "-Partie 1 ", replace.TextContent_())
										if n, err := node.FirstChild(); test.AssertErr(t, err) {

											test.AssertExpect(t, "Partie 3 -", n.TextContent_())

										}

									}

								}
							}

						}
					}

				}
			}

		}
	}

}
