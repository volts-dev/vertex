package stylesheet

import (
	"sync"

	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var stylesheetinterface js.Value

// StyleSheet struct
type StyleSheet struct {
	js.Object
}

type StyleSheetFrom interface {
	StyleSheet_() StyleSheet
}

func (s StyleSheet) StyleSheet_() StyleSheet {
	return s
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if stylesheetinterface = js.Global().Get("StyleSheet"); stylesheetinterface.Error() != nil {
			stylesheetinterface = js.Undefined()
		}
		js.Register(stylesheetinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return stylesheetinterface
}

func NewFromJSObject(obj js.Value) (StyleSheet, error) {
	var s StyleSheet
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				s.SetObjectValue(obj)

			} else {
				err = ErrNotAnStyleSheet
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return s, err
}

func (s StyleSheet) Disabled() (bool, error) {
	return s.GetAttributeBool("disabled")
}

func (s StyleSheet) SetDisabled(value bool) error {
	return s.SetAttributeBool("disabled", value)
}

func (s StyleSheet) Href() (string, error) {
	return s.GetAttributeString("href")
}

func (s StyleSheet) OwnerNode() (*node.Node, error) {
	var err error
	var obj js.Value
	var n *node.Node
	if obj = s.GetValueByKey("ownerNode"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			n, err = node.NewFromJSObject(obj)
		}
	}
	return n, err
}

func (s StyleSheet) ParentStyleSheet() (StyleSheet, error) {
	var err error
	var obj js.Value
	var ps StyleSheet
	if obj = s.GetValueByKey("parentStyleSheet"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = js.ErrNotAnObject

		} else {
			ps, err = NewFromJSObject(obj)
		}
	}
	return ps, err
}

/*
func (s StyleSheet) Media() {
//TODO IMPLEMENT
}*/

func (s StyleSheet) Title() (string, error) {
	return s.GetAttributeString("title")
}

func (s StyleSheet) Type() (string, error) {
	return s.GetAttributeString("type")
}
