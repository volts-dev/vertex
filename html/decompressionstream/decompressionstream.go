package decompressionstream

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/response"
	"github.com/volts-dev/vertex/html/stream"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var decompressionstreaminterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		if decompressionstreaminterface = js.Global().Get("DecompressionStream"); decompressionstreaminterface.Error() != nil {
			decompressionstreaminterface = js.Undefined()
		}

		response.GetInterface()
		js.Register(decompressionstreaminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return decompressionstreaminterface
}

// DecompressionStream struct
type DecompressionStream struct {
	js.Object
}

type DecompressionStreamFrom interface {
	DecompressionStream_() DecompressionStream
}

func (d DecompressionStream) DecompressionStream_() DecompressionStream {
	return d
}

func New(format string) (DecompressionStream, error) {
	var d DecompressionStream
	var err error
	var obj js.Value

	if decompressionstreami := GetInterface(); !decompressionstreami.IsUndefined() {

		if obj = decompressionstreami.New(js.ValueOf(format)); obj.Error() == nil {
			d.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err

}

func NewFromJSObject(obj js.Value) (DecompressionStream, error) {
	var d DecompressionStream
	var err error
	if decompressionstreami := GetInterface(); !decompressionstreami.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(decompressionstreami) {

				d.SetObjectValue(obj)
			} else {
				err = ErrNotADecompressionStream
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DecompressionStream) Readable() (stream.ReadableStream, error) {
	var err error
	var obj js.Value

	if obj = d.GetValueByKey("readable"); obj.Error() == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return stream.ReadableStream{}, err
}

func (d DecompressionStream) Writable() (stream.WritableStream, error) {
	var err error
	var obj js.Value

	if obj = d.GetValueByKey("writable"); obj.Error() == nil {
		return stream.NewWriteableStreamFromJSObject(obj)

	}
	return stream.WritableStream{}, err
}
