package component

import (
	"errors"
	"reflect"
	"strings"

	"github.com/volts-dev/vertex/core/js"
	"github.com/volts-dev/vertex/element"
)

type (
	___IComponent interface {
		String() string
		// JSValue returns the javascript value linked to the element.
		JSValue() js.Value

		// Reports whether the element is mounted.
		Mounted() bool

		parent() ___IComponent
		setParent(___IComponent) ___IComponent
	}

	Component struct {
		treeDepth     uint
		ref           element.IElement
		parentElement element.IElement
		rootElement   element.IElement
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
func (c *Component) Render() IElement {
	componentName := reflect.TypeOf(c.ref).Name()

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

func (c *Component) setRef(v IElement) IElement {
	c.ref = v
	return v
}

func (c *Component) depth() uint {
	return c.treeDepth
}

func (c *Component) setDepth(v uint) Composer {
	c.treeDepth = v
	return c.ref
}

func (c *Component) parent() IElement {
	return c.parentElement
}

func (c *Component) setParent(p IElement) IElement {
	c.parentElement = p
	return c.ref
}

func (c *Component) root() IElement {
	return c.rootElement
}

func (c *Component) setRoot(v IElement) IElement {
	c.rootElement = v
	return c.ref
}
