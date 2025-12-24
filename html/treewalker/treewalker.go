package treewalker

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
)

var ErrNotImplemented = errors.New("Browser not implemented TreeWalker")
var singleton sync.Once
var treeWalkerInterface js.Value

type (
	TreeWalker struct {
		js.Object
		filter js.Value
	}

	TreeWalkerFrom interface {
		TreeWalker_() TreeWalker
	}
)

func (t TreeWalker) TreeWalker_() TreeWalker {
	return t
}

func GetInterface() js.Value {
	singleton.Do(func() {
		if treeWalkerInterface = js.Global().Get("TreeWalker"); treeWalkerInterface.Error() != nil {
			treeWalkerInterface = js.Undefined()
			return
		}

		js.Register(treeWalkerInterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return treeWalkerInterface
}

func New() (TreeWalker, error) {
	var t TreeWalker
	var err error
	if twi := GetInterface(); !twi.IsUndefined() {
		if obj := twi.New(); obj.Error() == nil {
			t.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return t, err
}

func NewFromJSObject(obj js.Value) (TreeWalker, error) {
	var t TreeWalker
	var err error
	if eti := GetInterface(); !eti.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(eti) {
				t.SetObjectValue(obj)
			} else {
				err = ErrNotImplemented
			}
		}
	}

	return t, err
}

func (t TreeWalker) SetCurrentNode(node js.Value) error {
	t.SetValueByKey("currentNode", node)
	return nil
}

func (t TreeWalker) GetCurrentNode() (*node.Node, error) {
	v := t.GetValueByKey("currentNode")
	if v.Error() != nil {
		return nil, v.Error()
	}

	return node.NewFromJSObject(v)
}

func (t TreeWalker) NextNode() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("nextNode"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}

	return n, err
}

func (t TreeWalker) FirstChild() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("firstChild"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}

	return n, err
}

func (t TreeWalker) LastChild() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("lastChild"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}

	return n, err
}

func (t TreeWalker) NextSibling() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("nextSibling"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}

	return n, err
}

func (t TreeWalker) ParentNode() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("parentNode"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}
	return n, err
}

func (t TreeWalker) PreviousSibling() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("previousSibling"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}
	return n, err
}

func (t TreeWalker) PreviousNode() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node

	if obj = t.Call("previousNode"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}
	return n, err
}

func (t TreeWalker) Root() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node
	if obj = t.GetValueByKey("root"); obj.Error() == nil {
		n, err = node.NewFromJSObject(obj)
	}
	return n, err
}
