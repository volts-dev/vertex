package html

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

var (
	ErrNotAnElement     = errors.New("Object is not an Element")
	ErrElementNoChilds  = errors.New("Element has no childs")
	ErrAttributeEmpty   = errors.New("Attribute is empty")
	ErrInsertAdjacent   = errors.New("Insert Adjacent failed")
	ErrElementNotFound  = errors.New("Element not Found")
	ErrElementsNotFound = errors.New("Elements not Found")
	ErrSendUnknownType  = errors.New("Unknown type send data provide to send method")
)

type (

	// Element stucture for basic strucutre nodes (html elements, etc)
	IHTMLElement interface {
		Node

		//GetAssignedSlot() HTMLSlotElement
		Attributes() NamedNodeMap
		//GetBaseURI() string
		ChildElementCount() (int, error)
		//ChildNodes() NodeList
		Children() (HTMLCollection, error)
		ClassList() DOMTokenList
		ClassName() (string, error)
		SetClassName(val string)
		//CloneNode(args ...interface{}) Node
		//CompareDocumentPosition(args ...interface{}) int
		//Contains(args ...interface{}) bool
		//GetFirstChild() Node
		FirstElementChild() (IHTMLElement, error)
		//	GetRootNode(args ...interface{}) Node
		//	HasChildNodes(args ...interface{}) bool
		//Id(...string) string
		InnerHTML() (string, error)
		//InsertBefore(args ...interface{}) Node
		//GetIsConnected() bool
		//IsDefaultNamespace(args ...interface{}) bool
		//IsEqualNode(args ...interface{}) bool
		//IsSameNode(args ...interface{}) bool
		//LastChild() Node
		//LastElementChild() IHTMLElement
		LocalName() (string, error)
		//LookupNamespaceURI(args ...interface{}) string
		//LookupPrefix(args ...interface{}) string
		NamespaceURI() (string, error)
		NextElementSibling() (IHTMLElement, error)
		OuterHTML() (string, error)
		//Element.part
		Prefix() (string, error)
		///GetPreviousElementSibling() IHTMLElement
		//Element.scrollHeight
		//Element.scrollLeft
		//Element.scrollTop
		//Element.scrollWidth
		ShadowRoot() ShadowRoot
		Slot() (string, error)
		TagName() (string, error)
		//Instance methods
		//After(args ...interface{})
		//Element.animate()
		//Append(args ...interface{})
		//AppendChild(args ...interface{}) Node
		AttachShadow(args ...interface{}) ShadowRoot
		//Before(args ...interface{})
		//Element.checkVisibility()
		Closest(args ...interface{}) (IHTMLElement, error)
		//Element.computedStyleMap()
		//Element.getAnimations()
		GetAttribute(args ...interface{}) (string, error)
		GetAttributeNS(args ...interface{}) (string, error)
		GetAttributeNames(args ...interface{})
		GetAttributeNode(args ...interface{}) Attr
		GetAttributeNodeNS(args ...interface{}) Attr
		//Element.getBoundingClientRect()
		//Element.getClientRects()
		GetElementsByClassName(args ...interface{}) (HTMLCollection, error)
		GetElementsByTagName(args ...interface{}) (HTMLCollection, error)
		GetElementsByTagNameNS(args ...interface{}) (HTMLCollection, error)
		//Element.getHTML()
		HasAttribute(args ...interface{}) (bool, error)
		HasAttributeNS(args ...interface{}) (bool, error)
		HasAttributes(args ...interface{}) (bool, error)
		//Element.hasPointerCapture()
		InsertAdjacentElement(args ...interface{}) (IHTMLElement, error)
		//InsertAdjacentHTML(args ...interface{})
		InsertAdjacentText(args ...interface{})
		Matches(args ...interface{}) (bool, error)
		//Element.prepend()
		QuerySelector(args string) (IHTMLElement, error)
		QuerySelectorAll(args string) (NodeList, error)
		//Element.releasePointerCapture()
		Remove()
		RemoveAttribute(args ...interface{})
		RemoveAttributeNS(args ...interface{})
		RemoveAttributeNode(args ...interface{}) Attr
		ReplaceChildren(args ...interface{}) error
		//ReplaceWith(args ...interface{})
		//Element.requestFullscreen()
		//Element.requestPointerLock()
		//Element.scroll()
		//Element.scrollBy()
		//Element.scrollIntoView()
		//Element.scrollTo()
		//SetAttribute(args ...interface{})
		//SetAttributeNS(args ...interface{})
		SetAttributeNode(args ...interface{}) Attr
		SetAttributeNodeNS(args ...interface{}) Attr
		//Element.setHTMLUnsafe()
		//Element.setPointerCapture()
		//Element.toggleAttribute()
		GetTextContent() (string, error)
		SetTextContent(string)
		ToggleAttribute(args ...interface{}) (bool, error)
		WebkitMatchesSelector(args ...interface{}) (bool, error)
	}

	HTMLElement struct {
		Node
		parent   IHTMLElement // if element is root, parent is nil
		tag      string
		xmlns    string
		children []IHTMLElement

		//Attrs    func() Map
		//RAttrs   Map                // last rendered Attrs
		//Handlers map[string]Handler // events handlers: onClick, onHover

		HTML  func() string
		RHTML string

		Childes, OldChildes []interface{}

		RefName string

		//Component *Component // can be nil
	}
)

