//go:build localtest
// +build localtest

package eventsource

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/messageevent"
)

// use sse echo test from heroku

var sseurl string = "https://sse-echo.herokuapp.com/events?delay=3&data=data%3A%20msg1%0A%0Adata%3A%20msg2%0A%0Aevent%3A%20close%0Adata%3A%20msgend%0A%0A"

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`sse=new EventSource("https://sse-echo.herokuapp.com")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if sse, err := New(sseurl); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object EventSource]", sse.ToString_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj := js.Global().Get("sse"); test.AssertErr(t, obj.Error()) {
		if sse, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object EventSource]", sse.ToString_())

		}
	}

}

func TestOnError(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New("htt://error.com"); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object EventSource]", sse.ToString_())
		sse.SetOnError(func(e event.Event) {

			success <- false
		})

		select {
		case <-success:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No error receive")

		}

	}

}

func TestOnOpen(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New(sseurl); test.AssertErr(t, err) {

		sse.SetOnOpen(func(e event.Event) {

			success <- true

		})

		select {
		case <-success:
		case <-time.After(time.Duration(10000) * time.Millisecond):
			t.Errorf("No open receive")

		}

	}

}

var str_expect []string = []string{"msg1", "msg2"}

func TestOnMessage(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New(sseurl); test.AssertErr(t, err) {
		var i int = 0
		sse.SetOnMessage(func(e messageevent.MessageEvent) {

			if str, err := e.Data(); test.AssertErr(t, err) {
				test.AssertExpect(t, str, str_expect[i])
				i++
			}
			if i == 2 {
				success <- true
			}

		})

		select {
		case <-success:
		case <-time.After(time.Duration(10000) * time.Millisecond):
			t.Errorf("No open receive")

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "ReadyState", "resultattempt": 0},
	{"method": "Url", "resultattempt": "https://sse-echo.herokuapp.com/"},
	{"method": "WithCredentials", "resultattempt": false},
	{"method": "Close"},
}

func TestMethods(t *testing.T) {

	js.Eval(`sse1=new EventSource("https://sse-echo.herokuapp.com")`)

	if obj := js.Global().Get("sse1"); test.AssertErr(t, obj.Error()) {

		if anchor, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				test.InvokeCheck(t, anchor, result)
			}

		}

	}
}
