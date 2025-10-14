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

	if str, err := Stringify(1, "hello", true); test.AssertErr(t, err) {

		test.AssertExpect(t, "[1,\"hello\",true]", str)

	}

}

func TestStringifyObject(t *testing.T) {

	if str, err := StringifyObject(map[string]interface{}{"hello": "world"}); test.AssertErr(t, err) {

		test.AssertExpect(t, "{\"hello\":\"world\"}", str)

	}
}
