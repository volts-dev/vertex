import { html, css } from 'lit-element';

/*
import '@polymer/app-layout/app-header-layout/app-header-layout.js';
import '@polymer/app-layout/app-header/app-header.js';
import '@polymer/app-layout/app-toolbar/app-toolbar.js';

import '@polymer/iron-selector/iron-selector.js';
import '@polymer/iron-icons/iron-icons.js';
import '@polymer/iron-icons/av-icons.js';
import '@polymer/iron-media-query/iron-media-query.js';
import '@polymer/paper-drawer-panel/paper-drawer-panel.js';
import '@polymer/paper-button/paper-button.js';
import '@polymer/paper-item/paper-item.js';
import '@polymer/paper-listbox/paper-listbox.js';
import '@polymer/paper-menu-button/paper-menu-button.js';
*/

import { Page } from './page.js';
import { DataManager } from '@/service/data_manager';
import { Concurrency } from '@/core/concurrency.js';
import '@/components/app-icon/app-icon.js';
import '@/components/app-toolbar/app-toolbar';
import '@/elements/ve-dropdown-button/ve-dropdown-button.js';
import { ActionManager } from '@/views/action-manager/action_manager';
import store from '@/store';
import router from '@/router';

// a web client to show application content
class PageMain extends Page {
  static get styles() {
    return css`
      a {
        /*color: #222;*/
        text-decoration: none;
      }

      app-header {
        z-index: 1;
        font-size: 13px;
        border: 1px solid #5f5e97;
        background-color: #7c7bad;
      }

      .application_switcher {
        height: 100%;
        background-position: center;
        /*  background-image: url(/web_enterprise/static/src/img/application-switcher-bg.jpg);*/
        -webkit-background-size: cover;
        -moz-background-size: cover;
        -o-background-size: cover;
        background-size: cover;
        background-color: #797083;
        background: -moz-linear-gradient(135deg, #797083, #c59b9c);
        background: -o-linear-gradient(135deg, #c59b9c, #797083);
        background: -webkit-gradient(
          linear,
          left top,
          right bottom,
          from(#c59b9c),
          to(#797083)
        );
        background: -ms-linear-gradient(top, #c59b9c, #797083);
      }

      .home-grid {
        display: flex;
        max-width: 768px;
        margin: auto;
      }

      .menu_sections {
        height: 100%;
      }
      .menu_sections > * {
        height: 100%;
      }
      .menu_systray {
        position: relative;
        float: right;
      }

      .menu_brand {
        display: block;
        float: left;
        margin-right: 35px;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        color: white;
        font-size: 22px;
        font-weight: 500;
        line-height: 44px;
      }
      .iron-pages {
        width: 100%;
        height: 100%;
      }
    `;
  }

  render() {
    return html`
      <div class="iron-pages">
        <app-toolbar></app-toolbar>
        <!-- main pages -->
        <div class="app_main application_switcher">
          ${!store.app.activeAppId
            ? html`
                <div class="home-grid ">
                  ${Object.values(store.app.apps).map(
                    item => html`
                      <app-icon .item="${item}" class="package flex-none">
                      </app-icon>
                    `
                  )}
                </div>
              `
            : html` <!-- Action Manager --> `}
        </div>
      </div>
    `;
  }

  get renderDrawer() {
    return html`
      ${this.smallScreen
        ? html`
            <app-toolbar class="profileBar">
              <img
                class="profilePic"
                src="//app-layout-assets.appspot.com/assets/shrine/profile_pic.jpg"
                width="30"
                height="30"
              />
              <div class="profileName">Stella</div>
              <paper-icon-button
                icon="settings"
                aria-label="Settings"
              ></paper-icon-button>
            </app-toolbar>
            <div class="list">
              ${this.sections.map(
                section =>
                  html`
                    <a
                      href="#${section}"
                      class$="${this.getSectionClass(index, selectedTab)}"
                      >${section}</a
                    >
                  `
              )}
            </div>
          `
        : null}
    `;
  }

