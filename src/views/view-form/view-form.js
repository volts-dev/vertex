import { html, css } from 'lit-element';
//import '../../element/input-editor/input-editor.js';
//import './form-header/form-header.js';
//import "./form-sheet/form-sheet.js";
//import "./form-group/form-group.js";
//import "./form-tabs/form-tabs.js";
import './form-input/form-input.js';
import './form-button/form-button.js';
import './form-selection/form-selection.js';

//import './view-form-icon.js';
import './view-form-buttons.js';
import { View } from '../view.js';
import { formStyles } from './view-form-styles.js';

function isEmpty(map) {
  for (var key in map) {
    if (map.hasOwnProperty(key)) {
      return false;
    }
  }
  return true;
}

export class ViewForm extends View {
  static get styles() {
    return [
      formStyles,
      css`
        :host {
          position: relative;
          display: block;
          padding-bottom: 50px;
        }

        ::slotted(form-header, form-sheet) {
          position: relative;
          display: block;
        }

        ::slotted(form-header) {
          border-bottom: 1px solid #cacaca;
          padding-left: 2px;
          padding-left: 14px;
        }

        ::slotted(form-sheet) {
          margin-right: auto;
          margin-left: auto;
          padding-top: 24px;
          padding-right: 16px;
          padding-bottom: 24px;
          padding-left: 16px;
          position: relative;
          background-color: #8f8f8f;
        }

        ::slotted(.title) {
          padding-left: 95px;
        }

        ::slotted(.btn) {
          display: inline-block;
          margin-bottom: 0;
          font-weight: 400;
          text-align: center;
          vertical-align: middle;
          touch-action: manipulation;
          cursor: pointer;
          background-image: none;
          border: 1px solid transparent;
          white-space: nowrap;
          padding: 6px 12px;
          font-size: 14px;
          line-height: 1.42857143;
          border-radius: 4px;
          -webkit-user-select: none;
          -moz-user-select: none;
          -ms-user-select: none;
          user-select: none;
        }

        ::slotted(.button_box) {
          background-color: #e2e2e0;
          margin: -24px -16px 24px -16px;
          text-align: right;
        }

        ::slotted(.stat_button) {
          font-weight: normal;
          width: 132px !important;
          height: 40px;
          color: #666;
          margin: 0px -1px -1px 0px;
          padding: 0;
          border-radius: 0;
          box-shadow: none;
          background-color: transparent;

          border-radius: 0px;
          border: none;
          border-left: 1px solid #8f8f8f;
        }
      `,
    ];
  }

  static get properties() {
    return {
      arch: { type: Object },
      inited: { type: Boolean },
      editMode: { type: Boolean }, // 修改模式
      name: { type: String }, // 名称用于标题
      node_name: { type: Object }, // 导航名称节点
      data: { type: Object, notify: true },
      //-----------
      // actual_mode create/edit
      modelName: { type: String, value: 'view', notify: true },

      // 修改前的数据
      OrgData: { type: Object },

      // 修改过的字段  TODO改变量名称
      fieldsList: { type: Array },
      params: { type: Object, notify: true },
    };
  }

  render() {
    return html` ${this.renderForm} `;
  }

  constructor(...args) {
    super(...args);
    this.searchable = false;
    this.innerGroupColspan = 2;
    this.outerGroupColspan = 2;
    this.idsForLabels = {};
  }

  firstUpdated() {
    super.firstUpdated();

    this.editMode = false;

    // 监听输入控件事件
    this.addEventListener('show', function (e) {});
    this.addEventListener(
      'on-value-changed',
      this.onInputValueChanged.bind(this)
    );
  }

  show(mgr) {
    super.show();

    if (this.app.router.Get('id')) {
      //根据lViewMgr组织查询参数 {"view_id":[[viewId]],"model=":"[[action.res_model]]","view_type":"[[mode]]"}
      this.params = {};
      this.params.id = this.app.router.Get('id');
      this.params.fields = Object.keys(this.controlPanel.fields); //lViewMgr.view.fields.keys();
      this.params.context = {};

      // 执行API
      this.modelName = this.controlPanel.action.res_model;
      this.datasource.action = '/dataset/call_kw/' + this.modelName + '/read';
      this.datasource.params = this.params;
      this.datasource.read(0).then(this.onDataChanged.bind(this));
      //  console.log(e.detail.kicked); // true
    } else {
      // # 当ID未被指定时 为[创建]记录状态 无需获取数据
      //this.updateData(undefined);
      //this.setMode("edit");
    }
  }

  renderHeader(parent) {
    // 添加导航标题
    if (parent.children.length == 0) {
      var li = document.createElement('li');
      li.classList.add('breadcrumb-item');
      var a = document.createElement('a');
      a.innerText = this.controlPanel.action.name;
      a.addEventListener('tap', this.onBreadcrumbTap.bind(this));
      li.appendChild(a);
      parent.appendChild(li);

      this.node_name = document.createElement('li');
      //li.classList.add('breadcrumb');
      if (this.name) {
        // this.onNameChanged(this.name);
        this.node_name.innerText = this.name;
      }
      parent.appendChild(this.node_name);
    }
  }

