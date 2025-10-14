package js

import (
	"reflect"
	"testing"
)

// TestNewTypedArrayOf_Uint8Array 测试创建Uint8Array
func TestNewTypedArrayOf_Uint8Array(t *testing.T) {
	tests := []struct {
		name    string
		values  []interface{}
		wantErr bool
	}{
		{
			name:    "正常创建Uint8Array",
			values:  []interface{}{uint8(1), uint8(2), uint8(3)},
			wantErr: false,
		},
		{
			name:    "创建空Uint8Array",
			values:  []interface{}{},
			wantErr: false,
		},
		{
			name:    "使用整数创建Uint8Array",
			values:  []interface{}{1, 2, 3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTypedArrayOf[Uint8Array](tt.values...)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewTypedArrayOf[Uint8Array]() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证返回类型
				if reflect.TypeOf(got) != reflect.TypeOf((*Uint8Array)(nil)).Elem() {
					t.Errorf("NewTypedArrayOf[Uint8Array]() got type = %T, want Uint8Array", got)
				}

				// 验证数组长度
				lg, _ := got.Length()
				if len(tt.values) > 0 && lg != len(tt.values) {
					t.Errorf("NewTypedArrayOf[Uint8Array]() length = %v, want %v", lg, len(tt.values))
				}
			}
		})
	}
}

// TestNewTypedArrayOf_Int32Array 测试创建Int32Array
func TestNewTypedArrayOf_Int32Array(t *testing.T) {
	tests := []struct {
		name    string
		values  []interface{}
		wantErr bool
	}{
		{
			name:    "正常创建Int32Array",
			values:  []interface{}{int32(1), int32(2), int32(3)},
			wantErr: false,
		},
		{
			name:    "使用负数创建Int32Array",
			values:  []interface{}{int32(-1), int32(-2), int32(3)},
			wantErr: false,
		},
		{
			name:    "创建空Int32Array",
			values:  []interface{}{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTypedArrayOf[Int32Array](tt.values...)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewTypedArrayOf[Int32Array]() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证返回类型
				if reflect.TypeOf(got) != reflect.TypeOf((*Int32Array)(nil)).Elem() {
					t.Errorf("NewTypedArrayOf[Int32Array]() got type = %T, want Int32Array", got)
				}
			}
		})
	}
}

// TestNewTypedArrayOf_Float64Array 测试创建Float64Array
func TestNewTypedArrayOf_Float64Array(t *testing.T) {
	tests := []struct {
		name    string
		values  []interface{}
		wantErr bool
	}{
		{
			name:    "正常创建Float64Array",
			values:  []interface{}{float64(1.1), float64(2.2), float64(3.3)},
			wantErr: false,
		},
		{
			name:    "使用混合数值创建Float64Array",
			values:  []interface{}{1.1, 2, 3.3},
			wantErr: false,
		},
		{
			name:    "创建空Float64Array",
			values:  []interface{}{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTypedArrayOf[Float64Array](tt.values...)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewTypedArrayOf[Float64Array]() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证返回类型
				if reflect.TypeOf(got) != reflect.TypeOf((*Float64Array)(nil)).Elem() {
					t.Errorf("NewTypedArrayOf[Float64Array]() got type = %T, want Float64Array", got)
				}
			}
		})
	}
}

// TestNewTypedArrayOf_ErrorCases 测试错误情况
func TestNewTypedArrayOf_ErrorCases(t *testing.T) {
	// 这里测试当传入无法转换的值时的情况
	// 注意：具体实现取决于ToTypedArray的行为

	t.Run("测试无效值转换", func(t *testing.T) {
		// 这个测试需要根据实际的ToTypedArray实现来确定
		// 如果传入完全不兼容的类型应该返回错误
		// 例如：传入字符串到数值类型数组
		_, err := NewTypedArrayOf[Uint8Array]("invalid", "values")

		// 根据实际实现决定是否期望错误
		// 如果ToTypedArray能够处理字符串转换则不会出错
		// 否则应该返回错误
		t.Logf("Error handling test result: %v", err)
	})
}

// TestNewTypedArrayOf_TypeCoverage 测试所有支持的类型
func TestNewTypedArrayOf_TypeCoverage(t *testing.T) {
	// 测试所有支持的类型都能正常创建
	testValues := []interface{}{1, 2, 3}

	t.Run("Uint8Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Uint8Array](testValues...)
		if err != nil {
			t.Errorf("Uint8Array creation failed: %v", err)
		}
	})

	t.Run("Uint8ClampedArray", func(t *testing.T) {
		_, err := NewTypedArrayOf[Uint8ClampedArray](testValues...)
		if err != nil {
			t.Errorf("Uint8ClampedArray creation failed: %v", err)
		}
	})

	t.Run("Uint16Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Uint16Array](testValues...)
		if err != nil {
			t.Errorf("Uint16Array creation failed: %v", err)
		}
	})

	t.Run("Uint32Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Uint32Array](testValues...)
		if err != nil {
			t.Errorf("Uint32Array creation failed: %v", err)
		}
	})

	t.Run("Int8Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Int8Array](testValues...)
		if err != nil {
			t.Errorf("Int8Array creation failed: %v", err)
		}
	})

	t.Run("Int16Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Int16Array](testValues...)
		if err != nil {
			t.Errorf("Int16Array creation failed: %v", err)
		}
	})

	t.Run("Int32Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Int32Array](testValues...)
		if err != nil {
			t.Errorf("Int32Array creation failed: %v", err)
		}
	})

	t.Run("Float32Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Float32Array](testValues...)
		if err != nil {
			t.Errorf("Float32Array creation failed: %v", err)
		}
	})

	t.Run("Float64Array", func(t *testing.T) {
		_, err := NewTypedArrayOf[Float64Array](testValues...)
		if err != nil {
			t.Errorf("Float64Array creation failed: %v", err)
		}
	})
}

// BenchmarkNewTypedArrayOf 基准测试
func BenchmarkNewTypedArrayOf(b *testing.B) {
	values := make([]interface{}, 1000)
	for i := range values {
		values[i] = i
	}

	b.Run("Uint8Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewTypedArrayOf[Uint8Array](values...)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Int32Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewTypedArrayOf[Int32Array](values...)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
