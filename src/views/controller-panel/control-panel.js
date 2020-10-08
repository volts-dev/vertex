import { html, css } from 'lit-element';
import '@/core/pyeval.js';
//import '@/view-search/view-search.js';
//import '@/elements/dataset/data-source.js';
import { viewUtils } from '../view_utils';
import { breadcrumbStyles } from './styles-breadcrumb.js';
import { AbstractView } from '@/views/abstract_view';
import { ControlPanelController } from './control-panel-controller';
import { ControlPanelModel } from './control-panel-model';
import { ControlPanelRenderer } from './control-panel-renderer';
import controlPanelViewParameters from './control_panel_parameters';
import SearchBar from './search-bar';

var DEFAULT_INTERVAL = controlPanelViewParameters.DEFAULT_INTERVAL;
var DEFAULT_PERIOD = controlPanelViewParameters.DEFAULT_PERIOD;
var INTERVAL_OPTIONS = controlPanelViewParameters.INTERVAL_OPTIONS;
const OPTION_GENERATORS = controlPanelViewParameters.OPTION_GENERATORS;

export class ControlPanel extends AbstractView {
  static get styles() {
    return [breadcrumbStyles, css``];
  }

  render() {
    return html`
      <div class="o_control_panel">
        <ol class="breadcrumb" role="navigation">
          ${this.withBreadcrumbs ? html`${this._renderBreadcrumbs}` : null}
        </ol>
        <div class="o_cp_searchview" role="search" />
        <div class="o_cp_left">
          <div
            class="o_cp_buttons"
            role="toolbar"
            aria-label="Control panel toolbar"
          />
          <aside class="o_cp_sidebar" />
        </div>
        <div class="o_cp_right">
          <div class="btn-group o_search_options" role="search" />
          <nav class="o_cp_pager" role="search" aria-label="Pager" />
          <nav
            class="btn-group o_cp_switch_buttons"
            role="toolbar"
            aria-label="View switcher"
          />
        </div>
      </div>
    `;
  }

  static get properties() {
    return {
      //app: { type: Object },
      //controlPanel: { type: Object }, // 控制面板
      //viewManager: { type: Object }, // 视图控制器
      viewMode: { type: String }, // 显示模式 Can be "grid", "form", "create" or undefined.
      action: { type: Object },
      views: { type: Object }, // 授权访问的视图集

      query: { type: Object }, // URL查询参数对象
      //fields: { type: Object }, // 授权访问的字段
      //searchView: { type: Object }, // 搜索视图
      modes: { type: Object },
      controlElements: { type: Object }, // 保存控制器元素，Pager,Actions,sidebar
      modeOfPage: { type: Object, notify: true }, // mode 对应的Page idex
      viewByMode: { type: Object },
      viewIdOfMode: { type: Object, notify: true }, // mode 对应的 View ID
      defaultView: { type: String, notify: true }, // # action 的默认模板

      // 过滤器环境参数
      domains: { type: Array },
      contexts: { type: Array },
      groupbys: { type: Array },

      showingSearch: {
        type: Boolean,
        value: false,
      },
    };
  }

