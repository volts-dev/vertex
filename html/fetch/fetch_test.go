package fetch

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/html/abortcontroller"
	"github.com/volts-dev/vertex/html/json"
	"github.com/volts-dev/vertex/html/promise"
	"github.com/volts-dev/vertex/html/response"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	//Start promise and wait result
	t.Run("Get ", func(t *testing.T) {
		io := make(chan bool)
		if f, err := New("http://localhost/get"); test.AssertErr(t, err) {
			f.Then(func(r response.Response) *promise.Promise {

				if status, err := r.Status(); test.AssertErr(t, err) {
					if status != 200 {
						t.Errorf("Status must be 200 , give %d", status)
					}
					io <- true
				}
				return nil
			}, func(e error) {

				t.Error(e.Error())
			})
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	})
	t.Run("Get with custom headers", func(t *testing.T) {
		io := make(chan bool)
		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/json",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "GET", "headers": headers}

		//Start promise and wait result
		if f, err := New("http://localhost/get", fetchOpts); test.AssertErr(t, err) {
			textpromise, _ := f.Then(func(r response.Response) *promise.Promise {
				if status, err := r.Status(); test.AssertErr(t, err) {
					if status != 200 {
						t.Errorf("Status must be 200 , give %d", status)
					} else {

						if promise, err := r.Text(); test.AssertErr(t, err) {
							return &promise
						}

					}
				}
				return nil
			}, func(e error) {

				t.Error(e.Error())
			})

			textpromise.Then(func(i interface{}) *promise.Promise {

				if j, err := json.Parse(i.(string)); test.AssertErr(t, err) {
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

				return nil
			}, func(e error) {
				t.Error(e.Error())
			})
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Post with custom headers", func(t *testing.T) {
		io := make(chan bool)
		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/json",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers}

		//Start promise and wait result
		if f, err := New("http://localhost/post", fetchOpts); test.AssertErr(t, err) {
			textpromise, _ := f.Then(func(r response.Response) *promise.Promise {
				if status, err := r.Status(); test.AssertErr(t, err) {
					if status != 200 {
						t.Errorf("Status must be 200 , give %d", status)
					} else {

						if promise, err := r.Text(); test.AssertErr(t, err) {
							return &promise
						}

					}
				}
				return nil
			}, func(e error) {

				t.Error(e.Error())
			})

			textpromise.Then(func(i interface{}) *promise.Promise {

				if j, err := json.Parse(i.(string)); test.AssertErr(t, err) {
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

				return nil
			}, func(e error) {
				t.Error(e.Error())
			})
		}

		select {
		case <-io:
		case <-time.After(time.Duration(3000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	js.GetObjectInterface()

	t.Run("Post with custom headers and json response and form data ", func(t *testing.T) {
		io := make(chan bool)
		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test"}

		//Start promise and wait result
		if f, err := New("http://localhost/post", fetchOpts); test.AssertErr(t, err) {
			jsonpromise, _ := f.Then(func(r response.Response) *promise.Promise {
				if status, err := r.Status(); test.AssertErr(t, err) {
					if status != 200 {
						t.Errorf("Status must be 200 , give %d", status)
					} else {

						if promise, err := r.Json(); test.AssertErr(t, err) {
							return &promise
						}

					}
				}
				return nil
			}, func(e error) {

				t.Error(e.Error())
			})

			jsonpromise.Then(func(i interface{}) *promise.Promise {

				if obj, ok := i.(js.ObjectFrom); ok {
					j, _ := json.NewFromJSObject(obj.GetObjectValue())
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

				} else {
					t.Error("No a json")
				}

				return nil
			}, func(e error) {
				t.Error(e.Error())
			})
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

}

func TestNewCancelable(t *testing.T) {
	var io chan bool = make(chan bool)

	t.Run("Post with custom headers and json response and form data ", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test", "mode": "no-cors"}

		if f, err := NewCancelable("http://localhost/post", fetchOpts); test.AssertErr(t, err) {
			f.Then(func(r response.Response) *promise.Promise {

				t.Error("Must not get response")
				return nil
			}, func(e error) {
				if e.Error() != "signal is aborted without reason" {
					t.Error("Error mismatch:", e.Error())
				}
				io <- true
			})

			f.Abort()
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Post with custom headers and json response and custom signal ", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		abortctrl, _ := abortcontroller.New()

		s, _ := abortctrl.Signal()

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test", "mode": "no-cors", "signal": s.GetObjectValue()}

		if f, err := NewCancelable("http://localhost/post", fetchOpts); test.AssertErr(t, err) {
			f.Then(func(r response.Response) *promise.Promise {

				t.Error("Must not get response")
				return nil
			}, func(e error) {
				if e.Error() != "signal is aborted without reason" {
					t.Error("Error mismatch:", e.Error())
				}
				io <- true
			})

			if err := f.Abort(); err != ErrSignalNotManaged {

				t.Error("Must throw an error")
			}
			abortctrl.Abort()
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

}
