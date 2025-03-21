package vertex

import (
	"sync"

	"github.com/volts-dev/vertex/core/js"
)

type (
	Node struct {
		EventListener
	}
)

func init() {
	RegisterInterface(NodeInterface)
}

var nodeinterface js.Value

// GetInterface Get the js node interface
func NodeInterface() js.Value {
	sync.OnceFunc(func() {
		if nodeinterface = js.Global().Get("Node"); nodeinterface.IsUndefined() {
			nodeinterface = js.Undefined()
		}

		Register(nodeinterface, func(v js.Value) (interface{}, error) {
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
			err = ErrUndefinedValue
		} else {
			if obj.InstanceOf(ni) {
				n.SetObject(obj)
			} else {
				err = ErrNotANode
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return n, err
}

func (n Node) getAttributeNode(attribute string) (Node, error) {
	var nodeObject js.Value
	var newNode Node
	var err error

	nodeObject = n.Get(attribute)
	if nodeObject.IsUndefined() {
		err = ErrNodeNoChilds

	} else {
		newNode, err = ToNode(nodeObject)

	}

	return newNode, err
}

func (n Node) BaseURI() (string, error) {
	return n.GetAttributeString("baseURI")
}

func (n Node) FirstChild() (Node, error) {
	return n.getAttributeNode("firstChild")
}

func (n Node) IsConnected() (bool, error) {
	return n.GetAttributeBool("isConnected")
}

func (n Node) LastChild() (Node, error) {
	return n.getAttributeNode("lastChild")
}

func (n Node) NextSibling() (Node, error) {
	return n.getAttributeNode("nextSibling")
}

func (n Node) NodeName() (string, error) {
	return n.GetAttributeString("nodeName")

}

func (n Node) NodeType() (int, error) {
	return n.GetAttributeInt("nodeType")
}

func (n Node) NodeValue() (interface{}, error) {
	var err error
	var obj js.Value
	var v interface{}

	if obj = n.Get("nodeValue"); !obj.IsUndefined() {
		v, err = GoValue(obj)
	}

	return v, err
}

func (n Node) SetNodeValue(i interface{}) {
	n.Set("nodeValue", js.ValueOf(i))
}

func (n Node) OwnerDocument() (Node, error) {
	return n.getAttributeNode("ownerDocument")
}

func (n Node) ParentNode() (Node, error) {
	return n.getAttributeNode("parentNode")
}

func (n Node) ParentElement() (Node, error) {
	return n.getAttributeNode("parentElement")
}

func (n Node) PreviousSibling() (Node, error) {
	return n.getAttributeNode("previousSibling")
}

func (n Node) TextContent() (string, error) {
	return n.GetAttributeString("textContent")
}

func (n Node) SetTextContent(content string) {
	n.Set("textContent", js.ValueOf(content))
}

func (n Node) AppendChild(add Node) {
	n.Call("appendChild", add.JSObject())
}

func (n Node) CloneNode(deep bool) (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("cloneNode", js.ValueOf(deep)); !obj.IsUndefined() {
		return ToNode(obj)
	}

	return newNode, err
}

func (n Node) CompareDocumentPosition(node Node) (int, error) {
	var err error
	var obj js.Value
	var result int

	if obj = n.Call("compareDocumentPosition", node.JSObject()); !obj.IsUndefined() {
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = ErrObjectNotNumber
		}
	}
	return result, err

}

func (n Node) Contains(node Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = n.Call("contains", node.JSObject()); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err
}

func (n Node) GetRootNode() (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("getRootNode"); !obj.IsUndefined() {
		newNode, err = ToNode(obj)
	}
	return newNode, err
}

func (n Node) HasChildNodes() (bool, error) {
	return n.CallBool("hasChildNodes")
}

func (n Node) InsertBefore(elem, before Node) (Node, error) {
	var err error
	n.Call("insertBefore", elem.JSObject(), before.JSObject())
	return elem, err

}

func (n Node) IsDefaultNamespace(namespace string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isDefaultNamespace", js.ValueOf(namespace)); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) IsEqualNode(n1 Node) (bool, error) {

	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isEqualNode", n1.JSObject()); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) IsSameNode(n1 Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isSameNode", n1.JSObject()); !obj.IsUndefined() {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) LookupPrefix(prefix string) (string, error) {
	var err error
	var obj js.Value
	var result string

	if obj = n.Call("lookupPrefix", js.ValueOf(prefix)); !obj.IsUndefined() {
		if obj.Type() == js.TypeString {
			result = obj.String()
		}
	}

	return result, err

}

func (n Node) LookupNamespaceURI(prefix string) {
	n.Call("lookupNamespaceURI", js.ValueOf(prefix))

}

func (n *Node) Normalize() {
	n.Call("normalize")
}

func (n Node) RemoveChild(node Node) Node {
	n.Call("removeChild", node.JSObject())
	return node
}

func (n Node) ReplaceChild(new, old Node) (Node, error) {
	var err error
	n.Call("replaceChild", new.JSObject(), old.JSObject())
	return old, err

}
