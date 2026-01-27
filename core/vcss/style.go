package vcss

import (
	"fmt"

	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

// StyleAwareComponent 为组件添加样式管理功能
type StyleAwareComponent struct {
	//Component
	styleInitializer *ComponentStyleInitializer
	styleInjector    *StyleInjector
}

// NewStyleAwareComponent 创建一个支持样式的组件
func NewStyleAwareComponent(componentName string) *StyleAwareComponent {
	return &StyleAwareComponent{
		styleInitializer: NewComponentStyleInitializer(componentName),
		styleInjector:    NewStyleInjector(),
	}
}

// RegisterStyle 注册组件样式
func (sac *StyleAwareComponent) RegisterStyle(style *CSSResult) error {
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}
	return sac.styleInitializer.RegisterComponentStyle(style)
}

// RegisterGlobalStyle 注册全局样式
func (sac *StyleAwareComponent) RegisterGlobalStyle(style *CSSResult) error {
	if style == nil {
		return fmt.Errorf("style cannot be nil")
	}
	return sac.styleInitializer.RegisterGlobalStyle(style)
}

// GetComponentStyles 获取组件样式
func (sac *StyleAwareComponent) GetComponentStyles() []*CSSResult {
	return sac.styleInitializer.GetComponentStyles()
}

// GetAllStyles 获取所有样式
func (sac *StyleAwareComponent) GetAllStyles() []*CSSResult {
	styles := make([]*CSSResult, 0)
	styles = append(styles, sac.styleInitializer.GetComponentStyles()...)
	styles = append(styles, sac.styleInitializer.GetGlobalStyles()...)
	return styles
}

// StyleInitializer 返回样式初始化器
func (sac *StyleAwareComponent) StyleInitializer() *ComponentStyleInitializer {
	return sac.styleInitializer
}

// ------- StyleInjectionManager 样式注入管理器 -------

// StyleInjectionManager 负责在 WebComponent 注册时处理样式
type StyleInjectionManager struct {
	globalInjector     *StyleInjector
	componentInjectors map[string]*StyleInjector
}

// NewStyleInjectionManager 创建样式注入管理器
func NewStyleInjectionManager() *StyleInjectionManager {
	return &StyleInjectionManager{
		globalInjector:     NewStyleInjector(),
		componentInjectors: make(map[string]*StyleInjector),
	}
}

// InjectComponentStyles 注入组件样式到 Shadow DOM
func (sim *StyleInjectionManager) InjectComponentStyles(
	shadowRoot *node.Node,
	styles []*CSSResult,
	options *StyleRenderOptions,
) error {
	if len(styles) == 0 {
		return nil
	}

	if options == nil {
		options = &StyleRenderOptions{
			Minify:                false,
			UseAdoptedStyleSheets: SupportsAdoptingStyleSheets,
		}
	}

	// 如果支持 adoptedStyleSheets，使用它
	if options.UseAdoptedStyleSheets && SupportsAdoptingStyleSheets {
		return sim.injectViaAdoptedStyleSheets(shadowRoot, styles, options)
	}

	// 否则，使用传统的 <style> 标签方式
	return sim.injectViaStyleElements(shadowRoot, styles, options)
}

// injectViaStyleElements 通过 <style> 元素注入样式
func (sim *StyleInjectionManager) injectViaStyleElements(
	shadowRoot *node.Node,
	styles []*CSSResult,
	options *StyleRenderOptions,
) error {
	for i, style := range styles {
		cssText := style.CSSText
		if options.Minify {
			cssText = NormalizeCSS(cssText)
		}

		// 创建 <style> 元素
		doc, err := window.Default().Document()
		if err != nil {
			return err
		}

		styleElement, err := doc.CreateElement("style")
		if err != nil {
			return err
		}
		styleElement.SetTextContent(cssText)
		styleElement.SetAttribute("id", fmt.Sprintf("style-%d", i))

		// 插入到 Shadow DOM
		firstChild, _ := shadowRoot.FirstChild()
		shadowRoot.InsertBefore(&styleElement.Node, firstChild)
	}

	return nil
}

// injectViaAdoptedStyleSheets 通过 adoptedStyleSheets API 注入样式
func (sim *StyleInjectionManager) injectViaAdoptedStyleSheets(
	shadowRoot *node.Node,
	styles []*CSSResult,
	options *StyleRenderOptions,
) error {
	// 这需要在支持 adoptedStyleSheets 的环境中运行
	// 通常在现代浏览器中
	cssStyleSheetConstructor := js.Global().Get("CSSStyleSheet")
	if cssStyleSheetConstructor.IsNull() || cssStyleSheetConstructor.IsUndefined() {
		// 降级到 <style> 元素方式
		return sim.injectViaStyleElements(shadowRoot, styles, options)
	}

	adoptedSheets := make([]any, 0)
	for _, style := range styles {
		cssText := style.CSSText
		if options.Minify {
			cssText = NormalizeCSS(cssText)
		}

		// 创建 CSSStyleSheet
		sheet := cssStyleSheetConstructor.New()
		sheet.Call("replaceSync", cssText)

		// 添加到数组
		adoptedSheets = append(adoptedSheets, sheet)
	}

	// 设置 adoptedStyleSheets
	if !shadowRoot.IsNull() {
		shadowRoot.SetAttribute("adoptedStyleSheets", adoptedSheets)
	}

	return nil
}

// GetComponentInjector 获取或创建组件的样式注入器
func (sim *StyleInjectionManager) GetComponentInjector(componentName string) *StyleInjector {
	if injector, exists := sim.componentInjectors[componentName]; exists {
		return injector
	}

	injector := NewStyleInjector()
	sim.componentInjectors[componentName] = injector
	return injector
}

// ------- 辅助函数 -------
// PrepareComponentStyles 准备组件样式供注入
func PrepareComponentStyles(styles []*CSSResult, options *StyleRenderOptions) string {
	if len(styles) == 0 {
		return ""
	}

	return RenderStylesToDOMString(styles, options)
}

// GenerateStyleInitializationCode 生成样式初始化代码
func GenerateStyleInitializationCode(
	componentName string,
	styles []*CSSResult,
	options *StyleRenderOptions,
) string {
	if len(styles) == 0 {
		return ""
	}

	var code string

	// 合并所有样式
	combinedCSS := ""
	for _, style := range styles {
		combinedCSS += style.CSSText + "\n"
	}

	if options != nil && options.Minify {
		combinedCSS = NormalizeCSS(combinedCSS)
	}

	// 生成初始化代码
	code = fmt.Sprintf(`
	// Initialize styles for %s component
	(function() {
		const style = document.createElement('style');
		style.id = '%s-styles';
		style.textContent = \'%s\';
		return style;
	})();
`, componentName, componentName, escapeBackticks(combinedCSS))

	return code
}

// escapeBackticks 转义反引号
func escapeBackticks(s string) string {
	return fmt.Sprintf("%q", s)
}
