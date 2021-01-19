import { createStore, createLogger } from 'vuex'

import sys, { SysStoreType } from './sys/sys-store'
import crumbs, { CrumbsStoreType } from './sys/sys-crumbs'

const debug = process.env.NODE_ENV !== 'production'

interface RootState {
  sys: SysStoreType
  crumbs: CrumbsStoreType
}
export default createStore<RootState>({
  modules: {
    // @ts-ignore
    sys,
    crumbs,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})
