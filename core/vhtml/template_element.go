package vhtml

import (
	"fmt"
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
	ATTRIBUTE_PART TemplatePartType = iota
	CHILD_PART
	PROPERTY_PART
	BOOLEAN_ATTRIBUTE_PART
	EVENT_PART
	ELEMENT_PART
	COMMENT_PART
	IF_PART
	ELSE_PART
	END_PART
)

type (
	TemplatePart struct {
		Type    TemplatePartType
		Index   int
		Name    string
		Strings []string
		Ctor    func(element *node.Node, name string, value any, strs []string, parent Disconnectable) IPart
	}

	TemplateElement struct {
		el    *htmltemplateelement.HTMLTemplateElement
		parts []*TemplatePart
	}
)

var walker *treewalker.TreeWalker
var createMarker *node.Node

func newTemplateElement(result *TemplateResult) (*TemplateElement, error) {
	tmpl := &TemplateElement{}
	html, expreParts := getTemplateHtml(result.html)

	var err error
	tmpl.el, err = tmpl.createElement(html, nil)
	if err != nil {
		return nil, err
	}

	documentfragment, err := tmpl.el.Content()
	if err != nil {
		return nil, err
	}
	fmt.Println("documentfragment:", documentfragment)
	//if walker.GetObjectValue().IsNull() && walker.GetObjectValue().IsUndefined() {
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
	fmt.Println("walker:", walker)
	var attrNameIndex = 0
	var nodeIdex = 0 //treewalker 上 node 的索引
	for {
		node, err := walker.NextNode()
		if err != nil {
			console.Error("walker.NextNode", err)
			break
		}

		fmt.Println("node:", node.GetObjectValue())
		if node.IsNull() {
			break
		}

		console.Info("Visiting node:", node.GetObjectValue())
		// Element 节点：检查带 marker 的属性
		element, err := element.NewFromJSObject(node.GetObjectValue())
		if err != nil {
			console.Error("element.NewFromJSObject", err)
			return nil, err
		}

		nodeType, err := node.NodeType()
		if err != nil {
			console.Error("element.NewFromJSObject", err)
			return nil, err
		}

		if nodeType == 1 {
			if attrs, err := element.Attributes(); err == nil && attrs.Length() > 0 {
				for i := 0; i < attrs.Length(); i++ {
					fmt.Println(attrs.Length(), i)
					attr, err := attrs.Item(i)
					if err != nil {
						console.Error("attrs.Item(i)", err)
					}

					name, err := attr.Name()
					if err != nil {
						console.Error("attrs.Name()", err)
					}

					if strings.HasSuffix(name, boundAttributePrefix) {
						value, err := element.GetAttribute(name)
						if err != nil {
							console.Error("GetAttribute()", err)
						}

						attrNameIndex++
						expr := expreParts[attrNameIndex]

						part := &TemplatePart{
							Index:   nodeIdex,
							Type:    expr.Type,
							Strings: strings.Split(value, boundAttributePrefix),
						}

						// part 的构造函数
						switch expr.Type {
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
					} else if name == marker {
						tmpl.parts = append(tmpl.parts, &TemplatePart{
							Index: nodeIdex,
							Type:  ELEMENT_PART,
						})
						element.RemoveAttribute(marker)
					}
				}
				fmt.Println("localName error:", element.GetObjectValue(), err)

				localName, err := element.LocalName()
				if err != nil {
					fmt.Println("localName error:", element, err)
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

			if v == marker {
				// Comment 节点：检查是否为 marker comment
				tmpl.parts = append(tmpl.parts, &TemplatePart{
					Index: nodeIdex,
					Type:  CHILD_PART,
				})
			} else {
				// Comment 节点：检查是否为 marker comment
				idx := -1
				for {
					if idx = strings.Index(v[idx+1:], marker); idx != -1 {
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

func (self *TemplateElement) createElement(html string, options *RenderOptions) (*htmltemplateelement.HTMLTemplateElement, error) {
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
