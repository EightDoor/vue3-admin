import { SETCRUMBSLIST, MENUTABS } from '@/store/mutation-types'
import _ from 'lodash'
export interface PanesType {
  id: string
  title: string
  closable?: boolean
}
export interface CrumbsStoreType {
  list: string[]
  panes: PanesType[]
}
export default {
  namespace: true,
  state: {
    list: '',
    panes: [],
  },
  mutations: {
    [SETCRUMBSLIST](state: CrumbsStoreType, payload: string[]) {
      state.list = payload
    },
    [MENUTABS](state: CrumbsStoreType, payload: PanesType) {
      const data: PanesType[] = state.panes
      data.push(payload)
      state.panes = _.unionBy<PanesType>(data)
    },
  },
}
