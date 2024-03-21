export default {
  path: '/:pathMatch(.*)*',
  component: () => import('@/pages/PageNotFound.vue'),
}
