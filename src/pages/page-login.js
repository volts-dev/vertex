import { LitElement, html, css } from 'lit-element';

import Ajax from '@/core/ajax.js';

class PageLogin extends LitElement {
  static get styles() {
    return css`
      #login-panel {
        margin-top: 5rem !important;
        max-width: 300px;
        max-height: 350px;
        min-height: 200px;
        margin-left: -240px;
        margin: auto;
      }

      .border-bottom {
        border-bottom: 1px solid #dee2e6 !important;
      }
      .form-group {
        margin-bottom: 1rem;
      }
      .form-control {
        box-sizing: border-box;
        display: block;
        width: 100%;
        height: calc(2.0625rem + 2px);
        padding: 0.375rem 0.75rem;
        font-size: 0.875rem;
        line-height: 1.5;
        color: #495057;
        background-color: #ffffff;
        background-clip: padding-box;
        border: 1px solid #ced4da;
        border-radius: 0.25rem;
        transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
        background-color: rgb(232, 240, 254) !important;
      }

      @media (max-height: 500px) {
        #LoginPanel {
          position: absolute;
          left: 50%;
          top: 10%;
          width: 480px;
          min-height: 200px;
          max-height: 300px;
          margin-left: -240px;
          margin-top: 0;
        }
      }

      #login-button {
        padding-left: 15px;
        padding-right: 15px;
        background: #00b4f0;
        color: #ffffff;
      }
      paper-header-panel {
        background: #fafafa;
      }

      .content {
        display: block;
        padding: 50px 0;
      }

      .packages,
      #guides-container {
        max-width: 1100px;
        margin: 0 auto;
      }

      cart-icon {
        margin-left: 8px;
      }

      @media (max-width: 1132px) {
        .packages,
        #guides-container {
          max-width: 880px;
        }
      }

      @media (max-width: 912px) {
        .packages,
        #guides-container {
          max-width: 660px;
        }
      }

      @media (max-width: 692px) {
        .packages,
        #guides-container {
          max-width: 440px;
        }
      }

      a.package {
        width: 220px;
        margin: 8px;
      }

      @media (max-width: 489px) {
        a.package {
          width: calc(100% - 16px);
        }
      }

      .search.active {
        left: 15px;
      }

      @media (max-width: 639px) {
        paper-toolbar #topBar {
          margin-top: 4px;
          padding: 0 16px;
        }
      }

      guide-card,
      #coming-soon {
        width: 456px;
        cursor: pointer;
        margin: 8px;
      }

      @media (max-width: 489px) {
        guide-card,
        #coming-soon {
          width: 100%;
        }
      }

      #guides-container h3 {
        margin: 36px 8px 8px;
      }

      #coming-soon {
        cursor: normal;
        line-height: 76px;
        text-align: center;
        font-size: 16px;
        color: rgba(0, 0, 0, var(--dark-primary-opacity));
        border: 1px dashed #e5e5e5;
      }
    `;
  }

  render() {
    return html`
      <div id="login-panel" class="fit">
        <div class="text-center border-bottom">
          <img src="/app/binary/company_logo" />
        </div>

        <div id="login-form" class="login-form" vertical layout>
          <div style="display:none;">
            <input
              type="hidden"
              id="_authentication_token"
              name="_authentication_token"
              value="154501175410026827027342550694003663033"
            />
            <input
              id="back_url"
              type="hidden"
              name="back_url"
              value="${this.backUrl}"
            />
          </div>
          <div class="form-group" vertical layout>
            <label for="username">Account</label>
            <input
              id="username"
              type="text"
              name="username"
              class="form-control"
              placeholder="name@example.com"
              required
            />
            <div class="text-error"></div>
          </div>
          <div class="form-group" vertical layout>
            <label for="password">Password</label>
            <input
              id="password"
              type="password"
              class="form-control"
              name="password"
              required
            />
            <div id="login_status" class="text-status">${this.status}</div>
          </div>
          <div class="form-group">
            <input type="checkbox" name="rememberme" />
          </div>
          <br />
          <div class="form-group">
            <button
              id="login-button"
              raised
              @click="${this.onSubmit}"
              type="submit"
              name="button"
              class="btn btn-primary"
            >
              Sign in
            </button>
            <a
              class="forgot-your-password"
              href="/auth/forgot?continue_to=https%3A%2F%2Foctopart.com%2F"
              >Forgot your password?</a
            >
          </div>
        </div>
      </div>
    `;
  }

  static get properties() {
    return {
      username: { type: String },
      password: { type: String },
      status: { type: String, notify: true },
      backUrl: { type: String, notify: true },
    };
  }

  firstUpdated() {
    super.firstUpdated();
    this.username = this.shadowRoot.querySelector('#username');
    this.password = this.shadowRoot.querySelector('#password');

    var path = window.location.search.slice(1);
    var idx = path.search('([^back_url=].*)');
    if (idx != -1) {
      this.backUrl = path.substr(path.search('([^back_url=].*)')); // # 获取Url路径
    } else {
      this.backUrl = '';
    }
  }

  onSubmit(event) {
    var self = this;
    Ajax.rpc('/app/registry/login', {
      username: this.username.value,
      password: this.password.value,
      back_url: this.backUrl.value, //this.shadowRoot.querySelector('#back_url').value,
    }).then(resp => {
      if (resp.status == 'ERROR') {
        self.status = resp.message;
      } else if (resp.status == 'SUCCESS') {
        window.location.href = '.' + resp.message;
      } else {
        self.status = 'unknow error!';
      }
    });
  }
}

customElements.define('page-login', PageLogin);
