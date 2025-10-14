package document

import (
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/html/nodelist"
)

func New_() Document {
	d, _ := New()
	return d
}

func (d Document) Body_() htmlelement.HtmlElement {
	body, _ := d.Body()
	return body
}

func (d Document) QuerySelectorAll_(selector string) nodelist.NodeList {

	n, _ := d.QuerySelectorAll(selector)
	return n
}
