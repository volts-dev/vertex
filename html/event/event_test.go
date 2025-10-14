package event

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`event=new Event("test",{"bubbles":true, "cancelable":true, "composed":true})
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if e, err := New("test", map[string]interface{}{"bubbles": true, "cancelable": true, "composed": true}); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object Event]", e.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("event"); test.AssertErr(t, obj.Error()) {
		if event, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Event]", event.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Bubbles", "resultattempt": true},
	{"method": "Cancelable", "resultattempt": true},
	{"method": "Composed", "resultattempt": true},
	{"method": "EventPhase", "resultattempt": 0},
	{"method": "Type", "resultattempt": "test"},
	{"method": "IsTrusted", "resultattempt": false},
	{"method": "Target", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "CurrentTarget", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "PreventDefault", "type": "error", "resultattempt": nil},
	{"method": "StopImmediatePropagation", "type": "error", "resultattempt": nil},
	{"method": "StopPropagation", "type": "error", "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("event"); test.AssertErr(t, obj.Error()) {

		if event, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, event, result)
			}

		}

	}
}
