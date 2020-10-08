import { html, LitElement } from 'lit-element';

window.vectors = window.vectors || {};
vectors.elements = vectors.elements || {};
vectors.elements.search = vectors.elements.search || {};

vectors.elements.search.CompoundContext = {
    properties: {
        __ref: {
            type: String,
            value: "compound_domain",
        },

        __domains: {
            type: Array,
        },

        __eval_context: {
            type: Object,
        },
    },

    factoryImpl() {
        this.__ref = "compound_context";
        this.__domains = [];
        this.__eval_context = null;

        // # 添加参数
        var self = this;
        var args = [...arguments]
        args.forEach(function (x) {
            self.add(x);
        });
    },

    add(context) {
        this.__contexts.push(context);
        return this;
    },
    set_eval_context(eval_context) {
        this.__eval_context = eval_context;
        return this;
    },
    get_eval_context() {
        return this.__eval_context;
    },
    eval() {
        return pyeval.eval('context', this, undefined, { no_user_context: true });
    },
}


vectors.elements.search.CompoundDomain = {
    properties: {
        __ref: {
            type: String,
            value: "compound_domain",
        },

        __domains: {
            type: Array,
        },

        __eval_context: {
            type: Object,
        },
    },

    // factoryImpl created
    factoryImpl() {
        this.__ref = "compound_domain";
        this.__domains = [];
        this.__eval_context = null;

        // # 添加参数
        var self = this;
        var args = [...arguments]
        args.forEach(function (x) {
            self.add(x);
        });
    },

    add(domain) {
        this.__domains.push(domain);
        return this;
    },

    set_eval_context(eval_context) {
        this.__eval_context = eval_context;
        return this;
    },

    get_eval_context() {
        return this.__eval_context;
    },

    eval() {
        return pyeval.eval('domain', this);
    },
};


export class Field extends LitElement {
    static get properties() {
        return {
            host: { type: Object },
            attrs: { type: Object },
            default_operator: { type: String, value: "=" },
            style: { type: String },
            defaults: { type: Object }
        };
    }

    render() {
        return html`
            <!--显示Label-->
            <label t-att-class="'oe_label' + ${this.getClass}" title="${this.attrs.help}" for="${this.element_id}" style="${this.style}">
                ${this.attrs.string
                ? html`this.{attrs.string}`
                : html`${this.attrs.name
                    ? html`${this.attrs.name}`
                    : html``}`}
                                        
                <!--显示Help-->
                ${this.attrs.help
                ? html`${this.attrs.help}`
                : html`?`}
            </label>

            <div style="${this.style}">
                <input type="text" size="15" name="${this.attrs.name}" ?autofocus="${this.getFocus}" id="element_id"
                    .value="${this.getDefaults}" />
                <t t-if="filters.length" t-raw="filters.render(defaults)"></t>
            </div>
            `;
    }

    constructor() {
        super();
        /* 废弃
        //  this._super(parent);
        this.host = parent;
        //this.load_attrs(_.extend({}, field, view_section.attrs));
        var attr = {}

        for (var att in field) {
            attr[att] = field[att];
        }

        // # 遍历Field Html节点
        for (var key in view_section.attributes) {
            attr[view_section.attributes[key].name] = view_section.attributes[key].value;
        }
        this.loadAttrs(attr);
        */
    }

    init(parent) {
        // this._super(parent);
        this.searchview = parent;
        this.loadAttrs({});
    }

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
        return Promise.resolve(null);
    }

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
        if (!this.attrs || !(this.attrs.name in defaults && defaults[this.attrs.name])) {
            return Promise.resolve(null);
        }
        return this.facet_for(defaults[this.attrs.name]);
    }

    /**
     * Returns whether the input is "visible". The default behavior is to
     * query the ``modifiers.invisible`` flag on the input's description or
     * view node.
     *
     * @returns {Boolean}
     */
    visible() {
        return !this.attrs.modifiers.invisible;
    }

    getClass() {
        if (this.attrs.help) {
            return "_help";
        }
        return "";
    }

    getFocus() {
        if (attrs.default_focus === '1') {
            return 'autofocus'
        }
        return undefined;
    }

    getDefaults() {
        if (this.defaults[this.attrs.name]) {
            return this.defaults[attrs.name]
        }
        return "";
    }



    loadAttrs(attrs) {
        var type = typeof attrs.modifiers;
        if (type != 'function' || type != 'object' && !!attrs.modifiers) {
            //  if (!_.isObject(attrs.modifiers)) {
            // # 变更
            attrs.modifiers = attrs.modifiers ? JSON.parse(attrs.modifiers) : {};
        }
        this.attrs = attrs;
    }

    facet_for(value) {
        return Promise.resolve({
            field: this,
            category: this.attrs.string || this.attrs.name,
            values: [{ label: String(value), value: value }]
        });
    }

    value_from(facetValue) {
        return facetValue['value'];
    }

    get_context(facet) {
        var self = this;
        // A field needs a context to send when active
        var context = this.attrs.context;
        if (vectors.utils.isEmpty(context) || !facet.values.length) {
            return;
        }
        var contexts = facet.values.map(function (facetValue) {
            return new vectors.elements.search.CompoundContext(context)
                .set_eval_context({ self: self.value_from(facetValue) });
        });

        if (contexts.length === 1) { return contexts[0]; }

        var domain = vectors.elements.search.CompoundContext();
        domain['__domains'] = domains;
        return domain;
        return vectors.utils.extend(new vectors.elements.search.CompoundContext(), {
            __contexts: contexts
        });
    }

    /**
     * Function creating the returned domain for the field, override this
     * methods in children if you only need to customize the field's domain
     * without more complex alterations or tests (and without the need to
     * change override the handling of filter_domain)
     *
     * @param {String} name the field's name
     * @param {String} operator the field's operator (either attribute-specified or default operator for the field
     * @param {Number|String} facet parsed value for the field
     * @returns {Array<Array>} domain to include in the resulting search
     */
    make_domain(name, operator, facet) {
        return [[name, operator, this.value_from(facet)]];
    }

    get_domain(facet) {
        if (!facet.values.length) { return; }

        var value_to_domain;
        var self = this;
        var domain = this.attrs.filter_domain;
        if (domain) {
            value_to_domain = function (facetValue) {
                var cd = new vectors.elements.search.CompoundDomain(domain)
                return cd.set_eval_context({ self: self.value_from(facetValue) });
            };
        } else {
            value_to_domain = function (facetValue) {
                return self.make_domain(
                    self.attrs.name,
                    self.attrs.operator || self.default_operator,
                    facetValue);
            };
        }

        var domains = facet.values.map(value_to_domain);

        if (domains.length === 1) { return domains[0]; }
        for (var i = domains.length; --i;) {
            domains.unshift(['|']);
        }

        // var domain = vectors.elements.search.CompoundDomain();
        // domain['__domains'] = domains;
        // return domain;
        return vectors.utils.extend(new vectors.elements.search.CompoundDomain, {
            __domains: domains
        });
    }

    get_groupby() {
        throw new Error(
            "get_groupby not implemented for widget " + this.attrs.type);
    }
}
