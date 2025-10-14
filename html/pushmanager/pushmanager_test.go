package pushmanager

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/promise"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	js.Eval(`swregistercontainer=navigator.serviceWorker;
	promise=swregistercontainer.register('testpushmanager.js');

		p=promise.then(function(registration) {
			pm=registration.pushManager;

		  });


	  
	
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm := js.Global().Get("pm"); test.AssertErr(t, objpm.Error()) {

					if pm, err := NewFromJSObject(objpm); test.AssertErr(t, err) {

						test.AssertExpect(t, "PushManager", pm.ConstructName_())

					}

				}
				wchan <- true

				return nil
			}, func(e error) {

				t.Errorf(e.Error())
			})

			select {
			case <-wchan:
			case <-time.After(time.Duration(200) * time.Millisecond):
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SupportedContentEncodings", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "GetSubscription", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "PermissionState", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Subscribe", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm := js.Global().Get("pm"); test.AssertErr(t, objpm.Error()) {

					if pm, err := NewFromJSObject(objpm); test.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							test.InvokeCheck(t, pm, result)
						}

					}

				}
				wchan <- true

				return nil
			}, func(e error) {

				t.Errorf(e.Error())
			})

			select {
			case <-wchan:
			case <-time.After(time.Duration(200) * time.Millisecond):
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}
