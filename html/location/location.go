package location

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var locationinterface js.Value

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

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

func NewFromJSObject(obj js.Value) (Location, error) {
	var l Location

	if li := GetInterface(); !li.IsUndefined() {
		if obj.InstanceOf(li) {
			l.SetObjectValue(obj)
			return l, nil

		}
	}

	return l, ErrNotImplemented
}

func (l Location) Hash() string {

	return l.GetAttributeString("hash")
}

func (l Location) Host() string {

	return l.GetAttributeString("host")
}

func (l Location) Hostname() string {

	return l.GetAttributeString("hostname")
}

func (l Location) Href() string {

	return l.GetAttributeString("href")
}

func (l Location) Origin() string {

	return l.GetAttributeString("origin")
}

func (l Location) Pathname() string {

	return l.GetAttributeString("pathname")
}

func (l Location) Port() string {

	return l.GetAttributeString("port")
}

func (l Location) Protocol() string {

	return l.GetAttributeString("protocol")
}

func (l Location) Search() string {

	return l.GetAttributeString("search")
}

func (l Location) Username() string {

	return l.GetAttributeString("username")
}

func (l Location) Password() string {

	return l.GetAttributeString("password")
}

func (l Location) Assign(url string) error {
	var err error
	err = l.Call("assign", js.ValueOf(url)).Error()
	return err
}

func (l Location) Reload(value bool) error {
	var err error
	err = l.Call("reload", js.ValueOf(value)).Error()
	return err
}

func (l Location) Replace(url string) error {
	var err error
	err = l.Call("replace", js.ValueOf(url)).Error()
	return err
}