  // 添加菜单按钮到CtrlPanel
  renderButtons(parent) {
    //buttons.innerHTML=document.createElement("view-form-buttons").outerHTML;
    // 清空节点
    /*  while (buttons.children.length > 0) {
              buttons.removeChild(buttons.firstChild);
          }
*/
    this.buttons = document.createElement('view-form-buttons');
    this.buttons.action_buttons = true;
    parent.appendChild(this.buttons);
    this.buttons.addEventListener(
      'clickButton',
      function (event) {
        ev.stopPropagation(); // Prevent x2m lines to be auto-saved

        var action = event.detail;
        switch (action) {
          case 'edit':
            this.fieldsList = [];
            this.setMode('edit');
            break;
          case 'create':
            this.createRecord();
            break;
          case 'save':
            this.saveRecord();
            this.setMode('readonly');
            break;
          case 'discard':
            this.setMode('readonly');
            break;
        }
      }.bind(this)
    );
  }

  renderSidebar(bar) {
    // TODO
  }

  on_invalid() {
    /* var warnings = _(this.fields).chain()
             .filter(function (f) { return !f.isValid(); })
             .map(function (f) {
                 return _.str.sprintf('<li>%s</li>',
                     _.escape(f.string));
             }).value();
         warnings.unshift('<ul>');
         warnings.push('</ul>');
         this.do_warn(_t("The following fields are invalid:"), warnings.join(''));
         */
  }

  // 执行保存
  processSave(save_obj) {
    var self = this;
    //try {
    var form_invalid = false, // 记录Form是否有效
      values = {},
      readonly_values = {},
      first_invalid_field = null, // 第一个无效值的Input
      deferred = [];

    this.inputs.forEach(input => {
      if (input.isValid && !input.isValid()) {
        form_invalid = true;
        if (!first_invalid_field) {
          first_invalid_field = f;
        }
      } else if (input.name !== 'id') {
        //&& (!self.data)&&(!self.data.id)
        // Special case 'id' field, do not save this field
        // on 'create' : save all non readonly fields
        // on 'edit' : save non readonly modified fields
        if (!input.getAttribute('readonly')) {
          // 获取字段是否只读
          if (input.field && self.fieldsList.indexOf(input.field) > -1) {
            // 查看Field是否在以修改的列表里
            values[input.name] = input.get_value(true);
          }
        } else {
          readonly_values[input.name] = input.get_value(true);
        }
      }
    });

    // 初始化Data记录
    if (!this.data) {
      this.data = {};
    }

    // Heuristic to assign a proper sequence number for new records that
    // are added in a dataset containing other lines with existing sequence numbers
    /*if (!self.data.id && self.fields.sequence &&
            !_.has(values, 'sequence') && !_.isEmpty(self.dataset.cache)) {
            // Find current max or min sequence (editable top/bottom)
            var current = _[prepend_on_create ? "min" : "max"](
                _.map(self.dataset.cache, function(o){return o.values.sequence})
            );
            values['sequence'] = prepend_on_create ? current - 1 : current + 1;
        }*/

    if (form_invalid) {
      //self.set({'display_invalid_fields': true});
      first_invalid_field.focus();
      self.on_invalid();
      //def_process_save.reject();
    } else {
      //self.set({'display_invalid_fields': false});
      var save_deferral;
      if (!self.data.id) {
        // 创建保存 Creation save
        self.datasource.action =
          '/dataset/call_kw/' + this.modelName + '/create';
        save_deferral = self.datasource.create(values, {
          readonly_fields: readonly_values,
        }); //.then(function(r) {
        //     return self.record_created(r, prepend_on_create);
        // }, null);
      } else if (isEmpty(values)) {
        // 无修改不保存 Not dirty, noop save
        save_deferral = new Promise();
      } else {
        // 保存修改 Write save
        self.datasource.action =
          '/dataset/call_kw/' + this.modelName + '/write';
        save_deferral = self.datasource.write(self.data.id, values, {
          readonly_fields: readonly_values,
        }); //.then(function(r) {
        //      return self.record_saved(r);
        //  }, null);
      }
      save_deferral
        .then(function (result) {
          //         def_process_save.resolve(result);
        })
        .catch(function () {
          //        def_process_save.reject();
        });
    }

    // } catch (e) {
    //     console.error(e);
    //  return def_process_save.reject();
    //  }
    //    return def_process_save;

    /*
              var self = this;
              var prepend_on_create = save_obj.prepend_on_create;
              var def_process_save = $.Deferred();
              try {
                  var form_invalid = false,
                      values = {},
                      first_invalid_field = null,
                      readonly_values = {},
                      deferred = [];

                  _.each(self.fields, function (f) {
                      var res = f.before_save();
                      if (res) {
                          deferred.push(res);
                          res.fail(function () {
                              form_invalid = true;
                              if (!first_invalid_field) {
                                  first_invalid_field = f;
                              }
                          });
                      }
                  });

                  $.when.apply($, deferred).always(function () {

                      _.each(self.fields, function (f) {
                          if (!f.isValid()) {
                              form_invalid = true;
                              if (!first_invalid_field) {
                                  first_invalid_field = f;
                              }
                          } else if (f.name !== 'id' && (!self.datarecord.id || f._dirty_flag)) {
                              // Special case 'id' field, do not save this field
                              // on 'create' : save all non readonly fields
                              // on 'edit' : save non readonly modified fields
                              if (!f.get("readonly")) {
                                  values[f.name] = f.get_value(true);
                              } else {
                                  readonly_values[f.name] = f.get_value(true);
                              }
                          }

                      });

                      // Heuristic to assign a proper sequence number for new records that
                      // are added in a dataset containing other lines with existing sequence numbers
                      if (!self.datarecord.id && self.fields.sequence &&
                          !_.has(values, 'sequence') && !_.isEmpty(self.dataset.cache)) {
                          // Find current max or min sequence (editable top/bottom)
                          var current = _[prepend_on_create ? "min" : "max"](
                              _.map(self.dataset.cache, function(o){return o.values.sequence})
                          );
                          values['sequence'] = prepend_on_create ? current - 1 : current + 1;
                      }
                      if (form_invalid) {
                          self.set({'display_invalid_fields': true});
                          first_invalid_field.focus();
                          self.on_invalid();
                          def_process_save.reject();
                      } else {
                          self.set({'display_invalid_fields': false});
                          var save_deferral;
                          if (!self.datarecord.id) {
                              // Creation save
                              save_deferral = self.dataset.create(values, {readonly_fields: readonly_values}).then(function(r) {
                                  return self.record_created(r, prepend_on_create);
                              }, null);
                          } else if (_.isEmpty(values)) {
                              // Not dirty, noop save
                              save_deferral = $.Deferred().resolve({}).promise();
                          } else {
                              // Write save
                              save_deferral = self.dataset.write(self.datarecord.id, values, {readonly_fields: readonly_values}).then(function(r) {
                                  return self.record_saved(r);
                              }, null);
                          }
                          save_deferral.then(function(result) {
                              def_process_save.resolve(result);
                          }).fail(function() {
                              def_process_save.reject();
                          });
                      }
                  });
              } catch (e) {
                  console.error(e);
                  return def_process_save.reject();
              }
              return def_process_save;
              */
  }

