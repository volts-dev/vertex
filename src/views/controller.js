import { AbstractController } from './abstract_controller';
import { Concurrency } from '../core/concurrency';

export class Controller extends AbstractController {
  /*  custom_events: _.extend({}, ActionMixin.custom_events, {
        navigation_move: '_onNavigationMove',
        open_record: '_onOpenRecord',
        search: '_onSearch',
        switch_view: '_onSwitchView',
        search_panel_domain_updated: '_onSearchPanelDomainUpdated',
    }),
    events: {
        'click a[type="action"]': '_onActionClicked',
    }
*/
  /**
   * @override
   * @param {string} params.modelName
   * @param {string} [params.controllerID] an id to ease the communication
   *   with upstream components
   * @param {ControlPanelController} [params.controlPanel]
   * @param {any} [params.handle] a handle that will be given to the model (some id)
   * @param {boolean} params.isMultiRecord
   * @param {Object[]} params.actionViews
   * @param {string} params.viewType
   */
  constructor(parent, model, renderer, params) {
    super(...arguments);
    //this._super.apply(this, arguments);
    this._controlPanel = params.controlPanel;
    this._title = params.displayName;
    this.modelName = params.modelName;
    this.activeActions = params.activeActions;
    this.controllerID = params.controllerID;
    this.initialState = params.initialState;
    this.bannerRoute = params.bannerRoute;
    this.isMultiRecord = params.isMultiRecord;
    this.actionViews = params.actionViews;
    this.viewType = params.viewType;
    // use a DropPrevious to correctly handle concurrent updates
    this.dp = new Concurrency.DropPrevious();

    // the following attributes are used when there is a searchPanel
    this._searchPanel = params.searchPanel;
    this.controlPanelDomain = params.controlPanelDomain || [];
    this.searchPanelDomain = this._searchPanel
      ? this._searchPanel.getDomain()
      : [];
  }

  /**
   * Simply renders and updates the url.
   *
   * @returns {Promise}
   */
  firstUpdated() {
    var self = this;
    if (this._searchPanel) {
      var content = this.shadowRoot.querySelector('.o_content');
      content.classList.add('o_controller_with_searchpanel');
      content.append(this._searchPanel);
    }

    this.classList.add('o_view_controller');
    //this.$el.addClass('o_view_controller');

    super
      .firstUpdated()
      .then(function () {
        var prom;
        if (self._controlPanel) {
          // render the ControlPanel elements (buttons, pager, sidebar...)
          prom = self._renderControlPanelElements().then(function (elements) {
            self.controlPanelElements = elements; // 传递元件
            //self._controlPanel.$el.prependTo(self.$el);
            self.shadowRoot.append(self._controlPanel); // 插入节点
          });
        }
        return Promise.resolve(prom);
      })
      .then(function () {
        // 渲染更新控制板
        return self._update(self.initialState);
      });

    return;
  }

  /**
   * Called each time the controller is attached into the DOM.
   */
  connectedCallback() {
    super.connectedCallback();

    if (this._controlPanel) {
      this._controlPanel.connectedCallback();
    }
    if (this._searchPanel) {
      this._searchPanel.connectedCallback();
    }
    this.renderer.connectedCallback();
  }

  /**
   * Called each time the controller is detached from the DOM.
   */
  disconnectedCallback() {
    super.disconnectedCallback();

    if (this._controlPanel) {
      this._controlPanel.disconnectedCallback();
    }
    this.renderer.disconnectedCallback();
  }

