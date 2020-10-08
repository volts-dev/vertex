/**
 * This module defines the view_registry. Web views are added to the registry
 * in the 'web._view_registry' module to avoid cyclic dependencies.
 * Views defined in other addons should be added in this registry as well,
 * ideally in another module than the one defining the view, in order to
 * separate the declarative part of a module (the view definition) from its
 * 'side-effects' part.
 */
import { Registry } from '@/core/registry';
import { ViewSearch } from '@/views/view-search/view-search';
import { ViewForm } from '@/views/view-form/view-form';
//import { ViewTree } from "@/views/view-tree";

export var view_registry = new Registry();
view_registry.add('view-search', ViewSearch).add('view-form', ViewForm);
//.add("view-tree", ViewTree);
