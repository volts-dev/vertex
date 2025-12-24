package headers

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Headers

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var headersinterface js.Value

// History struct
type Headers struct {
	js.Object
}

type HeadersFrom interface {
	Headers_() Headers
}

func (h Headers) Headers_() Headers {
	return h
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if headersinterface = js.Global().Get("Headers"); headersinterface.Error() != nil {
			headersinterface = js.Undefined()
		}

		js.Register(headersinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return headersinterface
}

func New() (Headers, error) {

	var h Headers
	var err error
	var obj js.Value

	if hci := GetInterface(); !hci.IsUndefined() {

		if obj = hci.New(); obj.Error() == nil {
			h.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func NewFromJSObject(obj js.Value) (Headers, error) {
	var h Headers
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHeaders
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h Headers) Append(name, value string) error {
	var err error
	err = h.Call("append", js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}

func (h Headers) Delete(name string) error {
	var err error
	err = h.Call("delete", js.ValueOf(name)).Error()
	return err
}

func (h Headers) Entries() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = h.Call("entries"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (h Headers) Get(name string) (string, error) {

	var err error
	var obj js.Value
	var result string

	if obj = h.Call("get", js.ValueOf(name)); obj.Error() == nil {

		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrUndefinedValue
		}

	}

	return result, err
}

func (h Headers) Has(name string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = h.Call("has", js.ValueOf(name)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (h Headers) Keys() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = h.Call("keys"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (h Headers) Set(name, value string) error {
	var err error
	err = h.Call("set", js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}

func (h Headers) Values() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = h.Call("values"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}
