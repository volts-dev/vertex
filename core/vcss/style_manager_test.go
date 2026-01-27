package vcss

import (
	"testing"
)

// TestStyleManager 样式管理器测试
func TestStyleManager(t *testing.T) {
	sm := NewStyleManager()

	// 测试添加单个样式
	style1 := CSS("div { color: red; }")
	if err := sm.AddStyle(style1); err != nil {
		t.Errorf("AddStyle failed: %v", err)
	}

	// 验证样式已添加
	if sm.Count() != 1 {
		t.Errorf("Expected 1 style, got %d", sm.Count())
	}

	// 测试添加相同样式（应该被过滤）
	if err := sm.AddStyle(style1); err != nil {
		t.Errorf("AddStyle duplicate failed: %v", err)
	}

	if sm.Count() != 1 {
		t.Errorf("Expected 1 style after duplicate, got %d", sm.Count())
	}

	// 测试获取样式
	styles := sm.GetStyles()
	if len(styles) != 1 {
		t.Errorf("Expected 1 style in GetStyles, got %d", len(styles))
	}

	// 测试合并样式
	style2 := CSS("p { color: blue; }")
	sm.AddStyle(style2)

	combined := sm.CombineStyles()
	if !contains(combined.CSSText, "color: red") {
		t.Error("Combined styles should contain 'color: red'")
	}

	if !contains(combined.CSSText, "color: blue") {
		t.Error("Combined styles should contain 'color: blue'")
	}

	// 测试清空
	sm.Clear()
	if sm.Count() != 0 {
		t.Errorf("Expected 0 styles after clear, got %d", sm.Count())
	}
}

// TestStyleRegistry 全局样式注册表测试
func TestStyleRegistry(t *testing.T) {
	// 清空注册表
	ClearGlobalStyles()

	// 注册样式
	style := CSS("body { margin: 0; }")
	if err := RegisterGlobalStyle("app-reset", style); err != nil {
		t.Errorf("RegisterGlobalStyle failed: %v", err)
	}

	// 获取样式
	retrieved, err := GetGlobalStyle("app-reset")
	if err != nil {
		t.Errorf("GetGlobalStyle failed: %v", err)
	}

	if retrieved.CSSText != style.CSSText {
		t.Errorf("Retrieved style doesn't match: %q != %q", retrieved.CSSText, style.CSSText)
	}

	// 获取不存在的样式
	_, err = GetGlobalStyle("non-existent")
	if err == nil {
		t.Error("GetGlobalStyle should return error for non-existent style")
	}

	// 获取所有样式
	allStyles := GetAllGlobalStyles()
	if len(allStyles) != 1 {
		t.Errorf("Expected 1 global style, got %d", len(allStyles))
	}

	// 移除样式
	if err := RemoveGlobalStyle("app-reset"); err != nil {
		t.Errorf("RemoveGlobalStyle failed: %v", err)
	}

	// 验证已移除
	_, err = GetGlobalStyle("app-reset")
	if err == nil {
		t.Error("Style should be removed")
	}

	// 清空所有样式
	RegisterGlobalStyle("style1", CSS("a { }"))
	RegisterGlobalStyle("style2", CSS("b { }"))
	ClearGlobalStyles()

	allStyles = GetAllGlobalStyles()
	if len(allStyles) != 0 {
		t.Errorf("Expected 0 styles after clear, got %d", len(allStyles))
	}
}

// TestStyleRenderOptions 样式渲染选项测试
func TestStyleRenderOptions(t *testing.T) {
	style := CSS(`
		div {
			color: red;
			margin: 10px;
		}
	`)

	// 不压缩
	html := RenderStyleElement(style, &StyleRenderOptions{Minify: false})
	if !contains(html, "<style>") || !contains(html, "</style>") {
		t.Error("Should render style element")
	}

	// 压缩
	html = RenderStyleElement(style, &StyleRenderOptions{Minify: true})
	if !contains(html, "color:red") {
		t.Error("Minified CSS should remove spaces around colons")
	}
}

