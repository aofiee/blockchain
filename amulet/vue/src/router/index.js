import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Post from "../views/Post.vue";
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/post",
    component: Post,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
