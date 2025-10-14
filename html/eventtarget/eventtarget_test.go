package eventtarget

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/event"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`event=new EventTarget()
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if e, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object EventTarget]", e.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("event"); test.AssertErr(t, obj.Error()) {
		if event, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object EventTarget]", event.ToString_())

		}
	}

}

func TestEvent(t *testing.T) {

	//var io chan bool = make(chan bool)
	var eventreceive bool = false

	if e, err := New(); test.AssertErr(t, err) {

		if _, err := e.AddEventListener("test", func(e event.Event) {
			eventreceive = true
		}); test.AssertErr(t, err) {

			ev, _ := event.New("test")
			e.DispatchEvent(ev)
			test.AssertExpect(t, true, eventreceive)
		}

	}

}

func TestRemoveEventListener(t *testing.T) {

	if e, err := New(); test.AssertErr(t, err) {

		if f, err := e.AddEventListener("test", func(e event.Event) {

		}); test.AssertErr(t, err) {

			test.AssertErr(t, e.RemoveEventListener(f, "test"))

		}

	}

}
