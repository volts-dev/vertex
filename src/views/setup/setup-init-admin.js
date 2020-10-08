import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import { microTask } from '@polymer/polymer/lib/utils/async.js';
import '@polymer/paper-input/paper-input.js';
import '@polymer/paper-button/paper-button.js';
import '@polymer/iron-form/iron-form.js';

/**
 * `view-tree`
 *
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class SetupInitAdmin extends PolymerElement {
  static get template() {
    return html`
      <style>
        #view-creator-init-admin {
          height: 800px;
          width: 1024px;
        }
      </style>
      <div id="view-creator-init-admin">
        <json-rpc
          id="rpc"
          url="/app/registry/change_password"
          method="call"
          on-result="handleResult"
          last-result="{{rpcResult}}"
          method="listEntries"
          params="[]"
        ></json-rpc>
        <div id="form" class="form" vertical layout>
          <div style="">更改超级用户默认密码</div>

          <div class="form-group" vertical layout>
            <paper-input
              type="password"
              id="password"
              label="New Password"
              class="form-control"
              required
            ></paper-input>
            <div class="text-error"></div>
          </div>

          <div class="form-group" vertical layout>
            <paper-input
              type="password"
              id="password2"
              label="Confirm Password"
              class="form-control"
              required
            ></paper-input>
            <div id="status" class="text-status"></div>
          </div>
          <div class="form-group">
            <paper-button
              type="submit"
              name="button"
              class="btn btn-primary"
              on-tap="confirmBtn"
              >Confirm</paper-button
            >
          </div>
        </div>
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

  connectedCallback() {}

  handleResult(e, result) {
    if (result.status == 'ERROR') {
      var status = this.$.status; // document.getElementById('login-status');
      status.innerHTML = result.message;
    } else {
      var self = this;
      microTask.run(function () {
        self.dispatchEvent(new CustomEvent('tap-next')); // 跳转到下一页
      });
    }
  }

  confirmBtn(e) {
    //var target = Polymer.dom(event).localTarget;
    //var form = document.getElementById('login-form');
    //form.submit();
    var rpc = this.$.rpc;
    var data = {};
    data.password = this.$.password.value;
    data.password2 = this.$.password2.value;
    rpc.params = data;
    rpc.generateRequest();
  }
}

customElements.define('setup-init-admin', SetupInitAdmin);
