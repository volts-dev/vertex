//go:build localtest
// +build localtest

package websocket

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/html/messageevent"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if ws, err := New("wss://ws.ifelse.io"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object WebSocket]", ws.ToString_())
	}
}

func TestEcho(t *testing.T) {
	var io chan bool = make(chan bool)
	var nbmsg int = 0
	if w, err := New("wss://ws.ifelse.io"); test.AssertErr(t, err) {

		w.SetOnMessage(func(e messageevent.MessageEvent) {
			if nbmsg == 0 {
				w.Send("hogosuru")
			} else {
				if message, err := e.Data(); test.AssertErr(t, err) {
					if s, ok := message.(string); ok {
						if s == "hogosuru" {
							io <- true
						} else {
							t.Errorf("Serveur must send hogosuru pong receive %s", s)
						}

					} else {
						t.Error("Response must be a string")
					}

				}
			}
			nbmsg++

		})

	}
	select {
	case <-io:
	case <-time.After(time.Duration(2000) * time.Millisecond):
		t.Errorf("No message channel receive")
	}
}
