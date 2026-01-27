package vhtml

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/global"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
)

const (
	// Base62 字符集 (去掉了容易混淆的符号)
	vertexPrefix = "$vtx"
	letters      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
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
		Context context.Context
		typ     ResultType
		html    string //bytes.Buffer
		comp    Component
	}
)

func HTML(html string) *TemplateResult {
	return &TemplateResult{
		html: html,
		typ:  HTML_RESULT,
	}
}

func Render(value any, hostNode *node.Node, ctx context.Context) (*ContentPart, error) {
	nodeObj := hostNode.GetValueByKey(vertexPrefix)
	contentPart, _ := NewContentPartFromJSObject(nodeObj)
	if contentPart == nil {
		endNode, _ := ctx.Value("renderBefore").(*node.Node)
		/*container, err := element.NewFromJSObject(hostNode.GetObjectValue())
		if err != nil {
			return nil, err
		}*/

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
		contentPart = NewContentPart(&TemplatePart{}, cnode, nil, nil).(*ContentPart)
		contentPart.endNode = endNode
		// 保存到 hostNode 上
		hostNode.SetValueByKey(vertexPrefix, contentPart.object.GetObjectValue())
	}

	defer func() {
		if r := recover(); r != nil {
			js.RecoverHandler(r)
		}
	}()

	contentPart.SetValue(value, ctx)
	return contentPart, nil
}

func NewMarker() string {
	marker, _ := GenerateShortUID(8)
	return vertexPrefix + marker
}

func GenerateShortUID(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}
