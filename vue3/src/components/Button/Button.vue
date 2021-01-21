<template>
  <a-button @click="CommonClick()" :type="type">
    <template v-if="iconName" #icon>
      <PlusOutlined v-if="iconName === 'add'" />
    </template>
    {{ title }}
  </a-button>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import { PlusOutlined } from '@ant-design/icons-vue'

const CommonButton = defineComponent({
  components: {
    PlusOutlined,
  },
  name: 'CommonButton',
  props: {
    title: {
      type: String,
      required: true,
      default: '默认',
    },
    iconName: {
      type: String,
    },
    type: {
      type: String,
      required: false,
      default: 'primary',
    },
  },
  setup(props, { emit }) {
    function CommonClick() {
      emit('change')
    }
    return {
      CommonClick,
    }
  },
})

export default CommonButton
</script>
