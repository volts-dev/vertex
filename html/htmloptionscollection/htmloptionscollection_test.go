package htmloptionscollection

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`s=document.createElement("select")
	options=s.options`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("options"); test.AssertErr(t, obj.Error()) {

		if b, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLOptionsCollection", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Length", "resultattempt": 0},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("options"); test.AssertErr(t, obj.Error()) {

		if area, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, area, result)
			}

		}

	}
}
