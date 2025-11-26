package not_found

import (
	"context"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js/helper"
)

var (
	// NotFound is the ui element that is displayed when a request is not
	// routed.
	NotFound component.Component = &notFound{}
)

type notFound struct {
	component.Component
	Icon string
}

func init() {
	component.Register("not-found", New)
}

func New() component.Component {
	return &notFound{}
}

func (n *notFound) ObservedAttributes() []string {
	return []string{"ABC", "BCD"}
}

func (n *notFound) Render(ctx context.Context) *vhtml.TemplateResult {
	return vhtml.HTML(`<div class="goapp-app-info"> Not Found </div>`)
}

func (n *notFound) ConnectedCallback() {
	console.Log("notFound connected to the DOM")
}

func (n *notFound) DisconnectedCallback() {
	console.Log("notFound disconnected from the DOM")
}

func (n *notFound) AttributeChangedCallback(name, oldValue, newValue string) {
	console.Log(`notFound attribute changed: %v from %v to %v`, name, oldValue, newValue)
}

func (n *notFound) OnMount(context.Context) {
	links := window.Default().GetValueByKey("document").Call("getElementsByTagName", "link")

	for i := 0; i < links.Length(); i++ {
		link := links.Index(i)
		rel := link.Call("getAttribute", "rel")

		if helper.ValueToString(rel) == "icon" {
			favicon := link.Call("getAttribute", "href")
			n.Icon = helper.ValueToString(favicon)
			return
		}
	}
}

func (n *notFound) ___Render() component.Component {
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
