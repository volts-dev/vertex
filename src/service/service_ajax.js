import Ajax from '@/core/ajax';
import store from '@/store';
import { AbstractService } from './service';

// ajax 服务提供系统的rpc 调用
export class AjaxService extends AbstractService {
  /**
   * @param {Object} libs - @see ajax.loadLibs
   * @param {Object} [context] - @see ajax.loadLibs
   * @param {Object} [tplRoute] - @see ajax.loadLibs
   */
  loadLibs(libs, context, tplRoute) {
    return Ajax.loadLibs(libs, context, tplRoute);
  }

  // 服务方法 RPC
  rpc(route, args, options, target) {
    var rpcPromise;
    var promise = new Promise(function (resolve, reject) {
      rpcPromise = store.session.rpc(route, args, options);
      rpcPromise
        .then(function (result) {
          //if (!target.isDestroyed()) {
          resolve(result);
          //}
        })
        .guardedCatch(function (reason) {
          //if (!target.isDestroyed()) {
          reject(reason);
          //}
        });
    });

    promise.abort = rpcPromise.abort.bind(rpcPromise);
    return promise;
  }
}
