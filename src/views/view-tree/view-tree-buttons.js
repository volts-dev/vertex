import { html, css, LitElement } from "lit-element";
//import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import '@polymer/iron-flex-layout/iron-flex-layout.js';

class ViewTreeButtons extends LitElement {
    static get styles() {
        return css`
        paper-button{height:28px;}
        `;
    }

    render() {
        return html`
        <div class="o_list_buttons horizontal layout center">
            <!--dom-if if="[[action_buttons]]" t-if="!widget.no_leaf and widget.options.action_buttons !== false and widget.options.addable and widget.is_action_enabled('create')">
                <template></template>
            </dom-if!-->
            <paper-button type="button" class="list_button_add btn btn-sm btn-primary">New<t t-esc="widget.options.addable"/></paper-button>
        </div>
    `;
    }
    static get properties() {
        return {
            action_buttons: { type: Boolean },
        };
    }
}

customElements.define('view-tree-buttons', ViewTreeButtons);
