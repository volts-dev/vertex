package baseobject

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonnumber sync.Once

var numberinterface js.Value

// GetInterfaceNumber get the JS interface Array
func GetInterfaceNumber() js.Value {

	singletonnumber.Do(func() {

		if numberinterface, err = Get(js.Global(), "Number"); err != nil {
			numberinterface = js.Undefined()
		}
	})

	return numberinterface
}

// Number struct
type Number struct {
	BaseObject
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
	if ni := GetInterfaceNumber(); !ni.IsUndefined() {

		if obj, err := Call(ni, "isInteger", obj); err == nil {

			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				err = ErrObjectNotBool
			}
		}

	}
	return result, err
}
