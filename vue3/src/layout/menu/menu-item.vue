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
import { defineComponent, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { MenuItem } from '@/types/alyout/menu'
export default defineComponent({
  name: 'sub-menu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  setup: function () {
    const router = useRouter()
    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        router.push({
          path: item.path,
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
