package navigationpreloadmanager

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
	promise=swregistercontainer.register('testnavigationpreloadmanager.js');

		p=promise.then(function(registration) {
			np=registration.navigationPreload;

		  });


	  
	
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm := js.Global().Get("np"); test.AssertErr(t, objpm.Error()) {

					if pm, err := NewFromJSObject(objpm); test.AssertErr(t, err) {

						test.AssertExpect(t, "NavigationPreloadManager", pm.ConstructName_())

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
	{"method": "Enable", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Disable", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "SetHeaderValue", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "GetState", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm := js.Global().Get("np"); test.AssertErr(t, objpm.Error()) {

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
