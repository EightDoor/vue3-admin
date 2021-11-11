<template>
  <a-button :type="type" @click="CommonClick()">
    <template v-if="iconName" #icon>
      <PlusOutlined v-if="iconName === 'add'" />
    </template>
    {{ title }}
  </a-button>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { PlusOutlined } from '@ant-design/icons-vue';

const CommonButton = defineComponent({
  name: 'CommonButton',
  components: {
    PlusOutlined,
  },
  props: {
    title: {
      type: String,
      required: true,
      default: '默认',
    },
    iconName: {
      type: String,
      default: '',
    },
    type: {
      type: String,
      required: false,
      default: 'primary',
    },
  },
  emits: ['change'],
  setup(props, { emit }) {
    function CommonClick() {
      emit('change');
    }
    return {
      CommonClick,
    };
  },
});

export default CommonButton;
</script>
