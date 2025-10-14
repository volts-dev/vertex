package app

import (
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/js"
	"github.com/volts-dev/vertex/core/js/reflect"
)

var (
	ErrNotAPromise        = errors.New("Object is not a Promise")
	ErrResultPromiseError = errors.New("Result error promise")
)

func init() {

	js.RegisterInterface(GetPromiseInterface)
}

var promiseinterface js.Value

// GetInterface get the JS interface
func GetPromiseInterface() js.Value {

	singleton.Do(func() {

		if promiseinterface = js.Global().Get("Promise"); promiseinterface.IsUndefined() {
			promiseinterface = js.Undefined()
		}
		js.Register(promiseinterface, func(v js.Value) (interface{}, error) {
			return ToPromise(v)
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

func NewPromise(handler func(resolvefunc, errfunc js.Value) (interface{}, error)) (Promise, error) {

	var p Promise
	var err error
	var obj js.Value
	if pi := GetPromiseInterface(); !pi.IsUndefined() {
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

		if obj, err = reflect.New(pi, fh); err == nil {
			p.SetValue(obj)
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

	p, err = NewPromise(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		timeout.Invoke(resolvefunc, js.ValueOf(ms))

		return nil, nil
	})

	return p, err
}

func ToPromise(obj js.Value) (Promise, error) {
	var p Promise
	var err error
	if pi := GetPromiseInterface(); !pi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(pi) {
				p.SetValue(obj)

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
					return retp.Object
				}
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var errRejected error
		var exception DomException
		if exception, errRejected = ToDomException(args[0]); errRejected == nil {
			message, _ := exception.Message()
			errRejected = errors.New(message)
		} else {

			var strerr string
			b, _ := js.ToObject(args[0])
			if target := b.Get("target"); target.Error() == nil {
				t, _ := js.ToObject(target)
				if targeterror := t.Get("error"); targeterror.Error() == nil {
					if exception, errRejected = ToDomException(targeterror); errRejected == nil {
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
		newp, err = ToPromise(newpromiseobj)
	}
	return newp, err
}

func iterablePromises(method string, values ...interface{}) (Promise, error) {
	var err error
	var pr Promise
	var promiseobj js.Value
	var arr js.Array

	var arrayJS []interface{}
	if pi := GetPromiseInterface(); !pi.IsUndefined() {
		for _, value := range values {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}
		if arr, err = js.NewArray(arrayJS...); err == nil {

			if promiseobj, err = reflect.Call(pi, method, arr.Value()); err == nil {
				pr, err = ToPromise(promiseobj)
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
		var err error
		var exception DomException
		if exception, err = ToDomException(args[0]); err == nil {
			message, _ := exception.Message()
			err = errors.New(message)
		} else {
			v, err := args[0].String()
			if err == nil {
				err = errors.New(v)
			}
			err = errors.New(v)
		}

		if reject != nil {
			reject(err)
		}
		return nil
	})
	var newpromiseobj js.Value
	if newpromiseobj = p.Call("catch", rejectedFunc); newpromiseobj.IsNull() {
		newp, err = ToPromise(newpromiseobj)
	}

	return newp, err
}

func (p Promise) Finally(f func()) error {
	var err error
	finallyFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	p.Call("finally", finallyFunc)
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
	var jserr js.Error

	var err error
	if pi := GetPromiseInterface(); !pi.IsUndefined() {

		if jserr, err = js.NewError(reason.Error()); err == nil {
			if obj, err = reflect.Call(pi, "reject", jserr.Value()); err == nil {

				p, err = ToPromise(obj)
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
	if pi := GetPromiseInterface(); !pi.IsUndefined() {
		if obj, err = reflect.Call(pi, "resolve", js.ValueOf(result)); err == nil {
			p, err = ToPromise(obj)
		}
	} else {
		err = js.ErrNotImplemented
	}
	return p, err
}
