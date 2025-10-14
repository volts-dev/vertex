package htmlprogresselement

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/document"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`
	l= document.createElement("label")
	p= document.createElement("progress")
	l.appendChild(p)`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLProgressElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if meter, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLProgressElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Max", "resultattempt": float64(1)},
	{"method": "SetMax", "args": []interface{}{float64(10.0)}, "gettermethod": "Max", "resultattempt": float64(10.0)},
	{"method": "Position", "resultattempt": float64(-1)},
	{"method": "Value", "resultattempt": float64(0.0)},
	{"method": "SetValue", "args": []interface{}{float64(5.3)}, "gettermethod": "Value", "resultattempt": float64(5.3)},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
