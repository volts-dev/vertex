import { html, css, LitElement } from 'lit-element';
import '@/elements/ve-dropdown/ve-dropdown.js';
import '@material/mwc-list/mwc-list-item';

class VeDropdownButton extends LitElement {
  static get styles() {
    return css`
      :host {
        display: inline-block;
        position: relative;
        outline: none;
        border: solid 1px transparent;
      }

      ::slotted([slot='trigger']) {
        height: 100%;
        list-style-type: none;
        white-space: nowrap;
      }

      ::slotted([slot='content']) {
        margin: 0 0;
        padding: 0;
        list-style-type: none;
        white-space: nowrap;
      }
    `;
  }

  render() {
    const itemRoles = this.innerRole === 'menu' ? 'menuitem' : 'option';

    return html`
      <div class="button" @click="${this.toggle}">
        <slot name="dropdown-trigger"></slot>
      </div>
      <ve-dropdown
        ?hidden=${!this.open}
        .anchor=${this.anchor}
        .open=${this.open}
        .quick=${this.quick}
        .corner=${this.corner}
        .x=${this.x}
        .y=${this.y}
        .absolute=${this.absolute}
        .fixed=${this.fixed}
        .fullwidth=${this.fullwidth}
        class="mdc-menu mdc-menu-surface"
      >
        <slot name="dropdown-content" class="dropdown-content"></slot>
      </ve-dropdown>
    `;
  }

  static get properties() {
    return {
      open: { type: Boolean },
      quick: { type: Boolean },
      absolute: { type: Boolean },
      fixed: { type: Boolean },
      multi: { type: Boolean },
      activatable: { type: Boolean },
      fullwidth: { type: Boolean },
      wrapFocus: { type: Boolean },
      corner: { type: String },
      innerRole: { type: String },
      anchor: { type: Object },

      opened: {
        type: Boolean,
        notify: true,
        value: false,
        observer: '_openedChanged1',
      },
      /**
       * Set to true to enable automatically closing the dropdown after an
       * item has been activated, even if the selection did not change.
       */
      closeOnActivate: { type: Boolean, value: false },
      disabled: Boolean,
      noAnimations: { type: Boolean, value: true },
      _dropdownContent: { type: Object },

      /**
       * By default, the dropdown will constrain scrolling on the page
       * to itself when opened.
       * Set to true in order to prevent scroll from being constrained
       * to the dropdown when it opens.
       * This property is a shortcut to set `scrollAction` to lock or refit.
       * Prefer directly setting the `scrollAction` property.
       */
      allowOutsideScroll: { type: Boolean, value: true },
    };
  }

  constructor() {
    super();
    this.listElement_ = null;
    this.anchor = null;
    this.open = false;
    this.quick = false;
    this.wrapFocus = false;
    this.innerRole = 'menu';
    this.corner = 'BOTTOM_START';
    this.absolute = false;
    this.multi = false;
    this.activatable = false;
    this.fixed = false;
    this.fullwidth = false;
    this.dynamicAlign = false;
    this.horizontalAlign = undefined;
    this.defaultFocus = 'LIST_ROOT';
  }

  firstUpdated() {
    super.firstUpdated();

    this.dropdown = this.shadowRoot.querySelector('ve-dropdown');
    this.button = this.shadowRoot.querySelector('.button');
    this.button = this.querySelector('[slot=dropdown-trigger]');
    this.anchor = this.button;

    this.addEventListener('mouseover', this.onMouseover);
    this.addEventListener('mouseleave', this.onMouseleave);
  }

  connectedCallback() {
    super.connectedCallback();
    // NOTE(cdata): Due to timing, a preselected value in a `IronSelectable`
    // child will cause an `iron-select` event to fire while the element is
    // still in a `DocumentFragment`. This has the effect of causing
    // handlers not to fire. So, we double check this value on attached:
    //var contentElement = this.contentElement;
    // if (contentElement && contentElement.selectedItem) {
    //   this._setSelectedItem(contentElement.selectedItem);
    // }
  }

  select(index) {
    const listElement = this.listElement;
    if (listElement) {
      listElement.select(index);
    }
  }

  /**
   * The content element that is contained by the dropdown menu, if any.
   */
  static get contentElement() {
    // Polymer 2.x returns slot.assignedNodes which can contain text nodes.
    var nodes = this.$.content.getDistributedNodes();
    for (var i = 0, l = nodes.length; i < l; i++) {
      if (nodes[i].nodeType === Node.ELEMENT_NODE) {
        return nodes[i];
      }
    }
  }

  toggle(e) {
    this.dropdown.open = !this.dropdown.open;
    this.open = this.dropdown.open;
  }

  /**
   * Hide the dropdown content.
   */
  close() {
    this.open = false;
  }

  show() {
    this.open = true;
  }

  onMouseover() {
    this.show();
  }

  onMouseleave() {
    if (this.open) {
      this.close();
    }
  }

  _openedChanged1(newValue, oldValue) {
    var openState = this.opened ? 'true' : 'false';
    var e = this.contentElement;
    if (e) {
      // e.setAttribute('aria-expanded', openState);
    }
  }
}

customElements.define('ve-dropdown-button', VeDropdownButton);
