package vertex

import "github.com/volts-dev/vertex/core/js"

var (
	registry   map[string]func(js.Value) (interface{}, error)
	interfaces []func() js.Value
)

func RegisterInterface(f func() js.Value) {
	interfaces = append(interfaces, f)
}

// Register Register a construct func for type Object given
func Register(inter js.Value, contruct func(js.Value) (interface{}, error)) error {
	var err error
	var constructname string
	if registry == nil {
		registry = make(map[string]func(js.Value) (interface{}, error))
	}

	//registry[inter.Get("prototype").Call("toString").String()] = contruct
	if constructname, err = GetFuncName(inter); err == nil {
		registry[constructname] = contruct
	}
	return err
}

// Discover Discover the Object Given and return a Hogosuru Class if the construct is registered
func Discover(obj js.Value) (interface{}, error) {
	var err error
	var bobj interface{}
	var objname js.Value
	var objconstructor js.Value

	if objconstructor = obj.Get("constructor"); objconstructor.IsUndefined() {
		if objname = objconstructor.Get("name"); objconstructor.IsUndefined() {
			if f, ok := registry[objname.String()]; ok {
				var obji interface{}
				var ok bool

				if obji, err = f(obj); err == nil {
					if bobj, ok = obji.(ObjectFrom); !ok {
						err = ErrNotABaseObject
					}
				}

			} else {

				bobj, err = ToObject(obj)

			}

		} else {
			bobj, err = ToObject(obj)
		}

	} else {
		bobj, err = ToObject(obj)
	}

	return bobj, err
}
