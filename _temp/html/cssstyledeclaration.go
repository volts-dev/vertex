package html

import (
	//"github.com/volts-dev/vertex/html/cssrule"
	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/js"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented CSSStyleDeclaration")
	ErrNotAnCSSStyleDeclaration = errors.New("Object is not a CSSStyleDeclaration")
)

func init() {

	js.RegisterInterface(GetInterface)
}

var cssstyledeclarationinterface js.Value

// CSSStyleDeclaration struct
type CSSStyleDeclaration struct {
	js.Object
}

type CSSStyleDeclarationFrom interface {
	CSSStyleDeclaration_() CSSStyleDeclaration
}

func (c CSSStyleDeclaration) CSSStyleDeclaration_() CSSStyleDeclaration {
	return c
}

func GetInterface() js.Value {
	singleton.Do(func() {
		if cssstyledeclarationinterface = js.Global().Get("CSSStyleDeclaration"); cssstyledeclarationinterface.IsUndefined() {
			cssstyledeclarationinterface = js.Undefined()
		}

		js.Register(cssstyledeclarationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return cssstyledeclarationinterface
}

func NewFromJSObject(obj js.Value) (CSSStyleDeclaration, error) {
	var c CSSStyleDeclaration
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = errors.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				c.SetValue(obj)

			} else {
				err = ErrNotAnCSSStyleDeclaration
			}
		}
	} else {
		err = errors.ErrNotImplemented
	}
	return c, err
}

/*
	func (c CSSStyleDeclaration) ParentRule() (cssrule.CSSRule, error) {

		var obj js.Value
		var cr cssrule.CSSRule
		if obj = c.Get("parentRule"); obj.IsUndefined() {
			err = errors.ErrNotAnObject

		} else {
			cr, err = cssrule.NewFromJSObject(obj)
		}

		return cr, err
	}
*/

func (c CSSStyleDeclaration) SetProperty(propertyName string, opts ...string) error {
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(propertyName))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	c.Call("setProperty", arrayJS...)
	return nil

}

func (c CSSStyleDeclaration) Item(index int) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj = c.Call("item", js.ValueOf(index)); !obj.IsUndefined() {
		return obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) GetPropertyPriority(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj = c.Call("getPropertyPriority", js.ValueOf(property)); !obj.IsUndefined() {
		return obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) GetPropertyValue(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj = c.Call("getPropertyValue", js.ValueOf(property)); !obj.IsUndefined() {
		return obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) RemoveProperty(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj = c.Call("removeProperty", js.ValueOf(property)); !obj.IsUndefined() {
		return obj.String()
	}
	return ret, err
}
