package element

import (
	"errors"

	"github.com/volts-dev/vertex/html/domrect"
	"github.com/volts-dev/vertex/html/domrectlist"
	"github.com/volts-dev/vertex/html/htmlcollection"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/nodelist"
	"github.com/volts-dev/vertex/js"
)

func (e Element) attachShadow() {
	//TODO IMPLEMENT
}

func (e Element) Animate(keyframes, options interface{}) error {
	var argCall []interface{}

	var err error
	if keyframesObject, ok := keyframes.(js.ArrayFrom); ok {
		argCall = append(argCall, keyframesObject.Array_().GetObjectValue())

	}

	if keyframesObject, ok := keyframes.(js.ObjectFrom); ok {
		argCall = append(argCall, keyframesObject.BaseObject_().GetObjectValue())
	}

	if optionsObject, ok := keyframes.(js.ObjectFrom); ok {
		argCall = append(argCall, optionsObject.BaseObject_().GetObjectValue())
	} else {
		argCall = append(argCall, js.ValueOf(options))
	}
	err = e.Call("animate").Error()
	return err
}

func (e Element) After(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.GetObjectValue())
	}

	err = e.Call("after", arrayJS...).Error()

	return err
}

func (e Element) Append(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.GetObjectValue())
	}

	err = e.Call("append", arrayJS...).Error()

	return err
}

func (e Element) Before(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.GetObjectValue())
	}

	err = e.Call("before", arrayJS...).Error()

	return err
}

func (e Element) Closest(query string) (Element, error) {
	var err error
	var obj js.Value
	var elem Element

	if obj = e.Call("closest", js.ValueOf(query)); obj.Error() == nil {

		elem, err = NewFromJSObject(obj)
	}

	return elem, err
}

func (e Element) computedStyleMap() {
	//TODO IMPLEMENT
}

func (e Element) getAnimations() {
	//TODO IMPLEMENT
}

func (e Element) GetAttribute(attributename string) (string, error) {

	var err error
	var obj js.Value
	var newstr string

	if obj = e.Call("getAttribute", js.ValueOf(attributename)); obj.Error() == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			return obj.String()
		}

	}
	return newstr, err
}

func (e Element) GetAttributeNames() (js.Array, error) {

	var err error
	var obj js.Value
	var arr js.Array

	if obj = e.Call("getAttributeNames"); obj.Error() == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			arr, err = js.NewArrayFromJSObject(obj)
		}

	}
	return arr, err
}
func (e Element) GetAttributeNS(namespace, name string) (js.Object, error) {
	var err error
	var obj js.Value
	var newobj js.Object

	if obj = e.Call("getAttributeNS", js.ValueOf(namespace), js.ValueOf(name)); obj.Error() == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			newobj, err = js.ToObject(obj)
		}

	}
	return newobj, err
}

func (e Element) GetBoundingClientRect() (domrect.DOMRect, error) {
	var err error
	var obj js.Value
	var newdomrect domrect.DOMRect

	if obj = e.Call("getBoundingClientRect"); obj.Error() == nil {

		newdomrect, err = domrect.NewFromJSObject(obj)

	}
	return newdomrect, err
}

// retourne un DOMRectList
func (e Element) GetClientRects() (domrectlist.DOMRectList, error) {
	var err error
	var obj js.Value
	var arr domrectlist.DOMRectList

	if obj = e.Call("getClientRects"); obj.Error() == nil {

		arr, err = domrectlist.NewFromJSObject(obj)

	}
	return arr, err
}

func (e Element) GetElementsByClassName(classname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = e.Call("getElementsByClassName", js.ValueOf(classname)); obj.Error() == nil {

		if !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (e Element) GetElementsByTagName(tagname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = e.Call("getElementsByTagName", js.ValueOf(tagname)); obj.Error() == nil {
		if obj.IsNull() || obj.IsUndefined() {
			err = ErrElementsNotFound

		} else {
			collection, err = htmlcollection.NewFromJSObject(obj)
		}
	}

	return collection, err
}

func (e Element) GetElementsByTagNameNS(namespace, tagname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj = e.Call("getElementsByTagNameNS", js.ValueOf(namespace), js.ValueOf(tagname)); obj.Error() == nil {
		if obj.IsNull() || obj.IsUndefined() {
			err = ErrElementsNotFound
		} else {
			collection, err = htmlcollection.NewFromJSObject(obj)

		}
	}

	return collection, err
}

func (e Element) HasAttribute(attributename string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = e.Call("hasChildNodes", js.ValueOf(attributename)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}

	return result, err

}

func (e Element) HasPointerCapture(pointerid int) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = e.Call("hasPointerCapture", js.ValueOf(pointerid)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}
	return result, err
}

func (e Element) InsertAdjacentElement(position string, elem Element) (Element, error) {
	var elemObject js.Value
	var newelem Element
	var err error

	if elemObject = e.Call("insertAdjacentElement", js.ValueOf(position), elem.GetObjectValue()); elemObject.Error() == nil {

		if elemObject.IsNull() {
			err = ErrInsertAdjacent

		} else {
			newelem = elem
		}

	}
	return newelem, err
}

func (e Element) InsertAdjacentHTML(position string, textHTML string) error {

	var err error

	err = e.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(textHTML)).Error()
	return err
}