  processOperations() {
    this.processSave();
    /*   return this.mutating_mutex.exec(function() {
          function iterate() {

              var mutex = new utils.Mutex();
              _.each(self.fields, function(field) {
                  self.onchanges_mutex.def.then(function(){
                      mutex.exec(function(){
                          return field.commit_value();
                      });
                  });
              });

              return mutex.def.then(function () { return self.onchanges_mutex.def; }).then(function() {
                  var save_obj = self.save_list.pop();
                  if (save_obj) {
                      return self.processSave(save_obj).then(function() {
                          save_obj.ret = _.toArray(arguments);
                          return iterate();
                      }, function() {
                          save_obj.error = true;
                      });
                  }
                  return $.when();
              }).fail(function() {
                  self.save_list.pop();
                  return $.when();
              });
          }
          return iterate();
      });*/
  }

  /**
   * Triggers saving the form's record. Chooses between creating a new
   * record or saving an existing one depending on whether the record
   * already has an id property.
   *
   * @param {Boolean} [prepend_on_create=false] if ``save`` creates a new
   * record, should that record be inserted at the start of the dataset (by
   * default, records are added at the end)
   */
  saveRecord(prepend_on_create) {
    this.processOperations();
    /*      var self = this;
              var save_obj = {prepend_on_create: prepend_on_create, ret: null};
              this.save_list.push(save_obj);
              return self.processOperations().then(function() {
                  if (save_obj.error)
                      return $.Deferred().reject();
                  return $.when.apply($, save_obj.ret);
              }).done(function(result) {
                  self.$el.removeClass('oe_form_dirty');
              });*/
  }

