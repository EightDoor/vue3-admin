import { SET_SYS } from '@/store/mutation-types'
import store from '@/store/index'
import {
  NavigationGuardNext,
  RouteLocationNormalized,
  Router,
  RouteRecordRaw,
} from 'vue-router'

// 路由白名单
const routerWhiteList = ['/login']
export const canUserAccess = async (
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext,
  router: Router
) => {
  const status = routerWhiteList.find((item) => to.path === item)
  if (!status) {
    if (store.state.userInfo) {
      next()
    } else {
      try {
        // 判断vuex是否存在值
        const { menus } = await store.dispatch(`${SET_SYS}`)
        console.log(to, 'To')
        menus.forEach((item: RouteRecordRaw) => {
          router.addRoute(item)
        })
        console.log(router.getRoutes())
        // TODO
        if (router.getRoutes().length > 5) {
          next()
        } else {
          next({ ...to, replace: true })
        }
      } catch (err) {
        if (to.name !== 'login') {
          next({ name: 'login', replace: true })
        } else {
          next()
        }
      }
    }
  } else {
    next()
  }
}
