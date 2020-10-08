import { html, css } from 'lit-element';
import { View } from '../view.js';
import { viewUtils } from '../view_utils';
import { pyUtils } from '@/core/py_utils';

//import "../../core/core.js";
import { FieldBool } from './field-bool.js';
import { FieldChar } from './field-char.js';
import { FieldInt, FieldFloat } from './field-int.js';
import { FieldSelection } from './field-selection.js';
import './view-search-facet.js';
import './view-search-input.js';
import './view-search-autocomplete.js';

/*
var list1 = [[0, 1], [2, 3], [4, 5]];
var list2 = [0, [1, [2, [3, [4, [5]]]]]];
flatten(list1); // returns [0, 1, 2, 3, 4, 5]
flatten(list2); // returns [0, 1, 2, 3, 4, 5]
*/
const flatten = arr =>
  arr.reduce((a, b) => a.concat(Array.isArray(b) ? flatten(b) : b), []);

// # 删除false,0
const compact = arr =>
  arr.reduce(function (a, b) {
    if (a != null && b != null) {
      return a.concat(b);
    }
    return a;
  }, []);

export class ViewSearch extends View {
  static get styles() {
    return css`
      :host {
        /*# 对于子元素Position定位关系非常重要*/
        position: relative;
        font-family: Roboto;
        line-height: 25px;
        padding: 0 30px 1px 0;
        display: block;
        background-color: #fff;
        border-bottom: 1px solid #afafb6;
      }

      .container-fluid {
        width: 100%;
        margin-right: auto;
        margin-left: auto;
        padding-left: 15px;
        padding-right: 15px;
      }

      .searchview {
        display: -ms-flexbox;
        display: -moz-box;
        display: -webkit-box;
        display: -webkit-flex;
        display: flex;
        flex-wrap: wrap;
      }

      .searchview_more {
        position: absolute;
        padding: 0;
        margin: 0;
        top: auto;
        left: auto;
        bottom: auto;
        right: 5px;
        font-size: 16px;
        cursor: pointer;
      }

      /**** */
      :host(.search-on) {
        left: 0;
        background: inherit;
        z-index: 1001;
      }

      .breadcrumb {
        font-size: 18px;
        margin: 0;
      }

      .active {
        color: #8f8f8f;
      }

      #app-view-control-panel {
        position: relative;
      }

      #app-view-control-panel iron-icon {
        margin-right: 0;
      }

      #search[show] {
        position: absolute;
        width: 100%;
        height: 100%;
        left: 0;
        top: 0;
        padding: 0 16px;
        background: #fff;
      }

      #search input {
        display: none;
        font-family: var(--primary-font-family);
        font-size: 15px;
        width: 100%;
        padding: 10px;
        border: 0;
        border-radius: 2px;
        -webkit-appearance: none;
      }

      #search[show] input {
        display: block;
      }

      #search input:focus {
        outline: 0;
      }

      .control_panel {
        display: -ms-flexbox;
        display: -moz-box;
        display: -webkit-box;
        display: -webkit-flex;
        display: flex;
        -ms-flex-flow: row wrap;
        -moz-flex-flow: row wrap;
        -webkit-flex-flow: row wrap;
        flex-flow: row wrap;
        flex: 0 0 auto;
      }

      .row {
        margin-left: -15px;
        margin-right: -15px;
      }

      .col {
        width: 50%;
        float: left;
      }

      .search-options {
        position: relative;
        display: inline-block;
        vertical-align: middle;
      }

      .right-toolbar {
        float: right;
      }

      .right-toolbar > div {
        display: inline-block;
      }

      .fa {
        /* display: inline-block; */
        font: normal normal normal 14px/1 FontAwesome;
        font-size: inherit;
        text-rendering: auto;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
      }
    `;
  }

  render() {
    return html`
      <style is="custom-style" include="iron-flex iron-flex-alignment"></style>
      <div
        t-name="SearchView"
        class="searchview form-control input-sm horizontal layout center"
      >
        <!--自动完成菜单-->
        <view-search-autocomplete
          class="dropdown-menu searchview_autocomplete"
          .host="${this}"
          .source="${this.source}"
        ></view-search-autocomplete>
        <!--放大镜-->
        <iron-icon
          icon="search"
          class="searchview_more fa fa-search-minus"
          title="Advanced Search..."
          @click="${this.moreTapped}"
        ></iron-icon>
        <div class="searchview_facets"></div>
        <slot></slot>
      </div>
    `;
  }

