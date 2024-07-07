import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

if (process.env.NODE_ENV == "production") {
  document.addEventListener(
    "deviceready",
    () => {
      new Vue({
        router,
        render: (h) => h(App),
      }).$mount("#app");
    },
    false
  );
} else {
  new Vue({
    router,
    render: (h) => h(App),
  }).$mount("#app");
}
