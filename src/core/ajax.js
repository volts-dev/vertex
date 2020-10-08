import { Time } from './time';
import { Concurrency } from './concurrency';
import { Translation } from './translation';
import { Download } from '@/libs/download';
import { Parse } from '@/libs/content-disposition';
var _t = Translation._t;

// 交互式调用
// Create the final object containing all the functions first to allow monkey
// patching them correctly if ever needed.
let Ajax = {};

// helper function to make a rpc with a function name hardcoded to 'call'
function rpc(url, params, settings) {
  return jsonRpc(url, 'call', params, settings);
}

/**
 * Perform a RPC.  Please note that this is not the preferred way to do a
 * rpc if you are in the context of a widget.  In that case, you should use
 * the this._rpc method.
 *
 * @param {Object} params @see buildQuery for a description
 * @param {Object} options
 * @returns {Promise<any>}
 */
function query(params, options) {
  var query = this.buildQuery(params);
  return rpc(query.route, query.params, options);
}

/**
 * @param {Object} options
 * @param {any[]} [options.args]
 * @param {Object} [options.context]
 * @param {any[]} [options.domain]
 * @param {string[]} [options.fields]
 * @param {string[]} [options.groupBy]
 * @param {Object} [options.kwargs]
 * @param {integer|false} [options.limit]
 * @param {string} [options.method]
 * @param {string} [options.model]
 * @param {integer} [options.offset]
 * @param {string[]} [options.orderBy]
 * @param {Object} [options.params]
 * @param {string} [options.route]
 * @returns {Object} with 2 keys: route and params
 */
function buildQuery(options) {
  var route;
  var params = options.params || {};
  var orderBy;
  if (options.route) {
    route = options.route;
  } else if (options.model && options.method) {
    route = '/app/dataset/call_kw/' + options.model + '/' + options.method;
  }
  if (options.method) {
    params.args = options.args || [];
    params.model = options.model;
    params.method = options.method;
    params.kwargs = _.extend(params.kwargs || {}, options.kwargs);
    params.kwargs.context =
      options.context || params.context || params.kwargs.context;
  }

  if (options.method === 'read_group' || options.method === 'web_read_group') {
    /*
        if (!(params.args && params.args[0] !== undefined)) {
          params.kwargs.domain = options.domain || params.domain || params.kwargs.domain || [];
        }
        if (!(params.args && params.args[1] !== undefined)) {
          params.kwargs.fields = options.fields || params.fields || params.kwargs.fields || [];
        }
        if (!(params.args && params.args[2] !== undefined)) {
          params.kwargs.groupby = options.groupBy || params.groupBy || params.kwargs.groupby || [];
        }
        params.kwargs.offset = options.offset || params.offset || params.kwargs.offset;
        params.kwargs.limit = options.limit || params.limit || params.kwargs.limit;
        // In kwargs, we look for "orderby" rather than "orderBy" (note the absence of capital B),
        // since the Python argument to the actual function is "orderby".
        orderBy = options.orderBy || params.orderBy || params.kwargs.orderby;
        params.kwargs.orderby = orderBy ? this._serializeSort(orderBy) : orderBy;
        params.kwargs.lazy = "lazy" in options ? options.lazy : params.lazy;

        if (options.method === "web_read_group") {
          params.kwargs.expand = options.expand || params.expand || params.kwargs.expand;
          params.kwargs.expand_limit = options.expand_limit || params.expand_limit || params.kwargs.expand_limit;
          var expandOrderBy = options.expand_orderby || params.expand_orderby || params.kwargs.expand_orderby;
          params.kwargs.expand_orderby = expandOrderBy ? this._serializeSort(expandOrderBy) : expandOrderBy;
        }
        */
  }

  if (options.method === 'search_read') {
    // call the model method
    params.kwargs.domain =
      options.domain || params.domain || params.kwargs.domain;
    params.kwargs.fields =
      options.fields || params.fields || params.kwargs.fields;
    params.kwargs.offset =
      options.offset || params.offset || params.kwargs.offset;
    params.kwargs.limit = options.limit || params.limit || params.kwargs.limit;
    // In kwargs, we look for "order" rather than "orderBy" since the Python
    // argument to the actual function is "order".
    orderBy = options.orderBy || params.orderBy || params.kwargs.order;
    params.kwargs.order = orderBy ? _serializeSort(orderBy) : orderBy;
  }

  if (options.route === '/app/dataset/search_read') {
    // specifically call the controller
    params.model = options.model || params.model;
    params.domain = options.domain || params.domain;
    params.fields = options.fields || params.fields;
    params.limit = options.limit || params.limit;
    params.offset = options.offset || params.offset;
    orderBy = options.orderBy || params.orderBy;
    params.sort = orderBy ? _serializeSort(orderBy) : orderBy;
    params.context = options.context || params.context || {};
  }

  return {
    route: route,
    params: JSON.parse(JSON.stringify(params)),
  };
}

