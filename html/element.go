//go:build js && wasm

package html

import (
	"github.com/volts-dev/volts/vertex/core/js"
)

type (
	IElement interface {
		// JSValue returns the javascript value linked to the element.
		Value() js.Value

		// Reports whether the element is mounted.
		Mounted() bool

		parent() IElement
		setParent(IElement) IElement
	}

	HtmlElement[T any] struct {
		value         *T
		tag           string
		xmlns         string
		treeDepth     uint
		isSelfClosing bool
		jsElement     js.Value
		//attributes    attributes
		//eventHandlers eventHandlers
		parentElement IElement
		children      []IElement
	}
)

func (e HtmlElement[T]) New() *T {
	a := new(HtmlElement[T]) //.(HtmlElement[T])
	a.value = new(T)
	return a.value
}
func (e *HtmlElement[T]) setBody(v []IElement) *T {
	e.children = v
	return e.value
}
