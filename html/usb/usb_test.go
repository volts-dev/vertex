package usb

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/clipboarditem"
)

var clipitem clipboarditem.ClipboardItem

var methodsAttempt []map[string]interface{}

func TestMain(m *testing.M) {

	reflect.SetSyscall()

	methodsAttempt = []map[string]interface{}{
		{"method": "GetDevices", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "RequestDevices", "args": []interface{}{map[string]interface{}{"vendorId": 0x11}}, "type": "constructnamechecking", "resultattempt": "Promise"},
	}

	js.Eval(`usbobj=navigator.usb`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("usbobj"); test.AssertErr(t, obj.Error()) {
		if usb, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "USB", usb.ConstructName_())

		}
	}
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("usbobj"); test.AssertErr(t, obj.Error()) {

		if clip, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, clip, result)
			}

		}

	}
}
