import { css } from "lit-element";

export const styles = css`
  * {
    padding: 0;
    margin: 0;
  }

  input {
    border: none;
    width: 100%;
  }
  /* remove select border */
  input:focus {
    -webkit-box-shadow: none;
    box-shadow: none;
    outline: none;
  }

  :host {
    /*# 对于子元素Position定位关系非常重要*/
    position: relative;
    font-family: Roboto;
    line-height: 25px;
    padding: 0 0 1px 0;
    display: block;
    background-color: #fff;
  }

  .dropdown {
    background-image: none;
    background-color: #fff;
    border: 1px solid #dee2e6;
    font-size: 1.08333333rem;
    padding: 5px 0px;
    box-shadow: 0px 6px 12px rgba(0, 0, 0, 0.176);
    z-index: 1051;
  }

  .dropdown-list {
  }

  .option {
    font-weight: normal;
    display: block;
    white-space: pre;
    min-height: 1.2em;
    padding: 0px 2px 1px;
  }

  .option {
    position: relative;
    margin: 0;
    padding: 0px 1em 0px 0.4em;
    cursor: pointer;
    min-height: 0;
    list-style-image: url(data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7);
  }
  .input_box {
    width: 100%;
    width: 100px;
    -ms-flex: 1 0 auto;
    -moz-flex: 1 0 auto;
    -webkit-box-flex: 1;
    -webkit-flex: 1 0 auto;
    flex: 1 0 auto;
  }
  .searchview {
    width: 100%;
    border: 1px solid #ccc;
    border-radius: 3px;
    display: inline-block;
    padding: 2px 4px;
  }

  .searchview_facets {
    display: -ms-flexbox;
    display: -moz-box;
    display: -webkit-box;
    display: -webkit-flex;
    display: flex;
    flex-wrap: wrap;
  }

  .searchview_more {
    position: absolute;
    padding: 0;
    margin: 0;
    top: auto;
    left: auto;
    bottom: auto;
    right: 5px;
    font-size: 16px;
    cursor: pointer;
  }

  .searchview_facet {
    color: #8f8f8f;
    -ms-flex: 0 0 auto;
    -moz-flex: 0 0 auto;
    -webkit-box-flex: 0;
    -webkit-flex: 0 0 auto;
    flex: 0 0 auto;
    max-width: max-content;
    display: -ms-flexbox;
    display: -moz-box;
    display: -webkit-box;
    display: -webkit-flex;
    display: flex;
    position: relative;
    margin: 1px 3px 0 0;
    padding: 0.25em 0.4em;
    font-size: 12px;
    border-radius: 10rem;
    line-height: 1;
    box-shadow: inset 0 0 0 2px #777777;
  }

  .searchview_facet_label {
    -ms-flex: 0 0 auto;
    -moz-flex: 0 0 auto;
    -webkit-box-flex: 0;
    -webkit-flex: 0 0 auto;
    flex: 0 0 auto;
    display: inline-block;
    max-width: 100%;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    vertical-align: top;
    padding: 0 3px;
    color: white;
    display: -ms-flexbox;
    display: -moz-box;
    display: -webkit-box;
    display: -webkit-flex;
    display: flex;
    -webkit-align-items: center;
    align-items: center;
  }

  .searchview_facet_label {
    background-color: #777777;
  }

  .facet_remove {
    height: 12px;
    margin: auto;
    color: #777777;
  }
`;
