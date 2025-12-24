package blob

// Full implemented
// https://developer.mozilla.org/fr/docs/Web/API/Blob

import (
	"sync"

	"github.com/volts-dev/vertex/html/arraybuffer"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/stream"
	readablestream "github.com/volts-dev/vertex/html/stream"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var blobinterface js.Value

// GetInterface get the JS interface Blob
func GetInterface() js.Value {

	singleton.Do(func() {

		if blobinterface = js.Global().Get("Blob"); blobinterface.Error() != nil {
			blobinterface = js.Undefined()
		}
		//autodiscover
		arraybuffer.GetInterface()
		js.Register(blobinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return blobinterface
}

type Blob struct {
	js.Object
}

type BlobFrom interface {
	Blob_() Blob
}

func (b Blob) Blob_() Blob {
	return b
}

func New(values ...interface{}) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	if bi := GetInterface(); !bi.IsUndefined() {

		if obj = bi.New(arrayJS); obj.Error() == nil {
			b.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewWithObject(o js.Value) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {

		if obj = bi.New(o); obj.Error() == nil {
			b.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewWithArrayBuffer(a arraybuffer.ArrayBuffer) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {

		if obj = bi.New([]interface{}{a.GetObjectValue()}); obj.Error() == nil {
			b.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewFromJSObject(obj js.Value) (Blob, error) {
	var b Blob
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				b.SetObjectValue(obj)

			} else {
				err = ErrNotABlob
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return b, err
}

func (b Blob) IsClosed() (bool, error) {
	var err error
	var obj js.Value
	var ret bool

	if obj = b.GetValueByKey("isClosed"); obj.Error() == nil {
		if !obj.IsUndefined() {
			return obj.Bool()
		} else {
			err = js.ErrNotImplementedFunc
		}

	}
	return ret, err
}

func (b Blob) Size() (int64, error) {

	return b.GetAttributeInt64("size")
}
func (b Blob) Type() (string, error) {
	var err error
	var obj js.Value

	if obj = b.GetValueByKey("type"); obj.Error() == nil {

		return obj.String()
	}
	return "", err
}

func (b Blob) Close() error {
	err := b.Call("close").Error()
	return err
}

func (b Blob) Slice(begin, end int64) (Blob, error) {
	var blob js.Value
	var err error
	if blob = b.Call("slice", js.ValueOf(begin), js.ValueOf(end)); blob.Error() == nil {
		var newblob Blob
		//object := newblob.SetObjectValue(blob)
		//newblob.BaseObject = object
		newblob.SetObjectValue(blob)
		return newblob, nil
	}
	return Blob{}, err
}

func (b Blob) Stream() (stream.ReadableStream, error) {

	var err error
	var obj js.Value

	if obj = b.Call("stream"); obj.Error() == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return readablestream.ReadableStream{}, err
}

func (b Blob) ArrayBuffer() (promise.Promise, error) {

	var err error
	var promisebuffer js.Value
	var p promise.Promise

	if promisebuffer = b.Call("arrayBuffer"); promisebuffer.Error() == nil {

		p, err = promise.NewFromJSObject(promisebuffer)

	}

	return p, err
}

func (b Blob) Text() (promise.Promise, error) {
	var err error
	var promisetext js.Value
	var p promise.Promise

	if promisetext = b.Call("text"); promisetext.Error() == nil {
		p, err = promise.NewFromJSObject(promisetext)
	}

	return p, err
}
