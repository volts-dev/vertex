package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {
	js.RegisterInterface(GetClipboardItemInterface)
	js.RegisterInterface(GetClipboardInterface)
}

var ErrNotAClipboard = errors.New("Object is not a Clipboard")

var singletonClipboard sync.Once
var singletonClipboardItem sync.Once
var clipboardinterface js.Value
var clipboarditeminterface js.Value

type (
	ClipboardItem struct {
		js.Object
	}

	ClipboardItemFrom interface {
		ClipboardItem_() ClipboardItem
	}

	// Clipboard struct
	Clipboard struct {
		html.EventTarget
	}

	ClipboardFrom interface {
		Clipboard_() Clipboard
	}
)

func (c Clipboard) Clipboard_() Clipboard {
	return c
}

// GetJSInterface get the JS interface
func GetClipboardInterface() js.Value {

	singletonClipboard.Do(func() {
		if clipboardinterface = js.Global().Get("Clipboard"); clipboardinterface.IsUndefined() {
			clipboardinterface = js.Undefined()
		}

		js.Register(clipboardinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})

		GetClipboardItemInterface()
		GetPromiseInterface()

	})

	return clipboardinterface
}

func ToClipboard(obj js.Value) (Clipboard, error) {
	var c Clipboard

	if ci := GetClipboardInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			c.SetValue(obj)
			return c, nil

		}
	}

	return c, js.ErrNotImplemented
}

func (c Clipboard) Read() (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = c.Call("read"); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}

func (c Clipboard) ReadText() (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = c.Call("readText"); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}

func (c Clipboard) Write(data []ClipboardItem) (Promise, error) {
	var err error
	var obj js.Value
	var p Promise
	var arrayJS []interface{}

	for _, value := range data {
		arrayJS = append(arrayJS, value.Value())
	}

	if obj = c.Call("write", arrayJS); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}

func (c Clipboard) WriteText(data string) (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = c.Call("writeText", js.ValueOf(data)); obj.Error() == nil {

		p, err = ToPromise(obj)
	}

	return p, err
}

// GetInterface get the JS interface of clipboarditem
func GetClipboardItemInterface() js.Value {
	singletonClipboardItem.Do(func() {
		if clipboarditeminterface = js.Global().Get("ClipboardItem"); clipboarditeminterface.Error() != nil {
			clipboarditeminterface = js.Undefined()
		}

		js.Register(clipboarditeminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		js.GetArrayInterface()
		GetBlobInterface()
		GetPromiseInterface()
	})

	return clipboarditeminterface
}

func (c ClipboardItem) ClipboardItem_() ClipboardItem {
	return c
}

func ToClipboardItem(obj js.Value) (ClipboardItem, error) {
	var c ClipboardItem

	if ci := GetClipboardItemInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			//c.BaseObject = c.SetObject(obj)
			c.SetValue(obj)
			return c, nil

		}
	}

	return c, js.ErrNotImplemented
}

func NewClipboardItem(data map[string]Blob) (ClipboardItem, error) {

	var c ClipboardItem

	var obj js.Value
	var err error

	var arg map[string]interface{} = make(map[string]interface{})

	for t, b := range data {

		arg[t] = b.Value()
		break

	}

	if ci := GetClipboardItemInterface(); !ci.IsUndefined() {

		if obj = ci.New(arg); obj.Error() == nil {
			//c.BaseObject = c.SetObject(obj)
			c.SetValue(obj)
		}

	} else {
		err = js.ErrNotImplemented

	}
	return c, err
}

func (c ClipboardItem) Types() (js.Array, error) {

	var err error
	var obj interface{}
	var newArr js.Array
	var ok bool

	if obj, err = c.GetAttributeGlobal("types"); err == nil {

		if newArr, ok = obj.(js.Array); !ok {

			err = js.ErrNotAnArray
		}

	}
	return newArr, err
}

// support safari only
func (c ClipboardItem) PresentationStyle() (string, error) {
	return c.GetAttributeString("presentationStyle")
}

func (c ClipboardItem) GetType(mimetype string) (Promise, error) {
	var err error
	var obj js.Value
	var p Promise

	if obj = c.Call("getType", js.ValueOf(mimetype)); obj.Error() == nil {
		p, err = ToPromise(obj)
	}

	return p, err
}
