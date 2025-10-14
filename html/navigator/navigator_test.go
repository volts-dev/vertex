package navigator

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`nav=window.navigator`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("nav"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Navigator", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "CookieEnabled", "resultattempt": true},
	//	{"method": "DeviceMemory", "resultattempt": float64(8.0)},
	{"method": "UserAgent", "type": "contains", "resultattempt": "HeadlessChrome"},
	//	{"method": "Language", "resultattempt": "fr"},
	{"method": "Vendor", "resultattempt": "Google Inc."},
	{"method": "JavaEnabled", "resultattempt": false},
	{"method": "Permissions", "type": "constructnamechecking", "resultattempt": "Permissions"},
	{"method": "Clipboard", "type": "constructnamechecking", "resultattempt": "Clipboard"},
	{"method": "ServiceWorker", "type": "constructnamechecking", "resultattempt": "ServiceWorkerContainer"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("nav"); test.AssertErr(t, obj.Error()) {

		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, nav, result)
			}

		}

	}
}
