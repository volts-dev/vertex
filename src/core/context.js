import { pyUtils } from "./py_utils";
import _ from "underscore";

export class Context {
  constructor() {
    this.__ref = "compound_context";
    this.__contexts = [];
    this.__eval_context = null;
    var self = this;
    _.each(arguments, function(x) {
      self.add(x);
    });
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  add(context) {
    this.__contexts.push(context);
    return this;
  }

  eval() {
    return pyUtils.eval("context", this);
  }
  /**
   * Set the evaluation context to be used when we actually eval.
   *
   * @param {Object} evalContext
   * @returns {Context}
   */
  set_eval_context(evalContext) {
    // a special case needs to be done for moment objects.  Dates are
    // internally represented by a moment object, but they need to be
    // converted to the server format before being sent. We call the toJSON
    // method, because it returns the date with the format required by the
    // server
    for (var key in evalContext) {
      if (evalContext[key] instanceof moment) {
        evalContext[key] = evalContext[key].toJSON();
      }
    }
    this.__eval_context = evalContext;
    return this;
  }
}
