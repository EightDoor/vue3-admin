import {
  COLLAPSED,
  LOGIN,
  LOGINRESET,
  SET_MENUS_MUTATION,
  SET_SYS,
  SETUSERINFO,
  USERINFOMENUS,
  RESET,
} from '@/store/mutation-types'
import { LoginType, MenuType, UserInfoType, UserType } from '@/types/sys'
import { http } from '@/utils/request'
import { PERMISSIONBUTTONS, TOKEN } from '@/utils/constant'
import { Commit } from 'vuex'
import { MenuItem } from '@/types/layout/menu'
import { RouteRecordRaw } from 'vue-router'

import { Key } from 'ant-design-vue/es/_util/type'
import { ListObjCompare, ListToTree } from '@/utils'
import store from '../index'
import { cloneDeep } from 'lodash'
export interface SysStoreType {
  menus: MenuItem[]
  userInfo: UserType
  userInfoMenus: MenuType[]
  permissionButtons: MenuType[]
  collapsed: boolean
}
export type CustomMenus = RouteRecordRaw & { id: string; parent_id: string }

const getUserInfo = (): Promise<UserInfoType> => {
  return new Promise((resolve, reject) => {
    http({
      url: '/user/getWithTheMenu',
      method: 'get',
    })
      .then((res) => {
        // @ts-ignore
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
      path: item.path || '',
      name: item.name,
      component: () => import(`@/${item.component}`),
      redirect: item.redirect,
      meta: {
        title: item.title,
        icon: item.icon,
      },
      id: item.id || '',
      parent_id: item.parent_id,
    })
  })
  // @ts-ignore
  result = ListToTree(result)
  return result
}
function ListToTreeMenus(jsonData, id = 'id', pid = 'parent_id'): MenuItem[] {
  const result: MenuItem[] = [],
    temp = {}
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
      // 组合路由跳转地址
      currentElement.path = `${tempCurrentElementParent.path}/${currentElement.path}`
      // 组合面包屑
      currentElement.crumbs = `${tempCurrentElementParent.title},${currentElement.title}`
      tempCurrentElementParent['children'].push(currentElement) // 给父元素加上当前元素作为子元素
    } else {
      // 不存在父元素，意味着当前元素是一级元素
      result.push(currentElement)
    }
  }
  return result
}
function formatMenus(menus: MenuType[], status = false) {
  return menus.filter((item) => (status ? item.type === 3 : item.type !== 3))
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfoMenus: [],
    userInfo: null,
    collapsed: false,
  },
  mutations: {
    [COLLAPSED](state: SysStoreType): void {
      state.collapsed = !state.collapsed
    },
    [SET_MENUS_MUTATION](state: SysStoreType, payload: UserInfoType): void {
      const menus = formatMenus(payload.menu)
      let result: MenuItem[] = []
      const list = cloneDeep(menus.sort(ListObjCompare('order_num')))
      list.forEach((item) => {
        if (!item.hidden) {
          result.push({
            key: item.id || '',
            title: item.title,
            path: item.path,
            icon: item.icon,
            id: item.id,
            parent_id: item.parent_id,
            crumbs: `${item.title}`,
          })
        }
      })
      result = ListToTreeMenus(result)
      state.menus = result
      store.dispatch(RESET, payload.menu)
    },
    [SETUSERINFO](state: SysStoreType, payload: UserInfoType): void {
      state.userInfo = payload.user_info
    },
    [USERINFOMENUS](state: SysStoreType, payload: UserInfoType): void {
      state.userInfoMenus = payload.menu
    },
    [PERMISSIONBUTTONS](state: SysStoreType, payload: UserInfoType): void {
      const menus = cloneDeep(payload.menu)
      const data = formatMenus(menus, true)
      data.map((item) => {
        const r = menus.find((v) => v.id === item.parent_id)
        if (r) {
          item.name = r.id
        }
      })
      state.permissionButtons = data
    },
    [LOGINRESET](state: SysStoreType): void {
      state.permissionButtons = []
      state.userInfoMenus = []
      state.menus = []
      state.userInfo = {
        nick_name: '',
        status: 0,
        pass_word: '',
        account: '',
        dept_id: '',
      }
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
            commit(PERMISSIONBUTTONS, res)
            commit(USERINFOMENUS, res)
            commit(SETUSERINFO, res)
            commit(SET_MENUS_MUTATION, res)
            resolve({
              userInfo: res.user_info,
              menus: FormatMenuTree(formatMenus(res.menu)),
            })
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    [LOGIN](s: unknown, payload: LoginType): Promise<Key> {
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
          .then(async (res) => {
            // @ts-ignore
            localStorage.setItem(TOKEN, res)
            // @ts-ignore
            resolve(res)
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
  },
}
