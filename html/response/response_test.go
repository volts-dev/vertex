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

		test.AssertExpect(t, "[object Response]", d.ToString())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("r=new Response()")

	if obj := js.Global().Get("r"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object Response]", d.ToString())

		}
	}

}

func TestStatus(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())
			test.AssertExpect(t, 200, response.Status())
		}
	}

}

func TestStatusText(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, "", response.StatusText())
		}
	}

}

func TestRedirected(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, false, response.Redirected())

		}
	}

}

func TestOk(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, true, response.Ok())

		}
	}

}

func TestType(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, "cors", response.Type())

		}
	}

}

func TestUrl(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, "http://localhost/get", response.Url())

		}
	}

}

func TestBody(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if stream, err := response.Body(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object ReadableStream]", stream.ToString())
			}

		}
	}

}

func TestBodyUsed(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			test.AssertExpect(t, false, response.BodyUsed())

		}
	}

}

func TestClone(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if clone, err := response.Clone(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Response]", clone.ToString())
			}

		}
	}

}

func TestError(t *testing.T) {

	if response, err := Error(); test.AssertErr(t, err) {
		test.AssertExpect(t, "[object Response]", response.ToString())

		test.AssertExpect(t, "error", response.Type())

	}

}

func TestArrayBuffer(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if p, err := response.ArrayBuffer(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString())
			}

		}
	}

}

func TestBlob(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if p, err := response.Blob(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString())
			}

		}
	}

}

func TestJson(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if p, err := response.Json(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString())
			}

		}
	}

}

func TestText(t *testing.T) {

	if obj := js.Global().Get("resp"); test.AssertErr(t, obj.Error()) {

		if response, err := NewFromJSObject(obj); test.AssertErr(t, err) {
			test.AssertExpect(t, "[object Response]", response.ToString())

			if p, err := response.Text(); test.AssertErr(t, err) {
				test.AssertExpect(t, "[object Promise]", p.ToString())
			}

		}
	}

}
