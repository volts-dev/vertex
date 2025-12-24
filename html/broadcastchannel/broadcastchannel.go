package broadcastchannel

//Full implemented
// https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/eventtarget"
	"github.com/volts-dev/vertex/html/messageevent"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var bcinterface js.Value

// BroadcastChannel struct
type BroadcastChannel struct {
	eventtarget.EventTarget
}

type BroadcastChannelFrom interface {
	BroadcastChannel_() BroadcastChannel
}

func (b BroadcastChannel) BroadcastChannel_() BroadcastChannel {
	return b
}

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		if bcinterface = js.Global().Get("BroadcastChannel"); bcinterface.Error() != nil {
			bcinterface = js.Undefined()
		}

		messageevent.GetInterface()

	})

	return bcinterface
}

// New Get a new channel broadcast
func New(channelname string) (BroadcastChannel, error) {
	var channel BroadcastChannel
	var err error
	if bci := GetInterface(); !bci.IsUndefined() {
		channel.SetObjectValue(bci.New(js.ValueOf(channelname)))
	} else {
		err = ErrNotImplemented
	}
	return channel, err
}

// PostMessage Post a message on channel
func (c BroadcastChannel) PostMessage(message interface{}) error {
	var err error
	err = c.Call("postMessage", js.ValueOf(message)).Error()
	return err
}

// Close Close the channel
func (c BroadcastChannel) Close() error {
	var err error
	err = c.Call("close").Error()

	return err
}

func (c BroadcastChannel) Name() (string, error) {

	return c.GetAttributeString("name")
}
