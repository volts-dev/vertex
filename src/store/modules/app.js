import Cookies from '@/libs/js.cookie.min.mjs';
import { makeObservable, observable } from 'mobx';

export class App {
  apps = []; // 所有程序
  isHomePge = false;
  activeApp = undefined; // 当前程序
  activeAppId = undefined; // 当前程序Id

  menus = []; // 服务器返回的 raw 数据
  menusMap = []; // 经过字典索引

  sidebar = {
    opened: Cookies.get('sidebarStatus')
      ? !!+Cookies.get('sidebarStatus')
      : true,
    withoutAnimation: false,
  };
  device = 'desktop';
  size = Cookies.get('size') || 'medium';

  constructor() {
    makeObservable(this, {
      activeAppId: observable,
    });
  }

  GetMenus() {
    return this.menus;
  }

  getMenuById(id) {
    return this.menusMap[id];
  }

  // 废弃
  GetSubMenus(id) {
    var data = [];
    for (var i = 0; i < this.menus.length; i++) {
      var m = this.menus[i];
      if (m.parent_id === id) {
        data.push(m);
      }
    }

    return data;
  }

  UpdateRootMenu(id) {
    this.app = this.menusMap[id];
  }

  // 更新菜单
  UpdateMenus(data) {
    this.menus = data;
    this.menusMap = {}; //清空
    // 备份索引
    for (var i = 0; i < data.length; i++) {
      var menu = data[i];
      if (
        menu.parent_id === null ||
        menu.parent_id === '' ||
        menu.parent_id === 0
      ) {
        // 获取APP基本信息
        var app = {
          id: menu.id,
          name: menu.name,
          menus: genMenus(menu.id, menu.id, data),
        };

        this.apps[menu.id] = app;

        // 跟新激活App
        if (menu.id === this.activeAppId) {
          this.activeApp = app;
        }
      }
      this.menusMap[data[i].id] = data[i];
    }
  }

  // 设置激活的App
  setActiveApp(id) {
    var app = this.apps[id];
    if (id && !app) {
      var menu = this.getMenuById(id);
      if (menu && menu.app) {
        app = this.apps[menu.app];
      } else {
        return;
      }
    }
    this.activeApp = app;
    this.activeAppId = app ? app.id : null;
  }

  // 获取当前激活的程序状态
  getActiveApp() {
    var app = this.activeApp;
    if (!app) {
      if (this.activeAppId) {
        return this.apps[this.activeAppId];
      }

      return null;
    }

    return app;
  }
}

function genMenus(app, id, data) {
  var ms = [];
  for (var i = 0; i < data.length; i++) {
    var m = data[i];
    if (m.parent_id === id) {
      m.subMenus = genMenus(app, m.id, data);
      m.app = app;
      ms.push(m);
    }
  }
  return ms;
}
