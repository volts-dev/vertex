package clipboarditem

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/blob"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/js/array"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var clipboarditeminterface js.Value

// GetInterface get the JS interface of clipboarditem
func GetInterface() js.Value {

	singleton.Do(func() {

		if clipboarditeminterface = js.Global().Get("ClipboardItem"); clipboarditeminterface.Error() != nil {
			clipboarditeminterface = js.Undefined()
		}
		js.Register(clipboarditeminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		array.GetInterface()
		blob.GetInterface()
		promise.GetInterface()
	})

	return clipboarditeminterface
}

type ClipboardItem struct {
	js.Object
}

type ClipboardItemFrom interface {
	ClipboardItem_() ClipboardItem
}

func (c ClipboardItem) ClipboardItem_() ClipboardItem {
	return c
}

func NewFromJSObject(obj js.Value) (ClipboardItem, error) {
	var c ClipboardItem

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			c.SetObjectValue(obj)
			return c, nil

		}
	}

	return c, ErrNotImplemented
}

func New(data map[string]blob.Blob) (ClipboardItem, error) {

	var c ClipboardItem

	var obj js.Value
	var err error

	var arg map[string]interface{} = make(map[string]interface{})

	for t, b := range data {

		arg[t] = b.GetObjectValue()
		break

	}

	if ci := GetInterface(); !ci.IsUndefined() {

		if obj = ci.New(arg); obj.Error() == nil {
			c.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented

	}
	return c, err
}

func (c ClipboardItem) Types() (array.Array, error) {

	var err error
	var obj interface{}
	var newArr array.Array
	var ok bool

	if obj, err = c.GetAttributeGlobal("types"); err == nil {

		if newArr, ok = obj.(array.Array); !ok {

			err = array.ErrNotAnArray
		}

	}
	return newArr, err
}

// support safari only
func (c ClipboardItem) PresentationStyle() (string, error) {

	return c.GetAttributeString("presentationStyle")
}

func (c ClipboardItem) GetType(mimetype string) (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj = c.Call("getType", js.ValueOf(mimetype)); obj.Error() == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
