import { SET_SYS } from '@/store/mutation-types'
import store from '@/store/index'
import { RouteLocationNormalized, Router, RouteRecordRaw } from 'vue-router'

// 路由白名单
const routerWhiteList = ['/login']
export const canUserAccess = async (
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  router: Router
) => {
  const status = routerWhiteList.find((item) => to.path === item)
  if (!status) {
    if (store.state.sys.userInfo) {
      return true
    } else {
      try {
        // 判断vuex是否存在值
        const { menus } = await store.dispatch(`${SET_SYS}`)
        menus.forEach((item: RouteRecordRaw) => {
          router.addRoute(item)
        })
        // console.log(router.getRoutes())
        // TODO 待验证
        await router.push({ path: to.path, replace: true })
        return true
      } catch (err) {
        if (to.name !== 'login') {
          return false
        } else {
          return true
        }
      }
    }
  } else {
    return true
  }
}
