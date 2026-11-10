package typedarray

import (
	"testing"

	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/reflect"
	"github.com/volts-dev/vertex/test"
)

func TestMain(m *testing.M) {
	reflect.SetSyscall()
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {
	js.Eval(`
aint8 = new Int8Array();
auint8 = new Uint8Array();
auint8clamped = new Uint8ClampedArray();
aint16 = new Int16Array();
auint16 = new Uint16Array();
aint32 = new Int32Array();
auint32 = new Uint32Array();
afloat32 = new Float32Array();
afloat64 = new Float64Array();
`)

	if obj := js.Global().Get("aint8"); test.AssertErr(t, obj.Error()) {
		if d, err := NewInt8FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Int8Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("auint8"); test.AssertErr(t, obj.Error()) {
		if d, err := NewUint8FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Uint8Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("auint8clamped"); test.AssertErr(t, obj.Error()) {
		if d, err := NewUint8ClampedFromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Uint8ClampedArray", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("aint16"); test.AssertErr(t, obj.Error()) {
		if d, err := NewInt16FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Int16Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("auint16"); test.AssertErr(t, obj.Error()) {
		if d, err := NewUint16FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Uint16Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("aint32"); test.AssertErr(t, obj.Error()) {
		if d, err := NewInt32FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Int32Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("auint32"); test.AssertErr(t, obj.Error()) {
		if d, err := NewUint32FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Uint32Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("afloat32"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFloat32FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Float32Array", d.ConstructName_())

		}
	}

	if obj := js.Global().Get("afloat64"); test.AssertErr(t, obj.Error()) {
		if d, err := NewFloat64FromJSObject(obj); test.AssertErr(t, err) {

			test.AssertExpect(t, "Float64Array", d.ConstructName_())

		}
	}

}

func TestNewUint8Array(t *testing.T) {
	if a, err := NewUint8Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 10 {
			t.Errorf("Size must be 10 have %d", l)
		}

		if b, err := a.Subarray(2, 8); test.AssertErr(t, err) {
			if bb, ok := b.(Uint8ArrayFrom); ok {
				l1 := bb.Uint8Array_().ByteLength()
				if l1 != 6 {
					t.Error("Must be equal 6")
				}

			} else {
				t.Error("Not a Uint8array")
			}

		}

	}

}

func TestNewUint8ArrayFrom(t *testing.T) {

	if struint8, err := NewUint8ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := struint8.ByteLength()
		test.AssertExpect(t, int64(11), l)
	}
}

func TestNewUint8ArrayOf(t *testing.T) {

	if ofuint8, err := NewUint8ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := ofuint8.ByteLength()
		test.AssertExpect(t, int64(3), l)

	}
}

func TestNewInt8Array(t *testing.T) {
	if a, err := NewInt8Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 10 {
			t.Errorf("Size must be 10 have %d", l)
		}

		a.Fill(200)
		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(int64) > 0 {
				t.Error("Must be a negative value")
			}
		}

	}
}

func TestNewInt8ArrayFrom(t *testing.T) {

	if strint8, err := NewInt8ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := strint8.ByteLength()
		test.AssertExpect(t, int64(11), l)
	}
}

func TestNewInt8ArrayOf(t *testing.T) {

	if offint8, err := NewInt8ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := offint8.ByteLength()
		test.AssertExpect(t, int64(3), l)

	}
}

func TestNewUInt8ClampedArray(t *testing.T) {
	if a, err := NewUint8ClampedArray(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 10 {
			t.Errorf("Size must be 10 have %d", l)
		}

		a.SetValue(0, 3.6)
		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(int64) != 4 {
				t.Error("Must be a egal 4")
			}
		}

	}

}

func TestNewUint8ClampedArrayFrom(t *testing.T) {

	if struintclamp8, err := NewUint8ClampedArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := struintclamp8.ByteLength()
		test.AssertExpect(t, int64(11), l)
	}
}

func TestNewUint8ClampedArrayOf(t *testing.T) {

	if ofuintclamped8, err := NewUint8ClampedArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := ofuintclamped8.ByteLength()
		test.AssertExpect(t, int64(3), l)

	}
}

