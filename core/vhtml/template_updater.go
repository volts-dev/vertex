package vhtml

import (
	"context"
	"fmt"
	"reflect"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/global"
	"github.com/volts-dev/vertex/html/node"
)

type (
	// TemplateInstance 执行模板更新对象
	TemplateInstance struct {
		template  *TemplateElement
		parent    *ContentPart
		parts     []IPart
		compValue reflect.Value // 每个handler的控制器必须是唯一的
		comp      interface{}   // 提供Ctx特殊调用
	}
)

func NewTemplateInstance(tmpl *TemplateElement, parent *ContentPart, comp interface{}) *TemplateInstance {
	return &TemplateInstance{
		template:  tmpl,
		parent:    parent,
		compValue: reflect.ValueOf(comp),
		comp:      comp,
	}
}

// 从 template 克隆一个新的实例
func (self *TemplateInstance) CloneTemplate() (*node.Node, error) {
	fmt.Println(self.template)
	node, _ := self.template.el.Content()
	doc, err := global.Document()
	if err != nil {
		return nil, err
	}

	fragment, err := doc.ImportNode(node.Node, true)
	if err != nil {
		return nil, err
	}

	if walker.GetObjectValue().IsNull() && walker.GetObjectValue().IsUndefined() {
		walker, err = doc.CreateTreeWalker(doc.Node)
		if err != nil {

		}
	}

	err = walker.SetCurrentNode(fragment.GetObjectValue())
	if err != nil {
		return nil, err
	}

	nnode, err := walker.NextNode()
	nodeIndex := 0
	var templatePart *TemplatePart
	compValue := self.compValue.Elem()

	for partIndex := 0; partIndex < len(self.template.parts); partIndex++ {
		templatePart = self.template.parts[partIndex]

		if templatePart.Index != nodeIndex {
			nnode, err = walker.NextNode()
			nodeIndex++
		}

		if nodeIndex == templatePart.Index {
			var part IPart

			switch templatePart.Type {
			case ATTRIBUTE_PART, BOOLEAN_ATTRIBUTE_PART, PROPERTY_PART:
				f := compValue.FieldByName(templatePart.Name)
				part = templatePart.Ctor(nnode, templatePart.Name, f, templatePart.Strings, self)
			case EVENT_PART:
				m := compValue.MethodByName(templatePart.Name)
				part = NewEventPart(nnode, templatePart.Name, m, templatePart.Strings, self)
			case CHILD_PART:
				endNode, err := nnode.NextSibling()
				if err != nil {
					return nil, err
				}
				part = NewContentPart(nnode, endNode, self)
			case ELEMENT_PART:
				part = NewElementPart(nnode, self)

			}

			self.parts = append(self.parts, part)
		}
	}

	if err = walker.SetCurrentNode(doc.GetObjectValue()); err != nil {
		return nil, err
	}

	return fragment, nil
}

// 更新克隆的实例
func (self *TemplateInstance) Update(ctx ...context.Context) {
	for _, part := range self.parts {
		err := part.SetValue(nil, nil)
		if err != nil {
			console.Error(err)
		}
	}
}
