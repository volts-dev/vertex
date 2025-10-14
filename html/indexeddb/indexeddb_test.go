package indexeddb

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`iddb=window.indexedDB
	`)
	m.Run()
}
