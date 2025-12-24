package vhtml

import (
	//"github.com/volts-dev/cacher"
	//_ "github.com/volts-dev/cacher/memory"

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
		SetValue(value any, opts *RenderOptions) error
	}

	Disconnectable interface {
		//Disconnect()
	}

	AttributePart struct {
		element        *node.Node
		parent         Disconnectable
		name           string
		value          any
		committedValue any
		strings        []string // 多值 比如 class
	}

	EventPart struct {
		AttributePart
		method string
	}

	BooleanAttributePart struct {
		AttributePart
	}

	PropertyPart struct {
		AttributePart
	}

	ElementPart struct {
		parent Disconnectable
		tag    string
		name   string
		raw    bool
	}

	ContentPart struct {
		object         js.Object
		parent         Disconnectable
		startNode      *node.Node
		endNode        *node.Node
		committedValue any
	}
)

//var tmplCacher cacher.ICacher

func init() {
	//tmplCacher, _ = cacher.New("memory")
}

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
		if enableShadydomNoPatch, err := shadyDOM.Get("noPatch").Bool(); enableShadydomNoPatch && err != nil {
			// 如果所有条件都满足，使用 ShadyDOM 的 wrap 方法
			if v := shadyDOM.Call("wrap", n.GetObjectValue()); v.Error() == nil {
				vv, err := node.NewFromJSObject(v)
				if err != nil {
					//return n
				}
				return vv
			}
		}
	}

	// 否则直接返回节点
	return n
}

func NewContentPart(startNode, endNode *node.Node, parent Disconnectable) *ContentPart {
	obj := js.GetObjectInterface().New()
	o, err := js.ToObject(obj)
	if err != nil {
		return nil
	}
	return &ContentPart{
		parent:    parent,
		object:    o,
		startNode: startNode,
		endNode:   endNode,
	}
}

func NewAttributePart(element *node.Node, name string, value any, strs []string, parent Disconnectable) IPart {
	return &AttributePart{
		element: element,
		name:    name,
		value:   value,
		parent:  parent,
	}
}

func NewBooleanAttributePart(element *node.Node, name string, value any, strs []string, parent Disconnectable) IPart {
	return &BooleanAttributePart{
		AttributePart: AttributePart{
			element: element,
			name:    name,
			value:   value,
			parent:  parent,
		},
	}
}

func NewEventPart(element *node.Node, name string, value any, strs []string, parent Disconnectable) IPart {
	return &EventPart{
		AttributePart: AttributePart{
			element: element,
			name:    name,
			value:   value,
			parent:  parent,
		},
	}
}

func NewPropertyPart(element *node.Node, name string, value any, strs []string, parent Disconnectable) IPart {
	return &PropertyPart{
		AttributePart: AttributePart{
			element: element,
			name:    name,
			value:   value,
			parent:  parent,
		},
	}
}

func NewElementPart(element *node.Node, parent Disconnectable) IPart {
	return &ElementPart{
		parent: parent,
		//tag:     element.NodeName(),
	}
}

func NewContentPartFromJSObject(obj js.Value) (*ContentPart, error) {
	object, err := js.ToObject(obj)
	if err != nil {
		return nil, err
	}

	return &ContentPart{
		object: object,
	}, nil
}

func (self *ContentPart) SetValue(value any, opts *RenderOptions) error {
	var err error
	switch v := value.(type) {
	case *TemplateResult:
		// 清空当前内容
		tmplElement := self.getTemplateElement(v)

		if instance, ok := self.committedValue.(*TemplateInstance); ok {
			if instance.template != tmplElement {
				instance.Update()
			}
		} else {
			instance := NewTemplateInstance(tmplElement, self, opts.Component)
			fragment, err := instance.CloneTemplate()
			if err != nil {
				return err
			}

			if err = self.commitNode(fragment); err != nil {
				return err
			}

			instance.Update()
			self.committedValue = instance
		}
	case *node.Node:
		err = self.commitNode(v)
	default:
		err = self.commitText(v)
	}

	return err
}

