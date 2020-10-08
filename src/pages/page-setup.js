import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import '@polymer/iron-pages/iron-pages.js';
import '@polymer/paper-toolbar/paper-toolbar.js';
import '@polymer/iron-ajax/iron-ajax.js';

import '../view/setup/setup-start.js';
import '../view/setup/setup-init-data.js';
import '../view/setup/setup-init-company.js';
import '../view/setup/setup-init-admin.js';
import { PageMixin } from './page.js';

class PageSetup extends PageMixin(PolymerElement) {
  static get template() {
    return html`
      <style></style>
      <!---获取相关数据--->
      <!---赋值--->
      <div
        id="creatorPanel"
        class="creator-panel {{ {wide: wide} | tokenList }}"
        flex
        vertical
        layout?="{{!wide}}"
        slide-up?="{{parentElement.selected !== 'empty'}}"
        slide-down?="{{parentElement.selected === 'empty'}}"
      >
        <!---标题--->
        <iron-pages
          id="creatorViews"
          notap
          selected="{{selected}}"
          attr-for-selected="name"
          selectedItem="{{selectedItem}}"
          transitions="cross-fade cross-fade-delayed slide-up slide-down hero-transition"
        >
          <!---Pannel--->
          <setup-start
            id="start"
            name="start"
            on-tap-next="tapStart"
          ></setup-start>
          <!---初始化数据库--->
          <setup-init-data
            id="initData"
            name="initData"
            on-tap-next="tapData"
          ></setup-init-data>
          <!---创建企业信息--->
          <setup-init-company
            id="initCompany"
            name="initCompany"
            on-tap-next="tapCompany"
          ></setup-init-company>
          <!---修改超级密码--->
          <setup-init-admin
            id="initAdmin"
            name="initAdmin"
            on-tap-next="tapAdmin"
          ></setup-init-admin>
        </iron-pages>
      </div>
    `;
  }

  static get properties() {
    return {
      selected: {
        type: String,
        notify: true,
        value: 'start',
      },
      selectedItem: {
        type: String,
      },
    };
  }

  static get observers() {
    return [
      'updateURL(q,package,tag,view,elements)',
      'updateMeta(packageInfo)',
      'scrollToTop(package)',
    ];
  }

  ready() {
    super.ready();
    // this.view = this._forceCards ? 'cards' : 'table';
  }

  attached() {
    //this.updateMeta();
  }

  // 执行下一步
  tapStart() {
    this.selected = 'initData';
    this.$.initData.run();
  }

  tapData() {
    this.selected = 'initCompany';
    this.$.initCompany.run(); // 执行CompanyView 里的Run()跳到Admin
  }

  tapCompany() {
    this.selected = 'initAdmin';
  }

  tapAdmin() {
    //  this.selected = "initFinish";
    window.location.href = '/';
  }

  tapfinish() {
    window.location.href = '/';
  }
}

customElements.define('page-setup', PageSetup);
