import { LitElement } from 'lit-element';
import '../../core/class.js';
var Vectors = window.Vectors = window.Vectors || {};
var Class = Vectors.Class(); // # 执行构造函数

class SearchInputMixin extends LitElement {
    constructor() {
        super();
    }

    static get properties() {
        return {
            FilterGroup: FilterGroup,
            Filter: Filter,
            Field: Field,
            GroupbyGroup: GroupbyGroup,
        }
    }
}

