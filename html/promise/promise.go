package promise

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/html/domexception"
	"github.com/volts-dev/vertex/html/jserror"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var promiseinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		if promiseinterface = js.Global().Get("Promise"); promiseinterface.Error() != nil {
			promiseinterface = js.Undefined()
		}

		js.Register(promiseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return promiseinterface
}

// Promise struct
type Promise struct {
	js.Object
}

type PromiseFrom interface {
	Promise_() Promise
}

func (p Promise) Promise_() Promise {
	return p
}

func New(handler func(resolvefunc, errfunc js.Value) (interface{}, error)) (Promise, error) {

	var p Promise
	var err error
	var obj js.Value
	if pi := GetInterface(); !pi.IsUndefined() {
		fh := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if result, err := handler(args[0], args[1]); err == nil {
				if result != nil {
					args[0].Invoke(result)
				}

			} else {
				args[1].Invoke(err.Error())
			}

			return nil
		})

		if obj = pi.New(fh); obj.Error() == nil {
			p.SetObjectValue(obj)
		}

	} else {
		err = js.ErrNotImplemented
	}

	return p, err
}

func SetTimeout(ms int) (Promise, error) {

	var p Promise
	var err error

	timeout := js.Global().Get("window").Get("setTimeout")

	p, err = New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		timeout.Invoke(resolvefunc, js.ValueOf(ms))

		return nil, nil
	})

	return p, err
}

func NewFromJSObject(obj js.Value) (Promise, error) {
	var p Promise
	var err error
	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(pi) {
				p.SetObjectValue(obj)

			} else {
				err = ErrNotAPromise
			}
		}
	} else {
		err = js.ErrNotImplemented
	}
	return p, err
}

func (p Promise) Then(resolve func(interface{}) *Promise, reject func(error)) (Promise, error) {

	var err error
	var obj interface{}
	var newp Promise
	resolveFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) > 0 {
			obj, err = js.GoValue(args[0])
			if resolve != nil {
				if retp := resolve(obj); retp != nil {
					return retp.GetObjectValue()
				}
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var errRejected error
		var exception domexception.DomException
		if exception, errRejected = domexception.NewFromJSObject(args[0]); errRejected == nil {
			message, _ := exception.Message()
			errRejected = errors.New(message)
		} else {

			var strerr string
			b, _ := js.ToObject(args[0])
			if target := b.GetValueByKey("target"); target.Error() == nil {
				t, _ := js.ToObject(target)
				if targeterror := t.GetValueByKey("error"); targeterror.Error() == nil {
					if exception, errRejected = domexception.NewFromJSObject(targeterror); errRejected == nil {
						message, _ := exception.Message()
						errRejected = errors.New(message)
					}

					if reject != nil {
						reject(errRejected)
					}
					return nil

				}
			}

			if strerr, errRejected = reflect.ToStringWithErr(args[0]); errRejected == nil {
				errRejected = errors.New(strerr)
			}

		}

		if reject != nil {
			reject(errRejected)
		}

		return nil
	})
	var newpromiseobj js.Value
	if newpromiseobj = p.Call("then", resolveFunc, rejectedFunc); newpromiseobj.Error() == nil {
		newp, err = NewFromJSObject(newpromiseobj)
	}
	return newp, err
}

func iterablePromises(method string, values ...interface{}) (Promise, error) {
	var err error
	var pr Promise
	var promiseobj js.Value
	var arr js.Array

	var arrayJS []interface{}
	if pi := GetInterface(); !pi.IsUndefined() {
		for _, value := range values {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}
		if arr, err = js.NewArray(arrayJS...); err == nil {

			if promiseobj = pi.Call(method, arr.GetObjectValue()); promiseobj.Error() == nil {
				pr, err = NewFromJSObject(promiseobj)
			}
		}
	} else {
		err = js.ErrNotImplemented
	}

	return pr, err
}

func All(values ...interface{}) (Promise, error) {
	return iterablePromises("all", values...)

}
func AllSettled(values ...interface{}) (Promise, error) {
	return iterablePromises("allSettled", values...)
}

func Any(values ...interface{}) (Promise, error) {
	return iterablePromises("any", values...)
}

func Race(values ...interface{}) (Promise, error) {
	return iterablePromises("race", values...)
}

func (p Promise) Catch(reject func(error)) (Promise, error) {
	var err error
	var newp Promise
	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var exception domexception.DomException
		if exception, err = domexception.NewFromJSObject(args[0]); err == nil {
			message, _ := exception.Message()
			err = errors.New(message)
		} else {
			err = errors.New(js.ValueToString(args[0]))
		}

		if reject != nil {
			reject(err)
		}
		return nil
	})
	var newpromiseobj js.Value
	if newpromiseobj = p.Call("catch", rejectedFunc); newpromiseobj.Error() == nil {
		newp, err = NewFromJSObject(newpromiseobj)
	}

	return newp, err
}

func (p Promise) Finally(f func()) error {
	var err error
	finallyFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	err = p.Call("finally", finallyFunc).Error()
	return err
}

// avoid used it can deadlocks
func (p Promise) Await() (interface{}, error) {
	var obj interface{}
	var err error
	var ok bool

	ch := make(chan interface{})

	_, err = p.Then(func(i interface{}) *Promise {

		ch <- i
		return nil

	}, func(e error) {
		ch <- e
	})

	returnvalue := <-ch

	if err, ok = returnvalue.(error); !ok {

		obj = returnvalue
	}

	return obj, err
}

func Reject(reason error) (Promise, error) {
	var p Promise
	var obj js.Value
	var jserr jserror.JSError

	var err error
	if pi := GetInterface(); !pi.IsUndefined() {

		if jserr, err = jserror.New(reason.Error()); err == nil {
			if obj = pi.Call("reject", jserr.GetObjectValue()); obj.Error() == nil {

				p, err = NewFromJSObject(obj)
			}
		}

	} else {
		err = js.ErrNotImplemented
	}

	return p, err
}

func Resolve(result interface{}) (Promise, error) {
	var p Promise
	var obj js.Value
	var err error
	if pi := GetInterface(); !pi.IsUndefined() {
		if obj = pi.Call("resolve", js.ValueOf(result)); obj.Error() == nil {
			p, err = NewFromJSObject(obj)
		}
	} else {
		err = js.ErrNotImplemented
	}
	return p, err
}
