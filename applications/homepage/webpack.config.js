const path = require('path')

const webpack = require('webpack')

module.exports = {
  mode: 'development',
  entry: {
    homepage: path.resolve(__dirname, 'src/main.js')
  },
  resolve: {
    extensions: ['.js', '.vue', '.json', '.css'],
    alias: {
      'vue$': 'vue/dist/vue.esm.js',
      '@': path.join(__dirname, 'src')
    }
  },
  output: {
    path: path.resolve(__dirname, 'public/dist'),
    filename: '[name].min.js'
  },
  module: {
    rules: [{
      test: /\.vue$/,
      loader: 'vue-loader'
    }, {
      test: /\.js$/,
      loader: 'babel-loader',
      exclude: /node_modules/
    }]
  },
  plugins: []
}