func TestNewUInt16Array(t *testing.T) {
	if a, err := NewUint16Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 20 {
			t.Errorf("Size must be 20 have %d", l)
		}

		a.SetValue(0, 1000)
		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(int64) != 1000 {
				t.Error("Must be a egal 1000")
			}
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 2 {
			t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
		}

	}

}
func TestNewUint16ArrayFrom(t *testing.T) {

	if struint16, err := NewUint16ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := struint16.ByteLength()
		test.AssertExpect(t, int64(22), l)
	}
}

func TestNewUint16ArrayOf(t *testing.T) {

	if ofuint16, err := NewUint16ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := ofuint16.ByteLength()
		test.AssertExpect(t, int64(6), l)

	}
}

func TestNewInt16Array(t *testing.T) {
	if a, err := NewInt16Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 20 {
			t.Errorf("Size must be 20 have %d", l)
		}

		a.SetValue(0, 45000)
		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(int64) > 0 {
				t.Error("Must be <0")
			}
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 2 {
			t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
		}

	}

}

func TestNewInt16ArrayFrom(t *testing.T) {

	if strint16, err := NewInt16ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := strint16.ByteLength()
		test.AssertExpect(t, int64(22), l)
	}
}

func TestNewInt16ArrayOf(t *testing.T) {

	if offint16, err := NewInt16ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := offint16.ByteLength()
		test.AssertExpect(t, int64(6), l)

	}
}

func TestNewUint32Array(t *testing.T) {
	if a, err := NewUint32Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 40 {
			t.Errorf("Size must be 20 have %d", l)
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 4 {
			t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
		}

	}

}

func TestNewUint32ArrayFrom(t *testing.T) {

	if struint32, err := NewUint32ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := struint32.ByteLength()
		test.AssertExpect(t, int64(44), l)
	}
}

func TestNewUint32ArrayOf(t *testing.T) {

	if offuint32, err := NewUint32ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := offuint32.ByteLength()
		test.AssertExpect(t, int64(12), l)

	}
}

func TestNewInt32Array(t *testing.T) {
	if a, err := NewInt32Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 40 {
			t.Errorf("Size must be 20 have %d", l)
		}

		a.SetValue(0, 0xEFFFFFFF)

		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(int64) > 0 {
				t.Error("Must be <0")
			}
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 4 {
			t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
		}

	}

}

func TestNewInt32ArrayFrom(t *testing.T) {

	if strint32, err := NewInt32ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := strint32.ByteLength()
		test.AssertExpect(t, int64(44), l)
	}
}

func TestNewInt32ArrayOf(t *testing.T) {

	if offint32, err := NewInt32ArrayOf(3, 8, 4); test.AssertErr(t, err) {

		l := offint32.ByteLength()
		test.AssertExpect(t, int64(12), l)

	}
}

func TestNewFloat32Array(t *testing.T) {
	if a, err := NewFloat32Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 40 {
			t.Errorf("Size must be 20 have %d", l)
		}

		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if float32(v.(float64)) != float32(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 4 {
			t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
		}

	}

}

func TestNewFloat32ArrayFrom(t *testing.T) {

	if strfloat32, err := NewFloat32ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := strfloat32.ByteLength()
		test.AssertExpect(t, int64(44), l)
	}
}

func TestNewFloat32ArrayOf(t *testing.T) {

	if offloat32, err := NewFloat32ArrayOf(3.2, 8.3, 4.3); test.AssertErr(t, err) {

		l := offloat32.ByteLength()
		test.AssertExpect(t, int64(12), l)

	}
}

func TestNewFloat64Array(t *testing.T) {
	if a, err := NewFloat64Array(10); test.AssertErr(t, err) {
		l := a.ByteLength()
		if l != 80 {
			t.Errorf("Size must be 20 have %d", l)
		}

		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); test.AssertErr(t, err) {

			if v.(float64) != float64(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		}

		n := a.BYTES_PER_ELEMENT()
		if n != 8 {
			t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
		}

	}

}

func TestNewFloat64ArrayFrom(t *testing.T) {

	if strfloat64, err := NewFloat64ArrayFrom(js.ArrayFrom_("Hello World")); test.AssertErr(t, err) {

		l := strfloat64.ByteLength()
		test.AssertExpect(t, int64(88), l)
	}
}

func TestNewFloat64ArrayOf(t *testing.T) {

	if offloat64, err := NewFloat64ArrayOf(3.2, 8.3, 4.3); test.AssertErr(t, err) {

		l := offloat64.ByteLength()
		test.AssertExpect(t, int64(24), l)

	}
}
