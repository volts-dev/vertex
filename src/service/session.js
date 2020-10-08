import { EventDispatcherMixin } from '@/mixins/EventDispatcherMixin';
import Ajax from '@/core/ajax';

// 加载插件
import './service_ajax';
import './service_notification';
import './service_session_storage';

// 用户会话实例 语言地区等
export class Session extends EventDispatcherMixin(Object) {
  session_is_valid() {
    /* var db = $.deparam.querystring().db;
    if (db && this.db !== db) {
      return false;
    }
    return !!this.uid;*/
    return true;
  }

  /**
   * Executes an RPC call, registering the provided callbacks.
   *
   * Registers a default error callback if none is provided, and handles
   * setting the correct session id and session context in the parameter
   * objects
   *
   * @param {String} url RPC endpoint
   * @param {Object} params call parameters
   * @param {Object} options additional options for rpc call
   * @returns {Promise}
   */
  rpc(url, params, options) {
    var self = this;
    options = _.clone(options || {});
    options.headers = _.extend({}, options.headers);

    // we add here the user context for ALL queries, mainly to pass
    // the allowed_company_ids key
    if (params && params.kwargs) {
      params.kwargs.context = _.extend(
        params.kwargs.context || {},
        this.user_context
      );
    }

    // TODO: remove
    if (!_.isString(url)) {
      _.extend(options, url);
      url = url.url;
    }
    if (self.use_cors) {
      url = self.url(url, null);
    }

    return Ajax.jsonRpc(url, 'call', params, options);
  }

  /**
   * Setup a session
   */
  // 用户会话绑定确认
  session_bind(origin) {
    this.setup(origin);
    //qweb.default_dict._s = this.origin;
    this.uid = null; // 用户UID
    this.username = null; // 用户名称
    this.user_context = {}; // 用户数据
    this.db = null; // 数据库名称
    this.active_id = null; // 用户激活状态
    return this.session_init(); // 初始化
  }

  setup(origin, options) {
    // must be able to customize server
    var window_origin = location.protocol + '//' + location.host;
    origin = origin ? origin.replace(/\/+$/, '') : window_origin;
    if (!_.isUndefined(this.origin) && this.origin !== origin)
      throw new Error('Session already bound to ' + this.origin);
    else this.origin = origin;
    this.prefix = this.origin;
    this.server = this.origin; // keep chs happy
    options = options || {};
    if ('use_cors' in options) {
      this.use_cors = options.use_cors;
    }
  }

  /**
   * Init a session, reloads from cookie, if it exists
   */
  session_init() {
    var self = this;
    var prom = this.session_reload();

    if (this.is_frontend) {
      return prom.then(function () {
        return self.load_translations();
      });
    }

    return prom.then(function () {
      // 使用Lit-element 不需要渲染器Qweb
      /*var modules = self.module_list.join(",");
      var promise = self.load_qweb(modules);
      if (self.session_is_valid()) {
        return promise.then(function () {
          return self.load_modules();
        });
      }*/
      return Promise.all([
        //promise,
        //self.rpc("/web/webclient/bootstrap_translations", { mods: self.module_list }).then(function (trans) {
        //  _t.database.set_bundle(trans);
        // }),
      ]);
    });
  }

  /**
   * (re)loads the content of a session: db name, username, user id, session
   * context and status of the support contract
   *
   * @returns {Promise} promise indicating the session is done reloading
   */
  session_reload() {
    var result = _.extend({}); //_.extend({}, window.odoo.session_info);
    _.extend(this, result);
    return Promise.resolve();
  }

  load_translations() {
    //return _t.database.load_translations(this, this.module_list, this.user_context.lang, this.translationURL);
  }

  /**
   * Returns the time zone difference (in minutes) from the current locale
   * (host system settings) to UTC, for a given date. The offset is positive
   * if the local timezone is behind UTC, and negative if it is ahead.
   *
   * @param {string | moment} date a valid string date or moment instance
   * @returns {integer}
   */
  getTZOffset(date) {
    return -new Date(date).getTimezoneOffset();
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------
  /**
   * Replaces the value of a key in cache_hashes (the hash of some resource computed on the back-end by a unique value
   * @param {string} key the key in the cache_hashes to invalidate
   */
  invalidateCacheKey(key) {
    if (this.cache_hashes && this.cache_hashes[key]) {
      this.cache_hashes[key] = Date.now();
    }
  }
}
