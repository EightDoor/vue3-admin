import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import { canUserAccess } from '@/authority'
import nprogress from 'nprogress'
import 'nprogress/nprogress.css' // progress bar style

import Login from '@/views/login/login.vue'
import NotFound from '@/views/other/not-found.vue'

const staticRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
]
const app = createRouter({
  history: createWebHashHistory(),
  routes: staticRoutes,
})

// 全局路由前置钩子
app.beforeEach(async (to, from) => {
  nprogress.start()
  await canUserAccess(to, from, app)
})

app.afterEach(() => {
  nprogress.done()
})

export default app
