package app

import (
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

func init() {

	js.RegisterInterface(GetDocumentFragmentInterface)
}

var ErrNotADocumentFragment = errors.New("The given value must be a DocumentFragment")

var documentfragementinterface js.Value

type DocumentFragment struct {
	html.Node
}

type DocumentFragmentFrom interface {
	DocumentFragment_() DocumentFragment
}

func (d DocumentFragment) DocumentFragment_() DocumentFragment {
	return d
}

func GetDocumentFragmentInterface() js.Value {

	sync.OnceFunc(func() {
		if documentfragementinterface = js.Global().Get("DocumentFragment"); documentfragementinterface.Error() != nil {
			documentfragementinterface = js.Undefined()
		}
		js.Register(documentfragementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return documentfragementinterface
}

func NewDocumentFragment() (DocumentFragment, error) {

	var d DocumentFragment
	var err error
	var obj js.Value
	if di := GetDocumentFragmentInterface(); !di.IsUndefined() {

		if obj = di.New(); obj.Error() == nil {
			//d.BaseObject = d.SetObject(obj)
			d.SetValue(obj)
		}

	} else {

		err = js.ErrNotImplemented
	}

	return d, err
}

func ToDocumentFragment(obj js.Value) (DocumentFragment, error) {
	var d DocumentFragment
	var err error
	if dci := GetDocumentFragmentInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {

				//	d.BaseObject = d.SetObject(obj)
				d.SetValue(obj)
			} else {
				err = ErrNotADocumentFragment
			}
		}
	} else {
		err = js.ErrNotImplemented
	}
	return d, err
}

func (d DocumentFragment) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (e DocumentFragment) Children() (html.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection html.HTMLCollection

	if obj = e.Get("children"); obj.Error() == nil {

		return html.ToHTMLCollection(obj)
	}

	return collection, err
}

func (e DocumentFragment) getAttributeElement(attribute string) (html.IHTMLElement, error) {
	var nodeObject js.Value
	var newElement html.IHTMLElement
	var err error

	if nodeObject = e.Get(attribute); nodeObject.Error() == nil {
		if nodeObject.IsUndefined() {
			err = html.ErrElementNoChilds

		} else {

			newElement, err = html.ToElement(nodeObject)

		}

	}

	return newElement, err
}

func (d DocumentFragment) FirstElementChild() (html.IHTMLElement, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d DocumentFragment) LastElementChild() (html.IHTMLElement, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d DocumentFragment) nodesMethod(method string, elems ...interface{}) error {
	var arrayJS []interface{}

	for _, elem := range elems {
		arrayJS = append(arrayJS, js.ValueOf(elem))
	}

	return d.Call(method, arrayJS...).Error()

}

func (d DocumentFragment) Prepend(elems ...interface{}) error {
	return d.nodesMethod("prepend", elems...)
}

func (d DocumentFragment) Append(elems ...interface{}) error {
	return d.nodesMethod("append", elems...)
}

func (d DocumentFragment) QuerySelector(selector string) (html.IHTMLElement, error) {

	var err error
	var obj js.Value
	var elem html.IHTMLElement

	if obj = d.Call("querySelector", js.ValueOf(selector)); obj.Error() == nil {

		elem, err = html.ToElement(obj)
	}
	return elem, err
}

func (d DocumentFragment) QuerySelectorAll(selector string) (html.NodeList, error) {

	var err error
	var obj js.Value
	var nlist html.NodeList

	if obj = d.Call("querySelectorAll", js.ValueOf(selector)); obj.Error() == nil {

		nlist, err = html.ToNodeList(obj)
	}
	return nlist, err
}

func (d DocumentFragment) ReplaceChild(new, old html.Node) (html.Node, error) {

	v := d.Call("replaceChild", new.Value(), old.Value())

	return old, v.Error()

}

func (d DocumentFragment) GetElementById(id string) (html.IHTMLElement, error) {

	var err error
	var obj js.Value
	var elem html.IHTMLElement

	if obj = d.Call("getElementById", js.ValueOf(id)); obj.Error() == nil {

		elem, err = html.ToElement(obj)
	}

	return elem, err
}
