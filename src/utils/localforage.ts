import * as localforage from 'localforage'

const localForage = localforage.createInstance({
  name: 'vue3-gin-admin',
})

export { localForage }
