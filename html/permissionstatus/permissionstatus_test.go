package permissionstatus

import (
	"testing"
	"time"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"

	"github.com/volts-dev/vertex/html/promise"
)

func TestMain(m *testing.M) {

	reflect.SetSyscall()

	js.Eval(`
	p=navigator.permissions.query({name:'clipboard-read'}).then(function(permissionStatus) {
		permstatus=permissionStatus;
		});
		`)
	m.Run()

}

func TestNewFromJSObject(t *testing.T) {
	var wchan chan bool = make(chan bool)
	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); test.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objperm := js.Global().Get("permstatus"); test.AssertErr(t, objperm.Error()) {

					if permstatus, err := NewFromJSObject(objperm); test.AssertErr(t, err) {

						test.AssertExpect(t, "PermissionStatus", permstatus.ConstructName_())

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
				t.Errorf("Permission request timeout")

			}

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": "clipboard_read"},
	{"method": "State", "resultattempt": "prompt"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)
	if obj := js.Global().Get("p"); test.AssertErr(t, obj.Error()) {

		if p, err := promise.NewFromJSObject(obj); console.AssertErr(err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objperm := js.Global().Get("permstatus"); test.AssertErr(t, objperm.Error()) {

					if permstatus, err := NewFromJSObject(objperm); test.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							test.InvokeCheck(t, permstatus, result)
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
				t.Errorf("Permission request timeout")

			}

		}

	}

}
