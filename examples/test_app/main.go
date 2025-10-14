//go:build js && wasm

package main

import (
	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/component/not_found"
)

func main() {
	//v := vertex.New()

	component.Register("not-found", not_found.New)
}
