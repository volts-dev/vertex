// 加载服务
import { service_registry } from '@/service';
import '@/service/service_ajax';
import '@/service/service_notification';
import '@/service/service_session_storage';

// TODO　用插件替代
// rpc服务 打印服务监听 与webclient
export function ServicesListenerMixin(superClass) {
  return class extends superClass {
    constructor() {
      super();
      this.services = {}; // dict containing deployed service instances
      this.UndeployedServices = {}; // dict containing classes of undeployed services

      var self = this;
      // to properly instantiate services with this as parent, this mixin
      // assumes that it is used along the EventDispatcherMixin, and that
      // EventDispatchedMixin's init is called first
      // as EventDispatcherMixin's init is already called, this handler has
      // to be bound manually
      this.addEventListener('call_service', this._call_service.bind(this));

      // add already registered services from the service registry
      _.each(service_registry.map, function (Service, serviceName) {
        if (serviceName in self.UndeployedServices) {
          throw new Error('Service "' + serviceName + '" is already loaded.');
        }
        self.UndeployedServices[serviceName] = Service;
      });
      this._deployServices();

      // listen on newly added services
      service_registry.onAdd(function (serviceName, Service) {
        if (
          serviceName in self.services ||
          serviceName in self.UndeployedServices
        ) {
          throw new Error('Service "' + serviceName + '" is already loaded.');
        }
        self.UndeployedServices[serviceName] = Service;
        self._deployServices();
      });
    }

    //--------------------------------------------------------------------------
    // Private
    //--------------------------------------------------------------------------

    /**
     * @private
     */
    _deployServices() {
      var self = this;
      var done = false;
      while (!done) {
        var serviceName = _.findKey(this.UndeployedServices, function (
          Service
        ) {
          // no missing dependency
          return !_.some(Service.prototype.dependencies, function (depName) {
            return !self.services[depName];
          });
        });
        if (serviceName) {
          var service = new this.UndeployedServices[serviceName](this);
          this.services[serviceName] = service;
          delete this.UndeployedServices[serviceName];
          service.start();
        } else {
          done = true;
        }
      }
    }

    //--------------------------------------------------------------------------
    // Handlers
    //--------------------------------------------------------------------------

    /**
     * Call the 'service', using data from the 'event' that
     * has triggered the service call.
     *
     * For the ajax service, the arguments are extended with
     * the target so that it can call back the caller.
     *
     * @private
     * @param  {OdooEvent} event
     */
    _call_service(event) {
      var args = event.detail.args || [];
      if (event.detail.service === 'ajax' && event.detail.method === 'rpc') {
        // ajax service uses an extra 'target' argument for rpc
        args = args.concat(event.target);
      }
      var service = this.services[event.detail.service];
      var result = service[event.detail.method].apply(service, args);
      event.detail.callback(result); // 调用回调函数返回值
    }
  };
}