/**
 * Helper method, generates a string to describe a ordered by sequence for
 * SQL.
 *
 * For example, [{name: 'foo'}, {name: 'bar', asc: false}] will
 * be converted into 'foo ASC, bar DESC'
 *
 * @param {Object[]} orderBy list of objects {name:..., [asc: ...]}
 * @returns {string}
 */
function _serializeSort(orderBy) {
  return _.map(orderBy, function (order) {
    return order.name + (order.asc !== false ? ' ASC' : ' DESC');
  }).join(', ');
}

function _genericJsonRpc(fct_name, params, settings, fct) {
  var shadow = settings.shadow || false;
  delete settings.shadow;
  if (!shadow) {
    //OSV.bus.trigger("rpc_request");
  }

  var data = {
    id: Math.floor(Math.random() * 1000 * 1000 * 1000),
    jsonrpc: '2.0',
    method: fct_name,
    params: params,
  };

  var xhr = fct(data);
  var promise = xhr
    .then(response => response.json())
    .then(function (result) {
      if (result) {
        //OSV.bus.trigger("rpc:result", data, result);
        if (result.error !== undefined) {
          if (
            result.error.data.arguments[0] !==
            'bus.Bus not available in test mode'
          ) {
            console.debug(
              'Server application error\n',
              'Error code:',
              result.error.code,
              '\n',
              'Error message:',
              result.error.message,
              '\n',
              'Error data message:\n',
              result.error.data.message,
              '\n',
              'Error data debug:\n',
              result.error.data.debug
            );
          }
          return Promise.reject({ type: 'server', error: result.error });
        } else {
          if (!shadow) {
            //OSV.bus.trigger("rpc_response");
          }
          //resolve(result);
          return result.result;
        }
        // self.lastResponse = data.result;
        // self.dispatchEvent(new CustomEvent("response", { detail: self.lastResponse }));
        // 返回数据
        //resolve(data.result);
      }
    })
    .catch(reason => {
      console.error('JsonRPC communication error', _.toArray(arguments));
      reason = {
        type: 'communication',
        error: arguments[0],
        textStatus: arguments[1],
        errorThrown: arguments[2],
      };
      return Promise.reject(reason);
    });

  var rejection;

  var new_promise = new Promise(function (resolve, reject) {
    rejection = reject;

    promise
      .then(function (result) {
        if (!shadow) {
          //OSV.bus.trigger("rpc_response");
        }
        resolve(result);
      })
      .catch(reason => {
        //self.lastError = err;
        var type = reason.type;
        var error = reason.error;
        var textStatus = reason.textStatus;
        var errorThrown = reason.errorThrown;
        if (type === 'server') {
          if (!shadow) {
            //OSV.bus.trigger("rpc_response");
          }
          if (error.code === 100) {
            //OSV.bus.trigger("invalidate_session");
          }
          reject({ message: error, event: new Event('') });
        } else {
          if (!shadow) {
            //OSV.bus.trigger("rpc_response_failed");
          }
          var nerror = {
            code: -32098,
            message: 'XmlHttpRequestError ' + errorThrown,
            data: {
              type: 'xhr' + textStatus,
              debug: error, // error.responseText,
              objects: [error, errorThrown],
              arguments: [reason || textStatus],
            },
          };
          reject({ message: nerror, event: new Event('') });
        }
      });
  });

  // FIXME: jsonp?
  new_promise.abort = function () {
    rejection({
      message: 'XmlHttpRequestError abort',
      event: new Event('abort'),
    });
    if (xhr.abort) {
      xhr.abort();
    }
  };

  new_promise.guardedCatch(function (reason) {
    // Allow promise user to disable rpc_error call in case of failure
    setTimeout(function () {
      // we want to execute this handler after all others (hence
      // setTimeout) to let the other handlers prevent the event
      if (!reason.event.defaultPrevented) {
        //OSV.bus.trigger("rpc_error", reason.message, reason.event);
      }
    }, 0);
  });

  return new_promise;
}

