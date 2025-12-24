package vhtml

import (
	"context"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/global"
	"github.com/volts-dev/vertex/html/node"
)

type (
	Component interface {
		Render(context.Context) *TemplateResult
	}

	RenderOptions struct {
		Component    Component
		host         *node.Node
		renderBefore *node.Node
	}

	TemplateResult struct {
		typ  ResultType
		html string //bytes.Buffer
		comp Component
	}
)

func HTML(html string) *TemplateResult {
	return &TemplateResult{
		html: html,
		typ:  HTML_RESULT,
	}
}

func Render(value any, hostNode *node.Node, opts *RenderOptions) (*ContentPart, error) {
	node := hostNode.GetValueByKey("$vtx")
	contentPart, _ := NewContentPartFromJSObject(node)
	if contentPart == nil {
		endNode := opts.renderBefore
		/*container, err := element.NewFromJSObject(hostNode.GetObjectValue())
		if err != nil {
			return nil, err
		}*/

		//if createMarker.GetObjectValue().IsNull() || createMarker.GetObjectValue().IsUndefined() {
		if createMarker == nil {
			doc, err := global.Document()
			if err != nil {
				console.Error(err)
				return nil, err
			}

			createMarker, err = doc.CreateComment("")
			if err != nil {
				console.Error(err)
				return nil, err
			}
		}

		cnode, err := hostNode.InsertBefore(createMarker, endNode)
		if err != nil {
			return nil, err
		}

		// 创建一个新的模板实例
		contentPart = NewContentPart(cnode, endNode, nil)
		// 保存到 hostNode 上
		hostNode.SetValueByKey("$vtx", contentPart.object.GetObjectValue())
	}

	contentPart.SetValue(value, opts)
	return contentPart, nil
}
