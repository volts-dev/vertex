package messageevent

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`mevent=new MessageEvent("test",{data:Object(),origin:"or",lastEventId:"123"})
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object MessageEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("mevent"); test.AssertErr(t, obj.Error()) {
		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object MessageEvent]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Data", "type": "constructnamechecking", "resultattempt": "Object"},
	{"method": "Source", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "Origin", "resultattempt": "or"},
	{"method": "LastEventId", "resultattempt": "123"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("mevent"); test.AssertErr(t, obj.Error()) {

		if mevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, mevent, result)
			}

		}

	}
}
