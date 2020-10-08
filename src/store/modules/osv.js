/* 
    系统对象注册服务
*/
import { Registry } from "@/core/registry";
import { Translation } from "@/core/translation.js";
import { service_registry } from "@/service";
export default {
  state: {},

  _t: Translation._t,
  _lt: Translation._lt,
  form_custom_registry: new Registry(),
  form_tag_registry: new Registry(),
  form_widget_registry: new Registry(),
  list_widget_registry: new Registry(),
  one2many_view_registry: new Registry(),
  search_filters_registry: new Registry(),
  search_widgets_registry: new Registry(),
  // registries
  view_registry: new Registry(),
  action_registry: new Registry(),
  crash_registry: new Registry(),
  serviceRegistry: service_registry,
};
