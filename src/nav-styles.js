import { html } from "lit-element";
const styles = html`
  <custom-style>
    <style is="custom-style">
      .nav {
        display: flex;
        flex-wrap: wrap;
        padding-left: 0;
        margin-bottom: 0;
        list-style: none;
      }

      .nav-link {
        display: block;
        padding: var(--nav-link-padding-y) var(--nav-link-padding-x);
      }
      .nav-link :hover,
      .nav-link :focus {
        text-decoration: none;
      }

      /* Disabled state lightens text*/
      .nav-link.disabled {
        color: var(--nav-link-disabled-color);
      }

      .nav-tabs {
        border-bottom: var(--nav-tabs-border-width) solid var(--nav-tabs-border-color);
      }

      .nav-item {
        margin-bottom: var(--nav-tabs-border-width);
      }

      .nav-link {
        border: var(--nav-tabs-border-width) solid transparent;
        border-top-left-radius: var(--nav-tabs-border-radius);
        border-top-right-radius: var(--nav-tabs-border-radius);
      }
      .nav-link:hover,
      .nav-link:focus {
        border-color: var(--nav-tabs-link-hover-border-color);
      }

      .nav-link.disabled {
        color: var(--nav-link-disabled-color);
        background-color: transparent;
        border-color: transparent;
      }

      .nav-link.active,
      .nav-item.show .nav-link {
        color: var(--nav-tabs-link-active-color);
        background-color: var(--nav-tabs-link-active-bg);
        border-color: var(--nav-tabs-link-active-border-color);
      }
      
    </style>
  </custom-style>
`;
const template = /** @type {!HTMLTemplateElement} */ (document.createElement("template"));
template.setAttribute("style", "display: none;");
template.innerHTML = styles.getHTML();
document.head.appendChild(template.content);
