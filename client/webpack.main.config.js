const CopyPlugin = require("copy-webpack-plugin");

module.exports = {
  entry: './src/main.js',
  module: {
    rules: require('./webpack.rules'),
  },
  plugins: [
    // 复制图标文件信息
    new CopyPlugin({
      patterns: [{ from: "./icon", to: "icon" }]
    })
  ],
};
