//go:build js && wasm

package popover

import (
	"context"
	"fmt"
	"strings"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/core/popper"
	"github.com/volts-dev/vertex/core/vcss"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

var (
	// Store popovers for targets
	popovers = make(map[js.Value]js.Value)
)

func GetPopoverForTarget(target js.Value) js.Value {
	return popovers[target]
}

type (
	Popover struct {
		component.Component
		Props Props

		// Internal state
		popperInstance *popper.Instance
		popoverRef     element.Element
		targetElement  interface{} // Element or VirtualElement

		// Cleanup functions
		cleanupFns []func()
	}
)

func init() {
	component.Register("web-popover", func() component.IComponent {
		return &Popover{
			Props: Props{
				Animation:     true,
				Arrow:         true,
				CloseOnEscape: true,
				Position:      "bottom",
				CloseOnClickAway: func(target interface{}) bool {
					return true
				},
			},
		}
	})
}

func (p *Popover) Styles() *vcss.CSSResult {
	return vcss.CSS(`
.o_popover.o-popover--with-arrow {
    margin: 0.5rem;
}
.o_popover .popover-arrow {
    position: absolute;
}
.o_popover.o-popover--ts .popover-arrow,
.o_popover.o-popover--bs .popover-arrow {
    left: 0.3rem;
}
.o_popover.o-popover--te .popover-arrow,
.o_popover.o-popover--be .popover-arrow {
    right: 0.3rem;
}
.o_popover.o-popover--ls .popover-arrow,
.o_popover.o-popover--rs .popover-arrow {
    top: 0.3rem;
}
.o_popover.o-popover--le .popover-arrow,
.o_popover.o-popover--re .popover-arrow {
    bottom: 0.3rem;
}
	`)
}

func (p *Popover) Render(ctx context.Context) *vhtml.TemplateResult {
	classes := []string{"o_popover", "popover", "mw-100"}
	if p.Props.Arrow {
		classes = append(classes, "o-popover--with-arrow")
	}
	if p.Props.Class != "" {
		classes = append(classes, p.Props.Class)
	}

	role := p.Props.Role
	if role == "" {
		role = "tooltip"
	}

	arrowHtml := ""
	if p.Props.Arrow {
		arrowHtml = `<div class="popover-arrow"></div>`
	}

	return vhtml.HTML(`
		<div class="` + strings.Join(classes, " ") + `" role="` + role + `">
			<div class="popover-content">
				<slot></slot>
			</div>
			` + arrowHtml + `
		</div>
	`)
}

func (p *Popover) ConnectedCallback() {
	p.Component.ConnectedCallback()

	// 1. Find Target
	target := p.findTarget()
	if target == nil {
		fmt.Println("Popover: Target not found")
		return
	}
	p.targetElement = target

	// 2. Initialize Popper (must happen after target is found)
	p.initPopper()
}

func (p *Popover) FirstUpdate() error {
	p.setupEventListeners()
	return nil
}

func (p *Popover) DisconnectedCallback() {
	if p.popperInstance != nil {
		instance := *p.popperInstance
		if instance.Destroy != nil {
			instance.Destroy()
		}
		p.popperInstance = nil
	}

	for _, fn := range p.cleanupFns {
		fn()
	}
	p.cleanupFns = []func(){}

	if p.targetElement != nil {
		if el, ok := p.targetElement.(element.Element); ok {
			delete(popovers, el.GetObjectValue())
		}
	}

	p.Component.DisconnectedCallback()
}

func (p *Popover) findTarget() interface{} {
	if str, ok := p.Props.Target.(string); ok {
		w := window.Default()
		doc, err := w.Document()
		if err != nil {
			return nil
		}
		el, err := doc.QuerySelector(str)
		if err == nil {
			return el
		}
	}

	if el, ok := p.Props.Target.(element.Element); ok {
		return el
	}

	return nil
}

func (p *Popover) ObservedAttributes() []string {
	return []string{"target", "position", "arrow", "animation", "fixedposition", "class"}
}

func (p *Popover) AttributeChangedCallback(name, oldValue, newValue string) {
	switch name {
	case "target":
		p.Props.Target = newValue
	case "position":
		p.Props.Position = newValue
	case "arrow":
		p.Props.Arrow = newValue != "false"
	case "animation":
		p.Props.Animation = newValue != "false"
	case "fixedposition":
		p.Props.FixedPosition = newValue == "true"
	case "class":
		p.Props.Class = newValue
	}

	// Trigger popper re-init if needed
	if p.targetElement != nil {
		p.initPopper()
	}
}

func (p *Popover) initPopper() {
	if p.targetElement == nil {
		return
	}

	if p.popperInstance != nil {
		(*p.popperInstance).Destroy()
		p.popperInstance = nil
	}

	shadowRoot := p.Base().ShadowRoot
	if shadowRoot == nil {
		fmt.Println("Popover: ShadowRoot is nil")
		return
	}

	popoverEl, err := shadowRoot.QuerySelector(".o_popover")
	if err != nil {
		fmt.Println("Popover: element not found in shadow root")
		return
	}
	p.popoverRef = popoverEl

	opts := popper.Options{
		Placement: popper.Placement(p.Props.Position),
		Strategy:  popper.Absolute,
		Modifiers: []popper.Modifier{
			{
				Name: "arrow",
				Options: map[string]interface{}{
					"element": ".popover-arrow",
				},
			},
			{
				Name: "offset",
				Options: map[string]interface{}{
					"offset": []int{0, 10},
				},
			},
		},
		OnFirstUpdate: func(state popper.State) {
			p.onPositioned(state)

			// Animation logic
			if p.Props.Animation {
				p.runAnimation(state)
			}

			if p.Props.OnPositioned != nil {
				p.Props.OnPositioned(state.Elements.Popper, state.Placement)
			}
		},
	}

	if p.Props.FixedPosition {
		opts.Strategy = popper.Fixed
	}

	instance := popper.CreatePopper(p.targetElement, p.popoverRef, opts)
	p.popperInstance = &instance

	if el, ok := p.targetElement.(element.Element); ok {
		popovers[el.GetObjectValue()] = p.popoverRef.GetObjectValue()
	}

	// 3. Setup Target Observer
	p.setupTargetObserver()
}

func (p *Popover) runAnimation(state popper.State) {
	// Only animate once
	p.Props.Animation = false

	direction := strings.Split(string(state.Placement), "-")[0]
	transforms := map[string][]string{
		"top":    {"translateY(-5%)", "translateY(0)"},
		"right":  {"translateX(5%)", "translateX(0)"},
		"bottom": {"translateY(5%)", "translateY(0)"},
		"left":   {"translateX(-5%)", "translateX(0)"},
	}

	transform, ok := transforms[direction]
	if !ok {
		transform = []string{"translateY(0)", "translateY(0)"}
	}

	// Use Web Animations API via JS
	p.popoverRef.GetObjectValue().Call("animate", map[string]interface{}{
		"opacity":   []interface{}{0, 1},
		"transform": []interface{}{transform[0], transform[1]},
	}, map[string]interface{}{
		"duration": 200,
		"easing":   "ease-out",
	})
}

func (p *Popover) setupTargetObserver() {
	if el, ok := p.targetElement.(element.Element); ok {
		parent, err := el.ParentNode()
		if err != nil || parent == nil {
			return
		}

		// Create MutationObserver
		obsCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			// Check if target is still connected
			connected := el.GetObjectValue().Get("isConnected").Bool()
			if !connected {
				if p.Props.Close != nil {
					p.Props.Close()
				}
			}
			return nil
		})

		observer := js.Global().Get("MutationObserver").New(obsCallback)
		observer.Call("observe", parent.GetObjectValue(), map[string]interface{}{
			"childList": true,
		})

		p.cleanupFns = append(p.cleanupFns, func() {
			observer.Call("disconnect")
			obsCallback.Release()
		})
	}
}

