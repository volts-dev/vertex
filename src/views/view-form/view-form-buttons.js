import { LitElement, html, css } from "lit-element";

class ViewFormButtons extends LitElement {
  static get properties() {
    return {
      action_buttons: { type: Boolean, notify: false },
      mode: String
    };
  }

  static get styles() {
    return css`
      :host {
        display: flex;
      }
      .form_buttons {
        display: flex;
      }
    `;
  }

  render() {
    return html`
      <div class="form_buttons" t-name="FormView.buttons">
        ${this.action_buttons
          ? html`
              <div class="form_buttons_view">
                ${this.mode == "readonly"
                  ? html`
                      <button
                        t-if="widget.is_action_enabled('edit')"
                        type="button"
                        .action="${"edit"}"
                        @click="${this.onClick}"
                        class="form_button_edit btn btn-default btn-sm"
                        accesskey="E"
                      >
                        Edit
                      </button>
                      <button
                        t-if="widget.is_action_enabled('create')"
                        type="button"
                        .action="${"create"}"
                        @click="${this.onClick}"
                        class="form_button_create btn btn-default btn-sm"
                        accesskey="C"
                      >
                        Create
                      </button>
                    `
                  : html`
                      <button type="button" .action="${"save"}" @click="${this.onClick}" class="form_button_save btn btn-primary btn-sm" accesskey="S">
                        Save
                      </button>
                      <button type="button" .action="${"discard"}" @click="${this.onClick}"" class="form_button_cancel btn btn-sm btn-default" accesskey="D">
                        Discard
                      </button>
                    `}
              </div>
            `
          : null}
      </div>
    `;
  }

  constructor() {
    super();
    this.mode = "readonly";
  }

  onClick(event) {
    var el = event.srcElement;
    if (el.action) {
      this.dispatchEvent(new CustomEvent("clickButton", { detail: el.action }));
    }
  }
}

customElements.define("view-form-buttons", ViewFormButtons);
