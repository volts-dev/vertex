import { html, css, LitElement } from 'lit-element';
import store from '@/store';

class AppIcon extends store.Element {
  static get styles() {
    return css`
      :host {
        position: relative;
        width: 80%;
        max-width: 70px;
        overflow: hidden;
        margin: auto;
        padding: 10px 10px;
        border-radius: 4%;
        transition: all 0.3s ease 0s;
        box-shadow: 0 8px 0 -10px black;
        background-repeat: no-repeat;
        background-position: center;
        background-size: cover;
      }

      .deep-2 {
        max-width: 100%;
        height: auto;
      }

      :host(.active) {
        transform: translate(0, 0);
        opacity: 1;
      }

      :host(.active) h2 {
        transform: translate(0, 0);
        opacity: 1;
      }

      :host(.active) h3 {
        transform: translate(0, 0);
        opacity: 1;
      }

      :host(.active) hr {
        width: 100%;
        background: #606060;
        opacity: 1;
      }

      :host(.active) .version {
        transform: translate(0, 0);
        opacity: 0.8;
      }

      :host(.active) .tagline {
        transform: translate(0, 0);
        opacity: 1;
      }

      :host(:hover) #content {
        border-color: rgba(255, 255, 255, 0.08);
        background-color: rgba(255, 255, 255, 0.05);
      }

      .icon {
        border-radius: 4px;
        width: 75px;
        height: 75px;
        background-color: white;
      }
      .icon:hover {
        box-shadow: 0 8px 15px -10px black;
        transform: translateY(-1px);
      }

      .caption {
        margin: 0;
        font-size: 12px;
        white-space: nowrap;
        text-align: center;
        display: block;
        max-width: 100%;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        margin: 4px 0;
        color: white;
        text-shadow: 0 1px 1px rgba(0, 0, 0, 0.8);
      }

      .deep-2 {
        box-shadow: 0 6px 10px 0 rgba(0, 0, 0, 0.1),
          0 2px 2px 0 rgba(0, 0, 0, 0.05);
      }

      a {
        color: #875a7b;
        text-decoration: none;
      }
      #content {
        display: block;
        position: relative;
        width: 100%;
        box-sizing: border-box;
        overflow: hidden;
        cursor: pointer;
        transition: box-shadow 200ms;
        transition-timing-function: var(--material-curve);
        color: #606060;
        padding: 16px;
        border-radius: 6%;
        border: 1px solid transparent;
      }

      h2 {
        font-weight: 400;
        font-size: 48px;
        margin: 20px 0;
        transform: translate(-50px, 0);
        opacity: 0;
        transition: all 500ms 320ms;
        transition-timing-function: var(--material-curve-320);
      }

      h3 {
        transition: all 600ms 320ms;
        transition-timing-function: var(--material-curve-320);
      }

      hr {
        border: 0;
        background: #fff;
        width: 0;
        height: 1px;
        opacity: 0.2;
        margin: 0;
        transition: width 400ms 320ms;
        transition-timing-function: var(--material-curve-320);
      }

      .version {
        position: absolute;
        top: 16px;
        right: 16px;
        font-size: 13px;
        transition: all 400ms 320ms;
        transition-timing-function: var(--material-curve-320);
        transform: translate(50px, 0);
        opacity: 0;
      }

      .title {
        display: table-caption;
        margin: 0 0 15px;
        height: 64px;
        font-size: 20px;
        font-weight: 500;
        line-height: 28px;
      }

      .title[extended] {
        display: block;
      }

      .tagline {
        transition: all 700ms 320ms;
        transition-timing-function: var(--material-curve-320);
        transform: translate(-50px, 0);
        opacity: 0;
        font-size: 13px;
        margin: 10px 0 0 0;
        line-height: 20px;
        height: 40px;
      }
    `;
  }

  render() {
    return html`
      <a .href="${this.menuLink()}" @click="${this.onAppClicked}">
        <div class="vertical layout">
          <img
            class=" deep-2"
            src="//download.odoocdn.com/icons/repair/static/description/icon.png"
            alt="Website"
          />
        </div>
        <div class="caption" .extended="${this._extendedTitle(this.item.name)}">
          ${this.item.name}
        </div>
      </a>
    `;
  }

  static get properties() {
    return {
      item: { type: Object },
      href: String,
      skipNav: Boolean,
      name: { type: String },
      text: { type: String },
      module: { type: Object },
    };
  }

  connectedCallback() {
    super.connectedCallback();
    var tiles = this.parentNode.querySelectorAll('module-tile');
    var index = Array.prototype.indexOf.call(tiles, this);
    setTimeout(
      function () {
        this.classList.add('active');
      }.bind(this),
      (index + 1) * 50
    );
  }

  onAppClicked(e) {
    // allow for ctrl+click to open in new window
    if (e.ctrlKey || e.metaKey) {
      return true;
    } else {
      // 防止重载网页
      // e.preventDefault();
      if (!this.skipNav) {
        // 更新如有路径
        // this.fire('nav', { url: this.href });
        this.dispatchEvent(new CustomEvent('nav', { url: this.href }));
      }
    }
  }

  _extendedTitle(p) {
    return p.length > 50;
  }

  // 组织Link的参数
  menuLink() {
    //model = '';
    var action = '';
    var url = '/app?menu=' + this.item.id;
    if (this.item.action_id) {
      url = url + '&action=' + this.item.action_id;
    }

    return url;
  }
}

customElements.define('app-icon', AppIcon);
