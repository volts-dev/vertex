package component

import (
	"errors"

	"github.com/volts-dev/vertex/core/js"
	"github.com/volts-dev/vertex/element"
)

type (
	Composer interface {
		element.INode

		// Render constructs and returns the visual representation of the component
		// as a node tree.
		Render() element.INode

		setRef(Composer) Composer
		depth() uint
		setDepth(uint) Composer
		parent() element.INode
		root() element.INode
		setRoot(element.INode) Composer
	}

	Component struct {
		treeDepth     uint
		ref           element.INode
		parentElement element.INode
		rootElement   element.INode
	}
)

func NewComponent() *Component {
	return &Component{}
}

func (c *Component) String() string {
	return "Component"
}

// JSValue retrieves the JavaScript value associated with the component's root.
// If the root element isn't defined, it returns a nil JavaScript value.
func (c *Component) Value() js.Value {
	if c.rootElement == nil {
		return js.ValueOf(nil)
	}
	return c.rootElement.Value()
}

// Mounted checks if the component is currently mounted within the UI.
func (c *Component) Mounted() bool {
	return c.ref != nil
}

// Render produces a visual representation of the component's content. This
// default implementation ensures the app.Composer interface is satisfied
// when app.Component is embedded. However, developers are encouraged to redefine
// this method to customize the component's appearance.
func (c *Component) Render() element.INode {
	//componentName := reflect.TypeOf(c.ref).Name()
	return nil
	/*
		return Div().
			DataSet("compo-type", componentName).
			Style("border", "1px solid currentColor").
			Style("padding", "12px 0").
			Body(
				H1().Text("Component "+strings.TrimPrefix(componentName, "*")),
				P().Body(
					Text("Change appearance by implementing: "),
					Code().
						Style("color", "deepskyblue").
						Style("margin", "0 6px").
						Text("func (c "+componentName+") Render() app.UI"),
				),
			)
	*/
}

// ValueTo captures the value of the DOM element (if it exists) that triggered
// an event, and assigns it to the provided receiver. The receiver must be a
// pointer pointing to either a string, integer, unsigned integer, or a float.
// This method panics if the provided value isn't a pointer.
func (c *Component) ValueTo(v any) EventHandler {
	return func(ctx Context, e Event) {
		value := ctx.JSSrc().Get("value")
		if err := stringTo(value.String(), v); err != nil {
			Log(errors.New("storing dom element value failed").Wrap(err))
			return
		}
	}
}

func (c *Component) setRef(v element.IHTMLElement) element.IHTMLElement {
	c.ref = v
	return v
}

func (c *Component) depth() uint {
	return c.treeDepth
}

func (c *Component) setDepth(v uint) element.INode {
	c.treeDepth = v
	return c.ref
}

func (c *Component) parent() element.INode {
	return c.parentElement
}

func (c *Component) setParent(p element.INode) element.INode {
	c.parentElement = p
	return c.ref
}

func (c *Component) root() element.INode {
	return c.rootElement
}

func (c *Component) setRoot(v element.INode) element.INode {
	c.rootElement = v
	return c.ref
}