  onDataChanged(data) {
    var self = this;

    this.inputs.forEach(input => {
      input.mode = this.mode;
      if (data) {
        var fieldName = input.getAttribute('name');
        var value = data[fieldName];

        if (fieldName == 'name') {
          this.name = value;
        }

        switch (input.type) {
          case 'selection':
          case 'many2one':
          case 'one2one':
          case 'many2many':
            // for relation selection
            if (value && value.length == 2) {
              input.value = [{ id: parseInt(value[0]), name: value[1] }];
            } else {
              input.value = [{ id: parseInt(value), name: value }];
            }
            break;
          default:
            input.value = value;
        }
      } else {
        input.value = null;
      }
    });

    // 赋值
    this.data = data;
    // 确保字段存在
    /*
    if (!this.fields) {
      //this.fields = Polymer.dom(this).querySelectorAll('*[field]:not(field)');
      // 使用  Array.prototype.slice.call(); 转换到Arry
      this.fields = Array.prototype.slice.call(document.querySelectorAll("*[name]:not(field)"));
    }*/

    // 赋值每个字段元件
    switch (this.mode) {
      case 'create':
        // 返回视图模式
        this.setMode('readonly');

        break;
      case 'read':
        break;
      case 'write':
        // 返回视图模式
        this.setMode('readonly');
        break;
      case 'unlink':
        // 返回列表
        break;
    }
  }

  // 更新Data数据到控件
  updateData(data) {
    /* if (!this.fields) {
      //this.fields = Polymer.dom(this).querySelectorAll('*[field]:not(field)');
      this.fields = Array.prototype.slice.call(document.querySelectorAll("*[name]:not(field)"));
    }*/

    this.inputs.forEach(input => {
      if (data) {
        input.value = data[input.field.name];
      } else {
        input.value = '';
      }
    });
  }

  // 更新所有Imput Mode
  actualizeInputModel(switch_to) {
    //跟新Mode
    this.inputs.forEach(input => {
      input.mode = switch_to;
    });
  }
  /**
   * Ask the view to switch to a precise mode if possible. The view is free to
   * not respect this command if the state of the dataset is not compatible with
   * the new mode. For example, it is not possible to switch to edit mode if
   * the current record is not yet saved in database.
   *
   * @param {string} [new_mode] Can be "edit", "view", "create" or undefined. If
   * undefined the view will test the actual mode to check if it is still consistent
   * with the dataset state.
   */
  actualizeFormMode(switch_to) {
    var mode = switch_to || this.mode;
    /*
    if (!this.data || !this.data.id) {
      // 当Id为空时 数据为新添加的
      mode = "create";
    } else if (mode === "create") {
      mode = "edit";
    }*/
    this.render_value_defs = [];
    this.mode = mode;
  }

  /**
   * Ask the view to switch to edit mode if possible. The view may not do it
   * if the current record is not yet saved. It will then stay in create mode.
   */

  setMode(mode) {
    this.actualizeFormMode(mode);
    this.actualizeInputModel(mode);
  }

  /**
   * This method switches the form view in edit mode, with a new record.
   *
   * @todo make record creation a basic controller feature
   * @param {string} [parentID] if given, the parentID will be used as parent
   *                            for the new record.
   * @returns {Deferred}
   */
  createRecord(parentID) {
    this.fieldsList = [];
    this.datasource.index = undefined;

    if (this.data === undefined) {
      // null index means we should start a new record
      //return self.on_button_new();
    }

    // # 先转为修改模式 再设置Data为空
    // this.actualizeInputModel
    /*     return $.when(this.has_been_loaded).then(function() {
                 return self.can_be_discarded().then(function() {
                     return self.load_defaults();
                 });
             });*/

    //this.actualizeInputModel("edit") // 转换
    this.data = undefined;
    this.updateData(undefined);
    //this.actualizeFormMode("edit");
    this.setMode('edit');
  }

  ___onCreateRecord(event) {
    this.createRecord();
  }

  ___onEditRecord(event) {
    this.fieldsList = [];
    this.setMode('edit');
  }

  ___onSaveRecord(e) {
    var self = this;
    //  e.target.setAttribute("disabled", true);
    this.save();
    this.setMode('readonly');

    /*  return this.save().done(function(result) {
              self.trigger("save", result);
              self.reload().then(function() {
                  self.toViewMode();
                  core.bus.trigger('do_reload_needaction');
                  core.bus.trigger('form_view_saved', self);
              });
          }).always(function(){
              $(e.target).attr("disabled", false);
          });*/

    //TODO
    //this.fieldsList=[];
  }

  ___onDiscardRecord(event) {
    var self = this;
    //  this.can_be_discarded().then(function() {
    if (self.getAttribute('mode') === 'create') {
      //  self.trigger('history_back');
    } else {
      this.setMode('readonly');

      //  $.when.apply(null, self.render_value_defs).then(function(){
      //      self.trigger('load_record', self.datarecord);
      //  });
    }
    //  });
    //  this.trigger('onCancelButtonClick');*/
    return false;
  }

  // banding按钮事件
  ___on_buttons_changed(e) {
    //lv=Polymer.dom(e.currentTarget).querySelector("div");
    //lv=e.currentTarget.querySelector(".form_button_create");
  }

