package vhtml

import (
	"context"

	"github.com/volts-dev/logger"
	"github.com/volts-dev/vertex/html/node"
)

func HTML(html string, ctx ...context.Context) *TemplateResult {
	var c context.Context
	if len(ctx) == 0 {
		c = context.Background()
	} else {
		c = ctx[0]
	}

	parser := newParser(html)
	tmpl := newTemplateResult(c)
	err := tmpl.parseHtml(parser)
	if err != nil {
		logger.Errf("HTML parse error: %v", err)
		return nil
	}

	return tmpl
}

func Render(value any, hostNode node.Node) error {
	node := hostNode.GetValueByKey("$vtx")
	if node.IsUndefined() || node.IsNull() {
		// 创建一个新的模板实例
		node = NewContentPart().object.GetObjectValue()
		// 保存到 hostNode 上
		hostNode.SetValueByKey("$vtx", node)
	}

	content, err := NewContentPartFromJSObject(node)
	if err != nil {
		return err
	}

	content.SetValue(value)
	return nil
}