  static get properties() {
    return {
      // # 查询条件 Facets 组成
      // odoo 里的 SearchQuery
      queryFacets: { type: Array },
      viewMgr: { type: Object },
      autocomplete: { type: Object }, // # 自动完成菜单

      input_subviews: { type: Array }, // # 存储facet，input视图 提供清除

      input_node: { type: Object },
      action: { type: Object },
      modes: { type: Object },
      string: { type: String },
      title: { type: String },
      search_fields: { type: Array }, // 可查询的字段
      filters: { type: Array }, // 可过滤的字段

      groupbys: { type: Array }, // 可排序的字段
      filter_menu: { type: Object }, // 菜单对象
      groupby_menu: { type: Object }, // 菜单对象
      favorite_menu: { type: Object }, // 常用的配置
    };
  }

  firstUpdated() {
    super.firstUpdated();
    this.widgets_registry = new Registry();
    this.widgets_registry
      .add('bool', FieldBool)
      .add('char', FieldChar)
      .add('text', FieldChar)
      .add('html', FieldChar)
      .add('integer', FieldInt)
      .add('id', FieldInt)
      .add('float', FieldFloat)
      .add('monetary', FieldFloat)
      .add('selection', FieldSelection);

    // # 基础数组数据
    this.search_fields = []; // # 可查询字段
    this.filters = []; // # 可查询组条件
    this.groupbys = []; // # 可查询组条件
    this.queryFacets = [];

    this.filter_menu = undefined;
    this.groupby_menu = undefined;
    this.favorite_menu = undefined;
    this.visible_filters = true;

    if (!this.fields) {
      //this.fields = Polymer.dom(this).querySelectorAll('*[field]:not(field)');
      // 使用  Array.prototype.slice.call(); 转换到Arry
      this.fields = Array.prototype.slice.call(
        this.shadowRoot.querySelectorAll('*[name]:not(field)')
      );
    }

    // # 监听 Autocomplete 事件
    this.autocomplete = this.shadowRoot.querySelector(
      'view-search-autocomplete'
    );
    this.addEventListener('queryChanged', this.complete_global_search); // # 搜索输入字符串事件
    this.addEventListener('selected', this.select_completion); // # 选择事件
    this.addEventListener('keyup', this.keyupHandller); // # 监听Host事件
    this.addEventListener('removeFacet', this.onRemoveFacet);
  }

  // 显示
  show(mgr) {
    this.viewMgr = mgr;
    this.modes = this.viewMgr.action.view_mode.split(',');

    // # 默认More按钮为状态
    this.toggle_visibility(false);
    // # 创建Auto Complete
    this.setup_global_completion();

    // # 创建Query
    //  this.query = new SearchQuery()
    //  .on('add change reset remove', this.proxy('do_search'))
    //  .on('change', this.proxy('renderChangedFacets'))
    //  .on('add reset remove', this.proxy('renderFacets'));

    // # 初始化View的Field
    this.prepare_search_inputs();

    this.facets_container = this.shadowRoot.querySelector(
      'div.searchview_facets'
    );
    //  this.prepare_search_inputs().bind(this);
    // 获取parent
    //this.viewMgr = this.parentElement;
    //this.action = this.viewMgr.action;

    /*
            if (this.$buttons) {
                if (!this.options.disable_filters) {
                    this.filter_menu = new FilterMenu(this, this.filters);
                    menu_defs.push(this.filter_menu.appendTo(this.$buttons));
                }
                if (!this.options.disable_groupby) {
                    this.groupby_menu = new GroupByMenu(this, this.groupbys);
                    menu_defs.push(this.groupby_menu.appendTo(this.$buttons));
                }
                if (!this.options.disable_favorites) {
                    this.favorite_menu = new FavoriteMenu(this, this.query, this.dataset.model, this.action_id, this.favorite_filters);
                    menu_defs.push(this.favorite_menu.appendTo(this.$buttons));
                }
            }
            return $.when.apply($, menu_defs).then(this.set_default_filters.bind(this));
    */
    this.set_default_filters();
  }

