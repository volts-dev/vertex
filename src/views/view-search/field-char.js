import { Field } from './field.js';

export class FieldChar extends Field {
  static get properties() {
    return {
      default_operator: { type: String },
    };
  }

  constructor() {
    super();
    this.default_operator = 'ilike';
  }

  complete(value) {
    //if (isEmpty(value)) {
    if (!value) {
      return Promise.resolve(null);
    }

    //var label = _.str.sprintf(_.str.escapeHTML(
    //    _t("Search %(field)s for: %(value)s")), {
    //        field: '<em>' + _.escape(this.attrs.string) + '</em>',
    //        value: '<strong>' + _.escape(value) + '</strong>'
    //    });
    var label =
      'Search <em>' +
      this.attrs.string +
      '</em> for: <strong>' +
      value +
      '</strong>';

    /*
        return $.when([{
            label: label,
            facet: {
                category: this.attrs.string,
                field: this,
                values: [{ label: value, value: value }]
            }
        }]);
*/
    return Promise.resolve([
      {
        label: label,
        facet: {
          category: this.attrs.string,
          field: this,
          values: [{ label: value, value: value }],
        },
      },
    ]);
  }

  get_groupby() {
    //
  }
}

customElements.define('field-char', FieldChar);