func (p *Popover) onPositioned(state popper.State) {
	placement := string(state.Placement)
	parts := strings.Split(placement, "-")
	direction := parts[0]
	variant := "middle"
	if len(parts) > 1 {
		variant = parts[1]
	}

	// Calculate short position for classes (e.g. "bs" for bottom-start)
	posDir := ""
	if len(direction) > 0 {
		posDir = string(direction[0])
	}
	posVar := ""
	if len(variant) > 0 {
		posVar = string(variant[0])
	}
	position := posDir + posVar

	// Build final class list
	baseClasses := []string{"o_popover", "popover", "mw-100"}
	if p.Props.Arrow {
		baseClasses = append(baseClasses, "o-popover--with-arrow")
	}
	if p.Props.Class != "" {
		baseClasses = append(baseClasses, p.Props.Class)
	}

	directionMap := map[string]string{
		"top":    "top",
		"bottom": "bottom",
		"left":   "start",
		"right":  "end",
	}

	bsDir := directionMap[direction]
	if bsDir == "" {
		bsDir = direction
	}

	dynamicClasses := []string{
		"bs-popover-" + bsDir,
		"o-popover-" + direction,
		"o-popover--" + position,
	}

	finalClass := strings.Join(append(baseClasses, dynamicClasses...), " ")
	p.popoverRef.SetClassName(finalClass)

	// Arrow positioning
	if p.Props.Arrow {
		arrowEl, err := p.popoverRef.QuerySelector(".popover-arrow")
		if err == nil {
			arrowEl.SetClassName("popover-arrow")
			arrowList, _ := arrowEl.ClassList()

			switch position {
			case "tm", "bm", "tf", "bf":
				arrowList.Add("start-0", "end-0", "mx-auto")
			case "lm", "rm", "lf", "rf":
				arrowList.Add("top-0", "bottom-0", "my-auto")
			case "ts", "bs":
				arrowList.Add("end-auto")
			case "te", "be":
				arrowList.Add("start-auto")
			case "ls", "rs":
				arrowList.Add("bottom-auto")
			case "le", "re":
				arrowList.Add("top-auto")
			}
		}
	}
}

