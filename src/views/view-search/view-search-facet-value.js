import { html,  LitElement } from 'lit-element';

class ViewSearchFacetValue extends LitElement {
    render(){
        return html`
        <span t-name="SearchView.FacetView.Value">
            <t t-esc="widget.model.get('label')"></t>
        </span>
    `;
    }

    static get properties() {
        return {
            string: {
                type: String,
                notify: true
            },

            title: {
                type: String,
                notify: true
            }
        }
    }
}

customElements.define('view-search-facet-value', ViewSearchFacetValue);