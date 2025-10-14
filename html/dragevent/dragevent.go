package dragevent

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/datatransfer"
	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/mouseevent"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var drageventinterface js.Value

// DragEvent DragEvent struct
type DragEvent struct {
	mouseevent.MouseEvent
}

type DragEventFrom interface {
	DragEvent_() DragEvent
}

func (d DragEvent) DragEvent_() DragEvent {
	return d
}

// GetInterface get teh JS interface of event
func GetInterface() js.Value {

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

func New(typeevent string, init ...map[string]interface{}) (DragEvent, error) {

	var d DragEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj = pei.New(arrayJS...); obj.Error() == nil {
			d.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func NewFromJSObject(obj js.Value) (DragEvent, error) {
	var e DragEvent
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				e.SetObjectValue(obj)

			} else {
				err = ErrNotAnDragEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func (d DragEvent) DataTransfer() (datatransfer.DataTransfer, error) {

	var err error
	var obj js.Value

	if obj = d.GetValueByKey("dataTransfer"); obj.Error() == nil {

		return datatransfer.NewFromJSObject(obj)
	}
	return datatransfer.DataTransfer{}, err

}
