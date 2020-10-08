import { css, html } from "lit-element";
import { FormElement } from "./../form-element";

class FormGroup extends FormElement {
  static get styles() {
    return css`
      :host {
        position: relative;
        display: block;
      }

      .form_group {
        width: 100%;
        margin: 9px 0 9px 0;
      }

      .horizontal_separator {
        font-weight: bold;
        font-size: 20px;
        margin: 15px 0px 10px 0px;
        color: #7c7bad;
      }
    `;
  }

  render() {
    return html`
      ${this.isOuterGroup
        ? html`
            <div class="form_group">
              <slot></slot>
            </div>
          `
        : html`
            <table class="form_group inner_group">
              ${this.string
                ? html`
                    <div class="horizontal_separator oe_clear "></div>
                    <tr>
                      <td colspan="${this.innerGroupColspan}" style="width: 100%;">
                        <div class="o_horizontal_separator">${this.string}</div>
                      </td>
                    </tr>
                  `
                : null}
            </table>
          `}
    `;
  }

  static get properties() {
    return {
      string: { type: String },
      name: { type: String },

      isOuterGroup: { type: Boolean },
      innerGroupColspan: { type: Number }
    };
  }

  constructor() {
    super();
    this.isOuterGroup = false;
    this.innerGroupColspan = 2;
  }

  firstUpdated() {
    super.firstUpdated();
    var groups = this.shadowRoot.querySelectorAll("form-group");
    this.innerGroupColspan = parseInt(this.getAttribute("col"), 10) || this.innerGroupColspan;
    this.isOuterGroup = groups.length > 0;
    if (!this.isOuterGroup) {
      return this.renderInnerGroup();
    }
    return this.renderOuterGroup();
  }

  renderInnerGroup() {
    var rows = [];
    this.children.forEach(child => {
      var colspan = parseInt(child.getAttribute("colspan"), 10);
      var isLabeledField = child.tag === "field" && child.getAttribute("nolabel") !== "1";
      if (!colspan) {
        if (isLabeledField) {
          colspan = 2;
        } else {
          colspan = 1;
        }
      }
      var finalColspan = colspan - (isLabeledField ? 1 : 0);
      currentColspan += colspan;
      if (currentColspan > this.innerGroupColspan) {
        rows.push(document.createElement("tr"));
        //$currentRow = $('<tr/>');
        currentColspan = colspan;
      }

      var $tds;
      if (child.tag === "field") {
        $tds = self.renderInnerGroupField(child);
      } else if (child.tag === "label") {
        $tds = self.renderInnerGroupLabel(child);
      } else {
        $tds = $("<td/>").append(self._renderNode(child));
      }
      if (finalColspan > 1) {
        $tds.last().attr("colspan", finalColspan);
      }
    });
  }
  renderOuterGroup() {}

  renderInnerGroupField(node) {
    var tds = document.createElement("tr");
    if (node.getAttribute("nolabel") !== "1") {
      var labelTd = this.renderInnerGroupLabel(node);
      tds = labelTd.add(tds);
    }

    return tds;
  }
  renderInnerGroupLabel(node) {
    var td = document.createElement("td");
    td.append(this.renderTagLabel(node));
  }

  renderTagLabel(node) {
    // 获取字段名称
    var fieldName = node.tag === "label" ? node.getAttribute("for") : node.getAttribute("name");
    var text = node.getAttribute("string");
    if (!text && fieldName) {
      text =fieldName// this.state.fields[fieldName].string;
    } else {
      return this.renderGenericTag(node);
    }
    var result = document.createElement("label");
    result.classList.add("form_label");
    result.innerHTML=text;
    result.for=this.getIDForLabel(fieldName);
 
  }

  // 渲染其他标签
  renderGenericTag(node) {}

  getIDForLabel(name) {
    var idForLabel = this.idsForLabels[name];
    if (!idForLabel) {
        idForLabel = _.uniqueId('o_field_input_');
        this.idsForLabels[name] = idForLabel;
    }
    return idForLabel;
},
}

customElements.define("form-group", FormGroup);
