package domexception

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

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
func GetInterface() js.Value {

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

	if ei := GetInterface(); !ei.IsUndefined() {

		if obj = ei.New(arrayJS...); obj.Error() == nil {
			e.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (DomException, error) {
	var d DomException
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetObjectValue(obj)
			} else {
				err = ErrNotADOMException
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DomException) Message() (string, error) {
	return d.GetAttributeString("message")
}

func (d DomException) Name() (string, error) {
	return d.GetAttributeString("name")
}
