package jserror

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var errorinterface js.Value

// JSError JSError struct
type JSError struct {
	js.Object
}

type JSErrorFrom interface {
	JSError_() JSError
}

func (e JSError) JSError_() JSError {
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

func New(values ...interface{}) (JSError, error) {
	var e JSError
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
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (JSError, error) {
	var e JSError
	var err error
	if ei := GetInterface(); !ei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ei) {
				e.SetObjectValue(obj)
			} else {
				err = ErrNotAnError
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func (j JSError) Message() (string, error) {
	return j.GetAttributeString("message")
}

func (j JSError) SetMessage(value string) error {
	return j.SetAttributeString("message", value)
}

func (j JSError) Name() (string, error) {
	return j.GetAttributeString("name")
}
func (j JSError) SetName(value string) error {
	return j.SetAttributeString("name", value)
}

func (j JSError) Stack() (string, error) {
	return j.GetAttributeString("stack")
}
