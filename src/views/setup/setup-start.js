import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import '@polymer/paper-button/paper-button.js';
import { microTask } from '@polymer/polymer/lib/utils/async.js';

/**
 * `view-tree`
 *
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class SetupStart extends PolymerElement {
  static get template() {
    return html`
      <style>
        #view-creator-start {
          height: 800px;
          width: 1024px;
        }
      </style>
      <div id="view-creator-start">
        <paper-button raised on-tap="startBtn">Start</paper-button>
        <paper-button raised on-tap="cancelBtn">Cancel</paper-button>
      </div>
    `;
  }

  startBtn(e) {
    var self = this;
    microTask.run(function () {
      self.dispatchEvent(new CustomEvent('tap-next'));
    });
  }

  cancelBtn() {}
}

customElements.define('setup-start', SetupStart);
