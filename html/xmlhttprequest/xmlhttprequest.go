package xmlhttprequest

// https://developer.mozilla.org/fr/docs/Web/API/XMLHttpRequest/XMLHttpRequest

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/progressevent"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var xhrinterface js.Value

// XMLHTTPRequest XMLHTTPRequest struct
type XMLHTTPRequest struct {
	js.Object
}

type XMLHTTPRequestFrom interface {
	XMLHTTPRequest_() XMLHTTPRequest
}

func (x XMLHTTPRequest) XMLHTTPRequest_() XMLHTTPRequest {
	return x
}

// GetInterface Get the JS XMLHTTPRequest Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		if xhrinterface = js.Global().Get("XMLHttpRequest"); xhrinterface.Error() != nil {
			xhrinterface = js.Undefined()
		}
		js.Register(xhrinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return xhrinterface
}

func NewFromJSObject(obj js.Value) (XMLHTTPRequest, error) {
	var x XMLHTTPRequest
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				x.SetObjectValue(obj)

			} else {
				err = ErrNotAXMLHTTPRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return x, err
}

// New Get an XML HTTP Request
func New() (XMLHTTPRequest, error) {
	var request XMLHTTPRequest
	var objnew js.Value
	var err error
	if xhri := GetInterface(); !xhri.IsUndefined() {

		if objnew = xhri.New(); objnew.Error() == nil {
			request.SetObjectValue(objnew)
		}

	} else {
		err = ErrNotImplemented
	}
	return request, err
}

func (x XMLHTTPRequest) Open(method string, url string) error {
	var err error
	err = x.Call("open", js.ValueOf(method), js.ValueOf(url)).Error()
	return err
}

func (x XMLHTTPRequest) SetRequestHeader(header string, value string) error {
	var err error
	err = x.Call("setRequestHeader", js.ValueOf(header), js.ValueOf(value)).Error()
	return err
}

// Send the form. Can accept a form data in args
func (x XMLHTTPRequest) Send(value ...interface{}) error {
	var err error
	var arrayJS []interface{}
	if len(value) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(value[0]))
	}
	err = x.Call("send", arrayJS...).Error()
	return err
}

func (x XMLHTTPRequest) Abort() error {
	var err error
	err = x.Call("abort").Error()
	return err
}

func (x XMLHTTPRequest) setHandler(jshandlername string, handler func(i interface{})) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var i interface{}
		if len(args) > 0 {
			i, _ = js.Discover(args[0])
		}
		handler(i)

		return nil
	})

	x.GetObjectValue().Set(jshandlername, jsfunc)
}

// SetOnload Set OnLoad
func (x XMLHTTPRequest) SetOnload(handler func(i interface{})) {
	x.setHandler("onload", handler)
}

// SetOnAbort Set SetOnAbort
func (x XMLHTTPRequest) SetOnAbort(handler func(i interface{})) {
	x.setHandler("onabort", handler)
}

// SetOnError Set SetOnError
func (x XMLHTTPRequest) SetOnError(handler func(i interface{})) {
	x.setHandler("onerror", handler)
}

// SetOnReadyStateChange Set SetOnReadyStateChange
func (x XMLHTTPRequest) SetOnReadyStateChange(handler func(i interface{})) {
	x.setHandler("onreadystatechange", handler)
}

// SetOnProgress Set  OnProgress
func (x XMLHTTPRequest) SetOnProgress(handler func(progressevent.ProgressEvent)) {
	onprogress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
			handler(pe)
		}

		return nil
	})

	x.GetObjectValue().Set("onprogress", onprogress)

}

func (x XMLHTTPRequest) ReadyState() (int, error) {

	return x.GetAttributeInt("readyState")
}

func (x XMLHTTPRequest) ResponseText() (string, error) {

	return x.GetAttributeString("responseText")
}

// GetResponseHeader https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (x XMLHTTPRequest) GetResponseHeader(header string) (string, error) {
	var responseHeader js.Value
	var err error
	if responseHeader = x.Call("getResponseHeader", js.ValueOf(header)); responseHeader.Error() == nil {

		if responseHeader.Type() == js.TypeString {
			return responseHeader.String()
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}

// Response
func (x XMLHTTPRequest) Response() (js.Value, error) {
	return x.GetValueByKey("response"), nil
}

func (x XMLHTTPRequest) SetResponseType(typeResponse string) {

	x.GetObjectValue().Set("responseType", js.ValueOf(typeResponse))

}

func (x XMLHTTPRequest) SetWithCredentials(withcredentials bool) {

	x.GetObjectValue().Set("withCredentials", js.ValueOf(withcredentials))

}

func (x XMLHTTPRequest) ResponseURL() (string, error) {

	return x.GetAttributeString("responseURL")
}

func (x XMLHTTPRequest) ResponseXML() (js.Value, error) {
	var responseXML js.Value
	var err error
	if responseXML = x.GetValueByKey("responseXML"); responseXML.Error() == nil {
		//return a document object : TO DO IMPLEMENTATION
		return responseXML, nil

	}
	return responseXML, err
}

func (x XMLHTTPRequest) Status() (int, error) {
	var readystate js.Value
	var err error
	if readystate = x.GetValueByKey("status"); readystate.Error() == nil {
		if readystate.Type() == js.TypeNumber {
			return readystate.Int()
		} else {
			return 0, object.ErrObjectNotNumber
		}

	}
	return 0, err
}

func (x XMLHTTPRequest) StatusText() (string, error) {
	var responseUrl js.Value
	var err error
	if responseUrl = x.GetValueByKey("statusText"); responseUrl.Error() == nil {

		if responseUrl.Type() == js.TypeString {
			return responseUrl.String()
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}

func (x XMLHTTPRequest) uploadSetHandler(jshandlername string, handler func(XMLHTTPRequest)) {
	var uploadAbstractObject js.Value

	if uploadAbstractObject = x.GetValueByKey("upload"); uploadAbstractObject.Error() == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			handler(x)

			return nil
		})

		uploadAbstractObject.Set(jshandlername, jsfunc)
	}

}

// UploadSetOnloadstart
func (x XMLHTTPRequest) UploadSetOnloadstart(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadstart", handler)

}

// UploadSetOnabort
func (x XMLHTTPRequest) UploadSetOnabort(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onabort", handler)

}

// UploadSetOnerror
func (x XMLHTTPRequest) UploadSetOnerror(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onerror", handler)

}

// UploadSetOnload
func (x XMLHTTPRequest) UploadSetOnload(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onload", handler)

}

// UploadSetOntimeout
func (x XMLHTTPRequest) UploadSetOntimeout(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("ontimeout", handler)

}

// UploadSetOnloadend
func (x XMLHTTPRequest) UploadSetOnloadend(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadend", handler)

}

// UploadSetOnprogress
func (x XMLHTTPRequest) UploadSetOnprogress(handler func(XMLHTTPRequest, progressevent.ProgressEvent)) {

	var uploadAbstractObject js.Value

	if uploadAbstractObject = x.GetValueByKey("upload"); uploadAbstractObject.Error() == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
				handler(x, pe)
			}

			return nil
		})

		uploadAbstractObject.Set("onprogress", jsfunc)
	}

}
