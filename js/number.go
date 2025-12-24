package js

import (
	"sync"
)

var (
// ErrNotImplemented = errors.New("Browser not implemented Number")
)

func init() {
	//js.RegisterInterface(GetNumberInterface)
}

var singletonnumber sync.Once

var numberinterface Value

// GetInterfaceNumber get the JS interface Array
func GetNumberInterface() Value {
	singletonnumber.Do(func() {
		if numberinterface = Global().Get("Number"); numberinterface.Error() != nil {
			numberinterface = Undefined()
		}
	})

	return numberinterface
}

// Number struct
type Number struct {
	Object
}

type NumberFrom interface {
	Number_() Number
}

func (n Number) Number_() Number {
	return n
}

func IsInteger(obj Value) (bool, error) {
	var err error
	var result bool
	if ni := GetNumberInterface(); !ni.IsUndefined() {

		if obj := ni.Call("isInteger", obj); obj.Error() == nil {

			if obj.Type() == TypeBoolean {
				return obj.Bool()
			} else {
				err = ErrObjectNotBool
			}
		}

	}
	return result, err
}
