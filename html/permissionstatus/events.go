package permissionstatus

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/event"
)

func (p PermissionStatus) OnChange(handler func(e event.Event)) (js.Func, error) {

	return p.AddEventListener("change", handler)
}
