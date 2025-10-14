package reflect

import (
	"errors"

	"github.com/volts-dev/vertex/js"
)

var setFunc js.Value
var getFunc js.Value
var callFunc js.Value
var invokeFunc js.Value
var newFunc js.Value
var errorInterface js.Value

func SetSyscall() {
	//Set Set and get function
	eval_(`hSet = (obj, set , value) => { try { Reflect.set(obj,set,value) ; return }catch(err){ return err } }`)
	eval_(`hGet = (obj, get ) => { try { return [true,Reflect.get(obj,get)] }catch(err){ return [false,err] } }`)
	eval_(`hCall = (obj,method,args) => { try { func=Reflect.get(obj,method); return [true,Reflect.apply(func,obj,args)] } catch (err) { return [false,err] } }`)
	eval_(`hInvoke = (func,args) => { try { return [true,Reflect.apply(func,undefined,args)] } catch (err) { return [false,err] } }`)
	eval_(`hNew = (func,args) => { try { return [true,Reflect.construct(func,args)] } catch (err) { return [false,err] } }`)
	setFunc = js.Global().Get("hSet")
	getFunc = js.Global().Get("hGet")
	callFunc = js.Global().Get("hCall")
	invokeFunc = js.Global().Get("hInvoke")
	newFunc = js.Global().Get("hNew")

	errorInterface = js.Global().Get("Error")
}

func Set(obj js.Value, name string, val interface{}) error {
	var err error
	ret := setFunc.Invoke(obj, js.ValueOf(name), val)

	if !ret.IsUndefined() {
		msg, _ := ret.Get("message").String()
		err = errors.New(msg)
	}
	return err
}

func Get(obj js.Value, i interface{}) (js.Value, error) {
	var invokvar interface{}

	if s, ok := i.(string); ok {

		invokvar = js.ValueOf(s)
	} else {
		invokvar = i
	}
	ret := getFunc.Invoke(obj, invokvar)

	if _, err := ret.Index(0).Bool(); err != nil {
		return ret.Index(1), err

	} else {
		return ret.Index(1), nil
	}
}

func GetIndex(obj js.Value, index int) (js.Value, error) {

	ret := getFunc.Invoke(obj, js.ValueOf(index))
	if _, err := ret.Index(0).Bool(); err != nil {
		return ret.Index(1), err

	} else {
		return ret.Index(1), nil
	}
}

func New(obj js.Value, args ...interface{}) (js.Value, error) {
	var jsargs []interface{}

	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := newFunc.Invoke(obj, jsargs)

	if _, err := ret.Index(0).Bool(); err != nil {
		return ret.Index(1), err

	} else {
		return ret.Index(1), nil
	}
}

func Call(obj js.Value, name string, args ...interface{}) (js.Value, error) {

	var jsargs []interface{}

	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := callFunc.Invoke(obj, js.ValueOf(name), jsargs)

	if _, err := ret.Index(0).Bool(); err != nil {
		return ret.Index(1), err

	} else {
		return ret.Index(1), nil
	}
}

func Invoke(f js.Value, args ...interface{}) (js.Value, error) {

	var jsargs []interface{}

	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := invokeFunc.Invoke(f, jsargs)

	if _, err := ret.Index(0).Bool(); err != nil {
		return ret.Index(1), err

	} else {
		return ret.Index(1), nil
	}
}

func eval_(str string) {
	js.Global().Call("eval", str)
}

func GetFuncName(inter js.Value) (string, error) {
	var obj js.Value
	var err error
	var name string

	if obj, err = Get(inter, "name"); err == nil {
		if !obj.IsUndefined() {
			return obj.String()
		} else {
			err = js.ErrUnableGetFunctName
		}
	}

	return name, err
}

// String return the string representation of the given Object
func String(object js.Value) (string, error) {
	return object.String()
}

// ToStringWithErr return the ToString representation of the given Object
func ToStringWithErr(obj js.Value) (string, error) {
	if obj.Type() == js.TypeObject {
		if value, err := Call(obj, "toString"); err == nil {
			return value.String()
		} else {
			return "", err
		}

	}

	return "", js.ErrNotAnObject
}

 