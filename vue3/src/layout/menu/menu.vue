<template>
  <a-menu
    v-model:openKeys="menusInfo.openKeys"
    v-model:selectedKeys="menusInfo.selectedKeys"
    mode="inline"
    theme="dark"
    :inline-collapsed="collapsed"
    :collapsed="collapsed"
  >
    <template v-for="item in getMenus" :key="item.key">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <!--图标-->
          <span>{{ item.title }}</span>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :menu-info="item" :key="item.key" />
      </template>
    </template>
  </a-menu>
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted, computed, toRaw } from 'vue'
import SubMenu from './menu-item.vue'
import { MenuItem, MenusInfo } from '@/types/layout/menu'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { STORELETMENUPATH } from '@/utils/constant'
import { MenuType } from '@/types/sys'
import { localForage } from '@/utils/localforage'
import { SETCRUMBSLIST } from '@/store/mutation-types'
import { MenuFormatBrumb } from './menu-common'
import { PanesType } from '@/store/sys/sys-crumbs'

export default defineComponent({
  name: 'common-menu',
  setup() {
    const router = useRouter()
    const store = useStore()
    const menusInfo = reactive<MenusInfo>({
      selectedKeys: [],
      list: [],
      openKeys: [],
    })
    const getUserInfoMenus = computed(() => store.state.sys.userInfoMenus)
    onMounted(() => {
      menusInfo.list = []
      localForage.getItem<MenuType>(STORELETMENUPATH).then((res: any) => {
        if (res) {
          menusInfo.selectedKeys = [res.key || '']
          const parent_id = res.parent_id
          const data: MenuType[] = toRaw(getUserInfoMenus.value)
          const r = data.find((item) => item.id === parent_id)
          if (r) {
            menusInfo.openKeys = [r.id || '']
          }
        }
      })
    })

    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        store.commit(SETCRUMBSLIST, toRaw(item.crumbs))
        localForage.setItem(STORELETMENUPATH, toRaw(item))
        router.push({
          path: item.path,
        })
        const data: PanesType = {
          id: item.id || '',
          title: item.title,
        }
        MenuFormatBrumb(data)
      }
    }
    return {
      // data
      menusInfo,
      getMenus: computed(() => store.state.sys.menus),
      collapsed: computed(() => store.state.sys.collapsed),
      //methods
      jumpTo,
    }
  },
  components: {
    SubMenu,
  },
})
</script>
<style scoped lang="less">
@import './menu.less';
</style>
