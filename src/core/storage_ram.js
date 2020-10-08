/**
 * This module defines an alternative of the Storage objects (localStorage,
 * sessionStorage), stored in RAM. It is used when those native Storage objects
 * are unavailable (e.g. in private browsing on Safari).
 */

//import { EventDispatcherMixin } from "../mixins/EventDispatcherMixin";

export class RamStorage {
  /**
   * @constructor
   */
  constructor() {
    //mixins.EventDispatcherMixin.init.call(this);
    if (!this.storage) {
      this.clear();
    }
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * Removes all data from the storage
   */
  clear() {
    this.storage = Object.create(null);
    this.length = 0;
  }
  /**
   * Returns the value associated with a given key in the storage
   *
   * @param {string} key
   * @returns {string}
   */
  getItem(key) {
    return this.storage[key];
  }
  /**
   * @param {integer} index
   * @return {string}
   */
  key(index) {
    return _.keys(this.storage)[index];
  }
  /**
   * Removes the given key from the storage
   *
   * @param {string} key
   */
  removeItem(key) {
    if (key in this.storage) {
      this.length--;
    }
    delete this.storage[key];
    //this.trigger("storage", { key: key, newValue: null });
  }
  /**
   * Adds a given key-value pair to the storage, or update the value of the
   * given key if it already exists
   *
   * @param {string} key
   * @param {string} value
   */
  setItem(key, value) {
    if (!(key in this.storage)) {
      this.length++;
    }
    this.storage[key] = value;
    //this.trigger("storage", { key: key, newValue: value });
  }
}
