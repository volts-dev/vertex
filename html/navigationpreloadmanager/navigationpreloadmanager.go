package navigationpreloadmanager

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var navigationpreloadmanagerinterface js.Value

// GetInterface get the JS interface navigationptpreloadmanager
func GetInterface() js.Value {

	singleton.Do(func() {

		if navigationpreloadmanagerinterface = js.Global().Get("NavigationPreloadManager"); navigationpreloadmanagerinterface.Error() != nil {
			navigationpreloadmanagerinterface = js.Undefined()
		}
		js.Register(navigationpreloadmanagerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()

	})

	return navigationpreloadmanagerinterface
}

type NavigationPreloadManager struct {
	js.Object
}

type NavigationPreloadManagerFrom interface {
	NavigationPreloadManager_() NavigationPreloadManager
}

func (n NavigationPreloadManager) NavigationPreloadManager_() NavigationPreloadManager {
	return n
}

func NewFromJSObject(obj js.Value) (NavigationPreloadManager, error) {
	var n NavigationPreloadManager

	if ni := GetInterface(); !ni.IsUndefined() {
		if obj.InstanceOf(ni) {
			n.SetObjectValue(obj)
			return n, nil

		}
	}

	return n, ErrNotImplemented
}

func (n NavigationPreloadManager) Enable() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = n.Call("enable"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) Disable() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = n.Call("disable"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) SetHeaderValue(value string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = n.Call("setHeaderValue", js.ValueOf(value)); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) GetState() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = n.Call("getState"); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}
