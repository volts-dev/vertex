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
	if v == nil {
		return 0
	}

	value, err := v.Int()
	if err != nil {
		return 0
	}

	return value
}
func ValueToBool(v Value) bool {
	if v == nil {
		return false
	}

	value, err := v.Bool()
	if err != nil {
		return false
	}

	return value
}

func ValueToString(v Value) string {
	if v == nil {
		return ""
	}

	value, err := v.String()
	if err != nil {
		return ""
	}

	return value
}

// IsPrimitive 检查值是否为原始类型（不是对象或函数）
func IsPrimitive(value Value) bool {
	if value == nil {
		return false
	}

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
