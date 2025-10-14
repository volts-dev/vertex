package console

import (
	"testing"

	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

var expectMethods []string = []string{"assert", "clear", "count", "countReset", "debug",
	"dir", "dirxml", "error", "group", "groupCollapsed", "groupEnd", "info", "log", "time", "timeEnd", "timeLog",
	"trace", "warn"}

// infortunately it's not possible to acces js console , we just verify
func TestNew(t *testing.T) {

	if c, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object console]", c.ToString_())

		test.ImplementedExpect(t, c.Object, expectMethods)
	}
}
