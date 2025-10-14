package htmlmeterelement

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
	m= document.createElement("meter")
	l.appendChild(m)`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); test.AssertErr(t, err) {
		if b, err := New(doc); test.AssertErr(t, err) {
			test.AssertExpect(t, "HTMLMeterElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("m"); test.AssertErr(t, obj.Error()) {

		if meter, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "HTMLMeterElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Max", "resultattempt": float64(1)},
	{"method": "SetMax", "args": []interface{}{float64(10.0)}, "gettermethod": "Max", "resultattempt": float64(10.0)},
	{"method": "Min", "resultattempt": float64(0)},
	{"method": "SetMin", "args": []interface{}{float64(2.0)}, "gettermethod": "Min", "resultattempt": float64(2.0)},

	{"method": "High", "resultattempt": float64(10)},
	{"method": "SetHigh", "args": []interface{}{float64(2.3)}, "gettermethod": "High", "resultattempt": float64(2.3)},
	{"method": "Low", "resultattempt": float64(2.0)},
	{"method": "SetLow", "args": []interface{}{float64(2)}, "gettermethod": "Low", "resultattempt": float64(2)},

	{"method": "Optimum", "resultattempt": float64(6.0)},
	{"method": "SetOptimum", "args": []interface{}{float64(5.3)}, "gettermethod": "Optimum", "resultattempt": float64(5.3)},

	{"method": "Value", "resultattempt": float64(2.0)},
	{"method": "SetValue", "args": []interface{}{float64(5.3)}, "gettermethod": "Value", "resultattempt": float64(5.3)},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("m"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
