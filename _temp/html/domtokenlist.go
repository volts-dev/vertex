package html

import "github.com/volts-dev/vertex/js"

type DOMTokenListIFace interface {
	Add(args ...interface{})
	Contains(args ...interface{}) bool
	Item(args ...interface{}) string
	GetLength() int
	Remove(args ...interface{})
	Replace(args ...interface{}) bool
	Supports(args ...interface{}) bool
	Toggle(args ...interface{}) bool
	GetValue() string
	SetValue(string)
}
type DOMTokenList struct {
	js.Value
}

func ToDOMTokenList(val js.Value) DOMTokenList {
	return DOMTokenList{Value: val}
}

func NewDOMTokenList(args ...interface{}) DOMTokenList {
	return DOMTokenList{Value: js.Global().Get("DOMTokenList").New(args...)}
}
func (d DOMTokenList) Add(args ...interface{}) {
	d.Call("add", args...)
}
func (d DOMTokenList) Contains(args ...interface{}) (bool, error) {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DOMTokenList) Item(args ...interface{}) (string, error) {
	val := d.Call("item", args...)
	return val.String()
}
func (d DOMTokenList) GetLength() (int, error) {
	val := d.Get("length")
	return val.Int()
}
func (d DOMTokenList) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
func (d DOMTokenList) Replace(args ...interface{}) (bool, error) {
	val := d.Call("replace", args...)
	return val.Bool()
}
func (d DOMTokenList) Supports(args ...interface{}) (bool, error) {
	val := d.Call("supports", args...)
	return val.Bool()
}
func (d DOMTokenList) Toggle(args ...interface{}) (bool, error) {
	val := d.Call("toggle", args...)
	return val.Bool()
}
func (d DOMTokenList) GetValue() (string, error) {
	val := d.Get("value")
	return val.String()
}
func (d DOMTokenList) SetValue(val string) {
	d.Set("value", val)
}
