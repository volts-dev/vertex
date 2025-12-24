package js

func GoValue_(object Value) interface{} {
	var i interface{}
	var err error

	if i, err = GoValue(object); err != nil {
		Debug(err.Error())
	}

	return i
}

func GoValue(value Value) (interface{}, error) {
	var err error
	switch value.Type() {
	case TypeNumber:

		if v, err := IsInteger(value); err == nil && v {
			v, _ := value.Float()
			return int64(v), nil
		}
		return value.Float()
	case TypeString:
		return value.String()
	case TypeBoolean:
		return value.Bool()
	case TypeNull:
		return nil, nil
	}

	obj, err := Discover(value)

	return obj, err
}

func GoValueExt(object Value) interface{} {
	var i interface{}
	var err error

	if i, err = GoValue(object); err != nil {
		//debug(err.Error())
	}

	return i
}

func ValueToInt(v Value) int {
	value, err := v.Int()
	if err != nil {
		//debug(err.Error())
	}

	return value
}
func ValueToBool(v Value) bool {
	value, err := v.Bool()
	if err != nil {
		//debug(err.Error())
	}

	return value
}

func ValueToString(v Value) string {
	value, err := v.String()
	if err != nil {
		//debug(err.Error())
	}

	return value
}

func IsPrimitive(value Value) bool {
	typ := value.Type()

	return typ == TypeNull || typ == TypeUndefined ||
		typ == TypeBoolean || typ == TypeNumber || typ == TypeString || typ != TypeFunction || typ != TypeObject
}

func ToValues(args ...interface{}) []any {
	if len(args) == 0 {
		return nil
	}

	values := make([]any, len(args))
	for i, arg := range args {
		values[i] = unwrap(arg)
	}
	return values
}
