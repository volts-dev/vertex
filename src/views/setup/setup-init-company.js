import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import { microTask } from '@polymer/polymer/lib/utils/async.js';

/**
 * `view-tree`
 *
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class SetupInitCompany extends PolymerElement {
  static get template() {
    return html`
      <style></style>

      <!--iron-ajax id="req"  method="post" handle-as="json" on-response="_handleResponse"></iron-ajax -->
      <div id="setup-init-company">初始化公司数据</div>
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
    var self = this;
    microTask.run(function () {
      self.dispatchEvent(new CustomEvent('tap-next'));
    });
  }
}

customElements.define('setup-init-company', SetupInitCompany);
