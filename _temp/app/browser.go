package app

import (
	"time"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/context"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

type (
	nav              struct{}
	appUpdate        struct{}
	appInstallChange struct{}
	resize           struct{}
)

type browser struct {
	AppUpdatable bool

	anchorClick      js.Func
	popState         js.Func
	navigationFromJS js.Func
	appUpdate        js.Func
	appInstallChange js.Func
	appResize        js.Func
	resizeTimer      *time.Timer
}

func (b *browser) HandleEvents(ctx context.Context, notifyComponentEvent func(any)) {
	b.handleAnchorClick(ctx)
	b.handlePopState(ctx)
	b.handleNavigationFromJS(ctx)
	b.handleAppUpdate(ctx, notifyComponentEvent)
	b.handleAppInstallChange(ctx, notifyComponentEvent)
	b.handleAppResize(ctx, notifyComponentEvent)
}

func (b *browser) handleAnchorClick(ctx context.Context) {
	b.anchorClick = js.FuncOf(func(this js.Value, args []js.Value) any {
		ctx.Dispatch(func(context.Context) {
			obj, _ := js.ToObject(args[0])
			event := html.Event{Object: obj}

			for target := event.Get("target"); target.Truthy(); target = target.Get("parentElement") {
				switch js.ValueToString(target.Get("tagName")) {
				case "A":
					if meta := event.Get("metaKey"); meta.Truthy() && js.ValueToBool(meta) {
						return
					}

					if ctrl := event.Get("ctrlKey"); ctrl.Truthy() && js.ValueToBool(ctrl) {
						return
					}

					if download := target.Call("getAttribute", "download"); !download.IsNull() {
						return
					}

					switch js.ValueToString(target.Get("target")) {
					case "_blank":
						return
					}

					event.PreventDefault()
					if href := target.Get("href"); href.Truthy() {
						ctx.Navigate(js.ValueToString(target.Get("href")))
					}
					return

				case "BODY":
					return
				}
			}
		})
		return nil
	})
	DefaultWindow().Set("onclick", b.anchorClick)
}

func (b *browser) handlePopState(ctx context.Context) {
	b.popState = js.FuncOf(func(this js.Value, args []js.Value) any {
		ctx.Dispatch(func(c context.Context) {
			ctx.Navigate(DefaultWindow().URL().Path)
			//			ctx.Navigate(DefaultWindow().URL(), false)
		})
		return nil
	})
	DefaultWindow().Set("onpopstate", b.popState)
}

func (b *browser) handleNavigationFromJS(ctx context.Context) {
	b.navigationFromJS = js.FuncOf(func(this js.Value, args []js.Value) any {
		ctx.Dispatch(func(context.Context) {
			url, _ := args[0].String()
			ctx.Navigate(url)
		})
		return nil
	})
	DefaultWindow().Set("goappNav", b.navigationFromJS)
}

func (b *browser) handleAppUpdate(ctx context.Context, notifyComponentEvent func(any)) {
	appUpdate := func() {
		ctx.Dispatch(func(context.Context) {
			b.AppUpdatable = true
			notifyComponentEvent(appUpdate{})
		})
		ctx.Defer(func(context.Context) {
			console.Log(DefaultWindow().URL().Hostname() + " has been updated, reload to see changes")
		})
	}

	b.appUpdate = js.FuncOf(func(this js.Value, args []js.Value) any {
		appUpdate()
		return nil
	})
	DefaultWindow().Set("goappOnUpdate", b.appUpdate)

	if DefaultWindow().Get("goappUpdatedBeforeWasmLoaded").Truthy() {
		appUpdate()
	}
}

func (b *browser) handleAppInstallChange(ctx context.Context, notifyComponentEvent func(any)) {
	appInstallChange := func() {
		ctx.Dispatch(func(context.Context) {
			notifyComponentEvent(appInstallChange{})
		})
	}

	b.appInstallChange = js.FuncOf(func(this js.Value, args []js.Value) any {
		appInstallChange()
		return nil
	})
	DefaultWindow().Set("goappOnAppInstallChange", b.appInstallChange)

	if DefaultWindow().Get("goappAppInstallChangedBeforeWasmLoaded").Truthy() {
		appInstallChange()
	}
}

func (b *browser) handleAppResize(ctx context.Context, notifyComponentEvent func(any)) {
	const resizeCooldown = time.Millisecond * 250

	b.appResize = js.FuncOf(func(this js.Value, args []js.Value) any {
		ctx.Dispatch(func(context.Context) {
			if b.resizeTimer != nil {
				b.resizeTimer.Stop()
				b.resizeTimer.Reset(resizeCooldown)
				return
			}

			b.resizeTimer = time.AfterFunc(resizeCooldown, func() {
				ctx.Dispatch(func(context.Context) {
					notifyComponentEvent(resize{})
				})
			})
		})
		return nil
	})
	DefaultWindow().Set("onresize", b.appResize)
}
