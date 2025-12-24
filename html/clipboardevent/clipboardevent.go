package clipboardevent

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/datatransfer"
	"github.com/volts-dev/vertex/html/event"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var clipboardeventinterface js.Value

// ClipboardEvent ClipboardEvent struct
type ClipboardEvent struct {
	event.Event
}

type ClipboardEventFrom interface {
	ClipboardEvent_() ClipboardEvent
}

func (c ClipboardEvent) ClipboardEvent_() ClipboardEvent {
	return c
}

// GetInterface get the JS interface of ClipboardEvent
func GetInterface() js.Value {

	singleton.Do(func() {

		if clipboardeventinterface = js.Global().Get("ClipboardEvent"); clipboardeventinterface.Error() != nil {
			clipboardeventinterface = js.Undefined()
		}

		js.Register(clipboardeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})
		datatransfer.GetInterface()
	})

	return clipboardeventinterface
}

func NewFromJSObject(obj js.Value) (ClipboardEvent, error) {
	var c ClipboardEvent
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				c.SetObjectValue(obj)

			} else {
				err = ErrNotAnCustomEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return c, err
}

// New Create a ClipboardEvent
func New(data datatransfer.DataTransfer) (ClipboardEvent, error) {
	var event ClipboardEvent
	var obj js.Value
	var err error
	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj = eventi.New(data.GetObjectValue()); obj.Error() == nil {
			event.SetObjectValue(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return event, err
}

func (c ClipboardEvent) ClipboardData() (datatransfer.DataTransfer, error) {
	var obj interface{}
	var err error
	var d datatransfer.DataTransfer
	var ok bool
	if obj, err = c.GetAttributeGlobal("clipboardData"); err == nil {
		if d, ok = obj.(datatransfer.DataTransfer); !ok {

			err = datatransfer.ErrNotADataTransfer

		}

	}
	return d, err
}
