package array

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestNewEmpty(t *testing.T) {

	var err error

	var a Array
	var len int

	if a, err = NewEmpty(6); test.AssertErr(t, err) {
		if len, err = a.Length(); test.AssertErr(t, err) {

			test.AssertExpect(t, 6, len)

		}
	}

}

func TestFrom(t *testing.T) {

	var err error

	var a Array
	t.Run("From string", func(t *testing.T) {
		if a, err = From("test"); test.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "t,e,s,t", str)
			}
		}
	})
	t.Run("From array", func(t *testing.T) {
		if a, err = From(New_(1, 2, 3, 4), func(i interface{}) interface{} {
			if vi, ok := i.(int); ok {

				return vi * 3
			}
			return i
		}); test.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "3,6,9,12", str)

			}
		}

	})

}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value
	var a Array

	js.Eval("customarray=new Array(1,2,5)")
	if obj = js.Global().Get("customarray"); test.AssertErr(t, obj.Error()) {
		if a, err = NewFromJSObject(obj); test.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "1,2,5", str)
			}
		}
	}
}

func TestConcat(t *testing.T) {

	var a Array
	var err error

	if a, err = New(1, 2, 3); test.AssertErr(t, err) {
		if c, err := a.Concat(New_(6, 7, 8)); test.AssertErr(t, err) {
			var str string
			if str, err = c.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "1,2,3,6,7,8", str)
			}
		}
	}
}

func TestCopyWithin(t *testing.T) {

	var a Array
	var err error

	if a, err = New("a", "b", "c", "d", "e"); test.AssertErr(t, err) {
		if c, err := a.CopyWithin(0, 3, 4); test.AssertErr(t, err) {
			var str string
			if str, err = c.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "d,b,c,d,e", str)

			}
		}
	}

}

func TestEntries(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"a", "b", "c", "d", "e"}

	if a, err = New(goArray...); test.AssertErr(t, err) {

		if it, err := a.Entries(); test.AssertErr(t, err) {
			var loop int
			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

				if str, ok := value.(string); ok {
					if i, ok := index.(int); ok {

						if str != goArray[i] {
							test.AssertExpect(t, goArray[i], str)
						}
					} else {
						t.Errorf("Index is not int")
					}

				} else {
					t.Errorf("Value is not string")
				}

				loop++

			}

			if loop != len(goArray) {
				t.Errorf("Loop entries not match")
			}

		}

	}

}

func TestEvery(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4, 5}

	if a, err = New(goArray...); err == nil {

		if b, _ := a.Every(func(i interface{}) bool {

			if i.(int) < 13 {
				return true
			}

			return false
		}); !b {
			t.Errorf("number must be < 13")
		}

		if b, _ := a.Every(func(i interface{}) bool {

			if i.(int) < 2 {
				return true
			}

			return false
		}); b {
			t.Errorf("some number must be >  2")
		}
	}

}

func TestFill(t *testing.T) {
	var a Array
	var err error

	if a, err = NewEmpty(5); test.AssertErr(t, err) {

		if err := a.Fill(7); test.AssertErr(t, err) {

			if ok, _ := a.Every(func(i interface{}) bool {

				if i.(int) == 7 {
					return true
				}

				return false
			}); !ok {
				t.Errorf("must be fill with 7")
			}

			if ok, _ := a.Every(func(i interface{}) bool {

				if i.(int) != 7 {
					return true
				}

				return false
			}); ok {
				t.Errorf("must be fill with 7")
			}

		}
	}
}

func TestFilter(t *testing.T) {
	var a Array
	var err error

	if a, err = New("spray", "limit", "elite", "exuberant", "destruction", "present"); err == nil {

		//select all len word > 7
		if b, err := a.Filter(func(i interface{}) bool {
			if len(i.(string)) > 7 {
				return true
			}
			return false
		}); test.AssertErr(t, err) {
			if str, err := b.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "exuberant,destruction", str)

			}
		}

	}

}

