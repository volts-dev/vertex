package component

import (
	"github.com/volts-dev/vertex/js"
	"github.com/volts-dev/vertex/js/helper"
)

// WebComponent 定义了我们的组件需要实现的生命周期方法。
// 类似于 Web Component 的标准生命周期回调。
type (
	Component interface {
		Constructor()
		Styles() string
		// Render 返回组件的 HTML 内容字符串 (包括 <style> 标签)
		Render() string
		// ConnectedCallback 在组件被添加到 DOM 时调用
		ConnectedCallback()
		// DisconnectedCallback 在组件从 DOM 中移除时调用
		DisconnectedCallback()
		// Invoked when one of the element’s observedAttributes changes.
		AttributeChangedCallback(string, string, string)
		// Invoked when a component is moved to a new document.
		AdoptedCallback()
		// ObservedAttributes 返回需要监听的属性列表
		ObservedAttributes() []string
	}
)

// ComponentConstructor 是一个函数类型，用于创建 WebComponent 接口的新实例。
type Constructor func() Component

// RegisterComponent 是框架的核心函数。
// 它接收一个组件名称（如 "my-counter"）和一个构造函数，
// 然后在 JavaScript 中注册一个 Custom Element。
func Register(tagName string, constructor func() Component) {
	// 1. 创建一个 JavaScript 函数，它将作为 Custom Element 的构造函数
	jsConstructor := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// 创建 Go 组件实例
		vcom := constructor()
		vcom.Constructor()

		// 创建 Shadow DOM
		shadowRoot := this.Call("attachShadow", map[string]interface{}{"mode": "open"})

		// 将 Go 实例和 Shadow DOM 根节点附加到 JS `this` 上，以便后续访问
		this.Set("_vertex_component_instance_", vcom)
		this.Set("shadowRoot", shadowRoot)

		// 渲染初始UI
		html := vcom.Render()
		shadowRoot.Set("innerHTML", html)
		return this
	})

	// 2. 获取 HTMLElement 的原型
	htmlElementPrototype := js.Global().Get("HTMLElement").Get("prototype")

	// 3. 创建我们自定义元素的原型
	componentPrototype := js.Global().Get("Object").Call("create", htmlElementPrototype)

	// 4. 将生命周期回调绑定到原型上
	// connectedCallback
	componentPrototype.Set("connectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.ConnectedCallback()
			}

		}
		return nil
	}))

	// disconnectedCallback
	componentPrototype.Set("disconnectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.DisconnectedCallback()
				// 注意：在这里可以进行资源释放，比如释放回调函数			}
			}
		}
		return nil
	}))

	// attributeChangedCallback
	componentPrototype.Set("attributeChangedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				name := helper.ValueToString(args[0])
				oldValue := helper.ValueToString(args[1])
				newValue := helper.ValueToString(args[2])
				vcom.AttributeChangedCallback(name, oldValue, newValue)
			}
		}

		return nil
	}))

	// 5. 将原型与构造函数关联
	jsConstructor.Set("prototype", componentPrototype)

	// 6. 获取要监听的属性
	// 这是一个静态属性，所以我们直接在构造函数上设置它
	observedAttrs := constructor().ObservedAttributes()
	jsConstructor.Set("observedAttributes", js.ValueOf(observedAttrs))

	// 7. 使用 customElements.define 进行注册
	js.Global().Get("customElements").Call("define", tagName, jsConstructor)
}
