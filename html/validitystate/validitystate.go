package validitystate

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var validitystateinterface js.Value

// ValidityState struct
type ValidityState struct {
	js.Object
}

type ValidityStateFrom interface {
	ValidityState_() ValidityState
}

func (v ValidityState) ValidityState_() ValidityState {
	return v
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if validitystateinterface = js.Global().Get("ValidityState"); validitystateinterface.Error() != nil {
			validitystateinterface = js.Undefined()
		}
		js.Register(validitystateinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return validitystateinterface
}

func NewFromJSObject(obj js.Value) (ValidityState, error) {
	var v ValidityState
	var err error
	if hei := GetInterface(); !hei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hei) {

				v.SetObjectValue(obj)

			} else {
				err = ErrNotAnValidityState
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return v, err
}

func (v ValidityState) BadInput() (bool, error) {
	return v.GetAttributeBool("badInput")
}

func (v ValidityState) CustomError() (bool, error) {
	return v.GetAttributeBool("customError")
}

func (v ValidityState) PatternMismatch() (bool, error) {
	return v.GetAttributeBool("patternMismatch")
}

func (v ValidityState) RangeOverflow() (bool, error) {
	return v.GetAttributeBool("rangeOverflow")
}

func (v ValidityState) RangeUnderflow() (bool, error) {
	return v.GetAttributeBool("rangeUnderflow")
}

func (v ValidityState) StepMismatch() (bool, error) {
	return v.GetAttributeBool("stepMismatch")
}

func (v ValidityState) TooLong() (bool, error) {
	return v.GetAttributeBool("tooLong")
}

func (v ValidityState) TooShort() (bool, error) {
	return v.GetAttributeBool("tooShort")
}

func (v ValidityState) TypeMismatch() (bool, error) {
	return v.GetAttributeBool("typeMismatch")
}

func (v ValidityState) Valid() (bool, error) {
	return v.GetAttributeBool("valid")
}

func (v ValidityState) ValueMissing() (bool, error) {
	return v.GetAttributeBool("valueMissing")
}
