package xmlhttprequest

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/formdata"
	"github.com/volts-dev/vertex/html/json"
	"github.com/volts-dev/vertex/html/progressevent"
	"github.com/volts-dev/vertex/js/reflect"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if xhr, err := New(); test.AssertErr(t, err) {

		test.AssertExpect(t, "[object XMLHttpRequest]", xhr.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	js.Eval("xhr=new XMLHttpRequest()")

	if obj := js.Global().Get("xhr"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "[object XMLHttpRequest]", d.ToString_())

		}
	}

}

func TestGetRequest(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); test.AssertErr(t, err) {

		err := xhr.Open("GET", "http://localhost/get")
		test.AssertErr(t, err)

		xhr.SetOnload(func(i interface{}) {

			if status, err := xhr.Status(); test.AssertErr(t, err) {
				test.AssertExpect(t, status, 200)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); test.AssertErr(t, err) {

				test.AssertExpect(t, "application/json", header)

			}
			if text, err := xhr.ResponseText(); test.AssertErr(t, err) {

				if j, err := json.Parse(text); test.AssertErr(t, err) {
					goValue := j.Map()

					url := goValue.(map[string]interface{})["url"]

					if url != nil {
						test.AssertExpect(t, url, "http://localhost/get")
						io <- true
					} else {
						t.Error("No url present")
					}

				}

			}

		})

		xhr.Send()

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestPostRequest(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); test.AssertErr(t, err) {

		err := xhr.Open("POST", "http://localhost/post")
		test.AssertErr(t, err)

		xhr.SetOnload(func(i interface{}) {
			progressevent.GetInterface()
			test.AssertExpect(t, "[object ProgressEvent]", i.(js.Object).ToString_())

			if status, err := xhr.Status(); test.AssertErr(t, err) {
				test.AssertExpect(t, status, 200)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); test.AssertErr(t, err) {

				test.AssertExpect(t, "application/json", header)

			}
			if text, err := xhr.ResponseText(); test.AssertErr(t, err) {

				if j, err := json.Parse(text); test.AssertErr(t, err) {
					goValue := j.Map()

					form := goValue.(map[string]interface{})["form"]

					if form != nil {
						customValue := form.(map[string]interface{})["data"]

						test.AssertExpect(t, customValue, "testing")
						io <- true

					} else {
						t.Error("No form present")
					}

				}

			}

		})
		f, _ := formdata.New()
		f.Append("data", "testing")

		xhr.Send(f)

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestSetRequestHeader(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); test.AssertErr(t, err) {

		err := xhr.Open("GET", "http://localhost/get")
		test.AssertErr(t, err)

		xhr.SetOnload(func(i interface{}) {
			progressevent.GetInterface()
			test.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			if status, err := xhr.Status(); test.AssertErr(t, err) {
				test.AssertExpect(t, status, 200)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); test.AssertErr(t, err) {

				test.AssertExpect(t, "application/json", header)

			}
			if text, err := xhr.ResponseText(); test.AssertErr(t, err) {

				if j, err := json.Parse(text); test.AssertErr(t, err) {
					goValue := j.Map()

					headers := goValue.(map[string]interface{})["headers"]

					if headers != nil {
						customValue := headers.(map[string]interface{})["Xcustomvalue"]

						if customValue != nil {
							if customValue.(string) == "Test" {
								io <- true
							} else {
								t.Errorf("Xcustomvalue not match %s", customValue.(string))
							}
						} else {
							t.Error("No Xcustomvalue headers present")
						}

					} else {
						t.Error("No headers present")
					}

				}

			}

		})

		xhr.SetRequestHeader("Xcustomvalue", "Test")
		xhr.Send()

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestOnError(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); test.AssertErr(t, err) {
		progressevent.GetInterface()
		err := xhr.Open("GET", "m://httpbin.org/get")
		test.AssertErr(t, err)
		xhr.SetOnError(func(i interface{}) {
			test.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			io <- true
		})

		xhr.Send()

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestOnAbort(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); test.AssertErr(t, err) {
		progressevent.GetInterface()
		err := xhr.Open("GET", "http://localhost/get")
		test.AssertErr(t, err)

		xhr.SetOnAbort(func(i interface{}) {
			test.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			go func() {
				io <- true
			}()

		})

		xhr.Send()
		xhr.Abort() //call SetOnAbort in the same current go routine (be carefull on deadlock channel)

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}