  ___buildQueryString() {
    var out = [];
    if (this.menu) out.push('menu=' + this.menu);
    if (this.model && this.model.length) out.push('model=' + this.model);
    if (this.action && this.action.id) out.push('action=' + this.action.id);
    if (this.action) {
      if (!this.view) {
        out.push('view=' + this.action.view_mode.split(',')[0]);
      } else {
        out.push('view=' + this.view);
      }
    }
    return out.join('&');
  }

  // 导航到
  onBreadcrumbTap(e) {
    if (this.controlPanel) {
      this.controlPanel.SetViewMode('tree'); // form 返回通常为Tree
      /*if (this.controlPanel.lastView == "form") {
                 this.controlPanel.SetViewMode("tree");
             } else {
                 this.controlPanel.SetViewMode(this.controlPanel.lastView);
             }*/
    }
  }

  get name() {
    return this._name;
  }

  // 更新Name 当Form显示时显示
  set name(name) {
    let oldValue = this._name;
    this._name = name;
    this.requestUpdate('name', oldValue);
    this.node_name.innerText = this.name;
  }

  get mode() {
    return this._mode;
  }

  set mode(value) {
    let oldValue = this._mode;
    this._mode = value;
    this.requestUpdate('mode', oldValue);
    if (this.buttons) {
      this.buttons.mode = this.mode;
    }
  }

  onInputValueChanged(e) {
    //TODO 保存模式可用
    if (this.mode == 'edit') {
      if (!this.fieldsList) {
        this.fieldsList = [];
      }

      if (this.fieldsList.indexOf(e.detail.field) == -1) {
        this.fieldsList.push(e.detail.field);
      }
    }
  }

  capitalize = s => {
    if (typeof s !== 'string') return '';
    return s.charAt(0).toUpperCase() + s.slice(1);
  };

  get renderForm() {
    if (!this.arch) {
      var form = document.createElement('div');
      // TODO form.classList.add(this.classList);
      Array.from(this.children).forEach(child => {
        form.append(this.renderNode(child));
      });

      if (this.hasSheet) {
        form.classList.add('form_sheet_bg');
      }

      // 备份模板
      this.arch = form;
    }
    this.toggleClass(this, 'form_editable', this.mode === 'edit');
    this.toggleClass(this, 'form_readonly', this.mode === 'readonly');

    return html` ${this.arch} `;
  }

  addOnClickAction($el, node) {
    var self = this;
    $el.click(function () {
      /* self.trigger_up("button_clicked", {
        attrs: node.attrs,
        record: self.state
      });*/
    });
  }

  // 渲染节点
  renderNode(node) {
    var rendererName =
      'renderTag' + this.capitalize(node.tagName.toLowerCase());
    var renderer = this[rendererName];
    if (renderer) {
      return renderer.call(this, node);
    }
    if (node.tag === 'div' && node.getAttribute('name') === 'button_box') {
      return this.renderButtonBox(node);
    }

    if (typeof node === 'string' || node instanceof String) {
      return node;
    }
    return this.renderGenericTag(node);
  }

  renderButtonBox(node) {
    var self = this;
    var result = document.createElement(node.tag);
    result.classList.add('not_full');
    var buttons = _.map(node.children, function (child) {
      if (child.tagName.toLowerCase() === 'button') {
        return self.renderStatButton(child);
      } else {
        return self.renderNode(child);
      }
    });
    var buttons_partition = _.partition(buttons, function (button) {
      if (button.classList.contains('invisible_modifier')) {
        return button;
      }
    });
    var invisible_buttons = buttons_partition[0];
    var visible_buttons = buttons_partition[1];

    // Get the unfolded buttons according to window size
    var nb_buttons = [2, 2, 4, 6][config.device.size_class] || 7;
    var unfolded_buttons = visible_buttons
      .slice(0, nb_buttons)
      .concat(invisible_buttons);

    // Get the folded buttons
    var folded_buttons = visible_buttons.slice(nb_buttons);
    if (folded_buttons.length === 1) {
      unfolded_buttons = buttons;
      folded_buttons = [];
    }

    // Toggle class to tell if the button box is full (CSS requirement)
    var full = visible_buttons.length > nb_buttons;
    this.toggleClass(result, 'full', full);
    this.toggleClass(result, 'not_full', !full);

    // Add the unfolded buttons
    _.each(unfolded_buttons, function (button) {
      button.appendTo(result);
    });

    // Add the dropdown with folded buttons if any
    if (folded_buttons.length) {
      result.append(
        this.renderButton({
          attrs: {
            class: 'stat_button button_more dropdown-toggle',
            'data-toggle': 'dropdown',
          },
          text: _t('More'),
        })
      );
      var dropdown = document.createElement('div');
      dropdown.classList.value = 'dropdown-menu dropdown_more';
      dropdown.setAttribute('role', 'menu');
      _.each(folded_buttons, function (button) {
        button.classList.add('dropdown-item');
        dropdown.append(button);
      });
      result.append(dropdown);
    }

    this.handleAttributes($result, node);
    //this._registerModifiers(node, this.state, $result);
    return result;
  }

