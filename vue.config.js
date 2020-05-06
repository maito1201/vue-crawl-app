module.exports = {
  devServer: {
    proxy: {
      '^/api': {
      target: 'http://api:3000',
      changeOrigin: true,
      secure: false,
      pathRewrite: {'^/api': ''},
      logLevel: 'debug' 
      }
    },
    watchOptions: {
      poll: true
    },
    disableHostCheck: true
  }
}
