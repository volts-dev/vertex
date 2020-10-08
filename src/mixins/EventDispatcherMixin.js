/**
 * Backbone's events. Do not ever use it directly, use EventDispatcherMixin instead.
 *
 * This class just handle the dispatching of events, it is not meant to be extended,
 * nor used directly. All integration with parenting and automatic unregistration of
 * events is done in EventDispatcherMixin.
 *
 * Copyright notice for the following Class:
 *
 * (c) 2010-2012 Jeremy Ashkenas, DocumentCloud Inc.
 * Backbone may be freely distributed under the MIT license.
 * For all details and documentation:
 * http://backbonejs.org
 *
 */
class Events {
  on(events, callback, context) {
    var ev;
    events = events.split(/\s+/);
    var calls = this._callbacks || (this._callbacks = {});
    while ((ev = events.shift())) {
      var list = calls[ev] || (calls[ev] = {});
      var tail = list.tail || (list.tail = list.next = {});
      tail.callback = callback;
      tail.context = context;
      list.tail = tail.next = {};
    }
    return this;
  }

  off(events, callback, context) {
    var ev, calls, node;
    if (!events) {
      delete this._callbacks;
    } else if ((calls = this._callbacks)) {
      events = events.split(/\s+/);
      while ((ev = events.shift())) {
        node = calls[ev];
        delete calls[ev];
        if (!callback || !node) continue;
        while ((node = node.next) && node.next) {
          if (
            node.callback === callback &&
            (!context || node.context === context)
          )
            continue;
          this.on(ev, node.callback, node.context);
        }
      }
    }
    return this;
  }

  callbackList() {
    var lst = [];
    _.each(this._callbacks || {}, function (el, eventName) {
      var node = el;
      while ((node = node.next) && node.next) {
        lst.push([eventName, node.callback, node.context]);
      }
    });
    return lst;
  }

  trigger(events) {
    var event, node, calls, tail, args, all, rest;
    if (!(calls = this._callbacks)) return this;
    all = calls.all;
    (events = events.split(/\s+/)).push(null);
    // Save references to the current heads & tails.
    while ((event = events.shift())) {
      if (all)
        events.push({
          next: all.next,
          tail: all.tail,
          event: event,
        });
      if (!(node = calls[event])) continue;
      events.push({
        next: node.next,
        tail: node.tail,
      });
    }
    rest = Array.prototype.slice.call(arguments, 1);
    while ((node = events.pop())) {
      tail = node.tail;
      args = node.event ? [node.event].concat(rest) : rest;
      while ((node = node.next) !== tail) {
        node.callback.apply(node.context || this, args);
      }
    }
    return this;
  }
}

/**
 * Mixin containing an event system. Events are also registered by specifying the target object
 * (the object which will receive the event when it is raised). Both the event-emitting object
 * and the target object store or reference to each other. This is used to correctly remove all
 * reference to the event handler when any of the object is destroyed (when the destroy() method
 * from ParentedMixin is called). Removing those references is necessary to avoid memory leak
 * and phantom events (events which are raised and sent to a previously destroyed object).
 *
 * @name EventDispatcherMixin
 * @mixin
 */
export function EventDispatcherMixin(superClass) {
  return class extends superClass {
    //export class EventDispatcherMixin extends ParentedMixin {
    constructor() {
      super();
      this.__eventDispatcherMixin = true;
      this.custom_events = {};

      // init() {
      this.__edispatcherEvents = new Events();
      this.__edispatcherRegisteredEvents = [];
      this._delegateCustomEvents();
    }

    /**
     * Proxies a method of the object, in order to keep the right ``this`` on
     * method invocations.
     *
     * This method is similar to ``Function.prototype.bind`` or ``_.bind``, and
     * even more so to ``jQuery.proxy`` with a fundamental difference: its
     * resolution of the method being called is lazy, meaning it will use the
     * method as it is when the proxy is called, not when the proxy is created.
     *
     * Other methods will fix the bound method to what it is when creating the
     * binding/proxy, which is fine in most javascript code but problematic in
     * OpenERP Web where developers may want to replace existing callbacks with
     * theirs.
     *
     * The semantics of this precisely replace closing over the method call.
     *
     * @param {String|Function} method function or name of the method to invoke
     * @returns {Function} proxied method
     */
    proxy(method) {
      var self = this;
      return function () {
        var fn = typeof method === 'string' ? self[method] : method;
        if (fn === void 0) {
          throw new Error(
            "Couldn't find method '" + method + "' in widget " + self
          );
        }
        return fn.apply(self, arguments);
      };
    }

    _delegateCustomEvents() {
      if (_.isEmpty(this.custom_events)) {
        return;
      }
      for (var key in this.custom_events) {
        if (!this.custom_events.hasOwnProperty(key)) {
          continue;
        }

        var method = this.proxy(this.custom_events[key]);
        this.on(key, this, method);
      }
    }

    on(events, dest, func) {
      var self = this;
      if (typeof func !== 'function') {
        throw new Error('Event handler must be a function.');
      }
      events = events.split(/\s+/);
      _.each(events, function (eventName) {
        self.__edispatcherEvents.on(eventName, func, dest);
        if (dest && dest.__eventDispatcherMixin) {
          dest.__edispatcherRegisteredEvents.push({
            name: eventName,
            func: func,
            source: self,
          });
        }
      });
      return this;
    }

    off(events, dest, func) {
      var self = this;
      events = events.split(/\s+/);
      _.each(events, function (eventName) {
        self.__edispatcherEvents.off(eventName, func, dest);
        if (dest && dest.__eventDispatcherMixin) {
          dest.__edispatcherRegisteredEvents = _.filter(
            dest.__edispatcherRegisteredEvents,
            function (el) {
              return !(
                el.name === eventName &&
                el.func === func &&
                el.source === self
              );
            }
          );
        }
      });
      return this;
    }

    once(events, dest, func) {
      // similar to this.on(), but func is executed only once
      var self = this;
      if (typeof func !== 'function') {
        throw new Error('Event handler must be a function.');
      }
      self.on(events, dest, function what() {
        func.apply(this, arguments);
        self.off(events, dest, what);
      });
    }

    trigger() {
      this.__edispatcherEvents.trigger.apply(
        this.__edispatcherEvents,
        arguments
      );
      return this;
    }

    trigger_up(name, info) {
      /*var event = new OdooEvent(this, name, info);
    //console.info('event: ', name, info);
    this._trigger_up(event);
    return event;*/
      this.dispatchEvent(
        new CustomEvent(name, { detail: info, bubbles: true, composed: true })
      );
    }

    _trigger_up(event) {
      var parent;
      this.__edispatcherEvents.trigger(event.name, event);
      if (!event.is_stopped() && (parent = this.getParent())) {
        parent._trigger_up(event);
      }
    }

    destroy() {
      var self = this;
      _.each(this.__edispatcherRegisteredEvents, function (event) {
        event.source.__edispatcherEvents.off(event.name, event.func, self);
      });
      this.__edispatcherRegisteredEvents = [];
      _.each(
        this.__edispatcherEvents.callbackList(),
        function (cal) {
          this.off(cal[0], cal[2], cal[1]);
        },
        this
      );
      this.__edispatcherEvents.off();
      ParentedMixin.destroy.call(this);
    }
  };
}
