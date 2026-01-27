# WebComponent CSS 快速开始指南

本指南帮助您快速上手 Vertex WebComponent 的 CSS 样式系统。

## 5 分钟快速开始

### 第 1 步：定义样式

```go
package myapp

import (
    "github.com/volts-dev/vertex/core/vhtml"
)

// 定义组件样式
func MyComponentStyles() *vhtml.CSSResult {
    return vhtml.CSS(`
        :host {
            display: block;
            padding: 16px;
        }

        .title {
            font-size: 20px;
            color: #333;
        }

        .content {
            color: #666;
            line-height: 1.6;
        }
    `)
}
```

### 第 2 步：创建组件

```go
import (
    "context"
    "github.com/volts-dev/vertex/component"
    "github.com/volts-dev/vertex/core/vhtml"
)

type MyComponent struct {
    component.StyleAwareComponent
}

func (mc *MyComponent) Constructor() {
    // 初始化时注册样式
    mc.RegisterStyle(MyComponentStyles())
}

func (mc *MyComponent) Styles() *vhtml.CSSResult {
    return MyComponentStyles()
}

func (mc *MyComponent) Render(ctx context.Context) *vhtml.TemplateResult {
    // 返回组件模板（HTML）
    return nil
}
```

### 第 3 步：注册组件

```go
func RegisterMyComponent() {
    component.RegisterWithStyles(
        "my-component",
        func() component.IComponent {
            return &MyComponent{}
        },
        &component.StyleRegistrationOptions{
            Minify: false,
            UseAdoptedStyleSheets: true,
        },
    )
}
```

### 第 4 步：在 HTML 中使用

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body>
    <!-- 使用组件 -->
    <my-component></my-component>

    <script src="app.js"></script>
</body>
</html>
```

## 常见任务

### 任务 1：添加多个样式

```go
func (mc *MyComponent) Constructor() {
    mc.RegisterStyle(BaseStyles())
    mc.RegisterStyle(ThemeStyles())
    mc.RegisterStyle(LayoutStyles())
}

func BaseStyles() *vhtml.CSSResult {
    return vhtml.CSS(`:host { display: block; }`)
}

func ThemeStyles() *vhtml.CSSResult {
    return vhtml.CSS(`:host { --primary: #007bff; }`)
}

func LayoutStyles() *vhtml.CSSResult {
    return vhtml.CSS(`.container { padding: 16px; }`)
}
```

### 任务 2：添加全局样式

```go
// 在应用初始化时
func init() {
    globalStyle := vhtml.CSS(`
        body {
            margin: 0;
            padding: 0;
            font-family: sans-serif;
        }
    `)

    // 注册全局样式
    vhtml.RegisterGlobalStyle("app-base", globalStyle)

    // 注入到文档
    component.InjectGlobalStyles([]*vhtml.CSSResult{globalStyle})
}
```

### 任务 3：使用响应式样式

```go
func ResponsiveStyles() *vhtml.CSSResult {
    return vhtml.CSS(`
        :host {
            display: block;
        }

        .grid {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 16px;
        }

        @media (max-width: 768px) {
            .grid {
                grid-template-columns: repeat(2, 1fr);
                gap: 12px;
            }
        }

        @media (max-width: 480px) {
            .grid {
                grid-template-columns: 1fr;
                gap: 8px;
            }
        }
    `)
}
```

### 任务 4：使用主题变量

```go
func ThemedStyles() *vhtml.CSSResult {
    return vhtml.CSS(`
        :host {
            --primary: #007bff;
            --secondary: #6c757d;
            --danger: #dc3545;
            --success: #28a745;
        }

        .btn {
            padding: 8px 16px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-weight: bold;
        }

        .btn-primary {
            background: var(--primary);
            color: white;
        }

        .btn-secondary {
            background: var(--secondary);
            color: white;
        }

        .btn-danger {
            background: var(--danger);
            color: white;
        }

        .btn-success {
            background: var(--success);
            color: white;
        }
    `)
}
```

### 任务 5：压缩 CSS

```go
component.RegisterWithStyles(
    "my-component",
    func() component.IComponent {
        return &MyComponent{}
    },
    &component.StyleRegistrationOptions{
        Minify: true,  // 启用压缩
    },
)
```

### 任务 6：使用 CSS 变量

```go
// 定义样式
func DynamicStyles(color string) *vhtml.CSSResult {
    return vhtml.CSS([]string{
        `:host { --custom-color: `, `; }`,
        `.element { color: var(--custom-color); }`,
    }, color)
}

