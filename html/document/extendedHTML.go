package document

import (
	"github.com/volts-dev/vertex/js"

	"github.com/volts-dev/vertex/html/nodelist"
)

// Close Closer interface
func (d Document) Close() error {

	err := d.Call("close").Error()
	return err

}

func (d Document) GetElementsByName(name string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj = d.Call("getElementsByName", js.ValueOf(name)); obj.Error() == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}
	return nlist, err
}

func (d Document) getSelection() {
	//TO IMPLEMENT
}

func (d Document) HasFocus() (bool, error) {

	return d.GetAttributeBool("hasFocus")
}

// Close Closer interface
func (d Document) Open() error {

	err := d.Call("open").Error()
	return err
}

/* TO IMPLEMENTED
document.queryCommandValue
document.queryCommandSupported
document.queryCommandState
document.queryCommandIndeterm
document.queryCommandEnabled
*/

// Write Writer interface
func (d Document) Write(p []byte) (n int, err error) {
	n = len(p)
	err = d.Call("write", js.ValueOf(string(p))).Error()
	return
}

func (d Document) Writeln(text string) error {

	err := d.Call("writeln", js.ValueOf(text)).Error()
	return err
}
