package file

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/js/array"
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

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object File]", f.ToString_())

	}

}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })")

	if obj := js.Global().Get("file"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object File]", d.ToString_())

		}
	}

}

func TestName(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		if name, err := f.Name(); test.AssertErr(t, err) {

			test.AssertExpect(t, "chucknorris.png", name)
		}

	}

}

func TestType(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		if typefile, err := f.Type(); test.AssertErr(t, err) {

			test.AssertExpect(t, "image/png", typefile)
		}

	}

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png"); test.AssertErr(t, err) {

		if typefile, err := f.Type(); test.AssertErr(t, err) {

			test.AssertExpect(t, "", typefile)
		}

	}

}

func TestLastModifiedDate(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		if lastmodified, err := f.LastModifiedDate(); test.AssertErr(t, err) {

			test.AssertExpect(t, "Date", lastmodified.ConstructName_())
		}

	}

}

func TestLastModified(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); test.AssertErr(t, err) {

		if lastmodified, err := f.LastModified(); test.AssertErr(t, err) {

			test.AssertExpect(t, true, lastmodified > 0)
		}

	}

}
