package broadcastchannel

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/messageevent"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New("alpha"); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object BroadcastChannel]", a.ToString_())

	}
}

func TestChannelIO(t *testing.T) {

	var io chan string = make(chan string)
	var c1, c2 BroadcastChannel
	var err error

	if c1, err = New("alpha"); err == nil {

		if c2, err = New("alpha"); err == nil {

			c2.OnMessage(func(ev messageevent.MessageEvent) {

				if dataObject, err := ev.Data(); test.AssertErr(t, err) {
					io <- dataObject.(string)
				}
			})
		} else {
			t.Errorf("Channel Alpha not created")
		}
		c1.PostMessage("hello")

		select {
		case r := <-io:
			test.AssertExpect(t, "hello", r)

		case <-time.After(time.Duration(100) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
		if err := c2.Close(); err != nil {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf("Channel Alpha not created %s", err.Error())
	}
	if err := c1.Close(); err != nil {
		t.Errorf(err.Error())
	}
}
