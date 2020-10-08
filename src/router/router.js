import { parseParams, parseQuery, testRoute } from './utils.js';
import { makeObservable, observable } from 'mobx';

export class Router {
  static get properties() {
    return {
      url: String,
      path: String,
      query: String,
      hash: String,
      route: String,
      queryParams: Object,
      hashParams: Object,
      canceled: { type: Boolean },
    };
  }

  //route = '';

  constructor(routes) {
    this.route = '';

    makeObservable(this, {
      route: observable,
    });

    this.canceled = false;
    this._location = window.location;

    this._routing(routes, (...args) => this.router(...args));
    window.addEventListener('inner_route', () => {
      this._routing(routes, (...args) => this.router(...args));
    });

    window.onpopstate = () => {
      window.dispatchEvent(new CustomEvent('inner_route'));
    };

    window.addEventListener('hashchange', this._hashChanged.bind(this));

    // 监听所有连接变动
    document.body.addEventListener('click', this._globalOnClick.bind(this));
  }

  firstUpdated() {
    super.firstUpdated();
    // this._routing(this.constructor.routes, (...args) => this.router(...args));
  }

  _hashChanged(e) {
    this.hash = window.decodeURIComponent(this._location.hash.substring(1));
  }

  /**
   * A necessary evil so that links work as expected. Does its best to
   * bail out early if possible.
   *
   * @param {MouseEvent} event .
   */
  _globalOnClick(event) {
    // If another event handler has stopped this event then there's nothing
    // for us to do. This can happen e.g. when there are multiple
    // iron-location elements in a page.
    if (event.defaultPrevented) {
      return;
    }

    var href = this._getSameOriginLinkHref(event);

    if (!href) {
      return;
    }

    event.preventDefault();

    // If the navigation is to the current page we shouldn't add a history
    // entry or fire a change event.
    if (href === this._location.href) {
      return;
    }

    window.history.pushState({}, null, href);
    window.dispatchEvent(new CustomEvent('inner_route'));
  }
  /**
   * Returns the absolute URL of the link (if any) that this click event
   * is clicking on, if we can and should override the resulting full
   * page navigation. Returns null otherwise.
   *
   * @param {MouseEvent} event .
   * @return {string?} .
   */
  _getSameOriginLinkHref(event) {
    // We only care about left-clicks.
    if (event.button !== 0) {
      return null;
    }

    // We don't want modified clicks, where the intent is to open the page
    // in a new tab.
    if (event.metaKey || event.ctrlKey || event.shiftKey) {
      return null;
    }

    var eventPath = event.path;
    var anchor = null;

    for (var i = 0; i < eventPath.length; i++) {
      var element = eventPath[i];

      if (element.tagName === 'A' && element.href) {
        anchor = element;
        break;
      }
    }

    // If there's no link there's nothing to do.
    if (!anchor) {
      return null;
    }

    // Target blank is a new tab, don't intercept.
    if (anchor.target === '_blank') {
      return null;
    }

    // If the link is for an existing parent frame, don't intercept.
    if (
      (anchor.target === '_top' || anchor.target === '_parent') &&
      window.top !== window
    ) {
      return null;
    }

    // If the link is a download, don't intercept.
    if (anchor.download) {
      return null;
    }

    var href = anchor.href;

    // It only makes sense for us to intercept same-origin navigations.
    // pushState/replaceState don't work with cross-origin links.
    var url;

    if (document.baseURI != null) {
      url = new URL(href, /** @type {string} */ (document.baseURI));
    } else {
      url = new URL(href);
    }

    var origin;

    // IE Polyfill
    if (this._location.origin) {
      origin = this._location.origin;
    } else {
      origin = this._location.protocol + '//' + this._location.host;
    }

    var urlOrigin;

    if (url.origin) {
      urlOrigin = url.origin;
    } else {
      // IE always adds port number on HTTP and HTTPS on <a>.host but not on
      // window.location.host
      var urlHost = url.host;
      var urlPort = url.port;
      var urlProtocol = url.protocol;
      var isExtraneousHTTPS = urlProtocol === 'https:' && urlPort === '443';
      var isExtraneousHTTP = urlProtocol === 'http:' && urlPort === '80';

      if (isExtraneousHTTPS || isExtraneousHTTP) {
        urlHost = url.hostname;
      }
      urlOrigin = urlProtocol + '//' + urlHost;
    }

    if (urlOrigin !== origin) {
      return null;
    }

    var normalizedHref = url.pathname + url.search + url.hash;

    // pathname should start with '/', but may not if `new URL` is not supported
    if (normalizedHref[0] !== '/') {
      normalizedHref = '/' + normalizedHref;
    }

    // If we've been configured not to handle this url... don't handle it!
    if (this._urlSpaceRegExp && !this._urlSpaceRegExp.test(normalizedHref)) {
      return null;
    }

    // Need to use a full URL in case the containing page has a base URI.
    var fullNormalizedHref = new URL(normalizedHref, this._location.href).href;
    return fullNormalizedHref;
  }

