import { html, css } from 'lit-element';
import '../../element/dataset/data-source.js';
import '../view-search/view-search.js';
import '../view-tree/view-tree.js';
import '../view-form/view-form.js';
import store from '@/store';

export class ActionViewer extends store.Element {
  static get styles() {
    return css`
      :host {
        display: block;
      }
    `;
  }

  static get properties() {
    return {
      viewId: { type: Number, value: 0 },
      pages: { type: Object },
      modeOfPage: { type: Object }, // mode 对应的Page idex

      //action: { type: Object },
      //fields: { type: Object }, // # 该Action可用的所有字段
      //views: { type: Object }, // # load_viewsLoad结果
      // lastUrl: { type: String },
      // lastView: { type: String }, // 之前的View名称
      // default_view: { type: String }, // # action 的默认模板
      // viewIdOfMode: { type: Object }, // mode 对应的 View ID
      //*********** nodes ************
      searchView: { type: Object }, // 废弃
      selectedView: { type: Object }, // 目前显示的视图
    };
  }

  render() {
    return html`
      <data-source
        id="datasource"
        action="/dataset/search_read"
        method="POST"
      ></data-source>
      <iron-pages id="pages" selected="0">
        <slot></slot>
      </iron-pages>
    `;
  }

  constructor() {
    super();
  }

  firstUpdated() {
    super.firstUpdated();

    if (!this.controlPanel) {
      this.controlPanel = this.app.shadowRoot.querySelector('#control_panel');
    }

    if (!this.viewManager) {
      this.viewManager = this;
    }

    self = this;
    self.modeOfPage = {};
    self.viewIdOfMode = {};
    self.datasource = self.shadowRoot.querySelector('#datasource');
    self.pages = self.shadowRoot.querySelector('#pages');
    //this.control_elements = {};
    //if (this.flags.search_view) {
    //this.search_view_loaded = this.setup_search_view();
    self.setup_search_view();
    //  }
    //if (this.flags.views_switcher) {
    self.render_switch_buttons();
    //  }

    self.controlPanel.addEventListener(
      'onViewModeChanged',
      this.onViewModeChanged.bind(this)
    );
    self.controlPanel.addEventListener('search', this.doSearch.bind(this)); // Switch to the default_view to load it
  }

  // 控制板视图更新触发
  onViewModeChanged(e) {
    if (
      this.pages.selected != e.detail.pageIdex ||
      this.selectedView != e.detail.view
    ) {
      this.selectedView = e.detail.view;
      this.selectedView.controlPanel = this.controlPanel;
      this.selectedView.viewManager = this;
      this.selectedView.datasource = this.datasource;

      this.pages.selected = e.detail.pageIdex; // 切换页面
      this.selectedView = e.detail.view;

      this.params = e.detail.params; // 初始化数据集查询参数
      this.searchAndShow();
    }
  }

  doSearch(e) {
    this.params = e.detail;
    this.searchAndShow(); // 显示搜索结果
  }

