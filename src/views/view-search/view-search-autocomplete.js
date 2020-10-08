import { html, css, LitElement } from 'lit-element';
import { unsafeHTML } from 'lit-html/directives/unsafe-html.js';

class ViewSearchAutocomplete extends LitElement {
  static get styles() {
    return css`
      :host {
        position: absolute;
        top: 100%;
        left: 0;
        z-index: 1;
        /*   display: none;*/
        float: left;
        min-width: 160px;
        padding: 5px 0px;
        margin: 2px 0 0;
        list-style: none;
        font-size: 13px;
        text-align: left;
        background-color: #ffffff;
        border: 1px solid #cccccc;
        border: 1px solid rgba(0, 0, 0, 0.15);
        border-radius: 4px;
        -webkit-box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
        box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
        background-clip: padding-box;
        top: 100%;
        left: auto;
        bottom: auto;
        right: auto;
        width: 100%;
      }

      .item {
        margin: 0 10px;
      }

      a {
        text-decoration: none;
      }

      #tree {
        padding-left: 25px;
        display: contents;
      }

      .o-indent {
        padding-left: 50px;
      }

      a.o-expanded {
        position: absolute;
        top: auto;
        left: 6px;
        bottom: auto;
        right: auto;
        padding: 3px;
      }
    `;
  }

  static get properties() {
    return {
      host: { type: Object }, // # search view
      source: { type: Object }, // # datasource
      select: { type: Array },
      current_search: { type: String }, // # 当前查询的字符串
      search_string: { type: String },
      get_search_string: { type: Object }, // # 获取Input值函数
      filters: { type: Array }, // # filters指下来菜单上可选的过滤条件
      searching: { type: Boolean },
      _selectedItem: { type: String, notify: true },
      //_items: { type: Object },
    };
  }

  render() {
    return html`
      <span
        id="tree"
        role="listbox"
        on-touchend="${this._preventDefault}"
        selected-item="${this._selectedItem}"
        selection-enabled
      >
        <!-- 菜单-->
        ${this.filters
          ? html`
              ${this.filters.map(
                (item, index) => html`
                  <div class="${this.getClass()}" id="it${index}">
                    <!--展开按钮-->
                    ${item && item.expand
                      ? html`
                          <a
                            class="o-expand"
                            href="#"
                            @on-tap="${this.tap_autocomplete_expand}"
                          ></a>
                        `
                      : html``}
                    <a
                      href="#"
                      .idx="${index}"
                      @tap="${this.tap_autocomplete_item}"
                    >
                      ${unsafeHTML(item.label)}</a
                    >
                  </div>
                `
              )}
            `
          : html``}
      </span>
    `;
  }

  constructor() {
    super();
    this.hidden = true; // # 隐藏
    this.host = this.domHost; // Polymer.dom(this).parentNode;              // this.source = options.source;
    this.current_result = null;
    this.searching = true;
    this.search_string = '';
    this.current_search = null;
  }

  // # 查询字符串
  initiate_search(query) {
    if (query === this.search_string && query !== this.current_search) {
      this.search(query);
    }
  }

  // 开始查询
  search(query) {
    this.current_search = query;
    // # 发动搜素事件让Host监听并返回results
    this.dispatchEvent(
      new CustomEvent('queryChanged', {
        detail: { term: query },
        bubbles: true,
        composed: true,
      })
    );
  }

  select_item(ev) {
    if (this.current_result.facet) {
      this.select(ev, { item: { facet: this.current_result.facet } });
      this.close();
    }
  }
  /*
            render_search_results(results) {
                var self = this;
                var $list = this.$el;
                $list.empty();
                results.forEach(function (result) {
                    var $item = self.make_list_item(result).appendTo($list);
                    result.$el = $item;
                });
                this.show();
        )
     
            make_list_item(result) {
                var self = this;
                var $li = $('<li>')
                    .hover(function () { self.focus_element($li); })
                    .mousedown(function (ev) {
                        if (ev.button === 0) { // left button
                            self.select(ev, { item: { facet: result.facet } });
                            self.close();
                        }
                    })
                    .data('result', result);
                if (result.expand) {
                    var $expand = $('<a class="o-expand" href="#">').appendTo($li);
                    $expand.mousedown(function (ev) {
                        ev.stopPropagation();
                        ev.preventDefault(); // to prevent dropdown from closing
                        if (result.expanded) {
                            self.fold();
                        } else {
                            self.expand();
                        }
                    });
                    $expand.click(function (ev) {
                        ev.preventDefault(); // to prevent url from changing due to href="#"
                    });
                    result.expanded = false;
                }
                if (result.indent) $li.addClass('o-indent');
                $li.append($('<a href="#">').html(result.label));
                return $li;
        )
    */

  // # 子菜单张开
  expand() {
    var self = this;
    var current_result = this.current_result;
    current_result.expand(this.get_search_string()).then(function (results) {
      (results || [{ label: '(no result)' }])
        .reverse()
        .forEach(function (result) {
          result.indent = true;
          var $li = self.make_list_item(result);
          current_result.$el.after($li);
        });
      self.current_result.expanded = true;
      self.current_result.$el
        .find('a.o-expand')
        .removeClass('o-expand')
        .addClass('o-expanded');
    });
  }

  // # 子菜单缩回
  fold() {
    var $next = this.current_result.$el.next();
    while ($next.hasClass('o-indent')) {
      $next.remove();
      $next = this.current_result.$el.next();
    }
    this.current_result.expanded = false;
    this.current_result.$el
      .find('a.o-expanded')
      .removeClass('o-expanded')
      .addClass('o-expand');
  }

  // # 显示
  show() {
    this.show();
  }

  // # 关闭显示
  close() {
    this.current_search = null;
    this.search_string = '';
    this.searching = true;
    this.hidden = true;
  }

  is_expandable() {
    // # 是否已经展出
    return !!this.$$('.o-selection-focus .o-expand').length;
  }

  // # 菜单点击响应
  tap_autocomplete_item(ev) {
    /*       if (ev.button === 0) { // left button
                    self.select(ev, {item: {facet: result.facet}});
                    self.close();
                }*/
    // # 触发选择事件
    this.dispatchEvent(
      new CustomEvent('selected', {
        detail: { item: { facet: this.filters[ev.currentTarget.idx].facet } },
        bubbles: true,
        composed: true,
      })
    );
    this.close();
  }

  // # 菜单张开按钮事件
  tap_autocomplete_expand(ev) {
    /*       ev.stopPropagation();
                ev.preventDefault(); // to prevent dropdown from closing
                if (result.expanded) {
                    self.fold();
                } else {
                    self.expand();
                }*/
  }

  // change filters
  set filters(filters) {
    const oldValue = this._filters;
    this._filters = filters;
    this.requestUpdate('filters', oldValue);

    if (filters.length > 0) {
      this.hidden = false;
    } else {
      this.hidden = true;
    }
  }

  get filters() {
    return this._filters;
  }

  getClass() {
    if (this.indent) {
      return 'item o-indent';
    }

    return 'item';
  }
}

customElements.define('view-search-autocomplete', ViewSearchAutocomplete);
