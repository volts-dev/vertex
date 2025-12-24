package hello_world

import (
	"context"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

var (
	// NotFound is the ui element that is displayed when a request is not
	// routed.
	NotFound component.Component = &helloWorld{}
)

type helloWorld struct {
	component.Component
	Icon string
}

func init() {
	component.Register("hello-world", New)
}

func New() component.Component {
	return &helloWorld{}
}

func (n *helloWorld) ObservedAttributes() []string {
	return []string{"ABC", "BCD"}
}

func (n *helloWorld) Render(ctx context.Context) *vhtml.TemplateResult {
	return vhtml.HTML(`<div class="goapp-app-info"><p>Hello World</p></div>`)
}

func (n *helloWorld) ConnectedCallback() {
	console.Log("helloWorld connected to the DOM")
}

func (n *helloWorld) DisconnectedCallback() {
	console.Log("helloWorld disconnected from the DOM")
}

func (n *helloWorld) AttributeChangedCallback(name, oldValue, newValue string) {
	console.Log(`helloWorld attribute changed: %v from %v to %v`, name, oldValue, newValue)
}

func (n *helloWorld) OnMount(context.Context) {
	links := window.Default().GetValueByKey("document").Call("getElementsByTagName", "link")

	for i := 0; i < links.Length(); i++ {
		link := links.Index(i)
		rel := link.Call("getAttribute", "rel")

		if js.ValueToString(rel) == "icon" {
			favicon := link.Call("getAttribute", "href")
			n.Icon = js.ValueToString(favicon)
			return
		}
	}
}

func (n *helloWorld) ___Render() component.Component {
	/*return Div().
	Class("goapp-app-info").
	Body(
		Div().
			Class("goapp-notfound-title").
			Body(
				Text("4"),
				Img().
					Class("goapp-logo").
					Alt("0").
					Src(n.Icon),
				Text("4"),
			),
		P().
			Class("goapp-label").
			Text("Not Found"),
	)*/
	return nil
}
