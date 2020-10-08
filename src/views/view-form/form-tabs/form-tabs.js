import { LitElement, html, css } from "lit-element";
import "@polymer/iron-pages/iron-pages";
import "@material/mwc-tab/mwc-tab.js";
import "@material/mwc-tab-bar/mwc-tab-bar.js";

class FormTabs extends LitElement {
  static get styles() {
    return css`
      :host {
        position: relative;
        display: block;
      }

      mwc-tab {
        webkit-flex: inherit1;
        flex: inherit;
        -webkit-flex-basis: 0.000000001px;
        flex-basis: initial;
        overflow: auto;
        margin-bottom: -1px;
      }

      ::slotted(mwc-tab) :hover {
        color: #555;
        background-color: #fff;
        border: 1px solid #ddd;
        border-bottom-color: transparent;
        cursor: default;
      }

      ::slotted(mwc-tab) {
        /* border-radius: 4px 4px 0 0;*/
        border: 1px solid transparent;
        padding: 0 20px;
      }

      ::slotted(#tabsContent)  {
        border-bottom: 1px solid #ddd;
        height: 98%;
      }

      ::slotted(#selectionBar) {
        height: 2px;
        bottom: -1px;
        background-color: #fff;
      }

      mwc-tab[active] {
        border: 1px solid #ddd;
        border-bottom-color: transparent;
      }
    `;
  }

  render() {
    return html`
      <mwc-tab-bar id="tabs" activeIndex="${this.selected}"> </mwc-tab-bar>

      <iron-pages id="pages" selected="${this.selected}">
        <slot></slot>
      </iron-pages>
    `;
  }

  static get properties() {
    return {
      selected: {
        type: Number
      }
    };
  }

  constructor() {
    super();
    this.selected = 0;
  }

  firstUpdated() {
    var tabs = this.shadowRoot.querySelector("#tabs");
    tabs.addEventListener("click", function(e) {
      this.selected =tabs.activeIndex;
    }.bind(this));

    //var pages = this.shadowRoot.querySelectorAll("page");
    let slots = this.shadowRoot.querySelector("slot");
    let pages = slots.assignedNodes({ flatten: true });

    // 生成标签和内容
    pages.forEach(page => {
      if (page.getAttribute) {
        var title = page.getAttribute("string");
        if (title && title != "") {
          var tab = document.createElement("mwc-tab");
          tab.setAttribute("label", title); // 添加控件属性"label"值
          tabs.appendChild(tab);
        }
      }
    });

    // 监听事件 获得新数据
    /*
        this.$.pages.addEventListener('iron-select', function(e) {
            if (e.srcElement == e.currentTarget) {
              //  e.stopImmediatePropagation();
            }
        });
        */
  }
}

customElements.define("form-tabs", FormTabs);
