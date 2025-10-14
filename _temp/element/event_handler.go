package html

import (
	"github.com/volts-dev/vertex/core/context"
)

type EventHandler func(ctx context.Context, e Event)

func _ToEventHandler(handler func(ctx context.Context, e Event)) EventHandler {
	return func(ctx context.Context, e Event) {
		if handler != nil {
			handler(ctx, e)
		}
	}
}
