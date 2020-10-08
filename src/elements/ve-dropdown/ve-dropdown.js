import { css } from 'lit-element';
import { MenuSurface } from '@material/mwc-menu/mwc-menu-surface.js';

class VeDropdown extends MenuSurface {
  static get styles() {
    return [super.styles, css``];
  }

  static get properties() {
    return {};
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
    this.x = null;
    this.y = null;
    this.absolute = false;
    this.multi = false;
    this.activatable = false;
    this.fixed = false;
    this.forceGroupSelection = false;
    this.fullwidth = false;
    this.defaultFocus = 'LIST_ROOT';

    this.addEventListener('keydown', this.onKeydown.bind(this));
    this.addEventListener('opened', this.onOpened.bind(this));
    this.addEventListener('closed', this.onClosed.bind(this));
  }

  onKeydown(evt) {
    if (this.mdcFoundation) {
      this.mdcFoundation.handleKeydown(evt);
    }
  }

  onOpened() {
    this.open = true;
    if (this.mdcFoundation) {
      //this.mdcFoundation.handleMenuSurfaceOpened();
    }
  }
  onClosed() {
    this.open = false;
  }
}

customElements.define('ve-dropdown', VeDropdown);
