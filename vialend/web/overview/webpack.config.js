var webpack = require('webpack');
var path = require('path');
module.exports = {
    entry:'./src/index.js',
    output:{
        path:path.resolve(__dirname, './dist/'),
        filename:'build.js'
    },
    devServer: {
        inline: true,
        port: 8181
    },
    module:{
        loaders:[
            {
                test:/\.js?$/,
                exclude:/node_modules/,
                loader:'babel-loader',
                query:{
                    presets:['es2015','react']
                }
            },
            {
                test:/\.css$/,
                loader:'style-loader!css-loader'
            }
        ]
    }
};