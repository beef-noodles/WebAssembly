const path = require("path");
const CopyPlugin = require("copy-webpack-plugin");
const WasmPackPlugin = require("@wasm-tool/wasm-pack-plugin");
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const dist = path.resolve(__dirname, "docs");

module.exports = {
  mode: "production",
  entry: {
    index: "./front/index.tsx"
  },
  output: {
    path: dist,
    filename: "[name].js"
  },
  devtool: "source-map",
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"]
  },
  devServer: {
    contentBase: dist,
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: "awesome-typescript-loader" },
      { enforce: "pre", test: /\.js$/, loader: "source-map-loader" },
      {
        test: /\.(le|sa|sc|c)ss$/,
        use: [
          'css-hot-loader',
          MiniCssExtractPlugin.loader,
          {
            loader: 'css-loader',
            options: {
              // importLoaders: 1,
              modules: true,
              modules: {
                localIdentName: '[local]_[hash:base64:6]'
              }
            }
          },
          // {
          //   loader: 'postcss-loader',
          //   options: {
          //     sourceMap: !isProduction,
          //     ident: 'postcss'
          //   }
          // },
          // {
          //   loader: 'resolve-url-loader'
          // },
          {
            loader: 'less-loader',
            options: {
              javascriptEnabled: true,
              sourceMap: true,
            }
          }
        ]
      }
    ]
  },
  // externals: {
  //   "react": "React",
  //   "react-dom": "ReactDOM"
  // },
  plugins: [
    new CopyPlugin([
      path.resolve(__dirname, "static")
    ]),

    new MiniCssExtractPlugin({
      filename: 'static/style/[name].css'
    }),
    new WasmPackPlugin({
      crateDirectory: __dirname,
      extraArgs: "--out-name index"
    }),
  ]
};
