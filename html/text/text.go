package text

import (
	"errors"
	"sync"

	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/js"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented Text")
	ErrNotAnTextElement = errors.New("Object is not an Text")
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once
var htmltextelementinterface js.Value

// HtmlTemplatelement struct
type Text struct {
	htmlelement.HtmlElement
}

type TextFrom interface {
	Text_() Text
}

func (h Text) Text_() Text {
	return h
}

func GetInterface() js.Value {
	singleton.Do(func() {
		if htmltextelementinterface = js.Global().Get("Text"); htmltextelementinterface.Error() != nil {
			htmltextelementinterface = js.Undefined()
		}
		js.Register(htmltextelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltextelementinterface
}

func New(d document.Document) (Text, error) {
	var err error

	var h Text
	var e element.Element

	if e, err = d.CreateElement("title"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (Text, error) {
	var h Text
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.GetObjectValue().InstanceOf(hci) {
			h.SetObjectValue(elem.GetObjectValue())

		} else {
			err = ErrNotAnTextElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (Text, error) {
	var h Text
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.SetObjectValue(obj)

			} else {
				err = ErrNotAnTextElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h Text) Text() (string, error) {
	text := h.GetValueByKey("data")
	return text.String()
}

func (h Text) SetText(value string) {
	h.SetValueByKey("data", value)
}
