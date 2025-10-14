package formdata

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/htmlformelement"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()

	m.Run()
}

func TestNew(t *testing.T) {

	if f, err := New(); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object FormData]", f.ToString_())
	}
	js.Eval(`f= document.createElement("form")
	intext=document.createElement("input")
	intext.type="text"
	intext.name="hello"
	intext.value="world"
	f.appendChild(intext)
	`)

	if obj := js.Global().Get("f"); test.AssertErr(t, obj.Error()) {

		if form, err := htmlformelement.NewFromJSObject(obj); test.AssertErr(t, err) {

			if f, err := New(form); test.AssertErr(t, err) {

				if v, err := f.Get("hello"); test.AssertErr(t, err) {

					test.AssertExpect(t, "world", v)
				}

				_, err := f.Get("hell")
				test.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

			}
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval(`f= new FormData()
	`)

	if obj := js.Global().Get("f"); test.AssertErr(t, obj.Error()) {

		if f, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object FormData]", f.ToString_())
		}

	}
}

func TestGet(t *testing.T) {

	js.Eval(`f= new FormData()
	f.append("hello","world")
	`)

	if obj := js.Global().Get("f"); test.AssertErr(t, obj.Error()) {

		if f, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			if v, err := f.Get("hello"); test.AssertErr(t, err) {

				test.AssertExpect(t, "world", v)
			}

			_, err := f.Get("hell")
			test.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

		}

	}

}

func TestAppend(t *testing.T) {

	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))
		if v, err := f.Get("data"); test.AssertErr(t, err) {

			test.AssertExpect(t, "test", v)
		}

	}
}

func TestDelete(t *testing.T) {

	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))

		test.AssertErr(t, f.Delete("data"))

		_, err := f.Get("data")
		test.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

	}
}

func TestEntries(t *testing.T) {
	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))
		test.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Entries(); test.AssertErr(t, err) {

			var expecdata map[string]string = map[string]string{"data": "test", "hello": "world"}

			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

				_, ok := expecdata[index.(string)]
				test.AssertExpect(t, true, ok)

				test.AssertExpect(t, expecdata[index.(string)], value.(string))

			}
		}
	}

}

func TestHas(t *testing.T) {

	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))

		if ok, err := f.Has("data"); test.AssertErr(t, err) {
			test.AssertExpect(t, true, ok)

		}

		if ok, err := f.Has("c"); test.AssertErr(t, err) {
			test.AssertExpect(t, false, ok)

		}

	}

}

func TestKeys(t *testing.T) {
	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))
		test.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Keys(); test.AssertErr(t, err) {

			var expecdata map[string]int = map[string]int{"data": 0, "hello": 1}
			var i int
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				v, ok := expecdata[value.(string)]
				test.AssertExpect(t, true, ok)
				test.AssertExpect(t, i, v)
				i++
			}
			test.AssertExpect(t, 2, i)

		}
	}

}

func TestSet(t *testing.T) {
	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))
		test.AssertErr(t, f.Append("hello", "world"))
		test.AssertErr(t, f.Set("hello", "you"))

		if v, err := f.Get("hello"); test.AssertErr(t, err) {

			test.AssertExpect(t, "you", v)
		}

	}

}

func TestValues(t *testing.T) {
	if f, err := New(); test.AssertErr(t, err) {
		test.AssertErr(t, f.Append("data", "test"))
		test.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Values(); test.AssertErr(t, err) {

			var expecdata map[string]int = map[string]int{"test": 0, "world": 1}
			var i int
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				v, ok := expecdata[value.(string)]
				test.AssertExpect(t, true, ok)
				test.AssertExpect(t, i, v)
				i++
			}
			test.AssertExpect(t, 2, i)

		}
	}

}
