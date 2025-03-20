//go:build !wasm

package html

import "testing"

func TestCreator(t *testing.T) {
	a := htmlA{}.New()
	t.Log(a)
}
