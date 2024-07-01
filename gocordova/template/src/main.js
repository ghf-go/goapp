import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import vant from "vant";
import "vant/lib/index.css";
import Base from "./base";

let app = createApp(App).use(router).use(vant);
app.mixin({
  onShow() {
    console.log("--onShow--");
  },
});
app.use(Base);

if (process.env.NODE_ENV == "production") {
  document.addEventListener(
    "deviceready",
    () => {
      app.mount("#app");
    },
    false
  );
} else {
  app.mount("#app");
}
