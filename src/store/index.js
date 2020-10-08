import { MobxLitElement } from '@adobe/lit-mobx';
import { App } from './modules/app';
import { Session } from './modules/session';

// 构建全局store
var app = new App();
var session = new Session();

// a Abstract element with Mobx store
export class Element extends MobxLitElement {
  // 为对象属性成员提供支持默认值"value"
  static getPropertyDescriptor(name, key, options) {
    return {
      get() {
        return this[key] === undefined ? options.value : this[key];
      },
      set(value) {
        const oldValue = this[key];
        if (value === undefined) {
          value = options.value;
        }
        this[key] = value;
        this.requestUpdate(name, oldValue);
      },
      configurable: true,
      enumerable: true,
    };
    return descriptor;
  }

  constructor() {
    super();
    this.$app = app; // <==传递store
    this.$session = session; // <==传递store
  }
}

export default {
  Element,
  app,
  session,
};
