package stream

import "github.com/volts-dev/vertex/js"

type TransformStream struct {
	js.Object
}

func TransfertToTransformStream(b js.Object) TransformStream {
	var t TransformStream
	t.SetObjectValue(b.GetObjectValue())
	return t
}
