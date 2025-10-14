package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	//ErrNotImplemented = errors.New("Browser not implemented Document")
	//ErrNotADocument ErrNotADocument
	ErrNotADocument     = errors.New("The given value must be a document")
	ErrElementNotFound  = errors.New("Element not Found")
	ErrElementsNotFound = errors.New("Elements not Found")
)

func init() {
	js.RegisterInterface(DocumentInterface)
}

var docinterface js.Value

type Document struct {
	html.IHTMLElement
	body html.IHTMLElement
}

type DocumentFrom interface {
	Document_() Document
}

func (d Document) Document_() Document {
	return d
}

func DocumentInterface() js.Value {
	sync.OnceFunc(func() {
		if docinterface = js.Global().Get("Document"); docinterface.IsUndefined() {
			docinterface = js.Undefined()
		}

		js.Register(docinterface, func(v js.Value) (interface{}, error) {
			return ToDocument(v)
		})

		html.NodeInterface()
		GetDocumentFragmentInterface()
		GetDragEventInterface()
	})

	return docinterface
}

func NewDocument() (Document, error) {
	var d Document
	var err error
	if di := DocumentInterface(); !di.IsUndefined() {
		if dobj := js.Global().Get("document"); !dobj.IsUndefined() {
			d.SetValue(dobj)
		}

	} else {
		err = errors.ErrNotImplemented
	}

	return d, err
}

func ToDocument(obj js.Value) (Document, error) {
	var d Document
	var err error
	if dci := DocumentInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = errors.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {
				d.SetValue(obj)

			} else {
				err = ErrNotADocument
			}
		}
	} else {
		err = errors.ErrNotImplemented
	}
	return d, err
}

func (d Document) getAttributeElement(attribute string) (html.IHTMLElement, error) {
	var elem html.IHTMLElement
	var elemObject js.Value
	var err error

	if elemObject = d.Get(attribute); !elemObject.IsUndefined() {
		if !elemObject.IsNull() {
			elem, err = html.ToElement(elemObject)
		}

	}

	return elem, err
}

func (d Document) getAttributeHTMLCollection(attribute string) (html.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection html.HTMLCollection

	if obj = d.Get(attribute); !obj.IsUndefined() {
		collection, err = html.ToHTMLCollection(obj)
	}

	return collection, err
}

func (d Document) ActiveElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("activeElement")
}

func (d Document) Body() html.IHTMLElement {
	if d.body != nil {
		return d.body
	}

	var err error
	d.body, err = html.ToElement(d.Get("body"))
	if err != nil {
		console.Error(err)
	}

	return d.body
}

func (d Document) CharacterSet() (string, error) {
	return d.GetAttributeString("characterSet")
}

func (d Document) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (d Document) Children() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("children")
}

func (d Document) CompatMode() (string, error) {
	return d.GetAttributeString("compatMode")
}

func (d Document) ContentType() (string, error) {
	return d.GetAttributeString("contentType")
}

func (d *Document) Doctype() {
	//TO IMPLEMENT
}

func (d Document) DocumentElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("documentElement")
}

func (d *Document) DocumentURI() (string, error) {
	return d.GetAttributeString("documentURI")
}

func (d Document) Embeds() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("embeds")
}

func (d Document) FirstElementChild() (html.IHTMLElement, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d Document) Fonts() {
	//TO IMPLEMENT
}

func (d Document) Forms() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("forms")
}

func (d Document) FullscreenElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("fullscreenElement")
}

func (d Document) Head() (html.IHTMLElement, error) {
	return d.getAttributeElement("head")
}

func (d Document) Hidden() (bool, error) {

	return d.GetAttributeBool("hidden")
}

func (d Document) Images() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("images")
}
func (d Document) Implementation() {
	//TO IMPLEMENT
}

func (d Document) LastElementChild() (html.IHTMLElement, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d Document) Links() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("links")
}

func (d Document) PictureInPictureElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("pictureInPictureElement")
}

func (d Document) PictureInPictureEnabled() (bool, error) {
	return d.GetAttributeBool("pictureInPictureEnabled")
}

func (d Document) Plugins() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("plugins")
}

func (d Document) PointerLockElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("pointerLockElement")
}

func (d Document) Scripts() (html.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("scripts")
}

func (d Document) ScrollingElement() (html.IHTMLElement, error) {
	return d.getAttributeElement("scrollingElement")
}

func (d Document) VisibilityState() (string, error) {
	return d.GetAttributeString("visibilityState")
}

func (d Document) Domain() (string, error) {
	return d.GetAttributeString("domain")
}

func (d Document) LastModified() (string, error) {
	return d.GetAttributeString("lastModified")
}

func (d Document) SetDomain(domain string) {
	d.SetAttributeString("domain", domain)
}

func (d Document) ReadyState() (string, error) {
	return d.GetAttributeString("readyState")
}

func (d Document) Referrer() string {
	referrer, err := d.GetAttributeString("referrer")
	if err != nil {
		console.Error(err)
	}

	return referrer
}

func (d Document) Title() string {
	title, err := d.GetAttributeString("title")
	if err != nil {
		console.Error(err)
	}

	return title
}

func (d Document) SetTitle(title string) {
	d.SetAttributeString("title", title)
}

func (d Document) URL() string {
	url, err := d.GetAttributeString("URL")
	if err != nil {
		console.Error(err)
	}

	return url
}

func (d Document) Cookie() (string, error) {
	return d.GetAttributeString("cookie")
}

func (d Document) SetCookie(cookie string) {
	d.SetAttributeString("cookie", cookie)
}
