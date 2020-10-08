import Ajax from './ajax';

var throbbers = [];

function blockUI() {
  /*var tmp = $.blockUI.apply($, arguments);
  var throbber = new Throbber();
  throbbers.push(throbber);
  throbber.appendTo($(".oe_blockui_spin_container"));
  $(document.body).addClass("o_ui_blocked");
  blockAccessKeys();
  return tmp;*/
}

function unblockUI() {
  /*_.invoke(throbbers, "destroy");
  throbbers = [];
  $(document.body).removeClass("o_ui_blocked");
  unblockAccessKeys();
  return $.unblockUI.apply($, arguments);*/
}

/**
 * Redirect to url by replacing window.location
 * If wait is true, sleep 1s and wait for the server i.e. after a restart.
 */
function redirect(url, wait) {
  // Dont display a dialog if some xmlhttprequest are in progress
  //disableCrashManager();

  var load = function () {
    var old = '' + window.location;
    var old_no_hash = old.split('#')[0];
    var url_no_hash = url.split('#')[0];
    location.assign(url);
    if (old_no_hash === url_no_hash) {
      location.reload(true);
    }
  };

  var wait_server = function () {
    Ajax.rpc('/web/webclient/version_info', {})
      .then(load)
      .guardedCatch(function () {
        setTimeout(wait_server, 250);
      });
  };

  if (wait) {
    setTimeout(wait_server, 1000);
  } else {
    load();
  }
}

export var Http = { redirect };
