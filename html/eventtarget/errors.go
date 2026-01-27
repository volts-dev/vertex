package eventtarget

import "errors"

var (
	ErrNotImplemented   = errors.New("browser not implemented EventTarget")
	ErrNotAnEventTarget = errors.New("object is not an EventTarget")
	ErrInvalidHandler   = errors.New("event handler cannot be nil")
	ErrInvalidEvent     = errors.New("event cannot be nil")
)
