package main

import (
	"fmt"

	"github.com/volts-dev/vertex/app"
	"github.com/volts-dev/vertex/html/dedicatedworkerglobalscope"
	"github.com/volts-dev/vertex/html/messageevent"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"
)

func main() {
	app.Init()
	fmt.Printf("Get self\n")
	if instance, err := js.Self(); test.AssertErr(nil, err) {

		//if d, ok := instance.(dedicatedworkerglobalscope.DedicatedWorkerGlobalScope); ok {
		if d, err := dedicatedworkerglobalscope.NewFromJSObject(instance); test.AssertErr(nil, err) {
			fmt.Printf("Install handler\n")
			d.PostMessage("installok")
			d.OnMessage(func(m messageevent.MessageEvent) {
				d.PostMessage("testok")
			})
		}

	}

	ch := make(chan struct{})
	<-ch

}
