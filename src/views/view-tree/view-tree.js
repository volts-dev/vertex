// lib
import '@vaadin/vaadin-grid/vaadin-grid-selection-column.js';
import '@vaadin/vaadin-grid/vaadin-grid.js';

import '../../element/dataset/data-source.js';
import './view-tree-icon.js';
import './view-tree-buttons.js';
import './view-tree-sidebar.js';
import './view-tree-pager.js';
import { View } from '../view.js';
import { define, Element, html, css } from '../../core/element';
import { OSV } from '../../core/osv';

import { TreeController } from './view-tree-controller';
import { TreeModel } from './view-tree-model';
import { TreeRenderer } from './view-tree-renderer';

class ViewTree extends View {
  static get styles() {
    return css`
      :host {
        position: relative;
        display: block;
      }

      #grid {
        --vaadin-grid-body-row-odd-cell: {
          background-color: #efeff8;
        }
      }
      [role='columnheader'] {
        min-height: auto;
      }
    `;
  }

  render() {
    return html`
      <!--data-source id="datasource" action="/dataset/search_read" method="POST" .params="${this
        .params}" .data="${this.data}"></data-source!-->

      <vaadin-grid id="grid" .items="${this.data}">
        <!-- 选择框 -->
        <vaadin-grid-selection-column
          auto-select
        ></vaadin-grid-selection-column>

        <!-- 列 -->
        ${this.fields
          ? html`
              ${this.fields.map(
                field => html`
                  ${field.sortable
                    ? html`
                        <vaadin-grid-sort-column
                          path="${field.name}"
                          header="${field.string}"
                        ></vaadin-grid-sort-column>
                      `
                    : html`
                        <vaadin-grid-column
                          path="${field.name}"
                          header="${field.string}"
                        ></vaadin-grid-column>
                      `}
                `
              )}
            `
          : html` <p>have not fields</p> `}
      </vaadin-grid>
    `;
  }

  static get properties() {
    return {
      accesskey: { type: String, value: 'l' },
      display_name: { type: String, value: OSV._lt('List') },
      icon: { type: String, value: 'lfa-list-ul' },
      icon: { type: String, value: 'list' },
      config: {
        type: Object,
        value: _.extend({}, View.prototype.config, {
          Model: TreeModel,
          Renderer: TreeRenderer,
          Controller: TreeController,
        }),
      },

      grid: { type: Object },
      datasource: { type: Object },
      datatable: { type: Object },
      viewMgr: { type: Object },
      query: { type: Object },
      model: { type: String },
      data: { type: Object },
      params: { type: Object },

      // 被选择的项目
      selectedItems: { type: Array },
      fields: { type: Array, value: [] },
      page: Number, // 当前页
      pages: Array, // 所有页序号 1，2，3...
      pageSize: { type: Number, value: 25 },
      nodeBreadcrumb: Object,
      nodeButtons: Object,
      nodeSidebar: Object,
      nodePager: Object,
    };
  }

  constructor() {
    super(...arguments);
    this.searchable = true; // 控制面板带搜索
  }

  firstUpdated() {
    super.firstUpdated();

    this.query = {};
    this.grid = this.shadowRoot.querySelector('#grid');
    for (var i = 1; i < this.children.length - 1; i++) {
      var el = this.children[i];
      if (el.tag === 'field') {
        // 新对象
        var field = new Object();
        field.element = el;
        field.text = this.children[i].getAttribute('string');
        field.field = this.children[i].getAttribute('name');
        //  self.search_fields.push(lObj);
      }

      //  if (filter.item.tag === 'filter') {
      //    current_group.push(new search_inputs.Filter(filter.item, this));
      //  }
    }

    this.grid.addEventListener(
      'active-item-changed',
      this.selectedItemChanged.bind(this)
    );
    /*
    this.datasource.addEventListener(
      "onResponse",
      function(e) {
        this.data = e.detail;
      }.bind(this)
    );
    */
  }

  selectedItemChanged(e) {
    const item = e.detail.value;
    if (item) {
      // 更新选择项
      this.selectRecord(item.id, 'form');
    }
  }

  /**
   * Used to handle a click on a table row, if no other handler caught the
   * event.
   *
   * The default implementation asks the list view's view manager to switch
   * to a different view (by calling
   * :js:func:`~instance.web.ViewManager.on_mode_switch`), using the
   * provided record index (within the current list view's dataset).
   *
   * If the index is null, ``switch_to_record`` asks for the creation of a
   * new record.
   *
   * @param {Number|void} index the record index (in the current dataset) to switch to
   * @param {String} [view="page"] the view type to switch to
   */
  selectRecord(recordId, view) {
    var view =
      view || recordId === null || recordId === undefined ? 'form' : 'form';

    // 跟新URL地址参数信息
    var query = {};
    query['id'] = recordId;
    query['view'] = view;

    ///this.controlPanel.SetQuery(query); // set id first
    ///this.controlPanel.SetViewMode(view); // active mode
    ///this.controlPanel.UpdateQuery(query);
    this.app.router.SetQuery(query);
  }

  /** 添加新纪录
   * Handles signal for the addition of a new record (can be a creation,
   * can be the addition from a remote source, ...)
   *
   * The default implementation is to switch to a new record on the form view
   */
  do_add_record() {
    this.selectRecord(null);
  }

