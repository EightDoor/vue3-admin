module.exports = {
  devServer: {
    host: "0.0.0.0",
    public: "0.0.0.0:8080",
    disableHostCheck: true,
    hot: true,
    clientLogLevel: "info",
    watchOptions: {
      poll: true,
    },
    proxy: {
      "/api": {
        target: "http://localhost:9102/api",
        changeOrigin: true,
        pathRewrite: { "^/api": "" },
      },
    },
  },
  css: {
    requireModuleExtension: true,
  },
};