  constructor(params) {
    super();
    params = params || {};

    this.config = _.extend({}, AbstractView.prototype.config, {
      Controller: ControlPanelController,
      Model: ControlPanelModel,
      Renderer: ControlPanelRenderer,
    });
    //renderer
    this.menusSetup = false;
    this.searchMenuTypes = params.searchMenuTypes || [];
    this.subMenus = {};

    var self = this;
    var viewInfo = params.viewInfo || { arch: '<search/>', fields: {} };
    var context = _.extend({}, params.context);
    var domain = params.domain || [];
    var action = params.action || {};

    this.searchDefaults = {};
    Object.keys(context).forEach(function (key) {
      var match = /^search_default_(.*)$/.exec(key);
      if (match) {
        self.searchDefaults[match[1]] = context[key];
        delete context[key];
      }
    });

    this.arch = viewUtils.parseArch(viewInfo.arch);
    this.fields = viewInfo.fields;

    // 时间
    this.referenceMoment = moment();

    // setDescriptions 函数
    const setDescriptions = options => {
      return options.map(o => {
        const oClone = JSON.parse(JSON.stringify(o));
        const description = o.description
          ? o.description.toString()
          : this.referenceMoment.clone().add(o.addParam).format(o.format);
        return _.extend(oClone, { description: description });
      });
    };

    // process 函数
    const process = options => {
      return options.map(o => {
        const date = this.referenceMoment
          .clone()
          .set(o.setParam)
          .add(o.addParam);
        delete o.addParam;
        o.setParam[o.granularity] = date[o.granularity]();
        o.defaultYear = date.year();
        return o;
      });
    };

    this.optionGenerators = process(setDescriptions(OPTION_GENERATORS));
    this.intervalOptions = setDescriptions(INTERVAL_OPTIONS);

    this.controllerParams.modelName = params.modelName;

    this.modelParams.context = context;
    this.modelParams.domain = domain;
    this.modelParams.modelName = params.modelName;
    this.modelParams.actionId = action.id;
    this.modelParams.fields = this.fields;

    this.rendererParams.action = action;
    this.rendererParams.breadcrumbs = params.breadcrumbs;
    this.rendererParams.context = context;
    this.rendererParams.searchMenuTypes = params.searchMenuTypes || [];
    this.rendererParams.template = params.template;
    this.rendererParams.title = params.title;
    this.rendererParams.withBreadcrumbs = params.withBreadcrumbs !== false;
    this.rendererParams.withSearchBar =
      'withSearchBar' in params ? params.withSearchBar : true;

    this.loadParams.withSearchBar =
      'withSearchBar' in params ? params.withSearchBar : true;
    this.loadParams.searchMenuTypes = params.searchMenuTypes || [];
    this.loadParams.activateDefaultFavorite = params.activateDefaultFavorite;
    if (this.loadParams.withSearchBar) {
      if (params.state) {
        this.loadParams.initialState = params.state;
      } else {
        // groups are determined in _parseSearchArch
        this.loadParams.groups = [];
        this.loadParams.timeRanges = context.time_ranges;
        this._parseSearchArch(this.arch);
      }
    }

    // add a filter group with the dynamic filters, if any
    if (params.dynamicFilters && params.dynamicFilters.length) {
      var dynamicFiltersGroup = params.dynamicFilters.map(function (filter) {
        return {
          description: filter.description,
          domain: JSON.stringify(filter.domain),
          isDefault: true,
          type: 'filter',
        };
      });
      this.loadParams.groups.unshift(dynamicFiltersGroup);
    }
  }

  connectedCallback() {
    var self = this;

    // exposed jQuery nodesets
    var shadow = this.shadowRoot; // # 初始化控制元素
    this.nodes = {
      buttons: shadow.querySelector('.o_cp_buttons'),
      pager: shadow.querySelector('.o_cp_pager'),
      sidebar: shadow.querySelector('.o_cp_sidebar'),
      switch_buttons: shadow.querySelector('.o_cp_switch_buttons'),
    };

    // if we don't use the default search bar and buttons, we expose the
    // corresponding areas for custom content
    if (!this.withSearchBar) {
      this.nodes.searchview = shadow.querySelector('.o_cp_searchview');
    }
    if (this.searchMenuTypes.length === 0) {
      this.nodes.searchview_buttons = shadow.querySelector('.o_search_options');
    }

    //if (this.withBreadcrumbs) {
    //  this._renderBreadcrumbs();
    //}

    var superDef = super.firstUpdated(...arguments);
    var searchDef = this._renderSearch();
    return Promise.all([superDef, searchDef]).then(function () {
      self._setSearchMenusVisibility();
    });
  }