func (p *Popover) setupEventListeners() {
	win := window.Default()

	// Pointerdown handler for click-away
	pointerDownHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			ev := args[0]
			// Use composedPath to get the actual target (even if in shadow DOM)
			path := ev.Call("composedPath")
			if path.Length() > 0 {
				target := path.Index(0)
				p.onClickAway(target)
			}
		}
		return nil
	})

	// Blur handler for iframes
	blurHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			ev := args[0]
			target := ev.Get("relatedTarget")
			if target.IsUndefined() || target.IsNull() {
				target = js.Global().Get("document").Get("activeElement")
			}

			if target.Truthy() && target.Get("tagName").String() == "IFRAME" {
				p.onClickAway(target)
			}
		}
		return nil
	})

	// Add event listeners with capture = true using direct JS call
	win.GetObjectValue().Call("addEventListener", "pointerdown", pointerDownHandler, true)
	win.GetObjectValue().Call("addEventListener", "blur", blurHandler, true)

	p.cleanupFns = append(p.cleanupFns, func() {
		win.GetObjectValue().Call("removeEventListener", "pointerdown", pointerDownHandler, true)
		pointerDownHandler.Release()
	})

	p.cleanupFns = append(p.cleanupFns, func() {
		win.GetObjectValue().Call("removeEventListener", "blur", blurHandler, true)
		blurHandler.Release()
	})

	if p.Props.CloseOnEscape {
		keydownFn := func(ev event.Event) error {
			if ev.GetObjectValue().Get("key").String() == "Escape" {
				if p.Props.Close != nil {
					p.Props.Close()
				}
			}
			return nil
		}
		cleanupKey, err := win.AddEventListener("keydown", keydownFn)
		if err == nil {
			p.cleanupFns = append(p.cleanupFns, func() {
				cleanupKey.Release()
				win.RemoveEventListenerWithFunc("keydown", cleanupKey)
			})
		}
	}
}

func (p *Popover) isInside(target js.Value) bool {
	// Check target
	if el, ok := p.targetElement.(element.Element); ok {
		if el.GetObjectValue().Call("contains", target).Bool() {
			return true
		}
	}

	// Check popover
	if p.popoverRef.GetObjectValue().Call("contains", target).Bool() {
		return true
	}

	return false
}

func (p *Popover) onClickAway(target js.Value) {
	if p.Props.CloseOnClickAway != nil && !p.Props.CloseOnClickAway(target) {
		return
	}

	if !p.isInside(target) {
		if p.Props.Close != nil {
			p.Props.Close()
		}
	}
}
