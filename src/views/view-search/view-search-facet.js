import { html, css, LitElement } from "lit-element";
import { nothing } from "lit-html";
import { unsafeHTML } from "lit-html/directives/unsafe-html.js";

class ViewSearchFacet extends LitElement {
  static get styles() {
    return css`
      :host {
        border: 1px solid #8f8f8f;
        background: #e2e2e0;
        color: #8f8f8f;
        -ms-flex: 0 0 auto;
        -moz-flex: 0 0 auto;
        -webkit-box-flex: 0;
        -webkit-flex: 0 0 auto;
        flex: 0 0 auto;
        max-width: max-content;
        display: -ms-flexbox;
        display: -moz-box;
        display: -webkit-box;
        display: -webkit-flex;
        display: flex;
        position: relative;
        margin: 1px 3px 0 0;
      }

      .searchview_facet_label {
        -ms-flex: 0 0 auto;
        -moz-flex: 0 0 auto;
        -webkit-box-flex: 0;
        -webkit-flex: 0 0 auto;
        flex: 0 0 auto;
        display: inline-block;
        max-width: 100%;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        vertical-align: top;
        padding: 0 3px;
        color: white;
        display: -ms-flexbox;
        display: -moz-box;
        display: -webkit-box;
        display: -webkit-flex;
        display: flex;
        -webkit-align-items: center;
        align-items: center;
      }

      .searchview_facet_label {
        background-color: #777777;
      }

      .fa-filter:before {
        content: "\f0b0";
      }

      .facet_values {
        padding-left: 5px;
      }

      .facet_remove {
        height: 18px;
        margin: auto;
        color: #777777;
      }
    `;
  }

  render() {
    return html`
      <!--图标-->
      ${this.facet
        ? html`
            ${this.facet && this.facet.icon
              ? html`
                  <span class="fa ${this.facet.icon} searchview_facet_label">${this.facet.category}</span>
                `
              : html`
                  <span class="searchview_facet_label">${this.facet.category}</span>
                `}

            <!--值-->
            <div class="facet_values">
              ${this.facet.values.map(
                (item, index) => html`
                  ${index > 0
                    ? html`
                        <span class="facet_values_sep">${unsafeHTML(this.facet.separator || " or " + item.label)}</span>
                      `
                    : html`
                        <span>${unsafeHTML(item.label || "")}</span>
                      `}
                `
              )}
            </div>
          `
        : nothing}

      <!--删除按钮-->
      <iron-icon icon="clear" class="fa fa-sm fa-remove facet_remove" @tap="${this.handleClick}"></iron-icon>
    `;
  }

  static get properties() {
    return {
      host: { type: Object }, // # View banding的 Facet
      facet: { type: Object },
      string: { type: String, notify: true },
      title: { type: String, notify: true }, // #显示的字符
      values: { type: Array } // # 存储值
    };
  }

  constructor() {
    super();
    //this.host = parent;
    //this._super(parent);
    //this.facet = facet_model;
    //this.model.on('change', this.model_changed, this);

    var self = this;
    var matches = this.shadowRoot.querySelectorAll(".facet_values");
    //var $e = this.$('.o_facet_values').last();
    var valnode = matches[matches.length - 1];

    /*    listeners: {
                'focus': 'handleFocus',
                    'blur': 'handleBlur',
                        'click': 'handleClick',
                            'keydown': 'handleKeydown',
    
        }*/
  }

  handleFocus() {
    this.dispatchEvent(new CustomEvent("focused", this));
  }

  handleBlur() {
    this.dispatchEvent(new CustomEvent("blurred", this));
  }

  handleClick(e) {
    this.dispatchEvent(new CustomEvent("removeFacet", { bubbles: true, composed: true }));
  }

  handleKeydown(e) {
    switch (e.keyCode) {
      case 8:
      case 46:
        this.destroy();
    }
  }

  model_changed() {
    this.$el.text(this.$el.text() + "*");
  }
}

customElements.define("view-search-facet", ViewSearchFacet);
