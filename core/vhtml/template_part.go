package vhtml

import (
	//"github.com/volts-dev/cacher"
	//_ "github.com/volts-dev/cacher/memory"

	"context"
	"fmt"
	"reflect"

	"github.com/expr-lang/expr"
	"github.com/volts-dev/vertex/core/cacher"
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/text"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

type (
	IPart interface {
		SetValue(value any, ctx context.Context) error
		commitValue(value any) error
		//commitNode(value *node.Node) error
	}

	Disconnectable interface {
		//Disconnect()
	}

	Part struct {
		super          IPart
		parent         *TemplateInstance
		element        *node.Node
		startNode      *node.Node
		endNode        *node.Node
		name           string
		value          any // reflect.Value
		com            reflect.Value
		strings        []string
		committedValue any
	}

	AttributePart struct {
		Part
	}

	BooleanAttributePart struct {
		AttributePart
	}

	PropertyPart struct {
		AttributePart
	}

	EventPart struct {
		AttributePart
		method js.Func
	}

	ElementPart struct {
		parent *TemplateInstance
		tag    string
		name   string
		raw    bool
	}

	// 非属性节点部分，比如文本节点、注释节点等
	ContentPart struct {
		Part
		object   js.Object // 废弃
		Instance *TemplateInstance
	}
)

// WrapNode 包装一个 DOM 节点，如果启用了 ShadyDOM noPatch 模式则使用 ShadyDOM.wrap，
// 否则直接返回节点本身。
//
// 这个函数用于处理 ShadyDOM 的兼容性，确保在使用 ShadyDOM 时能够正确地操作 DOM 节点。
func WrapNode(n *node.Node) *node.Node {
	// 检查三个条件：
	// 1. ENABLE_SHADYDOM_NOPATCH - 是否启用 ShadyDOM noPatch 模式
	// 2. global.ShadyDOM 是否存在
	// 3. global.ShadyDOM.noPatch 是否为 true
	if shadyDOM := js.Global().Get("ShadyDOM"); !shadyDOM.IsNull() && !shadyDOM.IsUndefined() {
		console.Info("ShadyDOM")
		if enableShadydomNoPatch, err := shadyDOM.Get("noPatch").Bool(); enableShadydomNoPatch && err != nil {
			// 如果所有条件都满足，使用 ShadyDOM 的 wrap 方法
			if v := shadyDOM.Call("wrap", n.GetObjectValue()); v.Error() == nil {
				vv, err := node.NewFromJSObject(v)
				if err != nil {
					return n
				}
				return vv
			}
		}
	}

	// 否则直接返回节点
	return n
}

func NewContentPart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	// 获取 endNode：优先从当前元素的下一个兄弟节点获取，否则继承父容器的 endNode
	endNode, _ := element.NextSibling()
	if endNode == nil && instance != nil {
		if p, ok := instance.parent.(*ContentPart); ok && p.endNode != nil {
			endNode = p.endNode
		}
	}

	p := &ContentPart{
		Part: Part{
			parent:    instance,
			name:      part.Name,
			value:     part,
			startNode: element,
			endNode:   endNode,
			strings:   part.Strings,
		},
	}

	p.super = p
	return p
}

func NewAttributePart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	p := &AttributePart{
		Part: Part{
			element: element,
			name:    part.Name,
			value:   part.Value,
			parent:  instance,
			strings: part.Strings,
		},
	}

	if com != nil {
		if f := com.FieldByName(part.Name); f.IsValid() {
			p.value = f
		}
	}

	p.super = p
	return p
}

func NewBooleanAttributePart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	p := &BooleanAttributePart{
		AttributePart: AttributePart{
			Part: Part{
				element: element,
				name:    part.Name,
				value:   part.Value,
				parent:  instance,
				strings: part.Strings,
			},
		},
	}

	if com != nil {
		if f := com.FieldByName(part.Name); f.IsValid() {
			p.value = f
		}
	}

	p.super = p
	return p
}

func NewPropertyPart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	p := &PropertyPart{
		AttributePart: AttributePart{
			Part: Part{
				element: element,
				name:    part.Name,
				value:   part.Value,
				parent:  instance,
				strings: part.Strings,
			},
		},
	}

	if com != nil {
		if f := com.FieldByName(part.Name); f.IsValid() {
			p.value = f
		}
	}

	p.super = p
	return p
}

func NewEventPart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	p := &EventPart{
		AttributePart: AttributePart{
			Part: Part{
				element: element,
				name:    part.Name,
				value:   part.Value,
				parent:  instance,
				strings: part.Strings,
			},
		},
	}

	if com != nil {
		if m := com.MethodByName(part.Name); m.IsValid() {
			p.value = m
		}
	}

	p.super = p
	return p
}

