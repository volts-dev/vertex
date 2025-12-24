package messageevent

// https://developer.mozilla.org/fr/docs/Web/API/MessageEvent

import (
	"sync"

	"github.com/volts-dev/vertex/html/arraybuffer"
	"github.com/volts-dev/vertex/html/blob"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var messageeventinterface js.Value

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if messageeventinterface = js.Global().Get("MessageEvent"); messageeventinterface.Error() != nil {
			messageeventinterface = js.Undefined()
		}
		//instance object for autodiscovery
		arraybuffer.GetInterface()
		blob.GetInterface()
		js.Register(messageeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return messageeventinterface
}

type MessageEvent struct {
	event.Event
}

type MessageEventFrom interface {
	MessageEvent_() MessageEvent
}

func (m MessageEvent) MessageEvent_() MessageEvent {
	return m
}

func New(typeevent string, init ...map[string]interface{}) (MessageEvent, error) {

	var m MessageEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj = pei.New(arrayJS...); obj.Error() == nil {
			m.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return m, err
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {

	var message MessageEvent
	var err error
	if mi := GetInterface(); !mi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(mi) {
				message.SetObjectValue(obj)
			} else {
				err = ErrNotAMessageEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return message, err

}

func (m MessageEvent) Data() (interface{}, error) {

	var jsObject js.Value
	var globalObj interface{}
	var err error
	if jsObject = m.GetValueByKey("data"); jsObject.Error() == nil {
		globalObj, err = js.GoValue(jsObject)
	}

	return globalObj, err
}

func (m MessageEvent) Source() (interface{}, error) {
	var obj js.Value
	var err error
	var i interface{}

	if obj = m.GetValueByKey("source"); obj.Error() == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			i, err = js.Discover(obj)
		}

	}

	return i, err
}
func (m MessageEvent) Origin() (string, error) {
	var err error
	var originObject js.Value

	if originObject = m.GetValueByKey("origin"); originObject.Error() == nil {
		return originObject.String()
	}
	return "", err
}

func (m MessageEvent) LastEventId() (string, error) {
	var err error
	var originObject js.Value

	if originObject = m.GetValueByKey("lastEventId"); originObject.Error() == nil {
		return originObject.String()
	}
	return "", err
}
