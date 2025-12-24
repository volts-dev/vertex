package stream

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/typedarray"
)

var singletonWritableStreamDefault sync.Once

var writeablestreamdefaultinterface js.Value

// GetWritableStreamDefaultWriterInterface
func GetWritableStreamDefaultWriterInterface() js.Value {

	singletonReadableStreamDefault.Do(func() {

		if writeablestreamdefaultinterface = js.Global().Get("WritableStreamDefaultWriter"); writeablestreamdefaultinterface.Error() != nil {
			writeablestreamdefaultinterface = js.Undefined()
		}
	})

	return writeablestreamdefaultinterface
}

type WritableStreamDefaultWriter struct {
	js.Object
}

type WritableStreamDefaultWriterFrom interface {
	WritableStreamDefaultWriter_() WritableStreamDefaultWriter
}

func (w WritableStreamDefaultWriter) WritableStreamDefaultWriter_() WritableStreamDefaultWriter {
	return w
}

func NewWritableStreamDefaultWriterFromJSObject(obj js.Value) (WritableStreamDefaultWriter, error) {
	var w WritableStreamDefaultWriter

	if rsi := GetWritableStreamDefaultWriterInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			w.SetObjectValue(obj)
			return w, nil

		}
	}

	return w, ErrNotAWritableStream
}

func (w WritableStreamDefaultWriter) Closed() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.GetValueByKey("closed"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) DesiredSize() (int, error) {
	return w.GetAttributeInt("desiredSize")
}

func (w WritableStreamDefaultWriter) Ready() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.GetValueByKey("ready"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) Abort(reason string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("abort", js.ValueOf(reason)); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) Close() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("close"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) ReleaseLock() error {

	err := w.Call("releaseLock").Error()
	return err
}

func (w WritableStreamDefaultWriter) Write(chunk typedarray.Uint8Array) (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("write", chunk.GetObjectValue()); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
