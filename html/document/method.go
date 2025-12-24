package document

import (
	"github.com/volts-dev/vertex/html/attr"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/event"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/nodelist"
	"github.com/volts-dev/vertex/html/treewalker"
	"github.com/volts-dev/vertex/js"
)

func (d Document) AdoptNode(externalNode node.Node) (interface{}, error) {
	var err error
	var obj js.Value
	var r interface{}

	if obj = d.Call("adoptNode", externalNode.GetObjectValue()); obj.Error() == nil {
		r, err = js.Discover(obj)
	}
	return r, err

}

func (d Document) Append(i interface{}) error {
	var err error
	err = d.Call("append", js.ValueOf(i)).Error()
	return err
}

func (d Document) CreateTreeWalker(node node.Node) (*treewalker.TreeWalker, error) {
	var err error
	var obj js.Value
	var tw treewalker.TreeWalker

	if obj = d.Call("createTreeWalker", node.GetObjectValue(), js.ValueOf(129)); obj.Error() == nil {
		tw, err = treewalker.NewFromJSObject(obj)
	}

	return &tw, err
}

func (d Document) CreateAttribute(name string) (attr.Attr, error) {
	var err error
	var obj js.Value
	var attribute attr.Attr

	if obj = d.Call("createAttribute", js.ValueOf(name)); obj.Error() == nil {

		attribute, err = attr.NewFromJSObject(obj)
	}

	return attribute, err
}

func (d Document) CreateComment(comment string) (*node.Node, error) {
	var err error
	var obj js.Value
	var nod *node.Node

	if obj = d.Call("createComment", js.ValueOf(comment)); obj.Error() == nil {
		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) CreateDocumentFragment() (*node.Node, error) {
	var err error
	var obj js.Value
	var nod *node.Node

	if obj = d.Call("createDocumentFragment"); obj.Error() == nil {
		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) CreateHTMLElement(tagname string) (htmlelement.HtmlElement, error) {
	var err error
	var htmlelm htmlelement.HtmlElement
	var elem element.Element

	if elem, err = d.CreateElement(tagname); err == nil {
		htmlelm, err = htmlelement.NewFromElement(elem)

	}
	return htmlelm, err
}

func (d Document) CreateElement(tagname string) (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("createElement", js.ValueOf(tagname)); obj.Error() == nil {
		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) CreateElementNS(namespaceURI string, qualifiedName string) (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("createElementNS", js.ValueOf(namespaceURI), js.ValueOf(qualifiedName)); obj.Error() == nil {
		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) CreateEvent(eventtype string) (event.Event, error) {
	var err error
	var obj js.Value
	var ev event.Event

	if obj = d.Call("createEvent", js.ValueOf(eventtype)); obj.Error() == nil {

		ev, err = event.NewFromJSObject(obj)
	}

	return ev, err
}

func (d Document) createNodeIterator() {
	//TO IMPLEMENT
}

func (d Document) createProcessingInstruction() {
	//TO IMPLEMENT
}

func (d Document) createRange() {
	//TO IMPLEMENT
}

func (d Document) createTreeWalker() {
	//TO IMPLEMENT
}

func (d Document) CreateTextNode(text string) (*node.Node, error) {
	var err error
	var obj js.Value
	var nod *node.Node

	if obj = d.Call("createTextNode", js.ValueOf(text)); obj.Error() == nil {
		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) ElementFromPoint(x, y int) (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("elementFromPoint", js.ValueOf(x), js.ValueOf(y)); obj.Error() == nil {

		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) ElementsFromPoint(x, y int) ([]element.Element, error) {
	var err error
	var obj js.Value
	var elems []element.Element

	if obj = d.Call("elementsFromPoint", js.ValueOf(x), js.ValueOf(y)); obj.Error() == nil {
		for i := 0; i < obj.Length(); {
			if el, err := element.NewFromJSObject(obj.Index(i)); err == nil {
				elems = append(elems, el)
			}

		}
	}

	return elems, err
}

func (d Document) exitPictureInPicture() {
	//TO IMPLEMENT
}

func (d Document) ExitPointerLock() error {
	err := d.Call("exitPointerLock").Error()
	return err
}

func (d Document) getAnimations() {
	//TO IMPLEMENT
}

func (d Document) GetElementsByClassName(classname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = d.Call("getElementsByClassName", js.ValueOf(classname)); obj.Error() == nil {

		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) GetElementsByTagName(tagname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = d.Call("getElementsByTagName", js.ValueOf(tagname)); obj.Error() == nil {

		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) GetElementsByTagNameNS(namespace, tagname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = d.Call("getElementsByTagNameNS", js.ValueOf(namespace), js.ValueOf(tagname)); obj.Error() == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) ImportNode(externalNode node.Node, deep bool) (*node.Node, error) {
	var err error
	var obj js.Value
	var r *node.Node

	if obj = d.Call("importNode", externalNode.GetObjectValue(), js.ValueOf(deep)); obj.Error() == nil {
		r, err = node.NewFromJSObject(obj)
	}
	err = obj.Error()

	return r, err
}

func (d Document) ReleaseCapture() error {
	err := d.Call("releaseCapture").Error()
	return err
}

func (d Document) GetElementById(id string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("getElementById", js.ValueOf(id)); obj.Error() == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			elem, err = element.NewFromJSObject(obj)
		} else {
			err = ErrElementNotFound
		}

	}

	return elem, err
}

func (d Document) QuerySelector(selector string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj = d.Call("querySelector", js.ValueOf(selector)); obj.Error() == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			elem, err = element.NewFromJSObject(obj)
		} else {
			err = ErrElementNotFound
		}
	}
	return elem, err
}

func (d Document) QuerySelectorAll(selector string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj = d.Call("querySelectorAll", js.ValueOf(selector)); obj.Error() == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			nlist, err = nodelist.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}
	}
	return nlist, err
}