var (
	ErrElementNotImplemented = errors.New("Browser not implemented Element")
)

var elementinterface js.Value
var customElements = make(map[string]IHTMLElement)

func init() {
	js.RegisterInterface(ElementInterface)
	DefineElement("html", HtmlElementConstructor())
}

func DefineElement(tag string, element IHTMLElement) error {
	var err *errors.Error
	if _, has := customElements[tag]; has {
		err = errors.New("Element already defined")
	}

	customElements[tag] = element
	return err
}

func CreateElement(tag string) (IHTMLElement, error) {
	if element, ok := customElements[tag]; ok {
		return element, nil
	}

	return nil, ErrElementNotImplemented
}

// GetJSInterface get the JS interface
func ElementInterface() js.Value {
	sync.OnceFunc(func() {
		if elementinterface = js.Global().Get("Element"); elementinterface.IsUndefined() {
			elementinterface = js.Undefined()
		}

		js.Register(elementinterface, func(v js.Value) (interface{}, error) {
			return ToElement(v)
		})

	})

	return elementinterface
}
func ToElement(obj js.Value) (IHTMLElement, error) {
	var e HTMLElement
	var err error
	if ei := ElementInterface(); !ei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ei) {
				e.SetValue(obj)

			} else {
				err = ErrNotAnElement
			}
		}

	} else {
		err = js.ErrNotImplemented
	}

	return e, err
}

func (h HTMLElement) GetAccessKey() (string, error) {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLElement) GetAccessKeyLabel() (string, error) {
	val := h.Get("accessKeyLabel")
	return val.String()
}

func (h HTMLElement) __AppendChild(args ...interface{}) (Node, error) {
	val := h.Call("appendChild", args...)
	return ToNode(val)
}
func (h HTMLElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return ToShadowRoot(val)
}
func (h HTMLElement) Attributes() NamedNodeMap {
	val := h.Get("attributes")
	return ToNamedNodeMap(val)
}

func (h HTMLElement) GetAutocapitalize() (string, error) {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}

func (h HTMLElement) Blur(args ...interface{}) {
	h.Call("blur", args...)
}

func (h HTMLElement) ClassList() DOMTokenList {
	val := h.Get("classList")
	return ToDOMTokenList(val)
}
func (h HTMLElement) ClassName() (string, error) {
	val := h.Get("className")
	return val.String()
}
func (h HTMLElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLElement) Click(args ...interface{}) {
	h.Call("click", args...)
}

func (h HTMLElement) Closest(args ...interface{}) (IHTMLElement, error) {
	val := h.Call("closest", args...)
	return ToElement(val)
}

func (d HTMLElement) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (h HTMLElement) Children() (HTMLCollection, error) {
	val := h.Get("children")
	return ToHTMLCollection(val)
}
func (e HTMLElement) InnerHTML() (string, error) {

	return e.GetAttributeString("innerHTML")
}

func (e HTMLElement) SetInnerHTML(value string) error {
	e.SetAttributeString("innerHTML", value)
	return nil
}

