import { createStore, createLogger } from 'vuex'

import sys, { SysStoreType } from './sys/sys-store'
const debug = process.env.NODE_ENV !== 'production'

export type StoreType = SysStoreType
export default createStore<StoreType>({
  modules: {
    sys,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})
