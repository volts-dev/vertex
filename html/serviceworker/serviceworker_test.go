package serviceworker

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
	promise=swregistercontainer.register('testserviceworker.js');
	p=promise.then(function(registration) {
		serviceworker=registration.installing

	  });
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objsw := js.Global().Get("serviceworker"); test.AssertErr(t, objsw.Error()) {

					if servicew, err := NewFromJSObject(objsw); test.AssertErr(t, err) {

						test.AssertExpect(t, "ServiceWorker", servicew.ConstructName_())

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
	{"method": "ScriptURL", "type": "contains", "resultattempt": "/testserviceworker.js"},
	{"method": "State", "resultattempt": "installing"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)
	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objsw := js.Global().Get("serviceworker"); test.AssertErr(t, objsw.Error()) {

					if servicew, err := NewFromJSObject(objsw); test.AssertErr(t, err) {

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
			case <-time.After(time.Duration(200) * time.Millisecond):
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}
