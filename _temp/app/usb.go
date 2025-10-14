package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var ErrNotAUSB = errors.New("Object is not a USB")
var singletonUSB sync.Once
var usbinterface js.Value

// USB struct
type USB struct {
	html.EventTarget
}

type USBFrom interface {
	USB_() USB
}

func (u USB) Clipboard_() USB {
	return u
}

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singletonUSB.Do(func() {
		if usbinterface = js.Global().Get("USB"); usbinterface.Error() != nil {
			usbinterface = js.Undefined()
		}

		js.Register(usbinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})
		GetPromiseInterface()

	})

	return usbinterface
}

func ToUSB(obj js.Value) (USB, error) {
	var u USB

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			//u.BaseObject = u.SetObject(obj)
			u.SetValue(obj)
			return u, nil

		}
	}

	return u, js.ErrNotImplemented
}

func (u USB) GetDevices() (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = u.Call("getDevices"); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}

func (u USB) RequestDevices(filter map[string]interface{}) (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = u.Call("requestDevice", js.ValueOf(filter)); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}
