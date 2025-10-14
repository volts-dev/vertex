package app

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetDragEventInterface)
}

var ErrNotAnDragEvent = errors.New("Object is not an DragEvent")

var drageventinterface js.Value

// DragEvent DragEvent struct
type DragEvent struct {
	MouseEvent
}

type DragEventFrom interface {
	DragEvent_() DragEvent
}

func (d DragEvent) DragEvent_() DragEvent {
	return d
}

// GetDragEventInterface get teh JS interface of event
func GetDragEventInterface() js.Value {

	singleton.Do(func() {

		if drageventinterface = js.Global().Get("DragEvent"); drageventinterface.Error() != nil {
			drageventinterface = js.Undefined()
		}
		js.Register(drageventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return drageventinterface
}

func NewDragEvent(typeevent string, init ...map[string]interface{}) (DragEvent, error) {

	var d DragEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetDragEventInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj = pei.New(arrayJS...); obj.Error() == nil {
			//d.BaseObject = d.SetObject(obj)
			d.SetValue(obj)
		}

	} else {
		err = js.ErrNotImplemented
	}
	return d, err
}

func ToDragEvent(obj js.Value) (DragEvent, error) {
	var e DragEvent
	var err error
	if di := GetDragEventInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				//	e.BaseObject = e.SetObject(obj)
				e.SetValue(obj)
			} else {
				err = ErrNotAnDragEvent
			}
		}
	} else {
		err = js.ErrNotImplemented
	}
	return e, err
}

func (d DragEvent) DataTransfer() (DataTransfer, error) {

	var err error
	var obj js.Value

	if obj = d.Get("dataTransfer"); obj.Error() == nil {

		return ToDataTransfer(obj)
	}
	return DataTransfer{}, err

}
