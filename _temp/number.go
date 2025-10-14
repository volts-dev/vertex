package js

import (
	"sync"
)

var singletonnumber sync.Once

var numberinterface Value

// GetInterfaceNumber get the JS interface Array
func GetNumberInterface() Value {
	singletonnumber.Do(func() {
		if numberinterface = Global().Get("Number"); numberinterface.IsUndefined() {
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
