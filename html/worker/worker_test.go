package worker

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/messageevent"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`myWorker = new Worker('assets_test/script_test.js');`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("myWorker"); test.AssertErr(t, obj.Error()) {
		if nav, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Worker", nav.ConstructName_())

		}
	}

}

func TestNew(t *testing.T) {
	var wchan chan string = make(chan string)

	if w, err := New("assets_test/script_test.js"); test.AssertErr(t, err) {

		w.OnMessage(func(m messageevent.MessageEvent) {

			if d, err := m.Data(); test.AssertErr(t, err) {

				if message, ok := d.(string); ok {
					wchan <- message
				}

			}

		})

		select {
		case message := <-wchan:
			test.AssertExpect(t, message, "installok")
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("ServiceWorker request timeout")

		}

		w.PostMessage("test")

		select {
		case message := <-wchan:
			test.AssertExpect(t, message, "testok")
		case <-time.After(time.Duration(20000) * time.Millisecond):
			t.Errorf("ServiceWorker request timeout")

		}

		test.AssertErr(t, w.Terminate())

	}

}
