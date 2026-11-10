package htmlareaelement

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/domtokenlist"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlareaelementinterface js.Value

// HtmlAreaElement struct
type HtmlAreaElement struct {
	htmlelement.HtmlElement
}

type HtmlAreaElementFrom interface {
	HtmlAreaElement_() HtmlAreaElement
}

func (h HtmlAreaElement) HtmlAreaElement_() HtmlAreaElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if htmlareaelementinterface = js.Global().Get("HTMLAreaElement"); htmlareaelementinterface.Error() != nil {
			htmlareaelementinterface = js.Undefined()
		}
		js.Register(htmlareaelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlareaelementinterface
}

func New(d document.Document) (HtmlAreaElement, error) {
	var err error

	var h HtmlAreaElement
	var e element.Element

	if e, err = d.CreateElement("area"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlAreaElement, error) {
	var h HtmlAreaElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {

		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnHtmlAreaElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlAreaElement, error) {
	var h HtmlAreaElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHtmlAreaElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlAreaElement) Alt() string {
	return h.GetAttributeString("alt")
}

func (h HtmlAreaElement) SetAlt(value string) error {
	return h.SetAttributeString("alt", value)
}

func (h HtmlAreaElement) Coords() string {
	return h.GetAttributeString("coords")
}

func (h HtmlAreaElement) SetCoords(value string) error {
	return h.SetAttributeString("coords", value)
}

func (h HtmlAreaElement) Download() string {
	return h.GetAttributeString("download")
}

func (h HtmlAreaElement) SetDownload(value string) error {
	return h.SetAttributeString("download", value)
}

func (h HtmlAreaElement) Hash() string {
	return h.GetAttributeString("hash")
}

func (h HtmlAreaElement) SetHash(value string) error {
	return h.SetAttributeString("hash", value)
}

func (h HtmlAreaElement) Host() string {
	return h.GetAttributeString("host")
}

func (h HtmlAreaElement) SetHost(value string) error {
	return h.SetAttributeString("host", value)
}

func (h HtmlAreaElement) Hostname() string {
	return h.GetAttributeString("hostname")
}

func (h HtmlAreaElement) SetHostname(value string) error {
	return h.SetAttributeString("hostname", value)
}

func (h HtmlAreaElement) Href() string {
	return h.GetAttributeString("href")
}

func (h HtmlAreaElement) SetHref(value string) error {
	return h.SetAttributeString("href", value)
}

func (h HtmlAreaElement) Origin() string {
	return h.GetAttributeString("origin")
}

func (h HtmlAreaElement) Password() string {
	return h.GetAttributeString("password")
}

func (h HtmlAreaElement) SetPassword(value string) error {
	return h.SetAttributeString("password", value)
}

func (h HtmlAreaElement) Pathname() string {
	return h.GetAttributeString("pathname")
}

func (h HtmlAreaElement) SetPathname(value string) error {
	return h.SetAttributeString("pathname", value)
}

func (h HtmlAreaElement) Port() string {
	return h.GetAttributeString("port")
}

func (h HtmlAreaElement) SetPort(value string) error {
	return h.SetAttributeString("port", value)
}

func (h HtmlAreaElement) Protocol() string {
	return h.GetAttributeString("protocol")
}

func (h HtmlAreaElement) SetProtocol(value string) error {
	return h.SetAttributeString("protocol", value)
}

func (h HtmlAreaElement) ReferrerPolicy() string {
	return h.GetAttributeString("referrerPolicy")
}

func (h HtmlAreaElement) SetReferrerPolicy(value string) error {
	return h.SetAttributeString("referrerPolicy", value)
}

func (h HtmlAreaElement) Rel() string {
	return h.GetAttributeString("rel")
}

func (h HtmlAreaElement) SetRel(value string) error {
	return h.SetAttributeString("rel", value)
}

func (h HtmlAreaElement) RelList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj = h.GetValueByKey("relList"); obj.Error() == nil {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (h HtmlAreaElement) Search() string {
	return h.GetAttributeString("search")
}

func (h HtmlAreaElement) SetSearch(value string) error {
	return h.SetAttributeString("search", value)
}

func (h HtmlAreaElement) TabIndex() int {
	return h.GetAttributeInt("tabIndex")
}

func (h HtmlAreaElement) SetTabIndex(value int) error {
	return h.SetAttributeInt("tabIndex", value)
}

func (h HtmlAreaElement) Target() string {
	return h.GetAttributeString("target")
}

func (h HtmlAreaElement) SetTarget(value string) error {
	return h.SetAttributeString("target", value)
}

func (h HtmlAreaElement) Shape() string {
	return h.GetAttributeString("shape")
}

func (h HtmlAreaElement) SetShape(value string) error {
	return h.SetAttributeString("shape", value)
}

func (h HtmlAreaElement) Username() string {
	return h.GetAttributeString("username")
}

func (h HtmlAreaElement) SetUsername(value string) error {
	return h.SetAttributeString("username", value)
}
