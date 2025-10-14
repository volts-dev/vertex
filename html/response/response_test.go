package response

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	var io chan bool = make(chan bool)

	js.Global().Set("waiting", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		io <- true
		return nil
	}))

	js.Eval(`resp = fetch("http://localhost/get")
	fetch('http://localhost/get')
	.then(function(response) {

		resp=response
		waiting()
	})
`)
	<-io
	m.Run()
}

func TestNew(t *testing.T) {

	if d, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object Response]", d.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("r=new Response()")

	if obj := js.Global().Get("r"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Response]", d.ToString_())

		}
	}

}

func TestStatus(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())
			if s, err := response.Status(); test.AssertErr(t, err) {
				test.AssertExpect(t, 200, s)
			}
		}
	}

}

func TestStatusText(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if s, err := response.StatusText(); test.AssertErr(t, err) {
				test.AssertExpect(t, "", s)
			}
		}
	}

}

func TestRedirected(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.Redirected(); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)
			}

		}
	}

}

func TestOk(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.Ok(); test.AssertErr(t, err) {
				test.AssertExpect(t, true, b)
			}

		}
	}

}

func TestType(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if typef, err := response.Type(); test.AssertErr(t, err) {
				test.AssertExpect(t, "cors", typef)
			}

		}
	}

}

func TestUrl(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if url, err := response.Url(); test.AssertErr(t, err) {
				test.AssertExpect(t, "http://localhost/get", url)
			}

		}
	}

}

func TestBody(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if stream, err := response.Body(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object ReadableStream]", stream.ToString_())
			}

		}
	}

}

func TestBodyUsed(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.BodyUsed(); test.AssertErr(t, err) {
				test.AssertExpect(t, false, b)
			}

		}
	}

}

func TestClone(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if clone, err := response.Clone(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Response]", clone.ToString_())
			}

		}
	}

}

func TestError(t *testing.T) {

	if response, err := Error(); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object Response]", response.ToString_())

		if resptype, err := response.Type(); test.AssertErr(t, err) {

			test.AssertExpect(t, "error", resptype)
		}

	}

}

func TestArrayBuffer(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.ArrayBuffer(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestBlob(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Blob(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestJson(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Json(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestText(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Text(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}
