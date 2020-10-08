/**
 * Owner of the state, this component is tasked with fetching data, processing
 * it, updating it, ...
 *
 * Note that this is not a Element: it is a class which has not UI representation.
 *
 * @class Model
 */

import { ServicesCallerMixin } from '@/mixins/ServicesCallerMixin';
import { EventDispatcherMixin } from '@/mixins/EventDispatcherMixin';

// 模型原型
export class AbstractModel extends ServicesCallerMixin(
  EventDispatcherMixin(Object)
) {
  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * This method should return the complete state necessary for the renderer
   * to display the current data.
   *
   * @returns {*}
   */
  get() {}
  /**
   * The load method is called once in a model, when we load the data for the
   * first time.  The method returns (a promise that resolves to) some kind
   * of token/handle.  The handle can then be used with the get method to
   * access a representation of the data.
   *
   * @param {Object} params
   * @returns {Promise} The promise resolves to some kind of handle
   */
  load() {
    return Promise.resolve();
  }
}
