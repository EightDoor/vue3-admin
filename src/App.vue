<template>
  <a-config-provider :locale="locale">
    <router-view></router-view>
  </a-config-provider>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import zhCN from 'ant-design-vue/es/locale/zh_CN'

export default defineComponent({
  name: 'App',
  setup() {
    return {
      locale: zhCN,
    }
  },
})
</script>
