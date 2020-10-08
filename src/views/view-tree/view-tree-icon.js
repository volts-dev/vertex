import '@polymer/iron-iconset-svg/iron-iconset-svg.js';
import {html} from '@polymer/polymer/lib/utils/html-tag.js';

const template = html`<iron-iconset-svg name="view-tree-icon" size="24">
<svg><defs>
<g id="tree"><path d="M4 14h4v-4H4v4zm0 5h4v-4H4v4zM4 9h4V5H4v4zm5 5h12v-4H9v4zm0 5h12v-4H9v4zM9 5v4h12V5H9z"/></g>
</defs></svg>
</iron-iconset-svg>

<iron-iconset-svg name="view-list-icon" size="24">
<svg><defs>
<g id="list"><path d="M4 14h4v-4H4v4zm0 5h4v-4H4v4zM4 9h4V5H4v4zm5 5h12v-4H9v4zm0 5h12v-4H9v4zM9 5v4h12V5H9z"/></g>
</defs></svg>
</iron-iconset-svg>`

document.head.appendChild(template.content);