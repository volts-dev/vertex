package vhtml

import (
	"context"
	"testing"
)

var (
	content = `
	<div {{.id "123"}} {{@onclick "123"}}>
	{{if .IsTrue}}
		Hello {{.name}}
		{{if .IsShow}}
		Welcome to Vertex!
		{{else}}
		Goodbye!
			{{if .IsVertex}}
			 Vertex!
			{{else}}
			 Unknow Man!
			{{end}}
		{{end}}
	{{else}}
		Hello World
	{{end}}
	</div>`
)

// TestGetTemplateHtml_Normal 测试正常情况下的HTML解析
func TestGetTemplateHtml_Normal(t *testing.T) {
	//content = `<div class="goapp-app-info"><p >Hello {{if .IsShowing}}{{.Name}}{{else}}Byte!{{end}}</p></div>`

	ctx := context.WithValue(context.Background(), "_vertex_marker", NewMarker())
	// 测试用例1: 有效HTML内容，无context参数
	result, parts := parseTemplateHtml(content, ctx)
	// 验证返回值
	if result == "" {
		t.Error("Expected non-empty result, got empty string")
	}
	if parts == nil {
		t.Error("Expected non-nil parts, got nil")
	}

	// 测试用例2: 有效HTML内容，有context参数
	result2, parts2 := parseTemplateHtml(content, ctx)
	if result2 == "" {
		t.Error("Expected non-empty result with context, got empty string")
	}

	if parts2 == nil {
		t.Error("Expected non-nil parts with context, got nil")
	}
}
