package usb

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/promise"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var usbinterface js.Value

// USB struct
type USB struct {
	eventtarget.EventTarget
}

type USBFrom interface {
	USB_() USB
}

func (u USB) Clipboard_() USB {
	return u
}

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if usbinterface = js.Global().Get("USB"); usbinterface.Error() != nil {
			usbinterface = js.Undefined()
		}

		js.Register(usbinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})
		promise.GetInterface()

	})

	return usbinterface
}

func NewFromJSObject(obj js.Value) (USB, error) {
	var u USB

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			u.SetObjectValue(obj)
			return u, nil

		}
	}

	return u, ErrNotImplemented
}

func (u USB) GetDevices() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = u.Call("getDevices"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (u USB) RequestDevices(filter map[string]interface{}) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = u.Call("requestDevice", js.ValueOf(filter)); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
