package storage

// https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var storageinterface js.Value

// GetInterface get the Storage interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if storageinterface = js.Global().Get("Storage"); storageinterface.Error() != nil {
			storageinterface = js.Undefined()
		}
		js.Register(storageinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return storageinterface
}

type Storage struct {
	js.Object
}

type StorageFrom interface {
	Storage_() Storage
}

func (s Storage) Storage_() Storage {
	return s
}

func NewFromJSObject(obj js.Value) (Storage, error) {
	var s Storage
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				s.SetObjectValue(obj)

			} else {
				err = ErrNotAStorage
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return s, err
}

func (l Storage) SetItem(key, value string) error {
	var err error
	err = l.Call("setItem", js.ValueOf(key), js.ValueOf(value)).Error()
	return err
}

func (l Storage) GetItem(key string) (interface{}, error) {
	var err error
	var itemObject js.Value
	var ret interface{}

	if itemObject = l.Call("getItem", js.ValueOf(key)); itemObject.Error() == nil {
		if !itemObject.IsUndefined() && !itemObject.IsNull() {
			if itemObject.Type() == js.TypeString {
				return itemObject.String()
			}
		}

	}
	return ret, err
}

func (l Storage) RemoveItem(key string) error {
	var err error
	err = l.Call("removeItem", js.ValueOf(key)).Error()
	return err
}

func (l Storage) Clear() error {
	var err error
	err = l.Call("clear").Error()
	return err
}
func (l Storage) Key(index int) (interface{}, error) {
	var err error
	var itemObject js.Value
	var ret interface{}

	if itemObject = l.Call("key", js.ValueOf(index)); itemObject.Error() == nil {
		if !itemObject.IsUndefined() && !itemObject.IsNull() {
			if itemObject.Type() == js.TypeString {
				return itemObject.String()
			}
		}
	}
	return ret, err
}
