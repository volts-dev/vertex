import { html, css, LitElement } from 'lit-element';
import '@polymer/iron-flex-layout/iron-flex-layout.js';
import '@polymer/paper-item/paper-item.js';
import '@polymer/paper-listbox/paper-listbox.js';
import '@polymer/paper-menu-button/paper-menu-button.js';
import '../../element/vs-drop-menu/vs-drop-menu.js';

class ViewTreeSidebar extends LitElement {
    static get styles() {
        return css`
        iron-icon {
            margin: 0;
            cursor: pointer;
            padding: 0;
            margin-left: -3px;
        }
        
        vs-drop-menu {
            --background-color: white;
            --color: #222222;
        }`;
    }

    render() {
        return html`
            <div class="horizontal layout center">
                <!--test!-->
                <ve-dropdown-button horizontal-align="undefined" vertical-align="undefined" dynamic-align=false>
                        <paper-button slot="dropdown-trigger">
                            <span>Test</span><iron-icon icon="arrow-drop-down"></iron-icon>
                        </paper-button>
                        <paper-listbox slot="dropdown-content" class="dropdown-content">
                            <paper-item>
                                <a class="menu package" is="app-link" href$="[[_menuLink(submenu)]]" active$="[[_isEqual(package,submenu.name)]]" tabindex="[[index]]">
                                    <div class="layout horizontal center">
                                        <span class="title flex">Test</span>
                                    </div>
                                </a>
                            </paper-item>
                        </paper-listbox>
                    </ve-dropdown-button>
                <dom-repeat items="[[toolbar]]" as="tool">
                    <vs-drop-menu>
                        <paper-menu-button class="dropdown-trigger">[[_fmtToolName(tool.name)]]</paper-menu-button>
                        <paper-listbox class="dropdown-content">
                            <dom-repeat items="[[tool.value]]" as="action">
                                <paper-item>
                                    <a class="menu package" is="app-link" href$="[[_menuLink(submenu)]]" active$="[[_isEqual(package,submenu.name)]]" tabindex="[[index]]">
                                        <div class="layout horizontal center">
                                            <span class="title flex">[[action.display_name]]</span>
                                        </div>
                                    </a>
                                </paper-item>
                            </dom-repeat>
                        </paper-listbox>
                    </vs-drop-menu>
                </dom-repeat>
        </div >
        `;
    }

    static get properties() {
        return {
            toolbar: { type: Array, notify: false },
            action: { type: Array, notify: false },
            print: { type: Array, notify: false },
            relate: { type: Array, notify: false },
        };
    }


    _fmtToolName(name) {
        if (name == "action") {
            return "Action"
        }
        if (name == "print") {
            return "Print"
        }

        if (name == "relate") {
            return "Relate"
        }
    }
}

customElements.define('view-tree-sidebar', ViewTreeSidebar);
