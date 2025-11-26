package vhtml

import (
	"github.com/volts-dev/cacher"
	_ "github.com/volts-dev/cacher/memory"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/object"
)

type (
	EventPart struct {
		name   string
		method string
	}

	AttributePart struct {
		name  string
		raw   bool
		value string
	}

	BooleanAttributePart struct {
	}

	PropertyPart struct {
	}

	// getTemplateHtml
	ElementPart struct {
		tag  string
		name string
		raw  bool
	}

	ContentPart struct {
		object         object.Object
		committedValue any
	}
)

var tmplCacher cacher.ICacher

func init() {
	tmplCacher, _ = cacher.New("memory")
}

func NewContentPart() *ContentPart {
	obj := object.GetInterface().New()
	o, err := object.ToObject(obj)
	if err != nil {
		return nil
	}
	return &ContentPart{object: o}
}
func NewAttributePart() IPart {
	return AttributePart{}
}

func NewBooleanAttributePart() IPart {
	return BooleanAttributePart{}
}

func NewEventPart() IPart {
	return EventPart{}
}

func NewPropertyPart() IPart {
	return PropertyPart{}
}

func NewContentPartFromJSObject(obj js.Value) (*ContentPart, error) {
	var part ContentPart
	part.object, _ = object.ToObject(obj)
	return &part, nil
}

func (self *ContentPart) SetValue(value any) {
	switch v := value.(type) {
	case *TemplateResult:
		self.commitTemplateResult(v)
	case *node.Node:
		self.commitNode(v)
	default:
		self.commitText(v)
	}
}

func (self *ContentPart) commitNode(*node.Node) error {

	return nil
}

func (self *ContentPart) commitText(value any) error {
	return nil
}

func (self *ContentPart) commitTemplateResult(result *TemplateResult) error {
	// 清空当前内容
	tmplElement := self.getTemplateElement(result)

	if instance, ok := self.committedValue.(*TemplateInstance); ok {
		if instance.template != tmplElement {
			instance.Update()
		}
	} else {
		instance := NewTemplateInstance(tmplElement, self)
		fragment := instance.CloneTemplate()
		self.commitNode(fragment)
		instance.Update()
		self.committedValue = instance
	}

	return nil
}

func (self *ContentPart) getTemplateElement(result *TemplateResult) *TemplateElement {
	var tmpl *TemplateElement
	if err := tmplCacher.Get(result.Id(), &tmpl); err != nil {
		if tmpl, err = newTemplateElement(result); err == nil {
			tmplCacher.Set(&cacher.CacheBlock{
				Key:   result.Id(),
				Value: tmpl,
			})
			return tmpl
		}
	}

	return tmpl
}
