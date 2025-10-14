package progressevent

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`pevent=new ProgressEvent("test",{lengthComputable:true,loaded:12345,total:1234567})
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object ProgressEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("pevent"); test.AssertErr(t, obj.Error()) {
		if k, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object ProgressEvent]", k.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "LengthComputable", "resultattempt": true},
	{"method": "Loaded", "resultattempt": 12345},
	{"method": "Total", "resultattempt": 1234567},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("pevent"); test.AssertErr(t, obj.Error()) {

		if pevent, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, pevent, result)
			}

		}

	}
}
