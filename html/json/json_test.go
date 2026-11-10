package json

import (
	"testing"

	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestParse(t *testing.T) {

	var str = `{
		"name":"John",
		"age":30,
		"cars":[ "Ford", "BMW", "Fiat" ]
		}`

	var badstr = `{
			"name":"John",
			"age":30,
			"cars:[ "Ford", "BMW", "Fiat" ]
			}`

	if j, err := Parse(str); test.AssertErr(t, err) {
		goValue := j.Map()

		test.AssertExpect(t, "John", goValue.(map[string]interface{})["name"])

	}

	if _, err := Parse(badstr); err == nil {
		t.Error("Must give an error")
	}
}
func TestStringify(t *testing.T) {

	test.AssertExpect(t, "[1,\"hello\",true]", Stringify(1, "hello", true))

}

func TestStringifyObject(t *testing.T) {

	test.AssertExpect(t, "{\"hello\":\"world\"}", StringifyObject(map[string]interface{}{"hello": "world"}))

}
