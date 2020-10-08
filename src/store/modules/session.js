// 用户会话
import service from '@/service';
import { makeObservable, observable } from 'mobx';

export class Session {
  uid = null; // 用户UID
  username = null; // 用户名称
  user_context = {}; // 用户数据
  db = null; // 数据库名称
  active_id = null; // 用户激活状态
  isBound = true;

  constructor() {
    makeObservable(this, {
      isBound: observable,
    });
  }

  rpc(url, params, options) {
    return service.session.rpc(url, params, options);
  }

  is_bound() {
    return service.session.session_bind();
  }

  session_is_valid() {
    return service.session.session_is_valid();
  }
}