  renderHeader(parent) {
    // 添加导航标题
    if (parent.children.length == 0) {
      var li = document.createElement('li');
      li.innerText = this.controlPanel.action.name;
      parent.appendChild(li);
    }
  }
  /*
  renderSearchview(parent) {
    if (parent.firstChild) {
      parent.firstChild.hidden = false;
    }

    parent.innerHTML = this.viewMgr.action.arch;
  }
*/
  renderSidebar(parent) {
    if (!this.nodeSidebar) {
      this.nodeSidebar = document.createElement('view-tree-sidebar');
    }

    parent.appendChild(this.nodeSidebar);
  }

  // # 控制板按钮事件banding
  on_buttons_changed(e) {
    //lv=Polymer.dom(e.currentTarget).querySelector("div");
    //lv=e.currentTarget.querySelector(".form_button_create");
    //this.listen(e.currentTarget.querySelector(".list_button_add"), 'click', "do_add_record"); // 监听Add 按钮
    var node = e.currentTarget.querySelector('.list_button_add');
    if (node) {
      node.addEventListener('click', this.do_add_record.bind(this));
    }

    //  this.listen(e.currentTarget.querySelector(".form_button_edit"), 'click', "on_button_edit");
    //  this.listen(e.currentTarget.querySelector(".form_button_save"), 'click', "on_button_save");
    //  this.listen(e.currentTarget.querySelector(".form_button_cancel"), 'click', "on_button_cancel");
  }

  renderButtons(parent) {
    //buttons.innerHTML=document.createElement("view-form-buttons").outerHTML;
    // 清空节点
    /*   while (buttons.children.length > 0) {
           buttons.removeChild(buttons.firstChild);
       }*/
    if (!this.nodeButtons) {
      this.nodeButtons = document.createElement('view-tree-buttons');
      this.nodeButtons.action_buttons = true;
      //this.listen(this.nodeButtons, 'dom-change', "on_buttons_changed");
      this.nodeButtons.addEventListener(
        'dom-change',
        this.on_buttons_changed.bind(this)
      );
    }

    //lButtons=lNode.querySelector('div[class*="form-buttons"]');
    parent.appendChild(this.nodeButtons);
  }

  on_page_changed(e) {
    var page = Math.max(0, e.detail - 1); // 避免-1页
    this.page = page;
    var start = page * this.pageSize;
    var end = (page + 1) * this.pageSize;
    this.grid.items = this.data.slice(start, end);
  }

  renderPager(parent) {
    if (!this.nodePager) {
      this.nodePager = document.createElement('view-tree-pager');
      //this.listen(this.nodePager, 'page-changed', "on_page_changed");
      //lButtons=lNode.querySelector('div[class*="form-buttons"]');
      this.nodePager.addEventListener(
        'page-changed',
        this.on_page_changed.bind(this)
      );
    }
    parent.appendChild(this.nodePager);
  }
  // 注册事件
  //  listeners: {
  //    'show': '_showData',
  //    'iron-resize': '_resizeListener'
  //  }
  show(mgr) {
    super.show();
    var self = this;

    ///this.controlPanel.UpdateQuery({"id":undefined})
    this.app.router.SetQuery({ id: undefined });

    /*
    //var table = this.datatable;
    // var fields = this.queryAllEffectiveChildren('field');
    var fields = this.querySelectorAll("field");
    var columns = [];
    fields.forEach(field => {
      var lFieldName = field.getAttribute("name"); // 字段名
      var lFieldRec = this.viewMgr.fields[lFieldName]; // View 里的Field记录
      if (lFieldRec) {
        columns.push(lFieldRec);
      }
    });

    this.fields = columns;*/

    //根据lViewMgr组织查询参数 {"view_id":[[viewId]],"model=":"[[action.res_model]]","view_type":"[[mode]]"}

    // 执行API
    this.datasource.action = '/dataset/search_read';
    this.datasource.params = this.params;
    this.datasource.read().then(function (data) {
      self.data = data;
    });
  }

  get data() {
    return this._data;
  }

  set data(data) {
    const oldValue = this._data;
    this._data = data;
    this.requestUpdate('data', oldValue);
    // 生成页序列号
    this.pages = Array.apply(null, {
      length: Math.ceil(data.length / this.grid.pageSize),
    }).map(function (item, index) {
      return index + 1;
    });

    // 更新
    if (this.nodePager) {
      this.nodePager.pages = this.pages;
      this.nodePager.size = data.length;
    }
  }

  valueByField(item, field) {
    if (item && field) {
      return item[field.name];
    }
  }

  _queryAndSetFields() {
    var fields = this.queryAllEffectiveChildren('field');
    var table = this.grid;
    fields.forEach(field => {
      // create an instance with createElement:
      var column = document.createElement('paper-datatable-column');
      column.header = field.getAttribute('name');
      column.property = field.getAttribute('name');
      column.type = 'string';
      column.tooltip = '';
      //  column.align = 'right';
      column.sortable = true;
      //  column.formatValue = app.toFixedOne;
      //var table = Polymer.dom(this).querySelector('paper-datatable');
      if (table) {
        //  var columns = table.queryAllEffectiveChildren('paper-datatable-column');
        //  var columns=Polymer.dom(document).querySelectorAll('paper-datatable-column');
        //  var cols =table.getEffectiveChildNodes()

        // 必须使用Polymer.dom()接口插入才能实现是Effective Children
        table.appendChild(column);
        //Polymer.dom(table).insertBefore(column, null);
      }
    });
  }
}

customElements.define('view-tree', ViewTree);