func (e Element) InsertAdjacentText(position string, text string) error {

	var err error

	err = e.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(text)).Error()
	return err
}

func (e Element) Matches(selector string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj = e.Call("matches", js.ValueOf(selector)); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}
	return result, err
}
func (e Element) pseudo() {
	//TODO IMPLEMENT
}

func (e Element) Prepend(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.GetObjectValue())
	}

	err = e.Call("prepend", arrayJS...).Error()

	return err
}

func (e Element) QuerySelector(selector string) (Element, error) {

	var err error
	var obj js.Value
	var nod Element

	if obj = e.Call("querySelector", js.ValueOf(selector)); obj.Error() == nil {
		if !obj.IsNull() {
			nod, err = NewFromJSObject(obj)
		} else {
			err = errors.New(ErrElementNotFound.Error() + " " + selector)
		}
	}
	return nod, err
}

func (e Element) QuerySelectorAll(selector string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj = e.Call("querySelectorAll", js.ValueOf(selector)); obj.Error() == nil {
		if !obj.IsNull() {
			nlist, err = nodelist.NewFromJSObject(obj)
		} else {
			err = errors.New(ErrElementsNotFound.Error() + " " + selector)
		}
	}
	return nlist, err
}

func (e Element) ReleasePointerCapture(pointerid int) error {
	var err error
	err = e.Call("releasePointerCapture", js.ValueOf(pointerid)).Error()
	return err
}
func (e Element) Remove() error {
	var err error
	err = e.Call("remove").Error()
	return err
}

func (e Element) RemoveAttribute(attrname string) error {
	var err error
	err = e.Call("removeAttribute", js.ValueOf(attrname)).Error()
	return err
}

func (e Element) RemoveAttributeNS(namespace, attrname string) error {
	var err error
	err = e.Call("removeAttributeNS", js.ValueOf(namespace), js.ValueOf(attrname)).Error()
	return err
}

func (e Element) ReplaceChildren(params ...interface{}) error {
	var err error
	var arrayJS []interface{}
	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.GetObjectValue())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	err = e.Call("replaceChildren", arrayJS...).Error()

	return err
}

func (e Element) RequestFullscreen() error {
	var err error
	err = e.Call("requestFullscreen").Error()
	return err
}

func (e Element) RequestPointerLock() error {
	var err error
	err = e.Call("requestPointerLock").Error()
	return err
}

func (e Element) Scroll(x, y int, opts ...map[string]interface{}) error {
	var err error
	var optJSValue []interface{}

	optJSValue = append(optJSValue, js.ValueOf(x))
	optJSValue = append(optJSValue, js.ValueOf(y))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}
	err = e.Call("scroll", optJSValue...).Error()
	return err
}
func (e Element) ScrollTo(x, y int, opts ...map[string]interface{}) error {
	var err error
	var optJSValue []interface{}

	optJSValue = append(optJSValue, js.ValueOf(x))
	optJSValue = append(optJSValue, js.ValueOf(y))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}
	err = e.Call("scrollTo", optJSValue...).Error()
	return err
}

func (e Element) SetAttribute(name, value string) error {
	var err error
	err = e.Call("setAttribute", js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}
func (e Element) SetAttributeNS(namespace, name, value string) error {
	var err error
	err = e.Call("setAttributeNS", js.ValueOf(namespace), js.ValueOf(name), js.ValueOf(value)).Error()
	return err
}

func (e Element) SetPointerCapture(pointerid int) error {
	var err error
	err = e.Call("setPointerCapture", js.ValueOf(pointerid)).Error()
	return err
}

func (e Element) ToggleAttribute(name string, opts ...interface{}) (bool, error) {
	var err error
	var optJSValue []interface{}
	var obj js.Value
	var result bool

	optJSValue = append(optJSValue, js.ValueOf(name))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}

	if obj = e.Call("toggleAttribute", optJSValue...); obj.Error() == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			err = js.ErrObjectNotBool
		}
	}
	return result, err
}
