package component

import (
	"context"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vcss"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/node"
)

type ReactiveComponent struct {
	shadowRoot        *node.Node
	ContentPart       *vhtml.ContentPart
	StyleInjectionMgr *vcss.StyleInjectionManager
}

func (self *ReactiveComponent) init(rootNode *node.Node) {
	self.shadowRoot = rootNode
}

func (self *ReactiveComponent) initStyle(css *vcss.CSSResult) {
	self.StyleInjectionMgr = vcss.NewStyleInjectionManager()
	// 注入样式到 Shadow DOM
	if css != nil {
		self.StyleInjectionMgr.InjectComponentStyles(
			self.shadowRoot,
			[]*vcss.CSSResult{css},
			&vcss.StyleRenderOptions{
				Minify:                false,
				UseAdoptedStyleSheets: vcss.SupportsAdoptingStyleSheets,
			},
		)
	}
}

func (self *ReactiveComponent) iniDom(htmlResult *vhtml.TemplateResult, ctx context.Context) {
	contentPart, err := vhtml.Render(htmlResult, self.shadowRoot, ctx)
	if err != nil {
		console.Error(err.Error())
	}
	self.ContentPart = contentPart
}

func (self *ReactiveComponent) RequestUpdate() {
	//if instance, ok := GetValue().(*vhtml.TemplateInstance); ok {
	self.ContentPart.Instance.Update()
	//}
}
