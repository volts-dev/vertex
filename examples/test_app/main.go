//go:build js && wasm

package main

//"github.com/volts-dev/vertex/component"
import (
	"github.com/volts-dev/vertex/app"
	_ "github.com/volts-dev/vertex/component/not_found"
)

func main() {
	app.Start()
}
