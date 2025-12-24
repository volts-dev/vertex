package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBIndex

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

var singletonIDBKeyRange sync.Once

var idbkeyrangeinterface js.Value

// GetIDBIndexInterface get the JS interface
func GetIDBKeyRangeInterface() js.Value {

	singletonIDBIndex.Do(func() {

		if idbkeyrangeinterface = js.Global().Get("IDBKeyRange"); idbkeyrangeinterface.Error() != nil {
			idbkeyrangeinterface = js.Undefined()
		}
		js.Register(idbkeyrangeinterface, func(v js.Value) (interface{}, error) {
			return IDBDKeyRangeNewFromJSObject(v)
		})
	})
	return idbkeyrangeinterface
}

// IDBKeyRange struct
type IDBKeyRange struct {
	js.Object
}

type IDBKeyRangeFrom interface {
	IDBKeyRange_() IDBKeyRange
}

func (i IDBKeyRange) IDBKeyRange_() IDBKeyRange {
	return i
}

func IDBDKeyRangeNewFromJSObject(obj js.Value) (IDBKeyRange, error) {
	var i IDBKeyRange
	var err error
	if ai := GetIDBKeyRangeInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.SetObjectValue(obj)
			} else {
				err = ErrNotAnIDBKeyRange
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func newKeyRange(method string, values ...interface{}) (IDBKeyRange, error) {
	var i IDBKeyRange
	var err error
	var obj js.Value
	if ii := GetIDBKeyRangeInterface(); !ii.IsUndefined() {

		if obj = ii.New(values...); obj.Error() == nil {
			i.SetObjectValue(obj)
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func Bound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("bound", values...)
}

func LowerBound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("lowerBound", values...)
}

func UpperBound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("upperBound", values...)
}

func Only(value interface{}) (IDBKeyRange, error) {
	return newKeyRange("only", value)
}

func (i IDBKeyRange) Includes(value interface{}) (bool, error) {
	var obj js.Value
	var err error
	var ret bool

	if obj = i.Call("includes", value); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return ret, err
}

func (i IDBKeyRange) lowerOpen() (bool, error) {
	return i.GetAttributeBool("lowerOpen")
}

func (i IDBKeyRange) upperOpen() (bool, error) {
	return i.GetAttributeBool("upperOpen")
}

func (i IDBKeyRange) Lower() (interface{}, error) {
	return i.GetValueByKey("lower"), i.GetObjectValue().Error()
}

func (i IDBKeyRange) Upper() (interface{}, error) {
	return i.GetValueByKey("upper"), i.GetObjectValue().Error()
}
