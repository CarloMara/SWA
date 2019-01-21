module.exports = {
	entry: './src/js/app.js',
	output: {
		path: __dirname + '/public',
		filename: 'bundle.js'
	},
	module:{
		rules:[
			{test: /\.css$/, loader: "style-loader!css-loader"},
			{
                test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
                use: [{
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'fonts/'
                    }
                }]
            }

		]

	}
}