package storage

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNewStorage(t *testing.T) {

	t.Run("session", func(t *testing.T) {
		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object Storage]", stor.ToString())
				}

			}
		}

	})

	t.Run("local", func(t *testing.T) {
		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {
					test.AssertExpect(t, "[object Storage]", stor.ToString())
				}

			}
		}

	})

}

func TestGetItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		js.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					test.AssertExpect(t, "world", stor.GetItem("hello"))
					test.AssertExpect(t, nil, stor.GetItem("hello2"))
				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		js.Eval("window.localStorage.setItem(\"hello\",\"world\")")
		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					test.AssertExpect(t, "world", stor.GetItem("hello"))
					test.AssertExpect(t, nil, stor.GetItem("hello2"))
				}

			}
		}

	})

}

func TestSetItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		js.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("hello", "you"); test.AssertErr(t, err) {
						test.AssertExpect(t, "you", stor.GetItem("hello"))
					}
				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		js.Eval("window.localStorage.setItem(\"hello\",\"world\")")
		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("hello", "you"); test.AssertErr(t, err) {
						test.AssertExpect(t, "you", stor.GetItem("hello"))
					}
				}

			}
		}

	})

}

func TestRemoveItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		js.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("objrmv", "yes"); test.AssertErr(t, err) {
						if err := stor.RemoveItem("objrmv"); test.AssertErr(t, err) {

							test.AssertExpect(t, nil, stor.GetItem("objrmv"))
						}
					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		js.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("objrmv", "yes"); test.AssertErr(t, err) {
						if err := stor.RemoveItem("objrmv"); test.AssertErr(t, err) {
							test.AssertExpect(t, nil, stor.GetItem("objrmv"))

						}
					}

				}

			}
		}

	})

}

func TestClear(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("objclear", "yes"); test.AssertErr(t, err) {
						if err := stor.Clear(); test.AssertErr(t, err) {
							test.AssertExpect(t, nil, stor.GetItem("objclear"))

						}
					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {

					if err := stor.SetItem("objclear", "yes"); test.AssertErr(t, err) {
						if err := stor.Clear(); test.AssertErr(t, err) {
							test.AssertExpect(t, nil, stor.GetItem("objclear"))

						}
					}

				}

			}
		}

	})

}

func TestKey(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("sessionStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {
					stor.Clear()

					if err := stor.SetItem("objkey", "yes"); test.AssertErr(t, err) {

						str := stor.Key(0)
						test.AssertExpect(t, "objkey", str)

					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {

		if obj := js.Global().Get("window"); test.AssertErr(t, obj.Error()) {

			if storageobj := obj.Get("localStorage"); test.AssertErr(t, storageobj.Error()) {
				if stor, err := NewFromJSObject(storageobj); test.AssertErr(t, err) {
					stor.Clear()

					if err := stor.SetItem("objkey", "yes"); test.AssertErr(t, err) {

						str := stor.Key(0)
						test.AssertExpect(t, "objkey", str)

					}

				}

			}
		}

	})

}
