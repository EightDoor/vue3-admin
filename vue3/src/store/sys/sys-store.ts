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
import Layout from '@/layout/layout/layout.vue'
import Home from '@/views/home/home.vue'
import LayoutChildren from '@/views/layout-children.vue'

import SysUser from '@/views/sys/user.vue'
import SysMenu from '@/views/sys/menu.vue'
import SysRole from '@/views/sys/role.vue'
import Depart from '@/views/sys/depart.vue'

export interface SysStoreType {
  menus: MenuType[]
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
      component: Layout,
      redirect: '/home',
      children: [
        {
          name: 'home',
          path: 'home',
          component: Home,
          meta: {
            title: '首页',
            icon: '23',
          },
        },
        {
          name: 'sys',
          path: 'sys',
          component: LayoutChildren,
          meta: {
            title: '系统管理',
            icon: '1',
          },
          children: [
            {
              name: 'user',
              path: 'user',
              component: SysUser,
              meta: {
                title: '用户管理',
                icon: '1',
              },
            },
            {
              name: 'menu',
              path: 'menu',
              component: SysMenu,
              meta: {
                title: '菜单管理',
                icon: '1',
              },
            },
            {
              name: 'role',
              path: 'role',
              component: SysRole,
              meta: {
                title: '角色管理',
                icon: '1',
              },
            },
            {
              name: 'depart',
              path: 'depart',
              component: Depart,
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
      state.menus = payload.menu
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
        // const list: MenuItem[] = [
        //   {
        //     key: '1',
        //     title: '首页',
        //     path: '/home',
        //   },
        //   {
        //     key: '3',
        //     title: '系统管理',
        //     path: '/sys',
        //     children: [
        //       {
        //         key: '5',
        //         title: '部门管理',
        //         path: '/sys/depart',
        //       },
        //       {
        //         key: '8',
        //         title: '菜单管理',
        //         path: '/sys/menu',
        //       },
        //       {
        //         key: '7',
        //         title: '角色管理',
        //         path: '/sys/role',
        //       },
        //       {
        //         key: '6',
        //         title: '用户管理',
        //         path: '/sys/user',
        //       },
        //     ],
        //   },
        // ]
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