  renderStatButton(node) {
    var button = this.renderButtonFromNode(node, {
      extraClass: 'stat_button',
    });

    _.map(node.children, function (child) {
      button.append(this.renderNode(child));
    });

    if (node.getAttribute('help')) {
      this.addButtonTooltip(node, button);
    }
    this.addOnClickAction(button, node);
    this.handleAttributes(button, node);
    //this._registerModifiers(node, this.state, $button);
    return button;
  }

  addButtonTooltip(node, button) {
    var self = this;
    /*button.tooltip({
      title: function () {
          return qweb.render('WidgetButton.tooltip', {
              debug: config.debug,
              state: self.state,
              node: node,
          });
      },
  });*/
  }

  // 渲染范节点元素
  renderGenericTag(node) {
    var result = document.createElement(node.tagName.toLowerCase());
    var r = node.attributes;
    for (var i = 0; i < r.length; i++) {
      if (r[i].name != 'modifiers') {
        result.setAttribute(r[i].name, r[i].value);
      }
    }

    this.handleAttributes(result, node);
    //this._registerModifiers(node, this.state, result);
    Array.from(node.children).forEach(child => {
      result.append(this.renderNode(child));
    });
    return result;
  }

  renderHeaderButton(node) {
    /*
    var $button = this._renderButtonFromNode(node);

    // Current API of odoo for rendering buttons is "if classes are given
    // use those on top of the 'btn' and 'btn-{size}' classes, otherwise act
    // as if 'btn-secondary' class was given". The problem is that, for
    // header buttons only, we allowed users to only indicate their custom
    // classes without having to explicitely ask for the 'btn-secondary'
    // class to be added. We force it so here when no bootstrap btn type
    // class is found.
    if ($button.not('.btn-primary, .btn-secondary, .btn-link, .btn-success, .btn-info, .btn-warning, .btn-danger').length) {
        $button.addClass('btn-secondary');
    }

    this._addOnClickAction($button, node);
    this._handleAttributes($button, node);
    this._registerModifiers(node, this.state, $button);

    // Display tooltip
    if (config.debug || node.attrs.help) {
        this._addButtonTooltip(node, $button);
    }*/
    var button = document.createElement('button');
    return button;
  }

  renderHeaderButtons(node) {
    var self = this;
    var buttons = document.createElement('div');
    buttons.classList.add('statusbar_buttons');
    Array.from(node.children).forEach(child => {
      if (child.tag === 'button') {
        buttons.append(self.renderHeaderButton(child));
      }
      if (child.tag === 'widget') {
        //buttons.append(self.renderTagWidget(child));
      }
    });

    return buttons;
  }

  renderTagButton(node) {
    var button = this.renderButtonFromNode(node);
    Array.from(node.children).forEach(child => {
      button.append(this.renderNode(child));
    });
    //this.addOnClickAction(button, node);
    this.handleAttributes(button, node);
    //this._registerModifiers(node, this.state, $button);

    // Display tooltip
    //if (config.debug || node.attrs.help) {
    //    this._addButtonTooltip(node, $button);
    //}

    return button;
  }

  renderTagHeader(node) {
    var self = this;
    var statusbar = document.createElement('div');
    statusbar.classList.add('form_statusbar');
    statusbar.append(this.renderHeaderButtons(node));
    Array.from(node.children).forEach(child => {
      if (child.tagName.toLowerCase() === 'field') {
        var el = self.renderFieldWidget(child, self.state);
        statusbar.append(el);
      }
    });
    this.handleAttributes(statusbar, node);
    //this.registerModifiers(node, this.state, $statusbar);
    return statusbar;
  }
  renderTagSheet(node) {
    this.hasSheet = true;
    var sheet = document.createElement('div');
    sheet.classList.add('clearfix', 'form_sheet');
    Array.from(node.children).forEach(child => {
      sheet.append(this.renderNode(child));
    });
    return sheet;
  }

  renderTagGroup(node) {
    var isOuterGroup = _.some(node.children, function (child) {
      return child.tagName.toLowerCase() === 'group';
    });
    if (!isOuterGroup) {
      return this.renderInnerGroup(node);
    }
    return this.renderOuterGroup(node);
  }

