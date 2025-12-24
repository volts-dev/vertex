package file

// https://developer.mozilla.org/fr/docs/Web/API/File
import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/blob"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var fileinterface js.Value

// GetInterface get the  JS interface File
func GetInterface() js.Value {

	singleton.Do(func() {

		if fileinterface = js.Global().Get("File"); fileinterface.Error() != nil {
			fileinterface = js.Undefined()
		}
		blob.GetInterface()
		js.Register(fileinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return fileinterface
}

type File struct {
	blob.Blob
}

type FileFrom interface {
	File_() File
}

func (f File) File_() File {
	return f
}

func New(bits interface{}, name string, value ...map[string]interface{}) (File, error) {

	var f File
	var obj js.Value
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(bits), js.ValueOf(name))
	if len(value) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(value[0]))
	}

	if fi := GetInterface(); !fi.IsUndefined() {

		if obj = fi.New(arrayJS...); obj.Error() == nil {
			f.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented

	}
	return f, err
}

func NewFromJSObject(obj js.Value) (File, error) {
	var f File
	var err error
	if fi := GetInterface(); !fi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fi) {
				f.SetObjectValue(obj)

			} else {
				err = ErrNotAFile
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return f, err
}

func (f File) Name() (string, error) {

	return f.GetAttributeString("name")
}

func (f File) Type() (string, error) {
	return f.GetAttributeString("type")
}

func (f File) LastModified() (int64, error) {
	return f.GetAttributeInt64("lastModified")
}

func (f File) LastModifiedDate() (js.Date, error) {
	var obj js.Value
	var d js.Date
	var err error
	if obj = f.GetValueByKey("lastModifiedDate"); obj.Error() == nil {
		d, err = js.NewDateFromJSObject(obj)
	}
	return d, err
}
