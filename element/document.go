package element

import (
	"errors"
	"sync"

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
	RegisterInterface(DocumentInterface)
}

var docinterface js.Value

type Document struct {
	Node
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

		Register(docinterface, func(v js.Value) (interface{}, error) {
			return ToDocument(v)
		})
		NodeInterface()
		documentfragment.GetInterface()
		dragevent.GetInterface()
	})

	return docinterface
}

func New() (Document, error) {
	var d Document
	var err error
	if di := DocumentInterface(); !di.IsUndefined() {
		if dobj := js.Global().Get("document"); !dobj.IsUndefined() {
			d.SetObject(dobj)
		}

	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func ToDocument(obj js.Value) (Document, error) {
	var d Document
	var err error
	if dci := DocumentInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {
				d.SetObject(obj)

			} else {
				err = ErrNotADocument
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d Document) getAttributeElement(attribute string) (Element, error) {
	var elem Element
	var elemObject js.Value
	var err error

	if elemObject = d.Get(attribute); !elemObject.IsUndefined() {
		if !elemObject.IsNull() {
			elem, err = element.NewFromJSObject(elemObject)
		}

	}

	return elem, err
}

func (d Document) getAttributeHTMLCollection(attribute string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = d.Get(attribute); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) ActiveElement() (Element, error) {

	return d.getAttributeElement("activeElement")

}

func (d Document) Body() (htmlelement.HtmlElement, error) {
	var body htmlelement.HtmlElement
	var bodyObject js.Value
	var err error

	if bodyObject, err = d.Get("body"); err == nil {

		body, err = htmlelement.NewFromJSObject(bodyObject)

	}

	return body, err
}

func (d Document) CharacterSet() (string, error) {
	return d.GetAttributeString("characterSet")
}

func (d Document) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (d Document) Children() (htmlcollection.HtmlCollection, error) {
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

func (d Document) DocumentElement() (Element, error) {
	return d.getAttributeElement("documentElement")
}

func (d *Document) DocumentURI() (string, error) {
	return d.GetAttributeString("documentURI")
}

func (d Document) Embeds() (htmlcollection.HtmlCollection, error) {

	return d.getAttributeHTMLCollection("embeds")
}

func (d Document) FirstElementChild() (Element, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d Document) Fonts() {
	//TO IMPLEMENT
}

func (d Document) Forms() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("forms")
}

func (d Document) FullscreenElement() (Element, error) {
	return d.getAttributeElement("fullscreenElement")
}

func (d Document) Head() (Element, error) {
	return d.getAttributeElement("head")
}

func (d Document) Hidden() (bool, error) {

	return d.GetAttributeBool("hidden")
}

func (d Document) Images() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("images")
}
func (d Document) Implementation() {
	//TO IMPLEMENT
}

func (d Document) LastElementChild() (Element, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d Document) Links() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("links")
}

func (d Document) PictureInPictureElement() (Element, error) {
	return d.getAttributeElement("pictureInPictureElement")
}

func (d Document) PictureInPictureEnabled() (bool, error) {
	return d.GetAttributeBool("pictureInPictureEnabled")
}

func (d Document) Plugins() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("plugins")
}

func (d Document) PointerLockElement() (Element, error) {
	return d.getAttributeElement("pointerLockElement")
}

func (d Document) Scripts() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("scripts")
}

func (d Document) ScrollingElement() (Element, error) {
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

func (d Document) Referrer() (string, error) {
	return d.GetAttributeString("referrer")
}

func (d Document) Title() (string, error) {
	return d.GetAttributeString("title")
}

func (d Document) SetTitle(title string) {
	d.SetAttributeString("title", title)
}

func (d Document) URL() (string, error) {
	return d.GetAttributeString("URL")
}

func (d Document) Cookie() (string, error) {
	return d.GetAttributeString("cookie")
}

func (d Document) SetCookie(cookie string) {
	d.SetAttributeString("cookie", cookie)
}
