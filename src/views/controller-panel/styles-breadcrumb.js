import { css } from "lit-element";

export const breadcrumbStyles = css`
  .breadcrumb {
    display: flex;
    flex-wrap: wrap;
    padding: var(--breadcrumb-padding-y) var(--breadcrumb-padding-x);
    margin-bottom: var(--breadcrumb-margin-bottom);
    list-style: none;
    background-color: var(--breadcrumb-bg);
    @include border-radius(var(--breadcrumb-border-radius));
  }

  /* The separator between breadcrumbs (by default, a forward-slash: "/")*/
  .breadcrumb-item +.breadcrumb-item {
    padding-left: var(--breadcrumb-item-padding);
  }

  .breadcrumb-item +.breadcrumb-item ::before {
    display: inline-block; /*Suppress underlining of the separator in modern browsers*/
    padding-right: var(--breadcrumb-item-padding);
    color: var(--breadcrumb-divider-color);
    content: var(--breadcrumb-divider);
  }

  .breadcrumb-item:hover::before {
    text-decoration: underline;
  }
  /* stylelint-disable-next-line no-duplicate-selectors*/
  .breadcrumb-item:hover::before {
    text-decoration: none;
  }

  .breadcrumb-item .active {
    color: var(--breadcrumb-active-color);
  }
`;
