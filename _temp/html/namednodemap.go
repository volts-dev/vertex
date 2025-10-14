package html

// Code generated DO NOT EDIT
// namednodemap.go

import (
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/js"
)

type NamedNodeMapIFace interface {
	GetNamedItem(args ...interface{}) Attr
	GetNamedItemNS(args ...interface{}) Attr
	Item(args ...interface{}) Attr
	GetLength() int
	RemoveNamedItem(args ...interface{}) Attr
	RemoveNamedItemNS(args ...interface{}) Attr
	SetNamedItem(args ...interface{}) Attr
	SetNamedItemNS(args ...interface{}) Attr
}
type NamedNodeMap struct {
	js.Value
}

func ToNamedNodeMap(val js.Value) NamedNodeMap {
	return NamedNodeMap{Value: val}
}
func NewNamedNodeMap(args ...interface{}) NamedNodeMap {
	return NamedNodeMap{Value: js.Global().Get("NamedNodeMap").New(args...)}
}
func (n NamedNodeMap) GetNamedItem(args ...interface{}) Attr {
	val := n.Call("getNamedItem", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) GetNamedItemNS(args ...interface{}) Attr {
	val := n.Call("getNamedItemNS", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) Item(args ...interface{}) Attr {
	val := n.Call("item", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) GetLength() int {
	val, err := n.Get("length").Int()
	if err != nil {
		console.Error(err)
		return 0
	}

	return val
}
func (n NamedNodeMap) RemoveNamedItem(args ...interface{}) Attr {
	val := n.Call("removeNamedItem", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) RemoveNamedItemNS(args ...interface{}) Attr {
	val := n.Call("removeNamedItemNS", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) SetNamedItem(args ...interface{}) Attr {
	val := n.Call("setNamedItem", args...)
	return ToAttr(val)
}
func (n NamedNodeMap) SetNamedItemNS(args ...interface{}) Attr {
	val := n.Call("setNamedItemNS", args...)
	return ToAttr(val)
}