  static get properties() {
    return {
      action_manager: Object,

      //-------------------------------------------------------
      //  使用参数监听Query变动
      //-------------------------------------------------------
      title: String,
    };
  }

  constructor() {
    super();
    this.menus = [];
    this.ds_action = document.createElement('data-source');
    this.ds_action.method = 'POST';
    this.ds_view = document.createElement('data-source');
    this.ds_view.method = 'POST';

    this.menu_dp = new Concurrency.DropPrevious();

    this.customEvents();
  }

  connectedCallback(...args) {
    super.connectedCallback(...args);
    router.Listener('route', this.route.bind(this));
  }

  customEvents() {
    this.addEventListener('load_action', this._onLoadAction.bind(this));
    this.addEventListener('load_views', this._onLoadViews.bind(this));
  }

  /**
   * Loads an action from the database given its ID.
   *
   * @private
   * @param {OdooEvent} event
   * @param {integer} event.data.actionID
   * @param {Object} event.data.context
   * @param {function} event.data.on_success
   */
  _onLoadAction(event) {
    // 取缓存数据
    DataManager.load_action(event.detail.actionID, event.detail.context).then(
      event.detail.on_success
    );
  }

  _onLoadViews(event) {
    var params = {
      model: event.detail.modelName,
      context: event.detail.context,
      views_descr: event.detail.views,
    };

    return DataManager.load_views(params, event.detail.options || {}).then(
      event.detail.on_success
    );
  }

  // 开始整个Web Client主线
  firstUpdated() {
    super.firstUpdated();
    // start client
    this.startClient();
  }

  // --------------------------------------------------------------
  // URL state handling
  // --------------------------------------------------------------
  route(event) {
    if (this._ignore_hashchange) {
      this._ignore_hashchange = false;
      return Promise.resolve();
    }

    // # 切换页面
    store.app.setActiveApp(
      router.hash === '#home' || (!router.query && !router.hash)
        ? null
        : router.Query().menu
    );

    var self = this;
    return this.clear_uncommitted_changes().then(
      function () {
        //var stringstate = $.bbq.getState(false);
        var stringstate = router.url; // 获取当前URL
        if (!_.isEqual(self._current_state, stringstate)) {
          //对比是否更新
          var state = router.Query();
          if (state.action || (state.model && (state.view_type || state.id))) {
            // 加载ACTION_MANAGER
            return self.menu_dp
              .add(self.action_manager.LoadState(state, !!self._current_state))
              .then(function () {
                if (state.menu) {
                  // if (state.menu !== self.menu.current_primary_menu) {
                  //if (state.menu !== self.parentMenu) {
                  //OSV.bus.trigger('change_menu_section', state.menu);
                  //}
                } else {
                  var action = self.action_manager.getCurrentAction();
                  if (action) {
                    // TODO
                    //var menu_id = self.menu.action_id_to_primary_menu_id(action.id);
                    //OSV.bus.trigger("change_menu_section", menu_id);
                  }
                }
              });
          } else if (state.menu) {
            // 如果只有menu改变
            var menu = store.app.getMenuById(state.menu);
            var action_id = menu.action_id; //self.menu.menu_id_to_action_id(state.menu);
            if (action_id) {
              return self.menu_dp
                .add(self.do_action(action_id, { clear_breadcrumbs: true }))
                .then(function () {
                  // 废弃 触发菜单更新
                  //.bus.trigger('change_menu_section', state.menu);
                });
            }
          } else {
            store.app.setActiveApp(null);
          }
        }
        self._current_state = stringstate;
      } /*,
      function () {
        if (event) {
          self._ignore_hashchange = true;
          window.location = event.originalEvent.oldURL;
        }
      }*/
    );
  }