  /**
   * Given a <searchpanel> arch node, iterate over its children to generate the
   * description of each section (being either a category or a filter).
   *
   * @param {Object} node a <searchpanel> arch node
   * @param {Object} fields the fields of the model
   * @returns {Object}
   */
  _processSearchPanelNode(node, fields) {
    var sections = {};
    node.children.forEach((childNode, index) => {
      if (childNode.tag !== 'field') {
        return;
      }
      if (childNode.attrs.invisible === '1') {
        return;
      }
      var fieldName = childNode.attrs.name;
      var type = childNode.attrs.select === 'multi' ? 'filter' : 'category';

      var sectionId = _.uniqueId('section_');
      var section = {
        color: childNode.attrs.color,
        description: childNode.attrs.string || fields[fieldName].string,
        fieldName: fieldName,
        icon: childNode.attrs.icon,
        id: sectionId,
        index: index,
        type: type,
      };
      if (section.type === 'category') {
        section.icon = section.icon || 'fa-folder';
      } else if (section.type === 'filter') {
        section.disableCounters = !!pyUtils.py_eval(
          childNode.attrs.disable_counters || '0'
        );
        section.domain = childNode.attrs.domain || '[]';
        section.groupBy = childNode.attrs.groupby;
        section.icon = section.icon || 'fa-filter';
      }
      sections[sectionId] = section;
    });
    return sections;
  }

  /**
   * Parse a given search view arch to extract the searchpanel information
   * (i.e. a description of each filter/category). Note that according to the
   * 'view_types' attribute on the <searchpanel> node, and the given viewType,
   * it may return undefined, meaning that no searchpanel should be rendered
   * for the current view.
   *
   * Note that this is static method, called by AbstractView, *before*
   * instantiating the SearchPanel, as depending on what it returns, we may
   * or may not instantiate a SearchPanel.
   *
   * @static
   * @params {Object} viewInfo the viewInfo of a search view
   * @params {string} viewInfo.arch
   * @params {Object} viewInfo.fields
   * @params {string} viewType the type of the current view (e.g. 'kanban')
   * @returns {Object|undefined}
   */
  computeSearchPanelParams(viewInfo, viewType) {
    var searchPanelSections;
    var classes;
    if (viewInfo) {
      var arch = viewUtils.parseArch(viewInfo.arch);
      viewType = viewType === 'list' ? 'tree' : viewType;
      arch.children.forEach(function (node) {
        if (node.tag === 'searchpanel') {
          var attrs = node.attrs;
          var viewTypes = defaultViewTypes;
          if (attrs.view_types) {
            viewTypes = attrs.view_types.split(',');
          }
          if (attrs.class) {
            classes = attrs.class.split(' ');
          }
          if (viewTypes.indexOf(viewType) !== -1) {
            searchPanelSections = this._processSearchPanelNode(
              node,
              viewInfo.fields
            );
          }
        }
      });
    }
    return {
      sections: searchPanelSections,
      classes: classes,
    };
  }

  keyupHandller(e) {
    //console.log("keyupHandller");
    var self = this.autocomplete;

    switch (e.keyCode) {
      case 8: // backspace
        // 当Input无字符串时开始删除Facet
        if (this.input_node.value === '') {
          //var preceding = this.parentElement.siblingSubview(this, -1);
          var preceding = this.lastElementChild; // # 获取This的Container并选择最后一个Element
          // # 删除获取的Facet
          if (preceding && preceding instanceof FacetView) {
            preceding.model.destroy();
          }
        }
        break;
      // TAB and direction keys are handled at KeyDown because KeyUp
      // is not guaranteed to fire.
      // See e.g. https://github.com/aef-/jquery.masterblaster/issues/13
      case 9: //$.ui.keyCode.TAB:
        var self = this.autocomplete;
        if (self.search_string.length) {
          self.select_item(event);
        }
        break;

      case 13: // enter
        if (self.search_string.length) {
          self.select_item(e);
        }
        return;
      case 16: // shift
      case 17: // ctrl
      case 18: // alt
      case 27: // escape
      case 37: // left arrow // Stop propagation to parent if not at beginning of input value
        // if (this.el.selectionStart > 0) {
        //     e.stopPropagation();
        // }
        break;

      case 38: // up arrow
      case 39: // right arrow // Stop propagation to parent if not at end of input value
        // if (this.el.selectionStart < this.$$(".searchview_input").val().length) {
        //     e.stopPropagation();
        // }
        self.searching = true;
        e.preventDefault();
        return;
        break;
      case 40: // down arrow
    }

    // # 其他执行查询 这里可以执行粘帖的字符
    var search_string = self.get_search_string();
    if (self.search_string !== search_string) {
      if (search_string.length) {
        self.search_string = search_string;
        self.initiate_search(search_string);
      } else {
        self.close();
      }
    }
  }

