<template>
  <a-menu
    :default-selected-keys="menusInfo.selectedKeys"
    :default-open-keys="menusInfo.selectedKeys"
    mode="inline"
    theme="dark"
    :inline-collapsed="menusInfo.collapsed"
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
import { defineComponent, reactive, onMounted, computed } from 'vue'
import SubMenu from './menu-item.vue'
import { MenuItem, MenusInfo } from '@/types/alyout/menu'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default defineComponent({
  name: 'common-menu',
  setup() {
    const router = useRouter()
    const store = useStore()
    const menusInfo = reactive<MenusInfo>({
      selectedKeys: [],
      collapsed: false,
      list: [],
    })
    onMounted(() => {
      menusInfo.list = []
    })

    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        router.push({
          path: item.path,
        })
      }
    }
    return {
      menusInfo,
      getMenus: computed(() => store.state.sys.menus),
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
