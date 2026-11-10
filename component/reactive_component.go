package component

import (
	"context"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vcss"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/documentfragment"
	"github.com/volts-dev/vertex/html/node"
)

type ReactiveComponent struct {
	ShadowRoot        *documentfragment.DocumentFragment
	ContentPart       *vhtml.ContentPart
	StyleInjectionMgr *vcss.StyleInjectionManager
	Disabled          bool
}

func (self *ReactiveComponent) init(rootNode *node.Node) error {
	self.Disabled = true

	var err error
	self.ShadowRoot, err = documentfragment.NewFromJSObject(rootNode.GetObjectValue())
	if err != nil {
		return err
	}

	return nil
}

func (self *ReactiveComponent) initStyle(css *vcss.CSSResult) {
	self.StyleInjectionMgr = vcss.NewStyleInjectionManager()
	// 注入样式到 Shadow DOM
	if css != nil {
		self.StyleInjectionMgr.InjectComponentStyles(
			&self.ShadowRoot.Node,
			[]*vcss.CSSResult{css},
			&vcss.StyleRenderOptions{
				Minify:                false,
				UseAdoptedStyleSheets: vcss.SupportsAdoptingStyleSheets,
			},
		)
	}
}

func (self *ReactiveComponent) iniDom(htmlResult *vhtml.TemplateResult, ctx context.Context) {
	contentPart, err := vhtml.Render(htmlResult, &self.ShadowRoot.Node, ctx)
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
