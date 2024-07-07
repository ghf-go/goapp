import api from "./api.js";
//https://www.axios-http.cn/docs/req_config
//https://v3.router.vuejs.org/zh/guide/
export default {
  install(app, options) {
    app.prototype.$post = api.$post;
    app.prototype.$get = api.$get;
    app.prototype.$delete = api.$delete;
  },
};
