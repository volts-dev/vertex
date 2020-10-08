import _ from "underscore";

//var Vectors = window.Vectors = window.Vectors || {};
//var Class = Vectors.Class(); // # 执行构造函数

// # Registry 构造函数
export class Registry {
  /**
   * @constructor
   * @param {Object} [mapping] the initial data in the registry
   */
  constructor(mapping) {
    this.map = Object.create(mapping || null);
    this._scoreMapping = Object.create(null);
    this._sortedKeys = null;
    this.listeners = []; // listening callbacks on newly added items.
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * Add a key (and a value) to the registry.
   *
   * Notify the listeners on newly added item in the registry.
   *
   * @param {string} key
   * @param {any} value
   * @param {number} [score] if given, this value will be used to order keys
   * @returns {Registry} can be used to chain add calls.
   */
  add(key, value, score) {
    this._scoreMapping[key] = score === undefined ? key : score;
    this._sortedKeys = null;
    this.map[key] = value;
    _.each(this.listeners, function(callback) {
      callback(key, value);
    });
    return this;
  }
  /**
   * Check if the registry contains the key
   *
   * @param {string} key
   * @returns {boolean}
   */
  contains(key) {
    return key in this.map;
  }
  /**
   * Returns the content of the registry (an object mapping keys to values)
   *
   * @returns {Object}
   */
  entries() {
    return Object.create(this.map);
  }
  /**
   * Returns the value associated to the given key.
   *
   * @param {string} key
   * @returns {any}
   */
  get(key) {
    return this.map[key];
  }
  /**
   * Tries a number of keys, and returns the first object matching one of
   * the keys.
   *
   * @param {string[]} keys a sequence of keys to fetch the object for
   * @returns {any} the first result found matching an object
   */
  getAny(keys) {
    for (var i = 0; i < keys.length; i++) {
      if (keys[i] in this.map) {
        return this.map[keys[i]];
      }
    }
    return null;
  }
  /**
   * Return the list of keys in map object.
   *
   * The registry guarantees that the keys have a consistent order, defined by
   * the 'score' value when the item has been added.
   *
   * @returns {string[]}
   */
  keys() {
    var self = this;
    if (!this._sortedKeys) {
      this._sortedKeys = _.sortBy(Object.keys(this.map), function(key) {
        return self._scoreMapping[key] || 0;
      });
    }
    return this._sortedKeys;
  }
  /**
   * Register a callback to execute when items are added to the registry.
   *
   * @param {function} callback function with parameters (key, value).
   */
  onAdd(callback) {
    this.listeners.push(callback);
  }
  /**
   * Return the list of values in map object
   *
   * @returns {string[]}
   */
  values() {
    var self = this;
    return this.keys().map(function(key) {
      return self.map[key];
    });
  }

  /**
   * Creates and returns a copy of the current mapping, with the provided
   * mapping argument added in (replacing existing keys if needed)
   *
   * Parent and child remain linked, a new key in the parent (which is not
   * overwritten by the child) will appear in the child.
   *
   * @param {Object} [mapping={}] a mapping of keys to object-paths
   */
  extend(mapping) {
    var child = new Registry(this.map);
    _.extend(child.map, mapping);
    return child;
  }
}
