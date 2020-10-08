/**
 * The controller has to coordinate between parent, renderer and model.
 *
 * @class Controller
 */
import { Element } from '@/mixins/Element';

export class AbstractController extends Element {
  static get properties() {
    return {
      _title: {
        type: String,
      },
      _controlPanel: {
        type: Object,
      },
    };
  }

  /**
   * @override
   * @param {Model} model
   * @param {Renderer} renderer
   * @param {Object} params
   */
  constructor(parent, model, renderer, params) {
    super();
    // init: function (parent, model, renderer, params) {
    //   this._super.apply(this, arguments);
    this.model = model;
    this.renderer = renderer;
  }

  /**
   * @returns {Promise}
   */
  firstUpdated() {
    return Promise.all([
      super.firstUpdated(...arguments),
      this._startRenderer(),
    ]);
  }

  //--------------------------------------------------------------------------
  // Private
  //--------------------------------------------------------------------------

  /**
   * Appends the renderer in the $el. To override to insert it elsewhere.
   *
   * @private
   */
  _startRenderer() {
    return; //this.renderer.appendTo(this.$el);
  }

  /**
   * Returns a title that may be displayed in the breadcrumb area.  For
   * example, the name of the record (for a form view). This is actually
   * important for the action manager: this is the way it is able to give
   * the proper titles for other actions.
   *
   * @returns {string}
   */
  getTitle() {
    return this._title;
  }

  // TODO: add hooks methods:
  // - onRestoreHook (on_reverse_breadcrumbs)

  //--------------------------------------------------------------------------
  // Private
  //--------------------------------------------------------------------------

  /**
   * @private
   * @param {string} title
   */
  _setTitle(title) {
    this._title = title;
    this.updateControlPanel({ title: this._title }, { clear: false });
  }

  /**
   * Renders the buttons to append, in most cases, to the control panel (in
   * the bottom left corner). When the action is rendered in a dialog, those
   * buttons might be moved to the dialog's footer.
   *
   * @param {jQuery Node} $node
   */
  renderButtons(node) {}

  /**
   * This is the main method to customize the controlpanel content.
   *
   * @see updateContents method in ControlPanelRenderer for more info
   *
   * @param {Object} [status]
   * @param {string} [status.title]
   * @param {Object} [options]
   * @param {boolean} [options.clear]
   */
  updateControlPanel(status, options) {
    if (this._controlPanel) {
      status = status || {};
      status.title = status.title || this.getTitle();
      this._controlPanel.updateContents(status, options || {});
    }
  }

  /**
   * In some situations, we need confirmation from the controller that the
   * current state can be destroyed without prejudice to the user.  For
   * example, if the user has edited a form, maybe we should ask him if we
   * can discard all his changes when we switch to another action.  In that
   * case, the action manager will call this method.  If the returned
   * promise is successfully resolved, then we can destroy the current action,
   * otherwise, we need to stop.
   *
   * @returns {Promise} resolved if the action can be removed, rejected
   *   otherwise
   */
  canBeRemoved() {
    return Promise.resolve();
  }
}
customElements.define('abstract-controller', AbstractController);
