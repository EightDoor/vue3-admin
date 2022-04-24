<template>
  <div>
    个人博客
    <a style="margin-right: 50px" href="http://www.start6.cn">
      http://www.start6.cn
    </a>
    github:
    <a href="http://www.github.com/EightDoor">
      http://www.github.com/EightDoor
    </a>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CommonFooter',
  setup() {
    return {};
  },
});
</script>
<style scoped lang="less">
@import 'footer.less';
</style>
