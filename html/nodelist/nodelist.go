package nodelist

// https://developer.mozilla.org/fr/docs/Web/API/NodeList

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/node"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var nodelistinterface js.Value

// NodeList struct
type NodeList struct {
	js.Object
}

type NodeListFrom interface {
	NodeList_() NodeList
}

func (n NodeList) NodeList_() NodeList {
	return n
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		if nodelistinterface = js.Global().Get("NodeList"); nodelistinterface.Error() != nil {
			nodelistinterface = js.Undefined()
		}
		js.Register(nodelistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return nodelistinterface
}

func NewFromJSObject(obj js.Value) (NodeList, error) {
	var n NodeList
	var err error
	if nli := GetInterface(); !nli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(nli) {
				n.SetObjectValue(obj)

			} else {
				err = ErrNotAnNodeList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return n, err
}

func (n NodeList) Item(index int) (node.Node, error) {

	var err error
	var nd node.Node

	obj := n.GetObjectValue().Index(index)

	if !obj.IsUndefined() {
		nd, err = node.NewFromJSObject(obj)
	}

	return nd, err
}
