package domtokenlist

//

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var domtokenlistinterface js.Value

// DOMTokenList struct
type DOMTokenList struct {
	js.Object
}

type DOMTokenListFrom interface {
	DOMTokenList_() DOMTokenList
}

func (d DOMTokenList) DOMTokenList_() DOMTokenList {
	return d
}

// GetInterface get the JS interface DOMTokenList
func GetInterface() js.Value {

	singleton.Do(func() {

		if domtokenlistinterface = js.Global().Get("DOMTokenList"); domtokenlistinterface.Error() != nil {
			domtokenlistinterface = js.Undefined()
		}
		js.Register(domtokenlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domtokenlistinterface
}

func NewFromJSObject(obj js.Value) (DOMTokenList, error) {
	var d DOMTokenList
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				d.SetObjectValue(obj)

			} else {
				err = ErrNotAnDOMTokenList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMTokenList) Item(index int) (interface{}, error) {
	var obj js.Value
	obj = d.GetObjectValue().Index(index)
	return js.GoValue(obj)

}

func (d DOMTokenList) methodGetValue(method string, value string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj = d.Call(method, js.ValueOf(value)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err
}

func (d DOMTokenList) Contains(search string) (bool, error) {

	return d.methodGetValue("contains", search)
}

func (d DOMTokenList) method(method string, tokens ...string) error {
	var err error
	var arrayJS []interface{}

	for _, token := range tokens {
		arrayJS = append(arrayJS, js.ValueOf(token))
	}

	err = d.Call(method, arrayJS...).Error()

	return err

}

func (d DOMTokenList) Add(tokens ...string) error {
	return d.method("add", tokens...)
}

func (d DOMTokenList) Remove(tokens ...string) error {
	return d.method("remove", tokens...)
}

func (d DOMTokenList) Replace(oldtoken, newtoken string) error {
	return d.method("replace", oldtoken, newtoken)
}

func (d DOMTokenList) Toggle(token string, force ...bool) (bool, error) {
	var err error
	var arrayJS []interface{}
	var result bool
	var obj js.Value

	arrayJS = append(arrayJS, js.ValueOf(token))
	if len(force) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(force[0]))
	}

	if obj = d.Call("toggle", arrayJS...); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err

}

func (d DOMTokenList) Supports(token string) (bool, error) {
	return d.methodGetValue("supports", token)
}

func (d DOMTokenList) Entries() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = d.Call("entries"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (d DOMTokenList) ForEach(f func(string)) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(js.ValueToString(args[0]))
		return nil
	})

	err = d.Call("forEach", jsfunc).Error()
	jsfunc.Release()
	return err
}

func (d DOMTokenList) Keys() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = d.Call("keys"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}

func (d DOMTokenList) Values() (js.Iterator, error) {
	var err error
	var obj js.Value
	var iter js.Iterator

	if obj = d.Call("values"); obj.Error() == nil {
		iter, err = js.NewIteratorFromJSObject(obj)
	}

	return iter, err
}
