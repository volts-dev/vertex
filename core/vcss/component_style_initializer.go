package vcss

import (
	"fmt"
	"strings"
	"sync"
)

// ComponentStyleInitializer 用于初始化和管理组件的样式
type ComponentStyleInitializer struct {
	mu                 sync.RWMutex
	componentName      string
	styles             []*CSSResult
	globalStyles       []*CSSResult
	defaultOptions     *StyleRenderOptions
	styleElementID     string
	adoptedStyleSheets []interface{}
}

// NewComponentStyleInitializer 创建一个新的组件样式初始化器
func NewComponentStyleInitializer(componentName string) *ComponentStyleInitializer {
	return &ComponentStyleInitializer{
		componentName:  componentName,
		styles:         make([]*CSSResult, 0),
		globalStyles:   make([]*CSSResult, 0),
		styleElementID: fmt.Sprintf("%s-styles", strings.ToLower(componentName)),
		defaultOptions: &StyleRenderOptions{
			Minify:                false,
			UseAdoptedStyleSheets: SupportsAdoptingStyleSheets,
		},
	}
}

// RegisterComponentStyle 注册组件的样式
func (csi *ComponentStyleInitializer) RegisterComponentStyle(style *CSSResult) error {
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}

	csi.mu.Lock()
	defer csi.mu.Unlock()

	csi.styles = append(csi.styles, style)
	return nil
}

// RegisterGlobalStyle 注册全局样式
func (csi *ComponentStyleInitializer) RegisterGlobalStyle(style *CSSResult) error {
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}

	csi.mu.Lock()
	defer csi.mu.Unlock()

	csi.globalStyles = append(csi.globalStyles, style)
	return nil
}

// GetComponentStyles 获取组件样式
func (csi *ComponentStyleInitializer) GetComponentStyles() []*CSSResult {
	csi.mu.RLock()
	defer csi.mu.RUnlock()

	result := make([]*CSSResult, len(csi.styles))
	copy(result, csi.styles)
	return result
}

// GetGlobalStyles 获取全局样式
func (csi *ComponentStyleInitializer) GetGlobalStyles() []*CSSResult {
	csi.mu.RLock()
	defer csi.mu.RUnlock()

	result := make([]*CSSResult, len(csi.globalStyles))
	copy(result, csi.globalStyles)
	return result
}

// SetDefaultRenderOptions 设置默认的渲染选项
func (csi *ComponentStyleInitializer) SetDefaultRenderOptions(options *StyleRenderOptions) {
	csi.mu.Lock()
	defer csi.mu.Unlock()

	if options != nil {
		csi.defaultOptions = options
	}
}

// InitializeStyles 初始化所有样式并返回 HTML 元素字符串
func (csi *ComponentStyleInitializer) InitializeStyles() string {
	csi.mu.RLock()
	defer csi.mu.RUnlock()

	var htmlBuilder strings.Builder

	// 渲染组件样式
	for _, style := range csi.styles {
		htmlBuilder.WriteString(RenderStyleElement(style, csi.defaultOptions))
	}

	// 渲染全局样式
	for _, style := range csi.globalStyles {
		htmlBuilder.WriteString(RenderStyleElement(style, csi.defaultOptions))
	}

	return htmlBuilder.String()
}

// CombineComponentStyles 合并所有组件样式
func (csi *ComponentStyleInitializer) CombineComponentStyles() *CSSResult {
	csi.mu.RLock()
	defer csi.mu.RUnlock()

	if len(csi.styles) == 0 {
		return CSS("")
	}

	items := make([]interface{}, len(csi.styles))
	for i, style := range csi.styles {
		items[i] = style
	}

	return ConcatCSS(items...)
}

// CombineAllStyles 合并所有样式（组件 + 全局）
func (csi *ComponentStyleInitializer) CombineAllStyles() *CSSResult {
	csi.mu.RLock()
	defer csi.mu.RUnlock()

	allStyles := make([]interface{}, 0, len(csi.styles)+len(csi.globalStyles))

	for _, style := range csi.styles {
		allStyles = append(allStyles, style)
	}

	for _, style := range csi.globalStyles {
		allStyles = append(allStyles, style)
	}

	if len(allStyles) == 0 {
		return CSS("")
	}

	return ConcatCSS(allStyles...)
}

// InjectIntoShadowRoot 将样式注入到 Shadow DOM 根节点
// 在 JavaScript 环境中调用
func (csi *ComponentStyleInitializer) InjectIntoShadowRoot() string {
	// 这返回应该在 JavaScript 中执行的代码
	// 用于在 Shadow DOM 中创建和插入样式元素
	//styleElement := RenderStyleElement(csi.CombineComponentStyles(), csi.defaultOptions)

	return fmt.Sprintf(`
(function() {
	const style = document.createElement('style');
	style.id = '%s';
	style.textContent = \'%s\';
	return style;
})()
`, csi.styleElementID, csi.CombineComponentStyles().CSSText)
}

