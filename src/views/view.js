import { html, css } from 'lit-element';
import { Controller } from './controller';
import { Renderer } from './renderer.js';
import { Model } from './model.js';
import { SearchPanel } from './search_panel';
import { AbstractView } from './abstract_view';
import { viewUtils } from './view_utils';
import { ActionControllerPanel } from './controller-panel/control-panel';

export class View extends AbstractView {
  static get styles() {
    return css`
      :host {
        display: block;
        color: red;
      }
    `;
  }

  render() {
    return html``;
  }

  static get properties() {
    return {
      accessKey: { type: String },
      display_name: { type: String, value: 'fa-question' },
      icon: { type: String },
      multi_record: { type: Boolean, value: true },
      withSearchBar: { type: Boolean, value: true },

      app: { type: Object },
      controlPanel: { type: Object }, // 控制面板
      viewManager: { type: Object }, // 视图控制器
      fields: { type: Array }, // View 的所有可用Fields对象集合
      inputs: { type: Array }, // 输入控件 [FORM]
      datasource: { type: Object, notify: true },
      params: { type: Object }, // 查询参数
      ///searchNode: { type: Object }, // # search view node
      mode: String, // #
      Searchable: { type: Boolean }, // 可搜索并显示搜索过滤器
    };
  }

  constructor(viewInfo, params) {
    super(viewInfo, params);

    // name displayed in view switchers
    this.display_name = '';
    // indicates whether or not the view is mobile-friendly
    this.mobile_friendly = false;
    // icon is the font-awesome icon to display in the view switcher
    this.icon = 'fa-question';

    // multi_record is used to distinguish views displaying a single record
    // (e.g. FormView) from those that display several records (e.g. ListView)
    this.multi_record = true;
    // viewType is the type of the view, like 'form', 'kanban', 'list'...
    this.viewType = undefined;
    // determines if a search bar is available
    this.withSearchBar = true;
    // determines the search menus available and their orders
    this.searchMenuTypes = ['filter', 'groupBy', 'favorite'];
    // determines if a control panel should be instantiated
    this.withControlPanel = true;
    // determines if a search panel could be instantiated
    this.withSearchPanel = true;

    var action = params.action || {};
    params = _.defaults(params, this._extractParamsFromAction(action));

    // in general, the fieldsView has to be processed by the View (e.g. the
    // arch is a string that needs to be parsed) ; the only exception is for
    // inline form views inside form views, as they are processed alongside
    // the main view, but they are opened in a FormViewDialog which
    // instantiates another FormView (unlike kanban or list subviews for
    // which only a Renderer is instantiated)
    if (typeof viewInfo.arch === 'string') {
      this.fieldsView = this._processFieldsView(viewInfo);
    } else {
      this.fieldsView = viewInfo;
    }
    this.arch = this.fieldsView.arch;
    this.fields = this.fieldsView.viewFields;
    this.userContext = params.userContext || {};
    this.withControlPanel = this.withControlPanel && params.withControlPanel;
    const searchPanelDisabled =
      'search_panel' in params.context && !params.search_panel;
    this.withSearchPanel =
      this.withSearchPanel &&
      this.multi_record &&
      params.withSearchPanel &&
      !searchPanelDisabled;

    // the boolean parameter 'isEmbedded' determines if the view should be
    // considered as a subview. For now this is only used by the graph
    // controller that appends a 'Group By' button beside the 'Measures'
    // button when the graph view is embedded.
    var isEmbedded = params.isEmbedded || false;

    this.rendererParams = {
      arch: this.arch,
      isEmbedded: isEmbedded,
      noContentHelp: params.noContentHelp,
    };

    this.controllerParams = {
      actionViews: params.actionViews,
      activeActions: {
        edit: this.arch.attrs.edit ? !!JSON.parse(this.arch.attrs.edit) : true,
        create: this.arch.attrs.create
          ? !!JSON.parse(this.arch.attrs.create)
          : true,
        delete: this.arch.attrs.delete
          ? !!JSON.parse(this.arch.attrs.delete)
          : true,
        duplicate: this.arch.attrs.duplicate
          ? !!JSON.parse(this.arch.attrs.duplicate)
          : true,
      },
      bannerRoute: this.arch.attrs.banner_route,
      controllerID: params.controllerID,
      displayName: params.displayName,
      isEmbedded: isEmbedded,
      isMultiRecord: this.multi_record,
      modelName: params.modelName,
      viewType: this.viewType,
    };

    var controllerState = params.controllerState || {};
    var currentId = controllerState.currentId || params.currentId;
    this.loadParams = {
      context: params.context,
      count:
        params.count ||
        (this.controllerParams.ids !== undefined &&
          this.controllerParams.ids.length) ||
        0,
      domain: params.domain,
      modelName: params.modelName,
      res_id: currentId,
      res_ids:
        controllerState.resIds ||
        params.ids ||
        (currentId ? [currentId] : undefined),
    };
    // default_order is like:
    //   'name,id desc'
    // but we need it like:
    //   [{name: 'id', asc: false}, {name: 'name', asc: true}]
    var defaultOrder = this.arch.attrs.default_order;
    if (defaultOrder) {
      this.loadParams.orderedBy = _.map(defaultOrder.split(','), function (
        order
      ) {
        order = order.trim().split(' ');
        return { name: order[0], asc: order[1] !== 'desc' };
      });
    }
    if (params.searchQuery) {
      this._updateMVCParams(params.searchQuery);
    }

    // determines the MVC components to use
    this.config = _.extend({}, AbstractView.prototype.config, {
      Model: Model,
      Renderer: Renderer,
      Controller: Controller,
      SearchPanel: SearchPanel,
    });

    // 控制面板初始化参数
    this.controlPanelParams = {
      action: action,
      activateDefaultFavorite: params.activateDefaultFavorite,
      dynamicFilters: params.dynamicFilters,
      breadcrumbs: params.breadcrumbs,
      context: this.loadParams.context,
      domain: this.loadParams.domain,
      modelName: params.modelName,
      searchMenuTypes: params.searchMenuTypes,
      state: controllerState.cpState,
      viewInfo: params.controlPanelFieldsView,
      withBreadcrumbs: params.withBreadcrumbs,
      withSearchBar: params.withSearchBar,
    };

    this.searchPanelParams = {
      defaultNoFilter: params.searchPanelDefaultNoFilter,
      fields: this.fields,
      model: this.loadParams.modelName,
      state: controllerState.spState,
    };
  }

