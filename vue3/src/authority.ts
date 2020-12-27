import { SET_SYS } from '@/store/mutation-types'
import store from '@/store/index'
import { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'

// 路由白名单
const routerWhiteList = ["/login"];
export const canUserAccess = async(to: RouteLocationNormalized,from: RouteLocationNormalized,next: NavigationGuardNext) => {
  const status = routerWhiteList.find((item)=>to.path === item)
  if(!status) {
    if (store.state.userInfo) {
      next()
    } else {
      // 判断vuex是否存在值
      const {userInfo} = await store.dispatch(`${SET_SYS}`)
      if(userInfo) {
        next()
      }else {
        if(to.name !== 'login') {
          next({name: "login", replace: true})
        }else {
          next()
        }
      }
    }
  }else {
    next()
  }
}