  // # 执行查询并显示更新
  searchAndShow() {
    this.datasource.action = ''; // 初始化地址由视图控件重新定义

    // # 渲染View 和 control_elements
    var self = this;
    if (this.selectedView) {
      // 清空初始化节点
      var breadcrumb_node = self.controlPanel.controlElements.breadcrumb;
      while (breadcrumb_node.children.length > 0) {
        breadcrumb_node.removeChild(breadcrumb_node.firstChild);
      }

      var searchview_node = self.controlPanel.controlElements.searchView;
      if (searchview_node.firstChild) {
        searchview_node.firstChild.hidden = true;
      }
      // while (searchview_node.children.length > 0) {
      //    searchview_node.firstChild.hidden=true;
      //}

      var button_node = self.controlPanel.controlElements.buttons;
      while (button_node.children.length > 0) {
        button_node.removeChild(button_node.firstChild);
      }

      var sidebar_node = self.controlPanel.controlElements.sidebar;
      while (sidebar_node.children.length > 0) {
        sidebar_node.removeChild(sidebar_node.firstChild);
      }

      var pager_node = self.controlPanel.controlElements.pager;
      while (pager_node.children.length > 0) {
        pager_node.removeChild(pager_node.firstChild);
      }

      var view = this.selectedView;
      view.app = this.controlPanel.app;
      view.params = this.params;
      //view.fields=this.controlPanel.fields;
      // act on element after it's been updated / rendered.
      view.updateComplete.then(() => {
        view.show(this);

        // 重建节点
        try {
          if (view.render_cp_one) {
            view.render_cp_one(self.controlPanel.controlElements.cp_one);
          }

          if (view.render_cp_two) {
            view.render_cp_two(self.controlPanel.controlElements.cp_two);
          }

          if (view.render_cp_tree) {
            view.render_cp_tree(self.controlPanel.controlElements.cp_tree);
          }

          if (view.render_cp_four) {
            view.render_cp_four(self.controlPanel.controlElements.cp_four);
          }

          // # 回调View 实现的接口函数
          if (view.renderHeader) {
            // # 渲染CRUD按钮
            view.renderHeader(breadcrumb_node);
          }
          if (view.renderSearchview) {
            // # 渲染CRUD按钮
            view.renderSearchview(searchview_node);
          }
          if (view.renderButtons) {
            // # 渲染CRUD按钮
            view.renderButtons(button_node);
          }
          if (view.renderSidebar) {
            view.renderSidebar(sidebar_node);
          }
          if (view.renderPager) {
            view.renderPager(pager_node);
          }
        } catch (err) {
          console.log(err);
          //在这里处理错误
        }
      });
    }
  }

  /**
   * Sets up the current viewmanager's search view.
   * Sets $searchView and $searchview_buttons in control_elements to send to the controlPanel
   *
   * @param {Number|false} view_id the view to use or false for a default one
   * @returns {jQuery.Deferred} search view startup deferred
   */
  setup_search_view() {
    var self = this;
    if (this.searchView) {
      this.searchView.destroy();
    }
    /*
        var view_id = (this.action && this.action.search_view_id && this.action.search_view_id[0]) || false;
    
        var search_defaults = {};
    
        var context = this.action ? this.action.context : [];
        _.each(context, function (value, key) {
            var match = /^search_default_(.*)$/.exec(key);
            if (match) {
                search_defaults[match[1]] = value;
            }
        });
    
    
        var options = {
            hidden: this.flags.search_view === false,
            disable_custom_filters: this.flags.search_disable_custom_filters,
            $buttons: $("<div>"),
            action: this.action,
        };
        // Instantiate the SearchView, but do not append it nor its buttons to the DOM as this will
        // be done later, simultaneously to all other controlPanel elements
        this.searchView = new SearchView(this, this.dataset, view_id, search_defaults, options);
    
        this.searchView.on('search_data', this, this.search.bind(this));
        return $.when(this.searchView.appendTo($("<div>"))).done(function() {
            self.control_elements.$searchView = self.searchView.$el;
            self.control_elements.$searchview_buttons = self.searchView.$buttons.contents();
        });
        */
  }

  /**
   * Renders the switch buttons and adds listeners on them but does not append them to the DOM
   * Sets $switch_buttons in control_elements to send to the controlPanel
   * @param {Object} [src] the source requesting the switch_buttons
   * @param {Array} [views] the array of views
   */
  render_switch_buttons() {
    /*
         if (this.flags.views_switcher && this.view_order.length > 1) {
             var self = this;
    
             // Render switch buttons but do not append them to the DOM as this will
             // be done later, simultaneously to all other controlPanel elements
             this.control_elements.$switch_buttons = $(QWeb.render('ViewManager.switch-buttons', {views: self.view_order}));
    
             // Create bootstrap tooltips
             _.each(this.views, function(view) {
                 self.control_elements.$switch_buttons.siblings('.oe-cp-switch-' + view.type).tooltip();
             });
    
             // Add onclick event listener
             this.control_elements.$switch_buttons.siblings('button').click(function(event) {
                 var view_type = $(event.target).data('view-type');
                 self.switch_mode(view_type);
             });
         }
         */
  }
}

customElements.define('action-viewer', ActionViewer);
