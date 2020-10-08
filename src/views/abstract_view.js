import { Element } from '@/mixins/Element';
import { AbstractModel } from './abstract_model';
import { AbstractRenderer } from './abstract_renderer';
import { AbstractController } from './abstract_controller';
import Ajax from '@/core/ajax';

// 渲染原型
export class AbstractView extends Element {
  constructor() {
    super(...arguments);
    this.rendererParams = {};
    this.controllerParams = {};
    this.modelParams = {};
    this.loadParams = {};

    // determines the MVC components to use
    this.config = {
      Model: AbstractModel,
      Renderer: AbstractRenderer,
      Controller: AbstractController,
    };
  }

  //--------------------------------------------------------------------------
  // Public
  //--------------------------------------------------------------------------

  /**
   * Main method of the Factory class. Create a controller, and make sure that
   * data and libraries are loaded.
   *
   * There is a unusual thing going in this method with parents: we create
   * renderer/model with parent as parent, then we have to reassign them at
   * the end to make sure that we have the proper relationships.  This is
   * necessary to solve the problem that the controller needs the model and
   * the renderer to be instantiated, but the model need a parent to be able
   * to load itself, and the renderer needs the data in its constructor.
   *
   * @param {Widget} parent the parent of the resulting Controller (most
   *      likely an action manager)
   * @returns {Promise<Controller>}
   */
  getController(parent) {
    //var _super = this._super.bind(this);
    var self = this;
    var model = this.getModel(parent);
    return Promise.all([this._loadData(model), Ajax.loadLibs(this)]).then(
      function (result) {
        var state = result[0];
        var renderer = self.getRenderer(parent, state);
        var Controller = self.Controller || self.config.Controller;
        var controllerParams = _.extend(
          {
            initialState: state,
          },
          self.controllerParams
        );

        var controller = new Controller(
          parent,
          model,
          renderer,
          controllerParams
        );
        //model.setParent(controller);
        //renderer.setParent(controller);
        return controller;
      }
    );
  }
  /**
   * Returns a new model instance
   *
   * @param {Widget} parent the parent of the model
   * @returns {Model} instance of the model
   */
  getModel(parent) {
    var Model = this.config.Model;
    return new Model(parent, this.modelParams);
  }
  /**
   * Returns a new renderer instance
   *
   * @param {Widget} parent the parent of the renderer
   * @param {Object} state the information related to the rendered data
   * @returns {Renderer} instance of the renderer
   */
  getRenderer(parent, state) {
    var Renderer = this.config.Renderer;
    return new Renderer(parent, state, this.rendererParams);
  }

  //--------------------------------------------------------------------------
  // Private
  //--------------------------------------------------------------------------

  /**
   * Loads initial data from the model
   *
   * @private
   * @param {Model} model a Model instance
   * @returns {Promise<*>} a promise that resolves to the value returned by
   *   the get method from the model
   * @todo: get rid of loadParams (use modelParams instead)
   */
  _loadData(model) {
    return model.load(this.loadParams).then(function () {
      return model.get.apply(model, arguments);
    });
  }
}

customElements.define('abstract-view', AbstractView);
