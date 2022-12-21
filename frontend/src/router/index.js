import { createRouter, createWebHashHistory } from "vue-router";
import MainView from "../views/MainView.vue";
import LoginView from "../views/LoginView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
        path:"/", 
        component: LoginView,
    },
    {
        path: "/main",
        name: "main",
        component: MainView,
    },
  ],
});

export default router;