package location

import (
	"strconv"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`l=document.location`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {
		if l, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Location", l.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Hash", "resultattempt": ""},
	{"method": "Host", "type": "contains", "resultattempt": "localhost"},
	{"method": "Hostname", "resultattempt": "localhost"},
	{"method": "Href", "type": "contains", "resultattempt": "localhost"},
	{"method": "Origin", "type": "contains", "resultattempt": "localhost"},
	{"method": "Pathname", "resultattempt": "/"},
	{"method": "Protocol", "resultattempt": "http:"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if location, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, location, result)
			}

		}

	}
}

func TestPort(t *testing.T) {

	if obj := js.Global().Get("l"); test.AssertErr(t, obj.Error()) {

		if location, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if str, err := location.Port(); test.AssertErr(t, err) {
				if i, err := strconv.Atoi(str); test.AssertErr(t, err) {

					test.AssertExpect(t, true, i > 0)

				}

			}

		}
	}
}
