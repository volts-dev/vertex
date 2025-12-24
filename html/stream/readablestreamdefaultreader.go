package stream

import (
	"sync"

	"github.com/volts-dev/vertex/html/jserror"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/typedarray"
	"github.com/volts-dev/vertex/js"
)

var singletonReadableStreamDefault sync.Once

var readablestreamdefaultinterface js.Value

// GetReadableStreamDefaultReaderInterface
func GetReadableStreamDefaultReaderInterface() js.Value {

	singletonReadableStreamDefault.Do(func() {

		if readablestreamdefaultinterface = js.Global().Get("ReadableStreamDefaultReader"); readablestreamdefaultinterface.Error() != nil {
			readablestreamdefaultinterface = js.Undefined()
		}
	})

	return readablestreamdefaultinterface
}

type ReadableStreamDefaultReader struct {
	js.Object
}

type ReadableStreamDefaultReaderFrom interface {
	ReadableStreamDefaultReader_() ReadableStreamDefaultReader
}

func (r ReadableStreamDefaultReader) ReadableStreamDefaultReader_() ReadableStreamDefaultReader {
	return r
}

func NewReadableStreamDefaultReaderFromJSObject(obj js.Value) (ReadableStreamDefaultReader, error) {
	var r ReadableStreamDefaultReader

	if rsi := GetReadableStreamDefaultReaderInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			r.SetObjectValue(obj)
			return r, nil

		}
	}

	return r, ErrNotAReadableStream
}

func (r ReadableStreamDefaultReader) newRead(data []byte, dataHandle func([]byte, int)) *promise.Promise {
	var pp *promise.Promise
	var err error
	var promiseread js.Value

	if promiseread = r.Call("read"); promiseread.Error() == nil {
		var p promise.Promise

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			newpromise, _ := p.Then(func(i interface{}) *promise.Promise {
				var obj js.Value
				if b, ok := i.(js.ObjectFrom); ok {
					obj = b.GetObjectValue()
					var done bool = false
					if js.ValueToBool(obj.Get("done")) == true {
						done = true
					}

					var u8array typedarray.Uint8Array
					var n int
					uint8arrayObject := obj.Get("value")

					if u8array, err = typedarray.NewUint8Array(uint8arrayObject); err == nil {

						if n, err = u8array.CopyBytes(data); err == nil {
							dataHandle(data, n)
						} else {
							rej, _ := promise.Reject(err)
							return &rej
						}

					}

					if done == false {
						return r.newRead(data, dataHandle)
					} else {
						return nil
					}

				}

				return nil
			}, nil)
			pp = &newpromise

		}
	}
	return pp
}
func (r ReadableStreamDefaultReader) AsyncRead(buffersize int, dataHandle func([]byte, int)) (promise.Promise, error) {

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		var data []byte = make([]byte, buffersize)
		var p *promise.Promise

		p = r.newRead(data, dataHandle)

		p.Then(func(i interface{}) *promise.Promise {
			resolvefunc.Invoke(nil)
			return nil
		}, func(e error) {
			if errjs, err := jserror.New(e); err == nil {
				errfunc.Invoke(errjs.GetObjectValue())
			}
		})

		return nil, nil

	})

}

func (r ReadableStreamDefaultReader) Closed() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = r.GetValueByKey("closed"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (r ReadableStreamDefaultReader) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj = r.Call("cancel"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (r ReadableStreamDefaultReader) ReleaseLock() error {

	err := r.Call("releaseLock").Error()
	return err
}

func (w ReadableStreamDefaultReader) Read() (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj = w.Call("read"); obj.Error() == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
