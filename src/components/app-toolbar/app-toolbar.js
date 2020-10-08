import { html, css } from 'lit-element';
import store from '@/store';
import router from '@/router';

import '@/elements/ve-dropdown-button/ve-dropdown-button';
//import '@/elements/ve-list-item';

export class AppToolbar extends store.Element {
  static get styles() {
    return css`
      .main_navbar {
        position: relative;
        height: var(--v-navbar-height);
        border-bottom: 1px solid var(--v-navbar-inverse-link-hover-bg);
        background-color: var(--v-brand-vertex);
        color: white;
      }

      a,
      button,
      .dropdownButton {
        float: left;
        height: var(--v-navbar-height);
        border: none;
        padding: 0 var(--v-horizontal-padding) - 4px 0
          var(--v-horizontal-padding);
        line-height: var(--v-navbar-height);
        background-color: transparent;
        text-align: center;
        color: inherit;

        font-size: 18px;
        user-select: none;
      }

      a:hover,
      button:hover,
      .dropdownButton:hover,
      a:focus,
      button:focus,
      .dropdownButton:focus {
        background-color: var(--v-navbar-inverse-link-hover-bg);
        color: inherit;
      }

      a:focus,
      button:focus,
      .dropdownButton:focus,
      a:active,
      button:active,
      .dropdownButton:focus:active {
        outline: none;
      }

      .app {
        cursor: pointer;
      }

      .menu_brand {
        display: block;
        float: left;
        margin-right: 35px;
        user-select: none;
        color: white;
        font-size: 22px;
        font-weight: 500;
        line-height: var(--v-navbar-height);
        cursor: pointer;
      }

      .menu_sections {
        display: block;
        margin: 0;
        padding: 0;
      }

      .menu_toggle {
        margin-right: 5px;
        width: 44px;
      }

      .menu_systray {
        float: right;
      }
    `;
  }

  render() {
    return html`
      <div class="main_navbar">
        <a
          class="menu_toggle"
          @click="${this.recordUrl}"
          tabindex="-1"
          alt="Back to the home"
        >
        </a>
        ${store.app.activeAppId
          ? html` <!-- name of app -->
              <div class="menu_brand navItem leftItem">
                ${store.app.activeApp.name}
              </div>
              <!-- menu sections of app -->
              <div class="menu_sections">
                ${store.app.activeApp.menus.map(
                  menu => html`<ve-dropdown-button>
                    <a slot="dropdown-trigger">${menu.name}</a>
                    <div slot="dropdown-content">
                      ${menu.subMenus.map(
                        sub => html`
                          <ve-list-item
                            ><a href="${this.menuLink(sub)}"
                              >${sub.name}</a
                            ></ve-list-item
                          >
                        `
                      )}
                    </div>
                  </ve-dropdown-button> `
                )}
              </div>`
          : null}

        <!-- menu of system -->
        <div class="horizontal layout flex">
          <div class="right-menu">
            ${store.app.device !== 'mobile'
              ? html`
                  <search id="header-search" class="right-menu-item" />

                  <error-log
                    class="errLog-container right-menu-item hover-effect"
                  />

                  <ve-screenfull
                    id="screenfull"
                    class="right-menu-item hover-effect"
                  ></ve-screenfull>

                  <el-tooltip
                    content="Global Size"
                    effect="dark"
                    placement="bottom"
                  >
                    <size-select
                      id="size-select"
                      class="right-menu-item hover-effect"
                    />
                  </el-tooltip>
                `
              : null}
          </div>
        </div>
        <!-- menu of user -->
        <ve-dropdown-button class="menu_systray">
          <a slot="dropdown-trigger">a</a>
        </ve-dropdown-button>
      </div>
    `;
  }

  static get properties() {
    return {
      backUrl: String,
    };
  }

  recordUrl() {
    if (store.app.activeAppId) {
      this.backUrl = router.Path();
      router.Push(`/app#home`);
    } else {
      router.Push(this.backUrl);
    }
  }

  // 组织Link的参数
  menuLink(item) {
    //model = '';
    var action = '';
    var url = '/app?menu=' + item.id;
    if (item.action_id) {
      url = url + '&action=' + item.action_id;
    }

    return url;
  }
}

customElements.define('app-toolbar', AppToolbar);
