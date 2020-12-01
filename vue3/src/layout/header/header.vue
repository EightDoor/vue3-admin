<template>
  <a-layout-header :style="{ background: '#fff', padding: 0 }" />
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "common-header",
  setup() {
    return {};
  }
});
</script>
<style scoped lang="less">
@import "header.less";
</style>