  //$$$
  _renderBreadcrumbs() {
    var breadcrumbsDescriptors = this._breadcrumbs.concat({
      title: this._title,
    });
    var len = breadcrumbsDescriptors.length;
    return html`
      ${breadcrumbsDescriptors.map((bc, index) => {
        html`${this._renderBreadcrumbsItem(bc, index, len)}`;
      })}
    `;
  }
  /**
   * Render a breadcrumbs' li element.
   *
   * @private
   * @param {Object} bc
   * @param {string} bc.title
   * @param {string} bc.controllerID
   * @param {integer} index
   * @param {integer} length
   * @returns {jQueryElement} $bc
   */
  _renderBreadcrumbsItem(bc, index, length) {
    var self = this;
    var is_last = index === length - 1;
    var second_Last = index === length - 2;
    var li_content =
      (bc.title && _.escape(bc.title.trim())) || data.noDisplayContent;
    const onclick = function (e) {
      e.preventDefault();
      self.trigger_up('breadcrumb_clicked', {
        controllerID: bc.controllerID,
      });
    };

    return html`${is_last
      ? html` <li class="breadcrumb-item active">${li_content}</li> `
      : html`
          <li class="breadcrumb-item" @click="${onclick}">
            <a href="#"></a>
          </li>
        `}`;
  }

  /**
   * Renderer the search bar and the search menus
   *
   * @private
   * @returns {Promise}
   */
  _renderSearch() {
    var defs = [];
    /*
    if (this.menusSetup) {
      this._updateMenus();
    } else {
      this.menusSetup = true;
      defs = defs.concat(this._setupMenus());
    }
    */
    if (this.withSearchBar) {
      defs.push(this._renderSearchBar());
    }
    return Promise.all(defs).then(this._focusSearchInput.bind(this));
  }

  _renderSearchBar() {
    // TODO: might need a reload instead of a destroy/instantiate
    var oldSearchBar = this.searchBar;
    this.searchBar = new SearchBar(this, {
      context: this.context,
      facets: this.state.facets,
      fields: this.state.fields,
      filterFields: this.state.filterFields,
    });
    return this.searchBar.appendTo(this.$('.o_searchview')).then(function () {
      if (oldSearchBar) {
        oldSearchBar.destroy();
      }
    });
  }

  //Instantiate the search menu determined by this.searchMenuTypes.
  _setupMenus() {
    this.$subMenus = this._getSubMenusPlace();
    return this.searchMenuTypes.map(this._setupMenu.bind(this));
  }

  _getSubMenusPlace() {
    return $('<div>').appendTo(this.$('.o_search_options'));
  }

  /**
   * Create a new menu of the given type and append it to this.$subMenus.
   * This menu is also added to this.subMenus.
   */
  _setupMenu(menuType) {
    var Menu;
    var menu;
    if (menuType === 'filter') {
      Menu = FilterMenu;
    }
    if (menuType === 'groupBy') {
      Menu = GroupByMenu;
    }
    if (menuType === 'timeRange') {
      Menu = TimeRangeMenu;
    }
    if (menuType === 'favorite') {
      Menu = FavoriteMenu;
    }
    if (_.contains(['filter', 'groupBy', 'timeRange'], menuType)) {
      menu = new Menu(this, this._getMenuItems(menuType), this.state.fields);
    }
    if (menuType === 'favorite') {
      menu = new Menu(this, this._getMenuItems(menuType), this.action);
    }
    this.subMenus[menuType] = menu;
    return menu.appendTo(this.$subMenus);
  }

  // Update the search menus.
  _updateMenus() {
    var self = this;
    this.searchMenuTypes.forEach(function (menuType) {
      self.subMenus[menuType].update(self._getMenuItems(menuType));
    });
  }

  _getMenuItems(menuType) {
    var menuItems;
    if (menuType === 'filter') {
      menuItems = this.state.filters;
    }
    if (menuType === 'groupBy') {
      menuItems = this.state.groupBys;
    }
    if (menuType === 'timeRange') {
      menuItems = this.state.timeRanges;
    }
    if (menuType === 'favorite') {
      menuItems = this.state.favorites;
    }
    return menuItems;
  }

  getLastFacet() {
    return this.state.facets.slice(-1)[0];
  }