// TestComponentStyleInitializer 组件样式初始化器测试
func TestComponentStyleInitializer(t *testing.T) {
	initializer := NewComponentStyleInitializer("test-component")

	// 注册样式
	style1 := CSS("div { color: red; }")
	style2 := CSS("span { color: blue; }")

	if err := initializer.RegisterComponentStyle(style1); err != nil {
		t.Errorf("RegisterComponentStyle failed: %v", err)
	}

	if err := initializer.RegisterComponentStyle(style2); err != nil {
		t.Errorf("RegisterComponentStyle failed: %v", err)
	}

	// 获取组件样式
	styles := initializer.GetComponentStyles()
	if len(styles) != 2 {
		t.Errorf("Expected 2 styles, got %d", len(styles))
	}

	// 初始化样式
	initHTML := initializer.InitializeStyles()
	if !contains(initHTML, "color: red") || !contains(initHTML, "color: blue") {
		t.Error("InitializeStyles should include all registered styles")
	}

	// 合并样式
	combined := initializer.CombineComponentStyles()
	if !contains(combined.CSSText, "color: red") || !contains(combined.CSSText, "color: blue") {
		t.Error("Combined styles should contain all styles")
	}

	// 清空
	initializer.Clear()
	styles = initializer.GetComponentStyles()
	if len(styles) != 0 {
		t.Errorf("Expected 0 styles after clear, got %d", len(styles))
	}
}

// TestStyleInjector 样式注入器测试
func TestStyleInjector(t *testing.T) {
	injector := NewStyleInjector()

	style := CSS("div { color: red; }")

	// 注入样式
	if err := injector.InjectStyle("style-1", style); err != nil {
		t.Errorf("InjectStyle failed: %v", err)
	}

	// 检查是否已注入
	if !injector.HasInjected("style-1") {
		t.Error("Style should be marked as injected")
	}

	// 获取注入的样式
	injectedCSS, err := injector.GetInjectedStyle("style-1")
	if err != nil {
		t.Errorf("GetInjectedStyle failed: %v", err)
	}

	if injectedCSS != style.CSSText {
		t.Errorf("Injected style doesn't match: %q != %q", injectedCSS, style.CSSText)
	}

	// 获取不存在的样式
	_, err = injector.GetInjectedStyle("non-existent")
	if err == nil {
		t.Error("Should return error for non-existent style")
	}

	// 生成 HTML
	html := injector.GenerateHTML()
	if !contains(html, "style-1") || !contains(html, "color: red") {
		t.Error("GenerateHTML should include style id and content")
	}

	// 移除样式
	if err := injector.RemoveStyle("style-1"); err != nil {
		t.Errorf("RemoveStyle failed: %v", err)
	}

	if injector.HasInjected("style-1") {
		t.Error("Style should not be injected after removal")
	}

	// 清空
	injector.InjectStyle("style-2", CSS("p { }"))
	injector.Clear()
	if injector.HasInjected("style-2") {
		t.Error("Styles should be cleared")
	}
}

// TestRenderStyleToDOMString DOM 字符串渲染测试
func TestRenderStyleToDOMString(t *testing.T) {
	style := CSS("div { color: red; }")

	// 基础渲染
	html := RenderStyleToDOMString(style, nil)
	if !contains(html, "<style>") || !contains(html, "</style>") {
		t.Error("Should render style tags")
	}

	if !contains(html, "color: red") {
		t.Error("Should include CSS content")
	}

	// 空样式
	emptyHTML := RenderStyleToDOMString(nil, nil)
	if emptyHTML != "" {
		t.Error("Empty style should return empty string")
	}
}

// TestCreateStyleElementCode JavaScript 代码生成测试
func TestCreateStyleElementCode(t *testing.T) {
	code := CreateStyleElementCode("my-style", "div { color: red; }")

	if !contains(code, "my-style") {
		t.Error("Code should include style id")
	}

	if !contains(code, "createElement") {
		t.Error("Code should create element")
	}

	if !contains(code, "textContent") {
		t.Error("Code should set text content")
	}
}

// TestInjectStyleToShadowDOM Shadow DOM 注入代码测试
func TestInjectStyleToShadowDOM(t *testing.T) {
	code := InjectStyleToShadowDOM("shadowRoot", "my-style", "div { }")

	if !contains(code, "my-style") {
		t.Error("Code should include style id")
	}

	if !contains(code, "shadowRoot") {
		t.Error("Code should reference shadow root")
	}

	if !contains(code, "insertBefore") {
		t.Error("Code should use insertBefore")
	}
}

// TestEscapeJavaScriptString JavaScript 字符串转义测试
func TestEscapeJavaScriptString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"simple", "simple"},
		{"with'quote", "with\\'quote"},
		{"with\nnewline", "with\\nnewline"},
		{"with\ttab", "with\\ttab"},
		{"with\\backslash", "with\\\\backslash"},
	}

	for _, tt := range tests {
		if got := escapeJavaScriptString(tt.input); got != tt.expected {
			t.Errorf("escapeJavaScriptString(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

// 辅助函数
func contains(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
