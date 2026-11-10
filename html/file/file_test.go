package file

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

/*
file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
*/

func TestNew(t *testing.T) {

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object File]", f.ToString())

	}

}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })")

	if obj := js.Global().Get("file"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object File]", d.ToString())

		}
	}

}

func TestName(t *testing.T) {

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "chucknorris.png", f.Name())

	}

}

func TestType(t *testing.T) {

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "image/png", f.Type())

	}

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png"); test.AssertErr(t, err) {

		test.AssertExpect(t, "", f.Type())

	}

}

func TestLastModifiedDate(t *testing.T) {

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		if lastmodified, err := f.LastModifiedDate(); test.AssertErr(t, err) {

			test.AssertExpect(t, "Date", lastmodified.ConstructName_())
		}

	}

}

func TestLastModified(t *testing.T) {

	if f, err := New(js.ArrayFrom_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		test.AssertExpect(t, true, f.LastModified() > 0)

	}

}
