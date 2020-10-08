/**
 * Only responsability of this component is to display the user interface, and
 * react to user changes.
 *
 * @class Renderer
 */

import { Element } from '@/mixins/Element';

// 渲染原型
export class AbstractRenderer extends Element {
  static get properties() {
    return {
      state: String,
    };
  }

  constructor(parent, state, params) {
    super();
    //init: function (parent, state, params) {
    // this._super(parent);
    this.state = state;
  }
}

customElements.define('v-abstract-renderer', AbstractRenderer);
