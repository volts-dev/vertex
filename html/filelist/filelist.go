package filelist

// https://developer.mozilla.org/fr/docs/Web/API/FileList

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/file"
	"github.com/volts-dev/vertex/html/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var filelistinterface js.Value

// FileList struct
type FileList struct {
	js.Object
}

type FileListFrom interface {
	FileList_() FileList
}

func (f FileList) FileList_() FileList {
	return f
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		if filelistinterface = js.Global().Get("FileList"); filelistinterface.Error() != nil {
			filelistinterface = js.Undefined()
		}
		js.Register(filelistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return filelistinterface
}

func NewFromJSObject(obj js.Value) (FileList, error) {
	var f FileList
	var err error
	if fli := GetInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fli) {
				f.SetObjectValue(obj)

			} else {
				err = ErrNotAnFileList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return f, err
}

func (f FileList) Item(index int) (file.File, error) {

	return file.NewFromJSObject(f.GetObjectValue().Index(index))

}
func (f FileList) Length() (int, error) {

	return f.GetAttributeInt("length")

}
