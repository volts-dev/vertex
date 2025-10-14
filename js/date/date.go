package date

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var dateinterface js.Value

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		if dateinterface = js.Global().Get("Date"); dateinterface.Error() != nil {
			dateinterface = js.Undefined()
		}
		js.Register(dateinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return dateinterface
}

// Date struct
type Date struct {
	js.Object
}

type DateFrom interface {
	Date_() Date
}

func (d Date) Date_() Date {
	return d
}

func New(values ...interface{}) (Date, error) {
	var d Date
	var err error
	var arrayJS []interface{}
	var obj js.Value

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}
	if di := GetInterface(); !di.IsUndefined() {

		if obj = di.New(arrayJS...); obj.Error() == nil {
			d.SetObjectValue(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (Date, error) {
	var d Date
	var err error

	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.SetObjectValue(obj)
			} else {
				err = ErrNotADate
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d Date) callString(method string, opts ...interface{}) (string, error) {

	var err error
	var obj js.Value
	var ret string

	var optJSValue []interface{}

	for _, opt := range opts {
		optJSValue = append(optJSValue, js.ValueOf(opt))
	}

	if obj = d.Call(method, optJSValue...); obj.Error() == nil {
		if obj.Type() == js.TypeString {
			return obj.String()
		} else {
			err = object.ErrObjectNotString
		}
	}
	return ret, err
}

func (d Date) GetDate() (int, error) {
	return d.CallInt("getDate")
}

func (d Date) GetDay() (int, error) {
	return d.CallInt("getDay")
}

func (d Date) GetFullYear() (int, error) {
	return d.CallInt("getFullYear")
}

func (d Date) GetHours() (int, error) {
	return d.CallInt("getHours")
}

func (d Date) GetMilliseconds() (int, error) {
	return d.CallInt("getMilliseconds")
}

func (d Date) GetMinutes() (int, error) {
	return d.CallInt("getMinutes")
}

func (d Date) GetSeconds() (int, error) {
	return d.CallInt("getSeconds")
}

func (d Date) GetTime() (int64, error) {
	return d.CallInt64("getTime")
}

func (d Date) GetTimezoneOffset() (int64, error) {
	return d.CallInt64("getTimezoneOffset")
}

func (d Date) GetUTCDate() (int, error) {
	return d.CallInt("getUTCDate")
}

func (d Date) GetUTCDay() (int, error) {
	return d.CallInt("getUTCDay")
}

func (d Date) GetUTCFullYear() (int, error) {
	return d.CallInt("getUTCFullYear")
}

func (d Date) GetUTCHours() (int, error) {
	return d.CallInt("getUTCHours")
}

func (d Date) GetUTCMilliseconds() (int, error) {
	return d.CallInt("getUTCMilliseconds")
}

func (d Date) GetUTCMinutes() (int, error) {
	return d.CallInt("getUTCMinutes")
}

func (d Date) GetUTCMonth() (int, error) {
	return d.CallInt("getUTCMonth")
}

func (d Date) GetUTCSeconds() (int, error) {
	return d.CallInt("getUTCSeconds")
}

func (d Date) SetDate(value int) error {

	return d.Call("setDate", js.ValueOf(value)).Error()
}

func (d Date) SetFullYear(value int) error {

	return d.Call("setFullYear", js.ValueOf(value)).Error()
}

func (d Date) SetHours(value int) error {

	return d.Call("setHours", js.ValueOf(value)).Error()
}

func (d Date) SetMilliseconds(value int) error {

	return d.Call("setMilliseconds", js.ValueOf(value)).Error()
}

func (d Date) SetMinutes(value int) error {

	return d.Call("setMinutes", js.ValueOf(value)).Error()
}

func (d Date) SetSeconds(value int) error {

	return d.Call("setSeconds", js.ValueOf(value)).Error()
}

func (d Date) SetTime(value int64) error {

	return d.Call("setTime", js.ValueOf(value)).Error()
}

func (d Date) SetUTCDate(value int) error {

	return d.Call("setUTCDate", js.ValueOf(value)).Error()
}

func (d Date) SetUTCFullYear(value int) error {

	return d.Call("setUTCFullYear", js.ValueOf(value)).Error()
}

func (d Date) SetUTCHours(value int) error {

	return d.Call("setUTCHours", js.ValueOf(value)).Error()
}

func (d Date) SetUTCMilliseconds(value int) error {

	return d.Call("setUTCMilliseconds", js.ValueOf(value)).Error()
}

func (d Date) SetUTCMinutes(value int) error {

	return d.Call("setUTCMinutes", js.ValueOf(value)).Error()
}

func (d Date) SetUTCMonth(value int) error {

	return d.Call("setUTCMonth", js.ValueOf(value)).Error()
}

func (d Date) SetUTCSeconds(value int) error {
	return d.Call("setUTCSeconds", js.ValueOf(value)).Error()
}

func Parse(value string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if di := GetInterface(); !di.IsUndefined() {
		if obj = di.Call("parse", js.ValueOf(value)); obj.Error() == nil {
			if obj.Type() == js.TypeNumber {
				v, _ := obj.Float()
				ret = int64(v)
			} else {
				err = object.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func Now() (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if di := GetInterface(); !di.IsUndefined() {

		if obj = di.Call("now"); obj.Error() == nil {
			if obj.Type() == js.TypeNumber {
				v, _ := obj.Float()
				ret = int64(v)
			} else {
				err = object.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func (d Date) ToDateString() (string, error) {
	return d.callString("toDateString")
}

func (d Date) ToISOString() (string, error) {
	return d.callString("toISOString")
}

func (d Date) ToJSON() (string, error) {
	return d.callString("toJSON")
}

func (d Date) ToLocaleDateString(opts ...interface{}) (string, error) {
	var arrayopts []interface{}
	if len(opts) > 0 {
		arrayopts = append(arrayopts, js.ValueOf(opts[0]))
	}

	if len(opts) > 1 {
		arrayopts = append(arrayopts, js.ValueOf(opts[1]))
	}
	return d.callString("toLocaleDateString", arrayopts...)
}

func (d Date) ToLocaleString(opts ...interface{}) (string, error) {

	var arrayopts []interface{}
	if len(opts) > 0 {
		arrayopts = append(arrayopts, js.ValueOf(opts[0]))
	}

	if len(opts) > 1 {
		arrayopts = append(arrayopts, js.ValueOf(opts[1]))
	}

	return d.callString("toLocaleString", arrayopts...)
}

func (d Date) ToLocaleTimeString(opts ...interface{}) (string, error) {

	var arrayopts []interface{}
	if len(opts) > 0 {
		arrayopts = append(arrayopts, js.ValueOf(opts[0]))
	}

	if len(opts) > 1 {
		arrayopts = append(arrayopts, js.ValueOf(opts[1]))
	}
	return d.callString("toLocaleTimeString", arrayopts...)
}

func (d Date) ToTimeString() (string, error) {
	return d.callString("toTimeString")
}

func (d Date) ToUTCString() (string, error) {
	return d.callString("toUTCString")
}

func (d Date) ValueOf() (int64, error) {
	return d.CallInt64("valueOf")
}

func UTC(values ...interface{}) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	if di := GetInterface(); !di.IsUndefined() {

		if obj = di.Call("UTC", arrayJS...); obj.Error() == nil {
			if obj.Type() == js.TypeNumber {
				v, _ := obj.Float()
				ret = int64(v)
			} else {
				err = object.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}
