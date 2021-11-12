<template>
  <router-view></router-view>
</template>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  isRouter: true,
  name: 'layoutChildren',
});
</script>
