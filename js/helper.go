package js

func GoValue(value Value) any {
	switch value.Type() {
	case TypeNumber:

		if IsInteger(value) {
			v := value.Float()
			return int64(v)
		}
		return value.Float()
	case TypeString:
		return value.String()
	case TypeBoolean:
		return value.Bool()
	case TypeNull:
		return nil
	}
	obj, _ := Discover(value)

	return obj
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
