import { html, css, LitElement } from 'lit-element';
import { Element } from '@/mixins/element';

export class FormElement extends Element {
  static get properties() {
    return {
      field: { type: Object }, // 存储字段对象
      invalid: { type: Boolean },
      invisible: { type: Boolean, value: false },
      editMode: { type: Boolean }, // 必须默认值是可供Template判断
      mode: { type: String }, // 显示模式 Can be "edit", "readonly"
      label: { type: String }, //The label for this element.
      help: { type: String },
      value: { type: String }, // The value for this input.
      displayValue: { type: String }, // M2M O2M关系时显示值
      /**
       * The type of the input. The supported types are `text`, `number` and `password`.
       * If you're using PaperInputBehavior to implement your own paper-input-like element,
       * bind this to the `<input is="iron-input">`'s `type` property.
       */
      type: { type: String },

      /**
       * Set to true to mark the input as required. If you're using PaperInputBehavior to
       * implement your own paper-input-like element, bind this to
       * the `<input is="iron-input">`'s `required` property.
       */
      required: {
        type: Boolean,
        value: false,
      },

      /**
       * The minimum length of the input value.
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `minlength` property.
       */
      minlength: {
        type: Number,
      },
      /**
       * The maximum length of the input value.
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `maxlength` property.
       */
      maxlength: {
        type: Number,
      },
      /**
       * The minimum (numeric or date-time) input value.
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `min` property.
       */
      min: {
        type: String,
      },
      /**
       * The maximum (numeric or date-time) input value.
       * Can be a String (e.g. `"2000-1-1"`) or a Number (e.g. `2`).
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `max` property.
       */
      max: {
        type: String,
      },

      /**
       * Limits the numeric or date-time increments.
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `step` property.
       */
      step: {
        type: String,
      },
      /**
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `name` property.
       */
      name: {
        type: String,
      },
      /**
       * A placeholder string in addition to the label. If this is set, the label will always float.
       */
      placeholder: {
        type: String,
        // need to set a default so _computeAlwaysFloatLabel is run
        value: '',
      },
      /**
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `readonly` property.
       */
      readonly: {
        type: Boolean,
        value: false,
      },
      /**
       * If you're using PaperInputBehavior to implement your own paper-input-like
       * element, bind this to the `<input is="iron-input">`'s `size` property.
       */
      size: {
        type: Number,
      },
    };
  }

  extendTemplate(tpl) {
    return html`
      ${this.mode == 'edit'
        ? html` ${tpl} `
        : html` <p>${this.displayValue || this.value || ''} ddd</p> `}
    `;
  }
  static get styles() {
    return css`
      [type='text'],
      [type='password'],
      [type='number'],
      textarea,
      select {
        width: 100%;
        display: block;
        outline: none;
      }

      .input {
        border: 1px solid #ccc;
        border-radius: 3px;
        padding: 2px 4px;
        color: #1f1f1f;
      }

      :host(.required_modifier) input {
        background-color: #d2d2ff !important;
      }
    `;
  }

  readonlyTemplate(tpl) {
    return html` <p>${this.displayValue || this.value || ''} ddd</p> `;
  }

  constructor() {
    super();
    this.mode = 'readonly';
  }

  firstUpdated() {
    super.firstUpdated();
  }

  get value() {
    return this._value;
  }

  set value(value) {
    let oldVal = this._value;
    this._value = value;
    this.requestUpdate('value', oldVal);
    this.dispatchEvent(
      new CustomEvent('on-value-changed', {
        detail: { field: this.field, value: this.value },
        bubbles: true,
        composed: true,
      })
    );
  }

  get mode() {
    return this._mode;
  }
  set mode(mode) {
    let oldVal = this._mode;
    this._mode = mode;
    this.requestUpdate('mode', oldVal);

    if (mode == 'create' || mode == 'edit') {
      this.editMode = true;
    } else {
      this.editMode = false;
    }

    if (this.modeChanged) {
      this.modeChanged(mode);
    }
  }

  set invisible(invisible) {
    let oldVal = this._invisible;
    this._invisible = invisible;
    this.requestUpdate('invisible', oldVal);

    if (invisible) {
      this.style.display = 'none';
    } else {
      this.style.display = 'block';
    }
  }

  get_value() {
    return this.value;
  }

  isValid() {
    return (
      this.is_syntax_valid() &&
      !(this.getAttribute('required') && this.is_false())
    );
  }

  is_syntax_valid() {
    return true;
  }
  /**
   * Method useful to implement to ease validity testing. Must return true if the current
   * value is similar to false in OpenERP.
   */
  is_false() {
    return this.getAttribute('value') === false;
  }

  IsEditMode(mode) {
    if (mode == 'create' || mode == 'edit') {
      return true;
    } else {
      return false;
    }
  }
}