  /**
   * @override
   */
  destroy() {
    /*
        if (this.$buttons) {
            this.$buttons.off();
        }
        if (this.controlPanelElements && this.controlPanelElements.$switch_buttons) {
            this.controlPanelElements.$switch_buttons.off();
        }
        this._super.apply(this, arguments);
        */
    super.destroy();
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * @override
   */
  canBeRemoved() {
    // AAB: get rid of 'readonlyIfRealDiscard' option when on_hashchange mechanism is improved
    return this.discardChanges(undefined, {
      noAbandon: true,
      readonlyIfRealDiscard: true,
    });
  }

  /**
   * Discards the changes made on the record associated to the given ID, or
   * all changes made by the current controller if no recordID is given. For
   * example, when the user opens the 'home' screen, the action manager calls
   * this method on the active view to make sure it is ok to open the home
   * screen (and lose all current state).
   *
   * Note that it returns a Promise, because the view could choose to ask the
   * user if he agrees to discard.
   *
   * @param {string} [recordID]
   *        if not given, we consider all the changes made by the controller
   * @param {Object} [options]
   * @returns {Promise} resolved if properly discarded, rejected otherwise
   */
  discardChanges(recordID, options) {
    return Promise.resolve();
  }

  /**
   * Export the state of the controller containing information that is shared
   * between different controllers of a same action (like the current
   * searchQuery of the controlPanel).
   *
   * @returns {Object}
   */
  exportState() {
    var state = {};
    if (this._controlPanel) {
      state.cpState = this._controlPanel.exportState();
    }
    if (this._searchPanel) {
      state.spState = this._searchPanel.exportState();
    }
    return state;
  }

  /**
   * Gives the focus to the renderer
   */
  giveFocus() {
    this.renderer.giveFocus();
  }

  /**
   * The use of this method is discouraged.  It is still snakecased, because
   * it currently is used in many templates, but we will move to a simpler
   * mechanism as soon as we can.
   *
   * @deprecated
   * @param {string} action type of action, such as 'create', 'read', ...
   * @returns {boolean}
   */
  is_action_enabled(action) {
    return this.activeActions[action];
  }

  /**
   * Short helper method to reload the view
   *
   * @param {Object} [params] This object will simply be given to the update
   * @returns {Promise}
   */
  async reload(params) {
    params = params || {};
    var searchPanelUpdateProm;
    var controllerState = params.controllerState || {};
    var cpState = controllerState.cpState;
    if (this._controlPanel && cpState) {
      await this._controlPanel
        .importState(cpState)
        .then(function (searchQuery) {
          params = _.extend({}, params, searchQuery);
        });
    }
    var postponeRendering = false;
    if (this._searchPanel) {
      this.controlPanelDomain = params.domain || this.controlPanelDomain;
      if (controllerState.spState) {
        this._searchPanel.importState(controllerState.spState);
        this.searchPanelDomain = this._searchPanel.getDomain();
      } else {
        const viewDomain = await this._getViewDomain();
        searchPanelUpdateProm = this._searchPanel.update({
          searchDomain: this.controlPanelDomain,
          viewDomain,
        });
        postponeRendering = !params.noRender;
        params.noRender = true; // wait for searchpanel to be ready to render
      }
      params.domain = this.controlPanelDomain.concat(this.searchPanelDomain);
    }
    await Promise.all([this.update(params, {}), searchPanelUpdateProm]);
    if (postponeRendering) {
      return this.renderer._render();
    }
  }
  /**
   * For views that require a pager, this method will be called to allow the
   * controller to instantiate and render a pager. Note that in theory, the
   * controller can actually render whatever he wants in the pager zone.  If
   * your view does not want a pager, just let this method empty.
   *
   * @param {jQuery Node} $node
   * @return {Promise}
   */
  renderPager(node) {
    return Promise.resolve();
  }
  /**
   * Same as renderPager, but for the 'sidebar' zone (the zone with the menu
   * dropdown in the control panel next to the buttons)
   *
   * @param {jQuery Node} $node
   * @return {Promise}
   */
  renderSidebar(node) {
    return Promise.resolve();
  }
  /**
   * This is the main entry point for the controller.  Changes from the search
   * view arrive in this method, and internal changes can sometimes also call
   * this method.  It is basically the way everything notifies the controller
   * that something has changed.
   *
   * The update method is responsible for fetching necessary data, then
   * updating the renderer and wait for the rendering to complete.
   *
   * @param {Object} params will be given to the model and to the renderer
   * @param {Object} [options]
   * @param {boolean} [options.reload=true] if true, the model will reload data
   *
   * @returns {Promise}
   */
  update(params, options) {
    var self = this;
    var shouldReload = options && 'reload' in options ? options.reload : true;
    var def = shouldReload
      ? this.model.reload(this.handle, params)
      : Promise.resolve();
    // we check here that the updateIndex of the control panel hasn't changed
    // between the start of the update request and the moment the controller
    // asks the control panel to update itself ; indeed, it could happen that
    // another action/controller is executed during this one reloads itself,
    // and if that one finishes first, it replaces this controller in the DOM,
    // and this controller should no longer update the control panel.
    // note that this won't be necessary as soon as each controller will have
    // its own control panel
    var cpUpdateIndex = this._controlPanel && this._controlPanel.updateIndex;
    return this.dp.add(def).then(function (handle) {
      if (
        self._controlPanel &&
        cpUpdateIndex !== self._controlPanel.updateIndex
      ) {
        return;
      }
      self.handle = handle || self.handle; // update handle if we reloaded
      var state = self.model.get(self.handle);
      var localState = self.renderer.getLocalState();
      return self.dp
        .add(self.renderer.updateState(state, params))
        .then(function () {
          if (
            self._controlPanel &&
            cpUpdateIndex !== self._controlPanel.updateIndex
          ) {
            return;
          }
          self.renderer.setLocalState(localState);
          return self._update(state, params);
        });
    });
  }

  //--------------------------------------------------------------------------
  // Private
  //--------------------------------------------------------------------------

  /**
   * Get the domain defined by the view. It is meant to be overridden.
   *
   * @private
   * @returns {Promise<Array[]>}
   */
  async _getViewDomain() {
    return [];
  }

  /**
   * This method is the way a view can notifies the outside world that
   * something has changed.  The main use for this is to update the url, for
   * example with a new id.
   *
   * @private
   * @param {Object} [state] information that will be pushed to the outside
   *   world
   */
  _pushState(state) {
    this.trigger_up('push_state', {
      controllerID: this.controllerID,
      state: state || {},
    });
  }

  /**
   * Renders the html provided by the route specified by the
   * bannerRoute attribute on the controller (banner_route in the template).
   * Renders it before the view output and add a css class 'o_has_banner' to it.
   * There can be only one banner displayed at a time.
   *
   * If the banner contains stylesheet links or js files, they are moved to <head>
   * (and will only be fetched once).
   *
   * Route example:
   * @http.route('/module/hello', auth='user', type='json')
   * def hello(self):
   *     return {'html': '<h1>hello, world</h1>'}
   *
   * @private
   * @returns {Promise}
   */
  _renderBanner() {
    /*
        if (this.bannerRoute !== undefined) {
            var self = this;
            return this.dp
                .add(this._rpc({route: this.bannerRoute}))
                .then(function (response) {
                    if (!response.html) {
                        self.$el.removeClass('o_has_banner');
                        return Promise.resolve();
                    }
                    self.$el.addClass('o_has_banner');
                    var $banner = $(response.html);
                    // we should only display one banner at a time
                    if (self._$banner && self._$banner.remove) {
                        self._$banner.remove();
                    }
                    // Css and js are moved to <head>
                    var defs = [];
                    $('link[rel="stylesheet"]', $banner).each(function (i, link) {
                        defs.push(ajax.loadCSS(link.href));
                        link.remove();
                    });
                    $('script[type="text/javascript"]', $banner).each(function (i, js) {
                        defs.push(ajax.loadJS(js.src));
                        js.remove();
                    });
                    return Promise.all(defs).then(function () {
                        $banner.insertBefore(self.$('> .o_content'));
                        self._$banner = $banner;
                    });
                });
        }
        */
    return Promise.resolve();
  }

  /**
   * Renders the control elements (buttons, pager and sidebar) of the current
   * view.
   *
   * @private
   * @returns {Promise<Object>} resolved with an object containing the control
   *   panel jQuery elements
   */
  // 渲染控制板元件并返回元件集elements
  _renderControlPanelElements() {
    var self = this;
    var elements = {
      buttons: document.createElement('div'),
      sidebar: document.createElement('div'),
      pager: document.createElement('div'),
    };

    this.renderButtons(elements.buttons);
    var sidebarProm = this.renderSidebar(elements.sidebar);
    var pagerProm = this.renderPager(elements.pager);

    return Promise.all([sidebarProm, pagerProm]).then(function () {
      // remove the unnecessary outer div
      elements = _.mapObject(elements, function (node) {
        //return node && node.contents();
        return node && node.innerHTML;
      });
      elements.switch_buttons = self._renderSwitchButtons();

      return elements;
    });
  }

  /**
   * Renders the switch buttons and binds listeners on them.
   *
   * @private
   * @returns {jQuery}
   */
  // 渲染视图选择列表
  _renderSwitchButtons() {
    return;
  }

  /**
   * @override
   * @private
   */
  _startRenderer() {
    return; //this.renderer.appendTo(this.$(".o_content"));
  }

  /**
   * This method is called after each update or when the start method is
   * completed.
   *
   * Its primary use is to be used as a hook to update all parts of the UI,
   * besides the renderer.  For example, it may be used to enable/disable
   * some buttons in the control panel, such as the current graph type for a
   * graph view.
   *
   * @private
   * @param {Object} state the state given by the model
   * @param {Object} [params]
   * @param {Object[]} [params.breadcrumbs]
   * @returns {Promise}
   */
  _update(state, params) {
    // AAB: update the control panel -> this will be moved elsewhere at some point
    var cpContent = _.extend({}, this.controlPanelElements);
    this.updateControlPanel({
      breadcrumbs: params && params.breadcrumbs,
      cp_content: cpContent,
      title: this.getTitle(),
    });

    this._pushState();
    return this._renderBanner();
  }

  //--------------------------------------------------------------------------
  // Handlers
  //--------------------------------------------------------------------------

  /**
   * When a user clicks on an <a> link with type="action", we need to actually
   * do the action. This kind of links is used a lot in no-content helpers.
   *
   * * if the link has both data-model and data-method attributes, the
   *   corresponding method is called, chained to any action it would
   *   return. An optional data-reload-on-close (set to a non-falsy value)
   *   also causes th underlying view to be reloaded after the dialog is
   *   closed.
   * * if the link has a name attribute, invoke the action with that
   *   identifier (see :class:`ActionManager.doAction` to not get the
   *   details)
   * * otherwise an *action descriptor* is built from the link's data-
   *   attributes (model, res-id, views, domain and context)
   *
   * @private
   * @param ev
   */
  _onActionClicked(ev) {
    // FIXME: maybe this should also work on <button> tags?
    ev.preventDefault();
    var $target = $(ev.currentTarget);
    var self = this;
    var data = $target.data();

    if (data.method !== undefined && data.model !== undefined) {
      var options = {};
      if (data.reloadOnClose) {
        options.on_close = function () {
          self.trigger_up('reload');
        };
      }
      this.dp
        .add(
          this._rpc({
            model: data.model,
            method: data.method,
            context: session.user_context,
          })
        )
        .then(function (action) {
          if (action !== undefined) {
            self.do_action(action, options);
          }
        });
    } else if ($target.attr('name')) {
      this.do_action(
        $target.attr('name'),
        data.context && { additional_context: data.context }
      );
    } else {
      this.do_action(
        {
          name: $target.attr('title') || _.str.strip($target.text()),
          type: 'ir.actions.act_window',
          res_model: data.model || this.modelName,
          res_id: data.resId,
          target: 'current', // TODO: make customisable?
          views:
            data.views ||
            (data.resId
              ? [[false, 'form']]
              : [
                  [false, 'list'],
                  [false, 'form'],
                ]),
          domain: data.domain || [],
        },
        {
          additional_context: _.extend({}, data.context),
        }
      );
    }
  }

  /**
   * Called either from the control panel to focus the controller
   * or from the view to focus the search bar
   *
   * @private
   * @param {OdooEvent} ev
   */
  _onNavigationMove(ev) {
    switch (ev.data.direction) {
      case 'up':
        ev.stopPropagation();
        this._controlPanel.focusSearchBar();
        break;
      case 'down':
        ev.stopPropagation();
        this.giveFocus();
        break;
    }
  }

  /**
   * When an Odoo event arrives requesting a record to be opened, this method
   * gets the res_id, and request a switch view in the appropriate mode
   *
   * Note: this method seems wrong, it relies on the model being a basic model,
   * to get the res_id.  It should receive the res_id in the event data
   * @todo move this to basic controller?
   *
   * @private
   * @param {OdooEvent} ev
   * @param {number} ev.data.id The local model ID for the record to be
   *   opened
   * @param {string} [ev.data.mode='readonly']
   */
  _onOpenRecord(ev) {
    ev.stopPropagation();
    var record = this.model.get(ev.data.id, { raw: true });
    this.trigger_up('switch_view', {
      view_type: 'form',
      res_id: record.res_id,
      mode: ev.data.mode || 'readonly',
      model: this.modelName,
    });
  }
  /**
   * Called when there is a change in the search view, so the current action's
   * environment needs to be updated with the new domain, context and groupby.
   *
   * @private
   * @param {OdooEvent} ev
   * @param {Array[]} ev.data.domain
   * @param {Object} ev.data.context
   * @param {string[]} ev.data.groupby
   */
  _onSearch(ev) {
    ev.stopPropagation();
    this.reload(_.extend({ offset: 0, groupsOffset: 0 }, ev.data));
  }
  /**
   * @private
   * @param {OdooEvent} ev
   * @param {Array[]} ev.data.domain the current domain of the searchPanel
   */
  _onSearchPanelDomainUpdated(ev) {
    this.searchPanelDomain = ev.data.domain;
    this.reload({ offset: 0 });
  }
  /**
   * Intercepts the 'switch_view' event to add the controllerID into the data,
   * and lets the event bubble up.
   *
   * @param {OdooEvent} ev
   */
  _onSwitchView(ev) {
    ev.data.controllerID = this.controllerID;
  }
}

export class BasicController extends Controller {
  static get properties() {
    return {
      app: { type: Object },
      controlPanel: { type: Object }, // 控制面板
      viewManager: { type: Object }, // 视图控制器
      fields: { type: Array }, // 授权访问的字段 View 的所有可用Fields对象集合
      searchable: { type: Boolean }, // 可搜索并显示搜索过滤器
      searchView: { type: Object }, // 搜索视图

      datasource: { type: Object },
      params: { type: Object }, // 查询参数
      ///searchNode: { type: Object }, // # search view node
    };
  }

