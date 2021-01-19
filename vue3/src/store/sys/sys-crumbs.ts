import {
  SETCRUMBSLIST,
  MENUTABS,
  DELETETABS,
  DELETETABSACTION,
  RESET,
  RESETMU,
} from '@/store/mutation-types'
import { MenuItem } from '@/types/layout/menu'
import { SysTabDel } from '@/types/sys/tab'
import { STORELETMENUPATH } from '@/utils/constant'
import { localForage } from '@/utils/localforage'
import _ from 'lodash'
import { toRaw } from 'vue'
import { Commit } from 'vuex'

export interface PanesType {
  id: string
  title: string
  path: string
  parent_id: string
  closable?: boolean
}
export interface CrumbsStoreType {
  list: string[]
  panes: PanesType[]
  delPane: PanesType
  selectPane: PanesType | string
}
interface ResetType {
  result: PanesType
  title: string
}
export default {
  namespace: true,
  state: {
    list: '',
    panes: [],
    selectPane: '',
  },
  mutations: {
    [SETCRUMBSLIST](state: CrumbsStoreType, payload: string[]): void {
      state.list = payload
    },
    [MENUTABS](state: CrumbsStoreType, payload: PanesType): void {
      const data: PanesType[] = state.panes
      data.push(toRaw(payload))
      state.panes = _.unionBy<PanesType>(data, 'title')
    },
    [DELETETABS](state: CrumbsStoreType, payload: SysTabDel): void {
      state.selectPane = payload.selectData
      if (payload.delData) {
        state.panes = state.panes.filter(
          (item: PanesType) => item.id !== payload.delData.id
        )
      }
    },
    // 设置当前的登录默认显示的选中项
    [RESETMU](state: CrumbsStoreType, payload: ResetType): void {
      // 默认首页只能是一级的
      state.list = [payload.title]
      state.panes = [payload.result]
      state.selectPane = payload.result
    },
  },
  actions: {
    [DELETETABSACTION](
      { commit }: { commit: Commit },
      payload: SysTabDel
    ): void {
      // 设置菜单选中项
      localForage
        .setItem(STORELETMENUPATH, toRaw(payload.selectData))
        .then(() => {
          commit(DELETETABS, payload)
        })
    },
    // 设置当前的登录默认显示的选中项
    // 默认首页只能是一级的
    // TODO 待完善  刷新页面直接选择对应的选中项
    [RESET]({ commit }: { commit: Commit }, payload: MenuItem[]): void {
      const data = payload.filter((item) => item.is_home)
      if (data.length > 0) {
        const r = data[0]
        const result = {
          id: r.id || '',
          title: r.title,
          path: r.path || '',
          parent_id: r.parent_id || '',
          closable: false,
        }
        commit(RESETMU, { result, title: r.title })
        localForage.setItem(STORELETMENUPATH, r)
      }
    },
  },
}
