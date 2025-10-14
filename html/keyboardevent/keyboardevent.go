package keyboardevent

// https://developer.mozilla.org/fr/docs/Web/API/CustomEvent
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

var customeventinterface js.Value

// KeyboardEvent KeyboardEvent struct
type KeyboardEvent struct {
	event.Event
}

type KeyboardEventFrom interface {
	KeyboardEvent_() KeyboardEvent
}

func (k KeyboardEvent) KeyboardEvent_() KeyboardEvent {
	return k
}

// GetInterface get the KeyboardEvent interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if customeventinterface = js.Global().Get("KeyboardEvent"); customeventinterface.Error() != nil {
			customeventinterface = js.Undefined()
		}

		js.Register(customeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return customeventinterface
}

// New Create a KeyboardEvent
func New(typeK string, opt ...map[string]interface{}) (KeyboardEvent, error) {
	var event KeyboardEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if eventi := GetInterface(); !eventi.IsUndefined() {

		arrayJS = append(arrayJS, js.ValueOf(typeK))
		if len(opt) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(opt[0]))
		}

		if obj = eventi.New(arrayJS...); obj.Error() == nil {
			event.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return event, err
}

func NewFromJSObject(obj js.Value) (KeyboardEvent, error) {
	var k KeyboardEvent
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {

			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				k.SetObjectValue(obj)

			} else {
				err = ErrNotAKeyboardEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return k, err
}

func (k KeyboardEvent) AltKey() (bool, error) {

	return k.GetAttributeBool("altKey")
}

func (k KeyboardEvent) CtrlKey() (bool, error) {

	return k.GetAttributeBool("ctrlKey")
}
func (k KeyboardEvent) Key() (string, error) {

	return k.GetAttributeString("key")
}

func (k KeyboardEvent) Code() (string, error) {

	return k.GetAttributeString("code")
}

func (k KeyboardEvent) IsComposing() (bool, error) {

	return k.GetAttributeBool("isComposing")
}

func (k KeyboardEvent) Location() (int64, error) {

	return k.GetAttributeInt64("location")
}

func (k KeyboardEvent) MetaKey() (bool, error) {

	return k.GetAttributeBool("metaKey")
}

func (k KeyboardEvent) Repeat() (bool, error) {

	return k.GetAttributeBool("repeat")
}

func (k KeyboardEvent) ShiftKey() (bool, error) {

	return k.GetAttributeBool("shiftKey")
}
