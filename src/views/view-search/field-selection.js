import { Field } from './field.js';

/**
 * `view-tree`
 *
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
// This implementation is a basic <select> field, but it may have to be
// altered to be more in line with the GTK client, which uses a combo box
// (~ jquery.autocomplete):
// * If an option was selected in the list, behave as currently
// * If something which is not in the list was entered (via the text input),
//   the default domain should become (`ilike` string_value) but **any
//   ``context`` or ``filter_domain`` becomes falsy, idem if ``@operator``
//   is specified. So at least get_domain needs to be quite a bit
//   overridden (if there's no @value and there is no filter_domain and
//   there is no @operator, return [[name, 'ilike', str_val]]

export class FieldSelection extends Field {
  static get properties() {
    return {
      default_operator: { type: String },
    };
  }

  constructor() {
    super();
    this.default_operator = 'ilike';
  }

  init() {
    this._super.apply(this, arguments);
    // prepend empty option if there is no empty option in the selection list
    this.prepend_empty = !_(this.attrs.selection).detect(function (item) {
      return !item[1];
    });
  }

  complete(needle) {
    var self = this;
    var results = _(this.attrs.selection)
      .chain()
      .filter(function (sel) {
        var value = sel[0],
          label = sel[1];
        if (value === undefined || !label) {
          return false;
        }
        return label.toLowerCase().indexOf(needle.toLowerCase()) !== -1;
      })
      .map(function (sel) {
        return {
          label: _.escape(sel[1]),
          indent: true,
          facet: facet_from(self, sel),
        };
      })
      .value();
    if (_.isEmpty(results)) {
      return $.when(null);
    }
    return $.when.call(
      null,
      [
        {
          label: _.escape(this.attrs.string),
        },
      ].concat(results)
    );
  }

  facet_for(value) {
    var match = _(this.attrs.selection).detect(function (sel) {
      return sel[0] === value;
    });
    if (!match) {
      return $.when(null);
    }
    return $.when(facet_from(this, match));
  }
}

customElements.define('field-selection', FieldSelection);
