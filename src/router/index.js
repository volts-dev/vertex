import { Router } from './router';

var routes = [
  {
    name: 'app',
    pattern: 'app',
    data: { title: 'Home' },
  },
  {
    name: 'login',
    pattern: 'login',
  },
  {
    name: 'setup',
    pattern: 'setup',
  },
];

var router = new Router(routes);
export default router;
