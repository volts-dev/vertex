package gomap

import (
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/array"
	"github.com/volts-dev/vertex/js/helper"
	"github.com/volts-dev/vertex/js/object"
)

func MapFromJSObject(jsobj js.Value) interface{} {
	var retvalue interface{}

	if obj, err := object.ToObject(jsobj); err == nil {
		if ok, err := array.IsArray(obj); ok && err == nil {
			var arrayret []interface{}

			if a, err := array.NewFromJSObject(obj.GetObjectValue()); err == nil {
				if it, err := a.Entries(); err == nil {
					for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

						if obj1, ok := value.(js.ObjectFrom); !ok {
							arrayret = append(arrayret, value)

						} else {
							arrayret = append(arrayret, MapFromJSObject(obj1.Value()))
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
						i := helper.GoValue_(value)
						if obj1, ok := i.(js.ObjectFrom); !ok {
							vmap[key] = i
						} else {
							vmap[key] = MapFromJSObject(obj1.Value())
						}
					}
				}
			}
			retvalue = vmap
		}

	}

	return retvalue

}
