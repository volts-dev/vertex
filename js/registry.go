package js

var (
	registry   map[string]func(Value) (interface{}, error)
	interfaces []func() Value
)

func RegisterInterface(f func() Value) {
	interfaces = append(interfaces, f)
}

// Register Register a construct func for type Object given
func Register(inter Value, contruct func(Value) (interface{}, error)) error {
	var err error
	var constructname string
	if registry == nil {
		registry = make(map[string]func(Value) (interface{}, error))
	}

	//registry[inter.Get("prototype").Call("toString").String()] = contruct
	if constructname, err = GetFuncName(inter); err == nil {
		registry[constructname] = contruct
	}
	return err
}

// Discover Discover the Object Given and return a Hogosuru Class if the construct is registered
func Discover(obj Value) (interface{}, error) {
	var err error
	var bobj interface{}
	var objname Value
	var objconstructor Value

	if objconstructor = obj.Get("constructor"); objconstructor.IsUndefined() {
		if objname = objconstructor.Get("name"); objconstructor.IsUndefined() {
			name, _ := objname.String()
			if f, ok := registry[name]; ok {
				var obji interface{}
				var ok bool

				if obji, err = f(obj); err == nil {
					if bobj, ok = obji.(ObjectFrom); !ok {
						err = ErrNotABaseObject
					}
				}

			} else {

				bobj = obj

			}

		} else {
			bobj = obj
		}

	} else {
		bobj = obj
	}

	return bobj, err
}
