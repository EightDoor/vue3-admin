import { SETCRUMBSLIST, MENUTABS, DELETETABS } from '@/store/mutation-types'
import { SysTabDel } from '@/types/sys/tab'
import _ from 'lodash'
import { toRaw } from 'vue'
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
export default {
  namespace: true,
  state: {
    list: '',
    panes: [],
    selectPane: '',
  },
  mutations: {
    [SETCRUMBSLIST](state: CrumbsStoreType, payload: string[]) {
      state.list = payload
    },
    [MENUTABS](state: CrumbsStoreType, payload: PanesType) {
      const data: PanesType[] = state.panes
      data.push(toRaw(payload))
      state.panes = _.unionBy<PanesType>(data, 'title')
    },
    [DELETETABS](state: CrumbsStoreType, payload: SysTabDel) {
      state.selectPane = payload.selectData
      if (payload.delData) {
        state.panes = state.panes.filter(
          (item: PanesType) => item.id !== payload.delData.id
        )
      }
    },
  },
}
