import {
  LOGIN,
  SET_MENUS_MUTATION,
  SET_SYS,
  SETUSERINFO,
} from '@/store/mutation-types'
import { LoginType, MenuType, UserInfoType, UserType } from '@/types/sys'
import { http } from '@/utils/request'
import { TOKEN } from '@/utils/constant'
import { Commit } from 'vuex'
import { MenuItem } from '@/types/layout/menu'
import { RouteRecordRaw } from 'vue-router'

import { Key } from 'ant-design-vue/es/_util/type'
import { ListObjCompare, ListToTree } from '@/utils'
import { cloneDeep } from 'lodash'

export interface SysStoreType {
  menus: MenuItem[]
  userInfo: UserType
}
type CustomMenus = RouteRecordRaw & { id: string; parent_id: string }

const getUserInfo = (): Promise<UserInfoType> => {
  return new Promise((resolve, reject) => {
    http({
      url: '/user/getWithTheMenu',
      method: 'get',
    })
      .then((res: any) => {
        resolve(res.list)
      })
      .catch((err) => {
        reject(err)
      })
  })
}
const FormatMenuTree = (item: MenuType[]): RouteRecordRaw[] => {
  // {
  //   name: 'home',
  //     path: 'home',
  //   component: () => import('@/views/home/home.vue'),
  //   meta: {
  //   title: '首页',
  //     icon: '23',
  // },
  //   redirect: '/home',
  // },
  let result: CustomMenus[] = []
  item.forEach((item) => {
    result.push({
      path: item.path,
      name: item.name,
      component: () => import(`@/${item.component}`),
      redirect: item.redirect,
      meta: {
        title: item.title,
        icon: item.icon,
      },
      id: item.id,
      parent_id: item.parent_id,
    })
  })
  result = ListToTree(result)
  return result
}
function ListToTreeMenus(jsonData: any[], id = 'id', pid = 'parent_id') {
  const result: any = [],
    temp: any = {}
  for (let i = 0; i < jsonData.length; i++) {
    temp[jsonData[i][id]] = jsonData[i] // 以id作为索引存储元素，可以无需遍历直接定位元素
  }
  for (let j = 0; j < jsonData.length; j++) {
    const currentElement = jsonData[j]
    const tempCurrentElementParent = temp[currentElement[pid]] // 临时变量里面的当前元素的父元素
    if (tempCurrentElementParent) {
      // 如果存在父元素
      if (!tempCurrentElementParent['children']) {
        // 如果父元素没有chindren键
        tempCurrentElementParent['children'] = [] // 设上父元素的children键
      }
      currentElement.path = `${tempCurrentElementParent.path}/${currentElement.path}`
      tempCurrentElementParent['children'].push(currentElement) // 给父元素加上当前元素作为子元素
    } else {
      // 不存在父元素，意味着当前元素是一级元素
      result.push(currentElement)
    }
  }
  return result
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfo: null,
  },
  mutations: {
    [SET_MENUS_MUTATION](state: SysStoreType, payload: UserInfoType) {
      let result: MenuItem[] = []
      const list = cloneDeep(payload.menu.sort(ListObjCompare('order_num')))
      list.forEach((item) => {
        if (!item.hidden) {
          result.push({
            key: item.id,
            title: item.title,
            path: item.path,
            icon: item.icon,
            id: item.id,
            parent_id: item.parent_id,
          })
        }
      })
      result = ListToTreeMenus(result)
      state.menus = result
    },
    [SETUSERINFO](state: SysStoreType, payload: UserInfoType) {
      state.userInfo = payload.user_info
    },
  },
  actions: {
    [SET_SYS]({
      commit,
    }: {
      commit: Commit
    }): Promise<{ userInfo: UserType | null; menus: RouteRecordRaw[] }> {
      return new Promise((resolve, reject) => {
        // 获取数据然后直接存储
        getUserInfo()
          .then((res) => {
            commit(SETUSERINFO, res)
            commit(SET_MENUS_MUTATION, res)
            resolve({
              userInfo: res.user_info,
              menus: FormatMenuTree(res.menu),
            })
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    [LOGIN](
      {
        commit,
      }: {
        commit: Commit
      },
      payload: LoginType
    ): Promise<Key> {
      return new Promise<string>((resolve, reject) => {
        const body = {
          account: payload.account,
          pass_word: payload.pass_word,
        }
        http({
          url: '/other/login',
          method: 'post',
          body,
        })
          .then(async (res: any) => {
            localStorage.setItem(TOKEN, res)
            resolve(res)
          })
          .catch((err) => {
            reject(err)
          })
      })
      // commit(LOGIN, LOGIN)
    },
  },
}
