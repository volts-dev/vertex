package vhtml

import (
	"context"
	"reflect"
	"strings"

	"github.com/volts-dev/utils"
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/global"
	"github.com/volts-dev/vertex/html/htmltemplateelement"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/treewalker"
	"github.com/volts-dev/vertex/html/window"
)

const (
	// 模板部分类型
	CHILD_PART TemplatePartType = iota //废弃
	CHILD_VARIANT_PART
	CHILD_FUNC_PART //废弃
	CHILD_IF_PART
	CHILD_ELSE_PART
	CHILD_END_PART //废弃
	ATTRIBUTE_PART
	PROPERTY_PART
	BOOLEAN_ATTRIBUTE_PART
	EVENT_PART
	ELEMENT_PART
	COMMENT_PART
)

type (
	TemplatePart struct {
		Type     TemplatePartType
		Index    int // node index in template
		Name     string
		Value    string
		Strings  []string
		Children []*TemplatePart
		Ctor     func(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart
	}

	TemplateElement struct {
		el    *htmltemplateelement.HTMLTemplateElement
		parts []*TemplatePart
	}
)

var walker *treewalker.TreeWalker
var createMarker *node.Node

// func newTemplateElement(result *TemplateResult) (*TemplateElement, error) {
func newTemplateElement(content string, parts []*TemplatePart, ctx context.Context) (*TemplateElement, error) {
	if len(parts) == 0 {
		content, parts = parseTemplateHtml(content, ctx)
	}
	tmpl := &TemplateElement{}

	var err error
	tmpl.el, err = tmpl.createElement(content)
	if err != nil {
		return nil, err
	}

	documentfragment, err := tmpl.el.Content()
	if err != nil {
		return nil, err
	}

	if walker == nil {
		doc, err := global.Document()
		if err != nil {
			return nil, err
		}

		walker, err = doc.CreateTreeWalker(doc.Node)
		if err != nil {
			console.Error(err)
			return nil, err
		}
	}

	//treewalker, err := doc.CreateTreeWalker(documentfragment.Node)
	err = walker.SetCurrentNode(documentfragment.GetObjectValue())
	if err != nil {
		return nil, err
	}

	var partIndex = 0
	var nodeIdex = 0 //treewalker 上 node 的索引
	marker := ctx.Value("_vertex_marker").(string)

	for {
		node, err := walker.NextNode()
		if err != nil {
			//console.Error("walker.NextNode", err)
			break
		}

		if node.IsNull() {
			break
		}

		nodeValue := node.GetObjectValue()
		nodeType, err := node.NodeType()
		if err != nil {
			console.Error("element.NewFromJSObject", err)
			return nil, err
		}

		if nodeType == 1 {
			// Element 节点：检查带 marker 的属性
			element, err := element.NewFromJSObject(nodeValue)
			if err != nil {
				console.Error(" NewFromJSObject", nodeValue, err)
				return nil, err
			}

			if attrs, err := element.Attributes(); err == nil && attrs.Length() > 0 {
				for i := 0; i < attrs.Length(); i++ {
					attr, err := attrs.Item(i)
					if err != nil {
						console.Error("attrs.Item(i)", err)
					}

					name, err := attr.Name()
					if err != nil {
						console.Error("attrs.Name()", err)
					}

					if strings.HasSuffix(name, vertexPrefix) {
						value, err := element.GetAttribute(name)
						if err != nil {
							console.Error("GetAttribute()", err)
						}

						part := parts[partIndex]
						part.Index = nodeIdex
						part.Strings = strings.Split(value, vertexPrefix)
						/*
							part := &TemplatePart{
								Index:   nodeIdex,
								Type:    expr.Type,
								Strings: strings.Split(value, vertexPrefix),
							}*/

						// part 的构造函数
						switch part.Type {
						case PROPERTY_PART:
							part.Ctor = NewPropertyPart
						case BOOLEAN_ATTRIBUTE_PART:
							part.Ctor = NewBooleanAttributePart
						case EVENT_PART:
							part.Ctor = NewEventPart
						default:
							part.Ctor = NewAttributePart
						}

						tmpl.parts = append(tmpl.parts, part)
						element.RemoveAttribute(name)
						partIndex++
					} else if name == marker {
						part := parts[partIndex]
						part.Index = nodeIdex
						part.Type = ELEMENT_PART
						part.Ctor = NewElementPart

						tmpl.parts = append(tmpl.parts, part)
						element.RemoveAttribute(marker)
						partIndex++
					}
				}

				localName, err := element.LocalName()
				if err != nil {
					console.Error("localName error:", err)
				}

				if utils.IndexOf(localName, "script", "style", "textarea", "title") != -1 {
					// TODO:处理原始文本元素（script/style/textarea/title）
				}
			}

		} else if nodeType == 8 {
			// 获取 comment 节点内容
			v, err := node.GetValueByKey("data").String()
			if err != nil {
				console.Error("GetValueByKey", err)
			}

			if strings.HasSuffix(v, marker) {
				part := parts[partIndex]
				part.Ctor = NewContentPart
				part.Index = nodeIdex
				tmpl.parts = append(tmpl.parts, part)
				partIndex++

			} else {
				// Comment 节点：检查是否为 marker comment
				idx := -1
				for {
					if idx = strings.Index(v[idx+1:], marker); idx != -1 {
						///partIndex++
						// Comment 节点：检查是否为 marker comment
						tmpl.parts = append(tmpl.parts, &TemplatePart{
							Index: nodeIdex,
							Type:  COMMENT_PART,
						})
						// 移动索引
						idx += len(marker) - 1
						continue
					}
					break
				}
			}
		}

		nodeIdex++
	}

	return tmpl, nil
}

func (self *TemplateElement) createElement(html string) (*htmltemplateelement.HTMLTemplateElement, error) {
	doc, err := window.Default().Document()
	if err != nil {
		return nil, err
	}

	// 创建模板元素
	tmpl, err := htmltemplateelement.New(doc)
	if err != nil {
		return nil, err
	}

	if err = tmpl.SetInnerHTML(html); err != nil {
		return nil, err
	}

	return &tmpl, nil
}