func NewElementPart(part *TemplatePart, element *node.Node, com *reflect.Value, instance *TemplateInstance) IPart {
	return &ElementPart{
		parent: instance,
		//tag:     element.NodeName(),
	}
}

func NewContentPartFromJSObject(obj js.Value) (*ContentPart, error) {
	object, err := js.ToObject(obj)
	if err != nil {
		return nil, err
	}

	part := &ContentPart{
		object: object,
	}

	part.super = part
	return part, nil
}

func (self *Part) SetValue(value any, ctx context.Context) error {
	return nil
}

func (self *Part) commitValue(value any) error {
	return nil
}

func (self *Part) getTemplateElement(content string, parts []*TemplatePart, ctx context.Context) *TemplateElement {
	tmplCacher := cacher.Default()
	value, err := tmplCacher.Get(content)
	if value != nil {
		console.Info("Template cache hit", value.(*TemplateElement).el.Node.GetObjectValue())
		return value.(*TemplateElement)
	}

	tmpl, err := newTemplateElement(content, parts, ctx)
	if err != nil {
		panic(err)
	}

	if err = tmplCacher.Set(&cacher.CacheBlock{Key: content, Value: tmpl}); err != nil {
		console.Error(err)
	}

	return tmpl
}

func (self *Part) insert(target *node.Node) *node.Node {
	parentNode, err := WrapNode(self.startNode).ParentNode()
	if err != nil {
		console.Error("Failed to get parent node:", err)
		return nil
	}

	if parentNode == nil {
		console.Error("Parent node is nil")
		return nil
	}

	/*
		// 验证 endNode 是否是父节点的子节点
		// 如果 endNode 为 nil，insertBefore 会将节点追加到末尾
		// 如果 endNode 不是父节点的子节点，需要修正引用
		var refNode *node.Node
		if self.endNode != nil {
			console.Error("insert parent:", self.endNode.GetObjectValue())

			// 检查 endNode 是否真的是 parentNode 的子节点
			isChild, _ := self.isChildOf(self.endNode, parentNode)
			if isChild {
				refNode = self.endNode
			} else {
				// endNode 不在 parentNode 中，尝试找到正确的参考节点
				console.Warn("endNode is not a child of parent, trying to find correct reference node")
				// 如果找不到，传 nil 让 insertBefore 追加到末尾
				refNode = nil
			}
		}*/
	insertedNode, err := WrapNode(parentNode).InsertBefore(target, self.endNode)
	if err != nil {
		console.Error("Failed to insert node:", err)
		return nil
	}

	return insertedNode
}

// isChildOf 检查 child 是否是 parent 的直接或间接子节点
func (self *Part) isChildOf(child, parent *node.Node) (bool, error) {
	if child == nil || parent == nil {
		return false, nil
	}

	current := child
	for current != nil {
		parent, err := WrapNode(current).ParentNode()
		if err != nil {
			return false, err
		}

		if parent == nil {
			return false, nil
		}

		if parent.Equal(parent) {
			return true, nil
		}

		current = parent
	}

	return false, nil
}

func (self *Part) clear() {
	if self.startNode == nil {
		return
	}

	startNode := WrapNode(self.startNode)
	nextNode, err := startNode.NextSibling()
	if err != nil {
		//console.Error("Failed to get next sibling:", err)
		return
	}

	/*
		n, _ := startNode.ParentNode()
		tt, _ := element.NewFromJSObject(n.GetObjectValue())
		str, _ := tt.OuterHTML()
		console.Info("clear1", self.endNode == nil, str)
	*/

	// 清除 startNode 和 endNode 之间的所有节点
	for nextNode != nil {
		// 检查是否到达了 endNode
		if self.endNode != nil && nextNode.Equal(self.endNode) {
			//console.Error("endNode Equal:", self.endNode.GetObjectValue(), nextNode.GetObjectValue())
			break
		}

		// 保存下一个节点的引用，因为删除后就无法获取
		tempNext, err := nextNode.NextSibling()
		if err != nil {
			//console.Error("Failed to get next sibling:", err)
		}

		// 删除当前节点
		if err := WrapNode(nextNode).Remove(); err != nil {
			console.Error("Failed to remove node:", err)
		}

		nextNode = tempNext
	}
}

