import { EventDispatcherMixin } from '@/mixins/EventDispatcherMixin';
import { ServicesCallerMixin } from '@/mixins/ServicesCallerMixin';

// 各项 功能服务原型 如Ajax调用服务 打印服务 报表服务等
export class AbstractService extends ServicesCallerMixin(
  EventDispatcherMixin(Object)
) {
  constructor(...args) {
    super(...args);

    this.dependencies = [];
    //this.setParent(parent);
  }

  /**
   * @abstract
   */
  start() {}
}