  /**
   * Provide auto-completion result for req.term (an array to `resp`)
   *
   * @param {Object} req request to complete
   * @param {String} req.term searched term to complete
   * @param {Function} resp response callback
   */
  // # 返回AutoComplete查询条件
  complete_global_search(e) {
    var self = this;

    // # 集合所有查询条件var
    if (!this.search_fields) {
      return;
    }

    var inputs = this.search_fields.concat(this.filters, this.groupbys),
      deferArray = [];

    inputs.forEach(function (input) {
      input.visible(); // # 变更为可视
      // var defer = input['complete'](detail.term); // # 返回延迟
      if (e.detail.term) {
        var defer = input['complete'](e.detail.term); // # 返回延迟
        deferArray.push(defer);
      }
    });

    var promise = Promise.all(deferArray);
    promise.then(function (results) {
      // # filters指下来菜单上可选的过滤条件
      var filters = flatten(results);
      filters = compact(results);

      if (self.autocomplete) {
        // # Autocomplete
        self.autocomplete.filters = filters; // # 更新filters
      }
    });
  }
  /**
   * Action to perform in case of selection: create a facet (model)
   * and add it to the search collection
   *
   * @param {Object} e selection event, preventDefault to avoid setting value on object
   * @param {Object} ui selection information
   * @param {Object} ui.item selected completion item
   */
  // # AutoComplete 选择事件
  select_completion(e) {
    e.preventDefault();
    var ui = e.detail;
    if (
      ui.item.facet.values &&
      ui.item.facet.values.length &&
      String(ui.item.facet.values[0].value).trim() !== ''
    ) {
      // # TODO 过滤重复
      var facets = [];
      var isUpdated;
      if (this.queryFacets) {
        // # 添加合并值到已有的Facet里
        this.queryFacets.forEach(function (facet) {
          if (
            facet['category'] === ui.item.facet['category'] &&
            facet.field === ui.item.facet.field
          ) {
            facet.values = facet.values.concat(ui.item.facet.values); // # 添加Facet新值
            //facet.values = flatten(facet.values);
            isUpdated = true; // 标记为更新
          }

          // 保存
          facets.push(facet);
        });
      }

      // 添加新的facet
      if (!isUpdated) {
        facets.push(ui.item.facet);
      }
      // mutate the array
      this.updateFacets(facets);
    } else {
      //this.query.trigger('add');
    }
  }

  toggle_visibility(is_visible) {
    /* this.do_toggle(!this.headless && is_visible);
     if (this.$buttons) {
         this.$buttons.toggle(!this.headless && is_visible && this.visible_filters);
     }
     if (!config.device.touch && config.device.size_class >= config.device.SIZES.SM) {
         this.$('input').focus();
     }*/
  }

  /**
   * Sets up search view's view-wide auto-completion widget
   */
  // # 设置自动完成下拉选项控件
  setup_global_completion() {
    var self = this;
    this.autocomplete.get_search_string = function () {
      return self.input_node.value.trim();
    };

    this.autocomplete.view = this;
  }

  // # 更多搜索
  moreTapped(e) {
    e.target.classList.add('fa-search-plus fa-search-minus');
    var visible_search_menu =
      localStorage.getItem('visible_search_menu') !== 'true';
    localStorage.setItem('visible_search_menu', visible_search_menu);
    this.toggle_buttons();
  }

  toggle_buttons(is_visible) {
    this.visible_filters = is_visible || !this.visible_filters;
    if (this.$buttons) {
      this.$buttons.toggle(this.visible_filters);
    }
  }

