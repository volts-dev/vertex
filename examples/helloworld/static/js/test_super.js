class TestSuper  {
    constructor() {
         console.log("TestSuper constructor called");
        // Attach a shadow DOM to this custom element
       // this.attachShadow({ mode: 'open' });
   
    }

    connectedCallback() {
        console.log("TestSuper connected to the DOM");
    }

    disconnectedCallback() {
        console.log("TestSuper disconnected from the DOM");
    }

    adoptedCallback() {
        console.log("TestSuper adopted into a new document");
    }

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(`TestSuper attribute changed: ${name} from ${oldValue} to ${newValue}`);
    }

    static get observedAttributes() {
        return ['data-test'];
    }
}
function registerTestSuper(namespace) {
  
    var jsconstructor = function() {
      let com=  new TestSuper();
       // let com = new  HTMLElement( ) ;
       instance = Reflect.construct(HTMLElement, [], jsconstructor);
       //com.attachShadow({ mode: "open" });
        instance.__go_component_instance__= com;
         const shadow = instance.attachShadow({ mode: 'open' }); 
        shadow.innerHTML = '<div>TestSuper</div>';
        //console.log("Custom element instance created:", instance);
        return instance
    };

  classPrototype=  new Object();
    htmlElementPrototype=HTMLElement.prototype  ;
    Object.setPrototypeOf(classPrototype, htmlElementPrototype);
    jsconstructor.prototype = classPrototype;
     customElements.define(namespace, jsconstructor);
}