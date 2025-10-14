package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetPermissionsInterface)
}

var singletonPermissions sync.Once

var permissionsinterface js.Value

// GetInterface get the JS interface of formdata
func GetPermissionsInterface() js.Value {

	singletonPermissions.Do(func() {
		if permissionsinterface = js.Global().Get("Permissions"); permissionsinterface.Error() != nil {
			permissionsinterface = js.Undefined()
		}
		js.Register(permissionsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return permissionsinterface
}

type Permissions struct {
	js.Object
}

type PermissionsFrom interface {
	Permissions_() Permissions
}

func (p Permissions) Permissions_() Permissions {
	return p
}

func ToPermissions(obj js.Value) (Permissions, error) {
	var p Permissions

	if pi := GetPermissionsInterface(); !pi.IsUndefined() {
		if obj.InstanceOf(pi) {
			//p.BaseObject = p.SetObject(obj)
			p.SetValue(obj)
			return p, nil

		}
	}

	return p, js.ErrNotImplemented
}

func (p Permissions) Query(permissiondescriptor map[string]interface{}) (Promise, error) {

	var err error
	var obj js.Value
	var pro Promise

	if obj = p.Call("query", js.ValueOf(permissiondescriptor)); obj.Error() == nil {

		pro, err = ToPromise(obj)
	}

	return pro, err

}
