import { SET_MENUS } from '@/store/mutation-types'
import store from '@/store/index'

export const canUserAccess = () => {
  // TODO
  if (store.state.menu.list && store.state.menu.list.length > 0) {
    // 直接判断是否有权限
  } else {
    // 判断vuex是否存在值,不存在的话
    store.dispatch(`${SET_MENUS}`)
  }
  return true
}
