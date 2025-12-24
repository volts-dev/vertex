package context

import (
	"context"
	"net/url"
	"time"

	"github.com/volts-dev/vertex/html/node"
	//"github.com/volts-dev/vertex/component"
	//"github.com/volts-dev/vertex/component"
)

type Context struct {
	context.Context

	//page         func() Page
	appUpdatable bool
	resolveURL   func(string) string
	navigate     func(*url.URL, bool)
	//localStorage          Storage
	//sessionStorage        Storage
	dispatch func(func())
	defere   func(func())
	async    func(func())
	//addComponentUpdate    func(component.Component, int)
	//removeComponentUpdate func(component.Component)
	//handleAction          func(string, UI, bool, ActionHandler)
	//postAction            func(Context, Action)
	//observeState          func(Context, string, any) Observer
	//getState              func(Context, string, any)
	//setState              func(Context, string, any) State
	delState func(Context, string)

	sourceElement node.Node
	//notifyComponentEvent func(Context, html.Node, any)
}

// Navigate transitions to the given URL string.
func (ctx Context) Navigate(rawURL string) {
	u, err := url.Parse(rawURL)
	if err != nil {
		//Log(errors.New("navigating to URL failed").
		//	WithTag("url", rawURL).
		//	Wrap(err))
		return
	}
	ctx.NavigateTo(u)
}

// NavigateTo transitions to the provided URL.
func (ctx Context) NavigateTo(u *url.URL) {
	ctx.navigate(u, true)
}

// Dispatch prompts the execution of a function on the UI goroutine,
// flagging the enclosing component for an update.
func (ctx Context) Dispatch(v func(Context)) {
	ctx.dispatch(func() {
		if c, _ := ctx.sourceElement.ParentNode(); c.Object.GetObjectValue().IsNull() || c.Object.GetObjectValue().IsUndefined() {
			return
		}

		//for c, err := ctx.sourceElement.ParentNode(); err == nil; c, err = ctx.sourceElement.ParentNode() {
		//ctx.addComponentUpdate(c, 1)
		//}

		if v != nil {
			v(ctx)
		}
	})
}

// Defer postpones the function execution on the UI goroutine until the
// current update cycle completes.
func (ctx Context) Defer(v func(Context)) {
	ctx.defere(func() {
		if c, _ := ctx.sourceElement.ParentNode(); c.Object.GetObjectValue().IsNull() || c.Object.GetObjectValue().IsUndefined() {
			return
		}

		if v != nil {
			v(ctx)
		}
	})
}

// Async initiates a function asynchronously. It enables go-app to monitor
// goroutines, ensuring they conclude when rendering server-side.
func (ctx Context) Async(v func()) {
	ctx.async(v)
}

// After pauses for a determined span, then triggers a specified function.
func (ctx Context) After(d time.Duration, f func(Context)) {
	ctx.async(func() {
		time.Sleep(d)
		ctx.Dispatch(f)
	})
}

// 要求更新
func (ctx Context) RequestUpdate() {
	//for c := ctx.sourceElement; c != nil; c, _ = c.ParentNode() {
	//	ctx.addComponentUpdate(c, 1)
	//}
}

// PreventUpdate halts updates for the enclosing component.
func (ctx Context) PreventUpdate() {
	//for c := ctx.sourceElement; c != nil; c, _ = c.ParentNode() {
	//	ctx.addComponentUpdate(c, -1)
	//}
}

// Update flags the enclosing component for an update.
func (ctx Context) Update() {
	ctx.Dispatch(nil)
}
