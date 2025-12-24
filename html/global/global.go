package global

import (
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/html/document"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

var (
	defaultWindow *window.Window
)

func Alert(message string) {
	js.Self()

}

func Window() *window.Window {
	if defaultWindow != nil {
		return defaultWindow
	}

	w, err := js.Self()
	if err != nil {
		return nil
	}
	win, err := window.NewFromJSObject(w)
	if err != nil {
		console.Error("cant allocate self")
		return nil
	}
	defaultWindow = &win
	return defaultWindow
}

func Document() (document.Document, error) {
	if defaultWindow == nil {
		Window()
	}

	return defaultWindow.Document()
}
