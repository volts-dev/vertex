import { html } from 'lit-element';

class ViewSearchInputFilterGroup extends SearchInputMixin {
  render() {
    return html`
    ${filters.map(filter =>
      html`${!filter.visible || filter.visible()
        ? html` <li title="getTitle(filter)" data-index="{{index}}">
                      <a href="#">${this.getContent(filter)}</a>
                </li>`
        : html``}
      `)}
    `;
  }

  static get properties() {
    return {
      prop1: {
        type: String,
        value: 'view-tree',
      },
    };
  }

  // # 返回Filter菜单标题
  getTitle(filter) {

  }

  // # 返回Filter菜单内容
  getContent(filter) {

  }
}

customElements.define('view-search-input-filter-group', ViewSearchInputFilterGroup);