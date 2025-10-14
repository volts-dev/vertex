package html

import (
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/js"
)

type NodeList struct {
	js.Value
}

func ToNodeList(val js.Value) (NodeList, error) {
	return NodeList{Value: val}, nil
}

func NewNodeList(args ...interface{}) NodeList {
	return NodeList{Value: js.Global().Get("NodeList").New(args...)}
}

func (n NodeList) Item(args ...interface{}) Node {
	node, err := ToNode(n.Call("item", args...))
	if err != nil {
		console.Error(err)
	}
	return node
}

func (n NodeList) GetLength() int {
	val, err := n.Get("length").Int()
	if err != nil {
		console.Error(err)
		return 0
	}

	return val
}
