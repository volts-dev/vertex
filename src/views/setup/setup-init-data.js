import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import '@polymer/paper-button/paper-button.js';
import { microTask } from '@polymer/polymer/lib/utils/async.js';
import '../../element/json-rpc/json-rpc.js';

/**
 * `view-tree`
 *
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class SetupInitData extends PolymerElement {
  static get template() {
    return html`
      <style>
      </style>
      <json-rpc id="rpc" url="" method="call" on-result="_handleResponse" last-result="{{rpcResult}}"
      <div id="view-creator-init-data">
        初始化数据库
      </div>
    `;
  }
  static get properties() {
    return {
      prop1: {
        type: String,
        value: 'view-tree',
      },
    };
  }

  run() {
    this.$.rpc.url = '/app/registry/setup';
    this.$.rpc.generateRequest();
  }

  _handleResponse(e, result) {
    if (result.status == 'SUCCESS') {
      var self = this;
      // 跳转到下一页
      microTask.run(function () {
        self.dispatchEvent(new CustomEvent('tap-next'));
      });
    } else if (result.status == 'ERROR') {
      //var status =this.$.login_status;// document.getElementById('login-status');
      // status.innerHTML = result.message;
    } else {
      // 跳转到下一页
    }
  }
}

customElements.define('setup-init-data', SetupInitData);
