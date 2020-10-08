import { html } from 'lit-element';
import { FormElement } from './../form-element';

class FormButton extends FormElement {
  render() {
    return html`
      <style>
        :host {
          position: relative;
          display: block;
        }
      </style>

        <button .toggles="${this.toggles}" .active="${this.active}}>
          <iron-icon icon="check"></iron-icon>
          ${this.name}
          <slot></slot>
        </button>
    `;
  }

  static get properties() {
    return {
      active: { type: Boolean },
      toggles: { type: Boolean },
    };
  }
}

customElements.define('form-button', FormButton);
