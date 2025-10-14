package cssrule

import (
	"sync"

	"github.com/volts-dev/vertex/html/initinterface"
	"github.com/volts-dev/vertex/html/stylesheet"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var cssruleinterface js.Value

// CSSRule struct
type CSSRule struct {
	js.Object
}

type CSSRuleFrom interface {
	CSSRule_() CSSRule
}

func (c CSSRule) CSSRule_() CSSRule {
	return c
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if cssruleinterface = js.Global().Get("CSSRule"); cssruleinterface.Error() != nil {
			cssruleinterface = js.Undefined()
		}
		js.Register(cssruleinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return cssruleinterface
}

func NewFromJSObject(obj js.Value) (CSSRule, error) {
	var c CSSRule
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				c.SetObjectValue(obj)

			} else {
				err = ErrNotAnCSSRule
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return c, err
}

func (c CSSRule) CssText() (string, error) {

	return c.GetAttributeString("cssText")
}

func (c CSSRule) ParentRule() (CSSRule, error) {
	var err error
	var obj js.Value
	var cr CSSRule
	if obj = c.GetValueByKey("parentRule"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = object.ErrNotAnObject

		} else {
			cr, err = NewFromJSObject(obj)
		}
	}
	return cr, err
}

func (c CSSRule) ParentStyleSheet() (stylesheet.StyleSheet, error) {
	var err error
	var obj js.Value
	var s stylesheet.StyleSheet
	if obj = c.GetValueByKey("parentStyleSheet"); obj.Error() == nil {

		if obj.IsUndefined() {
			err = object.ErrNotAnObject

		} else {
			s, err = stylesheet.NewFromJSObject(obj)
		}
	}
	return s, err
}
