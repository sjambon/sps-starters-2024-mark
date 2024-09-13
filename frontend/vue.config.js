const webpack = require('webpack');

module.exports = {
    transpileDependencies: ['true'],
    configureWebpack: {
        plugins: [
            new webpack.ProvidePlugin({
                process: 'process/browser',
            }),
        ],
        resolve: {
        },
    },
    lintOnSave: false
};