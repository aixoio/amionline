import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import("@/views/HomeView.vue")
    },
    {
      path: '/stats',
      name: 'stats',
      component: () => import("@/views/StatsView.vue")
    },
  ]
})

export default router
