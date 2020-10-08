/**
 * Whether the client is currently in "debug" mode
 *
 * @type Boolean
 */
//import { Bus } from "../store/modules/bus";
import { Registry } from '@/core/registry';
import { Translation } from '@/core/translation.js';
import { Session } from './session';

import { AjaxService } from '@/service/service_ajax';
import { NotificationService } from '@/service/service_notification';
import { SessionStorageService } from '@/service/service_session_storage';

export const service_registry = new Registry();
service_registry
  .add('ajax', AjaxService)
  .add('notification', NotificationService)
  .add('session_storage', SessionStorageService);

//var bus = new Bus();
var session = new Session(); // Session(undefined, undefined, { modules: modules, use_cors: false });
session.is_bound = session.session_bind();

export default {
  //qweb: new QWeb(config.isDebug()),

  // core classes and functions
  //bus: bus,
  session: session,
  //main_bus: new Bus(),
  _t: Translation._t,
  _lt: Translation._lt,

  // registries
  action_registry: new Registry(),
  crash_registry: new Registry(),
  serviceRegistry: service_registry,

  /**
   * @type {String}
   */
  //csrf_token: odoo.csrf_token,
};
