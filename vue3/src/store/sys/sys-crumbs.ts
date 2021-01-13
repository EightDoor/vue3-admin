import { SETCRUMBSLIST } from '@/store/mutation-types'

export interface CrumbsStoreType {
  list: string[]
}
export default {
  namespace: true,
  state: {
    list: '',
  },
  mutations: {
    [SETCRUMBSLIST](state: CrumbsStoreType, payload: string[]) {
      state.list = payload
    },
  },
}
