import { Element } from "../../core/element";
import { html, css } from "lit-element";
import "@polymer/app-layout/app-header/app-header.js";
import "@polymer/app-layout/app-toolbar/app-toolbar.js";

export class MenuNav extends Element {
  static get styles() {
    return css``;
  }

  render() {
    return html` <!-- toolbar of home  -->
      <app-header id="header" effects="waterfall" ?fixed="${this.smallScreen}" ?condenses="${!this.smallScreen}" ?reveals="${!this.smallScreen}">
        <app-toolbar id="mainToolbar" class="main_too_bar">
          ${this.page == "home"
            ? html`
                <!-- toolbar of home  -->
                <div class="navItem leftItem">
                  <a href="${this.lastUrl}" tabindex="-1">
                    <paper-icon-button icon="arrow-back" alt="Back to the home"></paper-icon-button>
                  </a>
                  <paper-icon-button icon="menu" drawer-toggle alt="Toogle navigation menu"></paper-icon-button>
                </div>
              `
            : html`
                <!-- module of toolbar -->
                <div class="navItem leftItem">
                  <a href="/app#home" @click="${this.recordUrl}" tabindex="-1">
                    <paper-icon-button icon="polymer" alt="Back to the home"></paper-icon-button>
                  </a>
                  <paper-icon-button icon="menu" drawer-toggle alt="Toogle navigation menu"></paper-icon-button>
                </div>

                <!-- name of app -->
                <div class="menu_brand navItem leftItem">${this.title}</div>

                <!-- menu of module -->
                <div class="menu_sections horizontal layout flex">
                  ${this.menus.map(
                    (item) => html`
                      ${item.parent_id == this.parentMenu
                        ? html`
                            <ve-dropdown-button class="" horizontal-align="undefined" vertical-align="undefined" dynamic-align="false">
                              <paper-button slot="dropdown-trigger">${item.name}</paper-button>
                              <ul slot="dropdown-content">
                                ${this.menus.map(
                                  (menu, index) => html`
                                    ${menu.parent_id == item.id
                                      ? html`
                                          <paper-item>
                                            <a
                                              class="menu package"
                                              is="app-link"
                                              href="${this.menuLink(menu)}"
                                              ?active=${menu.name == "package"}
                                              tabindex="${index}"
                                              @click="${this.onMenuClicked}"
                                            >
                                              <div class="layout horizontal center">
                                                <package-symbol aria-hidden="true" package="${menu}"></package-symbol>
                                                <span class="title flex">${menu.name}</span>
                                              </div>
                                            </a>
                                          </paper-item>
                                        `
                                      : null}
                                  `
                                )}
                              </ul>
                            </ve-dropdown-button>
                          `
                        : null}
                    `
                  )}
                </div>
              `}

          <div main-title></div>

          <!-- menu of system -->
          <div class="menu_systray menu_sections navItem">
            <paper-icon-button icon="shopping-cart" aria-label="Shopping cart"></paper-icon-button>
            <ve-dropdown-button horizontal-align="undefined" vertical-align="undefined" dynamic-align="false">
              <paper-icon-button slot="dropdown-trigger" icon="more-vert" aria-label="More options"></paper-icon-button>
              <paper-listbox slot="dropdown-content" class="dropdown-content">
                <paper-item>
                  <a class="menu package" is="app-link" href="/app/registry/logout">
                    <div class="layout horizontal center">
                      <package-symbol aria-hidden="true" package="submenu"></package-symbol>
                      <span class="title flex">Logout</span>
                    </div>
                  </a>
                </paper-item>
              </paper-listbox>
            </ve-dropdown-button>
          </div>
        </app-toolbar>
      </app-header>`;
  }

  static get properties() {
    return {
      menus_data: { type: Array },
    };
  }
}
