import {
  LOGIN,
  SET_MENUS_MUTATION,
  SET_SYS,
  SETUSERINFO,
} from '@/store/mutation-types'
import { LoginType, MenuType, UserInfoType, UserType } from '@/types/sys'
import { http } from '@/utils/request'
import { TOKEN } from '@/utils/constant'

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
      commit: any
    }): Promise<{ userInfo: UserType | null }> {
      return new Promise((resolve, reject) => {
        // 获取数据然后直接存储
        getUserInfo()
          .then((res) => {
            commit(SETUSERINFO, res)
            commit(SET_MENUS_MUTATION, res)
            resolve({
              userInfo: res.user_info,
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
        commit: any
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
