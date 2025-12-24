package url

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/urlsearchparams"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var urlinterface js.Value

// GetInterface get the JS interface of URL
func GetInterface() js.Value {

	singleton.Do(func() {

		if urlinterface = js.Global().Get("URL"); urlinterface.Error() != nil {
			urlinterface = js.Undefined()
		}
		js.Register(urlinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return urlinterface
}

type URL struct {
	js.Object
}

type URLFrom interface {
	URL_() URL
}

func (u URL) Location_() URL {
	return u
}

func New(value string) (URL, error) {
	var u URL
	var err error
	var obj js.Value

	if ui := GetInterface(); !ui.IsUndefined() {

		if obj = ui.New(js.ValueOf(value)); obj.Error() == nil {
			u.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return u, err
}

func NewFromJSObject(obj js.Value) (URL, error) {
	var u URL

	if ui := GetInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.SetObjectValue(obj)
			return u, nil

		}
	}

	return u, ErrNotImplemented
}

func (u URL) Hash() (string, error) {

	return u.GetAttributeString("hash")
}

func (u URL) SetHash(hash string) error {

	return u.SetAttributeString("hash", hash)
}

func (u URL) Host() (string, error) {

	return u.GetAttributeString("host")
}

func (u URL) SetHost(host string) error {

	return u.SetAttributeString("host", host)
}

func (u URL) Hostname() (string, error) {

	return u.GetAttributeString("hostname")
}

func (u URL) SetHostname(hostname string) error {

	return u.SetAttributeString("hostname", hostname)
}

func (u URL) Href() (string, error) {

	return u.GetAttributeString("href")
}

func (u URL) SetHref(href string) error {

	return u.SetAttributeString("href", href)
}

func (u URL) Origin() (string, error) {

	return u.GetAttributeString("origin")
}

func (u URL) Pathname() (string, error) {

	return u.GetAttributeString("pathname")
}

func (u URL) SetPathname(pathname string) error {

	return u.SetAttributeString("pathname", pathname)
}

func (u URL) Port() (string, error) {

	return u.GetAttributeString("port")
}

func (u URL) SetPort(port string) error {

	return u.SetAttributeString("port", port)
}

func (u URL) Protocol() (string, error) {

	return u.GetAttributeString("protocol")
}

func (u URL) SetProtocol(protocol string) error {

	return u.SetAttributeString("protocol", protocol)
}

func (u URL) Username() (string, error) {

	return u.GetAttributeString("username")
}

func (u URL) SetUsername(username string) error {

	return u.SetAttributeString("username", username)
}

func (u URL) Password() (string, error) {

	return u.GetAttributeString("password")
}

func (u URL) SetPassword(password string) error {

	return u.SetAttributeString("password", password)
}

func (u URL) Search() (string, error) {

	return u.GetAttributeString("search")
}

func (u URL) SetSearch(search string) error {

	return u.SetAttributeString("search", search)
}

func (u URL) SearchParams() (urlsearchparams.URLSearchParams, error) {
	var err error
	var obj js.Value
	var params urlsearchparams.URLSearchParams

	if obj = u.GetValueByKey("searchParams"); obj.Error() == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrNotAnObject

		} else {

			params, err = urlsearchparams.NewFromJSObject(obj)
		}
	}

	return params, err
}

func CreateObjectURL(ob interface{}) (string, error) {
	var err error
	var obj js.Value
	var ret string
	if ui := GetInterface(); !ui.IsUndefined() {
		if objGo, ok := ob.(js.ObjectFrom); ok {

			if obj = ui.Call("createObjectURL", objGo.GetObjectValue()); obj.Error() == nil {
				if obj.Type() == js.TypeString {
					return obj.String()
				} else {
					err = js.ErrObjectNotString
				}

			}

		}
	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func RevokeObjectURL(objecturl string) error {
	var err error

	if ui := GetInterface(); !ui.IsUndefined() {
		err = ui.Call("revokeObjectURL", objecturl).Error()
	} else {
		err = ErrNotImplemented
	}

	return err
}

func (u URL) ToJSON() (string, error) {

	var err error
	var obj js.Value
	var ret string

	if obj = u.Call("toJSON"); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrObjectNotString
		}

	}
	return ret, err
}
