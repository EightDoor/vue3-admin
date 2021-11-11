<template>
  <a-result status="404" title="404" sub-title="页面没有找到哦!">
    <template #extra>
      <a-button type="primary" @click="GoHome()"> 返回首页 </a-button>
    </template>
  </a-result>
</template>
<script lang="ts">
import { defineComponent, toRaw } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

const NotFound = defineComponent({
  setup() {
    const router = useRouter()
    const store = useStore()
    function GoHome() {
      const data = toRaw(store.state.sys.menus)
      if (data && data.length > 0) {
        router.replace('/')
      } else {
        router.replace('/login')
      }
    }
    return {
      // methods
      GoHome,
    }
  },
})
export default NotFound
</script>
