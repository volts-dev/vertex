package document

import (
	"sync"

	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/documentfragment"
	"github.com/volts-dev/vertex/html/dragevent"
	"github.com/volts-dev/vertex/html/node"
)

func init() {
	js.RegisterInterface(GetInterface)
}

var singleton sync.Once
var docinterface js.Value

type Document struct {
	node.Node
}

type DocumentFrom interface {
	Document_() Document
}

func (d Document) Document_() Document {
	return d
}

func GetInterface() js.Value {

	singleton.Do(func() {

		if docinterface = js.Global().Get("Document"); docinterface.Error() != nil {
			docinterface = js.Undefined()
		}
		js.Register(docinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		node.GetInterface()
		documentfragment.GetInterface()
		dragevent.GetInterface()

	})

	return docinterface
}

func New() (Document, error) {

	var d Document
	var err error
	if di := GetInterface(); !di.IsUndefined() {

		if dobj := js.Global().Get("document"); dobj.Error() == nil {

			d.SetObjectValue(dobj)
		}

	} else {

		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (Document, error) {
	var d Document
	var err error
	if dci := GetInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = js.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {

				d.SetObjectValue(obj)

			} else {
				err = ErrNotADocument
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}
