package stream

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/promise"
)

var singletonw sync.Once

var writablestreaminterface js.Value

// GetWInterface get the JS interface WritableStream.
func GetWInterface() js.Value {

	singletonw.Do(func() {

		if writablestreaminterface = js.Global().Get("WritableStream"); writablestreaminterface.Error() != nil {
			writablestreaminterface = js.Undefined()
		}
		js.Register(writablestreaminterface, func(v js.Value) (interface{}, error) {
			return NewWriteableStreamFromJSObject(v)
		})
	})

	return writablestreaminterface
}

type WritableStream struct {
	js.Object
}

type WritableStreamFrom interface {
	WritableStream_() WritableStream
}

func (r WritableStream) WritableStream_() WritableStream {
	return r
}

// NewWritableStream Create a new NewWritableStream
func NewWritableStream() (WritableStream, error) {
	var w WritableStream
	var obj js.Value
	var err error
	if wi := GetWInterface(); !wi.IsUndefined() {

		if obj = wi.New(); obj.Error() == nil {
			w.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplementedWritableStream
	}
	return w, err
}

func NewWriteableStreamFromJSObject(obj js.Value) (WritableStream, error) {
	var w WritableStream
	var err error
	if wsi := GetWInterface(); !wsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(wsi) {
				w.SetObjectValue(obj)

			} else {
				err = ErrNotAWritableStream
			}
		}
	} else {
		err = ErrNotImplementedWritableStream
	}

	return w, err
}

func (w WritableStream) Locked() (bool, error) {
	return w.GetAttributeBool("locked")
}

func (w WritableStream) Abort(reason string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("abort", js.ValueOf(reason)); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStream) Close() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("close"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStream) GetWriter() (WritableStreamDefaultWriter, error) {
	var err error
	var obj js.Value

	if obj = w.Call("getWriter"); obj.Error() == nil {
		return NewWritableStreamDefaultWriterFromJSObject(obj)

	}
	return WritableStreamDefaultWriter{}, err
}
