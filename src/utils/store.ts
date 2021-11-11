import localforage from 'localforage';

const localInstant = localforage.createInstance({
  name: 'vue3-admin-nest',
});

const store = {
  get<T = any>(key: string) {
    return localInstant.getItem<T>(key);
  },
  set(key: string, data: any) {
    return localInstant.setItem(key, data);
  },
  removeItem(key: string) {
    return localInstant.removeItem(key);
  },
  clear() {
    return localInstant.clear();
  },
};

export default store;