func (self *ContentPart) SetValue(value any, ctx context.Context) error {
	if value == nil {
		value = self.value
	}

	part, ok := self.super.(*ContentPart)
	if !ok {
		return fmt.Errorf("")
	}

	switch v := value.(type) {
	case *TemplateResult:
		v.Context = ctx
		// 清空当前内容
		tmplElement := self.getTemplateElement(v.html, nil, ctx)
		//if instance, ok := self.committedValue.(*TemplateInstance); ok {
		if part.Instance != nil {
			if part.Instance.template != tmplElement {
				part.Instance.Update()
			}
		} else {
			part.Instance = NewTemplateInstance(tmplElement, part, ctx)
			fragment, err := part.Instance.CloneTemplate()
			if err != nil {
				return err
			}

			if err = part.commitNode(fragment); err != nil {
				return err
			}

			part.Instance.Update()
		}

	case *TemplatePart:
		switch v.Type {
		case CHILD_IF_PART, CHILD_ELSE_PART:
			//instance, ok := self.committedValue.(*TemplateInstance)
			if part.Instance == nil {
				tmplElement := self.getTemplateElement(v.Value, v.Children, ctx)
				instance := NewTemplateInstance(tmplElement, part, ctx)
				/*	part.Instance.fragment, err = part.Instance.CloneTemplate()
					if err != nil {
						console.Error("CloneTemplate error:", err)
						return err
					}*/
				//console.Info("Showing part:", instance.fragment.GetObjectValue())
				//part.Instance.Value = false // 初始值为 false
				//self.committedValue = instance // 固定存储instance
				env := make(map[string]any)
				for _, str := range part.strings {
					val := ctx.Value("." + str)
					if v, ok := val.(reflect.Value); ok {
						val = v.Interface()
					}
					env[str] = val
				}
				fmt.Println(env)

				program, err := expr.Compile(part.name, expr.AsBool(), expr.Env(env))
				if err != nil {
					panic(err)
				}

				instance.env = env
				instance.conditon = program

				part.Instance = instance
			}

			// 优化一些
			for _, str := range part.strings {
				val := ctx.Value("." + str)
				if v, ok := val.(reflect.Value); ok {
					val = v.Interface()
				}
				part.Instance.env[str] = val
			}

			output, err := expr.Run(part.Instance.conditon, part.Instance.env)
			if err != nil {
				panic(err)
			}

			toShow := output.(bool)
			/*
				switch v := ctx.Value(part.name).(type) {
				case bool:
					toShow = v
				case reflect.Value:
					toShow = v.Bool()
				}

					toShow, ok := ctx.Value(self.name).(bool)
					if !ok {
						toShow = false

						v, ok := ctx.Value(self.name).(reflect.Value)
						if ok {
							toShow = v.Bool()
						}
					}
			*/

			if v.Type == CHILD_ELSE_PART {
				toShow = !toShow
			}

			//fmt.Println("PART:", v.Type, "toShow:", toShow)

			// 获取当前显示状态
			wasShowing := false
			if show, ok := part.Instance.Value.(bool); ok {
				wasShowing = show
			}

			if toShow {
				if !wasShowing {
					fragment, err := part.Instance.CloneTemplate()
					if err != nil {
						console.Error("CloneTemplate error:", err)
						return err
					}

					// 从隐藏状态切换到显示状态，重新插入节点

					part.clear()
					//part.committedValue = part.insert(fragment)
					part.commitNode(fragment)
				}

				// 更新内容
				if err := part.Instance.Update(ctx); err != nil {
					console.Error("Update error:", err)
					return err
				}

			} else {
				if wasShowing {
					// 从显示状态切换到隐藏状态，清除内容
					part.clear()
					part.committedValue = nil
				}
			}

			// 更新显示状态
			part.Instance.Value = toShow

		case CHILD_VARIANT_PART:
			var newValue any = "???"
			if vv, ok := ctx.Value("." + part.name).(reflect.Value); ok {
				newValue = vv.Interface()
			}

			var change, noCommit bool
			change = change || part.committedValue != newValue
			if change {
				part.committedValue = newValue
			}

			if change && !noCommit {
				part.commitValue(newValue)
			}
		}

	case *node.Node:
		return part.commitNode(v)
	default:
		return part.commitValue(v)
	}

	return nil
}

func (self *ContentPart) commitNode(value *node.Node) error {
	oldValue, ok := self.committedValue.(*node.Node)
	if !ok || (oldValue == (*node.Node)(nil)) {
		oldValue = nil
	}

	if oldValue == nil || !oldValue.Equal(value) {
		self.clear()
		self.committedValue = self.insert(value)
	}

	return nil
}