func TestFind(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		if found, err := a.Find(func(i interface{}) bool {

			if i.(int) > 10 {

				return true
			}

			return false

		}); test.AssertErr(t, err) {

			if found != nil {
				test.AssertExpect(t, 12, found)
			} else {
				t.Errorf("no element found")
			}

		}

	}

}

func TestFindIndex(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		if found, err := a.FindIndex(func(i interface{}) bool {

			if i.(int) == 12 {

				return true
			}

			return false

		}); test.AssertErr(t, err) {
			if found >= 0 {
				test.AssertExpect(t, 2, found)

			} else {
				t.Errorf("no element found")
			}

		}

	}

}

func TestFlat(t *testing.T) {

	var a Array
	var err error

	var goArray []interface{} = []interface{}{1, 2, []interface{}{3, 4}}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if b, err := a.Flat(); test.AssertErr(t, err) {

			if str, err := b.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "1,2,3,4", str)

			}
		}

	}

}

func TestFlatMap(t *testing.T) {

	var goArray []interface{} = []interface{}{1, 2, 3, 4}

	if a, err := New(goArray...); test.AssertErr(t, err) {

		if b, err := a.FlatMap(func(i1 interface{}, i2 int) interface{} {

			b1 := Of_(i1.(int) * 2)
			return b1.GetObjectValue()

		}); test.AssertErr(t, err) {

			if str, err := b.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "2,4,6,8", str)
			}
		}

	}

}

func TestForEach(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		var count int = 0
		a.ForEach(func(i interface{}) {

			test.AssertExpect(t, goArray[count], i)
			count++

		})

		if count != len(goArray) {
			t.Errorf("Bad number of element")
		}

	}

}

func TestIncludes(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if ok, err := a.Includes("limit"); test.AssertErr(t, err) {
			if !ok {

				t.Errorf("Must include limit")

			}
		}

		if ok, err := a.Includes("limit2"); test.AssertErr(t, err) {
			if ok {

				t.Errorf("Must not include limit")

			}
		}

	}

}

/*
func TestIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		obj := a.GetObjectValue().Index(2)
		b, _ := object.ToObject(obj)

		if index, err := a.IndexOf(b); test.AssertErr(t, err) {

			test.AssertExpect(t, 2, index)

			if index != 2 {
				t.Errorf("index must be 2 have %d when searching %s", index, helper.ValueToString(obj))
			}

		}

		if index, err := a.IndexOf("elite"); test.AssertErr(t, err) {

			test.AssertExpect(t, 2, index)

		}

	}
}
*/

/*
func TestIsArray(t *testing.T) {
	var a Array
	var err error

	if a, err = NewEmpty(3); test.AssertErr(t, err) {
		if ok, err := IsArray(a.Object); test.AssertErr(t, err) {
			if !ok {
				t.Errorf("Must be an array")
			}
		}

		if ok, err := IsArray(object.Object{}); test.AssertErr(t, err) {
			if ok {
				t.Errorf("Must not be an array")
			}
		}

	}

}
*/

func TestJoin(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		if str, err := a.Join("|"); test.AssertErr(t, err) {

			test.AssertExpect(t, "Hello|World|elite", str)

		}

	}

}

func TestKeys(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		var i int = 0
		if it, err := a.Keys(); test.AssertErr(t, err) {
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				test.AssertExpect(t, value, i)

				i++

			}
		}

	}
}

/*
func TestLastIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present", "limit"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		obj := a.GetObjectValue().Index(6)
		b, _ := object.ToObject(obj)

		if index, err := a.LastIndexOf(b); test.AssertErr(t, err) {

			test.AssertExpect(t, 6, index)
		}

		if index, err := a.LastIndexOf("limit"); test.AssertErr(t, err) {

			test.AssertExpect(t, 6, index)

		}

	}

}
*/

