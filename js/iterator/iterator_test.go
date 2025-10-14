package iterator

import (
	"errors"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`str = "hp";
	it=str[Symbol.iterator]()
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("it"); test.AssertErr(t, obj.Error()) {

		if it, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Iterator", it.ConstructName_())
		}

	}

}

func TestNext(t *testing.T) {

	if obj := js.Global().Get("it"); test.AssertErr(t, obj.Error()) {

		if it, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			_, val, err := it.Next()
			test.AssertExpect(t, "h", val)
			test.AssertExpect(t, true, errors.Is(err, nil))
			_, val, err = it.Next()
			test.AssertExpect(t, "p", val)
			test.AssertExpect(t, true, errors.Is(err, nil))
			_, val, err = it.Next()
			test.AssertExpect(t, true, errors.Is(err, EOI))
		}

	}

}