  startClient() {
    var self = this;
    this._title_changed();
    return store.session
      .is_bound()
      .then(function () {
        //self.$el.toggleClass('o_rtl', _t.database.parameters.direction === "rtl");
        self.bind_events();
        return Promise.all([self.set_action_manager(), self.set_loading()]);
      })
      .then(function () {
        if (store.session.session_is_valid()) {
          return self.show_application();
        } else {
          // database manager needs the webclient to keep going even
          // though it has no valid session
          return Promise.resolve();
        }
      })
      .then(function () {
        /*
                // Listen to 'scroll' event and propagate it on main bus
                self.action_manager.$el.on('scroll', core.bus.trigger.bind(core.bus, 'scroll'));
                core.bus.trigger('web_client_ready');
                odoo.isReady = true;
                if (session.uid === 1) {
                    self.$el.addClass('o_is_superuser');
                }
                */
      });
  }
  // 绑定整个APP事件
  bind_events() {}

  _title_changed() {}

  set_action_manager() {
    var self = this;
    // 创建ActionManager并插入主节点
    this.action_manager = new ActionManager(this, store.session.user_context);
    var fragment = document.createDocumentFragment();
    new Promise(function (resolve, reject) {
      // 组织节点
      fragment.append(self.action_manager);
      resolve();
    }).then(function () {
      // 插入节点到app_main主节点
      var node = self.shadowRoot.querySelector('.app_main');
      node.append(fragment);
    });
  }

  set_loading() {
    // this.loading = new Loading(this);
    //  return this.loading.appendTo(this.$el);
  }

  // --------------------------------------------------------------
  // Window title handling
  // --------------------------------------------------------------
  /**
   * Sets the first part of the title of the window, dedicated to the current action.
   */
  set_title(title) {
    //  this.set_title_part("action", title);
  }

  show_application() {
    var self = this;
    this.set_title();

    return this.menu_dp.add(this.instanciate_menu_widgets()).then(function () {
      // 监听URL变动
      // $(window).bind("hashchange", self.on_hashchange);
      // window.addEventListener("hashchange", self.on_hashchange.bind(this));
      // self.router.addEventListener('changed', self.on_hashchange.bind(self));

      // 跟新用户信息
      // If the url's state is empty, we execute the user's home action if there is one (we
      // show the first app if not)
      //var state = $.bbq.getState(true);
      var state = router.Query();
      if (_.keys(state).length === 1 && _.keys(state)[0] === 'cids') {
        return self.menu_dp
          .add(
            self._rpc({
              model: 'res.users',
              method: 'read',
              args: [session.uid, ['action_id']],
            })
          )
          .then(function (result) {
            var data = result[0];
            if (data.action_id) {
              return self.do_action(data.action_id[0]).then(function () {
                //self.menu.change_menu_section(self.menu.action_id_to_primary_menu_id(data.action_id[0]));
              });
            } else {
              //self.menu.openFirstApp();
            }
          });
      } else {
        return self.route();
      }
    });
  }

  // TODO
  instanciate_menu_widgets() {
    var self = this;
    //var defs = [];
    return this.load_menus().then(function (menuData) {
      // Here, we instanciate every menu widgets and we immediately append them into dummy
      // document fragments, so that their `start` method are executed before inserting them
      // into the DOM.
      /*if (self.menu) {
        self.menu.destroy();
      }
      self.menu = new Menu(self, menuData);
      defs.push(self.menu.prependTo(self.$el));
      
      return $.when.apply($, defs);
      */
      return Promise.resolve([]);
    });
  }

  // 产生菜单Map
  // 当界面直接由Url参数生成时需要
  // 1.确认当前菜单所归属的App
  // 2.确认Url是否执行Action并渲染View
  // 加载菜单数据
  load_menus() {
    var self = this;
    return this._rpc({
      route: '/app/menu/load',
      // args: [config.debug],
      //context: session.user_context,
    }).then(function (menuData) {
      store.app.UpdateMenus(menuData); // store API
      return menuData;
    });
  }

  clear_uncommitted_changes() {
    return this.action_manager.clearUncommittedChanges();
  }

  // --------------------------------------------------------------
  // do_*
  // --------------------------------------------------------------
  /**
   * When do_action is performed on the WebClient, forward it to the main ActionManager
   * This allows to widgets that are not inside the ActionManager to perform do_action
   */
  do_action() {
    return this.action_manager.doAction.apply(this.action_manager, arguments);
  }
}

customElements.define('page-main', PageMain);
