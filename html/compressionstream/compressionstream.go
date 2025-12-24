package compressionstream

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/response"
	"github.com/volts-dev/vertex/html/stream"
)

func init() {

	js.RegisterInterface(GetInterface)

}

var singleton sync.Once

var compressionstreaminterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		if compressionstreaminterface = js.Global().Get("CompressionStream"); compressionstreaminterface.Error() != nil {
			compressionstreaminterface = js.Undefined()
		}

		response.GetInterface()
		js.Register(compressionstreaminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return compressionstreaminterface
}

// CompressionStream struct
type CompressionStream struct {
	js.Object
}

type CompressionStreamFrom interface {
	CompressionStream_() CompressionStream
}

func (c CompressionStream) CompressionStream() CompressionStream {
	return c
}

func New(format string) (CompressionStream, error) {
	var c CompressionStream
	var err error
	var obj js.Value

	if compressionstreami := GetInterface(); !compressionstreami.IsUndefined() {

		if obj = compressionstreami.New(js.ValueOf(format)); obj.Error() == nil {
			c.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return c, err

}

func NewFromJSObject(obj js.Value) (CompressionStream, error) {
	var c CompressionStream
	var err error
	if compressionstreami := GetInterface(); !compressionstreami.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(compressionstreami) {

				c.SetObjectValue(obj)
			} else {
				err = ErrNotACompressionStream
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return c, err
}

func (c CompressionStream) Readable() (stream.ReadableStream, error) {
	var err error
	var obj js.Value

	if obj = c.GetValueByKey("readable"); obj.Error() == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return stream.ReadableStream{}, err
}

func (c CompressionStream) Writable() (stream.WritableStream, error) {
	var err error
	var obj js.Value

	if obj = c.GetValueByKey("writable"); obj.Error() == nil {
		return stream.NewWriteableStreamFromJSObject(obj)

	}
	return stream.WritableStream{}, err
}
