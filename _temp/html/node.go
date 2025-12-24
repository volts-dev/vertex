package html

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

type (
	// UI is the interface that describes a user interface element such as
	// components and HTML elements.
	___INode interface {
		EventTarget

		// Reports whether the element is mounted.
		Mounted() bool

		Parent() Node
		SetParent(node Node) Node
	}

	Node interface {
		EventTarget
		BaseURI() (string, error)                       //
		ChildNodes() (NodeList, error)                  //
		FirstChild() (Node, error)                      //
		IsConnected() (bool, error)                     //
		LastChild() (Node, error)                       //
		NextSibling() (Node, error)                     //
		NodeName() (string, error)                      //
		NodeType() (int, error)                         //
		NodeValue(value ...any) (any, error)            //
		OwnerDocument() (Node, error)                   //
		ParentNode() (Node, error)                      //
		ParentElement() (Node, error)                   //
		PreviousSibling() (Node, error)                 //
		TextContent(content ...string) (string, error)  //
		AppendChild(add Node)                           //
		CloneNode(deep bool) (Node, error)              //
		CompareDocumentPosition(node Node) (int, error) //
		Contains(node Node) (bool, error)               //
		GetRootNode() (Node, error)                     //
		HasChildNodes() (bool, error)                   //
		InsertBefore(elem, before Node) (Node, error)   //
		IsDefaultNamespace(namespace string) (bool, error)
		IsEqualNode(n1 Node) (bool, error)
		IsSameNode(n1 Node) (bool, error)
		LookupPrefix(prefix string) (string, error)
		LookupNamespaceURI(prefix string)
		Normalize()
		RemoveChild(node Node) Node
		ReplaceChild(new, old Node) (Node, error)
	}

	node struct {
		eventTarget
	}
)

func init() {
	js.RegisterInterface(NodeInterface)
}

var nodeinterface js.Value

// GetInterface Get the js node interface
func NodeInterface() js.Value {
	sync.OnceFunc(func() {
		if nodeinterface = js.Global().Get("Node"); nodeinterface.IsUndefined() {
			nodeinterface = js.Undefined()
		}

		js.Register(nodeinterface, func(v js.Value) (interface{}, error) {
			return ToNode(v)
		})
	})

	return nodeinterface
}

func ToNode(obj js.Value) (Node, error) {
	var n Node
	var err error

	if ni := NodeInterface(); !ni.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ni) {
				n.SetValue(obj)
			} else {
				err = errors.ErrNotANode
			}
		}

	} else {
		err = errors.ErrNotImplemented
	}

	return n, err
}

func (n node) getAttributeNode(attribute string) (Node, error) {
	var nodeObject js.Value
	var newNode Node
	var err error

	nodeObject = n.Get(attribute)
	if nodeObject.IsUndefined() {
		err = errors.ErrNodeNoChilds

	} else {
		newNode, err = ToNode(nodeObject)

	}

	return newNode, err
}

func (n node) BaseURI() (string, error) {
	return n.GetAttributeString("baseURI")
}

func (h HTMLElement) ChildNodes() (NodeList, error) {
	val := h.Get("childNodes")
	return ToNodeList(val)
}

func (n node) FirstChild() (Node, error) {
	return n.getAttributeNode("firstChild")
}

func (n node) IsConnected() (bool, error) {
	return n.GetAttributeBool("isConnected")
}

func (n node) LastChild() (Node, error) {
	return n.getAttributeNode("lastChild")
}

func (n node) NextSibling() (Node, error) {
	return n.getAttributeNode("nextSibling")
}

func (n node) NodeName() (string, error) {
	return n.GetAttributeString("nodeName")
}

func (n node) NodeType() (int, error) {
	return n.GetAttributeInt("nodeType")
}

func (n node) NodeValue() (interface{}, error) {
	var err error
	var obj js.Value
	var v interface{}

	if obj = n.Get("nodeValue"); !obj.IsUndefined() {
		v, err = js.GoValue(obj)
	}

	return v, err
}

func (n node) SetNodeValue(i interface{}) {
	n.Set("nodeValue", js.ValueOf(i))
}

func (n node) OwnerDocument() (Node, error) {
	return n.getAttributeNode("ownerDocument")
}

func (n node) ParentNode() (Node, error) {
	return n.getAttributeNode("parentNode")
}

func (n node) ParentElement() (Node, error) {
	return n.getAttributeNode("parentElement")
}

func (n node) PreviousSibling() (Node, error) {
	return n.getAttributeNode("previousSibling")
}

func (n node) TextContent() (string, error) {
	return n.GetAttributeString("textContent")
}

func (n node) SetTextContent(content string) {
	n.Set("textContent", js.ValueOf(content))
}

func (n node) AppendChild(add Node) {
	n.Call("appendChild", add.Value())
}

func (n node) CloneNode(deep bool) (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("cloneNode", js.ValueOf(deep)); !obj.IsUndefined() {
		return ToNode(obj)
	}

	return newNode, err
}

func (n node) CompareDocumentPosition(node Node) (int, error) {
	var err error
	var obj js.Value
	var result int

	if obj = n.Call("compareDocumentPosition", node.Value()); !obj.IsUndefined() {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		} else {
			err = js.ErrObjectNotNumber
		}
	}
	return result, err

}

func (n node) Contains(node Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = n.Call("contains", node.Value()); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (n node) GetRootNode() (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("getRootNode"); !obj.IsUndefined() {
		newNode, err = ToNode(obj)
	}
	return newNode, err
}

func (n node) HasChildNodes() (bool, error) {
	return n.CallBool("hasChildNodes")
}

func (n node) InsertBefore(elem, before Node) (Node, error) {
	var err error
	n.Call("insertBefore", elem.Value(), before.Value())
	return elem, err
}

func (n node) IsDefaultNamespace(namespace string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isDefaultNamespace", js.ValueOf(namespace)); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (n node) IsEqualNode(n1 Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isEqualNode", n1.Value()); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err

}

func (n node) IsSameNode(n1 Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isSameNode", n1.Value); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err

}

func (n node) LookupPrefix(prefix string) (string, error) {
	var err error
	var obj js.Value
	var result string

	if obj = n.Call("lookupPrefix", js.ValueOf(prefix)); !obj.IsUndefined() {
		if obj.Type() == js.TypeString {
			return obj.String()
		}
	}

	return result, err

}

func (n node) LookupNamespaceURI(prefix string) {
	n.Call("lookupNamespaceURI", js.ValueOf(prefix))

}

func (n node) Normalize() {
	n.Call("normalize")
}

func (n node) RemoveChild(node Node) Node {
	n.Call("removeChild", node.Value)
	return node
}

func (n node) ReplaceChild(new, old Node) (Node, error) {
	var err error
	n.Call("replaceChild", new.Value, old.Value)
	return old, err

}