func TestMap(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		if b, err := a.Map(func(i interface{}) interface{} {
			if vi, ok := i.(int); ok {

				return vi * 3
			}
			return i
		}); test.AssertErr(t, err) {
			if str, err := b.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "3,6,9,12", str)

			}
		}

	}
}

func TestPop(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"hello"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if err := a.Pop(); test.AssertErr(t, err) {
			if l, _ := a.Length(); l != 0 {

				t.Errorf("Array must be empty now")

			}
		}
	}

}

func TestPush(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"hello"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if length, err := a.Push("world"); test.AssertErr(t, err) {

			test.AssertExpect(t, length, 2)

			if str, err := a.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "hello,world", str)

			}

		}
	}

}

func TestReduce(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if value, err := a.Reduce(func(accumulateur, value interface{}, opts ...interface{}) interface{} {
			val := accumulateur.(int) + value.(int)

			return val
		}); test.AssertErr(t, err) {

			test.AssertExpect(t, 10, value)

		}

	}
}

func TestReduceRight(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if value, err := a.ReduceRight(func(accumulateur, value interface{}, opts ...interface{}) interface{} {
			val := accumulateur.(int) - value.(int)

			return val
		}); test.AssertErr(t, err) {

			test.AssertExpect(t, 17, value)

		}

	}
}

func TestReverse(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if err := a.Reverse(); test.AssertErr(t, err) {
			if str, err := a.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "40,8,6,9", str)

			}

		}
	}
}

func TestShift(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if v, err := a.Shift(); test.AssertErr(t, err) {

			test.AssertExpect(t, 9, v)

			if str, err := a.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "6,8,40", str)

			}
		}

	}
}

func TestSlice(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"ant", "bison", "camel", "duck", "elephant"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if v, err := a.Slice(2); test.AssertErr(t, err) {

			if str, err := v.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "camel,duck,elephant", str)
			}

		}

		if v, err := a.Slice(2, 4); test.AssertErr(t, err) {

			if str, err := v.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "camel,duck", str)
			}

		}
		if v, err := a.Slice(-2); test.AssertErr(t, err) {

			if str, err := v.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "duck,elephant", str)

			}

		}

	}
}

func TestSome(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if ok, err := a.Some(func(i interface{}) bool {

			if i.(int) == 40 {
				return true
			}

			return false
		}); test.AssertErr(t, err) {

			test.AssertExpect(t, true, ok)

		}

		if ok, err := a.Some(func(i interface{}) bool {

			if i.(int) == 42 {
				return true
			}

			return false
		}); test.AssertErr(t, err) {

			test.AssertExpect(t, false, ok)

		}

	}
}

func TestSort(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"March", "Jan", "Feb", "Dec"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if err := a.Sort(); test.AssertErr(t, err) {
			if str, err := a.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "Dec,Feb,Jan,March", str)

			}
		}
	}
}

func TestSplice(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Jan", "March", "April", "June"}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if err := a.Splice(1, 0, "Feb"); test.AssertErr(t, err) {
			if str, err := a.ToString(); test.AssertErr(t, err) {
				test.AssertExpect(t, "Jan,Feb,March,April,June", str)
			}
		}

		if err := a.Splice(4, 1, "May"); test.AssertErr(t, err) {
			if str, err := a.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "Jan,Feb,March,April,May", str)

			}
		}
	}
}

func TestUnshift(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3}
	if a, err = New(goArray...); test.AssertErr(t, err) {

		if l, err := a.Unshift(4, 5); test.AssertErr(t, err) {

			if str, err := a.ToString(); test.AssertErr(t, err) {

				test.AssertExpect(t, "4,5,1,2,3", str)

			}
			test.AssertExpect(t, 5, l)

		}

	}
}

func TestValues(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); test.AssertErr(t, err) {
		var i int = 0
		if it, err := a.Values(); test.AssertErr(t, err) {
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {
				test.AssertExpect(t, goArray[i], value)
				i++

			}
		}

	}

}

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}
