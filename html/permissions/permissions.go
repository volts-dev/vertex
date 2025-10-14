package permissions

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

var permissionsinterface js.Value

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

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

func NewFromJSObject(obj js.Value) (Permissions, error) {
	var p Permissions

	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.InstanceOf(pi) {
			p.SetObjectValue(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p Permissions) Query(permissiondescriptor map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var pro promise.Promise

	if obj = p.Call("query", js.ValueOf(permissiondescriptor)); obj.Error() == nil {

		pro, err = promise.NewFromJSObject(obj)
	}

	return pro, err

}
