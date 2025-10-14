package app

import (
	"errors"

	"github.com/volts-dev/vertex/core/js"
	"github.com/volts-dev/vertex/core/js/reflect"
)

func init() {

	js.RegisterInterface(GetDOMExceptionInterface)
}

var ErrNotADOMException = errors.New("The given value must be a DOMException")
var domexceptioninterface js.Value

// DomException DomException struct
type DomException struct {
	js.Object
}

type DomExceptionFrom interface {
	DomException_() DomException
}

func (d DomException) DomException_() DomException {
	return d
}

// GetJSInterface get the JS interface
func GetDOMExceptionInterface() js.Value {

	singleton.Do(func() {
		if domexceptioninterface = js.Global().Get("DOMException"); domexceptioninterface.Error() != nil {
			domexceptioninterface = js.Undefined()
		}

		js.Register(domexceptioninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})
	return domexceptioninterface
}

func New(opts ...string) (DomException, error) {

	var e DomException
	var arrayJS []interface{}
	var obj js.Value
	var err error
	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if ei := GetDOMExceptionInterface(); !ei.IsUndefined() {

		if obj, err = reflect.New(ei, arrayJS...); err == nil {
			e.SetValue(obj)
		}

	} else {
		err = js.ErrNotImplemented
	}
	return e, err
}

func ToDomException(obj js.Value) (DomException, error) {
	var d DomException
	var err error
	if di := GetDOMExceptionInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetValue(obj)
			} else {
				err = ErrNotADOMException
			}
		}
	} else {
		err = js.ErrNotImplemented
	}

	return d, err
}

func (d DomException) Message() (string, error) {
	return d.GetAttributeString("message")
}

func (d DomException) Name() (string, error) {
	return d.GetAttributeString("name")
}