  show(element) {
    // 获取该视图有效字段
    // 使用  Array.prototype.slice.call(); 转换到Arry
    var fields = Array.prototype.slice.call(this.querySelectorAll('field'));
    var availbleFields = [];
    fields.forEach(field => {
      var lFieldName = field.getAttribute('name'); // 字段名
      var lFieldRec = this.controlPanel.fields[lFieldName]; // View 里的Field记录
      if (lFieldRec) {
        availbleFields.push(lFieldRec);
      }
    });

    this.fields = availbleFields;
  }
  /**
   * @private
   * @param {Object} [action]
   * @param {Object} [action.context || {}]
   * @param {boolean} [action.context.no_breadcrumbs=false]
   * @param {integer} [action.context.active_id]
   * @param {integer[]} [action.context.active_ids]
   * @param {Object} [action.controlPanelFieldsView]
   * @param {string} [action.display_name]
   * @param {Array[]} [action.domain=[]]
   * @param {string} [action.help]
   * @param {integer} [action.id]
   * @param {integer} [action.limit]
   * @param {string} [action.name]
   * @param {string} [action.res_model]
   * @param {string} [action.target]
   * @returns {Object}
   */
  _extractParamsFromAction(action) {
    action = action || {};
    var context = action.context || {};
    var inline = action.target === 'inline';
    return {
      actionId: action.id || false,
      actionViews: action.views || [],
      activateDefaultFavorite: !context.active_id && !context.active_ids,
      context: action.context || {},
      controlPanelFieldsView: action.controlPanelFieldsView,
      currentId: action.res_id ? action.res_id : undefined, // load returns 0
      displayName: action.display_name || action.name,
      domain: action.domain || [],
      limit: action.limit,
      modelName: action.res_model,
      noContentHelp: action.help,
      searchMenuTypes: inline ? [] : this.searchMenuTypes,
      withBreadcrumbs:
        'no_breadcrumbs' in context ? !context.no_breadcrumbs : true,
      withControlPanel: this.withControlPanel,
      withSearchBar: inline ? false : this.withSearchBar,
      withSearchPanel: this.withSearchPanel,
    };
  }

