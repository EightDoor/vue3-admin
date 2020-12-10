import { createStore, createLogger } from 'vuex'

import menu, { MenusStoreInter } from './menu/menus-store'
const debug = process.env.NODE_ENV !== 'production'

export default createStore<MenusStoreInter>({
  modules: {
    menu,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})
