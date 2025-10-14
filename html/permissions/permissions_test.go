package permissions

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {

	reflect.SetSyscall()

	js.Eval(`perm=navigator.permissions`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("perm"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Permissions", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Query", "args": []interface{}{map[string]interface{}{"name": "clipboard-read"}}, "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("perm"); test.AssertErr(t, obj.Error()) {

		if clip, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, clip, result)
			}

		}

	}
}
