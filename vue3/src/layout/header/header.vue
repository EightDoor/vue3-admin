<template>
  <a-layout-header class="container">
    <div class="content">
      <a-dropdown>
        <a class="ant-dropdown-link" @click="(e) => e.preventDefault()">
          <a-avatar :size="50">
            <template #icon><UserOutlined /></template>
          </a-avatar>
        </a>
        <template #overlay>
          <a-menu>
            <a-menu-item v-for="(item, index) in data" :key="index">
              <a @click="GoTo(item)">{{ item }}</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>
</template>
<script lang="ts">
import { UserOutlined } from '@ant-design/icons-vue'
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'common-header',
  components: { UserOutlined },
  setup() {
    const data = ref<string[]>(['个人中心', '退出'])
    const router = useRouter()
    function GoTo(val: string) {
      console.log(val)
      switch (val) {
        case '退出':
          localStorage.clear()
          router.replace('/login')
          break
      }
    }
    return {
      // data
      data,
      // methods
      GoTo,
    }
  },
})
</script>
<style scoped lang="less">
@import 'header.less';
</style>
