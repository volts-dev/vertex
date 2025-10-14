package stream

// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/js/array"
)

func init() {

	initinterface.RegisterInterface(GetRInterface)
	initinterface.RegisterInterface(GetWInterface)
	initinterface.RegisterInterface(GetReadableStreamDefaultReaderInterface)
	initinterface.RegisterInterface(GetWritableStreamDefaultWriterInterface)

}

var singletonr sync.Once

var readablestreaminterface js.Value

// GetRInterface get the JS interface ReadableStream.
func GetRInterface() js.Value {

	singletonr.Do(func() {

		if readablestreaminterface = js.Global().Get("ReadableStream"); readablestreaminterface.Error() != nil {
			readablestreaminterface = js.Undefined()
		}
		js.Register(readablestreaminterface, func(v js.Value) (interface{}, error) {
			return NewReadableStreamFromJSObject(v)
		})
	})

	return readablestreaminterface
}

type ReadableStream struct {
	js.Object
}

type ReadableStreamFrom interface {
	ReadableStream_() ReadableStream
}

func (r ReadableStream) ReadableStream_() ReadableStream {
	return r
}

func (r ReadableStream) Locked() (bool, error) {
	return r.GetAttributeBool("locked")
}

// NewReadableStream Create a new ReadableStream
func NewReadableStream() (ReadableStream, error) {
	var r ReadableStream
	var obj js.Value
	var err error
	if ri := GetRInterface(); !ri.IsUndefined() {

		if obj = ri.New(); obj.Error() == nil {
			r.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplementedReadableStream
	}
	return r, err
}

func NewReadableStreamFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream
	var err error
	if rsi := GetRInterface(); !rsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(rsi) {
				r.SetObjectValue(obj)

			} else {
				err = ErrNotAReadableStream
			}
		}
	} else {
		err = ErrNotImplementedReadableStream
	}

	return r, err
}

func (r ReadableStream) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = r.Call("cancel"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (r ReadableStream) GetReader() (ReadableStreamDefaultReader, error) {
	var err error
	var obj js.Value

	if obj = r.Call("getReader"); obj.Error() == nil {
		return NewReadableStreamDefaultReaderFromJSObject(obj)

	}
	return ReadableStreamDefaultReader{}, err

}

func (r ReadableStream) Tee() ([]ReadableStream, error) {
	var err error
	var obj js.Value
	var ret []ReadableStream
	var a array.Array

	if obj = r.Call("tee"); obj.Error() == nil {

		if a, err = array.NewFromJSObject(obj); err == nil {

			a.ForEach(func(i interface{}) {

				if r, ok := i.(ReadableStreamFrom); ok {
					ret = append(ret, r.ReadableStream_())
				}

			})
		}

	}
	return ret, err

}

func (r ReadableStream) PipeThrough(t TransformStream, options ...map[string]string) (ReadableStream, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var transformread ReadableStream

	arrayJS = append(arrayJS, t.GetObjectValue())

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = r.Call("pipeThrough", arrayJS...); obj.Error() == nil {

		transformread, err = NewReadableStreamFromJSObject(obj)

	}

	return transformread, err
}

func (r ReadableStream) PipeTo(w WritableStream, options ...map[string]string) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var finalpromise promise.Promise

	arrayJS = append(arrayJS, w.GetObjectValue())

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj = r.Call("pipeTo", arrayJS...); obj.Error() == nil {

		finalpromise, err = promise.NewFromJSObject(obj)

	}

	return finalpromise, err
}
