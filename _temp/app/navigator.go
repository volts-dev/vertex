package app

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetNavigatorInterface)
}

var ErrNotAPermissions = errors.New("Object is not a Permissions")

var navigatorinterface js.Value

// GetInterface get the JS interface of formdata
func GetNavigatorInterface() js.Value {

	singleton.Do(func() {
		if navigatorinterface = js.Global().Get("Navigator"); navigatorinterface.Error() != nil {
			navigatorinterface = js.Undefined()
		}
		js.Register(navigatorinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		GetNavigatorInterface()
		GetPermissionsInterface()
		GetServiceWorkerContainerInterface()
	})

	return navigatorinterface
}

type Navigator struct {
	js.Object
}

type NavigatorFrom interface {
	Navigator_() Navigator
}

func (n Navigator) Navigator_() Navigator {
	return n
}

func ToNavigator(obj js.Value) (Navigator, error) {
	var n Navigator

	if ni := GetNavigatorInterface(); !ni.IsUndefined() {
		if obj.InstanceOf(ni) {
			//n.BaseObject = n.SetObject(obj)
			n.SetValue(obj)
			return n, nil

		}
	}

	return n, js.ErrNotImplemented
}

func (n Navigator) CookieEnabled() (bool, error) {

	return n.GetAttributeBool("cookieEnabled")
}

func (n Navigator) Permissions() (Permissions, error) {
	var err error
	var obj interface{}
	var p Permissions
	var ok bool

	if obj, err = n.GetAttributeGlobal("permissions"); err == nil {
		if p, ok = obj.(Permissions); !ok {
			err = ErrNotAPermissions
		}
	}

	return p, err
}

/*
func (n Navigator) Credentials() (Credentials, error) {

	return n.GetAttributeBool("credentials")
}
*/

func (n Navigator) Clipboard() (Clipboard, error) {

	var err error
	var obj interface{}
	var c Clipboard
	var ok bool

	if obj, err = n.GetAttributeGlobal("clipboard"); err == nil {
		if c, ok = obj.(Clipboard); !ok {
			err = ErrNotAClipboard
		}
	}

	return c, err
}

func (n Navigator) USB() (USB, error) {

	var err error
	var obj interface{}
	var u USB
	var ok bool

	if obj, err = n.GetAttributeGlobal("usb"); err == nil {
		if u, ok = obj.(USB); !ok {
			err = ErrNotAUSB
		}
	}

	return u, err
}

func (n Navigator) DeviceMemory() (float64, error) {

	return n.GetAttributeDouble("deviceMemory")
}

func (n Navigator) HardwareConcurrency() (int, error) {

	return n.GetAttributeInt("hardwareConcurrency")
}

func (n Navigator) UserAgent() (string, error) {

	return n.GetAttributeString("userAgent")
}

func (n Navigator) Language() (string, error) {

	return n.GetAttributeString("language")
}

func (n Navigator) ServiceWorker() (ServiceWorkerContainer, error) {

	var err error
	var obj interface{}
	var s ServiceWorkerContainer
	var ok bool

	if obj, err = n.GetAttributeGlobal("serviceWorker"); err == nil {

		if s, ok = obj.(ServiceWorkerContainer); !ok {
			err = ErrNotAServiceWorkerContainer
		}

	}

	return s, err
}

func (n Navigator) Vendor() (string, error) {

	return n.GetAttributeString("vendor")
}

func (n Navigator) JavaEnabled() (bool, error) {
	return n.CallBool("javaEnabled")
}
