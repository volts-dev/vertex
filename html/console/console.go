package console

import (
	"sync"

	"github.com/volts-dev/vertex/js"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var consoleinterface js.Value

// Console Console struct
type Console struct {
	js.Object
}

type ConsoleFrom interface {
	Console_() Console
}

func (c Console) Console_() Console {
	return c
}

// GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		if consoleinterface = js.Global().Get("console"); consoleinterface.Error() != nil {
			consoleinterface = js.Undefined()
		}

		js.Register(consoleinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return consoleinterface
}

func New() (Console, error) {

	var c Console
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		c.SetObjectValue(di)

	} else {

		err = ErrNotImplemented
	}

	return c, err
}
func NewFromJSObject(obj js.Value) (Console, error) {
	var c Console
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				c.SetObjectValue(obj)

			} else {
				err = ErrNotAConsole
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return c, err
}

func (c Console) Assert(assertion bool, opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	arrayJS = append(arrayJS, js.ValueOf(assertion))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("assert", arrayJS...).Error()
	return err
}

func (c Console) Clear() error {
	var err error
	err = c.Call("clear").Error()
	return err
}

func (c Console) Count(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	err = c.Call("count", arrayJS...).Error()
	return err
}

func (c Console) CountReset(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	err = c.Call("countReset", arrayJS...).Error()
	return err
}

func (c Console) Debug(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("debug", arrayJS...).Error()
	return err
}

func (c Console) Dir(obj js.Object) error {

	var err error
	err = c.Call("dir", obj.GetObjectValue()).Error()
	return err
}

func (c Console) DirXml(obj js.Object) error {

	var err error
	err = c.Call("dirxml", obj.GetObjectValue()).Error()
	return err
}

func (c Console) Error(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("error", arrayJS...).Error()
	return err
}

func (c Console) Exception(opts ...interface{}) error {

	return c.Error(opts...)
}

func (c Console) Group(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	err = c.Call("group", arrayJS...).Error()
	return err
}

func (c Console) GroupCollapsed(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	err = c.Call("groupCollapsed", arrayJS...).Error()
	return err
}

func (c Console) GroupEnd() error {

	var err error
	err = c.Call("groupEnd").Error()
	return err
}

func (c Console) Info(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("info", arrayJS...).Error()
	return err
}

func (c Console) Log(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("log", arrayJS...).Error()
	return err
}

func (c Console) Time(label string) error {

	var err error
	err = c.Call("time", js.ValueOf(label)).Error()
	return err
}

func (c Console) TimeEnd(label string) error {

	var err error
	err = c.Call("timeEnd", js.ValueOf(label)).Error()
	return err
}

func (c Console) TimeLog(label string) error {

	var err error
	err = c.Call("timeLog", js.ValueOf(label)).Error()
	return err
}

func (c Console) Trace() error {

	var err error
	err = c.Call("trace").Error()
	return err
}

func (c Console) Warn(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	err = c.Call("warn", arrayJS...).Error()
	return err
}
