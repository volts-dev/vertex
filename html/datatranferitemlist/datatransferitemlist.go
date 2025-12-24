package datatranferitemlist

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/datatransferitem"
	"github.com/volts-dev/vertex/html/file"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var datatransferitemlistinterface js.Value

// DataTransferItemList struct
type DataTransferItemList struct {
	js.Object
}

type DataTransferItemListFrom interface {
	DataTransferItemList_() DataTransferItemList
}

func (d DataTransferItemList) DataTransferItemList_() DataTransferItemList {
	return d
}

// GetJSInterface get the JS interface DataTransferItemList
func GetInterface() js.Value {

	singleton.Do(func() {

		if datatransferitemlistinterface = js.Global().Get("DataTransferItemList"); datatransferitemlistinterface.Error() != nil {
			datatransferitemlistinterface = js.Undefined()
		}
		js.Register(datatransferitemlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return datatransferitemlistinterface
}

func NewFromJSObject(obj js.Value) (DataTransferItemList, error) {
	var d DataTransferItemList
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				d.SetObjectValue(obj)

			} else {
				err = ErrNotADataTransferItemList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DataTransferItemList) Length() (int, error) {

	return d.GetAttributeInt("length")

}

// doc said input can be file or string but string not work
func (d DataTransferItemList) Add(f file.File) error {
	var err error
	err = d.Call("add", f.GetObjectValue()).Error()
	return err
}

func (d DataTransferItemList) Remove(index int) error {
	var err error
	err = d.Call("remove", js.ValueOf(index)).Error()
	return err
}

func (d DataTransferItemList) Clear() error {
	var err error
	err = d.Call("clear").Error()
	return err
}

// this func doesn't work but exist in doc
func (d DataTransferItemList) DataTransferItem(index int) (datatransferitem.DataTransferItem, error) {

	var err error
	var obj js.Value
	var dt datatransferitem.DataTransferItem

	if obj = d.GetValueByIndex(index); obj.Error() == nil {
		dt, err = datatransferitem.NewFromJSObject(obj)
	}
	return dt, err

}
