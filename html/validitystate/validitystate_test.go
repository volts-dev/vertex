package validitystate

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`b=document.createElement("button")
	validity=b.validity
	`)

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("validity"); test.AssertErr(t, obj.Error()) {
		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ValidityState]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "BadInput", "resultattempt": false},
	{"method": "CustomError", "resultattempt": false},
	{"method": "PatternMismatch", "resultattempt": false},
	{"method": "RangeOverflow", "resultattempt": false},
	{"method": "RangeUnderflow", "resultattempt": false},
	{"method": "StepMismatch", "resultattempt": false},
	{"method": "TooLong", "resultattempt": false},
	{"method": "TooShort", "resultattempt": false},
	{"method": "TypeMismatch", "resultattempt": false},
	{"method": "Valid", "resultattempt": true},
	{"method": "ValueMissing", "resultattempt": false},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("validity"); test.AssertErr(t, obj.Error()) {

		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, mevent, result)
			}

		}

	}
}
