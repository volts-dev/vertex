package js

import (
	"testing"
	//"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	//reflect.SetSyscall()
	m.Run()
}

/*
func TestBaseObjectString(t *testing.T) {
	var obj Value

	t.Run("1 is Int", func(t *testing.T) {
		Eval("intvalue=1")

		if obj = Global().Get("intvalue"); test.AssertErr(t, obj.Error()) {

			if b, err := IsInteger(obj); test.AssertErr(t, err) {
				test.AssertExpect(t, true, b)
			}

		}

	})

	t.Run("1.3 is Float", func(t *testing.T) {
		Eval("intvalue=1.3")

		if obj = Global().Get("intvalue"); test.AssertErr(t, obj.Error()) {

			if b, err := IsInteger(obj); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)
			}

		}

	})

	t.Run("str is not int", func(t *testing.T) {
		Eval("intvalue='hello'")

		if obj = Global().Get("intvalue"); test.AssertErr(t, obj.Error()) {

			if b, err := IsInteger(obj); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)
			}

		}

	})

}
*/
