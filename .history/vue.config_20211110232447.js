module.exports = {
  devServer: {

    proxy: {
      "/api": {
        target: "http://localhost:8081/api",
        changeOrigin: true,
        pathRewrite: { "^/api": "" },
      },
    },
    host: "localhost",
    hotOnly: true,
  },
  css: {
    requireModuleExtension: true,
  },
};
