// for details see https://polymer-library.polymer-project.org/3.0/docs/devguide/custom-elements#defining-mixins
import { Element } from '@/mixins/element';

export class Page extends Element {
  static get properties() {
    return {
      router: Object,
      route: Object,
      path: String,
      hash: String,
    };
  }

  constructor() {
    super();
  }
}
