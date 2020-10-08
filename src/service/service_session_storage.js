/**
 * This module defines a service to access the sessionStorage object.
 */

import { storage } from '@/core/storage_session';
import { AbstractService } from './service';

export class AbstractStorageService extends AbstractService {
  static get properties() {
    return {
      // the 'storage' attribute must be set by actual StorageServices extending
      // this abstraction
      storage: Object,
    };
  }

  constructor() {
    super();
  }

  /**
   * @override
   */
  destroy() {
    // storage can be permanent or transient, destroy transient ones
    if ((this.storage || {}).destroy) {
      this.storage.destroy();
    }
    this._super.apply(this, arguments);
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * Removes all data from the storage
   */
  clear() {
    this.storage.clear();
  }
  /**
   * Returns the value associated with a given key in the storage
   *
   * @param {string} key
   * @returns {string}
   */
  getItem(key, defaultValue) {
    var val = this.storage.getItem(key);
    return val ? JSON.parse(val) : defaultValue;
  }
  /**
   * @param {integer} index
   * @return {string}
   */
  key(index) {
    return this.storage.key(index);
  }
  /**
   * @return {integer}
   */
  length() {
    return this.storage.length;
  }
  /**
   * Removes the given key from the storage
   *
   * @param {string} key
   */
  removeItem(key) {
    this.storage.removeItem(key);
  }
  /**
   * Sets the value of a given key in the storage
   *
   * @param {string} key
   * @param {string} value
   */
  setItem(key, value) {
    this.storage.setItem(key, JSON.stringify(value));
  }
  /**
   * Add an handler on storage event
   *
   */
  onStorage() {
    this.storage.on.apply(
      this.storage,
      ['storage'].concat(Array.prototype.slice.call(arguments))
    );
  }
}

export class SessionStorageService extends AbstractStorageService {
  constructor() {
    super();
    this.storage = storage;
  }
}
