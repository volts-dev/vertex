import { css } from "lit-element";

export const formStyles = css`
  a {
    color: #7c7bad;
    text-decoration: none;
    background-color: transparent;
    -webkit-text-decoration-skip: objects;
  }

  a:hover {
    color: #555487;
    text-decoration: none;
  }

  :host {
    --sheet-max-width: 1140px;
    --sheet-min-width: 650px;
    --sheet-padding: 16px;
  }

  .field_widget {
    width: 100%;
  }

  /* No sheet*/
  .form_nosheet {
    display: block;
    @include o-webclient-padding($top: $o-sheet-vpadding, $bottom: $o-sheet-vpadding);
  }
  .form_nosheet .form_statusbar {
    margin: var(--sheet-vpadding) var(--horizontal-padding) var(--sheet-vpadding) var(--horizontal-padding);
  }

  .form_nosheet.form_nomargin {
    margin: 0;
  }

  .form_sheet_bg {
    border-bottom: 1px solid #ddd;
    background: url(/app/static/src/img/form_sheetbg.png);
  }
  .form_sheet_bg .form_sheet {
    margin: calc(var(--sheet-vpadding) * 0.5) auto;
  }

  .form_sheet {
    min-width: var(--sheet-min-width);
    max-width: var(--sheet-max-width);
    min-height: 330px;
    padding: 16px;
    border: 1px solid #c8c8d3;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
    background: white;
  }
  .form_sheet.ui-tabs {
    margin: 0 -16px;
  }
  .form_sheet .notebook_page {
    padding: 0 16px;
  }
  .form_statusbar {
    position: relative; /* Needed for the "More" dropdown*/
    display: flex;
    justify-content: space-between;
    padding-left: var(--horizontal-padding);
    border-bottom: 1px solid gray("400");
    background-color: var(--view-background-color);
  }

  .form_statusbar > .statusbar_buttons,
  .form_statusbar > .statusbar_status {
    display: flex;
    align-items: center;
    align-content: space-around;
  }

  /* Avatar*/
  .avatar {
    float: left;
  }
  .avatar > img {
    max-height: 90px;
    max-width: 90px;
    margin-bottom: 10px;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.4);
    border: none;
  }

  /* Title*/
  .title > h1,
  .title > h2,
  .title > h3 {
    width: 100%; /* Needed because inline-block when is a hx.o_row*/
    margin-top: 0;
    margin-bottom: 0;
    line-height: inherit;
  }
  .title .o_priority > .o_priority_star {
    font-size: inherit;
  }
  .avatar + .title {
    padding-left: 95px;
    margin-left: 5px;
  }
  /* Button box*/
  .button_box {
    width: 400px;
    text-align: right;
    float: right;
  }
  /* Separators*/
  .horizontal_separator {
    font-size: var(--h2-font-size);
    margin: var(--form-spacing-unit) 0;
  }
  .horizontal_separator:empty {
    height: calc(var(--form-spacing-unit) * 2);
  }
  /* Notebooks*/
  .tabs {
    clear: both; /* For the notebook not to have alongside floating elements*/
    margin-top: calc(var(--form-spacing-unit) * 2);
  }
  .tabs .nav-tabs > .tab-pane {
    min-height: 100px;
  }

  .nav {
    display: -ms-flexbox;
    display: flex;
    -ms-flex-wrap: wrap;
    flex-wrap: wrap;
    padding-left: 0;
    margin-bottom: 0;
    list-style: none;
  }

  .nav-link {
    display: block;
    padding: 0.5rem 1rem;
  }

  .nav-link:hover,
  .nav-link:focus {
    text-decoration: none;
  }

  .nav-link.disabled {
    color: #6c757d;
  }

  .nav-tabs {
    border-bottom: 1px solid #dee2e6;
  }

  .nav-tabs .nav-item {
    margin-bottom: -1px;
  }

  .nav-tabs .nav-link {
    border: 1px solid transparent;
    border-top-left-radius: 0.25rem;
    border-top-right-radius: 0.25rem;
  }

  .nav-tabs .nav-link:hover,
  .nav-tabs .nav-link:focus {
    border-color: #e9ecef #e9ecef #dee2e6;
  }

  .nav-tabs .nav-link.disabled {
    color: #6c757d;
    background-color: transparent;
    border-color: transparent;
  }

  .nav-tabs .nav-link.active,
  .nav-tabs .nav-item.show .nav-link {
    color: #495057;
    background-color: #fff;
    border-color: #dee2e6 #dee2e6 #fff;
  }

  .nav-tabs .dropdown-menu {
    margin-top: -1px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }

  .dropdown-menu {
    /* Make dropdown border overlap tab border*/
    margin-top: var(--nav-tabs-border-width);
    /* Remove the top rounded corners here since there is a hard edge above the menu*/
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }
  /* Tabbable tabs*/
  /* Hide tabbable panes to start, show them when .active*/
  .tab-content > .tab-pane {
    display: none;
  }
  .tab-content > .active {
    display: block;
  }
  /* Labels*/
  .form_label {
    margin: 0 var(--form-spacing-unit) 0 0;
    font-size: var(--font-size-base); /* The label muse have the same size whatever their position*/
    line-height: var(--line-height-base);
    font-weight: bold;
  }

  /* Groups*/
  .group {
    display: inline-block;
    width: 100%;
    margin: 10px 0;
  }

  .group_col_1 {
    display: inline-block;
    width: calc(100% / 12 * 1);
    vertical-align: top;
  }
  .group_col_6 {
    display: inline-block;
    width: calc(100% / 12 * 6);
    vertical-align: top;
  }
  .group_col_12 {
    display: inline-block;
    width: calc(100% / 12 * 12);
    vertical-align: top;
  }
  .group.td_label {
    border-right: 1px solid #ddd;
  }
  .group.td_label + td {
    padding: 2px 36px 2px 8px;
  }

  .group.field_widget.text_overflow {
    width: 1px !important; /* hack to make the table layout believe it is a small element (so that the table does not grow too much) ...*/
    min-width: 100%; /* ... but in fact it takes the whole table space*/
  }

  .group .td_label {
    border-right: 1px solid #ddd;
  }
  .group .td_label + td {
    padding: 2px 36px 2px 8px;
  }
  .group .field_widget.o_text_overflow {
    width: 1px !important; /* hack to make the table layout believe it is a small element (so that the table does not grow too much) ...*/
    min-width: 100%; /* ... but in fact it takes the whole table space*/
  }

  .inner_group {
    display: inline-table;
  }
  .inner_group > tbody > tr > td,
  .inner_group > tbody > tr > tdtd_label {
    width: 0%;
    padding: 0 15px 0 0;
    min-width: 150px;
    vertical-align: top;
  }
  .inner_group span,
  .inner_group .field_boolean,
  .inner_group .avatar,
  .inner_group .form_uri,
  .inner_group .field_widget {
    width: auto;
  }

  /*Editable specific rules*/

  .form_editable .read_only {
    display: none !important;
  }

  :host(.form_editable) .title {
    max-width: calc(720px - (2 * var(--horizontal-padding)));
  }
`;
