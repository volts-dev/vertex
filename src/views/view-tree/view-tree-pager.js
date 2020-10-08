import { html, PolymerElement } from '@polymer/polymer/polymer-element.js';
import '@polymer/iron-flex-layout/iron-flex-layout.js';
import '@polymer/paper-icon-button/paper-icon-button.js';

/**
 * `view-tree`
 * 
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class ViewTreePager extends PolymerElement {
    static get template() {
        return html`
   <style>
        .btn {
            padding: 0;
            width: 30px;
            height: 30px;
        }
    </style>
        <div>
            <span class="o_pager_value">1-[[count]]</span> / <span class="o_pager_limit">[[page]]</span>
            <span class="btn-group btn-group-sm">
<paper-icon-button icon="chevron-left" class="fa fa-chevron-left btn btn-icon o_pager_previous" type="button" accesskey="p"  on-tap="_prev"></paper-icon-button>
<paper-icon-button icon="chevron-right"class="fa fa-chevron-right btn btn-icon o_pager_next" type="button" accesskey="n"  on-tap="_next"></paper-icon-button>
</span>
        </div>
    `;
    }
    static get properties() {
        return {
            prop1: {
                type: String,
                value: 'view-tree',
            },

            page: Number,//当前也号
            pages: Array,//页序列号
            size: Number,//页大小
            count: Number,// 页数
        };
    }

    static get observers() {
        return [
            // 'onPageChanged(page)',
            'onPagesChanged(pages)',
        ];
    }


    ready() {
        super.ready();

        this.page = 1;
    }

    onPagesChanged(pages) {
        // 初始化
        this.page = 1;
        this.count = pages.length;
        this.dispatchEvent(new CustomEvent("page-changed",{detail:this.page}));
    }

    _isSelected(page, item) {
        return page === item - 1;
        this.dispatchEvent(new CustomEvent("page-changed",{detail:this.page}));
    }

    _select(e) {
        this.page = e.model.item - 1;
        this.dispatchEvent(new CustomEvent("page-changed",{detail:this.page}));
    }

    _next() {
        this.page = Math.min(this.size - 1, this.page + 1);
        this.dispatchEvent(new CustomEvent("page-changed",{detail:this.page}));
    }

    _prev() {
        this.page = Math.max(1, this.page - 1);
        //this.fire("page-changed", this.page);
        this.dispatchEvent(new CustomEvent("page-changed",this.page));

    }
}

window.customElements.define('view-tree-pager', ViewTreePager);
