<template>
  <a-sub-menu :key="menuInfo.key" v-bind="$attrs">
    <template #title>
      <span>
        <!--图标-->
        {{ menuInfo.title }}</span
      >
    </template>
    <template v-for="item in menuInfo.children" :key="item.key">
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
  </a-sub-menu>
</template>

<script lang="ts">
import { defineComponent, toRaw } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { MenuItem } from '@/types/layout/menu'
import { STORELETMENUPATH } from '@/utils/constant'
import { localForage } from '@/utils/localforage'
import { SETCRUMBSLIST } from '@/store/mutation-types'
import { MenuFormatBrumb } from './menu-common'

export default defineComponent({
  name: 'SubMenu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  setup: function () {
    const router = useRouter()
    const store = useStore()
    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        store.commit(SETCRUMBSLIST, toRaw(item.crumbs))
        localForage.setItem(STORELETMENUPATH, toRaw(item)).then(() => {
          MenuFormatBrumb(item)
          router.push({
            path: item.path || '',
          })
        })
      }
    }
    return {
      //methods
      jumpTo,
    }
  },
})
</script>