func (d HTMLElement) FirstElementChild() (IHTMLElement, error) {
	v := d.Get("firstElementChild")
	return ToElement(v)
}

func (h HTMLElement) ___CompareDocumentPosition(args ...interface{}) (int, error) {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}

func (e HTMLElement) QuerySelector(selector string) (IHTMLElement, error) {

	var err error
	var obj js.Value
	var nod IHTMLElement

	if obj = e.Call("querySelector", js.ValueOf(selector)); obj.IsUndefined() {
		if !obj.IsNull() {
			nod, err = ToElement(obj)
		} else {
			err = errors.New(ErrElementNotFound.Error() + " " + selector)
		}
	}
	return nod, err
}

func (e HTMLElement) QuerySelectorAll(selector string) (NodeList, error) {

	var err error
	var obj js.Value
	var nlist NodeList

	if obj = e.Call("querySelectorAll", js.ValueOf(selector)); obj.IsUndefined() {
		if !obj.IsNull() {
			return ToNodeList(obj)
		} else {
			err = errors.New(ErrElementsNotFound.Error() + " " + selector)
		}
	}
	return nlist, err
}

func (h HTMLElement) GetContentEditable() (string, error) {
	val := h.Get("contentEditable")
	return val.String()
}
func (h HTMLElement) SetContentEditable(val string) {
	h.Set("contentEditable", val)
}

/*
func (h HTMLElement) GetDataset() DOMStringMap {
	val := h.Get("dataset")
	return ToDOMStringMap(val)
}
*/

func (h HTMLElement) GetDir() (string, error) {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLElement) SetDir(val string) {
	h.Set("dir", val)
}

func (h HTMLElement) GetDraggable() (bool, error) {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLElement) GetEnterKeyHint() (string, error) {
	val := h.Get("enterKeyHint")
	return val.String()
}
func (h HTMLElement) SetEnterKeyHint(val string) {
	h.Set("enterKeyHint", val)
}
func (h HTMLElement) GetFirstChild() (Node, error) {
	val := h.Get("firstChild")
	return ToNode(val)
}
func (h HTMLElement) Focus(args ...interface{}) {
	h.Call("focus", args...)
}
func (h HTMLElement) GetAttribute(args ...interface{}) (string, error) {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLElement) GetAttributeNS(args ...interface{}) (string, error) {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return ToAttr(val)
}
func (h HTMLElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return ToAttr(val)
}
func (h HTMLElement) GetElementsByClassName(args ...interface{}) (HTMLCollection, error) {
	val := h.Call("getElementsByClassName", args...)
	return ToHTMLCollection(val)
}
func (h HTMLElement) GetElementsByTagName(args ...interface{}) (HTMLCollection, error) {
	val := h.Call("getElementsByTagName", args...)
	return ToHTMLCollection(val)
}
func (h HTMLElement) GetElementsByTagNameNS(args ...interface{}) (HTMLCollection, error) {
	val := h.Call("getElementsByTagNameNS", args...)
	return ToHTMLCollection(val)
}
func (h HTMLElement) ___GetRootNode(args ...interface{}) (Node, error) {
	val := h.Call("getRootNode", args...)
	return ToNode(val)
}
func (h HTMLElement) HasAttribute(args ...interface{}) (bool, error) {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLElement) HasAttributeNS(args ...interface{}) (bool, error) {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLElement) HasAttributes(args ...interface{}) (bool, error) {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLElement) ___HasChildNodes(args ...interface{}) (bool, error) {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLElement) GetHidden() (bool, error) {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLElement) GetId() (string, error) {
	val := h.Get("id")
	return val.String()
}
func (h HTMLElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLElement) GetInnerText() (string, error) {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLElement) GetInputMode() (string, error) {
	val := h.Get("inputMode")
	return val.String()
}
func (h HTMLElement) SetInputMode(val string) {
	h.Set("inputMode", val)
}
func (h HTMLElement) InsertAdjacentElement(args ...interface{}) (IHTMLElement, error) {
	val := h.Call("insertAdjacentElement", args...)
	return ToElement(val)
}
func (h HTMLElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLElement) ___InsertBefore(args ...interface{}) (Node, error) {
	val := h.Call("insertBefore", args...)
	return ToNode(val)
}
func (h HTMLElement) GetIsConnected() (bool, error) {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLElement) GetIsContentEditable() (bool, error) {
	val := h.Get("isContentEditable")
	return val.Bool()
}
func (h HTMLElement) ___IsDefaultNamespace(args ...interface{}) (bool, error) {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLElement) ___IsEqualNode(args ...interface{}) (bool, error) {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLElement) ___IsSameNode(args ...interface{}) (bool, error) {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLElement) GetLang() (string, error) {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLElement) ___LastChild() (Node, error) {
	val := h.Get("lastChild")
	return ToNode(val)
}
func (h HTMLElement) GetLocalName() (string, error) {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLElement) ___LookupNamespaceURI(args ...interface{}) (string, error) {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}

