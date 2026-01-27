package vcss

import (
	"testing"
)

func TestCSSBasic(t *testing.T) {
	result := CSS("div { color: red; }")
	if result.CSSText != "div { color: red; }" {
		t.Errorf("Expected 'div { color: red; }', got '%s'", result.CSSText)
	}
	if result.String() != "div { color: red; }" {
		t.Errorf("String() failed: expected 'div { color: red; }', got '%s'", result.String())
	}
}

func TestCSSWithValues(t *testing.T) {
	result := CSS([]string{"div { margin: ", "px; }"}, 16)
	if result.CSSText != "div { margin: 16px; }" {
		t.Errorf("Expected 'div { margin: 16px; }', got '%s'", result.CSSText)
	}
}

func TestCSSWithCSSResults(t *testing.T) {
	inner := CSS("color: blue")
	outer := CSS([]string{"div { ", "; }"}, inner)
	if outer.CSSText != "div { color: blue; }" {
		t.Errorf("Expected 'div { color: blue; }', got '%s'", outer.CSSText)
	}
}

func TestUnsafeCSS(t *testing.T) {
	result := UnsafeCSS("div { border: 2px solid red; }")
	if result.CSSText != "div { border: 2px solid red; }" {
		t.Errorf("Expected 'div { border: 2px solid red; }', got '%s'", result.CSSText)
	}
	if !result.isCSSResult {
		t.Error("UnsafeCSS result should have isCSSResult = true")
	}
}

func TestCSSMultipleValues(t *testing.T) {
	result := CSS(
		[]string{"div { width: ", "px; height: ", "px; }", "color: red;"},
		100,
		200,
	)
	expected := "div { width: 100px; height: 200px; }color: red;"
	if result.CSSText != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.CSSText)
	}
}

func TestCSSFloatValues(t *testing.T) {
	result := CSS([]string{"opacity: ", ";"}, 0.5)
	if result.CSSText != "opacity: 0.5;" {
		t.Errorf("Expected 'opacity: 0.5;', got '%s'", result.CSSText)
	}
}

func TestCSSPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid input")
		}
	}()
	CSS(struct{}{})
}

func TestTextFromCSSResult(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "CSSResult value",
			value:   CSS("color: red"),
			want:    "color: red",
			wantErr: false,
		},
		{
			name:    "integer value",
			value:   42,
			want:    "42",
			wantErr: false,
		},
		{
			name:    "float value",
			value:   3.14,
			want:    "3.14",
			wantErr: false,
		},
		{
			name:    "nil value",
			value:   nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "invalid string value",
			value:   "plain string",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := textFromCSSResult(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("textFromCSSResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("textFromCSSResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatCSS(t *testing.T) {
	result1 := CSS("color: red")
	result2 := CSS("font-size: 14px")
	concat := ConcatCSS(result1, "; ", result2)

	expected := "color: red; font-size: 14px"
	if concat.CSSText != expected {
		t.Errorf("Expected '%s', got '%s'", expected, concat.CSSText)
	}
}

func TestNormalizeCSS(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "div { color: red; }",
			expected: "div{color:red;}",
		},
		{
			input:    "/* comment */ div { color: red; }",
			expected: "div{color:red;}",
		},
		{
			input:    "div  {  color:  red  ;  }",
			expected: "div{color:red;}",
		},
		{
			input:    "body, div { color: red; }",
			expected: "body,div{color:red;}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := NormalizeCSS(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeCSS(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFlattenCSSResultArray(t *testing.T) {
	result1 := CSS("color: red")
	result2 := CSS("font-size: 14px")
	result3 := CSS("margin: 10px")

	flattened := FlattenCSSResultArray([][]*CSSResult{{result1, result2}, {result3}})

	if len(flattened) != 3 {
		t.Errorf("Expected 3 items, got %d", len(flattened))
	}

	if flattened[0].CSSText != "color: red" {
		t.Errorf("Expected first item to be 'color: red', got '%s'", flattened[0].CSSText)
	}
}

func TestCSSResultMarker(t *testing.T) {
	result := CSS("div { }")
	if !result.GetCSSResultMarker() {
		t.Error("CSSResult marker should be true")
	}
}

func TestCSSStyleSheet(t *testing.T) {
	result := CSS("div { }")
	sheet := result.StyleSheet()
	if sheet != nil {
		t.Error("StyleSheet should return nil when SupportsAdoptingStyleSheets is false")
	}

	mockSheet := "mock-sheet"
	result.SetStyleSheet(mockSheet)
	if result.styleSheet != mockSheet {
		t.Error("SetStyleSheet should update the internal stylesheet")
	}
}
