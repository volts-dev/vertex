import { html, css } from 'lit-element';
import { FormElement } from './../form-element';
//import { DataSource } from '../../../element/dataset/data-source.js';
import '@/elements/ve-dropdown/ve-dropdown';
import { styles } from './form-selection-styles.js';

class FormSelection extends FormElement {
  static get styles() {
    return [super.styles, styles, css``];
  }

  static get properties() {
    return {
      isEditable: { type: Boolean }, // 可创建并编辑
      isSearchable: Boolean,
      isReference: { type: Boolean },
      /**
       * An array of values to be displayed as options in the dropdown. The
       * options could be either `String` or `Object`.
       */
      items: { type: Array },
      value: { type: Array }, // 已经选择的值
      type: { type: String }, // 控件格式
      // 存储[KEY]NAME
      itemsMap: {
        type: Object,
      },

      /**
       * If `true`, user can input a value that is not present in the items list.
       * `value` property will be set to the input value in this case.
       * Also, when `value` is set programmatically, the input value will be set
       * to reflect that value.
       */
      allowCustomValue: {
        type: Boolean,
        // value: false
      },

      _positionTarget: {
        type: Object,
        // value: function() {
        //   return this.$.inputContainer;
        // }
      },
      _focusedIndex: {
        type: Number,
        // value: -1
      },
      _filter: {
        type: String,
        //  value: ""
      },
      /**
       * The selected item from the `items` array.
       */
      selectedItem: {
        type: Object,
        //  readOnly: true,
        //  notify: true
      },
    };
  }

  render() {
    return html`
      ${this.mode == 'edit'
        ? html`
            <!-- 修改模式 !-->
            <ve-dropdown id="dropdown" .anchor=${this} fullwidth="true">
              <ul class="dropdown-list">
                ${this.isSearchable
                  ? html` <li class="option"><a>Search More...</a></li> `
                  : null}
                ${this.isEditable
                  ? html` <li class="option"><a>Create and Edit</a></li> `
                  : null}
                ${this.items.map(
                  (item, index) => html`
                    <li class="option">
                      <a
                        @click="${this.onSelected}"
                        .value=${item.id}
                        .name=${item.name}
                        >${item.name}</a
                      >
                    </li>
                  `
                )}
              </ul>
            </ve-dropdown>
            <div
              id="SearchView"
              class="searchview form-control input-sm horizontal layout center"
            >
              <iron-icon
                icon="search"
                class="searchview_more fa fa-search-minus"
                title="Advanced Search..."
                @click="${this.onMoreTapped}"
              ></iron-icon>
              <div class="searchview_facets">
                ${this.value.map(
                  (item, index) => html`
                    <div class="searchview_facet" .index=${index}>
                      <span class="facet_values_sep" value=${item.value}
                        >${item.name}</span
                      >
                      <iron-icon
                        icon="clear"
                        class="facet_remove"
                        @tap="${this.onRemoveFacet}"
                      ></iron-icon>
                    </div>
                  `
                )}
                ${this.isSearchable
                  ? html`
                      <input
                        id="input"
                        class="input_box "
                        placeholder="Search..."
                        @change="${this.onChange}"
                      />
                    `
                  : null}
              </div>
            </div>
          `
        : html`
            <div class="searchview_facets">
              ${this.value.map(
                (item, index) => html`
                  <div class="searchview_facet" .index=${index}>
                    <a
                      ><span class="facet_values_sep" value=${item.value}
                        >${item.name}</span
                      ></a
                    >
                  </div>
                `
              )}
            </div>
          `}
    `;
  }

  constructor() {
    super();
    this.items = [];
    this.value = [];

    this.isSearchable = true;
    this.isEditable = true;
    this.value = [
      { name: 'Your Name', value: 1111 },
      { name: 'My Name', value: 2222 },
    ];
  }

  firstUpdated() {
    super.firstUpdated();
    this.dropdown = this.shadowRoot.querySelector('#dropdown');

    // The dom-change event signifies when the template has stamped its DOM.
    /* this.listen(this, 'dom-change', "onDomChange");
   // 监控Input值变化
   this.listen(this.$.input, 'change', "onInputValueChanged");
   this.$.input.addEventListener('value-changed', function () {
       var val = this.$.input.value;
       for (var key in this.itemsMap) {
           if (this.itemsMap[key] == val) {
               this.value = key; // id
           }
       }

   });*/
  }

  onSelected(event) {
    var sender = event.srcElement;
    if (sender) {
      var vals = _.clone(this.value);
      var has = false;
      vals.some(value => {
        if (value.value == sender.value) {
          value.name = sender.name;
          has = true;
          return true;
        }
      });
      if (!has) {
        vals.push({ value: sender.value, name: sender.name });
      }

      this.value = vals;
    }
  }

  onRemoveFacet(event) {
    var parent = event.srcElement.parentElement;
    if (parent) {
      var vals = _.clone(this.value);
      delete vals[parent.index];
      this.value = vals;
    }
  }

  onMoreTapped(event) {
    if (!this.dropdown) {
      this.dropdown = this.shadowRoot.querySelector('#dropdown');
    }

    if (!this.dropdown.open) {
      this.nameSearch();

      this.dropdown.open = true;
    }
  }

  onItemsChanged(items) {
    if (this.items) {
      // && this.items.indexSplices
      // 生成Map
      /* var data = [];
        this.items.forEach(function (item) {
            data[item.id] = item[1];
        });
        this.itemsMap = data;
*/
      // 更新显示值
      this.updateDisplayValue();
    }
  }

  get mode() {
    return this._mode;
  }

  set mode(value) {
    let oldValue = this._mode;
    this._mode = value;
    this.requestUpdate('mode', oldValue);
  }

  get value() {
    return this._value;
  }

  set value(value) {
    super.value = value;
    let oldVal = this._value;
    this._value = value;
    this.requestUpdate('value', oldVal);

    // 更新显示值
    this.dispatchEvent(
      new CustomEvent('on-value-changed', {
        detail: { field: this.field, value: this.value },
      })
    );
  }

  onEditMode(editMode) {
    // 更新下拉结果
    if (this.editMode) {
      if (!this.items || (this.items && this.items.length == 0)) {
        // 只当无数据时
        this.nameSearch();
      }
    }
  }

  onInputValueChanged(e) {
    var val = this.$.input.value;
    for (var key in this.itemsMap) {
      if (this.itemsMap[key] == val) {
        this.value = key; // id
      }
    }
  }

  nameSearch() {
    // 执行API
    if (this.field) {
      var self = this;

      // #'{"views":JSON.stringify([[act_win.views]]),"model":"[[act_win.res_model]]"}
      var data = {};
      //data.views = this.action.views;
      // # 添加SearchView
      /// if (this.action.search_view_id) {
      //     data.views.push([this.action.search_view_id, "search"]);
      // }
      data.model = this.field.relation; // model
      data.name = '';
      data.limit = 10 || 0;
      data.operator = 'ilike';
      return this._rpc({
        route: '/dataset/call_kw/' + this.field.relation + '/name_search',
        args: data,
        //context: session.user_context,
      }).then(function (data) {
        self.items = data;
      });

      ///this.items = elements;
      //var combobox = this.$$.combobox;
      //     combobox.items =ds.data;
      /// var combobox = this.$.demo4;
      // Array of String values.
      ///combobox.items = elements;
    } else {
      this.datasource.action = '';
    }
  }
}

customElements.define('form-selection', FormSelection);