  /**
   * Processes a fieldsView. In particular, parses its arch.
   *
   * @private
   * @param {Object} fieldsView
   * @param {string} fieldsView.arch
   * @returns {Object} the processed fieldsView
   */
  _processFieldsView(fieldsView) {
    var fv = _.extend({}, fieldsView);
    fv.arch = viewUtils.parseArch(fv.arch);
    fv.viewFields = _.defaults({}, fv.viewFields, fv.fields);
    return fv;
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * Main method of the Factory class. Create a controller, and make sure that
   * data and libraries are loaded.
   *
   * There is a unusual thing going in this method with parents: we create
   * renderer/model with parent as parent, then we have to reassign them at
   * the end to make sure that we have the proper relationships.  This is
   * necessary to solve the problem that the controller needs the model and
   * the renderer to be instantiated, but the model need a parent to be able
   * to load itself, and the renderer needs the data in its constructor.
   *
   * @param {Widget} parent the parent of the resulting Controller (most
   *      likely an action manager)
   * @returns {Promise<Controller>}
   */
  // 获取视图控制器 包括控制面板以及视图内容
  getController(parent) {
    var self = this;
    var cpDef = this.withControlPanel && this._createControlPanel(parent);
    var spDef;
    if (this.withSearchPanel) {
      var spProto = this.config.SearchPanel.prototype;
      var viewInfo = this.controlPanelParams.viewInfo;
      var searchPanelParams = spProto.computeSearchPanelParams(
        viewInfo,
        this.viewType
      );
      if (searchPanelParams.sections) {
        this.searchPanelParams.sections = searchPanelParams.sections;
        this.rendererParams.withSearchPanel = true;
        spDef = Promise.resolve(cpDef).then(
          this._createSearchPanel.bind(this, parent, searchPanelParams)
        );
      }
    }

    var _super = super.getController.bind(this);
    return Promise.all([cpDef, spDef]).then(function ([
      controlPanel,
      searchPanel,
    ]) {
      // get the parent of the model if it already exists, as _super will
      // set the new controller as parent, which we don't want
      //var modelParent = self.model && self.model.getParent();
      var modelParent = self.model;

      var prom = _super(parent);
      prom.then(function (controller) {
        if (controlPanel) {
          //controlPanel.setParent(controller);
        }
        if (searchPanel) {
          //searchPanel.setParent(controller);
        }
        if (modelParent) {
          // if we already add a model, restore its parent
          //self.model.setParent(modelParent);
        }
      });

      return prom;
    });
  }

  /**
   * Instantiates and starts a ControlPanelController.
   *
   * @private
   * @param {Widget} parent
   * @returns {Promise<ControlPanelController>} resolved when the controlPanel
   *   is ready
   */
  // 创建控制面板
  _createControlPanel(parent) {
    var self = this;
    var controlPanelView = new ActionControllerPanel(this.controlPanelParams);
    return controlPanelView.getController(parent).then(function (controlPanel) {
      self.controllerParams.controlPanel = controlPanel;
      var fragment = document.createDocumentFragment();
      new Promise(function (resolve, reject) {
        // 组织节点
        fragment.append(controlPanel);
        resolve();
      }).then(function () {
        //self._updateMVCParams(controlPanel.getSearchQuery());
        return controlPanel;
      });
    });
  }
  /**
   * @private
   * @param {Widget} parent
   * @returns {Promise<SearchPanel>} resolved when the searchPanel is ready
   */
  async _createSearchPanel(parent, params) {
    var defaultValues = {};
    Object.keys(this.loadParams.context).forEach(key => {
      let match = /^searchpanel_default_(.*)$/.exec(key);
      if (match) {
        defaultValues[match[1]] = this.loadParams.context[key];
      }
    });

    var controlPanelDomain = this.loadParams.domain;
    const viewDomain = await this._getViewDomain(parent);
    var spParams = _.extend({}, this.searchPanelParams, {
      defaultValues: defaultValues,
      searchDomain: controlPanelDomain,
      viewDomain,
      classes: params.classes || [],
    });

    //var searchPanel = new this.config.SearchPanel(parent, spParams);
    var searchPanel = new ViewSearch();
    this.controllerParams.searchPanel = searchPanel;
    this.controllerParams.controlPanelDomain = controlPanelDomain;
    var fragment = document.createDocumentFragment();
    await fragment.append(searchPanel);

    var searchPanelDomain = searchPanel.getDomain();
    this.loadParams.domain = controlPanelDomain.concat(searchPanelDomain);
    return searchPanel;
  }

  getModel(parent) {
    if (!this.model) {
      this.model = new Model(parent);
    }
    return this.model;
  }

  toggleClass(node, className, allowed) {
    if (node) {
      if (allowed) {
        node.classList.add(className);
      } else {
        node.classList.remove(className);
      }
    }
  }
}