  /**
   * Extract search data from the view's facets.
   *
   * Result is an object with 3 (own) properties:
   *
   * domains
   *     Array of domains
   * contexts
   *     Array of contexts
   * groupbys
   *     Array of domains, in groupby order rather than view order
   *
   * @return {Object}
   */
  buildSearchData() {
    var domains = [],
      contexts = [],
      groupbys = [];

    if (this.queryFacets) {
      this.queryFacets.forEach(function (facet) {
        var field = facet['field'];
        var domain = field.get_domain(facet);
        if (domain) {
          domains.push(domain);
        }
        var context = field.get_context(facet);
        if (context) {
          contexts.push(context);
        }
        var group_by = field.get_groupby(facet);
        if (group_by) {
          groupbys.push.apply(groupbys, group_by);
        }
      });
    }

    return {
      domains: domains,
      contexts: contexts,
      groupbys: groupbys,
    };
  }

  /**
   * Performs the search view collection of widget data.
   *
   * If the collection went well (all fields are valid), then triggers
   * :js:func:`instance.web.SearchView.on_search`.
   *
   * If at least one field failed its validation, triggers
   * :js:func:`instance.web.SearchView.on_invalid` instead.
   *
   * @param [_query]
   * @param {Object} [options]
   */
  doSearch() {
    var search = this.buildSearchData();
    // # 触发search_data 事件 提供Form,View监听
    this.dispatchEvent(
      new CustomEvent('search_data', {
        detail: search,
        bubbles: true,
        composed: true,
      })
    );
  }

  /**
   * @param {openerp.web.search.SearchQuery | undefined} Undefined if event is change
   * @param {openerp.web.search.Facet}
   * @param {Object} [options]
   */
  // # 更新所有Facets和Input输入框
  updateFacets(facets) {
    this.queryFacets = facets;
    var self = this;
    var started = [];

    // # 销毁视图
    // _.invoke(this.input_subviews, 'destroy');
    if (self.input_subviews) {
      self.input_subviews.forEach(function (input) {
        // input.remove();
        self.removeChild(input);
      });
    }
    self.input_subviews = [];

    // # 重新建立
    if (self.queryFacets) {
      self.queryFacets.forEach(function (facet) {
        //var f = new FacetView(this, facet);
        var f = document.createElement('view-search-facet'); // # 创建Facet视图
        f.host = self;
        f.facet = facet;
        //f.requestUpdate("facet",f.facet);

        // var f = new ViewSearchFacet(self, facet);
        self.appendChild(f); // # 插入Search
        started.push(f);
        self.input_subviews.push(f);
      }, self);
    }

    var i = document.createElement('view-search-input'); // #创建Input视图
    self.input_node = i;
    self.appendChild(i); // # 插入Search
    started.push(i);
    self.input_subviews.push(i);

    /* _.each(this.input_subviews, function (childView) {
         childView.on('focused', self, self.proxy('childFocused'));
         childView.on('blurred', self, self.proxy('childBlurred'));
     });

     // # 等等所有View载入完毕设置最后一个为焦点通常是Input
     $.when.apply(null, started).then(function () {
         _.last(self.input_subviews).$el.focus();
     });*/
    this.doSearch();
  }

  // query 数组变动
  // @ 搜索
  // @ 更新Facets
  set query(query) {
    const oldValue = this._query;
    this._query = query;
    this.requestUpdate('query', oldValue);

    //if (this.query) {
    // this.renderFacets();
    //this.doSearch(); //query.base
    //}
  }

  get query() {
    return this._query;
  }

