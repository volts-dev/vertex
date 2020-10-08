import path from 'path';
import proxy from 'koa-proxies';
//import commonjs from '@rollup/plugin-commonjs';
//import scss from 'rollup-plugin-scss';
import rollupAlias from '@rollup/plugin-alias'; // 别名插件
import { fromRollup } from '@web/dev-server-rollup';
//import { importMapsPlugin } from '@web/dev-server-import-maps';
// const projectRootDir = path.resolve(__dirname);
const alias = fromRollup(rollupAlias);

export default {
  port: 8080,
  //open:false, // #BUG 直要设置就默认值为true
  watch: true,
  nodeResolve: true,
  appIndex: './index.html',

  middleware: [
    proxy('/app/*', {
      target: 'http://admin.localhost:16888',
      changeOrigin: true,
    }),
  ],

  plugins: [
    //scss(), // will output compiled styles to output.css
    // 配置别名
    alias({
      entries: [
        {
          find: '@',
          replacement: path.resolve(path.dirname(''), 'src'),
          // OR place `customResolver` here. See explanation below.
        },
      ],
    }),
  ],
};
