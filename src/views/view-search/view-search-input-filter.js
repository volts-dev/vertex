import { html, LitElement } from 'lit-element';


window.vectors = window.vectors || {};
vectors.view.search = vectors.view - search || {};
vectors.view.search.InputElement = {
    /**
* @constructs instance.web.search.Input
* @extends instance.web.Widget
*
* @param parent
*/
    init(parent) {
        this._super(parent);
        this.searchview = parent;
        this.load_attrs({});
    },
    /**
  * Fetch auto-completion values for the widget.
  *
  * The completion values should be an array of objects with keys category,
  * label, value prefixed with an object with keys type=section and label
  *
  * @param {String} value value to complete
  * @returns {jQuery.Deferred<null|Array>}
  */
    complete(value) {
        // return $.when(null);
    },
    /**
     * Returns a Facet instance for the provided defaults if they apply to
     * this widget, or null if they don't.
     *
     * This default implementation will try calling
     * :js:func:`instance.web.search.Input#facet_for` if the widget's name
     * matches the input key
     *
     * @param {Object} defaults
     * @returns {jQuery.Deferred<null|Object>}
     */
    facet_for_defaults(defaults) {
        if (!this.attrs ||
            !(this.attrs.name in defaults && defaults[this.attrs.name])) {
            // return $.when(null);
        }
        return this.facet_for(defaults[this.attrs.name]);
    },
    get_context() {
        throw new Error(
            "get_context not implemented for widget " + this.attrs.type);
    },
    get_groupby() {
        throw new Error(
            "get_groupby not implemented for widget " + this.attrs.type);
    },
    get_domain() {
        throw new Error(
            "get_domain not implemented for widget " + this.attrs.type);
    },
    load_attrs(attrs) {
        if (!_.isObject(attrs.modifiers)) {
            attrs.modifiers = attrs.modifiers ? JSON.parse(attrs.modifiers) : {};
        }
        this.attrs = attrs;
    },
    /**
     * Returns whether the input is "visible". The default behavior is to
     * query the ``modifiers.invisible`` flag on the input's description or
     * view node.
     *
     * @returns {Boolean}
     */
    visible() {
        return !this.attrs.modifiers.invisible;
    },
};

class ViewSearchInputField extends LitElement {
    static get styles() {
        return css`

        .o_searchview_input {
            width: 100px;
        }
`;
    }
    render() {
        return html`
        <label t-att-class="'oe_label' + (attrs.help ? '_help' : '')" t-att-title="attrs.help" t-att-for="element_id" t-att-style="style">
            <t t-esc="attrs.string || attrs.name"></t>
            <span t-if="attrs.help">?</span>
        </label>
        <div t-att-style="style">
            <input type="text" size="15" t-att-name="attrs.name" t-att-autofocus="attrs.default_focus === '1' ? 'autofocus' : undefined" t-att-id="element_id" t-att-value="defaults[attrs.name] || ''" />
            <t t-if="filters.length" t-raw="filters.render(defaults)" ></t>
        </div>
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
}

customElements.define('view-search-input-field', ViewSearchInputField);

