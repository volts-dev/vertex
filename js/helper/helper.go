package helper

import (
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/number"
)

func GoValue_(object js.Value) interface{} {
	var i interface{}
	var err error

	if i, err = GoValue(object); err != nil {
		js.Debug(err.Error())
	}

	return i
}

func GoValue(object js.Value) (interface{}, error) {
	var err error
	switch object.Type() {
	case js.TypeNumber:

		if v, err := number.IsInteger(object); err == nil && v {
			v, _ := object.Float()
			return int64(v), nil
		}
		return object.Float()
	case js.TypeString:
		return object.String()
	case js.TypeBoolean:
		return object.Bool()
	case js.TypeNull:
		return nil, nil
	}

	obj, err := js.Discover(object)

	return obj, err
}

func GoValueExt(object js.Value) interface{} {
	var i interface{}
	var err error

	if i, err = GoValue(object); err != nil {
		//debug(err.Error())
	}

	return i
}

func ValueToInt(v js.Value) int {
	value, err := v.Int()
	if err != nil {
		//debug(err.Error())
	}

	return value
}
func ValueToBool(v js.Value) bool {
	value, err := v.Bool()
	if err != nil {
		//debug(err.Error())
	}

	return value
}

func ValueToString(v js.Value) string {
	value, err := v.String()
	if err != nil {
		//debug(err.Error())
	}

	return value
}
