import { ServicesCallerMixin } from './ServicesCallerMixin';
import { EventDispatcherMixin } from './EventDispatcherMixin';
import store from '@/store';

// 控件基类
export class Element extends ServicesCallerMixin(
  EventDispatcherMixin(store.Element)
) {}
