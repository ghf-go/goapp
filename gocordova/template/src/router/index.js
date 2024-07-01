import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "",
    redirect: "/home/index",
  },
  {
    path: "/home",
    name: "Home",
    component: () => import("@/views/HomePage.vue"),
    children: [
      {
        path: "",
        redirect: "/home/index",
      },
    ],
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
