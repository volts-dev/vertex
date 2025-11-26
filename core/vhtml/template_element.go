package vhtml

import (
	"context"
	"strings"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmltemplateelement"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/window"
)

type (
	TemplateElement struct {
		el    htmltemplateelement.HTMLTemplateElement
		parts []TemplatePart
	}

	// TemplateInstance 执行模板更新对象
	TemplateInstance struct {
		template *TemplateElement
		parent   *ContentPart
		parts    []IPart
	}
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
)

func NewTemplateInstance(tmpl *TemplateElement, parent *ContentPart) *TemplateInstance {
	return &TemplateInstance{
		template: tmpl,
		parent:   parent,
	}
}

// 从 template 克隆一个新的实例
func (t *TemplateInstance) CloneTemplate() *node.Node {
	clonedEl, _ := t.template.el.CloneNode(true)
	return &clonedEl
}

// 更新克隆的实例
func (self *TemplateInstance) Update(ctx ...context.Context) error {
	return nil //self.template.Update(ctx...)
}

func newTemplateElement(result *TemplateResult) (*TemplateElement, error) {
	tmpl := &TemplateElement{}
	doc, err := window.Default().Document()
	if err != nil {
		return nil, err
	}

	// 创建模板元素
	tmpl.el, err = htmltemplateelement.New(doc)
	if err != nil {
		return nil, err
	}

	// 解析模板
	//result := HTML(html)

	// 设置模板内容
	if Render(result, tmpl.el.Node) != nil {
		return nil, err
	}

	documentfragment, err := tmpl.el.Content()
	if err != nil {
		return nil, err
	}

	treewalker, err := doc.CreateTreeWalker(documentfragment.Node)
	if err != nil {
		return nil, err
	}

	var attrNameIndex = 0
	var nodeIdex = 0
	for {
		node, err := treewalker.NextNode()
		if err != nil {
			break
		}

		nodeType, _ := node.NodeType()
		if nodeType == 1 {
			// Element 节点：检查带 marker 的属性
			element, err := element.NewFromJSObject(node.GetObjectValue())
			if err != nil {
				break
			}

			if attrs, err := element.Attributes(); err == nil && attrs.Length() > 0 {
				for i := 0; i < attrs.Length(); i++ {
					attr, err := attrs.Item(i)
					if err != nil {
						console.Error(err)
					}

					name, err := attr.Name()
					if err != nil {
						console.Error(err)
					}

					if strings.HasPrefix(name, boundAttributePrefix) {
						value, err := element.GetAttribute(name)
						if err != nil {
							console.Error(err)
						}

						attrNameIndex++
						realName := result.attributeNames[attrNameIndex]

						part := TemplatePart{
							Index:   nodeIdex,
							Type:    ATTRIBUTE_PART,
							Strings: strings.Split(value, boundAttributePrefix),
						}

						// part 的构造函数
						switch realName[0] {
						case '.':
							part.Ctor = NewPropertyPart
						case '?':
							part.Ctor = NewBooleanAttributePart
						case '@':
							part.Ctor = NewEventPart
						default:
							part.Ctor = NewAttributePart
						}

						tmpl.parts = append(tmpl.parts, part)
						element.RemoveAttribute(name)
					}
				}
			}

		} else if nodeType == 8 {

			node.GetValueByKey("data")
			// Comment 节点：检查是否为 marker comment
			tmpl.parts = append(tmpl.parts, TemplatePart{
				Index:   nodeIdex,
				Type:    ATTRIBUTE_PART,
				Strings: strings.Split(value, boundAttributePrefix),
			})

		}

		nodeIdex++
	}
	return tmpl, nil
}