// commitValue as commitText
func (self *ContentPart) commitValue(value any) error {
	if old, ok := self.committedValue.(js.Value); ok {
		if js.IsPrimitive(old) {
			node, err := WrapNode(self.startNode).NextSibling()
			if err != nil {
				console.Error(err)
			}

			if text, err := text.NewFromJSObject(node.GetObjectValue()); err == nil {
				text.SetText(value.(string))
			}
		}
	} else {
		doc, err := window.Default().Document()
		if err != nil {
			return err
		}

		textNode, _ := doc.CreateTextNode("")
		//textNode.SetTextContent(utils.ToString(value))
		textNode.SetNodeValue(value)
		self.commitNode(textNode)
	}

	self.committedValue = value
	return nil
}

func (self *AttributePart) SetValue(value any, ctx context.Context) error {
	if value == nil {
		value = self.value
	}

	var change, noCommit bool
	if len(self.strings) == 0 {
		// Single-value binding case
		change = change || self.committedValue != value
		if change {
			self.committedValue = value
		}

	} else {
		// Interpolation case
		newValues, ok := value.([]any)
		if !ok {

		}

		values, ok := self.committedValue.([]any)
		if !ok {

		}

		newValue := self.strings[0]

		for i := 1; i < len(self.strings); i++ {
			v := newValues[i]

			change = change || values[i] != v
			if value != "" {
				newValue += v.(string) + self.strings[i]
			}

			// We always record each value, even if one is `nothing`, for future
			// change detection.
			values[i] = newValue
		}

		value = values
		self.committedValue = values
	}

	if change && !noCommit {
		return self.super.commitValue(value)
	}

	return nil
}

func (self *AttributePart) commitValue(value any) error {
	if value == nil {
		ele, _ := element.NewFromJSObject(self.element.GetObjectValue())
		return ele.RemoveAttribute(self.name)
	}

	return self.element.SetAttribute(self.name, value)
}

func (self *BooleanAttributePart) commitValue(value any) error {
	ele, _ := element.NewFromJSObject(self.element.GetObjectValue())
	v, _ := value.(bool)
	_, err := ele.ToggleAttribute(self.name, v)
	return err
}

func (self *PropertyPart) commitValue(value any) error {
	if value == nil {
		return self.element.SetAttribute(self.name, js.Undefined())

	}

	return self.element.SetAttribute(self.name, value)
}

func (self *EventPart) SetValue(value any, ctx context.Context) error {
	if methodName, ok := self.committedValue.(string); ok && methodName == self.value {
		return nil
	}

	if value == nil {
		//if m, ok := ctx.Value(self.name).(reflect.Method); ok {
		//	value = m.Func.Interface()
		//}
		value = ctx.Value(self.value)
	}

	// If the new value is nothing or any options change we have to remove the
	// part as a listener.
	shouldRemoveListener := (value == nil && self.committedValue == nil)
	/*
		(newListener === nothing && oldListener !== nothing) ||
		(newListener as EventListenerWithOptions).capture !==
			(oldListener as EventListenerWithOptions).capture ||
		(newListener as EventListenerWithOptions).once !==
			(oldListener as EventListenerWithOptions).once ||
		(newListener as EventListenerWithOptions).passive !==
			(oldListener as EventListenerWithOptions).passive;
	*/

	// If the new value is not nothing and we removed the listener, we have
	// to add the part as a listener.
	shouldAddListener := value != nil &&
		(self.committedValue == nil || shouldRemoveListener)

	console.Info("EventPart.SetValue", "eventName:", self.name, "shouldAdd:", shouldAddListener, "shouldRemove:", shouldRemoveListener)

	ele, err := element.NewFromJSObject(self.element.GetObjectValue())
	if err != nil {
		console.Error("Failed to get element:", err)
		return err
	}

	if shouldRemoveListener && self.method != nil {
		if err := ele.RemoveEventListenerWithFunc(self.name, self.method); err != nil {
			console.Error("Failed to remove event listener:", err)
		}
		self.method = nil
	}

	if shouldAddListener {
		var newListener func(e event.Event) error

		// Try to extract the handler from reflect.Value
		if v, ok := value.(reflect.Value); ok && v.IsValid() {
			if handler, ok := v.Interface().(func(e event.Event) error); ok {
				newListener = handler
			}
		} else if handler, ok := value.(func(e event.Event) error); ok {
			// Direct function type
			newListener = handler
		}

		if newListener == nil {
			console.Warn("EventPart: handler is not a valid function for event:", self.name)
			return nil
		}

		cb, err := ele.AddEventListener(self.name, newListener)
		if err != nil {
			console.Error("Failed to add event listener:", err)
			return err
		}

		self.method = cb
		console.Info("Event listener added for:", self.name)
	}

	self.committedValue = self.value
	return nil
}

func (self *ElementPart) SetValue(value any, ctx context.Context) error {
	return nil
}

func (self *ElementPart) commitValue(value any) error {
	return nil
}

func (self *ElementPart) commitNode(value *node.Node) error {
	return nil
}
