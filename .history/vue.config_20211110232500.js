module.exports = {
  devServer: {
    host: "0.0.0.0",
    public: "0.0.0.0:8080",
    disableHostCheck: true,

    proxy: {
      "/api": {
        target: "http://localhost:8081/api",
        changeOrigin: true,
        pathRewrite: { "^/api": "" },
      },
    },
  },
  css: {
    requireModuleExtension: true,
  },
};
