package response

// https://developer.mozilla.org/fr/docs/Web/API/Response

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/arraybuffer"
	"github.com/volts-dev/vertex/html/headers"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/stream"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var (
	ErrNotAnFResp = errors.New("The given value must be an fetch response")
)

var singleton sync.Once

var responseinterface js.Value

// FetchResponse struct
type Response struct {
	js.Object
}

type ResponseFrom interface {
	Response_() Response
}

func (r Response) Response_() Response {
	return r
}

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if responseinterface = js.Global().Get("Response"); responseinterface.Error() != nil {
			responseinterface = js.Undefined()
		}
		js.Register(responseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		arraybuffer.GetInterface()
	})

	return responseinterface
}

// New Create a response
func New(opts ...interface{}) (Response, error) {
	var r Response
	var obj js.Value
	var err error
	var array []interface{}

	for _, opt := range opts {
		if v, ok := opt.(js.ObjectFrom); ok {
			array = append(array, v.GetObjectValue())
		}

	}
	if ri := GetInterface(); !ri.IsUndefined() {

		if obj = ri.New(array...); obj.Error() == nil {
			r.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return r, err
}

func Error() (Response, error) {

	var response Response
	var err error
	var obj js.Value

	if ri := GetInterface(); !ri.IsUndefined() {

		if obj = ri.Call("error"); obj.Error() == nil {
			response.SetObjectValue(obj)

		} else {
			err = ErrNotAnFResp
		}

	} else {
		err = ErrNotImplemented
	}

	return response, err

}

func NewFromJSObject(obj js.Value) (Response, error) {
	var response Response
	var err error
	if ri := GetInterface(); !ri.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ri) {
				response.SetObjectValue(obj)

			} else {
				err = ErrNotAnFResp
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return response, err
}

func (r Response) Ok() (bool, error) {

	var err error
	var obj js.Value

	if obj = r.GetValueByKey("ok"); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return false, err
}

func (r Response) Redirected() (bool, error) {
	return r.GetAttributeBool("redirected")
}

func (r Response) Status() (int, error) {
	var code int
	var err error
	if code, err = r.GetAttributeInt("status"); err != nil {
		code = 456
	}
	return code, err
}

func (r Response) StatusText() (string, error) {

	var err error
	var obj js.Value

	if obj = r.GetValueByKey("statusText"); obj.Error() == nil {

		return obj.String()
	}
	return "", err
}

func (r Response) Type() (string, error) {

	var err error
	var obj js.Value

	if obj = r.GetValueByKey("type"); obj.Error() == nil {

		return obj.String()
	}
	return "", err
}

func (r Response) Url() (string, error) {

	var err error
	var obj js.Value

	if obj = r.GetValueByKey("url"); obj.Error() == nil {

		return obj.String()
	}
	return "", err
}

func (r Response) BodyUsed() (bool, error) {

	return r.GetAttributeBool("bodyUsed")
}

func (r Response) Text() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject = r.Call("text"); promiseObject.Error() == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err
}

func (r Response) Json() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject = r.Call("json"); promiseObject.Error() == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err
}

/* not exist on chrome
func (r Response) UseFinalURL() (bool, error) {

	return r.GetAttributeBool("useFinalURL")
}

func (r Response) SetUseFinalURL(b bool) {

	r.Value( )().Set("useFinalURL", js.ValueOf(b))
}*/

func (r Response) ArrayBuffer() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject = r.Call("arrayBuffer"); promiseObject.Error() == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err

}

func (r Response) Blob() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject = r.Call("blob"); promiseObject.Error() == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err

}

func (r Response) Headers() (headers.Headers, error) {
	var obj js.Value
	var err error
	var h headers.Headers
	if obj = r.GetValueByKey("headers"); obj.Error() == nil {
		h, err = headers.NewFromJSObject(obj)

	}
	return h, err
}

func (r Response) Body() (stream.ReadableStream, error) {
	var obj js.Value
	var err error
	var s stream.ReadableStream
	if obj = r.GetValueByKey("body"); obj.Error() == nil {
		s, err = stream.NewReadableStreamFromJSObject(obj)

	}
	return s, err
}

func (r Response) Clone() (Response, error) {

	var cloneObject js.Value
	var clone Response
	var err error
	if cloneObject = r.Call("clone"); cloneObject.Error() == nil {
		clone, err = NewFromJSObject(cloneObject)
	}
	return clone, err
}
