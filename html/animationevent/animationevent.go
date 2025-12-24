package animationevent

//https://developer.mozilla.org/fr/docs/Web/API/AnimationEvent

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var animationeventinterface js.Value

// AnimationEvent AnimationEvent struct
type AnimationEvent struct {
	event.Event
}

type AnimationEventFrom interface {
	AnimationEvent() AnimationEvent
}

func (a AnimationEvent) AnimationEvent_() AnimationEvent {
	return a
}

// GetInterface get the JS interface animationEvent
func GetInterface() js.Value {

	singleton.Do(func() {

		if animationeventinterface = js.Global().Get("AnimationEvent"); animationeventinterface.Error() != nil {
			animationeventinterface = js.Undefined()

		}

		js.Register(animationeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return animationeventinterface
}

func NewFromJSObject(obj js.Value) (AnimationEvent, error) {
	var a AnimationEvent
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {
			if obj.InstanceOf(di) {
				a.SetObjectValue(obj)

			} else {
				err = ErrNotAnAnimationEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func (a AnimationEvent) AnimationName() (string, error) {
	return a.GetAttributeString("animationName")
}

func (a AnimationEvent) ElapsedTime() (float64, error) {
	return a.GetAttributeDouble("elapsedTime")
}

func (a AnimationEvent) PseudoElement() (string, error) {
	return a.GetAttributeString("pseudoElement")
}
