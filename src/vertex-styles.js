import { html } from 'lit-element';
const styles = html`
  <custom-style>
    <style is="custom-style">
      html {
        /*Color system*/
        --white: #fff;
        --gray-100: #f8f9fa;
        --gray-200: #e9ecef;
        --gray-300: #dee2e6;
        --gray-400: #ced4da;
        --gray-500: #adb5bd;
        --gray-600: #6c757d;
        --gray-700: #495057;
        --gray-800: #343a40;
        --gray-900: #212529;
        --black: #000;
        --blue: #007bff;
        --indigo: #6610f2;
        --purple: #6f42c1;
        --pink: #e83e8c;
        --red: #dc3545;
        --orange: #fd7e14;
        --yellow: #ffc107;
        --green: #28a745;
        --teal: #20c997;
        --cyan: #17a2b8;

        /*
  /// This file regroups the variables that style components.
  /// They are available in every asset bundle.
  */

        /* Font sizes*/
        --root-font-size: 12px;
        --font-size-base: 13rem * calc(1px / var(--root-font-size));
        --line-height-base: 1.5; /* This is BS default*/

        /* Colors*/
        --v-community-color: #7c7bad;
        --v-enterprise-color: #875a7b;
        --v-enterprise-primary-color: #00a09d;

        --v-brand-vertex: var(--v-enterprise-color);
        --brand-primary: var(--v-community-color);

        --brand-secondary: #f0eeee;
        --brand-lightsecondary: #e2e2e0;
        --gray-100: #f8f9fa; /* This is BS default*/

        --main-color-muted: #a8a8a8;
        --main-text-color: #4c4c4c;

        --view-background-color: white;
        --shadow-color: #303030;

        --form-lightsecondary: #ccc;

        --list-footer-bg-color: #eee;
        --list-footer-font-weight: bold;

        --tooltip-background-color: black;
        --tooltip-color: white;
        --tooltip-arrow-color: black;
        --tooltip-text-color: #777777;
        --tooltip-title-text-color: #4c4c4c;

        /* Layout
//
// Extension of BS4. This is not redefining the BS4 variable directly as we only
// need the extra ones for media queries (not creating new breakpoint classes).
// Note: default BS4 values are hardcoded here while it should be possible to
// merge with the default BS variable (but we would have to take care of
// ordering & cie).
*/
        --v-navbar-height: 46px;
        --v-navbar-inverse-link-hover-bg: darken(--v-brand-vertex, 10%);
        --extra-grid-breakpoints: (
          xs: 0,
          vsm: 475px,
          sm: 576px,
          md: 768px,
          lg: 992px,
          xl: 1200px,
          xxl: 1534px
        );

        --form-group-cols: 12;
        --form-spacing-unit: 5px;
        --v-horizontal-padding: 16px;
        --v-innergroup-rpadding: 45px;
        --dropdown-hpadding: 20px;

        --sheet-vpadding: 24px;

        --notification-error-bg-color: #f16567;
        --notification-info-bg-color: #fcfbea;

        /* Needed for having no spacing between sheet and mail body in mass_mailing:*/
        /* Different required cancel paddings between web and web_enterprise*/
        --sheet-cancel-tpadding: 0px;

        --avatar-size: 90px;

        --statusbar-height: 33px;

        --label-font-size-factor: 0.8;
        --navbar-height: 46px;

        --nb-calendar-colors: 24;

        --base-settings-mobile-tabs-height: 40px;
        --base-settings-mobile-tabs-overflow-gap: 3%;

        --cp-breadcrumb-height: 30px;

        --datepicker-week-color: #8f8f8f;

        /* Kanban*/

        --kanban-default-record-width: 300px;
        --kanban-small-record-width: 240px;

        --kanban-header-title-height: 50px;

        --kanban-image-width: 64px;
        --kanban-image-fill-width: 95px;
        --kanban-inside-vgutter: 8px;
        --kanban-inside-hgutter: 8px;
        --kanban-color-border-width: 3px;
        --kanban-inner-hmargin: 5px;
        --kanban-progressbar-height: 20px;

        --kanban-mobile-tabs-height: 40px;

        /* ------- Kanban dashboard variables -------*/

        /* Used to manage spacing in complex dropdown menu*/
        --kanban-dashboard-dropdown-complex-gap: 5px;

        /* For the frontend part*/
        --theme-font-size-base: calc(14 / 16) * 1rem;

        /* Navs */
        --nav-link-padding-y: 0.5rem;
        --nav-link-padding-x: 1rem;
        --nav-link-disabled-color: var(--gray-600);

        --nav-tabs-border-color: var(--gray-300);
        --nav-tabs-border-width: var(--border-width);
        --nav-tabs-border-radius: var(--border-radius);
        --nav-tabs-link-hover-border-color: var(--gray-200) var(--gray-200)
          var(--nav-tabs-border-color);
        --nav-tabs-link-active-color: var(--gray-700);
        --nav-tabs-link-active-bg: var(--body-bg);
        --nav-tabs-link-active-border-color: var(--gray-300) var(--gray-300)
          var(--nav-tabs-link-active-bg);

        --nav-pills-border-radius: var(--border-radius);
        --nav-pills-link-active-color: var(--component-active-color);
        --nav-pills-link-active-bg: var(--component-active-bg);

        --nav-divider-color: var(--gray-200);
        --nav-divider-margin-y: calc(var(--spacer) / 2);
      }
    </style>
  </custom-style>
`;
const template = /** @type {!HTMLTemplateElement} */ (document.createElement(
  'template'
));
template.setAttribute('style', 'display: none;');
template.innerHTML = styles.getHTML();
document.head.appendChild(template.content);
