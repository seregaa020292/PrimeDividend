
export default {
  path: "/:pathMatch(.*)*",
  component: () => import("@/pages/404.vue")
}