  renderTagTabs(node) {
    var self = this;
    var headers = document.createElement('ul');
    headers.classList.value = 'nav nav-tabs';
    var pages = document.createElement('div');
    pages.classList.value = 'tab-content';

    var autofocusTab = -1;
    // renderedTabs is used to aggregate the generated $headers and $pages
    // alongside their node, so that their modifiers can be registered once
    // all tabs have been rendered, to ensure that the first visible tab
    // is correctly activated
    var renderedTabs = _.map(node.children, function (child, index) {
      var pageID = _.uniqueId('tabs_page_');
      var header = self.renderTabHeader(child, pageID);
      var page = self.renderTabPage(child, pageID);
      if (
        autofocusTab === -1 &&
        child.getAttribute('autofocus') === 'autofocus'
      ) {
        autofocusTab = index;
      }
      self.handleAttributes(header, child);
      headers.append(header);
      pages.append(page);
      return {
        header: header,
        page: page,
        node: child,
      };
    });
    if (renderedTabs.length) {
      var tabToFocus = renderedTabs[Math.max(0, autofocusTab)];
      tabToFocus.header.querySelector('.nav-link').classList.add('active');
      tabToFocus.page.classList.add('active');
    }
    // register the modifiers for each tab
    _.each(renderedTabs, function (tab) {
      tab.header.addEventListener('click', function (e) {
        if (!tab.header.classList.contains('active')) {
          var currenTab = headers.querySelector('a.active');
          if (currenTab != tab.header.firstChild) {
            var currenPage = pages.querySelector('div.active');

            //var firstVisibleTab = headers.querySelector("li:not(.o_invisible_modifier):first() > a");
            tab.header.firstChild.classList.add('active');
            pages
              .querySelector(tab.header.firstChild.getAttribute('href'))
              .classList.add('active');

            if (currenTab) {
              currenTab.classList.remove('active');
            }
            if (currenPage) {
              currenPage.classList.remove('active');
            }
          }
        }
      });
      /*
        self.registerModifiers(tab.node, self.state, tab.header, {
            callback: function (element, modifiers) {
                // if the active tab is invisible, activate the first visible tab instead
                var link = element.querySelector('.nav-link');
                if (modifiers.invisible && link.classList.contains('active')) {
                    link.classList.remove('active');
                    tab.page.classList.remove('active');
                    var firstVisibleTab = headers.querySelector('li:not(.o_invisible_modifier):first() > a');
                    firstVisibleTab.classList.add('active');
                    pages.querySelector(firstVisibleTab.getAttribute('href')).classList.add('active');
                }
            },
        });
        */
    });
    var tabs = document.createElement('div');
    tabs.classList.add('tabs');
    tabs.append(headers, pages);
    //var $notebook = $('<div class="o_notebook">')
    //        .data('name', node.attrs.name || '_default_')
    //        .append($headers, $pages);
    //this._registerModifiers(node, this.state, $notebook);
    this.handleAttributes(tabs, node);
    return tabs;
  }

  renderTabHeader(page, page_id) {
    var a = document.createElement('a');
    a.setAttribute('data-toggle', 'tab');
    a.setAttribute('disable_anchor', 'true');
    a.setAttribute('href', '#' + page_id);
    a.setAttribute('role', 'tab');
    a.classList.add('nav-link');
    a.innerText = page.getAttribute('string');
    var li = document.createElement('li');
    li.classList.add('nav-item');
    li.append(a);
    return li;
  }

  renderTabPage(page, page_id) {
    var self = this;
    var result = document.createElement('div');
    result.classList.add('tab-pane');
    result.setAttribute('id', page_id);
    _.map(page.children, function (child) {
      result.append(self.renderNode(child));
    });

    return result;
  }

  renderTagField(node) {
    return this.renderFieldWidget(node, this.state);
  }

  /**
   * Renders a 'group' node, which contains 'group' nodes in its children.
   */
  renderOuterGroup(node) {
    var self = this;
    var result = document.createElement('div');
    result.classList.add('group');
    var nbCols =
      parseInt(this.getAttribute('col'), 10) || this.outerGroupColspan;
    var colSize = Math.max(1, Math.round(12 / nbCols));
    if (node.getAttribute('string')) {
      var sep = document.createElement('div');
      sep.classList.add('horizontal_separator');
      sep.innerText = node.getAttribute('string');
      result.append(sep);
    }

    _.map(node.children, function (child) {
      if (child.tagName.toLowerCase() === 'newline') {
        result.append(document.createElement('br'));
      }
      var node = self.renderNode(child);
      // 计算Group宽度值1~12
      node.classList.add(
        'group_col_' +
          colSize * (parseInt(child.getAttribute('colspan'), 10) || 1)
      );
      result.append(node);
    });

    this.handleAttributes(result, node);
    // this._registerModifiers(node, this.state, $result);
    return result;
  }

