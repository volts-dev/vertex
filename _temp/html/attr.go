package html

// Code generated DO NOT EDIT
// attr.go

import (
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/js"
)

type AttrIFace interface {
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetChildNodes() NodeList
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DispatchEvent(args ...interface{}) bool
	GetFirstChild() Node
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetName() string
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() IHTMLElement
	GetOwnerElement() IHTMLElement
	GetParentElement() IHTMLElement
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSpecified() bool
	GetTextContent() string
	SetTextContent(string)
	GetValue() string
	SetValue(string)
}
type Attr struct {
	Node
}

func ToAttr(val js.Value) Attr {
	n, _ := ToNode(val)
	return Attr{Node: n}
}

func NewAttr(args ...interface{}) Attr {
	v := js.Global().Get("Attr").New(args...)
	n, _ := ToNode(v)
	return Attr{Node: n}
}
func (a Attr) AddEventListener(args ...interface{}) {
	a.Call("addEventListener", args...)
}
func (a Attr) AppendChild(args ...interface{}) Node {
	val, err := ToNode(a.Call("appendChild", args...))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetBaseURI() (string, error) {
	val := a.Get("baseURI")
	return val.String()
}

func (a Attr) GetChildNodes() (NodeList, error) {
	return ToNodeList(a.Get("childNodes"))
}

func (a Attr) CloneNode(args ...interface{}) Node {

	val, err := ToNode(a.Get("cloneNode"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) CompareDocumentPosition(args ...interface{}) (int, error) {
	val := a.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (a Attr) Contains(args ...interface{}) (bool, error) {
	val := a.Call("contains", args...)
	return val.Bool()
}
func (a Attr) DispatchEvent(args ...interface{}) (bool, error) {
	val := a.Call("dispatchEvent", args...)
	return val.Bool()
}
func (a Attr) GetFirstChild() Node {

	val, err := ToNode(a.Get("firstChild"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetRootNode(args ...interface{}) Node {

	val, err := ToNode(a.Get("getRootNode"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) HasChildNodes(args ...interface{}) (bool, error) {
	val := a.Call("hasChildNodes", args...)
	return val.Bool()
}
func (a Attr) InsertBefore(args ...interface{}) Node {

	val, err := ToNode(a.Get("insertBefore"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetIsConnected() (bool, error) {
	val := a.Get("isConnected")
	return val.Bool()
}
func (a Attr) IsDefaultNamespace(args ...interface{}) (bool, error) {
	val := a.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (a Attr) IsEqualNode(args ...interface{}) (bool, error) {
	val := a.Call("isEqualNode", args...)
	return val.Bool()
}
func (a Attr) IsSameNode(args ...interface{}) (bool, error) {
	val := a.Call("isSameNode", args...)
	return val.Bool()
}
func (a Attr) GetLastChild() Node {

	val, err := ToNode(a.Get("lastChild"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetLocalName() (string, error) {
	val := a.Get("localName")
	return val.String()
}
func (a Attr) LookupNamespaceURI(args ...interface{}) (string, error) {
	val := a.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (a Attr) LookupPrefix(args ...interface{}) (string, error) {
	val := a.Call("lookupPrefix", args...)
	return val.String()
}
func (a Attr) GetName() (string, error) {
	val := a.Get("name")
	return val.String()
}
func (a Attr) GetNamespaceURI() (string, error) {
	val := a.Get("namespaceURI")
	return val.String()
}
func (a Attr) GetNextSibling() Node {

	val, err := ToNode(a.Get("nextSibling"))
	if err != nil {
		console.Error(err)
	}
	return val

}
func (a Attr) GetNodeName() (string, error) {
	val := a.Get("nodeName")
	return val.String()
}
func (a Attr) GetNodeType() (int, error) {
	val := a.Get("nodeType")
	return val.Int()
}
func (a Attr) GetNodeValue() (string, error) {
	val := a.Get("nodeValue")
	return val.String()
}
func (a Attr) SetNodeValue(val string) {
	a.Set("nodeValue", val)
}
func (a Attr) Normalize(args ...interface{}) {
	a.Call("normalize", args...)
}
func (a Attr) GetOwnerDocument() IHTMLElement {

	val, err := ToElement(a.Get("ownerDocument"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetOwnerElement() IHTMLElement {
	val, err := ToElement(a.Get("ownerElement"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetParentElement() IHTMLElement {
	val, err := ToElement(a.Get("parentElement"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetParentNode() Node {
	val, err := ToNode(a.Get("parentNode"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetPrefix() (string, error) {
	val := a.Get("prefix")
	return val.String()
}
func (a Attr) GetPreviousSibling() Node {
	val, err := ToNode(a.Get("previousSibling"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) RemoveChild(args ...interface{}) Node {
	val, err := ToNode(a.Get("removeChild"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) RemoveEventListener(args ...interface{}) {
	a.Call("removeEventListener", args...)
}
func (a Attr) ReplaceChild(args ...interface{}) Node {
	val, err := ToNode(a.Get("replaceChild"))
	if err != nil {
		console.Error(err)
	}
	return val
}
func (a Attr) GetSpecified() (bool, error) {
	val := a.Get("specified")
	return val.Bool()
}
func (a Attr) GetTextContent() (string, error) {
	val := a.Get("textContent")
	return val.String()
}
func (a Attr) SetTextContent(val string) {
	a.Set("textContent", val)
}
func (a Attr) GetValue() (string, error) {
	val := a.Get("value")
	return val.String()
}
func (a Attr) SetValue(val string) {
	a.Set("value", val)
}
