const webpack = require('webpack');

module.exports = {
    transpileDependencies: true,
    configureWebpack: {
        plugins: [
            new webpack.ProvidePlugin({
                process: 'process/browser',
            }),
        ],
        resolve: {
            fallback: {
                path: require.resolve('path-browserify'),
                os: require.resolve('os-browserify/browser'),
                crypto: require.resolve('crypto-browserify'),
            },
        },
    },
    lintOnSave: false
};