  queryChanged(e) {
    this.action = this.app.action;
    this.views = this.app.views;
    this.viewMode = this.app.router.Get('view');
    var show = false;

    if (this.action) {
      // 初始化Action 的Mode
      this.viewIdOfMode = {};
      this.modes = this.action.view_mode.split(',');

      //  views : [[126, "kanban"], [121, "list"], [124, "form"]]
      if (this.action.views) {
        for (var i = 0; i < this.action.views.length; i++) {
          this.viewIdOfMode[this.action.views[i][1]] = this.action.views[i][0];

          if (
            this.action.view_id &&
            this.action.view_id == this.action.views[i][0]
          ) {
            this.defaultView = this.action.views[i][1]; // action 默认视图
          }
        }

        // 默认视图
        if (!this.defaultView || this.defaultView == '') {
          if (this.action.view_type) {
            this.defaultView = this.action.view_type;
          } else {
            this.defaultView = this.modes[0];
          }
        }

        // 当有可用的View时才写换mode
        if (!this.viewMode && this.action.views.length > 0) {
          //如果存在且支持该Mode
          //if (!this.viewMode.indexOf(this.action.view_mode)) {
          this.viewMode = this.defaultView;
          //}
        }
      }

      if (this.views && this.views.fields_views[this.viewMode]) {
        this.fields = this.views.fields;

        /*
      var pages = this.viewManager.shadowRoot.querySelector("#pages");
      pages.items.forEach(function(item) {
        item.remove();        //pages.removeChild(item);
      });
      */

        // # 重建View页面

        let slots = this.viewManager.shadowRoot.querySelector('slot');
        // for (var view_type in this.views.fields_views) {
        if (!this.viewByMode[this.viewMode] && this.viewMode != 'search') {
          // # 如果View　页面不存在于索引里
          var viewFrame = document.createElement('div');
          viewFrame.innerHTML = this.views.fields_views[this.viewMode].arch;
          let view = viewFrame.firstChild;
          view.controlPanel = this;
          view.viewManager = this.viewManager;
          view.datasource = this.viewManager.datasource;
          let index = slots.assignedNodes({ flatten: true }).length;
          this.viewByMode[this.viewMode] = view;
          this.modeOfPage[this.viewMode] = index; //+ 1;
          this.viewManager.appendChild(viewFrame); // 添加到Manger的Slot里
        }
        //}

        // 建立搜索视图控件
        var searchNode = this.shadowRoot.querySelector('div.cp-search-view'); //
        if (searchNode) {
          // 不可查询的删除节点反之创建
          var searchable = this.viewByMode[this.viewMode].searchable;
          if (searchNode.firstChild && !searchable) {
            searchNode.firstChild.remove();
          } else if (!searchNode.firstChild && searchable) {
            if ('search' in this.views.fields_views) {
              searchNode.innerHTML = this.views.fields_views['search'].arch;
              this.searchView = searchNode.firstChild; // 展示View
              this.searchView.controlPanel = this;
              this.searchView.viewManager = this.viewManager;
              this.searchView.updateComplete.then(() => {
                this.searchView.show(this);
              });
            }
          }
        }

        show = true;
      } else {
        let slots = this.viewManager.shadowRoot.querySelector('slot');
        let views = slots.assignedNodes({ flatten: true });
        views.forEach(function (item) {
          item.remove(); // 移除节点
        });
        this.viewByMode = {};
        this.modeOfPage = {};
      }
    }

    this.style.display = show ? 'block' : 'none'; //

    if (show && this.viewMode) {
      this.SetViewMode(this.viewMode);
    }
  }

