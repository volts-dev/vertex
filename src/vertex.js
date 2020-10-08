import 'underscore';
import { html, css } from 'lit-element';

// lib
import '@/libs/promise_extension.js';
// core
import '@/core/pyeval.js';
import '@/core/strings.js';
import '@/core/utils.js';

// elements
import { Element } from '@/mixins/Element.js';
import router from '@/router';
import { ServicesListenerMixin } from '@/mixins/ServicesListenerMixin';

// pages
//import '@/pages/page-setup.js';
import '@/pages/page-login.js';
import '@/pages/page-main.js';
import './vertex-styles.js';
import './nav-styles.js';
//import '@/styles/variables.scss';

// App 作为前端服务程序
// the main web client
class VertexApp extends ServicesListenerMixin(Element) {
  /*static get routes() {
    return [
      {
        name: 'app',
        pattern: 'app',
        data: { title: 'Home' },
      },
      {
        name: 'login',
        pattern: 'login',
      },
      {
        name: 'setup',
        pattern: 'setup',
      },
    ];
  }
*/
  static get styles() {
    return css`
      .iron-pages {
        width: 100%;
        height: 100%;
      }
    `;
  }

  render() {
    return html`
      <div class="iron-pages">
        <!-- app -->
        ${router.route == 'app'
          ? html` <page-main name="app"></page-main> `
          : null}

        <!-- login -->
        ${router.route == 'login'
          ? html` <page-login name="login"></page-login> `
          : null}

        <!-- setup -->
        ${router.route == 'setup'
          ? html` <page-setup name="setup"></page-setup> `
          : null}
      </div>
    `;
  }
}

customElements.define('vertex-app', VertexApp);
