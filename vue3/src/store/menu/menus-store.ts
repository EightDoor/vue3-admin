import { MenuItem } from '@/types/alyout/menu'
import { SET_MENUS, SET_MENUS_MUTATION } from '@/store/mutation-types'

export interface MenusStoreInter {
  menu: any
  list: MenuItem[]
}
export default {
  namespace: true,
  state: {
    menus: [],
  },
  mutations: {
    [SET_MENUS_MUTATION](state: MenusStoreInter, payload: MenuItem[]) {
      state.list = payload
    },
  },
  actions: {
    [SET_MENUS]({ commit }: { commit: any }) {
      // 获取数据然后直接存储
      const list: MenuItem[] = [
        {
          key: '1',
          title: '首页',
          path: '/home',
        },
        {
          key: "3",
          title: "系统管理",
          path: "/sys",
          children: [
            {
              key: "5",
              title: "部门管理",
              path: "/sys/depart"
            },
            {
              key: "7",
              title: "角色管理",
              path: "/sys/role"
            },
            {
              key: "6",
              title: "用户管理",
              path: "/sys/user"
            },
          ]
        },
        {
          key: '2',
          title: 'Navigation 2',
          path: '/test',
          children: [
            {
              key: '2.1',
              title: 'Navigation 3',
              path: '/test/demo',
            },
          ],
        },
      ]
      commit(SET_MENUS_MUTATION, list)
    },
  },
}
