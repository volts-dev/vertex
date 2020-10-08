import { html, css, LitElement } from 'lit-element';

class ViewSearchInput extends LitElement {
  static get styles() {
    return css`
      :host {
        width: 100px;
        -ms-flex: 1 0 auto;
        -moz-flex: 1 0 auto;
        -webkit-box-flex: 1;
        -webkit-flex: 1 0 auto;
        flex: 1 0 auto;
      }
      .searchview_input {
        width: 100%;
      }
      /* remove select border */
      input:focus {
        -webkit-box-shadow: none;
        box-shadow: none;
        outline: none;
      }

      input {
        border: none;
        width: 100%;
      }
    `;
  }

  render() {
    return html`
      <input
        type="text"
        class="searchview_input"
        @keydown="${this.keydownHandler}"
        @keyup="${this.keyupHandler}"
      />
      <input id="input" placeholder="Search..." @change="${this.onChange}" />
    `;
  }

  static get properties() {
    return {
      host: Object, // # search view

      value: {
        type: String,
        value: '',
        notify: true,
      },
    };
  }

  firstUpdated() {
    super.firstUpdated();
    this.input = this.shadowRoot.querySelector('#input');
  }

  onChange() {
    /**bug 当改变时必须变换焦点才会触发 */
    //this.value = this.input.value;
  }

  keydownHandler(e) {}

  keyupHandler(e) {
    // TODO　考虑其他位置
    // this.input.value= this.input.value+ String.fromCharCode(e.which);
    this.value = this.input.value;
  }
}

customElements.define('view-search-input', ViewSearchInput);
