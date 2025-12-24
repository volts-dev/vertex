package nodelist

import "github.com/volts-dev/vertex/html/node"

func (n NodeList) Item_(index int) *node.Node {
	node, _ := n.Item(index)
	return node
}
