package component

import (
	"context"

	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/js"
)

// WebComponent 定义了我们的组件需要实现的生命周期方法。
// 类似于 Web Component 的标准生命周期回调。
type (
	Component interface {
		Constructor()
		Styles() []any
		// Render 返回组件的 HTML 内容字符串 (包括 <style> 标签)
		// the component's render() method returns a single TemplateResult object
		Render(context.Context) *vhtml.TemplateResult
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

		FirstUpdate() error
	}
)

// ComponentConstructor 是一个函数类型，用于创建 WebComponent 接口的新实例。
type (
	Constructor func() Component
)

var ()

// https://medium.com/@avicsebooks/super-vs-reflect-construct-2445eefd3b3a
// RegisterComponent 是框架的核心函数。
// 它接收一个组件名称（如 "my-counter"）和一个构造函数，
// 然后在 JavaScript 中注册一个 Custom Element。
func Register(tagName string, constructor func() Component) {
	// 缓存常用全局对象
	global := js.Global()
	object := global.Get("Object")
	htmlElementConstructor := global.Get("HTMLElement")
	reflect := js.Reflect()
	ctx := context.Background()

	// 1. 创建一个 JavaScript 函数，它将作为 Custom Element 的构造函数
	var customConstructor js.Func
	customConstructor = js.FuncOf(func(this js.Value, args []js.Value) any {
		// 模拟调用 super()，这是最关键的一步
		// 在 JS 中，这相当于: const instance = Reflect.construct(HTMLElement, [], new.target);
		// 'new.target' 在这里就是 'this' (因为 'this' 指向被调用的构造函数本身)
		instance := reflect.Call("construct", htmlElementConstructor, []any{}, customConstructor)
		if instance.IsNull() || instance.IsUndefined() {
			panic("failed to construct HTMLElement instance")
		}

		// 创建 Go 组件实例
		vcom := constructor()
		//vcom.Constructor()
		//instance.Set("_vertex_component_instance_",  vcom. )

		// 创建 Shadow DOM
		shadowRoot := instance.Call("attachShadow", map[string]any{"mode": "open"})
		// 将 Go 实例和 Shadow DOM 根节点附加到 JS `this` 上，以便后续访问
		this.Set("shadowRoot", shadowRoot)

		// 6. 获取要监听的属性
		// 这是一个静态属性，所以我们直接在构造函数上设置它
		observedAttrs := vcom.ObservedAttributes()
		instance.Set("observedAttributes", observedAttrs)

		rootNode, err := node.NewFromJSObject(shadowRoot)
		if err != nil {
			console.Log(err)
			panic(err.Error())
		}

		// 渲染初始UI
		htmlResult := vcom.Render(ctx)
		if _, err = vhtml.Render(htmlResult, rootNode, &vhtml.RenderOptions{Component: vcom}); err != nil {
			console.Error(err.Error())
		}
		//vcom.FirstUpdate()
		//shadowRoot.Set("innerHTML", html)
		return instance
	})
	defer customConstructor.Release()

	// 2. 获取 HTMLElement 的原型
	customPrototype := object.New()
	htmlElementPrototype := htmlElementConstructor.Get("prototype")
	object.Call("setPrototypeOf", customPrototype, htmlElementPrototype)

	// 3. 创建我们自定义元素的原型
	//componentPrototype := js.Global().Get("Object").Call("create", htmlElementPrototype)

	// 4. 将生命周期回调绑定到原型上
	// connectedCallback
	connectedCallback := js.FuncOf(func(this js.Value, args []js.Value) any {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.ConnectedCallback()
			}
		}

		return nil
	})
	customPrototype.Set("connectedCallback", connectedCallback)
	defer connectedCallback.Release()

	// disconnectedCallback
	disconnectedCallback := js.FuncOf(func(this js.Value, args []js.Value) any {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				vcom.DisconnectedCallback()
				// 注意：在这里可以进行资源释放，比如释放回调函数			}
			}
		}
		return nil
	})
	customPrototype.Set("disconnectedCallback", disconnectedCallback)
	defer disconnectedCallback.Release()

	// attributeChangedCallback
	attributeChangedCallback := js.FuncOf(func(this js.Value, args []js.Value) any {
		if v := this.Get("_vertex_component_instance_"); !v.IsUndefined() {
			if vcom, ok := v.(Component); ok {
				name := js.ValueToString(args[0])
				oldValue := js.ValueToString(args[1])
				newValue := js.ValueToString(args[2])
				vcom.AttributeChangedCallback(name, oldValue, newValue)
			}
		}

		return nil
	})
	customPrototype.Set("attributeChangedCallback", attributeChangedCallback)
	defer attributeChangedCallback.Release()

	// 5. 将原型与构造函数关联
	customConstructor.Set("prototype", customPrototype)

	// 7. 使用 customElements.define 进行注册
	js.Global().Get("customElements").Call("define", tagName, customConstructor)
}
