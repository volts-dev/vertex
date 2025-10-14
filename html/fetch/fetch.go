package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/response"
	"github.com/volts-dev/vertex/js"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var fetchinterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		if fetchinterface = js.Global().Get("fetch"); fetchinterface.Error() != nil {
			fetchinterface = js.Undefined()
		}

		response.GetInterface()
		js.Register(fetchinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return fetchinterface
}

// Fetch struct
type Fetch struct {
	promise.Promise
}

type FetchFrom interface {
	Fetch_() Fetch
}

func (f Fetch) Fetch_() Fetch {
	return f
}

func New(urlfetch string, opts ...interface{}) (Fetch, error) {
	var arrayJS []interface{}
	var f Fetch
	var err error
	arrayJS = append(arrayJS, urlfetch)
	for _, value := range opts {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		promisefetchobj := fetchi.Invoke(arrayJS...)
		f.SetObjectValue(promisefetchobj)
	} else {
		err = ErrNotImplemented
	}
	return f, err

}

func NewFromJSObject(obj js.Value) (Fetch, error) {
	var h Fetch
	var err error
	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fetchi) {

				h.SetObjectValue(obj)
			} else {
				err = ErrNotAFetch
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (f Fetch) Then(resolve func(response.Response) *promise.Promise, reject func(error)) (promise.Promise, error) {

	return f.Promise.Then(func(obj interface{}) *promise.Promise {
		var resp interface{}

		var err error
		if bo, ok := obj.(js.ObjectFrom); ok {
			if resp, err = js.Discover(bo.Value()); err == nil {

				if r, ok := resp.(response.ResponseFrom); ok {
					return resolve(r.Response_())
				}
			} else {
				if reject != nil {
					reject(err)
				}

			}
		}

		return nil

	}, func(e error) {
		if reject != nil {
			reject(e)
		}

	})

}
