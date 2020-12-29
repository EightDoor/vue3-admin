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
import { MenuItem } from '@/types/alyout/menu'
import { RouteRecordRaw, useRouter } from 'vue-router'

import Depart from '@/views/sys/depart.vue'

export interface SysStoreType {
  menus: MenuItem[]
  userInfo: UserType
}

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
  const list: MenuItem[] = []
  const router = useRouter()
  // item.forEach((item) => {})
  const dynamicRoutes: RouteRecordRaw[] = [
    {
      path: '/',
      name: 'layout',
      component: () => import('@/layout/layout/layout.vue'),
      redirect: '/home',
      children: [
        {
          name: 'home',
          path: 'home',
          component: () => import('@/views/home/home.vue'),
          meta: {
            title: '首页',
            icon: '23',
          },
        },
        {
          name: 'sys',
          path: 'sys',
          component: () => import('@/views/layout-children.vue'),
          meta: {
            title: '系统管理',
            icon: '1',
          },
          children: [
            {
              name: 'user',
              path: 'user',
              component: () => import('@/views/sys/user.vue'),
              meta: {
                title: '用户管理',
                icon: '1',
              },
            },
            {
              name: 'menu',
              path: 'menu',
              component: () => import('@/views/sys/menu.vue'),
              meta: {
                title: '菜单管理',
                icon: '1',
              },
            },
            {
              name: 'role',
              path: 'role',
              component: () => import('@/views/sys/role.vue'),
              meta: {
                title: '角色管理',
                icon: '1',
              },
            },
            {
              name: 'depart',
              path: 'depart',
              component: () => import('@/views/sys/depart.vue'),
              meta: {
                title: '部门管理',
                icon: '1',
              },
            },
          ],
        },
      ],
    },
  ]

  console.log(item)
  return dynamicRoutes
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfo: null,
  },
  mutations: {
    [SET_MENUS_MUTATION](state: SysStoreType, payload: UserInfoType) {
      console.log(payload.menu, 'menu')
      const list: MenuItem[] = [
        {
          key: '1',
          title: '首页',
          path: '/home',
        },
        {
          key: '3',
          title: '系统管理',
          path: '/sys',
          children: [
            {
              key: '5',
              title: '部门管理',
              path: '/sys/depart',
            },
            {
              key: '8',
              title: '菜单管理',
              path: '/sys/menu',
            },
            {
              key: '7',
              title: '角色管理',
              path: '/sys/role',
            },
            {
              key: '6',
              title: '用户管理',
              path: '/sys/user',
            },
          ],
        },
      ]
      // state.menus = payload.menu
      state.menus = list
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
    ): Promise<string> {
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