  renderInnerGroup(node) {
    var self = this;
    var result = document.createElement('table');
    result.classList.value = 'group inner_group';
    this.handleAttributes(result, node);
    //this._registerModifiers(node, this.state, result);
    this.innerGroupColspan =
      parseInt(this.getAttribute('col'), 10) || this.innerGroupColspan;
    var title = node.getAttribute('string');
    if (title) {
      var tr = document.createElement('tr');
      var td = document.createElement('td');
      td.setAttribute('colspan', this.innerGroupColspan);
      td.style.value = 'width: 100%;';
      var div = document.createElement('div');
      div.classList.add('horizontal_separator');
      div.innerText = title;
      td.append(div);
      tr.append(td);
      result.append(tr);
      /*
      var sep = html`
        <tr>
          <td colspan="${this.innerGroupColspan}" style="width: 100%;">
            <div class="horizontal_separator">${title}</div>
          </td>
        </tr>
      `;
      const template = document.createElement("template");
      template.innerHTML = sep.getHTML();
      result.append(template.content);
      */
    }

    var rows = [];
    var currentRow = document.createElement('tr');
    var currentColspan = 0;
    _.each(node.children, function (child) {
      if (child.tagName.toLowerCase() === 'newline') {
        rows.push(currentRow);
        currentRow = document.createElement('tr');
        currentColspan = 0;
        return;
      }

      var colspan = parseInt(child.getAttribute('colspan'), 10);
      var isLabeledField =
        child.tagName.toLowerCase() === 'field' &&
        child.getAttribute('nolabel') !== '1';
      if (!colspan) {
        if (isLabeledField) {
          colspan = 2;
        } else {
          colspan = 1;
        }
      }
      var finalColspan = colspan - (isLabeledField ? 1 : 0);
      currentColspan += colspan;

      if (currentColspan > self.innerGroupColspan) {
        rows.push(currentRow);
        currentRow = document.createElement('tr');
        currentColspan = colspan;
      }

      var tds;
      if (child.tagName.toLowerCase() === 'field') {
        tds = self.renderInnerGroupField(child);
      } else if (child.tagName.toLowerCase() === 'label') {
        tds = self.renderInnerGroupLabel(child);
      } else {
        tds = document.createElement('tr');
        tds.append(self.renderNode(child));
      }
      if (finalColspan > 1) {
        tds.last().setAttribute('colspan', finalColspan);
      }
      currentRow.append(tds);
    });
    rows.push(currentRow);

    _.each(rows, function (tr) {
      var lbs = tr.querySelectorAll('.td_label');
      var nonLabelColSize = 100 / (self.innerGroupColspan - lbs.length);
      var lbs = tr.querySelectorAll(':not(.td_label)');
      _.each(lbs, function (el) {
        el.style.width =
          (parseInt(el.getAttribute('colspan'), 10) || 1) * nonLabelColSize +
          '%';
      });
      result.append(tr);
    });

    return result;
  }

  renderInnerGroupField(node) {
    var el = this.renderFieldWidget(node, this.state);
    var tds = document.createElement('td');
    tds.append(el);

    if (node.getAttribute('nolabel') !== '1') {
      var labelTd = this.renderInnerGroupLabel(node);
      var fragment = document.createDocumentFragment();
      fragment.append(labelTd, tds);
      tds = fragment; //[labelTd,tds];
    }

    return tds;
  }

  renderInnerGroupLabel(node) {
    var td = document.createElement('td');
    td.classList.add('td_label');
    td.append(this.renderTagLabel(node));
    return td;
  }

  renderTagLabel(node) {
    var self = this;
    var text;
    var fieldName =
      node.tagName.toLowerCase() === 'label'
        ? node.getAttribute('for')
        : node.getAttribute('name');
    var text = node.getAttribute('string');
    if (!text) {
      if (fieldName) {
        text = fieldName; // this.state.fields[fieldName].string;
      } else {
        return this.renderGenericTag(node);
      }
    }

    var result = document.createElement('label');
    result.classList.add('form_label');
    result.setAttribute('for', this.getIDForLabel(fieldName));
    result.innerHTML = text;

    if (node.tagName.toLowerCase() === 'label') {
      this.handleAttributes(result, node);
    }
    var modifiersOptions;
    if (fieldName) {
      /*
      modifiersOptions = {
        callback: function(element, modifiers, record) {
          var widgets = self.allFieldWidgets[record.id];
          var widget = _.findWhere(widgets, { name: fieldName });
          if (!widget) {
            return; // FIXME this occurs if the widget is created
            // after the label (explicit <label/> tag in the
            // arch), so this won't work on first rendering
            // only on reevaluation
          }
          element.$el.toggleClass(
            "o_form_label_empty",
            !!(
              // FIXME condition is evaluated twice (label AND widget...)
              (record.data.id && (modifiers.readonly || self.mode === "readonly") && !widget.isSet())
            )
          );
        }
      };
      */
    }
    // FIXME if the function is called with a <label/> node, the registered
    // modifiers will be those on this node. Maybe the desired behavior
    // would be to merge them with associated field node if any... note:
    // this worked in 10.0 for "o_form_label_empty" reevaluation but not for
    // "o_invisible_modifier" reevaluation on labels...
    //this._registerModifiers(node, this.state, result, modifiersOptions);
    return result;
  }

  getIDForLabel(name) {
    var idForLabel = this.idsForLabels[name];
    if (!idForLabel) {
      idForLabel = _.uniqueId('o_field_input_');
      this.idsForLabels[name] = idForLabel;
    }
    return idForLabel;
  }
}

customElements.define('view-form', ViewForm);
