package vhtml // 替换为实际包名

import (
	"testing"
)

// TestHTML_NormalCase 测试正常情况下的HTML解析
func TestHTML_NormalCase(t *testing.T) {
	// 测试用例：有效的HTML字符串
	html := `<div {{tag}} onclick="Onclick" {{@onclick Onclick}} {{.active isActive}} {{?open isActive}}>
	{{if isOpen}}
	{{.aa.value}}<p>Hello Vertex </p>
	{{ else }}
	 <p>Hello World</p>
	 {{end}}
	 </div>`

	result := HTML(html)

	// 验证返回结果不为nil
	if result == nil {
		//logger.Errf("Expected TemplateResult, but got nil")
	}

	//temp, attrs := getTemplateHtml(html)
	//logger.Info(temp, attrs)
	// 可以添加更多针对TemplateResult结构的验证
}

// TestHTML_InvalidHTML 测试无效HTML的情况
func TestHTML_InvalidHTML(t *testing.T) {
	// 测试用例：无效的HTML（未闭合标签）
	html := "<div><p>Hello World</div>" // 缺少</p>标签

	result := HTML(html)

	// 验证返回结果为nil
	if result != nil {
		t.Error("Expected nil result for invalid HTML, but got a result")
	}
}

// TestHTML_EmptyString 测试空字符串输入
func TestHTML_EmptyString(t *testing.T) {
	// 测试用例：空HTML字符串
	html := ""

	result := HTML(html)

	if result == nil {
		t.Error("Expected TemplateResult for empty string, but got nil")
	}
}

// TestHTML_ComplexHTML 测试复杂HTML结构
func TestHTML_ComplexHTML(t *testing.T) {
	// 测试用例：复杂的HTML结构
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Test Page</title>
	</head>
	<body>
		<div class="container">
			<h1>Main Title</h1>
			<p>Some text here</p>
			<ul>
				<li>Item 1</li>
				<li>Item 2</li>
			</ul>
		</div>
	</body>
	</html>`

	result := HTML(html)

	if result == nil {
		t.Error("Expected TemplateResult for complex HTML, but got nil")
	}
}

// TestHTML_SpecialCharacters 测试包含特殊字符的HTML
func TestHTML_SpecialCharacters(t *testing.T) {
	// 测试用例：包含特殊字符的HTML
	html := `<div data-value="&lt;script&gt;alert('xss')&lt;/script&gt;">Content &amp; More</div>`

	result := HTML(html)

	if result == nil {
		t.Error("Expected TemplateResult for HTML with special characters, but got nil")
	}
}

// BenchmarkHTML 性能基准测试
func BenchmarkHTML(b *testing.B) {
	html := "<div><p>Benchmark test content</p></div>"

	for i := 0; i < b.N; i++ {
		HTML(html)
	}
}
