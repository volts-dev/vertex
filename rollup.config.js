import replace from 'rollup-plugin-replace';
export default {
  input: 'src/vertex.js',
  output: {
    file: 'dist/bundle.js',
    format: 'cjs',
  },

  plugins: [
    replace({
      'process.env.NODE_ENV': JSON.stringify('production'),
      delimiters: ['', ''],
    }),
  ],
};
