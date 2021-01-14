import { SETCRUMBSLIST } from '@/store/mutation-types'

interface PanesType {
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
  },
}
