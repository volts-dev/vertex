package dropdown

import (
	"context"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/htmlelement"
)

type (
	Dropdown struct {
		component.Component
		arrow    bool
		position string
		//slots
		item           []any
		holdOnHover    bool
		direction      string //	"dropdown","dropup","dropstart","dropend",
		target         *htmlelement.HtmlElement
		OnOpened       func()
		OnStateChanged func(bool)
	}
)

func (self *Dropdown) Constructor() {

}

func (self *Dropdown) ConnectedCallback() {

}

func (self *Dropdown) Render(ctx context.Context) *vhtml.TemplateResult {
	return vhtml.HTML("<slot></slot>")
}

func (self *Dropdown) setTargetElement(target *htmlelement.HtmlElement) {

}

func (self *Dropdown) setTargetDirectionClass(direction string) {
	if self.target != nil {
		return
	}

	directionClasses := []string{
		"dropdown",
		"dropup",
		"dropstart",
		"dropend",
	}

	lst, _ := self.target.ClassList()
	lst.Remove(directionClasses...)
	lst.Add(direction)
}

func (self *Dropdown) openPopover() {

}

func (self *Dropdown) closePopover() {

}

func (self *Dropdown) onOpened() {
	//self.navigation.enable()
	self.OnOpened()
	self.OnStateChanged(true)

	if self.target != nil {
		self.target.SetAriaExpanded("true")
		lst, _ := self.target.ClassList()
		lst.Add("show")
	}
}

func (self *Dropdown) onClosed() {
	self.OnStateChanged(false)

	if self.target != nil {
		self.target.SetAriaExpanded("false")
		lst, _ := self.target.ClassList()
		lst.Remove("show")
		self.setTargetDirectionClass(self.direction)
	}
}
