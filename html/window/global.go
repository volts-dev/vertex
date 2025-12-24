package window

import (
	"github.com/volts-dev/vertex/js"
)

func Alert(message string) error {
	err := Default().Call("alert", js.ValueOf(message)).Error()
	return err
}

func Confirm(message string) (bool, error) {
	var ret = false
	var err error
	if obj := Default().Call("confirm", js.ValueOf(message)); obj.Error() == nil {

		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return ret, err
}

func Prompt(message, input string) (*string, error) {
	var ret *string = nil
	var err error

	if obj := Default().Call("prompt", js.ValueOf(message), js.ValueOf(input)); obj.Error() == nil {

		if obj.Type() == js.TypeString {
			v := js.ValueToString(obj)
			ret = &v
		}
	}
	return ret, err
}

func Atob(encoded string) (string, error) {

	var err error
	var result string
	var obj js.Value

	if obj = Default().Call("atob", js.ValueOf(encoded)); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrObjectNotString
		}
	}
	return result, err
}

func Btoa(message string) (string, error) {

	var err error
	var result string
	var obj js.Value

	if obj = Default().Call("btoa", js.ValueOf(message)); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = js.ErrObjectNotString
		}
	}
	return result, err
}

func Open(url string, opts ...interface{}) (Window, error) {
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(url))
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj := Default().Call("open", arrayJS...); obj.Error() == nil {
		var w Window
		w.SetObjectValue(obj)
		return w, nil
	} else {
		return Window{}, nil
	}

}
