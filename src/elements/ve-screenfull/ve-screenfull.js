import { html, css, LitElement } from 'lit-element';
import screenfull from 'screenfull';

export class Screenfull extends LitElement {
  static get styles() {
    return css`
      .screenfull-svg {
        display: inline-block;
        cursor: pointer;
        fill: #5a5e66;
        width: 20px;
        height: 20px;
        vertical-align: 10px;
      }
    `;
  }

  render() {
    return html`
      <div>
        <svg-icon
          :icon-class="isFullscreen?'exit-fullscreen':'fullscreen'"
          @click="${this.click}"
        ></svg-icon>
      </div>
    `;
  }

  static get properties() {
    return {
      isFullscreen: false,
    };
  }

  connectedCallback() {
    this.init();
  }

  disconnectedCallback() {
    this.destroy();
  }

  click() {
    if (!screenfull.enabled) {
      /*
      this.$message({
        message: 'you browser can not work',
        type: 'warning',
      });
      */
      return false;
    }
    screenfull.toggle();
  }

  change() {
    this.isFullscreen = screenfull.isFullscreen;
  }

  init() {
    if (screenfull.enabled) {
      screenfull.on('change', this.change);
    }
  }

  destroy() {
    if (screenfull.enabled) {
      screenfull.off('change', this.change);
    }
  }
}