// Clear 清空所有样式
func (csi *ComponentStyleInitializer) Clear() {
	csi.mu.Lock()
	defer csi.mu.Unlock()

	csi.styles = make([]*CSSResult, 0)
	csi.globalStyles = make([]*CSSResult, 0)
}

// ------- StyleInjector 样式注入器 -------

// StyleInjector 用于在 DOM 中注入样式
type StyleInjector struct {
	mu             sync.RWMutex
	injectedStyles map[string]bool
	styleElements  map[string]string
}

// NewStyleInjector 创建一个新的样式注入器
func NewStyleInjector() *StyleInjector {
	return &StyleInjector{
		injectedStyles: make(map[string]bool),
		styleElements:  make(map[string]string),
	}
}

// InjectStyle 注入一个样式
func (si *StyleInjector) InjectStyle(id string, style *CSSResult) error {
	if id == "" {
		return fmt.Errorf("style id cannot be empty")
	}
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}

	si.mu.Lock()
	defer si.mu.Unlock()

	if !si.injectedStyles[id] {
		si.injectedStyles[id] = true
		si.styleElements[id] = style.CSSText
	}

	return nil
}

// GetInjectedStyle 获取注入的样式
func (si *StyleInjector) GetInjectedStyle(id string) (string, error) {
	si.mu.RLock()
	defer si.mu.RUnlock()

	style, exists := si.styleElements[id]
	if !exists {
		return "", fmt.Errorf("style '%s' not found", id)
	}

	return style, nil
}

// HasInjected 检查样式是否已注入
func (si *StyleInjector) HasInjected(id string) bool {
	si.mu.RLock()
	defer si.mu.RUnlock()

	return si.injectedStyles[id]
}

// RemoveStyle 移除已注入的样式
func (si *StyleInjector) RemoveStyle(id string) error {
	si.mu.Lock()
	defer si.mu.Unlock()

	if !si.injectedStyles[id] {
		return fmt.Errorf("style '%s' not found", id)
	}

	delete(si.injectedStyles, id)
	delete(si.styleElements, id)

	return nil
}

// GenerateHTML 生成所有注入的样式的 HTML
func (si *StyleInjector) GenerateHTML() string {
	si.mu.RLock()
	defer si.mu.RUnlock()

	var htmlBuilder strings.Builder

	for id, cssText := range si.styleElements {
		htmlBuilder.WriteString(fmt.Sprintf(
			`<style id="%s">%s</style>`,
			id,
			cssText,
		))
	}

	return htmlBuilder.String()
}

// Clear 清空所有注入的样式
func (si *StyleInjector) Clear() {
	si.mu.Lock()
	defer si.mu.Unlock()

	si.injectedStyles = make(map[string]bool)
	si.styleElements = make(map[string]string)
}

// ------- DOM 渲染辅助函数 -------

// RenderStyleToDOMString 生成可以直接插入 DOM 的 style 标签字符串
func RenderStyleToDOMString(style *CSSResult, options *StyleRenderOptions) string {
	if style == nil {
		return ""
	}

	if options == nil {
		options = &StyleRenderOptions{Minify: false}
	}

	cssText := style.CSSText
	if options.Minify {
		cssText = NormalizeCSS(cssText)
	}

	return fmt.Sprintf(`<style>%s</style>`, cssText)
}

// RenderStylesToDOMString 生成多个样式的 DOM 字符串
func RenderStylesToDOMString(styles []*CSSResult, options *StyleRenderOptions) string {
	if len(styles) == 0 {
		return ""
	}

	var htmlBuilder strings.Builder

	for _, style := range styles {
		htmlBuilder.WriteString(RenderStyleToDOMString(style, options))
	}

	return htmlBuilder.String()
}

// CreateStyleElementCode 生成创建 style 元素的 JavaScript 代码
func CreateStyleElementCode(id string, cssText string) string {
	return fmt.Sprintf(`
(function() {
	const style = document.createElement('style');
	style.id = '%s';
	style.textContent = '%s';
	return style;
})()
`, id, escapeJavaScriptString(cssText))
}

// InjectStyleToShadowDOM 生成注入样式到 Shadow DOM 的 JavaScript 代码
func InjectStyleToShadowDOM(shadowRootVar string, id string, cssText string) string {
	return fmt.Sprintf(`
(function(shadowRoot) {
	const style = document.createElement('style');
	style.id = '%s';
	style.textContent = '%s';
	shadowRoot.insertBefore(style, shadowRoot.firstChild);
})(%s)
`, id, escapeJavaScriptString(cssText), shadowRootVar)
}

// escapeJavaScriptString 转义 JavaScript 字符串
func escapeJavaScriptString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "\\'")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}
