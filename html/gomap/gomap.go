package gomap

import (
	"github.com/volts-dev/vertex/js"
)

func MapFromJSObject(jsobj js.Value) interface{} {
	var retvalue interface{}

	if obj, err := js.ToObject(jsobj); err == nil {
		if ok, err := js.IsArray(obj); ok && err == nil {
			var arrayret []interface{}

			if a, err := js.NewArrayFromJSObject(obj.GetObjectValue()); err == nil {
				if it, err := a.Entries(); err == nil {
					for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

						if obj1, ok := value.(js.ObjectFrom); !ok {
							arrayret = append(arrayret, value)

						} else {
							arrayret = append(arrayret, MapFromJSObject(obj1.GetObjectValue()))
						}

					}

				}

			}

			retvalue = arrayret

		} else {
			vmap := make(map[string]interface{})
			keys, _ := obj.Keys()
			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {
				if key, ok := vkey.(string); ok {
					if value := jsobj.Get(key); value.Error() == nil {
						i := js.GoValue_(value)
						if obj1, ok := i.(js.ObjectFrom); !ok {
							vmap[key] = i
						} else {
							vmap[key] = MapFromJSObject(obj1.GetObjectValue())
						}
					}
				}
			}
			retvalue = vmap
		}

	}

	return retvalue

}