function jsonRpc(url, fct_name, params, settings) {
  settings = settings || {};
  return _genericJsonRpc(fct_name, params, settings, function (data) {
    // request object
    // Note that a request using the GET or HEAD method cannot have a body.

    const request = new Request(
      url,
      _.extend({}, settings, {
        method: 'POST',
        body: JSON.stringify(data, Time.date_to_utc),
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
      })
    );
    return fetch(request);

    /*
    var self = this;
    return new Promise(function (resolve, reject) {
      //setTimeout(function() {
      fetch(request)
        //.then(res => res.json())
        .then(function (resp) {
          // code to handle the response
          resp.json().then(function (data) {
            if (data) {
              // self.lastResponse = data.result;
              // self.dispatchEvent(new CustomEvent("response", { detail: self.lastResponse }));
              // 返回数据
              resolve(data.result);
            }
          });
        })
        .catch((err) => {
          //self.lastError = err;
          reject(err);
        });
      // }, self.timeout);
    });

    /* 废弃
    return $.ajax(
      url,
      _.extend({}, settings, {
        url: url,
        dataType: "json",
        type: "POST",
        data: JSON.stringify(data, Time.date_to_utc),
        contentType: "application/json",
      })
    );
    */
  });
}

/**
 * Load css asynchronously: fetch it from the url parameter and add a link tag
 * to <head>.
 * If the url has already been requested and loaded, the promise will resolve
 * immediately.
 *
 * @param {String} url of the css to be fetched
 * @returns {Promise} resolved when the css has been loaded.
 */
var loadCSS = (function () {
  /*ivar urlDefs = {};

  return function loadCSS(url) {
    f (url in urlDefs) {
      // nothing to do here
    } else if ($('link[href="' + url + '"]').length) {
      // the link is already in the DOM, the promise can be resolved
      urlDefs[url] = Promise.resolve();
    } else {
      var $link = $("<link>", {
        href: url,
        rel: "stylesheet",
        type: "text/css",
      });
      urlDefs[url] = new Promise(function(resolve, reject) {
        $link.on("load", function() {
          resolve();
        });
      });
      $("head").append($link);
    }
    return urlDefs[url];
   
  }; */
})();

var loadJS = (function () {
  /* var dependenciesPromise = {};

  var load = function loadJS(url) {
    // Check the DOM to see if a script with the specified url is already there
    var alreadyRequired = $('script[src="' + url + '"]').length > 0;

    // If loadJS was already called with the same URL, it will have a registered promise indicating if
    // the script has been fully loaded. If not, the promise has to be initialized.
    // This is initialized as already resolved if the script was already there without the need of loadJS.
    if (url in dependenciesPromise) {
      return dependenciesPromise[url];
    }
    var scriptLoadedPromise = new Promise(function(resolve, reject) {
      if (alreadyRequired) {
        resolve();
      } else {
        // Get the script associated promise and returns it after initializing the script if needed. The
        // promise is marked to be resolved on script load and rejected on script error.
        var script = document.createElement("script");
        script.type = "text/javascript";
        script.src = url;
        script.onload = script.onreadystatechange = function() {
          if ((script.readyState && script.readyState !== "loaded" && script.readyState !== "complete") || script.onload_done) {
            return;
          }
          script.onload_done = true;
          resolve(url);
        };
        script.onerror = function() {
          console.error("Error loading file", script.src);
          reject(url);
        };
        var head = document.head || document.getElementsByTagName("head")[0];
        head.appendChild(script);
      }
    });

    dependenciesPromise[url] = scriptLoadedPromise;
    return scriptLoadedPromise;
  };

  return load;*/
})();

