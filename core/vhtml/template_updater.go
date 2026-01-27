package vhtml

import (
	"context"
	"fmt"
	"reflect"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/global"
	"github.com/volts-dev/vertex/html/node"

	"github.com/expr-lang/expr/vm"
)

type (
	// TemplateInstance 执行模板更新对象
	TemplateInstance struct {
		template  *TemplateElement
		parent    IPart
		fragment  *node.Node // 克隆的节点片段
		conditon  *vm.Program
		env       map[string]any
		parts     []IPart       // 模板中的所有更新部分
		Value     any           // 当前 contentpart 的值
		compValue reflect.Value // 每个handler的控制器必须是唯一的
		comp      interface{}   // 提供Ctx特殊调用
		context   context.Context
	}
)

func NewTemplateInstance(tmpl *TemplateElement, parent IPart, ctx context.Context) *TemplateInstance {
	com := ctx.Value("component")
	value := reflect.ValueOf(com)

	// 获取所有方法信息
	// 注意：reflect.TypeOf(s) 如果 s 是值，只能取到值接收者方法；
	// 如果 s 是指针，能取到值接收者 + 指针接收者方法。
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		method := value.Method(i)
		ctx = context.WithValue(ctx, typ.Method(i).Name, method)
		//fmt.Println("NumMethod", pt.Method(i).Name, method.IsValid())
	}

	// 获取所有字段信息时使用非指针
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typ = value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		ctx = context.WithValue(ctx, "."+typ.Field(i).Name, field)
		//console.Info("Field:", "."+typ.Field(i).Name)
	}

	return &TemplateInstance{
		template:  tmpl,
		parent:    parent,
		Value:     false,
		compValue: value.Addr(),
		comp:      com,
		context:   ctx,
	}
}

// 从 template 克隆一个新的实例
func (self *TemplateInstance) CloneTemplate() (*node.Node, error) {
	tmpNode, _ := self.template.el.Content()
	doc, err := global.Document()
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}

	fragment, err := doc.ImportNode(tmpNode.Node, true)
	if err != nil {
		return nil, fmt.Errorf("failed to import node: %w", err)
	}

	// 初始化树遍历器
	if walker == nil {
		if walker, err = doc.CreateTreeWalker(doc.Node); err != nil {
			return nil, fmt.Errorf("failed to create tree walker: %w", err)
		}
	}

	if err = walker.SetCurrentNode(fragment.GetObjectValue()); err != nil {
		return nil, fmt.Errorf("failed to set current node: %w", err)
	}

	defer func() {
		if err = walker.SetCurrentNode(doc.GetObjectValue()); err != nil {
			console.Error("failed to reset walker: %w", err)
		}
	}()

	//console.Info("fragment:", node.GetObjectValue(), fragment.GetObjectValue())
	/*
		nnode, err := walker.NextNode()
		if err != nil {
			//console.Info("walker.NextNode", fragment.GetObjectValue(), err)
			return fragment, nil
		}
	*/
	var nnode *node.Node
	nodeIndex := -1
	for _, templatePart := range self.template.parts {
		for templatePart.Index != nodeIndex {
			nnode, err = walker.NextNode()
			if err != nil {
				goto DONE
			}

			nodeIndex++
		}

		if templatePart.Ctor == nil {
			fmt.Println("Ctor found for part:", templatePart.Name)
			panic("Miss templatePart.Ctor implementation")
		}

		self.parts = append(self.parts, templatePart.Ctor(templatePart, nnode, &self.compValue, self))
	}

DONE:
	return fragment, nil
}

// 更新克隆的实例
func (self *TemplateInstance) Update(ctx ...context.Context) error {
	for _, part := range self.parts {
		if err := part.SetValue(nil, self.context); err != nil {
			console.Error("Failed to update part: %v", err)
			//panic(err)
			return err
		}
	}

	return nil
}
