package domstringlist

import (
	"testing"

	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
