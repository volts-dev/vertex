package formdata

// https://developer.mozilla.org/fr/docs/Web/API/FormData

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/htmlformelement"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var formadatainterface js.Value

// FormData struct
type FormData struct {
	js.Object
}

type FormDataFrom interface {
	FormData_() FormData
}

func (f FormData) FormData_() FormData {
	return f
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if formadatainterface = js.Global().Get("FormData"); formadatainterface.Error() != nil {
			formadatainterface = js.Undefined()
		}

	})

	return formadatainterface
}

func NewFromJSObject(obj js.Value) (FormData, error) {
	var f FormData
	var err error
	if fi := GetInterface(); !fi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fi) {
				f.SetObjectValue(obj)

			} else {
				err = ErrNotAFormData
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return f, err
}

func New(f ...htmlformelement.HtmlFormElement) (FormData, error) {

	var formdata FormData
	var obj js.Value
	var err error
	var opt []interface{}

	if fci := GetInterface(); !fci.IsUndefined() {
		if len(f) > 0 {
			opt = append(opt, f[0].GetObjectValue())
		}
		if obj = fci.New(opt...); obj.Error() == nil {
			formdata.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return formdata, err
}

func (f FormData) Append(key string, value interface{}) error {
	var err error
	err = f.Call("append", js.ValueOf(key), js.ValueOf(value)).Error()
	return err
}

func (f FormData) Delete(key string) error {
	var err error

	err = f.Call("delete", js.ValueOf(key)).Error()
	return err
}

func (f FormData) Entries() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = f.Call("entries"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (f FormData) Get(key string) (interface{}, error) {

	var err error
	var obj js.Value
	var result interface{}

	if obj = f.Call("get", js.ValueOf(key)); obj.Error() == nil {
		if obj.IsNull() {
			err = ErrNotAFormValueNotFound
		} else {
			result, err = js.GoValue(obj)
		}

	}
	return result, err
}

func (f FormData) Has(key string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = f.Call("has", js.ValueOf(key)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (f FormData) Keys() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = f.Call("keys"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (f FormData) Set(key string, value interface{}) error {
	var err error
	err = f.Call("set", js.ValueOf(key), js.ValueOf(value)).Error()
	return err
}

func (f FormData) Values() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = f.Call("values"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}
