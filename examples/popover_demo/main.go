//go:build js && wasm

package main

import (
	"github.com/volts-dev/vertex/app"
	_ "github.com/volts-dev/vertex/component/popover"
)

func main() {
	app := app.Default()
	app.Start()
}