// 在组件中使用
colorStyle := DynamicStyles("#ff0000")
```

## 完整示例：计数器组件

```go
package components

import (
    "context"
    "fmt"
    "github.com/volts-dev/vertex/component"
    "github.com/volts-dev/vertex/core/vhtml"
)

// Counter 计数器组件
type Counter struct {
    component.StyleAwareComponent
    count int
}

// Styles 返回样式
func (c *Counter) Styles() *vhtml.CSSResult {
    return vhtml.CSS(`
        :host {
            display: block;
            padding: 20px;
            text-align: center;
            font-family: Arial, sans-serif;
        }

        .container {
            background: #f5f5f5;
            border-radius: 8px;
            padding: 20px;
        }

        .display {
            font-size: 48px;
            font-weight: bold;
            color: #007bff;
            margin: 20px 0;
        }

        button {
            padding: 10px 20px;
            margin: 0 5px;
            border: none;
            border-radius: 4px;
            background: #007bff;
            color: white;
            cursor: pointer;
            font-size: 16px;
            font-weight: bold;
        }

        button:hover {
            background: #0056b3;
        }

        button:active {
            transform: scale(0.95);
        }

        .reset {
            background: #6c757d;
        }

        .reset:hover {
            background: #5a6268;
        }
    `)
}

// Constructor 构造函数
func (c *Counter) Constructor() {
    c.count = 0
    c.RegisterStyle(c.Styles())
}

// Render 渲染方法
func (c *Counter) Render(ctx context.Context) *vhtml.TemplateResult {
    // 实现您的 HTML 模板
    return nil
}

// HandleIncrement 增加计数
func (c *Counter) HandleIncrement() {
    c.count++
}

// HandleDecrement 减少计数
func (c *Counter) HandleDecrement() {
    c.count--
}

// HandleReset 重置计数
func (c *Counter) HandleReset() {
    c.count = 0
}

// ObservedAttributes 观察的属性
func (c *Counter) ObservedAttributes() []string {
    return []string{}
}

// FirstUpdate 首次更新
func (c *Counter) FirstUpdate() error {
    return nil
}

// ConnectedCallback 连接回调
func (c *Counter) ConnectedCallback() {
    fmt.Println("Counter connected")
}

// DisconnectedCallback 断开连接回调
func (c *Counter) DisconnectedCallback() {
    fmt.Println("Counter disconnected")
}

// AttributeChangedCallback 属性变化回调
func (c *Counter) AttributeChangedCallback(name, oldValue, newValue string) {
}

// AdoptedCallback 采用回调
func (c *Counter) AdoptedCallback() {
}

// 注册组件
func RegisterCounter() {
    component.RegisterWithStyles(
        "my-counter",
        func() component.IComponent {
            return &Counter{}
        },
        &component.StyleRegistrationOptions{
            Minify: true,
            UseAdoptedStyleSheets: true,
        },
    )
}
```

## 常见问题

### Q: 样式如何限制在组件内？
A: 样式在 Shadow DOM 中自动隔离。您的样式不会影响其他组件。

### Q: 如何处理响应式设计？
A: 使用 `@media` 查询，如上面的响应式示例所示。

### Q: 能否在多个组件间共享样式？
A: 可以！将样式定义为独立函数，在多个组件中使用。

### Q: 如何动态改变样式？
A: 使用 CSS 变量和 `:host` 伪元素，通过 JavaScript 改变变量值。

### Q: 支持哪些浏览器？
A: 支持所有现代浏览器（Chrome, Firefox, Safari, Edge）。

## 下一步

1. 查看 [STYLE_GUIDE.md](./STYLE_GUIDE.md) 获取完整的 API 文档
2. 浏览 [examples_styles.go](./examples_styles.go) 查看更多示例
3. 查看测试文件了解更多用法

## 有用的链接

- [Lit 官方文档 - 样式](https://lit.dev/docs/components/styles/)
- [Shadow DOM 指南](https://developer.mozilla.org/en-US/docs/Web/Web_Components/Using_shadow_DOM)
- [CSS 自定义属性](https://developer.mozilla.org/en-US/docs/Web/CSS/--*)
- [CSS Grid 布局](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout)

---

**提示**: 开始时最好先参考完整示例，逐步增加功能！
