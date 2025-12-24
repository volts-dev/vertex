package pushmanager

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/promise"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var pushmanagerinterface js.Value

// GetInterface get the JS interface pushmanager
func GetInterface() js.Value {

	singleton.Do(func() {

		if pushmanagerinterface = js.Global().Get("PushManager"); pushmanagerinterface.Error() != nil {
			pushmanagerinterface = js.Undefined()
		}
		js.Register(pushmanagerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		js.GetArrayInterface()
		promise.GetInterface()
	})

	return pushmanagerinterface
}

type PushManager struct {
	js.Object
}

type PushManagerFrom interface {
	PushManagerManager_() PushManager
}

func (p PushManager) PushManager_() PushManager {
	return p
}

func NewFromJSObject(obj js.Value) (PushManager, error) {
	var p PushManager

	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.InstanceOf(pi) {
			p.SetObjectValue(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p PushManager) SupportedContentEncodings() (js.Array, error) {

	var err error
	var obj interface{}
	var a js.Array
	var ok bool

	if obj, err = p.GetAttributeGlobal("supportedContentEncodings"); err == nil {
		if a, ok = obj.(js.Array); !ok {
			err = js.ErrNotAnArray
		}
	}

	return a, err
}

func (p PushManager) GetSubscription() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj = p.Call("getSubscription"); obj.Error() == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}

func (p PushManager) PermissionState() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj = p.Call("permissionState"); obj.Error() == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}

func (p PushManager) Subscribe() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj = p.Call("subscribe"); obj.Error() == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}
