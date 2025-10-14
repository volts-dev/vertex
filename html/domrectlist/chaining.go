package domrectlist

import "github.com/volts-dev/vertex/html/domrect"

func (d DOMRectList) Item_(index int) domrect.DOMRect {
	domrect, _ := d.Item(index)
	return domrect
}
