package dragevent

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`devent=new DragEvent("test",{dataTransfer:new DataTransfer()})
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object DragEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("devent"); test.AssertErr(t, obj.Error()) {
		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DragEvent]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "DataTransfer", "type": "constructnamechecking", "resultattempt": "DataTransfer"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("devent"); test.AssertErr(t, obj.Error()) {

		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, mevent, result)
			}

		}

	}
}
