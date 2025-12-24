package documentfragment

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/nodelist"
)

func init() {

	js.RegisterInterface(GetInterface)
}

var singleton sync.Once

var documentfragementinterface js.Value

type DocumentFragment struct {
	node.Node
}

type DocumentFragmentFrom interface {
	DocumentFragment_() DocumentFragment
}

func (d DocumentFragment) DocumentFragment_() DocumentFragment {
	return d
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if documentfragementinterface = js.Global().Get("DocumentFragment"); documentfragementinterface.Error() != nil {
			documentfragementinterface = js.Undefined()
		}
		js.Register(documentfragementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return documentfragementinterface
}

func New() (DocumentFragment, error) {

	var d DocumentFragment
	var err error
	var obj js.Value
	if di := GetInterface(); !di.IsUndefined() {

		if obj = di.New(); obj.Error() == nil {
			d.SetObjectValue(obj)
		}

	} else {

		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (DocumentFragment, error) {
	var d DocumentFragment
	var err error
	if dci := GetInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {

				d.SetObjectValue(obj)

			} else {
				err = ErrNotADocumentFragment
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DocumentFragment) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (e DocumentFragment) Children() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = e.GetValueByKey("children"); obj.Error() == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (e DocumentFragment) getAttributeElement(attribute string) (element.Element, error) {
	var nodeObject js.Value
	var newElement element.Element
	var err error

	if nodeObject = e.GetValueByKey(attribute); nodeObject.Error() == nil {

		if nodeObject.IsUndefined() {
			err = element.ErrElementNoChilds

		} else {

			newElement, err = element.NewFromJSObject(nodeObject)

		}

	}

	return newElement, err
}

func (d DocumentFragment) FirstElementChild() (element.Element, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d DocumentFragment) LastElementChild() (element.Element, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d DocumentFragment) nodesMethod(method string, elems ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elems {
		arrayJS = append(arrayJS, js.ValueOf(elem))
	}
	err = d.Call(method, arrayJS...).Error()
	return err

}

func (d DocumentFragment) Prepend(elems ...interface{}) error {
	return d.nodesMethod("prepend", elems...)
}

func (d DocumentFragment) Append(elems ...interface{}) error {
	return d.nodesMethod("append", elems...)
}

func (d DocumentFragment) QuerySelector(selector string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("querySelector", js.ValueOf(selector)); obj.Error() == nil {

		elem, err = element.NewFromJSObject(obj)
	}
	return elem, err
}

func (d DocumentFragment) QuerySelectorAll(selector string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj = d.Call("querySelectorAll", js.ValueOf(selector)); obj.Error() == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}
	return nlist, err
}

func (d DocumentFragment) ReplaceChild(new, old node.Node) (node.Node, error) {
	var err error

	err = d.Call("replaceChild", new.GetObjectValue(), old.GetObjectValue()).Error()

	return old, err

}

func (d DocumentFragment) GetElementById(id string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("getElementById", js.ValueOf(id)); obj.Error() == nil {

		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}
