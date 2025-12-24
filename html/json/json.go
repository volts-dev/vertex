package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"github.com/volts-dev/vertex/html/gomap"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var jsoninterface js.Value

// Json  struct
type Json struct {
	js.Object
}

type JsonFrom interface {
	Json_() Json
}

func (i Json) Json_() Json {
	return i
}

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		if jsoninterface = js.Global().Get("JSON"); jsoninterface.Error() != nil {
			jsoninterface = js.Undefined()
		}
		js.Register(jsoninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return jsoninterface
}

func Parse(data string) (Json, error) {

	var jsonObject js.Value
	var err error
	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if jsonObject = jsoni.Call("parse", data); jsonObject.Error() != nil {
			return Json{}, err
		} else {
			return NewFromJSObject(jsonObject)
		}

	} else {
		err = ErrNotImplemented
	}

	return Json{}, err
}

func NewFromJSObject(obj js.Value) (Json, error) {
	var j Json

	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		j.SetObjectValue(obj)
		return j, nil

	}

	return j, ErrNotAJson

}

func (j Json) Map() interface{} {

	return gomap.MapFromJSObject(j.GetObjectValue())

}

func Stringify(opts ...interface{}) (string, error) {

	var arrayJS []interface{}
	var err error
	var stringObject js.Value

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if stringObject = jsoni.Call("stringify", arrayJS); stringObject.Error() != nil {
			return "", err
		} else {

			return stringObject.String()
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}

func StringifyObject(object interface{}) (string, error) {

	var err error
	var stringObject js.Value

	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if stringObject = jsoni.Call("stringify", js.ValueOf(object)); stringObject.Error() != nil {
			return "", err
		} else {

			return stringObject.String()
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
