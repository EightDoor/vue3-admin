import { MenuItem } from '@/types/alyout/menu'
import { SET_MENUS_MUTATION, SET_SYS, SETUSERINFO } from '@/store/mutation-types'
import { UserType } from '@/types/sys'

export interface SysStoreType {
  menus: MenuItem[],
  userInfo: UserType
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfo: null
  },
  mutations: {
    [SET_MENUS_MUTATION](state: SysStoreType, payload: MenuItem[]) {
      state.menus = payload
    },
    [SETUSERINFO](state: SysStoreType, payload: UserType){
      state.userInfo = payload
    }
  },
  actions: {
    [SET_SYS]({ commit }: { commit: any }):Promise<{userInfo: UserType}> {
      return new Promise((resolve, reject)=>{
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
                key: "8",
                title: "菜单管理",
                path: "/sys/menu"
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
        ]
        const userInfo: UserType = {
          id: "123",
          nick_name: "999",
          pass_word: "123",
          account: "123",
          status: 1,
          dept_id: "999"
        }
        commit(SET_MENUS_MUTATION, list)
        commit(SETUSERINFO, userInfo)
        resolve({
          userInfo
        })
      })
    },
  },
}
