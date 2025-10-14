package attr

import (
	"testing"

	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

//Attr can't be contructed with New. Test will be done in document
