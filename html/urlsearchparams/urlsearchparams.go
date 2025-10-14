package urlsearchparams

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/iterator"
	"github.com/volts-dev/vertex/js/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var urlsearchparamsinterface js.Value

// URLSearchParams struct
type URLSearchParams struct {
	js.Object
}

type URLSearchParamsFrom interface {
	URLSearchParams_() URLSearchParams
}

func (u URLSearchParams) URLSearchParams_() URLSearchParams {
	return u
}

// GetInterface get the JS interface URLSearchParams
func GetInterface() js.Value {

	singleton.Do(func() {

		if urlsearchparamsinterface = js.Global().Get("URLSearchParams"); urlsearchparamsinterface.Error() != nil {
			urlsearchparamsinterface = js.Undefined()
		}
		js.Register(urlsearchparamsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return urlsearchparamsinterface
}

func New(s ...string) (URLSearchParams, error) {

	var u URLSearchParams
	var err error
	var obj js.Value
	var arrayJS []interface{}

	if len(s) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(s[0]))
	}
	if hci := GetInterface(); !hci.IsUndefined() {

		if obj = hci.New(arrayJS...); obj.Error() == nil {
			u.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return u, err
}

func NewFromJSObject(obj js.Value) (URLSearchParams, error) {
	var u URLSearchParams
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				u.SetObjectValue(obj)

			} else {
				err = ErrNotAnURLSearchParams
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return u, err
}

func (u URLSearchParams) Append(name, value string) error {
	var err error
	err = u.Call("append", js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}

func (u URLSearchParams) Delete(name string) error {
	var err error
	err = u.Call("delete", js.ValueOf(name)).Error()
	return err
}

func (u URLSearchParams) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = u.Call("entries"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (u URLSearchParams) Get(name string) (string, error) {

	var err error
	var obj js.Value
	var result string

	if obj = u.Call("get", js.ValueOf(name)); obj.Error() == nil {

		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrUndefinedValue
		}

	}

	return result, err
}

func (u URLSearchParams) Has(name string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = u.Call("has", js.ValueOf(name)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return result, err
}

func (u URLSearchParams) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = u.Call("keys"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (u URLSearchParams) Set(name, value string) error {
	var err error
	err = u.Call("set", js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}

func (u URLSearchParams) Sort() error {
	var err error
	err = u.Call("sort").Error()
	return err
}

func (u URLSearchParams) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj = u.Call("values"); obj.Error() == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
