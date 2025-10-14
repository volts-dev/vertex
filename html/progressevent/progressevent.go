package progressevent

// https://developer.mozilla.org/en-US/docs/Web/API/ProgressEvent

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var progresseeventinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if progresseeventinterface = js.Global().Get("ProgressEvent"); progresseeventinterface.Error() != nil {
			progresseeventinterface = js.Undefined()
		}
		js.Register(progresseeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return progresseeventinterface
}

type ProgressEvent struct {
	event.Event
}

type ProgressEventFrom interface {
	ProgressEvent_() ProgressEvent
}

func (p ProgressEvent) ProgressEvent_() ProgressEvent {
	return p
}

func New(typeevent string, opts ...map[string]interface{}) (ProgressEvent, error) {

	var p ProgressEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(opts) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(opts[0]))
		}
		if obj = pei.New(arrayJS...); obj.Error() == nil {
			p.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return p, err
}

func NewFromJSObject(obj js.Value) (ProgressEvent, error) {
	var p ProgressEvent
	var err error
	if pei := GetInterface(); !pei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(pei) {
				p.SetObjectValue(obj)

			} else {
				err = ErrNotAnProgressEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func (p ProgressEvent) LengthComputable() (bool, error) {
	return p.GetAttributeBool("lengthComputable")
}

func (p ProgressEvent) Loaded() (int, error) {
	return p.GetAttributeInt("loaded")
}

func (p ProgressEvent) Total() (int, error) {
	return p.GetAttributeInt("total")
}
