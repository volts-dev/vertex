package node

import (
	"sync"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/helper"
	"github.com/volts-dev/vertex/js/object"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {
	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var nodeinterface js.Value

// GetInterface Get the js node interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if nodeinterface = js.Global().Get("Node"); nodeinterface.Error() != nil {
			nodeinterface = js.Undefined()
		}
		js.Register(nodeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return nodeinterface
}

type Node struct {
	eventtarget.EventTarget
}

type NodeFrom interface {
	Node_() Node
}

func (n Node) Node_() Node {
	return n
}

func NewFromJSObject(obj js.Value) (Node, error) {
	var n Node
	var err error

	if ni := GetInterface(); !ni.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ni) {
				n.SetObjectValue(obj)

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

	if nodeObject = n.GetValueByKey(attribute); nodeObject.Error() == nil {

		if nodeObject.IsUndefined() {
			err = ErrNodeNoChilds

		} else {

			newNode, err = NewFromJSObject(nodeObject)

		}

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

func (n Node) NodValue() (interface{}, error) {

	var err error
	var obj js.Value
	var v interface{}

	if obj = n.GetValueByKey("nodeValue"); obj.Error() == nil {
		v, err = helper.GoValue(obj)
	}

	return v, err
}

func (n Node) SetNodeValue(i interface{}) error {
	n.SetValueByKey("nodeValue", js.ValueOf(i))
	return nil
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

func (n Node) SetTextContent(content string) error {

	n.SetValueByKey("textContent", js.ValueOf(content))
	return n.GetObjectValue().Error()
}

func (n Node) AppendChild(add Node) error {

	err := n.Call("appendChild", add.GetObjectValue()).Error()
	return err
}

func (n Node) CloneNode(deep bool) (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("cloneNode", js.ValueOf(deep)); obj.Error() == nil {
		return NewFromJSObject(obj)
	}

	return newNode, err
}

func (n Node) CompareDocumentPosition(node Node) (int, error) {
	var err error
	var obj js.Value
	var result int

	if obj = n.Call("compareDocumentPosition", node.GetObjectValue()); obj.Error() == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		} else {
			err = object.ErrObjectNotNumber
		}
	}
	return result, err

}

func (n Node) Contains(node Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = n.Call("contains", node.GetObjectValue()); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err
}

func (n Node) GetRootNode() (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj = n.Call("getRootNode"); obj.Error() == nil {
		newNode, err = NewFromJSObject(obj)
	}
	return newNode, err
}

func (n Node) HasChildNodes() (bool, error) {
	return n.CallBool("hasChildNodes")
}

func (n Node) InsertBefore(elem, before Node) (Node, error) {
	var err error

	err = n.Call("insertBefore", elem.GetObjectValue(), before.GetObjectValue()).Error()

	return elem, err

}

func (n Node) IsDefaultNamespace(namespace string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isDefaultNamespace", js.ValueOf(namespace)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) IsEqualNode(n1 Node) (bool, error) {

	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isEqualNode", n1.GetObjectValue()); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) IsSameNode(n1 Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = n.Call("isSameNode", n1.GetObjectValue()); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) LookupPrefix(prefix string) (string, error) {
	var err error
	var obj js.Value
	var result string

	if obj = n.Call("lookupPrefix", js.ValueOf(prefix)); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		}
	}

	return result, err

}

func (n Node) LookupNamespaceURI(prefix string) error {
	var err error
	err = n.Call("lookupNamespaceURI", js.ValueOf(prefix)).Error()
	return err
}

func (n *Node) Normalize() error {
	var err error
	err = n.Call("normalize").Error()
	return err
}

func (n Node) RemoveChild(node Node) (Node, error) {
	var err error
	err = n.Call("removeChild", node.GetObjectValue()).Error()
	return node, err

}

func (n Node) ReplaceChild(new, old Node) (Node, error) {
	var err error

	err = n.Call("replaceChild", new.GetObjectValue(), old.GetObjectValue()).Error()

	return old, err

}
