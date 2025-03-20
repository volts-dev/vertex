//go:build js && wasm

package vertex

import "github.com/volts-dev/volts/vertex/core/js"

type (
	IHtmlElement interface {
		IElement
	}

	HtmlElement struct {
		Element
	}
)

func init() {
	DefineElement("html", HtmlElementConstructor)
}

func HtmlElementConstructor() IElement {
	fff := &(struct {
		Element
	})
	return fff{}
}
func (e *HtmlElement) Value() js.Value {
	return *e.object
}

func (e *HtmlElement) Mounted() bool {
	return e.object != nil
}

func (e *HtmlElement) Tag() string {
	return e.tag
}

func (e *HtmlElement) XMLNamespace() string {
	return e.xmlns
}

func (e *HtmlElement) SelfClosing() bool {
	return e.isSelfClosing
}

func (e *HtmlElement) depth() uint {
	return e.treeDepth
}

func (e *HtmlElement) attrs() attributes {
	return e.attributes
}

func (e *HtmlElement) setAttr(name string, value any) {
	if e.attributes == nil {
		e.attributes = make(attributes)
	}
	e.attributes.Set(name, value)
}

func (e *HtmlElement) events() eventHandlers {
	return e.eventHandlers
}

func (e *HtmlElement) setEventHandler(event string, h EventHandler, options ...EventOption) {
	if e.eventHandlers == nil {
		e.eventHandlers = make(eventHandlers)
	}
	e.eventHandlers.Set(event, h, options...)
}

func (e *HtmlElement) parent() UI {
	return e.parentElement
}

func (e *HtmlElement) body() []UI {
	return e.children
}
