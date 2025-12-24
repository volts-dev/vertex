package permissionstatus

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var permissionstatusinterface js.Value

// GetJSInterface get the JS interface of PermissionStatus
func GetInterface() js.Value {

	singleton.Do(func() {

		if permissionstatusinterface = js.Global().Get("PermissionStatus"); permissionstatusinterface.Error() != nil {
			permissionstatusinterface = js.Undefined()
		}

		js.Register(permissionstatusinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return permissionstatusinterface
}

type PermissionStatus struct {
	eventtarget.EventTarget
}

type PermissionStatusFrom interface {
	PermissionStatus_() PermissionStatus
}

func (p PermissionStatus) PermissionStatus_() PermissionStatus {
	return p
}

func NewFromJSObject(obj js.Value) (PermissionStatus, error) {
	var p PermissionStatus

	if psi := GetInterface(); !psi.IsUndefined() {
		if obj.InstanceOf(psi) {
			p.SetObjectValue(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p PermissionStatus) Name() (string, error) {

	return p.GetAttributeString("name")
}

func (p PermissionStatus) State() (string, error) {

	return p.GetAttributeString("state")
}

// deprecated
func (p PermissionStatus) Status() (string, error) {

	return p.GetAttributeString("status")
}