  doPrepareSearch() {
    var self = this;
    return new Promise(function (resolve, reject) {
      if (!self.fields) {
        resolve(undefined);
        return;
      }

      var action_context = {};
      var view_context = {}; //this.get_context();
      var domain = [];

      if (self.action) {
        action_context = self.action.context || {};
        domain = self.action.domain;
      }

      pyeval
        .eval_domains_and_contexts({
          domains: [domain || []].concat(self.domains || []),
          contexts: [action_context, view_context].concat(self.contexts || []),
          group_by_seq: self.groupbys || [],
        })
        .then(function (results) {
          if (results.error) {
            // self.active_search.resolve();
            // throw new Error(
            //         _.str.sprintf(_t("Failed to evaluate search criterions")+": \n%s",
            //                       JSON.stringify(results.error)));
            return;
          }
          //self.dataset._model = new Model(
          //   self.dataset.model, results.context, results.domain);

          var groupby = results.group_by.length
            ? results.group_by
            : action_context.group_by;
          if (vectors.utils.isString(groupby)) {
            groupby = [groupby];
          }

          //if (!controller.grouped && !vectors.utils.isEmpty(groupby)) {
          //    self.dataset.set_sort([]);
          // }

          var params = {};
          params.context = results.context;
          params.domain = JSON.stringify(results.domain);
          params.group_by = groupby;
          params.views = self.views ? self.views.fields_views : {}; // # 传递Views视图列表给ViewManager
          params.fields = Object.keys(self.fields); // # 传递Views字段列表给ViewManager
          params.model = self.action ? self.action.res_model : '';
          params.limit = self.action ? self.action.limit : 80;
          params.offset = '';
          params.sort = '';
          resolve(params);
        });
    });
  }

  // 执行搜索程序
  doSearch() {
    // # 触发查询事件
    var self = this;
    this.doPrepareSearch().then(function (params) {
      self.dispatchEvent(
        new CustomEvent('search', {
          detail: params,
          bubbles: true,
          composed: true,
        })
      );
    });
  }

  //-----------------------------------------

  onModeBtnClick(e) {
    this.SetViewMode(e.currentTarget.modeType); // 改变Mode并激活事件
  }

  onHotkeys(e) {
    // ESC
    if (e.keyCode === 27 && e.rootTarget === this.$.query) {
      this.showingSearch = false;
    }
  } /*
    return out;
  }

  // 修改且更新Query值
  SetQuery(query) {
    if (query) {
      this.app.SetQuery(query);
    }
  }

 
  // 从基础参数上添加或覆盖Map对象的查询参数
  // 如果值为undefined则删除该值
  UpdateQuery(query) {
    if (query) {
      var newQuery = this.GetQuery();
      for (var key in query) {
        var value = query[key];
        if (key) {
          if (!value) {
            //undefined
            delete newQuery[key];
            continue;
          }

          if (key == "view") {
            if (this.mode == value) {
              continue;
            } else {
              this.mode = value;
              var self = this;
              this.doPrepareSearch().then(function(params) {
                // params.mode = self.mode;
                // params.pageIdex = self.modeOfPage[self.mode];
                // params.view = self.viewByMode[self.mode];
                if (params) {
                  self.dispatchEvent(
                    new CustomEvent("onViewModeChanged", {
                      detail: { params: params, mode: self.mode, pageIdex: self.modeOfPage[self.mode], view: self.viewByMode[self.mode] }
                    })
                  );
                }
              });
            }
          }

          newQuery[key] = value;
        }
      }

      // 执行更新
      this.SetQuery(newQuery);
    }
  }
*/
  /*
  // 获取Query字典
  GetQuery() {
    var out = {};
    // 获取现有查询字符
    var query = this.app.GetQuery();
    out = query;
    /*
    // 只保留基本查询字符
    for (var key in query) {
      if (["menu", "model", "action", "view","id"].indexOf(key) > -1) {
        out[key] = query[key];
      }
    }
    */ GetViewMode() {
    return this.app.router.Get('view');
  }

  // 更新查询字符和事件
  SetViewMode(mode) {
    if (mode && mode != '') {
      this.app.router.SetQuery({ view: mode });
      var self = this;
      var index = self.modeOfPage[mode];
      var view = self.viewByMode[mode];
      // var viewFrame = self.viewByMode[mode];
      // if (!viewFrame.firstChild) {
      //   viewFrame.innerHTML = this.views.fields_views[mode].arch;
      // }

      this.doPrepareSearch().then(function (params) {
        if (params) {
          self.dispatchEvent(
            new CustomEvent('onViewModeChanged', {
              detail: {
                params: params,
                mode: mode,
                pageIdex: index,
                view: view,
              },
            })
          );
        }
      });
    }
  }
}

customElements.define('control-panel', ControlPanel);
