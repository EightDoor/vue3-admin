<template>
  <a-layout id="components-layout-demo-fixed-sider">
    <a-layout-sider class="sider-height">
      <div class="logo">vue3-gin-admin</div>
      <common-menu />
    </a-layout-sider>
    <a-layout>
      <common-header />
      <a-layout-content class="layout-content">
        <router-view />
      </a-layout-content>
      <a-layout-footer :style="{ textAlign: 'center' }">
        <common-footer />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

import CommonFooter from '@/layout/footer/footer.vue'
import CommonMenu from '@/layout/menu/menu.vue'
import CommonHeader from '@/layout/header/header.vue'

export default defineComponent({
  name: 'layout',
  setup() {
    const selectedKeys: string[] = ['4']
    return {
      selectedKeys,
    }
  },
  components: {
    CommonMenu,
    CommonFooter,
    CommonHeader,
  },
})
</script>
<style lang="less" scoped>
@import 'layout.less';
</style>
