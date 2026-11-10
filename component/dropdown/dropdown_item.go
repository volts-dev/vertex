package dropdown

import (
	"context"

	"github.com/volts-dev/vertex/core/vhtml"
)

type (
	DropdownItem struct {
	}
)

func (self *DropdownItem) Render(ctx context.Context) *vhtml.TemplateResult {
	return vhtml.HTML("")
}
