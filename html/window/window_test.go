package window

import (
	"strings"
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Init()
	m.Run()
}

func TestWindow(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if d, err := w.Document(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object HTMLDocument]", d.ToString())
		}

	}

}

func TestHistory(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if h, err := w.History(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object History]", h.ToString())
		}

	}

}
func TestLocation(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if l, err := w.Location(); test.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(l.ToString(), expect) {
				t.Errorf("Must contain %s have %s", expect, l.ToString())
			}
		}

	}

}

func TestLocalStorage(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if l, err := w.LocalStorage(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Storage]", l.ToString())
		}

	}

}

func TestSessionStorage(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if l, err := w.SessionStorage(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Storage]", l.ToString())
		}

	}

}

func TestIndexdedDB(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if i, err := w.IndexdedDB(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object IDBFactory]", i.ToString())
		}

	}

}

func TestNavigator(t *testing.T) {

	if w, err := New(); test.AssertErr(t, err) {

		if i, err := w.Navigator(); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Navigator]", i.ToString())
		}

	}

}

func TestAtob(t *testing.T) {

	v := Atob("SGVsbG93b3JsZA==")
	test.AssertExpect(t, "Helloworld", v)

}

func TestBtoa(t *testing.T) {

	v := Btoa("Helloworld")
	test.AssertExpect(t, "SGVsbG93b3JsZA==", v)

}
