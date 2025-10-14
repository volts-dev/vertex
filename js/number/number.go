package number

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
)

func init() {
	initinterface.RegisterInterface(GetNumberInterface)
}

var singletonnumber sync.Once

var numberinterface js.Value

// GetInterfaceNumber get the JS interface Array
func GetNumberInterface() js.Value {
	singletonnumber.Do(func() {
		if numberinterface = js.Global().Get("Number"); numberinterface.Error() != nil {
			numberinterface = js.Undefined()
		}
	})

	return numberinterface
}

// Number struct
type Number struct {
	js.Object
}

type NumberFrom interface {
	Number_() Number
}

func (n Number) Number_() Number {
	return n
}

func IsInteger(obj js.Value) (bool, error) {
	var err error
	var result bool
	if ni := GetNumberInterface(); !ni.IsUndefined() {

		if obj := ni.Call("isInteger", obj); obj.Error() == nil {

			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				err = js.ErrObjectNotBool
			}
		}

	}
	return result, err
}
