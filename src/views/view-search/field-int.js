import { Field } from './field.js';

class NumberFieldMixin extends Field {
  complete(value) {
    var val = this.parse(value);
    if (isNaN(val)) {
      return $.when();
    }
    /* var label = _.str.sprintf(
             _t("Search %(field)s for: %(value)s"), {
                 field: '<em>' + _.escape(this.attrs.string) + '</em>',
                 value: '<strong>' + _.escape(value) + '</strong>'
             });
             */
    var label =
      'Search <em>' +
      this.attrs.string +
      '</em> for: <strong>' +
      value +
      '</strong>';

    return Promise.resolve([
      {
        label: label,
        facet: {
          category: this.attrs.string,
          field: this,
          values: [{ label: value, value: val }],
        },
      },
    ]);
  }

  get_groupby() {
    //
  }
}

export class FieldInt extends NumberFieldMixin {
  static get properties() {
    return {
      error_message: 'not a valid integer',
    };
  }

  parse(value) {
    try {
      return formats.parse_value(value, { widget: 'integer' });
    } catch (e) {
      return NaN;
    }
  }
}

export class FieldFloat extends NumberFieldMixin {
  static get properties() {
    return {
      error_message: 'not a valid number',
    };
  }

  parse(value) {
    try {
      return formats.parse_value(value, { widget: 'float' });
    } catch (e) {
      return NaN;
    }
  }
}

customElements.define('field-int', FieldInt);
customElements.define('field-float', FieldFloat);
