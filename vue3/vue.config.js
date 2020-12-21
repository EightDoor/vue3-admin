module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8081/api/v1',
        changeOrigin: true,
        pathRewrite: { '^/api': '' },
      },
    },
  },
}
