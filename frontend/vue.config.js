const webpack = require('webpack');

module.exports = {
    transpileDependencies: ['true'],
    configureWebpack: {
        plugins: [
            new webpack.ProvidePlugin({
                process: 'process/browser',
            }),
        ],
    },
    lintOnSave: false,
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:1080',
                changeOrigin: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        }
    }
};