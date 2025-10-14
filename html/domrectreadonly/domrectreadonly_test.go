package domrectreadonly

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`r=new DOMRect()
	r.x=3
	r.y=8
	r.width=10
	r.height=13
	ro=DOMRectReadOnly.fromRect(r)
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("ro"); test.AssertErr(t, obj.Error()) {
		if rect, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object DOMRectReadOnly]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "X", "resultattempt": float64(3)},
	{"method": "Width", "resultattempt": float64(10)},
	{"method": "Right", "resultattempt": float64(13)},
	{"method": "Left", "resultattempt": float64(3)},
	{"method": "Y", "resultattempt": float64(8)},
	{"method": "Height", "resultattempt": float64(13)},
	{"method": "Top", "resultattempt": float64(8)},
	{"method": "Bottom", "resultattempt": float64(21)},
}

func TestMethods(t *testing.T) {

	if obj := js.Global().Get("ro"); test.AssertErr(t, obj.Error()) {

		if meta, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, meta, result)
			}

		}

	}
}