/**
 * Cooperative file download implementation, for ajaxy APIs.
 *
 * Requires that the server side implements an httprequest correctly
 * setting the `fileToken` cookie to the value provided as the `token`
 * parameter. The cookie *must* be set on the `/` path and *must not* be
 * `httpOnly`.
 *
 * It would probably also be a good idea for the response to use a
 * `Content-Disposition: attachment` header, especially if the MIME is a
 * "known" type (e.g. text/plain, or for some browsers application/json
 *
 * @param {Object} options
 * @param {String} [options.url] used to dynamically create a form
 * @param {Object} [options.data] data to add to the form submission. If can be used without a form, in which case a form is created from scratch. Otherwise, added to form data
 * @param {HTMLFormElement} [options.form] the form to submit in order to fetch the file
 * @param {Function} [options.success] callback in case of download success
 * @param {Function} [options.error] callback in case of request error, provided with the error body
 * @param {Function} [options.complete] called after both ``success`` and ``error`` callbacks have executed
 * @returns {boolean} a false value means that a popup window was blocked. This
 *   mean that we probably need to inform the user that something needs to be
 *   changed to make it work.
 */
function get_file(options) {
  var xhr = new XMLHttpRequest();

  var data;
  if (options.form) {
    xhr.open(options.form.method, options.form.action);
    data = new FormData(options.form);
  } else {
    xhr.open('POST', options.url);
    data = new FormData();
    _.each(options.data || {}, function (v, k) {
      data.append(k, v);
    });
  }
  data.append('token', 'dummy-because-api-expects-one');
  if (OSV.csrf_token) {
    data.append('csrf_token', OSV.csrf_token);
  }
  // IE11 wants this after xhr.open or it throws
  xhr.responseType = 'blob';

  // onreadystatechange[readyState = 4]
  // => onload (success) | onerror (error) | onabort
  // => onloadend
  xhr.onload = function () {
    var mimetype = xhr.response.type;
    if (xhr.status === 200 && mimetype !== 'text/html') {
      // replace because apparently we send some C-D headers with a trailing ";"
      // todo: maybe a lack of CD[attachment] should be interpreted as an error case?
      var header = (xhr.getResponseHeader('Content-Disposition') || '').replace(
        /;$/,
        ''
      );
      var filename = header ? Parse(header).parameters.filename : null;

      Download(xhr.response, filename, mimetype);
      // not sure download is going to be sync so this may be called
      // before the file is actually fetched (?)
      if (options.success) {
        options.success();
      }
      return true;
    }

    if (!options.error) {
      return true;
    }
    var decoder = new FileReader();
    decoder.onload = function () {
      var contents = decoder.result;

      var err;
      var doc = new DOMParser().parseFromString(contents, 'text/html');
      var nodes =
        doc.body.children.length === 0
          ? doc.body.childNodes
          : doc.body.children;
      try {
        // Case of a serialized Odoo Exception: It is Json Parsable
        var node = nodes[1] || nodes[0];
        err = JSON.parse(node.textContent);
      } catch (e) {
        // Arbitrary uncaught python side exception
        err = {
          message: nodes.length > 1 ? nodes[1].textContent : '',
          data: {
            name: String(xhr.status),
            title: nodes.length > 0 ? nodes[0].textContent : '',
          },
        };
      }
      options.error(err);
    };
    decoder.readAsText(xhr.response);
  };
  xhr.onerror = function () {
    if (options.error) {
      options.error({
        message: _(
          'Something happened while trying to contact the server, check that the server is online and that you still have a working network connection.'
        ),
        data: { title: _t('Could not connect to the server') },
      });
    }
  };
  if (options.complete) {
    xhr.onloadend = function () {
      options.complete();
    };
  }

  xhr.send(data);
  return true;
}

