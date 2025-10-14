package serviceworkerregistration

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
	promise=swregistercontainer.register('testserviceworkerregistration.js');

		p=promise.then(function(registration) {
			reg=registration;
		  });


	  
	
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objreg := js.Global().Get("reg"); test.AssertErr(t, objreg.Error()) {

					if servicew, err := NewFromJSObject(objreg); test.AssertErr(t, err) {

						test.AssertExpect(t, "ServiceWorkerRegistration", servicew.ConstructName_())

					}

				}
				wchan <- true

				return nil
			}, func(e error) {

				t.Errorf(e.Error())
			})

			select {
			case <-wchan:
			case <-time.After(time.Duration(500) * time.Millisecond):
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Installing", "type": "constructnamechecking", "resultattempt": "ServiceWorker"},
	{"method": "Active", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "Scope", "type": "contains", "resultattempt": "localhost"},
	{"method": "Waiting", "type": "error", "resultattempt": js.ErrUndefinedValue},
	{"method": "NavigationPreload", "type": "constructnamechecking", "resultattempt": "NavigationPreloadManager"},
	{"method": "PushManager", "type": "constructnamechecking", "resultattempt": "PushManager"},
	{"method": "GetNotifications", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "ShowNotification", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Unregister", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Update", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "UpdateViaCache", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objreg := js.Global().Get("reg"); test.AssertErr(t, objreg.Error()) {

					if servicew, err := NewFromJSObject(objreg); test.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							test.InvokeCheck(t, servicew, result)
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
			case <-time.After(time.Duration(500) * time.Millisecond):
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}