  // 根据现有参数重组路径
  _getUrl() {
    var partiallyEncodedPath = window
      .encodeURI(this.path)
      .replace(/\#/g, '%23')
      .replace(/\?/g, '%3F');
    var partiallyEncodedQuery = '';
    if (this.query) {
      partiallyEncodedQuery = '?' + this.query.replace(/\#/g, '%23');
      if (this.encodeSpaceAsPlusInQuery) {
        partiallyEncodedQuery = partiallyEncodedQuery
          .replace(/\+/g, '%2B')
          .replace(/ /g, '+')
          .replace(/%20/g, '+');
      } else {
        // required for edge
        partiallyEncodedQuery = partiallyEncodedQuery
          .replace(/\+/g, '%2B')
          .replace(/ /g, '%20');
      }
    }
    var partiallyEncodedHash = '';
    if (this.hash) {
      partiallyEncodedHash = '#' + window.encodeURI(this.hash);
    }
    return partiallyEncodedPath + partiallyEncodedQuery + partiallyEncodedHash;
  }

  router(route, params, query, data) {
    /* this.route = route;
      this.params = params;
      this.query = query;
      this.data = data;
      console.log(route, params, query, data);*/
  }

  routed(name, params, query, data, callback, localCallback) {
    this.url = this._location.href;
    this.route = name;
    this._hashChanged();

    localCallback && localCallback(name, params, query, data);
    callback(name, params, query, data);

    window.dispatchEvent(
      new CustomEvent('route', { bubbles: true, composed: true })
    );
  }

  _routing(routes, callback) {
    this.canceled = true;

    this.path = decodeURI(window.location.pathname);
    this.query = decodeURI(window.location.search.substring(1));
    this.hash = window.decodeURIComponent(this._location.hash.substring(1));

    this.queryParams = parseQuery(this.query);
    this.hashParams = parseQuery(this.hash);

    let notFoundRoute = routes.filter(route => route.pattern === '*')[0];
    let activeRoute = routes.filter(
      route => route.pattern !== '*' && testRoute(this.path, route.pattern)
    )[0];

    if (activeRoute) {
      activeRoute.params = parseParams(activeRoute.pattern, this.path);
      activeRoute.data = activeRoute.data || {};
      if (
        activeRoute.authentication &&
        activeRoute.authentication.authenticate &&
        typeof activeRoute.authentication.authenticate === 'function'
      ) {
        this.canceled = false;
        Promise.resolve(
          activeRoute.authentication.authenticate.bind(this).call()
        ).then(authenticated => {
          if (!this.canceled) {
            if (authenticated) {
              if (
                activeRoute.authorization &&
                activeRoute.authorization.authorize &&
                typeof activeRoute.authorization.authorize === 'function'
              ) {
                this.canceled = false;
                Promise.resolve(
                  activeRoute.authorization.authorize.bind(this).call()
                ).then(authorizatied => {
                  if (!this.canceled) {
                    if (authorizatied) {
                      this.routed(
                        activeRoute.name,
                        activeRoute.params,
                        this.queryParams,
                        activeRoute.data,
                        callback,
                        activeRoute.callback
                      );
                    } else {
                      this.routed(
                        activeRoute.authorization.unauthorized.name,
                        activeRoute.params,
                        this.queryParams,
                        activeRoute.data,
                        callback,
                        activeRoute.callback
                      );
                    }
                  }
                });
              } else {
                this.routed(
                  activeRoute.name,
                  activeRoute.params,
                  this.queryParams,
                  activeRoute.data,
                  callback,
                  activeRoute.callback
                );
              }
            } else {
              this.routed(
                activeRoute.authentication.unauthenticated.name,
                activeRoute.params,
                this.queryParams,
                activeRoute.data,
                callback,
                activeRoute.callback
              );
            }
          }
        });
      } else if (
        activeRoute.authorization &&
        activeRoute.authorization.authorize &&
        typeof activeRoute.authorization.authorize === 'function'
      ) {
        this.canceled = false;
        Promise.resolve(
          activeRoute.authorization.authorize.bind(this).call()
        ).then(authorizatied => {
          if (!this.canceled) {
            if (authorizatied) {
              this.routed(
                activeRoute.name,
                activeRoute.params,
                this.queryParams,
                activeRoute.data,
                callback,
                activeRoute.callback
              );
            } else {
              this.routed(
                activeRoute.authorization.unauthorized.name,
                activeRoute.params,
                this.queryParams,
                activeRoute.data,
                callback,
                activeRoute.callback
              );
            }
          }
        });
      } else {
        this.routed(
          activeRoute.name,
          activeRoute.params,
          this.queryParams,
          activeRoute.data,
          callback,
          activeRoute.callback
        );
      }
    } else if (notFoundRoute) {
      notFoundRoute.data = notFoundRoute.data || {};
      this.routed(
        notFoundRoute.name,
        {},
        this.queryParams,
        notFoundRoute.data,
        callback,
        notFoundRoute.callback
      );
    }
  }

  Listener(type, listener) {
    window.addEventListener(type, listener);
  }

  Push(url) {
    this.url = url;
    window.history.pushState({}, null, this.url);
    window.dispatchEvent(new CustomEvent('inner_route'));
  }

  Path() {
    return this._getUrl();
  }

  Hash() {
    return this.hashParams;
  }

  Query() {
    return this.queryParams;
  }

  // API set the map of query to change url
  SetQuery(query) {
    if (!query) {
      return;
    }

    // 匹对跟新内容
    var temp = _.clone(this.queryParams);
    for (var key in query) {
      var value = query[key];
      if (key) {
        if (!value) {
          //undefined
          delete temp[key];
          continue;
        }

        temp[key] = value;
      }
    }

    if (this.queryParams == temp) {
      return;
    }

    this.queryParams = temp; // this._rebuildQueryMap(temp);
    //this.lastQuery = this.query; // 记录变更前的Query

    // 格式化生成新Query
    var params = [];
    for (var key in this.queryParams) {
      params.push(key + '=' + this.queryParams[key]);
    }

    var querystring = params.join('&');
    if (this.query != querystring) {
      this.query = querystring;
      this.Push(this._getUrl());
    }
  }

  Get(key) {
    return this.queryParams[key];
  }
}