func (e HTMLElement) LocalName() (string, error) {

	return e.GetAttributeString("localname")
}

func (h HTMLElement) ___LookupPrefix(args ...interface{}) (string, error) {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLElement) Matches(args ...interface{}) (bool, error) {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLElement) NamespaceURI() (string, error) {
	val := h.Get("namespaceURI")
	return val.String()
}
func (e HTMLElement) NextElementSibling() (IHTMLElement, error) {
	val := e.Get("nextElementSibling")
	return ToElement(val)
}

func (e HTMLElement) OuterHTML() (string, error) {
	val := e.Get("outerHTML")
	return val.String()
}

func (h HTMLElement) GetNextSibling() (Node, error) {
	val := h.Get("nextSibling")
	return ToNode(val)
}
func (h HTMLElement) GetNodeName() (string, error) {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLElement) GetNodeType() (int, error) {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLElement) GetNodeValue() (string, error) {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLElement) GetNonce() (string, error) {
	val := h.Get("nonce")
	return val.String()
}
func (h HTMLElement) SetNonce(val string) {
	h.Set("nonce", val)
}
func (h HTMLElement) ___Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}

/*
func (h HTMLElement) GetOnabort() EventHandler {
	val := h.Get("onabort")
	return ToEventHandler(val)
}
func (h HTMLElement) SetOnabort(val EventHandler) {
	h.Set("onabort", val)
}
func (h HTMLElement) GetOnauxclick() EventHandler {
	val := h.Get("onauxclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnauxclick(val EventHandler) {
	h.Set("onauxclick", val)
}
func (h HTMLElement) GetOnblur() EventHandler {
	val := h.Get("onblur")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnblur(val EventHandler) {
	h.Set("onblur", val)
}
func (h HTMLElement) GetOncancel() EventHandler {
	val := h.Get("oncancel")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncancel(val EventHandler) {
	h.Set("oncancel", val)
}
func (h HTMLElement) GetOncanplay() EventHandler {
	val := h.Get("oncanplay")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncanplay(val EventHandler) {
	h.Set("oncanplay", val)
}
func (h HTMLElement) GetOncanplaythrough() EventHandler {
	val := h.Get("oncanplaythrough")
	return JSValueToEventHandler(val. JSValue())
}
func (h HTMLElement) SetOncanplaythrough(val EventHandler) {
	h.Set("oncanplaythrough", val)
}
func (h HTMLElement) GetOnchange() EventHandler {
	val := h.Get("onchange")
	return ToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnchange(val EventHandler) {
	h.Set("onchange", val)
}
func (h HTMLElement) GetOnclick() EventHandler {
	val := h.Get("onclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnclick(val EventHandler) {
	h.Set("onclick", val)
}
func (h HTMLElement) GetOnclose() EventHandler {
	val := h.Get("onclose")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnclose(val EventHandler) {
	h.Set("onclose", val)
}
func (h HTMLElement) GetOncontextmenu() EventHandler {
	val := h.Get("oncontextmenu")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncontextmenu(val EventHandler) {
	h.Set("oncontextmenu", val)
}
func (h HTMLElement) GetOncopy() EventHandler {
	val := h.Get("oncopy")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncopy(val EventHandler) {
	h.Set("oncopy", val)
}
func (h HTMLElement) GetOncuechange() EventHandler {
	val := h.Get("oncuechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncuechange(val EventHandler) {
	h.Set("oncuechange", val)
}
func (h HTMLElement) GetOncut() EventHandler {
	val := h.GetOnblur()Get("oncut")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncut(val EventHandler) {
	h.Set("oncut", val)
}
func (h HTMLElement) GetOndblclick() EventHandler {
	val := h.Get("ondblclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndblclick(val EventHandler) {
	h.Set("ondblclick", val)
}
func (h HTMLElement) GetOndrag() EventHandler {
	val := h.Get("ondrag")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndrag(val EventHandler) {
	h.Set("ondrag", val)
}
func (h HTMLElement) GetOndragend() EventHandler {
	val := h.Get("ondragend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragend(val EventHandler) {
	h.Set("ondragend", val)
}
func (h HTMLElement) GetOndragenter() EventHandler {
	val := h.Get("ondragenter")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragenter(val EventHandler) {
	h.Set("ondragenter", val)
}
func (h HTMLElement) GetOndragexit() EventHandler {
	val := h.Get("ondragexit")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragexit(val EventHandler) {
	h.Set("ondragexit", val)
}
func (h HTMLElement) GetOndragleave() EventHandler {
	val := h.Get("ondragleave")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragleave(val EventHandler) {
	h.Set("ondragleave", val)
}
func (h HTMLElement) GetOndragover() EventHandler {
	val := h.Get("ondragover")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragover(val EventHandler) {
	h.Set("ondragover", val)
}
func (h HTMLElement) GetOndragstart() EventHandler {
	val := h.Get("ondragstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragstart(val EventHandler) {
	h.Set("ondragstart", val)
}
func (h HTMLElement) GetOndrop() EventHandler {
	val := h.Get("ondrop")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndrop(val EventHandler) {
	h.Set("ondrop", val)
}
func (h HTMLElement) GetOndurationchange() EventHandler {
	val := h.Get("ondurationchange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndurationchange(val EventHandler) {
	h.Set("ondurationchange", val)
}
func (h HTMLElement) GetOnemptied() EventHandler {
	val := h.Get("onemptied")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnemptied(val EventHandler) {
	h.Set("onemptied", val)
}
func (h HTMLElement) GetOnended() EventHandler {
	val := h.Get("onended")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnended(val EventHandler) {
	h.Set("onended", val)
}
func (h HTMLElement) GetOnerror() OnErrorEventHandler {
	val := h.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnerror(val OnErrorEventHandler) {
	h.Set("onerror", val)
}
func (h HTMLElement) GetOnfocus() EventHandler {
	val := h.Get("onfocus")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnfocus(val EventHandler) {
	h.Set("onfocus", val)
}
func (h HTMLElement) GetOninput() EventHandler {
	val := h.Get("oninput")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOninput(val EventHandler) {
	h.Set("oninput", val)
}
func (h HTMLElement) GetOninvalid() EventHandler {
	val := h.Get("oninvalid")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOninvalid(val EventHandler) {
	h.Set("oninvalid", val)
}
func (h HTMLElement) GetOnkeydown() EventHandler {
	val := h.Get("onkeydown")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeydown(val EventHandler) {
	h.Set("onkeydown", val)
}
func (h HTMLElement) GetOnkeypress() EventHandler {
	val := h.Get("onkeypress")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeypress(val EventHandler) {
	h.Set("onkeypress", val)
}
func (h HTMLElement) GetOnkeyup() EventHandler {
	val := h.Get("onkeyup")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeyup(val EventHandler) {
	h.Set("onkeyup", val)
}
func (h HTMLElement) GetOnload() EventHandler {
	val := h.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnload(val EventHandler) {
	h.Set("onload", val)
}
func (h HTMLElement) GetOnloadeddata() EventHandler {
	val := h.Get("onloadeddata")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadeddata(val EventHandler) {
	h.Set("onloadeddata", val)
}
func (h HTMLElement) GetOnloadedmetadata() EventHandler {
	val := h.Get("onloadedmetadata")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadedmetadata(val EventHandler) {
	h.Set("onloadedmetadata", val)
}
func (h HTMLElement) GetOnloadend() EventHandler {
	val := h.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadend(val EventHandler) {
	h.Set("onloadend", val)
}
func (h HTMLElement) GetOnloadstart() EventHandler {
	val := h.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadstart(val EventHandler) {
	h.Set("onloadstart", val)
}
func (h HTMLElement) GetOnmousedown() EventHandler {
	val := h.Get("onmousedown")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmousedown(val EventHandler) {
	h.Set("onmousedown", val)
}
func (h HTMLElement) GetOnmouseenter() EventHandler {
	val := h.Get("onmouseenter")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseenter(val EventHandler) {
	h.Set("onmouseenter", val)
}
func (h HTMLElement) GetOnmouseleave() EventHandler {
	val := h.Get("onmouseleave")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseleave(val EventHandler) {
	h.Set("onmouseleave", val)
}
func (h HTMLElement) GetOnmousemove() EventHandler {
	val := h.Get("onmousemove")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmousemove(val EventHandler) {
	h.Set("onmousemove", val)
}
func (h HTMLElement) GetOnmouseout() EventHandler {
	val := h.Get("onmouseout")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseout(val EventHandler) {
	h.Set("onmouseout", val)
}
func (h HTMLElement) GetOnmouseover() EventHandler {
	val := h.Get("onmouseover")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseover(val EventHandler) {
	h.Set("onmouseover", val)
}
func (h HTMLElement) GetOnmouseup() EventHandler {
	val := h.Get("onmouseup")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseup(val EventHandler) {
	h.Set("onmouseup", val)
}
func (h HTMLElement) GetOnpaste() EventHandler {
	val := h.Get("onpaste")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnpaste(val EventHandler) {
	h.Set("onpaste", val)
}
func (h HTMLElement) GetOnpause() EventHandler {
	val := h.Get("onpause")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnpause(val EventHandler) {
	h.Set("onpause", val)
}
func (h HTMLElement) GetOnplay() EventHandler {
	val := h.Get("onplay")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnplay(val EventHandler) {
	h.Set("onplay", val)
}
func (h HTMLElement) GetOnplaying() EventHandler {
	val := h.Get("onplaying")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnplaying(val EventHandler) {
	h.Set("onplaying", val)
}
func (h HTMLElement) GetOnprogress() EventHandler {
	val := h.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnprogress(val EventHandler) {
	h.Set("onprogress", val)
}
func (h HTMLElement) GetOnratechange() EventHandler {
	val := h.Get("onratechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnratechange(val EventHandler) {
	h.Set("onratechange", val)
}
func (h HTMLElement) GetOnreset() EventHandler {
	val := h.Get("onreset")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnreset(val EventHandler) {
	h.Set("onreset", val)
}
func (h HTMLElement) GetOnresize() EventHandler {
	val := h.Get("onresize")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnresize(val EventHandler) {
	h.Set("onresize", val)
}
func (h HTMLElement) GetOnscroll() EventHandler {
	val := h.Get("onscroll")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnscroll(val EventHandler) {
	h.Set("onscroll", val)
}
func (h HTMLElement) GetOnsecuritypolicyviolation() EventHandler {
	val := h.Get("onsecuritypolicyviolation")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsecuritypolicyviolation(val EventHandler) {
	h.Set("onsecuritypolicyviolation", val)
}
func (h HTMLElement) GetOnseeked() EventHandler {
	val := h.Get("onseeked")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnseeked(val EventHandler) {
	h.Set("onseeked", val)
}
func (h HTMLElement) GetOnseeking() EventHandler {
	val := h.Get("onseeking")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnseeking(val EventHandler) {
	h.Set("onseeking", val)
}
func (h HTMLElement) GetOnselect() EventHandler {
	val := h.Get("onselect")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnselect(val EventHandler) {
	h.Set("onselect", val)
}
func (h HTMLElement) GetOnstalled() EventHandler {
	val := h.Get("onstalled")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnstalled(val EventHandler) {
	h.Set("onstalled", val)
}
func (h HTMLElement) GetOnsubmit() EventHandler {
	val := h.Get("onsubmit")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsubmit(val EventHandler) {
	h.Set("onsubmit", val)
}
func (h HTMLElement) GetOnsuspend() EventHandler {
	val := h.Get("onsuspend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsuspend(val EventHandler) {
	h.Set("onsuspend", val)
}
func (h HTMLElement) GetOntimeupdate() EventHandler {
	val := h.Get("ontimeupdate")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOntimeupdate(val EventHandler) {
	h.Set("ontimeupdate", val)
}
func (h HTMLElement) GetOntoggle() EventHandler {
	val := h.Get("ontoggle")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOntoggle(val EventHandler) {
	h.Set("ontoggle", val)
}
func (h HTMLElement) GetOnvolumechange() EventHandler {
	val := h.Get("onvolumechange")
	return ToEventHandler(val)
}
func (h HTMLElement) SetOnvolumechange(val EventHandler) {
	h.Set("onvolumechange", val)
}
func (h HTMLElement) GetOnwaiting() EventHandler {
	val := h.Get("onwaiting")
	return ToEventHandler(val)
}
func (h HTMLElement) SetOnwaiting(val EventHandler) {
	h.Set("onwaiting", val)
}
func (h HTMLElement) GetOnwheel() EventHandler {
	val := h.Get("onwheel")
	return ToEventHandler(val)
}
func (h HTMLElement) SetOnwheel(val EventHandler) {
	h.Set("onwheel", val)
}
func (h HTMLElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
*/