function post(controller_url, data) {
  /* var postData = new FormData();

  $.each(data, function(i, val) {
    postData.append(i, val);
  });
  if (OSV.csrf_token) {
    postData.append("csrf_token", OSV.csrf_token);
  }

  return new Promise(function(resolve, reject) {
    $.ajax(controller_url, {
      data: postData,
      processData: false,
      contentType: false,
      type: "POST",
    })
      .then(resolve)
      .fail(reject);
  });
  */
}

/**
 * Loads an XML file according to the given URL and adds its associated qweb
 * templates to the given qweb engine. The function can also be used to get
 * the promise which indicates when all the calls to the function are finished.
 *
 * Note: "all the calls" = the calls that happened before the current no-args
 * one + the calls that will happen after but when the previous ones are not
 * finished yet.
 *
 * @param {string} [url] - an URL where to find qweb templates
 * @param {QWeb} [qweb] - the engine to which the templates need to be added
 * @returns {Promise}
 *          If no argument is given to the function, the promise's state
 *          indicates if "all the calls" are finished (see main description).
 *          Otherwise, it indicates when the templates associated to the given
 *          url have been loaded.
 */
var loadXML = (function () {
  // Some "static" variables associated to the loadXML function
  var isLoading = false;
  var loadingsData = [];
  var seenURLs = [];

  return function (url, qweb) {
    function _load() {
      isLoading = true;
      if (loadingsData.length) {
        // There is something to load, load it, resolve the associated
        // promise then start loading the next one
        var loadingData = loadingsData[0];
        loadingData.qweb.add_template(loadingData.url, function () {
          // Remove from array only now so that multiple calls to
          // loadXML with the same URL returns the right promise
          loadingsData.shift();
          loadingData.resolve();
          _load();
        });
      } else {
        // There is nothing to load anymore, so resolve the
        // "all the calls" promise
        isLoading = false;
      }
    }

    // If no argument, simply returns the promise which indicates when
    // "all the calls" are finished
    if (!url || !qweb) {
      return Promise.resolve();
    }

    // If the given URL has already been seen, do nothing but returning the
    // associated promise
    if (_.contains(seenURLs, url)) {
      var oldLoadingData = _.findWhere(loadingsData, { url: url });
      return oldLoadingData ? oldLoadingData.def : Promise.resolve();
    }
    seenURLs.push(url);

    // Add the information about the new data to load: the url, the qweb
    // engine and the associated promise
    var newLoadingData = {
      url: url,
      qweb: qweb,
    };
    newLoadingData.def = new Promise(function (resolve, reject) {
      newLoadingData.resolve = resolve;
      newLoadingData.reject = reject;
    });
    loadingsData.push(newLoadingData);

    // If not already started, start the loading loop (reinitialize the
    // "all the calls" promise to an unresolved state)
    if (!isLoading) {
      _load();
    }

    // Return the promise associated to the new given URL
    return newLoadingData.def;
  };
})();

/**
 * Loads a template file according to the given xmlId.
 *
 * @param {string} [xmlId] - the template xmlId
 * @param {Object} [context]
 *        additionnal rpc context to be merged with the default one
 * @param {string} [tplRoute='/web/dataset/call_kw/']
 * @returns {Deferred} resolved with an object
 *          cssLibs: list of css files
 *          cssContents: list of style tag contents
 *          jsLibs: list of JS files
 *          jsContents: list of script tag contents
 */