func (self *ContentPart) insert(node *node.Node) *node.Node {
	n, err := WrapNode(self.startNode).ParentNode()
	if err != nil {
		console.Error("ContentPart.insert", err)
		return n
	}

	n, err = WrapNode(n).InsertBefore(node, self.endNode)
	if err != nil {
		console.Error("ContentPart.insert", err)
	}

	return n
}

func (self *ContentPart) clear() {
	if self.startNode == nil {
		return
	}

	startNode := WrapNode(self.startNode)
	nextNode, _ := startNode.NextSibling() // 无需错误处理
	for {
		if nextNode == nil || nextNode.Equal(self.endNode) {
			break
		}

		// Remove the node
		WrapNode(nextNode).Remove()

		// Get the next node
		nextNode, _ = startNode.NextSibling()
	}
}

func (self *ContentPart) commitNode(value *node.Node) error {
	n, ok := self.committedValue.(*node.Node)
	if !ok || !n.Equal(value) {
		self.clear()
		self.committedValue = self.insert(value)
	}

	return nil
}

func (self *ContentPart) commitText(value any) error {
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
		textNode.SetTextContent(value.(string))
		self.commitNode(textNode)
	}

	return nil
}

func (self *ContentPart) ___commitTemplateResult(result *TemplateResult) error {
	// 清空当前内容
	tmplElement := self.getTemplateElement(result)

	if instance, ok := self.committedValue.(*TemplateInstance); ok {
		if instance.template != tmplElement {
			instance.Update()
		}
	} else {
		instance := NewTemplateInstance(tmplElement, self, nil)
		fragment, err := instance.CloneTemplate()
		if err != nil {
			return err
		}

		self.commitNode(fragment)
		instance.Update()
		self.committedValue = instance
	}

	return nil
}

func (self *ContentPart) getTemplateElement(result *TemplateResult) *TemplateElement {
	var tmpl *TemplateElement
	var err error
	/*if err = tmplCacher.Get(result.html, &tmpl); err == nil {
		return tmpl
	}*/

	if tmpl, err = newTemplateElement(result); err != nil {
		panic(err)
		console.Error(err)
	}
	/*
		if err = tmplCacher.Set(&cacher.CacheBlock{Key: result.html, Value: tmpl}); err != nil {
			console.Error(err)
		}*/

	return tmpl
}

func (self *AttributePart) SetValue(value any, opts *RenderOptions) error {
	var change, noCommit bool

	if self.strings == nil || len(self.strings) == 0 {
		value = self.value
		// Single-value binding case
		change = change || self.committedValue != self.value

		if change {
			self.committedValue = self.value
		}
	} else {
		// Interpolation case
		values, ok := self.value.([]any)
		if !ok {

		}
		commitValues, ok := self.committedValue.([]any)
		if !ok {

		}

		value := self.strings[0]

		for i := 1; i < len(self.strings); i++ {
			v := values[i]

			change = change || commitValues[i] != v
			if value != "" {
				value += v.(string) + self.strings[i]
			}

			// We always record each value, even if one is `nothing`, for future
			// change detection.
			commitValues[i] = v
		}

		self.committedValue = commitValues
	}

	if change && !noCommit {
		self.commitValue(value)
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

func (self *EventPart) SetValue(value any, opts *RenderOptions) error {
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
	shouldAddListener := value == nil &&
		(self.committedValue == nil || shouldRemoveListener)

	ele, _ := element.NewFromJSObject(self.element.GetObjectValue())

	if shouldRemoveListener {
		oleLitener := self.committedValue.(func(e event.Event) error)
		ele.RemoveEventListener(self.name, oleLitener)
	}

	if shouldAddListener {
		newListener := value.(func(e event.Event) error)
		ele.AddEventListener(self.name, newListener)
	}

	self.committedValue = value
	return nil
}

func (self *ElementPart) SetValue(value any, opts *RenderOptions) error {
	return nil
}
