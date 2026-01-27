//go:build js && wasm

package main

import (
	"github.com/volts-dev/vertex/app"
	_ "github.com/volts-dev/vertex/component/hello_world"
)

func main() {
	app:=app.Default()
}
