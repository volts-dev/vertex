package component

import (
	"github.com/volts-dev/vertex/core/console"
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
		AttributeChangedCallback(name, oldValue, newValue string)
		// Invoked when a component is moved to a new document.
		AdoptedCallback()
		// ObservedAttributes 返回需要监听的属性列表
		ObservedAttributes() []string
	}
)

// ComponentConstructor 是一个函数类型，用于创建 WebComponent 接口的新实例。
type Constructor func() Component

// https://medium.com/@avicsebooks/super-vs-reflect-construct-2445eefd3b3a
// RegisterComponent 是框架的核心函数。
// 它接收一个组件名称（如 "my-counter"）和一个构造函数，
// 然后在 JavaScript 中注册一个 Custom Element。
func Register(tagName string, constructor func() Component) {
	// 1. 创建一个 JavaScript 函数，它将作为 Custom Element 的构造函数
	var jsConstructor js.Func
	jsConstructor = js.FuncOf(func(this js.Value, args []js.Value) any {
		// A. 获取 HTMLElement 构造函数和 Reflect 对象
		htmlElementConstructor := js.Global().Get("HTMLElement")
		//emptyArgs := js.Global().Get("Array").New() // 创建一个空的JS数组
		// B. 模拟调用 super()，这是最关键的一步
		//    在 JS 中，这相当于: const instance = Reflect.construct(HTMLElement, [], new.target);
		//    'new.target' 在这里就是 'this' (因为 'this' 指向被调用的构造函数本身)
		instance := js.Reflect().Call("construct", htmlElementConstructor, []any{}, jsConstructor)

		// 创建 Go 组件实例
		vcom := constructor()
		//vcom.Constructor()
		//instance.Set("_vertex_component_instance_",  vcom. )

		// 创建 Shadow DOM
		shadowRoot := instance.Call("attachShadow", map[string]interface{}{"mode": "open"})
		// 将 Go 实例和 Shadow DOM 根节点附加到 JS `this` 上，以便后续访问
		this.Set("shadowRoot", shadowRoot)

		// 6. 获取要监听的属性
		// 这是一个静态属性，所以我们直接在构造函数上设置它
		observedAttrs := vcom.ObservedAttributes()
		instance.Set("observedAttributes", observedAttrs)

		// 渲染初始UI
		html := vcom.Render()
		shadowRoot.Set("innerHTML", html)
		console.Log(instance, this)
		return instance
	})

	// 2. 获取 HTMLElement 的原型
	object := js.Global().Get("Object")
	customPrototype := object.New()
	htmlElementPrototype := js.Global().Get("HTMLElement").Get("prototype")
	object.Call("setPrototypeOf", customPrototype, htmlElementPrototype)

	// 3. 创建我们自定义元素的原型
	//componentPrototype := js.Global().Get("Object").Call("create", htmlElementPrototype)

	// 4. 将生命周期回调绑定到原型上
	// connectedCallback
	customPrototype.Set("connectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.ConnectedCallback()
			}
		}

		return nil
	}))

	// disconnectedCallback
	customPrototype.Set("disconnectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.DisconnectedCallback()
				// 注意：在这里可以进行资源释放，比如释放回调函数			}
			}
		}
		return nil
	}))

	// attributeChangedCallback
	customPrototype.Set("attributeChangedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
	jsConstructor.Set("prototype", customPrototype)

	// 7. 使用 customElements.define 进行注册
	js.Global().Get("customElements").Call("define", tagName, jsConstructor)
}
