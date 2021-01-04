import { createStore, createLogger } from 'vuex'

import sys, { SysStoreType } from './sys/sys-store'
const debug = process.env.NODE_ENV !== 'production'

interface RootState {
  sys: SysStoreType
}
export default createStore<RootState>({
  modules: {
    sys,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})
