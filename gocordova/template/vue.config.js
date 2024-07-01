const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  // define: {
  //   "process.env": {},
  //   __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: "true",
  // },
  transpileDependencies: true,
  outputDir: "www",
  publicPath: "/",
});