  // it should parse the arch field of the view, instantiate the corresponding
  // filters/fields, and put them in the correct variables:
  // * this.search_fields is a list of all the fields,
  // * this.filters: groups of filters
  // * this.group_by: group_bys
  prepare_search_inputs() {
    var self = this;
    function eval_item(item) {
      var category = 'filters';
      if (item.attributes.context) {
        try {
          var context = pyeval.eval('context', item.attributes.context);
          if (context.group_by) {
            category = 'group_by';
          }
        } catch (e) {}
      }
      return {
        item: item,
        category: category,
      };
    }

    var filters = []; //Array.prototype.concat.apply([], arr);
    var current_group = [],
      current_category = 'filters',
      categories = { filters: this.filters, group_by: this.groupbys };

    // 遍历所有
    for (var i = 0; i < this.children.length; i++) {
      var item = this.children[i];
      if (item.tagName && item.tagName.toLowerCase() === 'group') {
        for (var n = 0; n < item.children.length; n++) {
          filters.push(eval_item(item.children[n]));
        }
      } else {
        filters.push(eval_item(item));
      }
    }

    // # 遍历过滤对象表
    filters
      .concat({ category: 'filters', item: 'separator' })
      .forEach(function (filter) {
        // 创建Filter元素
        if (
          filter.item.tagName &&
          filter.item.tagName.toLowerCase() === 'filter' &&
          filter.category === current_category
        ) {
          //* return current_group.push(new search_inputs.Filter(filter.item, self));
        }

        if (current_group.length) {
          var group = new search_inputs.FilterGroup(current_group, self);
          categories[current_category].push(group); // # 保存过滤组
          current_group = []; // # 清空
        }

        // #　字段查询
        if (
          filter.item.tagName &&
          filter.item.tagName.toLowerCase() === 'field'
        ) {
          var attrs = filter.item.attributes; // # 字段属性
          var field = self.viewMgr.fields[attrs.name.value]; //获取View的字段
          if (field) {
            // M2O combined with selection widget is pointless and broken in search views,
            // but has been used in the past for unsupported hacks -> ignore it
            if (field.type === 'many2one' && attrs.widget === 'selection') {
              attrs.widget = undefined;
            }

            // # 获取字段对应的Input控件
            var Obj = Vectors.Core.search_widgets_registry.get_any([
              attrs.widget,
              field.type,
            ]);
            if (Obj) {
              // 创建字段对象
              var ele = new Obj(filter.item, field, self);
              ele.host = self;

              // 获取属性
              var attr = {};
              for (var att in field) {
                attr[att] = field[att];
              }
              // # 遍历Field Html节点
              for (var key in filter.item.attributes) {
                attr[filter.item.attributes[key].name] =
                  filter.item.attributes[key].value;
              }

              // 添加属性
              ele.loadAttrs(attr);

              // 添加
              self.search_fields.push(ele);
            }
          }
        }

        if (
          filter.item.tagName &&
          filter.item.tagName.toLowerCase() === 'filter'
        ) {
          // TODO
        }

        current_category = filter.category;
      });
  }

  // # 设置初始化默认过滤器
  set_default_filters() {
    /*  var self = this,
         default_custom_filter = this.$buttons && this.favorite_menu.get_default_filter();
       if (!self.options.disable_custom_filters && default_custom_filter) {
         return this.favorite_menu.toggle_filter(default_custom_filter, true);
       }
 
       if (!_.isEmpty(this.search_defaults)) {
         var inputs = this.search_fields.concat(this.filters, this.groupbys),
           search_defaults = _.invoke(inputs, 'facet_for_defaults', this.search_defaults);
         return $.when.apply(null, search_defaults).then(function () {
           self.query.reset(_(arguments).compact(), { preventSearch: true });
         });
       }
 */

    this.updateFacets([]);
    return; //$.when();
  }

  ___onModeBtnClick(e) {
    //lControlPanel = this.parentElement;
    //this.viewMgr.mode= event.currentTarget.modeType;
    this.viewMgr.SetViewMode(event.currentTarget.modeType);
    // mode-type 属性传递给ControlPanel
    //  var mgr = document.querySelector('app-view-manager');
    //  mgr.mode = event.currentTarget.modeType;
  }

  toggleSearch(e) {
    if (e) {
      e.stopPropagation();
    }
    if (e.target === this.$.query) {
      return;
    }

    this.showingSearch = !this.showingSearch;
  }

  clearSearch() {
    this.showingSearch = false;
  }

  hotkeys(e) {
    // ESC
    if (e.keyCode === 27 && e.rootTarget === this.$.query) {
      this.showingSearch = false;
    }
  }

  // 移除facet
  removeFacet(facet) {
    var index = this.queryFacets.indexOf(facet);
    var newq = [];
    newq.concat(this.queryFacets.splice(index, 1));
    this.updateFacets(newq);
  }

  onRemoveFacet(e) {
    // # --删除按钮--
    if (e.target) {
      this.removeFacet(e.target.facet);
    }

    //this.focus()
    e.stopPropagation();
  }
}

customElements.define('view-search', ViewSearch);
