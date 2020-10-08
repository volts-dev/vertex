import { Field } from './field.js';

export class FieldBool extends Field {
  static get properties() {
    return {
      default_operator: { type: String },
    };
  }

  constructor() {
    super();
    this.default_operator = '=';
  }

  init() {
    this._super.apply(this, arguments);
    this.attrs.selection = [
      [true, _t('Yes')],
      [false, _t('No')],
    ];
  }

  get_groupby() {
    //
  }
}

customElements.define('field-bool', FieldBool);
