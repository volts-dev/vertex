package serviceworkercontainer

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`sw=window.navigator.serviceWorker`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("sw"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "ServiceWorkerContainer", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Controller", "type": "error", "resultattempt": ErrControllerNotDefined},
	{"method": "Ready", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "GetRegistration", "args": []interface{}{"url"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Register", "args": []interface{}{"url"}, "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("sw"); test.AssertErr(t, obj.Error()) {

		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, nav, result)
			}

		}

	}
}
