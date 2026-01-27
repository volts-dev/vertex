package hello_world

import (
	"context"

	"github.com/volts-dev/vertex/component"
	"github.com/volts-dev/vertex/core/console"
	"github.com/volts-dev/vertex/core/vcss"
	"github.com/volts-dev/vertex/core/vhtml"
	"github.com/volts-dev/vertex/html/event"
)

var (
	// NotFound is the ui element that is displayed when a request is not
	// routed.
	NotFound component.IComponent = &helloWorld{}
)

type helloWorld struct {
	component.Component
	Name      string
	IsShowing bool
	Count     int
	IsAdmin   bool
	HasError  bool
}

func init() {
	component.Register("hello-world", New)
}

func New() component.IComponent {
	return &helloWorld{
		Name: "Vertex",
	}
}

func (n *helloWorld) ObservedAttributes() []string {
	return []string{"ABC", "BCD"}
}

func (n *helloWorld) AttributeChangedCallback(name, oldValue, newValue string) {
	console.Log(`helloWorld attribute changed: %v from %v to %v`, name, oldValue, newValue)
}

func (n *helloWorld) Styles() *vcss.CSSResult {
	return vcss.CSS(`
  		:host {
            --primary-color: #007bff;
            --secondary-color: #6c757d;
            --text-color: #333;
            --bg-color: #fff;
        }

        .btn-primary {
            background-color: var(--primary-color);
            color: white;
        }

        .btn-secondary {
            background-color: var(--secondary-color);
            color: white;
        }

        body {
            color: var(--text-color);
            background-color: var(--bg-color);
        }
	`,
		"")
}

func (n *helloWorld) Render(ctx context.Context) *vhtml.TemplateResult {
	return vhtml.HTML(`
	<div class="goapp-app-info">
		<h1>WebComponent 测试面板</h1>
		
		<!-- 信息展示区域 -->
		<div class="info-section">
			<p class="greeting">{{if IsShowing}}Hello {{.Name}}!{{else}}Welcome to Vertex Framework!{{end}}</p>
			<p class="count-display">当前计数: {{.Count}}</p>
		</div>
		
		<!-- 控制按钮区域 -->
		<div class="button-group">
			<button class="btn-primary" {{@click OnClick}}>
				<span class="btn-text">点击计数</span>
				<span class="btn-count">{{.Count}}</span>
			</button>
			
			<button class="btn-secondary" {{@click OnClick}}>
				切换显示状态
			</button>
		</div>
		
		<!-- 状态指示器 -->
		<div class="status-section">
			<div class="status-item">
				<span class="status-label">显示状态:</span>
				<span class="status-value">{{if IsShowing}}开启{{else}}关闭{{end}}</span>
			</div>
			<div class="status-item">
				<span class="status-label">组件名称:</span>
				<span class="status-value">{{.Name}}</span>
			</div>
		</div>
		
		<!-- 简单条件渲染 - if IsShowing -->
		{{if IsShowing}}
		<div class="conditional-content">
			<h3>✓ 条件内容已显示</h3>
			<p>这是一个只在 IsShowing 为 true 时显示的内容块</p>
		</div>
		{{else}}
		<div class="conditional-content empty">
			<h3>✗ 条件内容已隐藏</h3>
			<p>这是一个只在 IsShowing 为 false 时显示的内容块</p>
		</div>
		{{end}}
		
		<!-- 计数条件渲染 - 不同的计数范围显示不同内容 -->
		{{if   Count > 10}}
		<div class="count-milestone">
			<h3>🎉 恭喜！计数已超过10</h3>
			<p>您已达到重要里程碑！</p>
		</div>
		{{end}}
		{if Count < 5}}
		<div class="count-milestone progress">
			<h3>⚡ 继续加油！计数已达到5</h3>
			<p>即将达到下一个里程碑，继续努力！</p>
		</div>
		{{else}}
		<div class="count-milestone start">
			<p>开始计数，向里程碑迈进...</p>
		</div>
		{{end}}
		
		<!-- 管理员权限条件渲染 -->
		{{if IsAdmin}}
		<div class="admin-panel">
			<h3>👨‍💼 管理员面板</h3>
			<div class="admin-controls">
				<p>您拥有管理员权限，可以访问高级功能</p>
				<button class="btn-admin">重置计数</button>
				<button class="btn-admin">导出数据</button>
			</div>
		</div>
		{{end}}
		
		<!-- 错误状态条件渲染 -->
		{{if HasError}}
		<div class="error-section">
			<h3>⚠️ 错误提示</h3>
			<p class="error-message">发生了一个错误，请检查您的操作</p>
		</div>
		{{if IsShowing}}
		<div class="success-section">
			<h3>✅ 状态正常</h3>
			<p>所有功能运行正常，没有错误</p>
		</div>
		{{end}}
		
		<!-- 复杂条件渲染 - 多条件组合 -->
		{{if  IsShowing &&(  Count < 5)}}
		<div class="milestone-info">
			<p>🌟 您已解锁特殊内容（需要显示状态开启且计数≥5）</p>
		</div>
		{{end}}
		
		{{if   IsAdmin || (Count < 20)}}
		<div class="premium-content">
			<p>💎 高级功能已启用（管理员或计数≥20）</p>
		</div>
		{{end}}
	</div>
	`)
}

func (n *helloWorld) OnClick(e event.Event) error {
	n.Count++
	n.IsShowing = !n.IsShowing
	n.RequestUpdate()
	return nil
}

// lifecycle methods
func (n *helloWorld) FirstUpdate() error {
	n.Name = "Vertex Framework"
	console.Log(n.Name, " FirstUpdate")
	return nil
}

func (n *helloWorld) ConnectedCallback() {
	n.Name = n.Marker + " Vertex Framework"

	/*
		t := time.Tick(8 * time.Second)
			go func() {
				for range t {
					n.IsShowing = !n.IsShowing
					//n.RequestUpdate()
				}
			}()
	*/
}

func (n *helloWorld) DisconnectedCallback() {
	console.Log(n.Name, " disconnected from the DOM")
}
