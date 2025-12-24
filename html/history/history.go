package history

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

// https://developer.mozilla.org/fr/docs/Web/API/History_API

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

// var historyinterface *JSInterface

var historyinterface js.Value

//JSInterface JSInterface struct
// type JSInterface struct {
// 	objectInterface js.Value
// }

// History struct
type History struct {
	js.Object
}

type HistoryFrom interface {
	History_() History
}

func (h History) History_() History {
	return h
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if historyinterface = js.Global().Get("History"); historyinterface.Error() != nil {
			historyinterface = js.Undefined()
		}
		js.Register(historyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return historyinterface
}

func NewFromJSObject(obj js.Value) (History, error) {
	var h History
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnHistory
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h History) Forward() error {
	var err error
	err = h.Call("forward").Error()
	return err
}

func (h History) Back() error {
	var err error
	err = h.Call("back").Error()
	return err
}

func (h History) Go(position int) error {
	return h.Call("go", js.ValueOf(position)).Error()
}

func (h History) Length() (int, error) {
	return h.GetValueByKey("length").Int()
}

func (h History) PushState(obj interface{}, title string, url ...string) error {
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(obj), js.ValueOf(title))

	if len(url) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(url[0]))
	}

	err = h.Call("pushState", arrayJS...).Error()

	return err
}

func (h History) ReplaceState(obj interface{}, title string, url ...string) error {
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(obj), js.ValueOf(title))

	if len(url) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(url[0]))
	}

	err = h.Call("replaceState", arrayJS...).Error()

	return err
}

func (h History) State() (interface{}, error) {
	var err error
	var obj js.Value
	var ret interface{}

	if obj = h.GetValueByKey("state"); obj.Error() == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			ret, err = js.Discover(obj)
		}

	}

	return ret, err
}
