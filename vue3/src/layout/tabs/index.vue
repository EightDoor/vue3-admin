<template>
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      type="editable-card"
      @edit="OnEdit"
    >
      <a-tab-pane
        v-for="(pane, index) in panes"
        :key="index"
        :tab="pane.title"
        :closable="pane.closable"
      >
        <slot />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useStore } from 'vuex'

export default defineComponent({
  setup() {
    const store = useStore()
    const activeKey = ref(0)
    const OnEdit = (val: string) => {
      console.log(val)
    }
    return {
      activeKey,
      OnEdit,
    }
  },
})
</script>
<style lang="less" scoped>
@import 'index.less';
</style>
