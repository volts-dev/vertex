package app

import (
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetLocationInterface)
}

var locationinterface js.Value

// GetInterface get the JS interface of formdata
func GetLocationInterface() js.Value {

	singleton.Do(func() {
		if locationinterface = js.Global().Get("Location"); locationinterface.Error() != nil {
			locationinterface = js.Undefined()
		}
		js.Register(locationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return locationinterface
}

type Location struct {
	js.Object
}

type LocationFrom interface {
	Location_() Location
}

func (l Location) Location_() Location {
	return l
}

func ToLocation(obj js.Value) (Location, error) {
	var l Location

	if li := GetLocationInterface(); !li.IsUndefined() {
		if obj.InstanceOf(li) {
			//l.BaseObject = l.SetObject(obj)
			l.SetValue(obj)
			return l, nil

		}
	}

	return l, js.ErrNotImplemented
}

func (l Location) Hash() (string, error) {

	return l.GetAttributeString("hash")
}

func (l Location) Host() (string, error) {

	return l.GetAttributeString("host")
}

func (l Location) Hostname() (string, error) {

	return l.GetAttributeString("hostname")
}

func (l Location) Href() (string, error) {

	return l.GetAttributeString("href")
}

func (l Location) Origin() (string, error) {

	return l.GetAttributeString("origin")
}

func (l Location) Pathname() (string, error) {

	return l.GetAttributeString("pathname")
}

func (l Location) Port() (string, error) {

	return l.GetAttributeString("port")
}

func (l Location) Protocol() (string, error) {

	return l.GetAttributeString("protocol")
}

func (l Location) Search() (string, error) {

	return l.GetAttributeString("search")
}

func (l Location) Username() (string, error) {

	return l.GetAttributeString("username")
}

func (l Location) Password() (string, error) {

	return l.GetAttributeString("password")
}

func (l Location) Assign(url string) error {

	return l.Call("assign", js.ValueOf(url)).Error()
}

func (l Location) Reload(value bool) error {

	return l.Call("reload", js.ValueOf(value)).Error()
}

func (l Location) Replace(url string) error {

	return l.Call("replace", js.ValueOf(url)).Error()
}
