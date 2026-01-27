package vcss

import (
	"fmt"
	"sync"
)

// StyleManager 管理组件的样式初始化和渲染
type StyleManager struct {
	mu              sync.RWMutex
	styles          []*CSSResult
	processedStyles map[string]bool
	styleCache      map[string]*CSSResult
}

// NewStyleManager 创建一个新的样式管理器
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:          make([]*CSSResult, 0),
		processedStyles: make(map[string]bool),
		styleCache:      make(map[string]*CSSResult),
	}
}

// AddStyle 添加一个样式到管理器
func (sm *StyleManager) AddStyle(style *CSSResult) error {
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 检查是否已经添加过相同的样式
	if !sm.processedStyles[style.CSSText] {
		sm.styles = append(sm.styles, style)
		sm.processedStyles[style.CSSText] = true
	}

	return nil
}

// AddStyles 批量添加多个样式
func (sm *StyleManager) AddStyles(styles ...*CSSResult) error {
	for _, style := range styles {
		if err := sm.AddStyle(style); err != nil {
			return err
		}
	}
	return nil
}

// GetStyles 获取所有已添加的样式
func (sm *StyleManager) GetStyles() []*CSSResult {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// 返回副本以防止外部修改
	result := make([]*CSSResult, len(sm.styles))
	copy(result, sm.styles)
	return result
}

// CombineStyles 将多个样式合并为一个
func (sm *StyleManager) CombineStyles() *CSSResult {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	if len(sm.styles) == 0 {
		return CSS("")
	}

	// 将所有样式文本合并
	items := make([]interface{}, len(sm.styles))
	for i, style := range sm.styles {
		items[i] = style
	}

	return ConcatCSS(items...)
}

// Clear 清空所有样式
func (sm *StyleManager) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.styles = make([]*CSSResult, 0)
	sm.processedStyles = make(map[string]bool)
}

// Count 返回已添加的样式数量
func (sm *StyleManager) Count() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return len(sm.styles)
}

// ------- StyleRegistry 全局样式注册表 -------

// StyleRegistry 用于注册和管理全局样式
type StyleRegistry struct {
	mu     sync.RWMutex
	styles map[string]*CSSResult
}

// globalStyleRegistry 全局样式注册表实例
var globalStyleRegistry = &StyleRegistry{
	styles: make(map[string]*CSSResult),
}

// RegisterGlobalStyle 注册一个全局样式
func RegisterGlobalStyle(name string, style *CSSResult) error {
	if name == "" {
		return fmt.Errorf("style name cannot be empty")
	}
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}

	globalStyleRegistry.mu.Lock()
	defer globalStyleRegistry.mu.Unlock()

	globalStyleRegistry.styles[name] = style
	return nil
}

// GetGlobalStyle 获取一个全局样式
func GetGlobalStyle(name string) (*CSSResult, error) {
	globalStyleRegistry.mu.RLock()
	defer globalStyleRegistry.mu.RUnlock()

	style, exists := globalStyleRegistry.styles[name]
	if !exists {
		return nil, fmt.Errorf("style '%s' not found", name)
	}
	return style, nil
}

// GetAllGlobalStyles 获取所有全局样式
func GetAllGlobalStyles() map[string]*CSSResult {
	globalStyleRegistry.mu.RLock()
	defer globalStyleRegistry.mu.RUnlock()

	result := make(map[string]*CSSResult)
	for k, v := range globalStyleRegistry.styles {
		result[k] = v
	}
	return result
}

// RemoveGlobalStyle 移除一个全局样式
func RemoveGlobalStyle(name string) error {
	globalStyleRegistry.mu.Lock()
	defer globalStyleRegistry.mu.Unlock()

	if _, exists := globalStyleRegistry.styles[name]; !exists {
		return fmt.Errorf("style '%s' not found", name)
	}
	delete(globalStyleRegistry.styles, name)
	return nil
}

// ClearGlobalStyles 清空所有全局样式
func ClearGlobalStyles() {
	globalStyleRegistry.mu.Lock()
	defer globalStyleRegistry.mu.Unlock()

	globalStyleRegistry.styles = make(map[string]*CSSResult)
}

// ------- Style Rendering -------

// StyleRenderOptions 定义样式渲染选项
type StyleRenderOptions struct {
	// Minify 是否压缩 CSS
	Minify bool
	// UseAdoptedStyleSheets 是否使用 adoptedStyleSheets API
	UseAdoptedStyleSheets bool
	// Namespace 样式命名空间（用于作用域化）
	Namespace string
}

// RenderStyleElement 将样式渲染为 <style> 元素的 HTML
func RenderStyleElement(styles *CSSResult, options *StyleRenderOptions) string {
	if styles == nil {
		return ""
	}

	if options == nil {
		options = &StyleRenderOptions{}
	}

	cssText := styles.CSSText
	if options.Minify {
		cssText = NormalizeCSS(cssText)
	}

	if options.Namespace != "" {
		// 将命名空间添加到 CSS 规则
		cssText = applyNamespace(cssText, options.Namespace)
	}

	return fmt.Sprintf("<style>%s</style>", cssText)
}

// RenderStyleElements 将多个样式渲染为 <style> 元素列表
func RenderStyleElements(styles []*CSSResult, options *StyleRenderOptions) string {
	if len(styles) == 0 {
		return ""
	}

	var result string
	for _, style := range styles {
		result += RenderStyleElement(style, options)
	}
	return result
}

// CreateStyleSheet 创建一个 CSSStyleSheet 对象（JavaScript 环境）
func CreateStyleSheet(cssText string) interface{} {
	// 这个函数在 JavaScript 环境中应该调用：
	// const sheet = new CSSStyleSheet()
	// sheet.replaceSync(cssText)
	// 由于我们在 Go 中，这里返回一个占位符
	return cssText
}

// applyNamespace 给 CSS 规则应用命名空间
func applyNamespace(cssText string, namespace string) string {
	// 简单的命名空间实现
	// 更复杂的实现可能需要解析 CSS AST
	return fmt.Sprintf("%s { %s }", namespace, cssText)
}

// ScopedCSS 创建一个作用域化的 CSS（用于 Shadow DOM）
func ScopedCSS(cssText string) *CSSResult {
	// 在 Shadow DOM 中，样式自动作用域化，
	// 但这个函数可以用于标记意图
	return CSS(cssText)
}

// GlobalCSS 创建一个全局 CSS（应用于整个文档）
func GlobalCSS(cssText string) *CSSResult {
	// 全局样式通常在 <head> 中或通过样式表注入
	return CSS(cssText)
}