  constructor() {
    super();
    this.fields = [];
    this.searchable = false;
  }

  firstUpdated() {
    super.firstUpdated();
  }

  /**
   * Saves the record whose ID is given if necessary (@see _saveRecord).
   *
   * @param {string} [recordID] - default to main recordID
   * @param {Object} [options]
   * @returns {Deferred}
   *        Resolved with the list of field names (whose value has been modified)
   *        Rejected if the record can't be saved
   */
  saveRecord(recordID, options) {
    self._saveRecord(recordID, options);
  }

  _saveRecord(recordID, options) {
    recordID = recordID || this.handle;
    options = _.defaults(options || {}, {
      stayInEdit: false,
      reload: true,
      savePoint: false,
    });

    // Check if the view is in a valid state for saving
    // Note: it is the model's job to do nothing if there is nothing to save
    if (this.canBeSaved(recordID)) {
      var self = this;
      var saveDef = this.model.save(recordID, {
        // Save then leave edit mode
        reload: options.reload,
        savePoint: options.savePoint,
        viewType: options.viewType,
      });
      if (!options.stayInEdit) {
        saveDef = saveDef.then(function (fieldNames) {
          var def = fieldNames.length
            ? self._confirmSave(recordID)
            : self._setMode('readonly', recordID);
          return def.then(function () {
            return fieldNames;
          });
        });
      }
      return saveDef;
    } else {
      return $.Deferred().reject(); // Cannot be saved
    }
  }