func (h HTMLElement) GetParentElement() (IHTMLElement, error) {
	val := h.Get("parentElement")
	return ToElement(val)
}
func (h HTMLElement) GetParentNode() (Node, error) {
	val := h.Get("parentNode")
	return ToNode(val)
}
func (h HTMLElement) Prefix() (string, error) {
	return h.Get("prefix").String()
}
func (h HTMLElement) GetPreviousSibling() (Node, error) {
	val := h.Get("previousSibling")
	return ToNode(val)
}

func (e HTMLElement) Remove() {
	e.Call("remove")
}

func (h HTMLElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return ToAttr(val)
}

func (e HTMLElement) ReplaceChildren(params ...interface{}) error {
	var err error
	var arrayJS []interface{}
	for _, param := range params {
		switch p := param.(type) {
		case Node:
			arrayJS = append(arrayJS, p.Value())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	e.Call("replaceChildren", arrayJS...)

	return err
}
func (h HTMLElement) ___RemoveChild(args ...interface{}) (Node, error) {
	val := h.Call("removeChild", args...)
	return ToNode(val)
}

func (h HTMLElement) ___ReplaceChild(args ...interface{}) (Node, error) {
	val := h.Call("replaceChild", args...)
	return ToNode(val)
}
func (h HTMLElement) _SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return ToAttr(val)
}
func (h HTMLElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return ToAttr(val)
}
func (h HTMLElement) ShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return ToShadowRoot(val)
}
func (h HTMLElement) Slot() (string, error) {
	return h.Get("slot").String()
}
func (h HTMLElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLElement) GetSpellcheck() (bool, error) {
	return h.Get("spellcheck").Bool()
}
func (h HTMLElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLElement) GetStyle() (CSSStyleDeclaration, error) {
	val := h.Get("style")
	return NewFromJSObject(val)
}
func (h HTMLElement) GetTabIndex() (int, error) {
	return h.Get("tabIndex").Int()
}
func (h HTMLElement) SetTabIndex(val int) {
	h.Set("tabIndex", val)
}
func (h HTMLElement) TagName() (string, error) {
	return h.Get("tagName").String()
}
func (h HTMLElement) GetTextContent() (string, error) {
	return h.Get("textContent").String()
}
func (h HTMLElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLElement) GetTitle() (string, error) {
	return h.Get("title").String()
}
func (h HTMLElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLElement) ToggleAttribute(args ...interface{}) (bool, error) {
	return h.Call("toggleAttribute", args...).Bool()
}
func (h HTMLElement) GetTranslate() (bool, error) {
	return h.Get("translate").Bool()
}
func (h HTMLElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLElement) WebkitMatchesSelector(args ...interface{}) (bool, error) {
	return h.Call("webkitMatchesSelector", args...).Bool()
}
