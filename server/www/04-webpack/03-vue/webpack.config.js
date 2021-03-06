const path = require('path')
const webpack = require('webpack')
const htmlWebpack = require('html-webpack-plugin')
const gulifyjsWebpack = require('uglifyjs-webpack-plugin')

module.exports = {
	entry: './src/index.js',
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'bundle.js',
		// publicPath: 'dist/',
	},
	module: {
		rules: [
			{
				test: /\.css$/,
				// use: ['css-loader', 'style-loader']
				use: ['style-loader', 'css-loader']
			},
			{
				test: /\.less$/,
				// use: ['css-loader', 'style-loader']
				use: [
					{
						loader: 'style-loader'
					},
					{
						loader: 'css-loader'
					},
					{
						loader: 'less-loader'
					}
				]
			},
			{
				test: /\.(png|jpg|gif|jpeg)$/,
				use: [
					{
						loader: 'url-loader',
						options: {
							limit: 8192,
							name: 'img/[name].[hash:8].[ext]'
						}
					}
				]
			},
			{
				test: /\.js/,
				exclude: /(node_modules|bower_components)/,
				use: {
					loader: 'babel-loader',
					options: {
						presets: ['es2015']
					}
				}
			},
			{
				test: /\.vue$/,
				// use: ['css-loader', 'style-loader']
				use: [
					{
						loader:'vue-loader',
					},
				],
			},
		]
	},
	resolve: {
		alias: {
			'vue$': 'vue/dist/vue.esm.js' // 用 webpack 1 时需用 'vue/dist/vue.common.js'
		}
	},
	plugins: [
		// new webpack.BannerPlugin('版权归胖子所有'),
		new htmlWebpack({
			template: './template.html',
		}),
		new gulifyjsWebpack(),
	],
}