  _deleteRecords(ids) {
    var self = this;
    function doIt() {
      return self.model
        .deleteRecords(ids, self.modelName)
        .then(self._onDeletedRecords.bind(self, ids));
    }
    if (this.confirmOnDelete) {
      Dialog.confirm(
        this,
        _t('Are you sure you want to delete this record ?'),
        {
          confirm_callback: doIt,
        }
      );
    } else {
      doIt();
    }
  }

  /**
   * Discards the changes made to the record whose ID is given, if necessary.
   * Automatically leaves to default mode for the given record.
   *
   * @private
   * @param {string} [recordID] - default to main recordID
   * @param {Object} [options]
   * @param {boolean} [options.readonlyIfRealDiscard=false]
   *        After discarding record changes, the usual option is to make the
   *        record readonly. However, the action manager calls this function
   *        at inappropriate times in the current code and in that case, we
   *        don't want to go back to readonly if there is nothing to discard
   *        (e.g. when switching record in edit mode in form view, we expect
   *        the new record to be in edit mode too, but the view manager calls
   *        this function as the URL changes...) @todo get rid of this when
   *        the webclient/action_manager's hashchange mechanism is improved.
   * @returns {Deferred}
   */
  _discardChanges(recordID, options) {
    var self = this;
    recordID = recordID || this.handle;
    return this.canBeDiscarded(recordID).then(function (needDiscard) {
      if (options && options.readonlyIfRealDiscard && !needDiscard) {
        return;
      }
      self.model.discardChanges(recordID);
      if (self.model.canBeAbandoned(recordID)) {
        self._abandonRecord(recordID);
        return;
      }
      return self._confirmSave(recordID);
    });
  }
  /**
   * Disables buttons so that they can't be clicked anymore.
   *
   * @private
   */
  _disableButtons() {
    if (this.$buttons) {
      this.$buttons.find('button').attr('disabled', true);
    }
  }

  /**
   * Enables buttons so they can be clicked again.
   *
   * @private
   */
  _enableButtons() {
    if (this.$buttons) {
      this.$buttons.find('button').removeAttr('disabled');
    }
  }

  /**
   * Ask the renderer if all associated field widget are in a valid state for
   * saving (valid value and non-empty value for required fields). If this is
   * not the case, this notifies the user with a warning containing the names
   * of the invalid fields.
   *
   * Note: changing the style of invalid fields is the renderer's job.
   *
   * @param {string} [recordID] - default to main recordID
   * @return {boolean}
   */
  canBeSaved(recordID) {
    var fieldNames = this.renderer.canBeSaved(recordID || this.handle);
    if (fieldNames.length) {
      this._notifyInvalidFields(fieldNames);
      return false;
    }
    return true;
  }
}

customElements.define('v-controller', Controller);
customElements.define('b-controller', BasicController);
