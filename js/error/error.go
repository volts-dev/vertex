package errors

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var (
	ErrNotAnError = errors.New("Object is not Error")
	ErrNotAnArray = errors.New("The given value must be an Array")
)

var errorinterface js.Value
var singleton sync.Once

// Error Error struct
type Error struct {
	js.Object
}

type ErrorFrom interface {
	Error_() Error
}

func (e Error) Error_() Error {
	return e
}

// GetInterface get the Error interface
func GetInterface() js.Value {

	singleton.Do(func() {
		if errorinterface = js.Global().Get("Error"); errorinterface.Error() != nil {
			errorinterface = js.Undefined()
		}

		js.Register(errorinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})
	return errorinterface
}

func NewError(values ...interface{}) (Error, error) {
	var e Error
	var objs []interface{}
	var obj js.Value
	var err error
	if len(values) == 1 {
		switch values[0].(type) {
		case string:
			objs = append(objs, values[0])
		case error:
			objs = append(objs, values[0].(error).Error())
		}
	}

	if ei := GetInterface(); !ei.IsUndefined() {

		if obj = ei.New(objs...); obj.Error() == nil {
			e.SetObjectValue(obj)
		}
	} else {
		err = js.ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (Error, error) {
	var e Error
	var err error
	if ei := GetInterface(); !ei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ei) {
				//e.SetObjectValue(obj)
				e.SetObjectValue(obj)
			} else {
				err = ErrNotAnError
			}
		}
	} else {
		err = js.ErrNotImplemented
	}

	return e, err
}

func (j Error) Message() (string, error) {
	return j.GetAttributeString("message")
}

func (j Error) SetMessage(value string) error {
	return j.SetAttributeString("message", value)
}

func (j Error) Name() (string, error) {
	return j.GetAttributeString("name")
}
func (j Error) SetName(value string) error {
	return j.SetAttributeString("name", value)
}

func (j Error) Stack() (string, error) {
	return j.GetAttributeString("stack")
}
