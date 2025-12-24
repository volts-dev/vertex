package customevent

// https://developer.mozilla.org/fr/docs/Web/API/CustomEvent
import (
	"sync"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var customeventinterface js.Value

// CustomEvent CustomEvent struct
type CustomEvent struct {
	event.Event
}

type CustomEventFrom interface {
	CustomEvent_() CustomEvent
}

func (c CustomEvent) CustomEvent_() CustomEvent {
	return c
}

// GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		if customeventinterface = js.Global().Get("CustomEvent"); customeventinterface.Error() != nil {
			customeventinterface = js.Undefined()
		}

		js.Register(customeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return customeventinterface
}

// New Create a CustomEvent
func New(message, detail interface{}) (CustomEvent, error) {
	var event CustomEvent
	var obj js.Value
	var err error
	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj = eventi.New(js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": js.ValueOf(detail)})); obj.Error() == nil {
			event.SetObjectValue(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return event, err
}

func NewFromJSObject(obj js.Value) (CustomEvent, error) {
	var c CustomEvent
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

func (c CustomEvent) Detail() (interface{}, error) {
	var obj js.Value
	var err error
	var i interface{}

	if obj = c.GetValueByKey("detail"); obj.Error() == nil {
		i, err = js.GoValue(obj)
	}
	return i, err
}