/*
var loadAsset = (function() {
  var cache = {};

  var load = function loadAsset(xmlId, context, tplRoute = "/web/dataset/call_kw/") {
    if (cache[xmlId]) {
      return cache[xmlId];
    }
    context = _.extend({}, odoo.session_info.user_context, context);
    const params = {
      args: [
        xmlId,
        {
          debug: Config.isDebug(),
        },
      ],
      kwargs: {
        context: context,
      },
    };
    if (tplRoute === "/web/dataset/call_kw/") {
      Object.assign(params, {
        model: "ir.ui.view",
        method: "render_template",
      });
    }
    cache[xmlId] = rpc(tplRoute, params)
      .then(function(xml) {
        var $xml = $(xml);
        return {
          cssLibs: $xml
            .filter('link[href]:not([type="image/x-icon"])')
            .map(function() {
              return $(this).attr("href");
            })
            .get(),
          cssContents: $xml
            .filter("style")
            .map(function() {
              return $(this).html();
            })
            .get(),
          jsLibs: $xml
            .filter("script[src]")
            .map(function() {
              return $(this).attr("src");
            })
            .get(),
          jsContents: $xml
            .filter("script:not([src])")
            .map(function() {
              return $(this).html();
            })
            .get(),
        };
      })
      .guardedCatch((reason) => {
        reason.event.preventDefault();
        throw `Unable to render the required templates for the assets to load: ${reason.message.message}`;
      });
    return cache[xmlId];
  };

  return load;
})();
*/
/**
 * Loads the given js/css libraries and asset bundles. Note that no library or
 * asset will be loaded if it was already done before.
 *
 * @param {Object} libs
 * @param {Array<string|string[]>} [libs.assetLibs=[]]
 *      The list of assets to load. Each list item may be a string (the xmlID
 *      of the asset to load) or a list of strings. The first level is loaded
 *      sequentially (so use this if the order matters) while the assets in
 *      inner lists are loaded in parallel (use this for efficiency but only
 *      if the order does not matter, should rarely be the case for assets).
 * @param {string[]} [libs.cssLibs=[]]
 *      The list of CSS files to load. They will all be loaded in parallel but
 *      put in the DOM in the given order (only the order in the DOM is used
 *      to determine priority of CSS rules, not loaded time).
 * @param {Array<string|string[]>} [libs.jsLibs=[]]
 *      The list of JS files to load. Each list item may be a string (the URL
 *      of the file to load) or a list of strings. The first level is loaded
 *      sequentially (so use this if the order matters) while the files in inner
 *      lists are loaded in parallel (use this for efficiency but only
 *      if the order does not matter).
 * @param {string[]} [libs.cssContents=[]]
 *      List of inline styles to add after loading the CSS files.
 * @param {string[]} [libs.jsContents=[]]
 *      List of inline scripts to add after loading the JS files.
 * @param {Object} [context]
 *        additionnal rpc context to be merged with the default one
 * @param {string} [tplRoute]
 *      Custom route to use for template rendering of the potential assets
 *      to load (see libs.assetLibs).
 *
 * @returns {Promise}
 */
function loadLibs(libs, context, tplRoute) {
  var mutex = new Concurrency.Mutex();
  /*
  mutex.exec(function () {
    var defs = [];
    var cssLibs = [libs.cssLibs || []]; // Force loading in parallel
    defs.push(
      _loadArray(cssLibs, ajax.loadCSS).then(function () {
        if (libs.cssContents && libs.cssContents.length) {
          $("head").append(
            $("<style/>", {
              html: libs.cssContents.join("\n"),
            })
          );
        }
      })
    );
    defs.push(
      _loadArray(libs.jsLibs || [], ajax.loadJS).then(function () {
        if (libs.jsContents && libs.jsContents.length) {
          $("head").append(
            $("<script/>", {
              html: libs.jsContents.join("\n"),
            })
          );
        }
      })
    );
    return Promise.all(defs);
  });
  mutex.exec(function () {
    return _loadArray(libs.assetLibs || [], function (xmlID) {
      return ajax.loadAsset(xmlID, context, tplRoute).then(function (asset) {
        return ajax.loadLibs(asset);
      });
    });
  });

  function _loadArray(array, loadCallback) {
    var _mutex = new Concurrency.Mutex();
    array.forEach(function (urlData) {
      _mutex.exec(function () {
        if (typeof urlData === "string") {
          return loadCallback(urlData);
        }
        return Promise.all(urlData.map(loadCallback));
      });
    });
    return _mutex.getUnlockedDef();
  }
*/
  return mutex.getUnlockedDef();
}

export default _.extend(Ajax, {
  jsonRpc: jsonRpc,
  rpc: rpc,
  query: query,
  buildQuery: buildQuery,
  loadCSS: loadCSS,
  loadJS: loadJS,
  loadXML: loadXML,
  // loadAsset: loadAsset,
  loadLibs: loadLibs,
  get_file: get_file,
  post: post,
});
